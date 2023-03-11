package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
)

// Default config location
var configFile string = path.Join(os.Getenv("HOME"), ".config", "swirl", "config.json")

func main() {
	fmt.Println("========================= Swirl =========================")

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

	theme := swirlConfig.Theme
	background := swirlConfig.Background

	fmt.Printf("Changing theme to %s using %s background.\n\n", theme, background)

	for _, app := range swirlConfig.Applications {
		name := app.Name
		variables := app.Variables

		// Print app name
		fmt.Printf("%s\n", strings.ToUpper(name))

		// Command and arguments to run
		var cmdArgs []string

		// Print commands for the background
		for _, command := range app.BackgroundCommands[background] {
			command = replaceVariables(command, variables)
			cmdArgs = parseCommandString(command)
			runCommand(cmdArgs)
		}
	}
}

type SwirlConfig struct {
	Theme        string        `json:"theme"`
	Background   string        `json:"background"`
	Applications []Application `json:"applications"`
}

type Application struct {
	Name               string              `json:"name"`
	Variables          map[string]string   `json:"variables"`
	BackgroundCommands map[string][]string `json:"background_commands"`
}

func runCommand(cmdArgs []string) {
	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
	cmd.Run()
}

func replaceVariables(cmdString string, variables map[string]string) string {
	// Loop over variables map and replace them in the cmdString
	for key := range variables {
		cmdString = strings.ReplaceAll(cmdString, "%"+key+"%", variables[key])
	}

	// Change ~ for the actual home directory
	homeDir, _ := os.UserHomeDir()
	cmdString = strings.ReplaceAll(cmdString, "~", homeDir)

	return cmdString
}

func parseCommandString(cmdString string) []string {
	// Array of string slices to hold the command and args
	cmdArgs := []string{}
	quote := false
	arg := ""

	for _, c := range cmdString {
		// Keep track of " to allow spaces in args
		if c == '"' {
			quote = !quote
		} else if c == ' ' && !quote {
			// Append to the cmd args and reset the arg
			if arg != "" {
				cmdArgs = append(cmdArgs, arg)
			}
			arg = ""
		} else {
			// Keep adding to arg until the above conditions met
			arg += string(c)
		}
	}

	if arg != "" {
		// Add the remaining arg is exists
		cmdArgs = append(cmdArgs, arg)
	}

	return cmdArgs
}
