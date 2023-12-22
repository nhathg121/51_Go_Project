package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "po2ph09ofv26fm1g9ovw:pscale_pw_Und4Q2TR9pHfUGerOtK8pqowitTmrW9t7nIpCdC4fBc@tcp(aws.connect.psdb.cloud)/test?tls=true&interpolateParams=true")
	if err != nil {
		panic(err)
	}

	db = d
}

func GetDB() *gorm.DB {
	return db

}
