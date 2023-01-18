package themeLoader

type themeConfig struct {
	ThemeName string `toml:"themeName"`
	ThemeDir  string
	Default   struct {
		Header string `toml:"header"`
		Footer string `toml:"footer"`
		Mask   string `toml:"mask"`
	} `toml:"default"`
	Entrance struct {
		EntrancesMap []string `toml:"EntrancesMap"`
	} `toml:"Entrance"`
	Static struct {
		StaticMap []string `toml:"staticMap"`
	} `toml:"static"`
}

type EntranceConfig struct {
	Entrance struct {
		Header  string `toml:"header"`
		Content string `toml:"content"`
		Footer  string `toml:"footer"`
		Mask    string `toml:"mask"`
	} `toml:"entrance"`
}

type PageConfig struct {
	Custom struct {
		Content   string   `toml:"content"`
		Header    string   `toml:"header"`
		Footer    string   `toml:"footer"`
		Mask      string   `toml:"mask"`
		Extension []string `toml:"extension"`
	} `toml:"custom"`
}
