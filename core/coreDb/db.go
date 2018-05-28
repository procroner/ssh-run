package coreDb

import (
	"github.com/joho/godotenv"
	"os"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/procroner/ssh-run/core/coreError"
	"github.com/jinzhu/gorm"
)

func Connect() *gorm.DB {
	err := godotenv.Load()
	coreError.HandleError("load env file", err)

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	connectString := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbName)
	db, err := gorm.Open("mysql", connectString)
	coreError.HandleError("open db", err)

	return db
}
