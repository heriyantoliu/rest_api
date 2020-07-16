package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var materials []Material

// const usernameDB string = "root"
// const passwordDB string = "example"
// const addressDB string = "localhost"
// const portDB string = "3306"
// const dbName string = "rest_api"

// var db *gorm.DB

// func initDb() {
// 	dburl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", usernameDB, passwordDB, addressDB, portDB, dbName)
// 	db, _ = gorm.Open("mysql", dburl)
// 	// defer db.Close()

// 	db.AutoMigrate(&Material{})
// }
