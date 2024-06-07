package db

import (
	"net/url"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

func SetupDB() (*gorm.DB, error) {
	var (
		userName = "root"
		password = ""
		host     = "localhost"
		port     = "3306"
		dbName   = "test"
	)

	// connect to the database
	connStr := url.QueryEscape(userName) + ":" + url.QueryEscape(password) + "@tcp(" + host + ":" + port + ")/" + dbName

	db, err := gorm.Open(mysql.Open(connStr), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Info),
	})
	if err != nil {
		return nil, err
	}

	return db, nil
}
