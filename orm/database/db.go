package database

import (
	"github.com/NandanSatheesh/go-lang-playground/orm/utils"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func GetDatabaseHandler() *gorm.DB {
	dbHandler, err := gorm.Open("postgres", "postgres://rcwhbkhp:YwxQ_ElK7Rm9hVMWz3DkjW0qxdKX9o2h@salt.db.elephantsql.com:5432/rcwhbkhp")
	utils.CheckError(err)
	return dbHandler
}
