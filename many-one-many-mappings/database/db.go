package database

import (
	"fmt"
	"github.com/NandanSatheesh/go-lang-playground/many-one-many-mappings/utils"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func GetDatabaseHandler() *gorm.DB {
	connectionString := fmt.Sprintf("%s://%s:%s@%s:%s/%s", DATABASE_VENDOR, USERNAME, PASSWORD, CONNECTION_URL, PORT_NUMBER, DATABASE_NAME)
	dbHandler, err := gorm.Open("postgres", connectionString)
	utils.CheckError(err)
	return dbHandler
}
