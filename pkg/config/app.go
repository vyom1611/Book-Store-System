package config

import (
	"github.com/jinzhu/gorm"
)

var(
	db * gorm.DB
)

//Function to connect with MySQL database
func Connect() {
	d, err := gorm.Open("mysql", "DATABASE_INFO")
}