package tools

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type TemplateData struct {
	Filename string      `json:"filename" yaml:"filename"`
	Data     interface{} `json:"data" yaml:"data"`
}

func Unmarshal(path string) ([]*TemplateData, error) {
	if path == "" {
		return nil, fmt.Errorf("data wasn't specified")
	}

	dataBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var data []*TemplateData
	switch filepath.Ext(path) {
	case ".json":
		err = json.Unmarshal(dataBytes, &data)
	case ".yaml", ".yml":
		err = yaml.Unmarshal(dataBytes, &data)
	}
	if err != nil {
		return nil, err
	}

	return data, nil
}
