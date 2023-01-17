package model

type Theme struct {
	ThemeName      string
	ThemeUrl       string
	Default        Default
	Pages          map[string]Page
	TemplateDetail map[string]string
	StaticFiles    []StaticFile
}

type Default struct {
	TemplateString     string
	ControllerURL      string
	HeaderTemplateName string
	FooterTemplateName string
	MaskTemplateName   string
}

type Page struct {
	TomlDir             string
	ControllerURL       string
	HeaderTemplateName  string
	ContentTemplateName string
	FooterTemplateName  string
	MaskTemplateName    string
	CompileString       []byte
}

type StaticFile struct {
	FileDir       string
	ControllerURL string
}
