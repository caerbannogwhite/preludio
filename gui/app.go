package main

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
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

func (a *App) RunCommand(cmd string) string {
	fmt.Printf("Run command '%s'\n", cmd)

	return ""
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
