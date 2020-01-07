package db

import (
	"fmt"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"os"
)


var db *gorm.DB


func init() {
	err := godotenv.Load() //Load .env file
	if err != nil {
		panic("can not load .env file")
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbPort := os.Getenv("db_port")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	sslmode := os.Getenv("db_sslmode")

	dbUri := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		dbHost, dbPort, dbName, username, password, sslmode)

	db, err = gorm.Open("postgres", dbUri)
	if err != nil {
		panic("can not connect to db")
	}

	driver, err := postgres.WithInstance(db.DB(), &postgres.Config{})
	if err != nil {
		fmt.Print("cannot find driver for db")
		return
	}

	m, err := migrate.NewWithDatabaseInstance("file://db/migrations", "avito-adv", driver)
	if err != nil {
		fmt.Print("cannot make migrations")
		return
	}

	err = m.Up()
	if err != nil {
		fmt.Print("cannot make migrations", err)
		return
	}
}

func GetDB() *gorm.DB {
	return db
}
