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

		// Add global variables to the current app
		for k, v := range swirlVariables.Global {
			variables[k] = v
		}

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

// Finds the chuncks in the cmd that have valid dot variables.
func findDotVariables(cmdString string) []string {
	variables := make([]string, 0)

	variable := ""
	inPercent := false
	for _, c := range cmdString {
		if c == '%' {
			variable += string(c)
			if inPercent {
				if isValidVariable(variable) {
					variables = append(variables, variable)
				}
				variable = ""
			}
			inPercent = !inPercent
		} else if isValidKey(c) && inPercent {
			variable += string(c)
		} else {
			inPercent = false
			variable = ""
		}
	}

	return variables
}

func isValidVariable(variable string) bool {
	length := len(variable)
	if length < 2 || !strings.Contains(variable, ".") || variable[1] == '.' || variable[length-2] == '.' {
		return false
	}
	for i := 1; i < len(variable); i++ {
		if variable[i] == '.' && variable[i-1] == '.' {
			return false
		}
	}
	return true
}

func isValidKey(c rune) bool {
	return c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' || c == '_' || c == '.'
}

func replaceVariables(cmdString string, variables map[string]interface{}) string {
	// Loop over variables map and replace them in the cmdString
	for {
		replaced := false

		// Loop over every key and replace if exists in cmdString
		for key := range variables {
			placeholder := "%" + key + "%"
			if strings.Contains(cmdString, placeholder) {
				variable, ok := variables[key].(string)
				if !ok {
					continue
				}
				cmdString = strings.ReplaceAll(cmdString, placeholder, variable)
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

func replaceDotVariables(path string, variables Variable) (string, error) {
	// Chop prefixes until there are no more dots
	current := variables
	for strings.Contains(path, ".") {
		// Get current iteration prefix
		prefix := path[:strings.Index(path, ".")]
		prefixDot := fmt.Sprintf("%s.", prefix)

		// Descend to the next level using the prefix
		next, ok := current.GoDown(prefix)
		if !ok {
			return "", fmt.Errorf("Couldn't go down to %v", prefix)
		}

		// Update the variable and remove the prefix from the path
		path = strings.TrimPrefix(path, prefixDot)
		current = next
	}

	if path == "" {
		return "", fmt.Errorf("Invalid dot at the end of string")
	}

	// Get the final final
	value, ok := current.GetValue(path)
	if !ok {
		fmt.Printf("Got an error while getting the value\n")
		return "", fmt.Errorf("Couldn't get value %v", path)
	}

	return value, nil
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
