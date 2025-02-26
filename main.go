package main

import "strings"

var (
	mainTemplate        Template
	swagTemplate        Template
	usecaseTemplate     Template
	containerTemplate   Template
	baseHandlerTemplate Template
	routerTemplate      Template
)

func main() {
	config := LoadConfig("config.json")

	DeleteFilesInDir(RESULT_DIR)

	mapper := map[string]string{
		"module_name":           config.ModuleName,
		"method_name":           config.MethodName,
		"tag":                   config.Tag,
		"lower_tag":             strings.ToLower(config.Tag),
		"route_path":            config.RoutePath,
		"letter_request_method": ToCapFirstLetter(config.RequestMethod),
		"lower_request_method":  strings.ToLower(config.RequestMethod),
		"lower_module_name":     strings.ToLower(config.ModuleName),
		"request_method":        strings.ToLower(config.RequestMethod),
		"dto_model":             config.DtoModel,
		"unpointer_dto_model":   strings.ReplaceAll(config.DtoModel, "*", ""),
		"repository":            config.Repository,
		"camel_repository":      ToCapFirstLower(config.Repository),
	}

	ReadTemplate("handler", mapper, &mainTemplate)
	ReadTemplate("get_swag", mapper, &swagTemplate)
	ReadTemplate("usecase", mapper, &usecaseTemplate)
	ReadTemplate("container", mapper, &containerTemplate)
	ReadTemplate("base_handler", mapper, &baseHandlerTemplate)
	ReadTemplate("router", mapper, &routerTemplate)

	mainTemplate.MapTemplate("swag_template", &swagTemplate)

	mainTemplate.WriteResultFile("handler")
	usecaseTemplate.WriteResultFile("usecase")
	containerTemplate.WriteResultFile("container")
	baseHandlerTemplate.WriteResultFile("base_handler")
	routerTemplate.WriteResultFile("router")
}
