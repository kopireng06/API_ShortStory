package configs

import (
	"api_short_story/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "root:@tcp(database-1.csncfmrvzixw.us-east-2.rds.amazonaws.com:3306)/api_short_story?charset=utf8mb4&parseTime=True&loc=Local"
	// dsn := "root:@tcp(127.0.0.1:3306)/api_short_story?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Database tidak connect")
	}
	migrate()
}

func migrate() {
	DB.AutoMigrate(models.Author{}, models.Category{}, models.ShortStory{})
}
