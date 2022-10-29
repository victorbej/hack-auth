package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB //база данных

func init() {

	//e := godotenv.Load() //Загрузить файл .env
	//if e != nil {
	//	fmt.Print(e)
	//}

	//username := os.Getenv("DB_USER")
	//password := os.Getenv("DB_PASS")
	//dbName := os.Getenv("DB_NAME")
	//dbHost := os.Getenv("DB_HOST")
	//username := "postgres"
	//password := "1234"
	//dbName := "postgres"
	//dbHost := "localhost"

	//dbUri := fmt.Sprintf("host=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, username, dbName, password) //Создать строку подключения

	dbUri := fmt.Sprintf("postgres://jehoyfskadjiwk:c79ff9cab606cc84fbb2b7138257d10abb3937de4faa805962852e11539bae5f@ec2-23-20-140-229.compute-1.amazonaws.com:5432/d6rc58i6vp57i1")

	fmt.Println(dbUri)

	conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		fmt.Print(err)
	}

	db = conn
	db.Debug().AutoMigrate(&Account{}, &UserTable{}) //Миграция базы данных
}

// возвращает дескриптор объекта DB
func GetDB() *gorm.DB {
	return db
}
