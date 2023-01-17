package model

type Theme struct {
	ThemeName      string
	ThemeUrl       string
	Default        Default
	Pages          []Page
	TemplateDetail map[string]string
}

type Default struct {
	headerTemplateName string
	footerTemplateName string
	maskTemplateName   string
}

type Page struct {
	templateString     string
	controllerURL      string
	headerTemplateName string
	footerTemplateName string
	maskTemplateName   string
}
