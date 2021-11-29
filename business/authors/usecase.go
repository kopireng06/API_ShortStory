package authors

import (
	"api_short_story/helpers"
	"api_short_story/middlewares/token"
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

type AuthorUseCase struct {
	repo AuthorRepoInterface
}

func NewUseCase(authorRepo AuthorRepoInterface) AuthorUseCaseInterface {
	return &AuthorUseCase{
		repo: authorRepo,
	}
}

func (usecase *AuthorUseCase) Login(author AuthorEntity) (AuthorEntity, error) {
	if author.Email == "" {
		return AuthorEntity{}, errors.New("email empty")
	}
	if author.Password == "" {
		return AuthorEntity{}, errors.New("password empty")
	}
	password := author.Password
	author, err := usecase.repo.Login(author)
	if err != nil {
		return AuthorEntity{}, err
	}
	if helpers.CheckSamePassword(password, author.Password) {
		jwtClaims := token.JwtClaims{
			int(author.Id),
			author.Name,
			author.Email,
			author.Role,
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			},
		}
		author.Token = token.GenerateJWT(jwtClaims)
		return author, nil
	}
	return AuthorEntity{}, errors.New("wrong password")
}

func (usecase *AuthorUseCase) GetAllAuthors() ([]AuthorEntity, error) {
	authors, err := usecase.repo.GetAllAuthors()
	if err != nil {
		return authors, err
	}
	return authors, nil
}

func (usecase *AuthorUseCase) GetAuthorById(id int) (AuthorEntity, error) {
	author, err := usecase.repo.GetAuthorById(id)
	if err != nil {
		return author, err
	}
	return author, nil
}

func (usecase *AuthorUseCase) GetAuthorsByName(name string) ([]AuthorEntity, error) {
	author, err := usecase.repo.GetAuthorsByName(name)
	if err != nil {
		return author, err
	}
	return author, nil
}

func (usecase *AuthorUseCase) AddAuthor(author AuthorEntity) (AuthorEntity, error) {

	if author.Name == "" {
		return AuthorEntity{}, errors.New("name empty")
	}
	if author.Email == "" {
		return AuthorEntity{}, errors.New("email empty")
	}
	if author.Password == "" {
		return AuthorEntity{}, errors.New("password empty")
	}
	if author.ConfirmPassword == "" {
		return AuthorEntity{}, errors.New("confirm password empty")
	}
	if author.Profile == "" {
		return AuthorEntity{}, errors.New("profile empty")
	}
	if author.ConfirmPassword != author.Password {
		return AuthorEntity{}, errors.New("password must same with confirm password")
	}
	hashedPassword := helpers.HashPassword(author.Password)
	author.Password = hashedPassword
	author.Role = 1
	author, err := usecase.repo.AddAuthor(author)
	if err != nil {
		return AuthorEntity{}, err
	}
	return author, nil
}

func (usecase *AuthorUseCase) EditAuthor(id int, author AuthorEntity) (AuthorEntity, error) {
	if author.Name == "" {
		return AuthorEntity{}, errors.New("name empty")
	}
	if author.Email == "" {
		return AuthorEntity{}, errors.New("email empty")
	}
	if author.Password == "" {
		return AuthorEntity{}, errors.New("password empty")
	}
	if author.Profile == "" {
		return AuthorEntity{}, errors.New("profile empty")
	}
	hashedPassword := helpers.HashPassword(author.Password)
	author.Password = hashedPassword
	author, err := usecase.repo.EditAuthor(id, author)
	if err != nil {
		return AuthorEntity{}, err
	}
	return author, nil
}

func (usecase *AuthorUseCase) DeleteAuthor(id int) (AuthorEntity, error) {
	author, err := usecase.repo.DeleteAuthor(id)
	if err != nil {
		return AuthorEntity{}, err
	}
	return author, nil
}
