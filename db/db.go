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


func Connect() {
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
		panic(fmt.Errorf("can not connect to db: %q", err))
	}

	driver, err := postgres.WithInstance(db.DB(), &postgres.Config{})
	if err != nil {
		fmt.Print(fmt.Errorf("cannot find driver for db: %q", err))
		return
	}

	m, err := migrate.NewWithDatabaseInstance("file://db/migrations", "avito-adv", driver)
	if err != nil {
		fmt.Print(fmt.Errorf("cannot make migrations: %q", err))
		return
	}
	if result := m.Up(); result != nil && result.Error() != "no change" {
		fmt.Print(fmt.Errorf("something get wrong: %q", result.Error()))
	}
}

func GetDB() *gorm.DB {
	return db
}
