package models

import (
	"big_bang/configs"
	"fmt"
	"os"
	"strings"
)

type (
	Template struct {
		Name        string
		ResourceTxt string
	}
)

func (t *Template) MapValue(find, replace string) {
	t.ResourceTxt = strings.ReplaceAll(t.ResourceTxt, "{{{"+find+"}}}", replace)
}

func (t *Template) MapTemplate(find string, template *Template) {
	t.ResourceTxt = strings.ReplaceAll(t.ResourceTxt, "{{{"+find+"}}}", template.ResourceTxt)
}

func (t *Template) WriteResultFile() {
	if err := os.WriteFile(fmt.Sprintf("%s/%s.go", configs.RESULT_DIR, t.Name), []byte(t.ResourceTxt), 0644); err != nil {
		panic(err)
	}
}
