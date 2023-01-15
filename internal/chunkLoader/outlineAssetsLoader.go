package chunkLoader

import "html/template"

func outlineAssetsLoader(templates *template.Template) *template.Template {
	templates = template.Must(templates.ParseFiles(
		"./themes/frontStage/plain-sekai/pages/home/content.tmpl",
		"./themes/frontStage/plain-sekai/pages/post/content.tmpl",
		"./themes/frontStage/plain-sekai/template/footer/footer.tmpl",
		"./themes/frontStage/plain-sekai/template/header/header.tmpl",
		"./themes/frontStage/plain-sekai/template/mask/mask.tmpl",
		"./themes/backStage/plain-sekai/pages/login/content.tmpl",
	))
	return templates
}
