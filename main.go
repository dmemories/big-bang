package main

import (
	"big_bang/configs"
	"big_bang/models"
	"big_bang/utils"
	"fmt"
	"strings"
)

var (
	MAIN_CONFIG   = &models.Config{}
	MAIN_MAPPER   = map[string]string{}
	METHOD_MAPPER = map[string]string{}

/*
swagTemplate        models.Template
usecaseTemplate     models.Template
containerTemplate   models.Template
baseHandlerTemplate models.Template
routerTemplate      models.Template
*/
)

func init() {
	MAIN_CONFIG = utils.LoadConfig("configs/config.json")

	utils.DeleteFilesInDir(configs.RESULT_DIR)

	MAIN_MAPPER = map[string]string{
		"module_name":        MAIN_CONFIG.ModuleName,
		"plural_module_name": utils.GetPluralName(MAIN_CONFIG.ModuleName),
		"tag":                MAIN_CONFIG.Tag,
		"base_method_name":   MAIN_CONFIG.BaseMethodName,
		"lower_module_name":  strings.ToLower(MAIN_CONFIG.ModuleName),
		"route_path":         MAIN_CONFIG.RoutePath,

		"repository":         MAIN_CONFIG.Repository,
		"camel_repository":   utils.ToFirstLower(MAIN_CONFIG.Repository),
		"model_alias":        MAIN_CONFIG.ModelAlias,
		"letter_model_alias": utils.ToFirstUpper(MAIN_CONFIG.ModelAlias),
	}
}

func main() {
	//methodTemplates := []*models.Template{}
	for _, val := range MAIN_CONFIG.Methods {
		if val.Enable < 1 {
			for _, subFileName := range configs.METHOD_SUB_FILE_NAMES {
				methodFileName := utils.GetMethodFileName(val.Request, subFileName)
				METHOD_MAPPER[methodFileName] = ""
			}
			continue
		}

		var (
			prefixMethodName string
			suffixMethodName string
		)

		prefixMethodName = utils.ToFirstUpper(strings.Replace(val.Request, "post", "Create", 1))
		if val.Plural > 0 {
			suffixMethodName = utils.GetPluralName(MAIN_CONFIG.BaseMethodName)
		} else {
			suffixMethodName = MAIN_CONFIG.BaseMethodName
		}

		for _, subFileName := range configs.METHOD_SUB_FILE_NAMES {
			methodFileName := utils.GetMethodFileName(val.Request, subFileName)

			METHOD_MAPPER[methodFileName] = utils.ReadTemplateWithMapper(fmt.Sprintf("%s/%s", configs.METHOD_DIR, methodFileName), MAIN_MAPPER, METHOD_MAPPER, map[string]string{
				"method_name":         prefixMethodName + suffixMethodName,
				"suffix_method_name":  suffixMethodName,
				"dto_model":           val.Model,
				"unpointer_dto_model": strings.ReplaceAll(val.Model, "*", ""),
			})
		}
	}

	mainTemplate := utils.ReadMainTemplateFile("handler", METHOD_MAPPER, MAIN_MAPPER)
	usecaseTemplate := utils.ReadMainTemplateFile("usecase", METHOD_MAPPER, MAIN_MAPPER)
	repositoryTemplate := utils.ReadMainTemplateFile("repository", METHOD_MAPPER, MAIN_MAPPER)
	modelTemplate := utils.ReadMainTemplateFile("model", METHOD_MAPPER, MAIN_MAPPER)

	mainTemplate.WriteResultFile()
	usecaseTemplate.WriteResultFile()
	repositoryTemplate.WriteResultFile()
	modelTemplate.WriteResultFile()
}
