package db

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB

func Connect() error {
	var err error
	dns := fmt.Sprintf("%s&parseTime=True", os.Getenv("DSN"))

	Database, err = gorm.Open(
		mysql.Open(dns),
	)

	if err != nil {
		fmt.Printf("err")
	}
	fmt.Println("Successfully connected to PlanetScale")
	return err
}
