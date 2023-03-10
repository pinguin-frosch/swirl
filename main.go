package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path"
)

// Default config location
var configFile string = path.Join(os.Getenv("HOME"), ".config", "swirl", "config.json")

func main() {
	fmt.Println("===== Swirl =====")

	// Read config file
	file, err := os.Open(configFile)

	if err != nil {
		log.Fatal(err)
		return
	}

	defer file.Close()

	byteResult, err := io.ReadAll(file)

	if err != nil {
		log.Fatal(err)
	}

	var res map[string]interface{}
	json.Unmarshal([]byte(byteResult), &res)

	fmt.Println(res)
}
