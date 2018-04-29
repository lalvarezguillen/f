package main

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

type Reel struct {
	gorm.Model
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Pictures    []Picture `json:"pictures" gorm:"many2many:reel_pictures;"`
	UserID      uint      `json:"userID"`
}

type Picture struct {
	gorm.Model
	Caption string `json:"caption"`
	URL     string `json:"url"`
	UserID  uint   `json:"userID"`
}
