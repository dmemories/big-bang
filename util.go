package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	TEMPLATE_DIR = "templates"
	RESULT_DIR   = "results"
)

func LoadConfig(filename string) *Config {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		panic(err)
	}

	config.RequestMethod = ToCapFirstLetter(config.RequestMethod)
	switch config.RequestMethod {
	case "Get", "Post", "Put", "Delete":
		break
	default:
		panic("invalid method format")
	}

	return &config
}

func ReadTemplate(templateName string, mapper map[string]string, result *Template) {
	templateData, err := os.ReadFile(fmt.Sprintf("%s/%s.txt", TEMPLATE_DIR, templateName))
	if err != nil {
		panic(err)
	}

	result.ResourceTxt = string(templateData)
	for k, v := range mapper {
		result.MapValue(k, v)
	}
}

func DeleteFilesInDir(dirPath string) {
	files, err := os.ReadDir(dirPath)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		filePath := filepath.Join(dirPath, file.Name())

		if !file.IsDir() {
			err := os.Remove(filePath)
			if err != nil {
				panic(err)
			}
		}
	}
}

func ToCapFirstLetter(input string) string {
	if len(input) == 0 {
		return input
	}
	return strings.ToUpper(input[:1]) + strings.ToLower(input[1:])
}

func ToCapFirstLower(input string) string {
	if len(input) > 0 {
		input = strings.ToLower(string(input[0])) + input[1:]
	}

	return input
}
