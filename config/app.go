package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() {
	// refer: https://gorm.io/docs/connecting_to_the_database.html#MySQL
	dsn := "zohaib:1234@tcp(127.0.0.1:3306)/my_db?charset=utf8mb4&parseTime=True&loc=Local"
	mysqlDb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = mysqlDb
	//to create table of book  => here the table will be created with name= books
}

func GetDb() *gorm.DB {
	return db
}
