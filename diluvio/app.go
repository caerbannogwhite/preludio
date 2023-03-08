package main

import (
	"bytes"
	"compiler"
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"preludio"
)

// App struct
type App struct {
	ctx        context.Context
	preludioVm *preludio.ByteEater
}

// NewApp creates a new App application struct
func NewApp() *App {

	be := new(preludio.ByteEater).
		InitVM()
		// SetPrintWarning(args.Warnings).
		// SetDebugLevel(args.DebugLevel)

	return &App{preludioVm: be.InitVM()}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func IsValidPath(filePaht string) bool {
	// Check if file already exists
	if _, err := os.Stat(filePaht); err == nil {
		return true
	}

	// Attempt to create it
	var d []byte
	if err := ioutil.WriteFile(filePaht, d, 0644); err == nil {
		os.Remove(filePaht) // And delete it
		return true
	}

	return false
}

func (a *App) LookUpPath(path string) []string {

	info, err := os.Stat(path)
	if err == nil {
		if info.IsDir() {
			list, err := os.ReadDir(path)
			if err != nil {
				return []string{}
			}

			res := make([]string, len(list))
			for i, entry := range list {
				res[i] = entry.Name()
			}

			return res
		}
	}

	return []string{}
}

func (a *App) RunPreludioCode(code string) string {

	bytecode := compiler.CompileSource(code)
	out := a.preludioVm.RunBytecode(bytecode)

	res, err := json.Marshal(out)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(string(res))
	return string(res)
}

func (a *App) ParseCsv(blob string) string {

	file := []byte{}
	if err := json.Unmarshal([]byte(blob), &file); err != nil {
		log.Fatal(err)
	}

	reader := csv.NewReader(bytes.NewReader(file))

	tab, err := reader.ReadAll()
	if err != nil {
		res, _ := json.Marshal(err)
		return string(res)
	}

	res, _ := json.Marshal(tab)
	return string(res)
}

func (a *App) ImportCsv(path string) []byte {
	f, err := os.Open(path)
	if err != nil {
		res, _ := json.Marshal(err)
		return res
	}

	reader := csv.NewReader(f)

	tab, err := reader.ReadAll()
	if err != nil {
		res, _ := json.Marshal(err)
		return res
	}

	res, _ := json.Marshal(tab)
	return res
}

func (a *App) ImportExcel(path string) []byte {

	_, err := os.Open(path)
	if err != nil {
		res, _ := json.Marshal(err)
		return res
	}

	return []byte{}
}
