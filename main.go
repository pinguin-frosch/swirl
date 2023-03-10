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

	var swirlConfig SwirlConfig

	// Marshal json to a struct
	json.Unmarshal(byteResult, &swirlConfig)

	theme := swirlConfig.Theme
	background := swirlConfig.Background

	fmt.Printf("Changing theme to %s using %s background.\n", theme, background)

	for _, app := range swirlConfig.Applications {
		// Print app name
		fmt.Printf("\n%s\n", strings.ToUpper(app.Name))

		// Change background color
		background_commands := app.Background[background]

		for _, command := range background_commands {
			fmt.Println(command)

			// Form command with executable and the arguments
			cmd := exec.Command(command[0], command...)
			err := cmd.Run()

			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

type SwirlConfig struct {
	Theme        string        `json:"theme"`
	Background   string        `json:"background"`
	Applications []Application `json:"applications"`
}

type Application struct {
	Name       string                 `json:"name"`
	Variables  map[string]interface{} `json:"variables"`
	Background map[string][][]string  `json:"background"`
	Theme      map[string][][]string  `json:"theme"`
}
