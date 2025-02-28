package models

type (
	MethodConfig struct {
		Request          string `json:"request"`
		Enable           int    `json:"enable"`
		Plural           int    `json:"plural"`
		PrefixMethodName string `json:"prefix_method_name"`
		Model            string `json:"model,omitempty"`
	}

	Config struct {
		ModuleName     string         `json:"module_name"`
		Tag            string         `json:"tag"`
		RoutePath      string         `json:"route_path"`
		Repository     string         `json:"repository"`
		BaseMethodName string         `json:"base_method_name"`
		Methods        []MethodConfig `json:"methods"`
		ModelAlias     string         `json:"model_alias"`
	}
)
