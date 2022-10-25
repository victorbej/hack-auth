package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var db *gorm.DB //база данных

func init() {

	e := godotenv.Load() //Загрузить файл .env
	if e != nil {
		fmt.Print(e)
	}

	//username := os.Getenv("DB_USER")
	//password := os.Getenv("DB_PASS")
	//dbName := os.Getenv("DB_NAME")
	//dbHost := os.Getenv("DB_HOST")

	//dbUri := fmt.Sprintf("host=%s user=%s dbname=%s password=%s", dbHost, username, dbName, password) //Создать строку подключения

	dbUri := fmt.Sprintf("postgres://mzodgzxneleony:339c90b16d75e73c4001c1cb32d5efa64be126649ab91f33e644d4698e7369b1@ec2-54-227-248-71.compute-1.amazonaws.com:5432/d8f5abl5k4f51n")
	fmt.Println(dbUri)

	conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		fmt.Print(err)
	}

	db = conn
	db.Debug().AutoMigrate(&Account{}) //Миграция базы данных
}

// возвращает дескриптор объекта DB
func GetDB() *gorm.DB {
	return db
}
