package main

import (
	"encoding/json"
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

	// Parse command line arguments
	args := parseCommandArgs()

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
	globalVariables := &swirlConfig.Variables.Global

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
		panic(err)
	}
	err = ioutil.WriteFile(configFile, data, 0644)
	if err != nil {
		panic(err)
	}

	// Run commands
	for k := range args {
		runAppCommands(swirlConfig.Commands[k], &swirlConfig.Variables, fmt.Sprintf("Changing %s", k))
	}
}
