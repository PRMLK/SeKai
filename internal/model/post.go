package model

import (
	"SeKai/internal/logger"
	"SeKai/internal/util"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title         string
	Author        uint
	Content       string
	PostStatus    string
	CommentStatus string
}

func init() {
	if err := util.Datebase.AutoMigrate(&Post{}); err != nil {
		logger.ServerLogger.Panic(err)
	}
}
