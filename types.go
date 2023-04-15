package main

type SwirlConfig struct {
	Variables SwirlVariables           `json:"variables"`
	Commands  map[string][]Application `json:"commands"`
}

type SwirlVariables struct {
	Global       map[string]string            `json:"global"`
	Applications map[string]map[string]string `json:"applications"`
}

type Application struct {
	Name     string   `json:"name"`
	Commands []string `json:"commands"`
}
