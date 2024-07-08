package models

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"sync"
	"time"
)

var once sync.Once

type DBConnection struct {
	DB *gorm.DB
}

var dbInstance *DBConnection

func Init() *DBConnection {
	once.Do(func() {

		dbHost := os.Getenv("DB_HOST")
		dbPort := os.Getenv("DB_PORT")
		dbUser := os.Getenv("DB_USER")
		dbPassword := os.Getenv("DB_PASSWORD")
		dbName := os.Getenv("DB_NAME")

		//dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		//	dbHost, dbPort, dbUser, dbPassword, dbName)

		dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

		db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{
			NowFunc: func() time.Time {
				return time.Now().Local()
			},
		})

		if err != nil {
			log.Fatalln(err)
		}

		db.AutoMigrate(
			&City{}, &Location{}, &Owner{}, &Company{}, &FranchiseWebEndpoint{}, &FranchiseWebSite{}, &Franchise{},
		)

		sqlDB, err := db.DB()
		sqlDB.SetMaxOpenConns(50)
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetConnMaxLifetime(time.Hour)

		dbInstance = &DBConnection{DB: db}

	})

	return dbInstance

}
