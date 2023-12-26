package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

// DSN=kucpi4r8lmyedajtvbt4:pscale_pw_svmILsWoqOeguD4PytG81eDlU9HuIMoVgjb7N5YWeX3@tcp(aws.connect.psdb.cloud)/test?tls=true&interpolateParams=true

func Connect() {
	d, err := gorm.Open("mysql", "kucpi4r8lmyedajtvbt4:pscale_pw_svmILsWoqOeguD4PytG81eDlU9HuIMoVgjb7N5YWeX3@tcp(aws.connect.psdb.cloud)/test?tls=true&interpolateParams=true")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db

}
