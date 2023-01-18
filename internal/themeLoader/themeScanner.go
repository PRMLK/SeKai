package themeLoader

import (
	"SeKai/internal/config"
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
	// 遍历所有Entrance
	theme.Entrances = make(map[string]model.Entrances)
	for _, entranceTomlString := range themeMap[themeName].Entrance.EntrancesMap {
		wt := bytes.NewBufferString("")
		var entrance model.Entrances
		var tempTemplate *template.Template
		tempTemplate = template.New("")

		// 去除所有空格和换行符
		entranceTomlString = util.StandardizeSpaces(entranceTomlString)
		entrance.TomlDir = strings.Split(entranceTomlString, ":")[0]
		entrance.ControllerURL = strings.Split(entranceTomlString, ":")[1]

		// 加载Entrance toml
		LoadEntrance(basicDir, themeName, themeMap, tempTemplate, entrance.TomlDir)

		// 加载默认模板
		LoadDefaultPages(basicDir, themeName, themeMap, tempTemplate)

		err := tempTemplate.ExecuteTemplate(wt, "entrance", map[string]interface{}{
			"sekaiPageTitle": config.ApplicationConfig.SiteConfig.SiteName,
			"sekaiSiteRoot":  "localhost:12070",
		})
		if err != nil {
			logger.ServerLogger.Error("加载 " + entranceTomlString + " 页面失败 : " + err.Error())
			return model.Theme{}
		}
		entrance.CompileString = wt.Bytes()
		theme.Entrances[entranceTomlString] = entrance
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

func LoadEntrance(basicDir string, themeName string, themeMap map[string]themeConfig, tempTemplate *template.Template, TomlDir string) {
	themeDir := themeMap[themeName].ThemeDir

	// 加载内置的基础模板
	nowTemplateString := inlineTemplateStringLoader()

	var tomlData []byte
	var entranceConfig EntranceConfig

	// 读取当前toml
	if tempData, err := os.ReadFile(basicDir + "/" + themeDir + "/" + TomlDir); err != nil {
		logger.ServerLogger.Error()
		return
	} else {
		tomlData = tempData
	}

	// 解析当前toml
	if err := toml.Unmarshal(tomlData, &entranceConfig); err != nil {
		logger.ServerLogger.Error()
	} else {
		// 加载 head 组件
		if entranceConfig.Entrance.Header != "" {
			// 如果是默认的，就不需要改变
			if entranceConfig.Entrance.Header != "default" {
				nowTemplateString, _ = util.ReplaceStringByRegex(nowTemplateString, "{{\\s*template\\s*\"header\"\\s*.\\s*}}", "{{ template \""+"root#header"+"\" .}}")
				LoadPage(basicDir, themeDir, entranceConfig.Entrance.Header, tempTemplate, "header", "root"+"#header")
			}
		} else {
			// 清空当前页面所有header标识不加载
			nowTemplateString, _ = util.ReplaceStringByRegex(nowTemplateString, "{{\\s*template\\s*\" "+"header"+" \"\\s*.\\s*}}", "")
		}

		// 加载 content 组件
		if entranceConfig.Entrance.Content != "" {
			nowTemplateString, _ = util.ReplaceStringByRegex(nowTemplateString, "{{\\s*template\\s*\"content\"\\s*.\\s*}}", "{{ template \""+"root#content"+"\" .}}")
			LoadPage(basicDir, themeDir, entranceConfig.Entrance.Content, tempTemplate, "content", "root"+"#content")
		} else {
			logger.ServerLogger.Error("content不能为空")
		}

		// 加载 footer 组件
		if entranceConfig.Entrance.Footer != "" {
			// 如果是默认的，就不需要改变
			if entranceConfig.Entrance.Footer != "default" {
				nowTemplateString, _ = util.ReplaceStringByRegex(nowTemplateString, "{{\\s*template\\s*\"footer\"\\s*.\\s*}}", "{{ template \""+"root#footer"+"\" .}}")
				LoadPage(basicDir, themeDir, entranceConfig.Entrance.Footer, tempTemplate, "footer", "root"+"#footer")
			}
		} else {
			// 清空当前页面所有header标识不加载
			nowTemplateString, _ = util.ReplaceStringByRegex(nowTemplateString, "{{\\s*template\\s*\" "+"footer"+" \"\\s*.\\s*}}", "")
		}

		// 加载 mask 组件
		if entranceConfig.Entrance.Mask != "" {
			// 如果是默认的，就不需要改变
			if entranceConfig.Entrance.Mask != "default" {
				nowTemplateString, _ = util.ReplaceStringByRegex(nowTemplateString, "{{\\s*template\\s*\"mask\"\\s*.\\s*}}", "{{ template \""+"root#mask"+"\" .}}")
				LoadPage(basicDir, themeDir, entranceConfig.Entrance.Mask, tempTemplate, "mask", "root"+"#mask")
			}
		} else {
			// 清空当前页面所有header标识不加载
			nowTemplateString, _ = util.ReplaceStringByRegex(nowTemplateString, "{{\\s*template\\s*\" "+"mask"+" \"\\s*.\\s*}}", "")
		}

		if tempTemplate, err = tempTemplate.New("entrance").Parse(nowTemplateString); err != nil {
		}
	}
}

func LoadDefaultPages(basicDir string, themeName string, themeMap map[string]themeConfig, importTemplate *template.Template) {
	themeDir := themeMap[themeName].ThemeDir
	// 读取footer
	LoadPage(basicDir, themeDir, themeMap[themeName].Default.Footer, importTemplate, "footer", "footer")

	// 读取header
	LoadPage(basicDir, themeDir, themeMap[themeName].Default.Header, importTemplate, "header", "header")

	// 读取mask
	LoadPage(basicDir, themeDir, themeMap[themeName].Default.Mask, importTemplate, "mask", "mask")
}

func LoadPage(basicDir string, themeDir string, TomlDir string, template *template.Template, templateType string, templateLink string) {
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
					nowData, _ = util.ReplaceStringByRegex(nowData, "{{\\s*template\\s*\"header\"\\s*.\\s*}}", "{{ template \""+templateType+"#header"+"\" .}}")
					LoadPage(basicDir, themeDir, pageConfig.Custom.Header, template, "header", templateType+"#header")
				}
			} else {
				// 清空当前页面所有header标识不加载
				nowData, _ = util.ReplaceStringByRegex(nowData, "{{\\s*template\\s*\" "+"header"+" \"\\s*.\\s*}}", "")
			}
			if pageConfig.Custom.Footer != "" {
				if pageConfig.Custom.Footer != "default" {
					nowData, _ = util.ReplaceStringByRegex(nowData, "{{\\s*template\\s*\"footer\"\\s*.\\s*}}", "{{ template \""+templateType+"#footer"+"\" .}}")
					LoadPage(basicDir, themeDir, pageConfig.Custom.Footer, template, "footer", templateType+"#footer")
				}
			} else {
				// 清空当前页面所有footer标识不加载
				nowData, _ = util.ReplaceStringByRegex(nowData, "{{\\s*template\\s*\" "+"footer"+" \"\\s*.\\s*}}", "")
			}
			if pageConfig.Custom.Mask != "" {
				if pageConfig.Custom.Mask != "default" {
					nowData, _ = util.ReplaceStringByRegex(nowData, "{{\\s*template\\s*\"mask\"\\s*.\\s*}}", "{{ template \""+templateType+"#mask"+"\" .}}")
					LoadPage(basicDir, themeDir, pageConfig.Custom.Mask, template, "mask", templateType+"#mask")
				}
			} else {
				// 清空当前页面所有mask标识不加载
				nowData, _ = util.ReplaceStringByRegex(nowData, "{{\\s*template\\s*\" "+"mask"+" \"\\s*.\\s*}}", "")
			}

			// 加载自身
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
