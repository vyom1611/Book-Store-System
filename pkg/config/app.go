package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var(
	db * gorm.DB
)

//Function to connect with MySQL database
func Connect() {
	data, err := gorm.Open("mysql", "root:vijay1610@localhost/mysql?charset=utf8mb4&parseTime=True&loc=Local")
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