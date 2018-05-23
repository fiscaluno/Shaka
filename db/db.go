package db

import (
	"github.com/fiscaluno/shaka/util"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Conn connect on SQL Data Base Gorm
func Conn() (db *gorm.DB) {

	typeDB := util.GetOSEnvironment("DB", "sqlite3")

	pathDB := util.GetOSEnvironment("DATABASE_URL", "test.db")

	db, err := gorm.Open(typeDB, pathDB)
	if err != nil {
		panic("failed to connect database")
	}

	return db
}
