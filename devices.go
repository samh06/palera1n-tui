package main

// A simple example that shows how to retrieve a value from a Bubble Tea
// program after the Bubble Tea has exited.

import (
	"fmt"
	"os"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func devices() {
	writeFile("./tmp1", "\nasd2", false)
}

func writeFile(file string, text string, overwrite bool) {
	// If not overwriting, most of the time you will want a /n at the start of your string
	if overwrite {
		data := []byte(text)
    	err := os.WriteFile(file, data, 0644)
    	check(err)
	} else {
		readData, readError := os.ReadFile(file)
    	check(readError)
    	fmt.Print(string(readData))
		data := []byte(string(readData)+text)
    	writeError := os.WriteFile(file, data, 0644)
    	check(writeError)
	}
}