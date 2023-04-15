package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func runCommand(cmdArgs []string) {
	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
	cmd.Run()
}

func runAppCommands(apps []Application, swirlVariables *SwirlVariables, message string) {
	if len(apps) == 0 {
		return
	}

	fmt.Printf("\n%s\n", message)

	for _, app := range apps {
		name := app.Name
		variables := swirlVariables.Applications[name]

		// Add swirl variables to the current app
		variables["theme"] = (*swirlVariables).Theme
		variables["background"] = (*swirlVariables).Background
		variables["keyboard"] = (*swirlVariables).Keyboard
		variables["taskbar"] = (*swirlVariables).Taskbar

		// Print app name
		fmt.Printf("%s\n", strings.ToLower(name))

		// Command and arguments to run
		var cmdArgs []string

		// Parse commands and run them
		for _, command := range app.Commands {
			command = replaceVariables(command, variables)
			cmdArgs = parseCommandString(command)
			runCommand(cmdArgs)
		}
	}
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

func parseCommandArgs() map[string]string {
	rawArgs := os.Args[1:]
	parsedArgs := make(map[string]string)

	for i, arg := range rawArgs {
		if strings.HasPrefix(arg, "-") && i < len(rawArgs)-1 {
			if !strings.HasPrefix(rawArgs[i+1], "-") {
				parsedArgs[strings.TrimPrefix(arg, "-")] = rawArgs[i+1]
			}
		}
	}

	return parsedArgs
}
