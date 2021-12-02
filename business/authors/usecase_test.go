package authors_test

import (
	"api_short_story/business/authors"
	"api_short_story/business/authors/mocks"
	"api_short_story/helpers"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var authorRepoInterfaceMock mocks.AuthorRepoInterface
var authorUseCaseInterface authors.AuthorUseCaseInterface
var authorResponseFromRepo authors.AuthorEntity
var authorArrayResponseFromRepo []authors.AuthorEntity

func setupLogin() {
	authorUseCaseInterface = authors.NewUseCase(&authorRepoInterfaceMock)
	password := helpers.HashPassword("123")
	authorResponseFromRepo = authors.AuthorEntity{
		Id:       1,
		Name:     "Alterra",
		Email:    "alterra@gmail.com",
		Password: password,
		Role:     0,
	}
	// jwtClaims := token.JwtClaims{
	// 	int(authorResponseFromRepo.Id),
	// 	authorResponseFromRepo.Name,
	// 	authorResponseFromRepo.Email,
	// 	authorResponseFromRepo.Role,
	// 	jwt.StandardClaims{
	// 		ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
	// 	},
	// }
	// authorResponseFromRepo.Token = token.GenerateJWT(jwtClaims)
}

func setupAddAuthor() {
	authorUseCaseInterface = authors.NewUseCase(&authorRepoInterfaceMock)
	password := helpers.HashPassword("keong")
	authorResponseFromRepo = authors.AuthorEntity{
		Id:       1,
		Name:     "Alterra",
		Email:    "alterra@gmail.com",
		Password: password,
		Profile:  "story lovers",
	}
}

func setupEditAuthor() {
	authorUseCaseInterface = authors.NewUseCase(&authorRepoInterfaceMock)
	password := helpers.HashPassword("keong")
	authorResponseFromRepo = authors.AuthorEntity{
		Id:       1,
		Name:     "Alterra",
		Email:    "alterra@gmail.com",
		Password: password,
		Profile:  "story lovers",
	}
}

func setupGetArrayResponseFromRepo() {
	authorUseCaseInterface = authors.NewUseCase(&authorRepoInterfaceMock)
	password := helpers.HashPassword("keong")
	authorArrayResponseFromRepo = []authors.AuthorEntity{
		{
			Id:       1,
			Name:     "Alterra",
			Email:    "alterra@gmail.com",
			Password: password,
			Profile:  "story lovers",
		},
	}
}

func setupGetResponseFromRepo() {
	authorUseCaseInterface = authors.NewUseCase(&authorRepoInterfaceMock)
	password := helpers.HashPassword("keong")
	authorResponseFromRepo = authors.AuthorEntity{
		Id:       1,
		Name:     "Alterra",
		Email:    "alterra@gmail.com",
		Password: password,
		Profile:  "story lovers",
	}
}

func setupDeleteAuthor() {
	authorUseCaseInterface = authors.NewUseCase(&authorRepoInterfaceMock)
	authorResponseFromRepo = authors.AuthorEntity{
		Id: 1,
	}
}

func TestLogin(t *testing.T) {
	setupLogin()
	t.Run("Success Login", func(t *testing.T) {
		authorRepoInterfaceMock.On("Login", mock.AnythingOfType("authors.AuthorEntity"), mock.Anything).Return(authorResponseFromRepo, nil).Once()
		var requestLoginEntity = authors.AuthorEntity{
			Email:    "alterra@gmail.com",
			Password: "123",
		}
		entity, err := authorUseCaseInterface.Login(requestLoginEntity)
		authorResponseFromRepo.Token = entity.Token
		assert.Equal(t, nil, err)
		assert.Equal(t, authorResponseFromRepo, entity)
	})

	t.Run("Fail Login", func(t *testing.T) {
		authorRepoInterfaceMock.On("Login", mock.AnythingOfType("authors.AuthorEntity"), mock.Anything).Return(authors.AuthorEntity{}, errors.New("error bro")).Once()
		var requestLoginEntity = authors.AuthorEntity{
			Email:    "alterra@gmail.com",
			Password: "123",
		}
		entity, err := authorUseCaseInterface.Login(requestLoginEntity)
		assert.Equal(t, "error bro", err.Error())
		assert.Equal(t, authors.AuthorEntity{}, entity)
	})

	t.Run("Login with Email Empty", func(t *testing.T) {
		var requestLoginEntity = authors.AuthorEntity{
			Email:    "",
			Password: "123",
		}
		entity, err := authorUseCaseInterface.Login(requestLoginEntity)
		assert.Equal(t, "email empty", err.Error())
		assert.Equal(t, authors.AuthorEntity{}, entity)
	})

	t.Run("Login with Password Empty", func(t *testing.T) {
		var requestLoginEntity = authors.AuthorEntity{
			Email:    "alterra@gmail.com",
			Password: "",
		}
		entity, err := authorUseCaseInterface.Login(requestLoginEntity)
		assert.Equal(t, "password empty", err.Error())
		assert.Equal(t, authors.AuthorEntity{}, entity)
	})

	t.Run("Login with Wrong Password", func(t *testing.T) {
		authorRepoInterfaceMock.On("Login", mock.AnythingOfType("authors.AuthorEntity"), mock.Anything).Return(authorResponseFromRepo, nil).Once()
		var requestLoginEntity = authors.AuthorEntity{
			Email:    "alterra@gmail.com",
			Password: "keoooong",
		}
		entity, err := authorUseCaseInterface.Login(requestLoginEntity)
		assert.Equal(t, "wrong password", err.Error())
		assert.Equal(t, authors.AuthorEntity{}, entity)
	})
}

func TestAddAuthor(t *testing.T) {
	setupAddAuthor()
	t.Run("Success Add Author", func(t *testing.T) {
		authorRepoInterfaceMock.On("AddAuthor", mock.AnythingOfType("authors.AuthorEntity"), mock.Anything).Return(authorResponseFromRepo, nil).Once()
		var requestAddAuthorEntity = authors.AuthorEntity{
			Id:              1,
			Name:            "Alterra",
			Email:           "alterra@gmail.com",
			Password:        "keong",
			ConfirmPassword: "keong",
			Profile:         "story lovers",
		}
		entity, err := authorUseCaseInterface.AddAuthor(requestAddAuthorEntity)
		assert.Equal(t, nil, err)
		assert.Equal(t, authorResponseFromRepo, entity)
	})

	t.Run("Add Author with Name Empty", func(t *testing.T) {
		var requestAddAuthorEntity = authors.AuthorEntity{
			Id:              1,
			Name:            "",
			Email:           "alterra@gmail.com",
			Password:        "keong",
			ConfirmPassword: "keong",
			Profile:         "story lovers",
		}
		entity, err := authorUseCaseInterface.AddAuthor(requestAddAuthorEntity)
		assert.Equal(t, "name empty", err.Error())
		assert.Equal(t, authors.AuthorEntity{}, entity)
	})

	t.Run("Add Author with Email Empty", func(t *testing.T) {
		var requestAddAuthorEntity = authors.AuthorEntity{
			Id:              1,
			Name:            "alterra",
			Email:           "",
			Password:        "keong",
			ConfirmPassword: "keong",
			Profile:         "story lovers",
		}
		entity, err := authorUseCaseInterface.AddAuthor(requestAddAuthorEntity)
		assert.Equal(t, "email empty", err.Error())
		assert.Equal(t, authors.AuthorEntity{}, entity)
	})

	t.Run("Add Author with Password Empty", func(t *testing.T) {
		var requestAddAuthorEntity = authors.AuthorEntity{
			Id:              1,
			Name:            "alterra",
			Email:           "alterra@gmail.com",
			Password:        "",
			ConfirmPassword: "keong",
			Profile:         "story lovers",
		}
		entity, err := authorUseCaseInterface.AddAuthor(requestAddAuthorEntity)
		assert.Equal(t, "password empty", err.Error())
		assert.Equal(t, authors.AuthorEntity{}, entity)
	})

	t.Run("Add Author with ConfirmPassword Empty", func(t *testing.T) {
		var requestAddAuthorEntity = authors.AuthorEntity{
			Id:              1,
			Name:            "alterra",
			Email:           "alterra@gmail.com",
			Password:        "keong",
			ConfirmPassword: "",
			Profile:         "story lovers",
		}
		entity, err := authorUseCaseInterface.AddAuthor(requestAddAuthorEntity)
		assert.Equal(t, "confirm password empty", err.Error())
		assert.Equal(t, authors.AuthorEntity{}, entity)
	})

	t.Run("Add Author with Profile Empty", func(t *testing.T) {
		var requestAddAuthorEntity = authors.AuthorEntity{
			Id:              1,
			Name:            "alterra",
			Email:           "alterra@gmail.com",
			Password:        "keong",
			ConfirmPassword: "keong",
			Profile:         "",
		}
		entity, err := authorUseCaseInterface.AddAuthor(requestAddAuthorEntity)
		assert.Equal(t, "profile empty", err.Error())
		assert.Equal(t, authors.AuthorEntity{}, entity)
	})

	t.Run("Add Author with Different Password and Confirm Password", func(t *testing.T) {
		var requestAddAuthorEntity = authors.AuthorEntity{
			Id:              1,
			Name:            "alterra",
			Email:           "alterra@gmail.com",
			Password:        "keong",
			ConfirmPassword: "keong2",
			Profile:         "story lovers",
		}
		entity, err := authorUseCaseInterface.AddAuthor(requestAddAuthorEntity)
		assert.Equal(t, "password must same with confirm password", err.Error())
		assert.Equal(t, authors.AuthorEntity{}, entity)
	})

	t.Run("Fail Add Author", func(t *testing.T) {
		authorRepoInterfaceMock.On("AddAuthor", mock.AnythingOfType("authors.AuthorEntity"), mock.Anything).Return(authors.AuthorEntity{}, errors.New("error bro")).Once()
		var requestAddAuthorEntity = authors.AuthorEntity{
			Id:              1,
			Name:            "Alterra",
			Email:           "alterra@gmail.com",
			Password:        "keong",
			ConfirmPassword: "keong",
			Profile:         "story lovers",
		}
		entity, err := authorUseCaseInterface.AddAuthor(requestAddAuthorEntity)
		assert.Equal(t, "error bro", err.Error())
		assert.Equal(t, authors.AuthorEntity{}, entity)
	})
}

func TestEditAuthor(t *testing.T) {
	setupEditAuthor()
	t.Run("Success Edit Author", func(t *testing.T) {
		authorRepoInterfaceMock.On("EditAuthor", 1, mock.AnythingOfType("authors.AuthorEntity"), mock.Anything).Return(authorResponseFromRepo, nil).Once()
		var requestEditAuthorEntity = authors.AuthorEntity{
			Id:       1,
			Name:     "Alterra",
			Email:    "alterra@gmail.com",
			Password: "keong",
			Profile:  "story lovers",
		}
		entity, err := authorUseCaseInterface.EditAuthor(int(requestEditAuthorEntity.Id), requestEditAuthorEntity)
		assert.Equal(t, nil, err)
		assert.Equal(t, authorResponseFromRepo, entity)
	})

	t.Run("Edit Author with Name Empty", func(t *testing.T) {
		var requestEditAuthorEntity = authors.AuthorEntity{
			Id:              1,
			Name:            "",
			Email:           "alterra@gmail.com",
			Password:        "keong",
			ConfirmPassword: "keong",
			Profile:         "story lovers",
		}
		entity, err := authorUseCaseInterface.EditAuthor(int(requestEditAuthorEntity.Id), requestEditAuthorEntity)
		assert.Equal(t, "name empty", err.Error())
		assert.Equal(t, authors.AuthorEntity{}, entity)
	})

	t.Run("Edit Author with Email Empty", func(t *testing.T) {
		var requestEditAuthorEntity = authors.AuthorEntity{
			Id:              1,
			Name:            "alterra",
			Email:           "",
			Password:        "keong",
			ConfirmPassword: "keong",
			Profile:         "story lovers",
		}
		entity, err := authorUseCaseInterface.EditAuthor(int(requestEditAuthorEntity.Id), requestEditAuthorEntity)
		assert.Equal(t, "email empty", err.Error())
		assert.Equal(t, authors.AuthorEntity{}, entity)
	})

	t.Run("Edit Author with Password Empty", func(t *testing.T) {
		var requestEditAuthorEntity = authors.AuthorEntity{
			Id:              1,
			Name:            "alterra",
			Email:           "alterra@gmail.com",
			Password:        "",
			ConfirmPassword: "keong",
			Profile:         "story lovers",
		}
		entity, err := authorUseCaseInterface.EditAuthor(int(requestEditAuthorEntity.Id), requestEditAuthorEntity)
		assert.Equal(t, "password empty", err.Error())
		assert.Equal(t, authors.AuthorEntity{}, entity)
	})

	t.Run("Edit Author with Profile Empty", func(t *testing.T) {
		var requestEditAuthorEntity = authors.AuthorEntity{
			Id:              1,
			Name:            "alterra",
			Email:           "alterra@gmail.com",
			Password:        "keong",
			ConfirmPassword: "keong",
			Profile:         "",
		}
		entity, err := authorUseCaseInterface.EditAuthor(int(requestEditAuthorEntity.Id), requestEditAuthorEntity)
		assert.Equal(t, "profile empty", err.Error())
		assert.Equal(t, authors.AuthorEntity{}, entity)
	})

	t.Run("Fail Edit Author", func(t *testing.T) {
		var requestAddAuthorEntity = authors.AuthorEntity{
			Id:              1,
			Name:            "Alterra",
			Email:           "alterra@gmail.com",
			Password:        "keong",
			ConfirmPassword: "keong",
			Profile:         "story lovers",
		}
		authorRepoInterfaceMock.On("EditAuthor", 1, mock.AnythingOfType("authors.AuthorEntity"), mock.Anything).Return(authors.AuthorEntity{}, errors.New("error bro")).Once()
		entity, err := authorUseCaseInterface.EditAuthor(int(requestAddAuthorEntity.Id), requestAddAuthorEntity)
		assert.Equal(t, "error bro", err.Error())
		assert.Equal(t, authors.AuthorEntity{}, entity)
	})
}

func TestGetAllAuthors(t *testing.T) {
	setupGetArrayResponseFromRepo()
	t.Run("Success Get All Authors", func(t *testing.T) {
		authorRepoInterfaceMock.On("GetAllAuthors", mock.Anything).Return(authorArrayResponseFromRepo, nil).Once()
		entity, err := authorUseCaseInterface.GetAllAuthors()
		assert.Equal(t, nil, err)
		assert.Equal(t, authorArrayResponseFromRepo, entity)
	})

	t.Run("Fail Get All Authors", func(t *testing.T) {
		authorRepoInterfaceMock.On("GetAllAuthors", mock.Anything).Return([]authors.AuthorEntity{}, errors.New("error bro")).Once()
		entity, err := authorUseCaseInterface.GetAllAuthors()
		assert.Equal(t, "error bro", err.Error())
		assert.Equal(t, []authors.AuthorEntity{}, entity)
	})
}

func TestGetAuthorById(t *testing.T) {
	setupGetResponseFromRepo()
	t.Run("Success Get Author By Id", func(t *testing.T) {
		authorRepoInterfaceMock.On("GetAuthorById", int(authorResponseFromRepo.Id), mock.Anything).Return(authorResponseFromRepo, nil).Once()
		entity, err := authorUseCaseInterface.GetAuthorById(int(authorResponseFromRepo.Id))
		assert.Equal(t, nil, err)
		assert.Equal(t, authorResponseFromRepo, entity)
	})

	t.Run("Fail Get Author By Id", func(t *testing.T) {
		authorRepoInterfaceMock.On("GetAuthorById", int(authorResponseFromRepo.Id), mock.Anything).Return(authorResponseFromRepo, errors.New("error bro")).Once()
		entity, err := authorUseCaseInterface.GetAuthorById(int(authorResponseFromRepo.Id))
		assert.Equal(t, "error bro", err.Error())
		assert.Equal(t, authorResponseFromRepo, entity)
	})
}

func TestGetAuthorsByName(t *testing.T) {
	setupGetArrayResponseFromRepo()
	t.Run("Success Get Authors By Name", func(t *testing.T) {
		name := "naufal"
		authorRepoInterfaceMock.On("GetAuthorsByName", name, mock.Anything).Return(authorArrayResponseFromRepo, nil).Once()
		entity, err := authorUseCaseInterface.GetAuthorsByName(name)
		assert.Equal(t, nil, err)
		assert.Equal(t, authorArrayResponseFromRepo, entity)
	})

	t.Run("Fail Get Authors By Name", func(t *testing.T) {
		name := "naufal"
		authorRepoInterfaceMock.On("GetAuthorsByName", mock.Anything).Return([]authors.AuthorEntity{}, errors.New("error bro")).Once()
		entity, err := authorUseCaseInterface.GetAuthorsByName(name)
		assert.Equal(t, "error bro", err.Error())
		assert.Equal(t, []authors.AuthorEntity{}, entity)
	})
}

func TestDeleteAuthor(t *testing.T) {
	setupDeleteAuthor()
	id := 1
	t.Run("Success Delete Author", func(t *testing.T) {
		authorRepoInterfaceMock.On("DeleteAuthor", id, mock.Anything).Return(authorResponseFromRepo, nil).Once()
		entity, err := authorUseCaseInterface.DeleteAuthor(id)
		assert.Equal(t, nil, err)
		assert.Equal(t, authorResponseFromRepo, entity)
	})

	t.Run("Fail Delete Autho", func(t *testing.T) {
		authorRepoInterfaceMock.On("DeleteAuthor", id, mock.Anything).Return(authorResponseFromRepo, errors.New("error bro")).Once()
		entity, err := authorUseCaseInterface.DeleteAuthor(id)
		assert.Equal(t, "error bro", err.Error())
		assert.Equal(t, authors.AuthorEntity{}, entity)
	})
}
