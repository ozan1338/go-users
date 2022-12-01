package database

import (
	"fmt"
	"go-users/src/models"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

const (
	DB_HOST="DB_HOST"
	DB_PORT="DB_PORT"
	DB_USERNAME="DB_USERNAME"
	DB_PASSWORD="DB_PASSWORD"
	DB_DATABASE="DB_DATABASE"
)

var (
	db_host = os.Getenv(DB_HOST)
	db_port = os.Getenv(DB_PORT)
	db_username = os.Getenv(DB_USERNAME)
	db_password = os.Getenv(DB_PASSWORD)
	db_database = os.Getenv(DB_DATABASE)
)

func Connect() {
	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", db_username, db_password, db_host, db_port, db_database)
	fmt.Println(dsn)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Could not connect with the database!" + err.Error())
	}
}

func AutoMigrate() {
	DB.AutoMigrate(models.User{},models.UserToken{})
}
