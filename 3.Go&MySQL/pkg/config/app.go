package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "u6sl26lkhk6kyhke8oak:pscale_pw_yikMwWvMNlQTBITvYxIP6zTcvzK3OedgatMmIqwR0SI@tcp(aws.connect.psdb.cloud)/test?tls=true&interpolateParams=true")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db

}
