package config

import (
	"encoding/json"
	"os"
)

// File is optional batch configuration loaded from JSON.
type File struct {
	MinDate    string   `json:"min_date"`
	MaxDate    string   `json:"max_date"`
	Categories []string `json:"categories"`
}

func Load(path string) (File, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return File{}, err
	}
	var f File
	return f, json.Unmarshal(data, &f)
}
