package model

import (
	"SeKai/internal/logger"
	"SeKai/internal/util"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName string
	Email    string
	Password string
}

func init() {
	if err := util.Datebase.AutoMigrate(&User{}); err != nil {
		logger.ServerLogger.Panic(err)
	}
}
