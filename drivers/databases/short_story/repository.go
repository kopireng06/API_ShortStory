package shortstory

import (
	shortstorybussines "api_short_story/business/short_story"
	"api_short_story/controllers/short_story/request"
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
	result := repo.db.Offset(offset).Limit(limit).Preload("Author").Preload("Category").Find(&storyDB)
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
	result := repo.db.Preload("Author").Preload("Category").Find(&storyDB, id)
	if result.Error != nil {
		return shortstorybussines.ShortStoryEntity{}, result.Error
	}
	if result.RowsAffected == 0 {
		return storyDB.ToShortStoryEntity(), errors.New("record not found")
	}
	return storyDB.ToShortStoryEntity(), result.Error
}

func (repo *ShortStoryRepository) GetShortStoriesByTitle(title string) ([]shortstorybussines.ShortStoryEntity, error) {
	storyDB := []response.ShortStory{}
	result := repo.db.Where("title LIKE ?", "%"+title+"%").Preload("Author").Preload("Category").Find(&storyDB)
	if result.Error != nil {
		return []shortstorybussines.ShortStoryEntity{}, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("record not found")
	}
	storyResponse := response.ToArrayShortStoryEntity(storyDB)
	return storyResponse, result.Error
}

func (repo *ShortStoryRepository) GetShortStoriesByIdCategory(id int) ([]shortstorybussines.ShortStoryEntity, error) {
	storyDB := []response.ShortStory{}
	result := repo.db.Where("category_id = ?", id).Preload("Author").Preload("Category").Find(&storyDB)
	if result.Error != nil {
		return []shortstorybussines.ShortStoryEntity{}, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("record not found")
	}
	storyResponse := response.ToArrayShortStoryEntity(storyDB)
	return storyResponse, result.Error
}

func (repo *ShortStoryRepository) GetShortStoriesByIdAuthor(id int) ([]shortstorybussines.ShortStoryEntity, error) {
	storyDB := []response.ShortStory{}
	result := repo.db.Where("author_id = ?", id).Preload("Author").Preload("Category").Find(&storyDB)
	if result.Error != nil {
		return []shortstorybussines.ShortStoryEntity{}, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("record not found")
	}
	storyResponse := response.ToArrayShortStoryEntity(storyDB)
	return storyResponse, result.Error
}

func (repo *ShortStoryRepository) AddShortStory(story shortstorybussines.ShortStoryEntity) (shortstorybussines.ShortStoryEntity, error) {
	storyDB := FromAuthorEntity(story)
	result := repo.db.Create(&storyDB)
	if result.Error != nil {
		return shortstorybussines.ShortStoryEntity{}, result.Error
	}
	storyResponse := response.ShortStory{}
	repo.db.Preload("Author").Preload("Category").Find(&storyResponse, storyDB.Id)
	return storyResponse.ToShortStoryEntity(), nil
}

func (repo *ShortStoryRepository) EditShortStory(id int, story shortstorybussines.ShortStoryEntity) (shortstorybussines.ShortStoryEntity, error) {
	storyDB := response.ShortStory{}
	result := repo.db.Preload("Author").Preload("Category").Find(&storyDB, id)
	if story.AuthorID != storyDB.AuthorID {
		return shortstorybussines.ShortStoryEntity{}, errors.New("unauthorized")
	}
	if result.Error != nil {
		return shortstorybussines.ShortStoryEntity{}, result.Error
	}
	if result.RowsAffected == 0 {
		return shortstorybussines.ShortStoryEntity{}, errors.New("record not found")
	}
	storyDB.Title = story.Title
	storyDB.Content = story.Content
	storyDB.CategoryID = story.CategoryID
	result.Save(&storyDB)
	return storyDB.ToShortStoryEntity(), nil
}

func (repo *ShortStoryRepository) DeleteShortStory(id int, story shortstorybussines.ShortStoryEntity) (shortstorybussines.ShortStoryEntity, error) {
	storyDB := request.ShortStory{Id: uint(id)}
	result := repo.db.Find(&storyDB, id)
	if result.Error != nil {
		return shortstorybussines.ShortStoryEntity{}, result.Error
	}
	if result.RowsAffected == 0 {
		return shortstorybussines.ShortStoryEntity{}, errors.New("record not found")
	}
	if story.AuthorID != storyDB.AuthorID {
		return shortstorybussines.ShortStoryEntity{}, errors.New("unauthorized")
	}
	repo.db.Where(&storyDB).Delete(&storyDB)
	return storyDB.ToShortStoryEntity(), nil
}
