package db

import (
	"fmt"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)


var db *gorm.DB


func init() {
	db, _ = gorm.Open("postgres", "host=localhost port=5432 user=postgres sslmode=disable dbname=avito-adv password=postgres")
	//defer db.Close()

	driver, err := postgres.WithInstance(db.DB(), &postgres.Config{})
	if err != nil {
		fmt.Print("cannot find driver")
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
