package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID    uint
	Name  string
	Email string
}
