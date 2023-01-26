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
				continue
			} else {
				data = tempData
			}
			// 尝试读取到Config
			if err := toml.Unmarshal(data, &themeConfig); err != nil {
				logger.ServerLogger.Debug(config.LanguageConfig.ServerLogger.TomlUnmarshalError + ": " + err.Error())
				continue
			} else {
				if _, exist := themeMap[themeConfig.ThemeName]; exist == true {
					// 已经存在同名模板
					logger.ServerLogger.Debug(config.LanguageConfig.ServerLogger.SameThemeExist + ": " + themeConfig.ThemeName)
					continue
				} else {
					themeConfig.ThemeDir = dir.Name()
					themeMap[themeConfig.ThemeName] = themeConfig
				}
			}
		}
	}
}

func SingleThemeScan(basicDir string, themeName string, themeMap map[string]themeConfig, theme *model.Theme) {
	themeDir := themeMap[themeName].ThemeDir
	// 遍历所有Entrance
	for _, entranceTomlString := range themeMap[themeName].Entrance.EntrancesMap {
		wt := bytes.NewBufferString("")
		var entrance model.Entrances
		var tempTemplate *template.Template
		tempTemplate = template.New("")

		// 去除所有空格和换行符
		entranceTomlString = util.StandardizeSpaces(entranceTomlString)
		entrance.TomlDir = strings.Split(entranceTomlString, "::")[0]
		entrance.ControllerURL = strings.Split(entranceTomlString, "::")[1]

		// 加载Entrance toml
		LoadEntrance(basicDir, themeName, themeMap, tempTemplate, entrance.TomlDir)

		// 加载默认模板
		LoadDefaultPages(basicDir, themeName, themeMap, tempTemplate)

		err := tempTemplate.ExecuteTemplate(wt, "entrance", map[string]interface{}{
			"sekaiSiteRoot":        config.ApplicationConfig.SiteConfig.SiteRoot,
			"sekaiSiteHome":        config.ApplicationConfig.SiteConfig.SiteHome,
			"sekaiSiteName":        config.ApplicationConfig.SiteConfig.SiteName,
			"sekaiSiteDescription": config.ApplicationConfig.SiteConfig.SiteDescription,
		})
		if err != nil {
			logger.ServerLogger.Error(config.LanguageConfig.ServerLogger.LoadPageError + " " + entranceTomlString + ": " + err.Error())
			return
		}
		entrance.CompileString = wt.Bytes()
		exist := false
		// 判断是否已经存在同名entrances（主要用于重加载）
		for i := range theme.Entrances {
			if theme.Entrances[i].TomlDir == entrance.TomlDir {
				exist = true
				theme.Entrances[i].CompileString = entrance.CompileString
				break
			}
		}
		if !exist {
			theme.Entrances = append(theme.Entrances, entrance)
		}
	}
	// 加载静态文件
	for _, staticFileToml := range themeMap[themeName].Static.StaticMap {
		var staticFile model.StaticFile
		staticFileString := util.StandardizeSpaces(staticFileToml)
		staticFile.FileDir = basicDir + "/" + themeDir + "/" + strings.Split(staticFileString, "::")[0]
		staticFile.ControllerURL = strings.Split(staticFileString, "::")[1]
		theme.StaticFiles = append(theme.StaticFiles, staticFile)
	}
}

func LoadEntrance(basicDir string, themeName string, themeMap map[string]themeConfig, tempTemplate *template.Template, TomlDir string) {
	themeDir := themeMap[themeName].ThemeDir
	nowDir := basicDir + "/" + themeDir + "/" + TomlDir

	// 加载内置的基础模板
	nowTemplateString := inlineTemplateStringLoader()

	var tomlData []byte
	var entranceConfig EntranceConfig

	// 读取当前toml
	if tempData, err := os.ReadFile(nowDir); err != nil {
		logger.ServerLogger.Error(config.LanguageConfig.ServerLogger.ReadFileError + " " + nowDir + ": " + err.Error())
		return
	} else {
		tomlData = tempData
	}

	// 解析当前toml
	if err := toml.Unmarshal(tomlData, &entranceConfig); err != nil {
		logger.ServerLogger.Error(config.LanguageConfig.ServerLogger.TomlUnmarshalError + " " + nowDir + ": " + err.Error())
		return
	} else {
		// 加载 head 组件
		if entranceConfig.Entrance.Header != "" {
			// 如果是默认的，就不需要改变
			if entranceConfig.Entrance.Header != "default" {
				nowTemplateString, _ = util.ReplaceStringByRegex(nowTemplateString, "{{\\s*template\\s*\"header\"\\s*.\\s*}}", "{{ template \""+"root#header"+"\" .}}")
				if errorDir, err := LoadPage(basicDir, themeDir, entranceConfig.Entrance.Header, tempTemplate, "header", "root"+"#header"); err != nil {
					logger.ServerLogger.Error(config.LanguageConfig.ServerLogger.LoadPageError + " " + errorDir + ": " + err.Error())
					return
				}
			}
		} else {
			// 清空当前页面所有header标识不加载
			nowTemplateString, _ = util.ReplaceStringByRegex(nowTemplateString, "{{\\s*template\\s*\""+"header"+"\"\\s*.\\s*}}", "")
		}

		// 加载 content 组件
		if entranceConfig.Entrance.Content != "" {
			nowTemplateString, _ = util.ReplaceStringByRegex(nowTemplateString, "{{\\s*template\\s*\"content\"\\s*.\\s*}}", "{{ template \""+"root#content"+"\" .}}")
			if errorDir, err := LoadPage(basicDir, themeDir, entranceConfig.Entrance.Content, tempTemplate, "content", "root"+"#content"); err != nil {
				logger.ServerLogger.Error(config.LanguageConfig.ServerLogger.LoadPageError + " " + errorDir + ": " + err.Error())
				return
			}
		} else {
			logger.ServerLogger.Error("content不能为空")
		}

		// 加载 footer 组件
		if entranceConfig.Entrance.Footer != "" {
			// 如果是默认的，就不需要改变
			if entranceConfig.Entrance.Footer != "default" {
				nowTemplateString, _ = util.ReplaceStringByRegex(nowTemplateString, "{{\\s*template\\s*\"footer\"\\s*.\\s*}}", "{{ template \""+"root#footer"+"\" .}}")
				if errorDir, err := LoadPage(basicDir, themeDir, entranceConfig.Entrance.Footer, tempTemplate, "footer", "root"+"#footer"); err != nil {
					logger.ServerLogger.Error(config.LanguageConfig.ServerLogger.LoadPageError + " " + errorDir + ": " + err.Error())
					return
				}
			}
		} else {
			// 清空当前页面所有header标识不加载
			nowTemplateString, _ = util.ReplaceStringByRegex(nowTemplateString, "{{\\s*template\\s*\""+"footer"+"\"\\s*.\\s*}}", "")
		}

		// 加载 mask 组件
		if entranceConfig.Entrance.Mask != "" {
			// 如果是默认的，就不需要改变
			if entranceConfig.Entrance.Mask != "default" {
				nowTemplateString, _ = util.ReplaceStringByRegex(nowTemplateString, "{{\\s*template\\s*\"mask\"\\s*.\\s*}}", "{{ template \""+"root#mask"+"\" .}}")
				if errorDir, err := LoadPage(basicDir, themeDir, entranceConfig.Entrance.Mask, tempTemplate, "mask", "root"+"#mask"); err != nil {
					logger.ServerLogger.Error(config.LanguageConfig.ServerLogger.LoadPageError + " " + errorDir + ": " + err.Error())
					return
				}
			}
		} else {
			// 清空当前页面所有header标识不加载
			nowTemplateString, _ = util.ReplaceStringByRegex(nowTemplateString, "{{\\s*template\\s*\""+"mask"+"\"\\s*.\\s*}}", "")
		}

		if tempTemplate, err = tempTemplate.New("entrance").Parse(nowTemplateString); err != nil {
			logger.ServerLogger.Debug(err)
		}
	}
}

func LoadDefaultPages(basicDir string, themeName string, themeMap map[string]themeConfig, importTemplate *template.Template) {
	themeDir := themeMap[themeName].ThemeDir
	// 读取footer
	if errorDir, err := LoadPage(basicDir, themeDir, themeMap[themeName].Default.Footer, importTemplate, "footer", "footer"); err != nil {
		logger.ServerLogger.Error(config.LanguageConfig.ServerLogger.LoadPageError + " " + errorDir + ": " + err.Error())
		return
	}

	// 读取header
	if errorDir, err := LoadPage(basicDir, themeDir, themeMap[themeName].Default.Header, importTemplate, "header", "header"); err != nil {
		logger.ServerLogger.Error(config.LanguageConfig.ServerLogger.LoadPageError + " " + errorDir + ": " + err.Error())
		return
	}

	// 读取mask
	if errorDir, err := LoadPage(basicDir, themeDir, themeMap[themeName].Default.Mask, importTemplate, "mask", "mask"); err != nil {
		logger.ServerLogger.Error(config.LanguageConfig.ServerLogger.LoadPageError + " " + errorDir + ": " + err.Error())
		return
	}
}

func LoadPage(basicDir string, themeDir string, TomlDir string, template *template.Template, templateType string, templateLink string) (string, error) {
	var tomlData []byte
	var pageConfig PageConfig
	var nowDir = basicDir + "/" + themeDir + "/" + TomlDir

	// 读取当前toml
	if tempData, err := os.ReadFile(nowDir); err != nil {
		return nowDir, err
	} else {
		tomlData = tempData
	}

	// 解析当前toml
	if err := toml.Unmarshal(tomlData, &pageConfig); err != nil {
		return nowDir, err
	} else {
		// 读取当前页面的content
		if tempData, err := os.ReadFile(basicDir + "/" + themeDir + "/" + pageConfig.Custom.Content); err != nil {
			return nowDir, err
		} else {
			nowData := string(tempData)
			// 先加载子组件
			if pageConfig.Custom.Header != "" {
				// 如果是默认的，就不需要改变
				if pageConfig.Custom.Header != "default" {
					nowData, _ = util.ReplaceStringByRegex(nowData, "{{\\s*template\\s*\"header\"\\s*.\\s*}}", "{{ template \""+templateType+"#header"+"\" .}}")
					if errorDir, err := LoadPage(basicDir, themeDir, pageConfig.Custom.Header, template, "header", templateType+"#header"); err != nil {
						return errorDir, err
					}
				}
			} else {
				// 清空当前页面所有 header 标识不加载
				nowData, _ = util.ReplaceStringByRegex(nowData, "{{\\s*template\\s*\""+"header"+"\"\\s*.\\s*}}", "")
			}
			if pageConfig.Custom.Footer != "" {
				if pageConfig.Custom.Footer != "default" {
					nowData, _ = util.ReplaceStringByRegex(nowData, "{{\\s*template\\s*\"footer\"\\s*.\\s*}}", "{{ template \""+templateType+"#footer"+"\" .}}")
					if errorDir, err := LoadPage(basicDir, themeDir, pageConfig.Custom.Footer, template, "footer", templateType+"#footer"); err != nil {
						return errorDir, err
					}
				}
			} else {
				// 清空当前页面所有 footer 标识不加载
				nowData, _ = util.ReplaceStringByRegex(nowData, "{{\\s*template\\s*\""+"footer"+"\"\\s*.\\s*}}", "")
			}
			if pageConfig.Custom.Mask != "" {
				if pageConfig.Custom.Mask != "default" {
					nowData, _ = util.ReplaceStringByRegex(nowData, "{{\\s*template\\s*\"mask\"\\s*.\\s*}}", "{{ template \""+templateType+"#mask"+"\" .}}")
					if errorDir, err := LoadPage(basicDir, themeDir, pageConfig.Custom.Mask, template, "mask", templateType+"#mask"); err != nil {
						return errorDir, err
					}
				}
			} else {
				// 清空当前页面所有 mask 标识不加载
				nowData, _ = util.ReplaceStringByRegex(nowData, "{{\\s*template\\s*\""+"mask"+"\"\\s*.\\s*}}", "")
			}
			// 加载自身
			if template, err = template.New(templateLink).Parse(nowData); err != nil {
				return nowDir, err
			}
		}
	}
	return nowDir, nil
}
