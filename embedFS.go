package SeKai

import (
	"embed"
)

/*
	embed在导入的时候是以注释所在文件的相对目录来导入的，所以放根目录应该比较好调用到所有的文件目录，所以暂时放这
	Refer: https://stackoverflow.com/questions/66285635/how-do-you-use-go-1-16-embed-features-in-subfolders-packages
*/

//go:embed internal/chunkLoader/tmpl/*
var InlineTmpl embed.FS
