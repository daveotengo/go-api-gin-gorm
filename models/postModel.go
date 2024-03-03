package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title string // A regular string field
	Body  string // A pointer to a string, allowing for null values
}
