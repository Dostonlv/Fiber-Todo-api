package model

import "github.com/jinzhu/gorm"

type Todo struct {
	gorm.Model
	Title       string `json:"title"`
	Status      bool   `json:"status"`
	Description string `json:"description"`
}
