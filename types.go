package main

type SwirlConfig struct {
	Variables  SwirlVariables `json:"variables"`
	Background []Application  `json:"background"`
	Theme      []Application  `json:"theme"`
	Keyboard   []Application  `json:"keyboard"`
	Taskbar    []Application  `json:"taskbar"`
}

type SwirlVariables struct {
	Theme        string                       `json:"theme"`
	Background   string                       `json:"background"`
	Keyboard     string                       `json:"keyboard"`
	Taskbar      string                       `json:"taskbar"`
	Applications map[string]map[string]string `json:"applications"`
}

type Application struct {
	Name     string   `json:"name"`
	Commands []string `json:"commands"`
}
