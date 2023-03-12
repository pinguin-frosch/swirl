package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
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

	// Define command line arguments
	var theme string
	var background string
	flag.StringVar(&theme, "theme", "", "Theme to use")
	flag.StringVar(&background, "background", "", "Background to use")

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

	// Use config values if not provided in the command line
	if theme == "" {
		theme = swirlVariables.Theme
	}
	if background == "" {
		background = swirlVariables.Background
	}

	// Update config
	swirlConfig.Variables.Theme = theme
	swirlConfig.Variables.Background = background

	// Save config after changing the theme and background
	data, err := json.MarshalIndent(swirlConfig, "", "  ")
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(configFile, data, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Changing theme to %s using %s background.\n\n", theme, background)

	for _, app := range swirlConfig.Applications {
		name := app.Name
		variables := app.Variables

		// Add swirl variables to the current app
		variables["theme"] = theme
		variables["background"] = background

		// Print app name
		fmt.Printf("%s\n", strings.ToUpper(name))

		// Command and arguments to run
		var cmdArgs []string

		// Change theme color
		for _, command := range app.ThemeCommmands[theme] {
			command = replaceVariables(command, variables)
			cmdArgs = parseCommandString(command)
			runCommand(cmdArgs)
		}

		// Change backgrund color
		for _, command := range app.BackgroundCommands[background] {
			command = replaceVariables(command, variables)
			cmdArgs = parseCommandString(command)
			runCommand(cmdArgs)
		}
	}
}

type SwirlConfig struct {
	Variables    SwirlVariables `json:"variables"`
	Applications []Application  `json:"applications"`
}

type SwirlVariables struct {
	Theme      string `json:"theme"`
	Background string `json:"background"`
}

type Application struct {
	Name               string              `json:"name"`
	Variables          map[string]string   `json:"variables"`
	BackgroundCommands map[string][]string `json:"background_commands"`
	ThemeCommmands     map[string][]string `json:"theme_commands"`
}

func runCommand(cmdArgs []string) {
	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
	cmd.Run()
}

func replaceVariables(cmdString string, variables map[string]string) string {
	// Loop over variables map and replace them in the cmdString
	for {
		replaced := false

		// Loop over every key and replace if exists in cmdString
		for key := range variables {
			placeholder := "%" + key + "%"
			if strings.Contains(cmdString, placeholder) {
				cmdString = strings.ReplaceAll(cmdString, placeholder, variables[key])
				replaced = true
			}
		}

		// Exit if there was not any replacement
		if !replaced {
			break
		}
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
	quoteChar := rune(0)
	arg := ""

	for _, c := range cmdString {
		// Keep track of " and ' to allow spaces in args
		if c == '"' || c == '\'' {
			if quote && c == quoteChar {
				quote = false
				quoteChar = rune(0)
			} else if !quote {
				quote = true
				quoteChar = c
			} else {
				// Keep everything inside the current quote
				arg += string(c)
			}
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
