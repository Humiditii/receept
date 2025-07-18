package database

import (
	"fmt"
	"log"
	"time"

	"github.com/Humiditii/receept/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database interface {
	CloseDBInstance()
}

type database struct {
	conn *gorm.DB
}

func (d *database) CloseDBInstance() {
	db, _ := d.conn.DB()

	log.Println("Database connection closing... at:->", time.Now())

	db.Close()
}

func NewDatabase(config config.ConfigI, models ...interface{}) *database {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", config.GetDbHost(), config.GetDbUser(), config.GetDbPassword(), config.GetDbName(), config.GetDbPort(), config.GetDbSslMode())

	fmt.Println(dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	log.Println("Database connection established... at:->", time.Now())

	// AutoMigrate all provided models
	if err = db.AutoMigrate(models...); err != nil {
		panic(err)
	}

	if len(models) > 0 {
		log.Println("Database automigration done! at:->", time.Now())
	}else {
		log.Println("No table to migrate at:->", time.Now())
	}

	return &database{
		conn: db,
	}
}
