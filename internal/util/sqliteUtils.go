package util

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Datebase *gorm.DB

func init() {
	if tempDatebase, err := gorm.Open(sqlite.Open("sekai.db"), &gorm.Config{}); err != nil {
		panic(err)
	} else {
		Datebase = tempDatebase
	}
}
