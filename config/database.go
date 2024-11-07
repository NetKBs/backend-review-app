package config

import (
	"log"
	"os"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	DATABASE := os.Getenv("TURSO_DATABASE_URL")
	TOKEN := os.Getenv("TURSO_AUTH_TOKEN")
	var error error

	if DATABASE == "" || TOKEN == "" {
		DB, error = gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
		log.Default().Println("Token or Database URL is empty. Using local database.")
	} else {
		url := DATABASE + "?auth_token=" + TOKEN
		DB, error = gorm.Open(sqlite.New((sqlite.Config{
			DSN:        url,
			DriverName: "libsql",
		})), &gorm.Config{})
	}

	if error != nil {
		log.Fatal(error)
	}
}

func SyncDB() {
	models := []interface{}{}

	DB.AutoMigrate(models...)
}
