package main

import (
	"bufio"
	"bytefeeder"
	"fmt"
	"os"
	"preludiocore"
	"preludiometa"
	"strconv"
	"strings"

	"github.com/alexflint/go-arg"
	"github.com/charmbracelet/lipgloss"
)

const VERSION = "0.4.0"

const DEFAULT_PROMPT = ">>> "
const DEFAULT_INDENTAION = "    "
const DEFAULT_SUSPENSION_STRING = "... "
const DEFAULT_NULL_STRING = preludiometa.SYMBOL_NA
const DEFAULT_OUTPUT_COLUMN_SIZE = 12

var JUST_RIGHT_TYPES = map[string]bool{
	"Bool":    true,
	"Int64":   true,
	"Float64": true,
	"String":  false,
}

var (
	STYLE_BOLD    = lipgloss.NewStyle().Bold(true)
	STYLE_ITALIC  = lipgloss.NewStyle().Italic(true)
	STYLE_PROMPT  = STYLE_BOLD.Copy().Foreground(lipgloss.Color("#F87217"))
	STYLE_NA      = STYLE_BOLD.Copy().Foreground(lipgloss.Color("#C04000"))
	STYLE_BOOL    = lipgloss.NewStyle().Foreground(lipgloss.Color("#00BFFF"))
	STYLE_NUMERIC = lipgloss.NewStyle().Foreground(lipgloss.Color("#00BFFF"))
	STYLE_STRING  = STYLE_ITALIC.Copy().Foreground(lipgloss.Color("#98AFC7"))
	NO_STYLE      = lipgloss.NewStyle()
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

	outputColumnSize := DEFAULT_OUTPUT_COLUMN_SIZE

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
				case "ENV_OUTPUT_COLUMN_SIZE":
					l, err := strconv.Atoi(spt[2])
					if err != nil {
						fmt.Println("Error parsing output column size:", err)
						continue
					}
					outputColumnSize = l
					fmt.Println("Output column size set to", outputColumnSize)
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
				case preludiometa.LOG_DEBUG:
					if int(log.Level) < be.GetParamDebugLevel() {
						fmt.Println("[🐛] " + log.Message)
					}

				case preludiometa.LOG_INFO:
					fmt.Println("[ ℹ️ ] " + log.Message)

				case preludiometa.LOG_WARNING:
					fmt.Println("[⚠️] " + log.Message)

				case preludiometa.LOG_ERROR:
					fmt.Println("[❌] " + log.Message)
				}
			}

			for _, c := range res.Data {
				prettyPrint(DEFAULT_INDENTAION, outputColumnSize, c)
			}

			code = ""
			readerStart = true
		} else {
			readerStart = false
		}

		code += line + "\n"
	}
}

func truncate(s string, n int) string {
	if len(s) > n {
		return s[:n-3] + "..."
	}
	return s
}

func prettyPrint(indent string, colSize int, columnar []preludiometa.Columnar) {
	if len(columnar) == 0 {
		return
	}

	actualColSize := colSize + 3
	fmtStringLeft := fmt.Sprintf(" %%-%ds ", colSize)
	fmtStringRight := fmt.Sprintf(" %%%ds ", colSize)

	// header
	buffer := indent + "╭"
	for i := 1; i < len(columnar)*actualColSize; i++ {
		if i%actualColSize == 0 {
			buffer += "┬"
		} else {
			buffer += "─"
		}
	}
	buffer += "╮\n"

	// column names
	// check if there are any column names
	colNames := false
	for _, c := range columnar {
		if c.Name != "" {
			colNames = true
			break
		}
	}

	// only print column names if there are any
	if colNames {
		buffer += indent
		for _, c := range columnar {
			buffer += "│" + STYLE_BOLD.Render(fmt.Sprintf(fmtStringLeft, truncate(c.Name, colSize)))
		}
		buffer += "│\n"

		// separator
		buffer += indent + "├"
		for i := 1; i < len(columnar)*actualColSize; i++ {
			if i%actualColSize == 0 {
				buffer += "┼"
			} else {
				buffer += "─"
			}
		}
		buffer += "┤\n"
	}

	// column preludiometa
	buffer += indent
	for _, c := range columnar {
		buffer += "│" + STYLE_BOLD.Copy().
			Italic(true).
			Render(fmt.Sprintf(fmtStringLeft, truncate(c.Type, colSize)))
	}
	buffer += "│\n"

	// separator
	buffer += indent + "├"
	for i := 1; i < len(columnar)*actualColSize; i++ {
		if i%actualColSize == 0 {
			buffer += "┼"
		} else {
			buffer += "─"
		}
	}
	buffer += "┤\n"

	// data
	for i := 0; i < len(columnar[0].Data); i++ {
		buffer += indent
		for _, c := range columnar {
			fmtString := fmtStringLeft
			if JUST_RIGHT_TYPES[c.Type] {
				fmtString = fmtStringRight
			}

			if c.Nulls[i] {
				buffer += "│" + STYLE_NA.Render(fmt.Sprintf(fmtString, DEFAULT_NULL_STRING))
			} else {
				switch c.Type {
				case "Bool":
					buffer += "│" + STYLE_BOOL.Render(fmt.Sprintf(fmtString, c.Data[i]))
				case "Int64", "Float64":
					buffer += "│" + STYLE_NUMERIC.Render(fmt.Sprintf(fmtString, c.Data[i]))
				case "String":
					buffer += "│" + STYLE_STRING.Render(fmt.Sprintf(fmtString, truncate(c.Data[i], colSize)))
				default:
					buffer += "│" + STYLE_STRING.Render(fmt.Sprintf(fmtString, truncate(c.Data[i], colSize)))
				}
			}
		}
		buffer += "│\n"
	}

	if len(columnar[0].Data) < columnar[0].ActualLength {
		buffer += indent
		for _, c := range columnar {
			fmtString := fmtStringLeft
			if JUST_RIGHT_TYPES[c.Type] {
				fmtString = fmtStringRight
			}
			buffer += "│" + STYLE_STRING.Render(fmt.Sprintf(fmtString, "..."))
		}
		buffer += "│\n"
	}

	// separator
	buffer += indent + "╰"
	for i := 1; i < len(columnar)*actualColSize; i++ {
		if i%actualColSize == 0 {
			buffer += "┴"
		} else {
			buffer += "─"
		}
	}
	buffer += "╯\n"

	fmt.Print(buffer)
}
