package models

import (
	"mwp3000/database"

	"gorm.io/gorm"
)

func GetDB() *gorm.DB {
	return database.GetDataBase()
}
