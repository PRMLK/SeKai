package model

import (
	"SeKai/internal/logger"
	"SeKai/internal/util"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username         string
	Nickname         string
	Email            string
	Password         string
	FirstName        string
	LastName         string
	SiteUrl          string
	Language         string
	Bio              string
	ProfilePhoto     string
	GoogleAuthSecret string
}

func init() {
	if err := util.Datebase.AutoMigrate(&User{}); err != nil {
		logger.ServerLogger.Panic(err)
	}
}
