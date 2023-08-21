package mysql

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	_dbName     = "MYSQL_DATABASE"
	_dbHost     = "MYSQL_HOST"
	_dbUsername = "MYSQL_USER"
	_dbPassword = "MYSQL_PASSWORD"
)

func GetMySQLConnection() *gorm.DB {
	// open db dconection
	dbHost := os.Getenv(_dbHost)
	dbName := os.Getenv(_dbName)
	dbUsername := os.Getenv(_dbUsername)
	dbPassword := os.Getenv(_dbPassword)

	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True",
		dbUsername, dbPassword, dbHost, dbName)

	db, err := gorm.Open(mysql.Open(connectionString),
		&gorm.Config{
			//Logger: logger.Default.LogMode(logger.Info),
		},
	)

	if err != nil {
		log.Fatal("cannot create mysql connection: ", err)
	}

	return db
}
