package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// create a db with GORM
var (
	db *gorm.DB
)

//connect to db
func Connect() {
	// d, err := gorm.Open("mysql", "root:deep2002/simplerest?charset=utf8&parseTime=True")
	d, err := gorm.Open("mysql", "root:deep2002@tcp(127.0.0.1:3306)/simplerest?charset=utf8&parseTime=True")
	if err != nil {
		panic(err)
	}
	db = d
}

//get db after using gorm after connecting
func GetDB() *gorm.DB {
	return db
}
