package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
)

// Default config location
var configFile string = path.Join(os.Getenv("HOME"), ".config", "swirl", "config.json")

func main() {
	fmt.Println("========================= Swirl =========================")

	// Parse command line arguments
	args := parseCommandArgs()

	// Read config file
	file, err := os.ReadFile(configFile)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Create a struct to hold the json
	var swirlConfig SwirlConfig

	// Marshal json to a struct
	json.Unmarshal(file, &swirlConfig)

	// Get the variables from the config
	globalVariables := &swirlConfig.Variables.Global

	// Invert background when no args are provided
	if len(args) == 0 {
		if background, ok := (*globalVariables)["background"]; ok {
			if background == "dark" {
				args["background"] = "light"
			} else {
				args["background"] = "dark"
			}
		}
	}

	// Update config
	for k, v := range args {
		if _, ok := (*globalVariables)[k]; ok {
			(*globalVariables)[k] = v
		} else {
			delete(args, k)
		}
	}

	// Save config after updating it
	data, err := json.MarshalIndent(swirlConfig, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile(configFile, data, 0644)
	if err != nil {
		log.Fatal(err)
	}

	// Run commands
	for k := range args {
		runAppCommands(swirlConfig.Commands[k], &swirlConfig.Variables, fmt.Sprintf("Changing %s", k))
	}
}
