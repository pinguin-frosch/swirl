package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
)

// Default config location
var configFile string = path.Join(os.Getenv("HOME"), ".config", "swirl", "config.json")

func main() {
	fmt.Println("========================= Swirl =========================")

	// Define command line arguments
	var theme string
	var background string
	var keyboard string
	flag.StringVar(&theme, "theme", "", "Theme to use")
	flag.StringVar(&background, "background", "", "Background to use")
	flag.StringVar(&keyboard, "keyboard", "", "Keyboard layout to use")

	// Parse command line arguments
	flag.Parse()

	// Read config file
	file, err := os.Open(configFile)

	if err != nil {
		log.Fatal(err)
		return
	}

	// Close file when exiting
	defer file.Close()

	byteResult, err := io.ReadAll(file)

	if err != nil {
		log.Fatal(err)
	}

	// Create a struct to hold the json
	var swirlConfig SwirlConfig

	// Marshal json to a struct
	json.Unmarshal(byteResult, &swirlConfig)

	// Get the variables from the config
	swirlVariables := swirlConfig.Variables

	// Toggle dark and light theme if no commands are provided
	if flag.NFlag() == 0 {
		if swirlVariables.Background == "dark" {
			background = "light"
		} else {
			background = "dark"
		}
	}

	// Use config values if not provided in the command line
	if theme == "" {
		theme = swirlVariables.Theme
	}
	if background == "" {
		background = swirlVariables.Background
	}
	if keyboard == "" {
		keyboard = swirlVariables.Keyboard
	}

	// Update config
	swirlConfig.Variables.Theme = theme
	swirlConfig.Variables.Background = background
	swirlConfig.Variables.Keyboard = keyboard

	// Save config after changing the theme and background
	data, err := json.MarshalIndent(swirlConfig, "", "  ")
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(configFile, data, 0644)
	if err != nil {
		panic(err)
	}

	// Run commands
	runAppCommands(swirlConfig.Theme, &swirlConfig.Variables, "Changing theme...")
	runAppCommands(swirlConfig.Background, &swirlConfig.Variables, "Changing background...")
	runAppCommands(swirlConfig.Keyboard, &swirlConfig.Variables, "Changing keyboard layout...")
}
