package main

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Id           int
	Post_title   string
	Post_content string
}
