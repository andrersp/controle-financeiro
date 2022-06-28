package database

import (
	"log"

	"github.com/andrersp/controle-financeiro/src/core"
	"github.com/andrersp/controle-financeiro/src/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbConn *gorm.DB

func CreateDBConnection() error {
	if dbConn != nil {
		CloseConnection(dbConn)
	}

	db, err := gorm.Open(postgres.Open(core.DATABASE_URI), &gorm.Config{})

	if err != nil {
		return err
	}
	dbConn = db

	return err

}

func Connect() (*gorm.DB, error) {
	sqlDB, err := dbConn.DB()
	if err != nil {
		return nil, err
	}
	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}

	return dbConn, nil
}

func CloseConnection(conn *gorm.DB) {
	db, err := conn.DB()
	if err != nil {
		return
	}

	defer db.Close()
}

func MigrateTable() error {

	if err := dbConn.AutoMigrate(&models.Users{}); err != nil {
		return err
	}
	return nil
}

func SetupAPP() error {
	if err := CreateDBConnection(); err != nil {
		return err
	}

	if err := MigrateTable(); err != nil {
		log.Fatal(err)
	}

	return nil
}
