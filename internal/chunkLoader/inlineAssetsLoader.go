package chunkLoader

import (
	"SeKai"
	"html/template"
)

func inlineAssetsLoader(templates *template.Template) *template.Template {
	templates = template.Must(templates.ParseFS(SeKai.InlineTmpl, "internal/chunkLoader/tmpl/*.tmpl"))
	return templates
}
