package themeLoader

type themeConfig struct {
	ThemeName string `toml:"themeName"`
	ThemeDir  string
	Default   struct {
		Header string `toml:"header"`
		Footer string `toml:"footer"`
		Mask   string `toml:"mask"`
	} `toml:"default"`
	Pages struct {
		PagesMap []string `toml:"pagesMap"`
	} `toml:"pages"`
	Static struct {
		StaticMap []string `toml:"staticMap"`
	} `toml:"static"`
}

type PageConfig struct {
	Custom struct {
		Content string `toml:"content"`
		Header  string `toml:"header"`
		Footer  string `toml:"footer"`
		Mask    string `toml:"mask"`
	} `toml:"custom"`
}
