package dto

import "gorm.io/gorm"

type UserDto struct {
	gorm.Model
	Username     string
	Nickname     string
	Email        string
	FirstName    string
	LastName     string
	SiteUrl      string
	Language     string
	Bio          string
	ProfilePhoto string
	Role         string
}
