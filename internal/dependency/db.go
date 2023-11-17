package dependency

import (
	"activity-reporter/model"
	"activity-reporter/shared/helper"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	helper.LoadEnv()
	var (
		host     = os.Getenv("DB_HOST")
		port     = os.Getenv("DB_PORT")
		user     = os.Getenv("DB_USER")
		password = os.Getenv("DB_PASSWORD")
		dbname   = os.Getenv("DB_NAME")
	)
	// Getting and using a value from .env
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true})
	if err != nil {
		return nil
	}
	// // UNCOMMENT WHEN TRYING TO RESET DB
	// Down(db)

	db.AutoMigrate(model.User{}, model.Photo{}, model.Activity{}, model.UserPhoto{})

	return db
}

func Down(db *gorm.DB) {
	db.Migrator().DropTable(model.Activity{})
	db.Migrator().DropTable(model.Photo{})
	db.Migrator().DropTable(model.User{})
}