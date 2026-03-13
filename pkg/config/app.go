package config

import (
	"github.com/jinzhu/gorm"
	_	"github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func Connect() {

	d, err := gorm.Open(
		"mysql",
		"uu49b3v04yy1ijpr:y8NAtwVHlnFZ0nROT5Qp@tcp(bqpgwgr9a2z5wxnbvbe0-mysql.services.clever-cloud.com:3306)/bqpgwgr9a2z5wxnbvbe0?charset=utf8&parseTime=True&loc=Local",
	)

	if err != nil {
		panic("Failed to connect database")
	}

	DB = d
}

func GetDB() *gorm.DB {
	return DB
}
