package themeLoader

import (
	"SeKai"
)

func inlineTemplateStringLoader() string {
	file, err := SeKai.InlineTmpl.ReadFile("internal/themeLoader/tmpl/root.tmpl")
	if err != nil {
		return ""
	}
	return string(file)
}
