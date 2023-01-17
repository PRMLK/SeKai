package themeLoader

import (
	"SeKai"
	"html/template"
)

func inlineTemplateLoader(templateMap map[string]*template.Template) map[string]*template.Template {
	templateMap["root"] = template.New("")
	_, _ = templateMap["root"].ParseFS(SeKai.InlineTmpl, "internal/themeLoader/tmpl/root.tmpl")
	return templateMap
}
