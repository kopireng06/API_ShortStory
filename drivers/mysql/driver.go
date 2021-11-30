package mysql

import (
	"api_short_story/app/config"
	"api_short_story/helpers"
	"api_short_story/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitialDb() *gorm.DB {
	configs := config.ReadJsonConfig()
	dsn := "" + configs.Database.Username + ":" + configs.Database.Password + "@tcp(" + configs.Database.Host + ")/" + configs.Database.Name + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&models.Author{}, &models.Category{}, &models.ShortStory{})
	passwordAdmin := helpers.HashPassword(configs.Admin.Password)
	db.Create(&models.Author{
		Name:     configs.Admin.Name,
		Email:    configs.Admin.Email,
		Password: passwordAdmin,
		Profile:  configs.Admin.Profile,
		Role:     0,
	})
	if err != nil {
		log.Fatal(err)
	}
	return db
}
