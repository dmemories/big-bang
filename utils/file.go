package utils

import (
	"big_bang/configs"
	"big_bang/models"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func LoadConfig(filename string) *models.Config {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var config models.Config
	if err := json.Unmarshal(data, &config); err != nil {
		panic(err)
	}

	return &config
}

func ReadTemplateWithMapper(templateName string, mappers ...map[string]string) (result string) {
	templateData, err := os.ReadFile(fmt.Sprintf("%s/%s.txt", configs.TEMPLATE_DIR, templateName))
	if err != nil {
		panic(err)
	}

	result = string(templateData)
	for _, mapper := range mappers {
		for k, v := range mapper {
			result = strings.ReplaceAll(result, "{{{"+k+"}}}", v)
		}
	}

	return result
}

func ReadMainTemplateFile(templateName string, mappers ...map[string]string) (result *models.Template) {
	templateData, err := os.ReadFile(fmt.Sprintf("%s/%s.txt", configs.TEMPLATE_DIR, templateName))
	if err != nil {
		panic(err)
	}

	result = &models.Template{
		Name:        templateName,
		ResourceTxt: string(templateData),
	}
	for _, mapper := range mappers {
		for k, v := range mapper {
			result.MapValue(k, v)
		}
	}

	return result
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
