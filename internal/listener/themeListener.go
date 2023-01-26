package listener

import (
	"SeKai/internal/logger"
	"SeKai/internal/themeLoader"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"os"
)

func ThemeListener() {
	watcher, _ := fsnotify.NewWatcher()
	getThemeDirMap("./themes", watcher)
	go func() {
		defer watcher.Close()
		for true {
			select {
			case event := <-watcher.Events:
				{
					if event.Op&fsnotify.Write == fsnotify.Write {
						logger.ServerLogger.Debug("主题重新加载中...")
						themeLoader.InitThemeLoader()
						logger.ServerLogger.Debug("主题已重新加载...")
					}
				}
			case err := <-watcher.Errors:
				fmt.Println(err.Error())
			}
		}
	}()
}

// 获取主题所有目录并加入watcher
func getThemeDirMap(baseDirString string, watcher *fsnotify.Watcher) {
	baseDir, err := os.ReadDir(baseDirString)
	if err != nil {
		return
	}
	if err := watcher.Add(baseDirString); err != nil {
		logger.ServerLogger.Debug(err)
		return
	}
	for _, nextDirString := range baseDir {
		if nextDirString.IsDir() {
			getThemeDirMap(baseDirString+"/"+nextDirString.Name(), watcher)
		}
	}
}
