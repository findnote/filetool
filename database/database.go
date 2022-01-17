package database

import (
	"mwp3000/models/system"

	log "github.com/pion/ion-log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB
var (
	DB_FILE_SQLITE = "cyp3000.db"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func SetupDatabase() {
	// c := config.GetConf()
	// var err error

	// // 初始化数据库链接
	dbSqlite, err := gorm.Open(sqlite.Open(DB_FILE_SQLITE), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db = dbSqlite

	log.Infof(">>dbSqlite: %v, err: %v", db, err)

	// 数据库迁移
	err = db.AutoMigrate(
		new(system.TDeviceInfo),
	)

	if err != nil {
		panic(err)
	}
}

func GetDataBase() *gorm.DB {
	return db
}
