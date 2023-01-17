package themeLoader

import (
	"SeKai/internal/logger"
	"SeKai/internal/model"
	"SeKai/internal/util"
	"bytes"
	"github.com/pelletier/go-toml/v2"
	"html/template"
	"os"
	"strings"
)

func ThemeBasicScan(basicDir string, themeMap map[string]themeConfig) {
	// 扫描backStage目录
	backStageRootDir, err := os.ReadDir(basicDir)
	if err != nil {
		return
	}
	for _, dir := range backStageRootDir {
		// 如果是目录
		if dir.IsDir() {
			var data []byte
			var themeConfig themeConfig
			// 读取该目录下manifest.toml
			if tempData, err := os.ReadFile(basicDir + "/" + dir.Name() + "/manifest.toml"); err != nil {
				logger.ServerLogger.Debug()
				continue
			} else {
				data = tempData
			}
			// 尝试读取到Config
			if err := toml.Unmarshal(data, &themeConfig); err != nil {
				logger.ServerLogger.Debug()
				continue
			} else {
				if _, exist := themeMap[themeConfig.ThemeName]; exist == true {
					// 已经存在同名模板
					logger.ServerLogger.Debug()
					continue
				} else {
					themeConfig.ThemeDir = dir.Name()
					themeMap[themeConfig.ThemeName] = themeConfig
				}
			}
		}
	}
}

func SingleThemeScan(basicDir string, themeName string, themeMap map[string]themeConfig) model.Theme {
	var theme model.Theme

	themeDir := themeMap[themeName].ThemeDir
	// 遍历所有Pages
	theme.Pages = make(map[string]model.Page)
	for _, pageTomlString := range themeMap[themeName].Pages.PagesMap {
		wt := bytes.NewBufferString("")
		var page model.Page
		var tempTemplate *template.Template

		// 去除所有空格和换行符
		pageTomlString = util.StandardizeSpaces(pageTomlString)
		page.TomlDir = strings.Split(pageTomlString, ":")[0]
		page.ControllerURL = strings.Split(pageTomlString, ":")[1]

		// 加载内置的基础模板
		tempTemplate = InlineTemplateMap["root"]

		// 加载default模板
		LoadDefaultPages(basicDir, themeName, themeMap, tempTemplate)

		CompilePage(basicDir, themeDir, page.TomlDir, tempTemplate, "content", "content")

		err := tempTemplate.ExecuteTemplate(wt, "root", map[string]interface{}{
			"sekaiPageTitle": "123",
			"sekaiSiteRoot":  "http://localhost:12070",
		})
		if err != nil {
			logger.ServerLogger.Error("加载 " + pageTomlString + " 页面失败 : " + err.Error())
			return model.Theme{}
		}
		page.CompileString = wt.Bytes()
		theme.Pages[pageTomlString] = page
	}
	// 加载静态文件
	for _, staticFileToml := range themeMap[themeName].Static.StaticMap {
		var staticFile model.StaticFile
		staticFileString := util.StandardizeSpaces(staticFileToml)
		staticFile.FileDir = basicDir + "/" + themeDir + "/" + strings.Split(staticFileString, ":")[0]
		staticFile.ControllerURL = strings.Split(staticFileString, ":")[1]
		theme.StaticFiles = append(theme.StaticFiles, staticFile)
	}
	return theme
}

func LoadDefaultPages(basicDir string, themeName string, themeMap map[string]themeConfig, importTemplate *template.Template) {
	themeDir := themeMap[themeName].ThemeDir
	// 读取footer
	if tempData, err := os.ReadFile(basicDir + "/" + themeDir + "/" + themeMap[themeName].Default.Footer); err != nil {
		logger.ServerLogger.Error()
		return
	} else {
		_, err := importTemplate.New("footer").Parse(string(tempData))
		if err != nil {
			return
		}
	}

	// 读取header
	if tempData, err := os.ReadFile(basicDir + "/" + themeDir + "/" + themeMap[themeName].Default.Header); err != nil {
		logger.ServerLogger.Error()
		return
	} else {
		_, err := importTemplate.New("header").Parse(string(tempData))
		if err != nil {
			return
		}
	}

	// 读取mask
	if tempData, err := os.ReadFile(basicDir + "/" + themeDir + "/" + themeMap[themeName].Default.Mask); err != nil {
		logger.ServerLogger.Error()
		return
	} else {
		_, err := importTemplate.New("mask").Parse(string(tempData))
		if err != nil {
			return
		}
	}
}

func CompilePage(basicDir string, themeDir string, TomlDir string, template *template.Template, templateType string, templateLink string) {
	var tomlData []byte
	var pageConfig PageConfig

	// 读取当前toml
	if tempData, err := os.ReadFile(basicDir + "/" + themeDir + "/" + TomlDir); err != nil {
		logger.ServerLogger.Error()
		return
	} else {
		tomlData = tempData
	}

	// 解析当前toml
	if err := toml.Unmarshal(tomlData, &pageConfig); err != nil {
		logger.ServerLogger.Error()
	} else {
		// 读取当前页面的content
		if tempData, err := os.ReadFile(basicDir + "/" + themeDir + "/" + pageConfig.Custom.Content); err != nil {
			logger.ServerLogger.Error()
			return
		} else {
			nowData := string(tempData)
			// 先加载子组件
			if pageConfig.Custom.Header != "" {
				// 如果是默认的，就不需要改变
				if pageConfig.Custom.Header != "default" {
					CompilePage(basicDir, themeDir, pageConfig.Custom.Header, template, "header", templateType+"#header")
				}
			} else {
				// 清空当前页面所有header标识不加载
				nowData, _ = util.ReplaceStringByRegex(string(nowData), "{{\\s*template\\s*\" "+"header"+" \"\\s*.\\s*}}", "")
			}
			if pageConfig.Custom.Footer != "" {
				if pageConfig.Custom.Footer != "default" {
					CompilePage(basicDir, themeDir, pageConfig.Custom.Footer, template, "footer", templateType+"#footer")
				}
			} else {
				// 清空当前页面所有footer标识不加载
				nowData, _ = util.ReplaceStringByRegex(string(nowData), "{{\\s*template\\s*\" "+"footer"+" \"\\s*.\\s*}}", "")
			}
			if pageConfig.Custom.Mask != "" {
				if pageConfig.Custom.Mask != "default" {
					CompilePage(basicDir, themeDir, pageConfig.Custom.Mask, template, "mask", templateType+"#mask")
				}
			} else {
				// 清空当前页面所有mask标识不加载
				nowData, _ = util.ReplaceStringByRegex(string(nowData), "{{\\s*template\\s*\" "+"mask"+" \"\\s*.\\s*}}", "")
			}

			// 加载自身
			nowData, _ = util.ReplaceStringByRegex(string(tempData), "{{\\s*template\\s*\" "+templateType+" \"\\s*.\\s*}}", "{{\\s*template\\s*\" "+templateLink+" \"\\s*.\\s* .}}")
			if template, err = template.New(templateLink).Parse(nowData); err != nil {
			}
		}
	}
}

//func SinglePageLoader(basicDir string, themeDir string, pageTomlDir string, pageControllerURL string, templateMap map[string]string) (page model.Page) {
//	page.TomlDir = pageTomlDir
//	page.ControllerURL = pageControllerURL
//
//	// 读取page的toml文件
//	var pageConfig PageConfig
//	if tempData, err := os.ReadFile(basicDir + "/" + themeDir + "/" + page.TomlDir); err != nil {
//		logger.ServerLogger.Error(err)
//		return
//	} else {
//		// 尝试读取到Config
//		if err := toml.Unmarshal(tempData, &pageConfig); err != nil {
//			logger.ServerLogger.Error(err)
//		}
//	}
//	page.HeaderTemplateName = pageConfig.Custom.Header
//	page.ContentTemplateName = pageConfig.Custom.Content
//	page.FooterTemplateName = pageConfig.Custom.Footer
//	page.MaskTemplateName = pageConfig.Custom.Mask
//
//	SingleTemplateLoader(basicDir, themeDir, page.HeaderTemplateName, templateMap)
//	SingleTemplateLoader(basicDir, themeDir, page.ContentTemplateName, templateMap)
//	SingleTemplateLoader(basicDir, themeDir, page.FooterTemplateName, templateMap)
//	SingleTemplateLoader(basicDir, themeDir, page.MaskTemplateName, templateMap)
//
//	page.CompileTemplate = InlineTemplateMap["chunk"]
//	page.CompileTemplate.New("header").ParseGlob()
//
//	return page
//}
//
//func SingleTemplateLoader(basicDir string, themeDir string, templateName string, templateMap map[string]string) map[string]string {
//	if tempData, err := os.ReadFile(basicDir + "/" + themeDir + "/" + templateName); err != nil {
//		logger.ServerLogger.Error(err)
//	} else {
//		if _, exist := templateMap[templateName]; exist == true {
//			logger.ServerLogger.Debug()
//		} else {
//			templateMap[templateName] = string(tempData)
//		}
//	}
//	return templateMap
//}
