package shortstory

import (
	shortstorybussines "api_short_story/business/short_story"
	"api_short_story/controllers/short_story/response"
	"errors"

	"gorm.io/gorm"
)

type ShortStoryRepository struct {
	db *gorm.DB
}

func NewShortStoryRepository(gormDB *gorm.DB) shortstorybussines.ShortStoryRepoInterface {
	return &ShortStoryRepository{
		db: gormDB,
	}
}

func (repo *ShortStoryRepository) GetShortStories(offset int, limit int) ([]shortstorybussines.ShortStoryEntity, error) {
	storyDB := []response.ShortStory{}
	result := repo.db.Offset(offset).Limit(limit).Find(&storyDB)
	if result.Error != nil {
		return []shortstorybussines.ShortStoryEntity{}, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("record not found")
	}
	storyResponse := response.ToArrayShortStoryEntity(storyDB)
	return storyResponse, result.Error
}
func (repo *ShortStoryRepository) GetShortStoryById(id int) (shortstorybussines.ShortStoryEntity, error) {
	storyDB := response.ShortStory{}
	result := repo.db.Find(&storyDB, id)
	if result.Error != nil {
		return shortstorybussines.ShortStoryEntity{}, result.Error
	}
	if result.RowsAffected == 0 {
		return storyDB.ToShortStoryEntity(), errors.New("record not found")
	}
	return storyDB.ToShortStoryEntity(), result.Error
}
func (repo *ShortStoryRepository) AddShortStory(story shortstorybussines.ShortStoryEntity) (shortstorybussines.ShortStoryEntity, error) {
	storyDB := FromAuthorEntity(story)
	result := repo.db.Create(&storyDB)
	if result.Error != nil {
		return shortstorybussines.ShortStoryEntity{}, result.Error
	}
	return storyDB.ToShortStoryEntity(), nil
}
func (repo *ShortStoryRepository) EditShortStory(id int, story shortstorybussines.ShortStoryEntity) (shortstorybussines.ShortStoryEntity, error) {
	storyDB := response.ShortStory{}
	result := repo.db.Find(&storyDB, id)
	if result.Error != nil {
		return shortstorybussines.ShortStoryEntity{}, result.Error
	}
	if result.RowsAffected == 0 {
		return shortstorybussines.ShortStoryEntity{}, errors.New("record not found")
	}
	storyDB.Title = story.Title
	storyDB.Content = story.Content
	storyDB.AuthorID = story.AuthorID
	storyDB.CategoryID = story.CategoryID
	result.Save(&storyDB)
	return storyDB.ToShortStoryEntity(), nil
}

func (repo *ShortStoryRepository) DeleteShortStory(id int) (shortstorybussines.ShortStoryEntity, error) {
	storyDB := response.ShortStory{Id: uint(id)}
	result := repo.db.Where(&storyDB).Delete(&storyDB)
	if result.Error != nil {
		return shortstorybussines.ShortStoryEntity{}, result.Error
	}
	if result.RowsAffected == 0 {
		return shortstorybussines.ShortStoryEntity{}, errors.New("record not found")
	}
	return storyDB.ToShortStoryEntity(), nil
}
