package main

// A simple example that shows how to retrieve a value from a Bubble Tea
// program after the Bubble Tea has exited.

import (
	"os"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func devices() {
	parseDevice("1", "asd", "--tweaks 15.4 --semi-tethered")
}

func writeFile(file string, text string) {
	// If not overwriting, most of the time you will want a /n at the start of your string
	readData, readError := os.ReadFile(file)
	if readError != nil {
		os.Create(file)
		readData, readError = os.ReadFile(file)
		data := []byte(string(readData)+text)
		writeError := os.WriteFile(file, data, 0644)
		check(writeError)
	} else {
		data := []byte(string(readData)+"\n"+text)
		writeError := os.WriteFile(file, data, 0644)
		check(writeError)
	}
}

func parseDevice(id string, nickname string, commands string) {
	sep := ";"
	Device := id+sep+nickname+sep+commands
	writeFile("./tmp1", Device)
}