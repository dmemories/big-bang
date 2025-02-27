package main

import (
	"fmt"
	"os"
	"strings"
)

type (
	Config struct {
		ModuleName    string `json:"module_name"`
		RequestMethod string `json:"request_method"`
		MethodName    string `json:"method_name"`
		Tag           string `json:"tag"`
		RoutePath     string `json:"route_path"`
		DtoModel      string `json:"dto_model"`
		Repository    string `json:"repository"`
	}

	Template struct {
		ResourceTxt string
	}
)

func (t *Template) MapValue(find, replace string) {
	t.ResourceTxt = strings.ReplaceAll(t.ResourceTxt, "{{{"+find+"}}}", replace)
}

func (t *Template) MapTemplate(find string, template *Template) {
	t.ResourceTxt = strings.ReplaceAll(t.ResourceTxt, "{{{"+find+"}}}", template.ResourceTxt)
}

func (t *Template) WriteResultFile(resultPath string) {
	if err := os.WriteFile(fmt.Sprintf("%s/%s.go", RESULT_DIR, resultPath), []byte(t.ResourceTxt), 0644); err != nil {
		panic(err)
	}
}
