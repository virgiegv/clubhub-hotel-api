package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
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
		dbURL := "postgres://postgres:1234@localhost:5432/clubhub"
		//TO DO:    dsn := "host=localhost user=gorm password=gorm dbname=gorm port=5432 sslmode=disable TimeZone=Asia/Shanghai" instead?

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
