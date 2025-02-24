package main

import (
	"strings"
)

func main() {
	config := LoadConfig("config.json")

	DeleteFilesInDir(RESULT_DIR)

	var templateTxt string
	ReadTemplate("handler", &templateTxt)

	TemplateMapValue(&templateTxt, "module_name", config.ModuleName)
	TemplateMapValue(&templateTxt, "lower_module_name", strings.ToLower(config.ModuleName))
	TemplateMapValue(&templateTxt, "method_prefix", ToCapFirstLetter(config.RequestMethod))
	TemplateMapValue(&templateTxt, "tag", config.Tag)
	TemplateMapValue(&templateTxt, "route_path", config.RoutePath)
	TemplateMapValue(&templateTxt, "lower_method_prefix", strings.ToLower(config.RequestMethod))

	WriteResultFile("handler", &templateTxt)
}

type Config struct {
	ModuleName    string `json:"module_name"`
	RequestMethod string `json:"request_method"`
	Tag           string `json:"tag"`
	RoutePath     string `json:"route_path"`
}
