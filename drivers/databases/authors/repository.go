package authors

import (
	"api_short_story/business/authors"

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

	err := repo.db.Where("email = ? AND password = ?", authorDB.Email, authorDB.Password).First(&authorDB).Error
	if err != nil {
		return authors.AuthorEntity{}, err
	}
	return authorDB.ToAuthorEntity(), nil
}

func (repo *AuthorRepository) GetAllAuthors() ([]authors.AuthorEntity, error) {
	return []authors.AuthorEntity{}, nil
}

func (repo *AuthorRepository) AddAuthor(author authors.AuthorEntity) (authors.AuthorEntity, error) {
	authorDB := FromAuthorEntity(author)
	result := repo.db.Create(&authorDB)
	if result.Error != nil {
		return authors.AuthorEntity{}, result.Error
	}
	return authorDB.ToAuthorEntity(), nil
}
