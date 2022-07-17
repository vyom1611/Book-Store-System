package config

import (
	"github.com/jinzhu/gorm"
)

var(
	db * gorm.DB
)

//Function to connect with MySQL database
func Connect() {
	data, err := gorm.Open("mysql", "DATABASE_INFO")
	if err != nil {
		panic(err)
	}

	//Transfering the data to db database
	db = data
}

//Helper function to get database
func GetDB() *gorm.DB {
	return db
}