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

func ReadTemplate(templateName string, result *string) {
	templateData, err := os.ReadFile(fmt.Sprintf("%s/%s.txt", TEMPLATE_DIR, templateName))
	if err != nil {
		panic(err)
	}

	*result = string(templateData)
}

func TemplateMapValue(sourceTxt *string, find string, replace string) {
	if sourceTxt == nil {
		panic("")
	}

	*sourceTxt = strings.ReplaceAll(*sourceTxt, "{{{"+find+"}}}", replace)
}

func WriteResultFile(resultPath string, resultTxt *string) {
	if resultTxt == nil {
		panic("")
	}

	if err := os.WriteFile(fmt.Sprintf("%s/%s.go", RESULT_DIR, resultPath), []byte(*resultTxt), 0644); err != nil {
		panic(err)
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
