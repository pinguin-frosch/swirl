package main

type SwirlConfig struct {
	Variables SwirlVariables           `json:"variables"`
	Commands  map[string][]Application `json:"commands"`
}

type SwirlVariables struct {
	Global       map[string]string   `json:"global"`
	Applications map[string]Variable `json:"applications"`
}

type Variable map[string]interface{}

func (v *Variable) GoDown(key string) (Variable, bool) {
	if next, ok := (*v)[key].(map[string]interface{}); ok {
		return next, true
	}
	return nil, false
}

func (v *Variable) GetValue(key string) (string, bool) {
	if value, ok := (*v)[key].(string); ok {
		return value, true
	}
	return "", false
}

type Application struct {
	Name     string   `json:"name"`
	Commands []string `json:"commands"`
}
