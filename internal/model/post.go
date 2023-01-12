package model

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	ID            uint
	Title         string
	Content       string
	PostStatus    string
	CommentStatus string
}
