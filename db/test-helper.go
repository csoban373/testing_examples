package db

import (
	"net/url"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

func SetupTestDB(t *testing.T) *gorm.DB {
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
		t.Fatal("Failed to connect to database")
	}

	return db
}

func WithinTransaction(t *testing.T, db *gorm.DB, run func(t *testing.T, tx *gorm.DB)) {
	tx := db.Begin()
	if tx.Error != nil {
		t.Fatal("Failed to start transaction")
	}
	defer tx.Rollback()
	run(t, tx)
}
