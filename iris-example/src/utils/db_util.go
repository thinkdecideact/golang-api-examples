package utils

import (
	"fmt"
	"log"
	"thinkdecideact/src/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetDB(dbConfig *config.DBConfig) *gorm.DB {
	dbURI := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=Local",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Name,
		dbConfig.Charset)

	db, err := gorm.Open(mysql.Open(dbURI), &gorm.Config{})
	if err != nil {
		log.Fatal("Fail to connect database")
	}
	return db
}
