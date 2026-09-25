package main

import (
	"bufio"
	"bytefeeder"
	"fmt"
	"os"
	"preludiocore"
	"strconv"
	"strings"

	"github.com/alexflint/go-arg"
	"github.com/caerbannogwhite/enchanter/dataframe"
	"github.com/caerbannogwhite/enchanter/meta"
	"github.com/charmbracelet/lipgloss"
)

const VERSION = "0.6.0"

const DEFAULT_PROMPT = ">>> "
const DEFAULT_INDENTAION = "    "
const DEFAULT_SUSPENSION_STRING = "... "
const DEFAULT_TABLE_WIDTH = 150

var (
	STYLE_BOLD    = lipgloss.NewStyle().Bold(true)
	STYLE_PROMPT  = STYLE_BOLD.Copy().Foreground(lipgloss.Color("#F87217"))
	STYLE_NUMERIC = lipgloss.NewStyle().Foreground(lipgloss.Color("#00BFFF"))
)

type CliArgs struct {
	SourceCode string `arg:"-s, --source" help:"source code to execute" default:""`
	InputPath  string `arg:"-i, --input" help:"source file input path" default:""`
	DebugLevel int    `arg:"-d, --debug-level" help:"debug level" default:"0"`
	SdtOut     bool   `arg:"-o, --stdout" help:"print output to stdout" default:"false"`
	Verbose    bool   `arg:"-v, --verbose" help:"verbosity level" default:"false"`
	Warnings   bool   `arg:"-w, --warnings" help:"print warnings" defaut:"true"`
}

func main() {

	var args CliArgs

	arg.MustParse(&args)

	if args.InputPath != "" {
		be := new(preludiocore.ByteEater).
			InitVM().
			SetParamPrintWarning(args.Warnings).
			SetParamDebugLevel(args.DebugLevel).
			SetParamVerbose(args.Verbose).
			SetParamPrintToStdout(args.SdtOut)

		bytecode, logs, err := bytefeeder.CompileFile(args.InputPath)
		if err != nil {
			fmt.Println("Error compiling file:", err)
			os.Exit(1)
		}

		for _, log := range logs {
			fmt.Println(log)
		}

		be.RunBytecode(bytecode)

	} else if args.SourceCode != "" {
		be := new(preludiocore.ByteEater).
			InitVM().
			SetParamPrintWarning(args.Warnings).
			SetParamDebugLevel(args.DebugLevel).
			SetParamVerbose(args.Verbose).
			SetParamPrintToStdout(args.SdtOut)

		bytecode, logs, err := bytefeeder.CompileSource(args.SourceCode)
		if err != nil {
			fmt.Println("Error compiling file:", err)
			os.Exit(1)
		}

		for _, log := range logs {
			fmt.Println(log)
		}

		be.RunBytecode(bytecode)
	} else {
		LaunchRepl(args)
	}
}

func LaunchRepl(args CliArgs) {

	tableWidth := DEFAULT_TABLE_WIDTH

	fmt.Println("Welcome to the " + STYLE_PROMPT.Copy().Italic(true).Render("Preludio REPL") + "!")
	fmt.Println("Version:", STYLE_NUMERIC.Render(VERSION))

	be := new(preludiocore.ByteEater).
		InitVM().
		SetParamPrintWarning(args.Warnings).
		SetParamFullOutput(false).
		SetParamDebugLevel(args.DebugLevel).
		SetParamVerbose(args.Verbose)

	if args.Verbose {
		fmt.Printf("\nPreludio VM initialized\n")
		fmt.Printf("%15s %t\n", "Print warnings:", be.GetParamPrintWarning())
		fmt.Printf("%15s %t\n", "Full output:", be.GetParamFullOutput())
		fmt.Printf("%15s %d\n", "Debug level:", be.GetParamDebugLevel())
		fmt.Printf("%15s %t\n", "Verbose:", be.GetParamVerbose())
	}

	in := bufio.NewReader(os.Stdin)

	readerStart := true
	code := ""
	for {
		if readerStart {
			fmt.Print(STYLE_PROMPT.Render(DEFAULT_PROMPT))
		} else {
			fmt.Print(DEFAULT_SUSPENSION_STRING)
		}

		line, err := in.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				fmt.Println("Bye!")
				os.Exit(0)
			}
			fmt.Println("Error reading input:", err)
		}

		// look for magic commands
		// ignore the rest
		if strings.HasPrefix(line, "%") {
			spt := strings.Split(strings.Trim(line, "\t\n\r "), " ")
			switch spt[0] {
			case "%setenv":
				if len(spt) != 3 {
					fmt.Printf("Usage: %%setenv <key> <value>\n")
					continue
				}

				switch spt[1] {
				case "ENV_WARNINGS":
					if spt[2] == "true" {
						be.SetParamPrintWarning(true)
					} else if spt[2] == "false" {
						be.SetParamPrintWarning(false)
					}
					fmt.Printf("Print warnings set to \"%t\"", be.GetParamPrintWarning())
				case "ENV_DEBUG_LEVEL":
					l, err := strconv.Atoi(spt[2])
					if err != nil {
						fmt.Println("Error parsing debug level:", err)
						continue
					}
					be.SetParamDebugLevel(l)
					fmt.Println("Debug level set to", be.GetParamDebugLevel())
				case "ENV_VERBOSE":
					if spt[2] == "true" {
						be.SetParamVerbose(true)
					} else if spt[2] == "false" {
						be.SetParamVerbose(false)
					}
					fmt.Printf("Verbose set to \"%t\"", be.GetParamVerbose())
				case "ENV_FULL_OUTPUT":
					if spt[2] == "true" {
						be.SetParamFullOutput(true)
					} else if spt[2] == "false" {
						be.SetParamFullOutput(false)
					}
					fmt.Printf("Full output set to \"%t\"", be.GetParamFullOutput())
				case "ENV_OUTPUT_WIDTH":
					l, err := strconv.Atoi(spt[2])
					if err != nil {
						fmt.Println("Error parsing output width:", err)
						continue
					}
					tableWidth = l
					fmt.Println("Output width set to", tableWidth)
				default:
					fmt.Println("Unknown environment variable:", spt[1])
				}

			// case "%getenv":
			// 	if len(spt) != 2 {
			// 		fmt.Println("Usage: %getenv <key>")
			// 		continue
			// 	}
			// 	fmt.Println(be.GetEnv(spt[1]))

			default:
				fmt.Println("Unknown magic command:", spt[0])
			}
			continue
		}

		line = strings.TrimSpace(line)

		if line == "" {
			res := be.RunSource(code)
			for _, log := range res.Log {
				switch log.LogType {
				case meta.LOG_DEBUG:
					if int(log.Level) < be.GetParamDebugLevel() {
						fmt.Println("[🐛] " + log.Message)
					}

				case meta.LOG_INFO:
					fmt.Println("[ ℹ️ ] " + log.Message)

				case meta.LOG_WARNING:
					fmt.Println("[⚠️] " + log.Message)

				case meta.LOG_ERROR:
					fmt.Println("[❌] " + log.Message)
				}
			}

			params := dataframe.NewPPrintParams().
				SetIndent(DEFAULT_INDENTAION).
				SetUseLipGloss(true).
				SetWidth(tableWidth).
				SetNRows(be.GetParamOutputSnippetLength())

			for _, df := range res.Data {
				p := params
				if be.GetParamFullOutput() {
					p = p.SetNRows(df.NRows())
				}

				out := df.Table(p)
				if !strings.HasSuffix(out, "\n") {
					out += "\n"
				}
				fmt.Print(out)
			}

			code = ""
			readerStart = true
		} else {
			readerStart = false
		}

		code += line + "\n"
	}
}
