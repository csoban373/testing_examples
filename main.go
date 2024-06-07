package main

import (
	"fmt"

	"github.com/csoban373/testing_examples/db"
	"github.com/csoban373/testing_examples/model"
	"gorm.io/gorm"
)

func main() {
	db, err := db.SetupDB()
	if err != nil {
		return
	}
	fmt.Println(GetPeople(db))
}

func GetPeople(db *gorm.DB) (result []model.Person, err error) {
	err = db.Find(&result).Error
	return
}
