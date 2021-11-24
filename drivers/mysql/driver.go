package mysql

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitialDb() *gorm.DB {
	// dsn := "admin:Opangkers02@tcp(database-1.csncfmrvzixw.us-east-2.rds.amazonaws.com:3306)/api_short_story?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "root:@tcp(127.0.0.1:3306)/api_short_story?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}
