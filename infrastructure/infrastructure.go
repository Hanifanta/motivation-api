package infrastructure

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"motivation-api/domain/entity"
)

func NewDatabaseConnection(dbDsn string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dbDsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	if err := db.AutoMigrate(&entity.Quote{}); err != nil {
		log.Fatal(err)
	}
	return db
}
