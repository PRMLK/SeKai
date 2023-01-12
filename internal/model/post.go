package model

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title         string
	Content       string
	PostStatus    string
	CommentStatus string
}
