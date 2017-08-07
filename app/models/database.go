package models

import (
	"log"

	aah "aahframework.org/aah.v0"
	"github.com/jinzhu/gorm"

	// _ postgres
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db *gorm.DB
)

func initDb(_ *aah.Event) {
	var err error

	db, err = gorm.Open("postgres", "host=localhost user=postgres dbname=test sslmode=disable password=postgres")
	if err != nil {
		log.Fatalf("Got error when connect database, the error is '%v'", err)
	}

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	db.LogMode(true)

	initDbUser()
}

func closeDb(_ *aah.Event) {
	if db != nil {
		_ = db.Close()
	}
}

func init() {
	aah.OnStart(initDb)
	aah.OnShutdown(closeDb)
}
