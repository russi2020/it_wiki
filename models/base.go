package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var db *gorm.DB //база данных

func init() {

	e := godotenv.Load() //Загрузить файл .env
	if e != nil {
		log.Println(e)
	}

	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	dbUri := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		dbHost, dbPort, username, dbName, password) //Создать строку подключения

	conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		log.Println(err)
	}

	db = conn
	db.Debug().AutoMigrate(&Account{}, &Contact{}, &Category{}, Theme{}, Question{}, Answer{}) //Миграция базы данных
}

// GetDB возвращает дескриптор объекта DB
func GetDB() *gorm.DB {
	return db
}
