package authors

import (
	"api_short_story/business/authors"
	"api_short_story/controllers/authors/response"
	"api_short_story/models"
	"errors"

	"gorm.io/gorm"
)

type AuthorRepository struct {
	db *gorm.DB
}

func NewAuthorRepository(gormDB *gorm.DB) authors.AuthorRepoInterface {
	return &AuthorRepository{
		db: gormDB,
	}
}

func (repo *AuthorRepository) Login(author authors.AuthorEntity) (authors.AuthorEntity, error) {
	authorDB := FromAuthorEntity(author)
	result := repo.db.Where("email = ?", authorDB.Email).First(&authorDB)
	if result.Error != nil {
		return authors.AuthorEntity{}, result.Error
	}
	if result.RowsAffected == 0 {
		return authors.AuthorEntity{}, errors.New("email not registered")
	}
	return authorDB.ToAuthorEntity(), nil
}

func (repo *AuthorRepository) GetAllAuthors() ([]authors.AuthorEntity, error) {
	authorDB := []response.Author{}
	result := repo.db.Where("role = ?", 1).Find(&authorDB)
	if result.Error != nil {
		return []authors.AuthorEntity{}, result.Error
	}
	if result.RowsAffected == 0 {
		return []authors.AuthorEntity{}, errors.New("record not found")
	}
	authorResponse := response.ToArrayAuthorEntity(authorDB)
	return authorResponse, result.Error
}

func (repo *AuthorRepository) GetAuthorById(id int) (authors.AuthorEntity, error) {
	authorDB := response.Author{}
	result := repo.db.Where("role = ?", 1).Find(&authorDB, id)
	if result.Error != nil {
		return authors.AuthorEntity{}, result.Error
	}
	if result.RowsAffected == 0 {
		return authorDB.ToAuthorEntity(), errors.New("record not found")
	}
	return authorDB.ToAuthorEntity(), result.Error
}

func (repo *AuthorRepository) GetAuthorsByName(name string) ([]authors.AuthorEntity, error) {
	authorDB := []response.Author{}
	result := repo.db.Where("name LIKE ?", "%"+name+"%").Where("role = ?", 1).Find(&authorDB)
	if result.Error != nil {
		return []authors.AuthorEntity{}, result.Error
	}
	if result.RowsAffected == 0 {
		return []authors.AuthorEntity{}, errors.New("record not found")
	}
	authorResponse := response.ToArrayAuthorEntity(authorDB)
	return authorResponse, result.Error
}

func (repo *AuthorRepository) AddAuthor(author authors.AuthorEntity) (authors.AuthorEntity, error) {
	authorDB := FromAuthorEntity(author)
	result := repo.db.Where("email = ?", authorDB.Email).First(&authorDB)
	if result.RowsAffected != 0 {
		return authors.AuthorEntity{}, errors.New("email has applied")
	}
	result = repo.db.Create(&authorDB)
	if result.Error != nil {
		return authors.AuthorEntity{}, result.Error
	}
	return authorDB.ToAuthorEntity(), nil
}

func (repo *AuthorRepository) EditAuthor(id int, author authors.AuthorEntity) (authors.AuthorEntity, error) {
	authorDB := Author{}
	result := repo.db.Find(&authorDB, id)
	if result.Error != nil {
		return authors.AuthorEntity{}, result.Error
	}
	if result.RowsAffected == 0 {
		return authorDB.ToAuthorEntity(), errors.New("record not found")
	}
	if authorDB.Id != author.Id {
		return authorDB.ToAuthorEntity(), errors.New("unauthorized")
	}
	authorDB.Email = author.Email
	authorDB.Name = author.Name
	authorDB.Password = author.Password
	authorDB.Profile = author.Profile
	result.Save(&authorDB)
	return authorDB.ToAuthorEntity(), nil
}

func (repo *AuthorRepository) DeleteAuthor(id int) (authors.AuthorEntity, error) {
	storyDB := models.ShortStory{AuthorID: uint(id)}
	result := repo.db.Where(&storyDB).Delete(&storyDB)
	if result.Error != nil {
		return authors.AuthorEntity{}, result.Error
	}
	authorDB := response.Author{Id: uint(id)}
	result = repo.db.Where(&authorDB).Delete(&authorDB)
	if result.Error != nil {
		return authors.AuthorEntity{}, result.Error
	}
	if result.RowsAffected == 0 {
		return authorDB.ToAuthorEntity(), errors.New("record not found")
	}
	return authorDB.ToAuthorEntity(), nil
}
