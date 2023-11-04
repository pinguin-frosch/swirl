package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestReplaceDots(t *testing.T) {
	weztermConfig := `
    {
        "fonts": {
            "agave": "Agave Nerd Font",
            "cascadia": "CaskaydiaCove NF SemiLight",
            "fantasque": "FantasqueSansM Nerd Font",
            "hack": "Hack Nerd Font",
            "inconsolata": "Inconsolata Nerd Font",
            "iosevka": "Iosevka Nerd Font",
            "jetbrains": "JetBrainsMono Nerd Font",
            "monoid": "Monoid Nerd Font"
        },
        "path": "~/.config/wezterm/wezterm.lua"
    }
    `
	var variables Variable
	err := json.Unmarshal([]byte(weztermConfig), &variables)
	if err != nil {
		t.Fatalf("Got error: %v\n", err)
	}

	fonts := map[string]string{
		"agave":       "Agave Nerd Font",
		"cascadia":    "CaskaydiaCove NF SemiLight",
		"fantasque":   "FantasqueSansM Nerd Font",
		"hack":        "Hack Nerd Font",
		"inconsolata": "Inconsolata Nerd Font",
		"iosevka":     "Iosevka Nerd Font",
		"jetbrains":   "JetBrainsMono Nerd Font",
		"monoid":      "Monoid Nerd Font",
	}

	for font, fontName := range fonts {
		path := fmt.Sprintf("fonts.%s", font)
		replaced, err := replaceDotVariables(path, variables)
		if err != nil {
			t.Fatalf("Got an error: %v", err)
		}
		if replaced != fontName {
			t.Fatalf(`replaceDotVariables(%v, variables) should be "%v", but got "%v"`, path, fontName, replaced)
		}
	}

}

func TestFindDotVariables(t *testing.T) {
	weztermConfig := `
	{
		"fonts": {
			"iosevka": "Iosevka Nerd Font"
		},
		"very": {
			"long": {
				"path": {
					"to": {
						"annoy": "you"
					}
				}
			}
		},
		"font": "iosevka",
		"path": "~/.config/wezterm/wezterm.lua"
	}
	`
	var variables Variable
	err := json.Unmarshal([]byte(weztermConfig), &variables)
	if err != nil {
		t.Fatalf("Got error: %v\n", err)
	}

	pattern := "sed %a..b% -i \"s|\\(wezterm\\.font('\\).*')|\\1%fonts.%font%%')|\" %__.% %2 %%%3 %path% %very.long.path.to.annoy%"
	pattern = replaceVariables(pattern, variables)
	toReplace := findDotVariables(pattern)

	if len(toReplace) != 2 {
		t.Fatalf("Expected 2 dotVariables, but got: %v", len(toReplace))
	}
	expected := []string{"%fonts.iosevka%", "%very.long.path.to.annoy%"}
	for k, v := range toReplace {
		if expected[k] != v {
			t.Fatalf("Expected %v, got: %v", expected[k], v)
		}
	}
}
