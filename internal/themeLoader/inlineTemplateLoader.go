package themeLoader

import (
	"SeKai"
	"html/template"
)

func inlineTemplateLoader() *template.Template {
	var tempTemplate *template.Template
	tempTemplate = template.New("")
	_, _ = tempTemplate.ParseFS(SeKai.InlineTmpl, "internal/themeLoader/tmpl/root.tmpl")
	return tempTemplate
}
