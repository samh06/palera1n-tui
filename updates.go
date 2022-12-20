package main

// A simple example that shows how to retrieve a value from a Bubble Tea
// program after the Bubble Tea has exited.

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/blang/semver"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
)

const version = "0.0.2"

func updates() {
    selfupdate.DisableLog()
    latest, found, err := selfupdate.DetectLatest("samh06/palera1n-tui")
    if err != nil {
        log.Println("Error occurred while detecting version:", err)
        return
    }

    v := semver.MustParse(version)
    if !found || latest.Version.LTE(v) {
        log.Println("Current version is the latest")
        return
    }

    fmt.Print("Do you want to update to ", latest.Version, "? (y/n): ")
    input, err := bufio.NewReader(os.Stdin).ReadString('\n')
    if err != nil || (input != "y\n" && input != "n\n") {
        log.Println("Invalid input")
        return
    }
    if input == "n\n" {
        return
    }

    exe, err := os.Executable()
    if err != nil {
        log.Println("Could not locate executable path")
        return
    }
    if err := selfupdate.UpdateTo(latest.AssetURL, exe); err != nil {
        log.Println("Error occurred while updating binary: ", err)
        return
    }
    log.Println("Successfully updated to version", latest.Version)
}