package shortstory_test

import (
	shortstory "api_short_story/business/short_story"
	"api_short_story/business/short_story/mocks"
	"errors"
	"testing"

	"github.com/magiconair/properties/assert"
	"github.com/stretchr/testify/mock"
)

var shortStoryRepoInterfaceMock mocks.ShortStoryRepoInterface
var shortStoryUseCaseInterface shortstory.ShortStoryUseCaseInterface
var shortStoryArrayResponseFromRepo []shortstory.ShortStoryEntity
var shortStoryResponseFromRepo shortstory.ShortStoryEntity

func setupGetResponseShortStories() {
	shortStoryUseCaseInterface = shortstory.NewUseCase(&shortStoryRepoInterfaceMock)
	shortStoryArrayResponseFromRepo = []shortstory.ShortStoryEntity{
		{
			Id:         1,
			Title:      "Cinta dalam diam",
			Content:    "Lorem ipsum dolor sit amet",
			AuthorID:   1,
			CategoryID: 1,
		},
	}
}

func setupGetArrayResponseShortStory() {
	shortStoryUseCaseInterface = shortstory.NewUseCase(&shortStoryRepoInterfaceMock)
	shortStoryResponseFromRepo = shortstory.ShortStoryEntity{
		Id:         1,
		Title:      "Cinta dalam diam",
		Content:    "Lorem ipsum dolor sit amet",
		AuthorID:   1,
		CategoryID: 1,
	}
}

func TestGetShortStories(t *testing.T) {
	setupGetResponseShortStories()
	offset := 0
	limit := 5
	t.Run("Success Get Short Stories", func(t *testing.T) {
		shortStoryRepoInterfaceMock.On("GetShortStories", offset, limit, mock.Anything).Return(shortStoryArrayResponseFromRepo, nil).Once()
		entity, err := shortStoryUseCaseInterface.GetShortStories(offset, limit)
		assert.Equal(t, nil, err)
		assert.Equal(t, shortStoryArrayResponseFromRepo, entity)
	})

	t.Run("Fail Get Short Stories", func(t *testing.T) {
		shortStoryRepoInterfaceMock.On("GetShortStories", offset, limit, mock.Anything).Return([]shortstory.ShortStoryEntity{}, errors.New("error bro")).Once()
		entity, err := shortStoryUseCaseInterface.GetShortStories(offset, limit)
		assert.Equal(t, errors.New("error bro"), err)
		assert.Equal(t, []shortstory.ShortStoryEntity{}, entity)
	})
}

func TestGetShortStoryById(t *testing.T) {
	setupGetArrayResponseShortStory()
	id := 1
	t.Run("Success Get Short Stories By Id", func(t *testing.T) {
		shortStoryRepoInterfaceMock.On("GetShortStoryById", id, mock.Anything).Return(shortStoryResponseFromRepo, nil).Once()
		entity, err := shortStoryUseCaseInterface.GetShortStoryById(id)
		assert.Equal(t, nil, err)
		assert.Equal(t, shortStoryResponseFromRepo, entity)
	})

	t.Run("Fail Get Short Stories", func(t *testing.T) {
		shortStoryRepoInterfaceMock.On("GetShortStoryById", id, mock.Anything).Return(shortstory.ShortStoryEntity{}, errors.New("error bro")).Once()
		entity, err := shortStoryUseCaseInterface.GetShortStoryById(id)
		assert.Equal(t, errors.New("error bro"), err)
		assert.Equal(t, shortstory.ShortStoryEntity{}, entity)
	})
}

func TestGetShortStoriesByTitle(t *testing.T) {
	setupGetResponseShortStories()
	title := "halo"
	t.Run("Success Get Short Stories By Title", func(t *testing.T) {
		shortStoryRepoInterfaceMock.On("GetShortStoriesByTitle", title, mock.Anything).Return(shortStoryArrayResponseFromRepo, nil).Once()
		entity, err := shortStoryUseCaseInterface.GetShortStoriesByTitle(title)
		assert.Equal(t, nil, err)
		assert.Equal(t, shortStoryArrayResponseFromRepo, entity)
	})

	t.Run("Fail Get Short Stories By Title", func(t *testing.T) {
		shortStoryRepoInterfaceMock.On("GetShortStoriesByTitle", title, mock.Anything).Return([]shortstory.ShortStoryEntity{}, errors.New("error bro")).Once()
		entity, err := shortStoryUseCaseInterface.GetShortStoriesByTitle(title)
		assert.Equal(t, errors.New("error bro"), err)
		assert.Equal(t, []shortstory.ShortStoryEntity{}, entity)
	})
}

func TestGetShortStoriesByIdCategory(t *testing.T) {
	setupGetResponseShortStories()
	idCategory := 1
	t.Run("Success Get Short Stories By Id Category", func(t *testing.T) {
		shortStoryRepoInterfaceMock.On("GetShortStoriesByIdCategory", idCategory, mock.Anything).Return(shortStoryArrayResponseFromRepo, nil).Once()
		entity, err := shortStoryUseCaseInterface.GetShortStoriesByIdCategory(idCategory)
		assert.Equal(t, nil, err)
		assert.Equal(t, shortStoryArrayResponseFromRepo, entity)
	})

	t.Run("Fail Get Short Stories", func(t *testing.T) {
		shortStoryRepoInterfaceMock.On("GetShortStoriesByIdCategory", idCategory, mock.Anything).Return([]shortstory.ShortStoryEntity{}, errors.New("error bro")).Once()
		entity, err := shortStoryUseCaseInterface.GetShortStoriesByIdCategory(idCategory)
		assert.Equal(t, errors.New("error bro"), err)
		assert.Equal(t, []shortstory.ShortStoryEntity{}, entity)
	})
}

func TestGetShortStoriesByIdAuthor(t *testing.T) {
	setupGetResponseShortStories()
	idAuthor := 1
	t.Run("Success Get Short Stories By Id Author", func(t *testing.T) {
		shortStoryRepoInterfaceMock.On("GetShortStoriesByIdAuthor", idAuthor, mock.Anything).Return(shortStoryArrayResponseFromRepo, nil).Once()
		entity, err := shortStoryUseCaseInterface.GetShortStoriesByIdAuthor(idAuthor)
		assert.Equal(t, nil, err)
		assert.Equal(t, shortStoryArrayResponseFromRepo, entity)
	})

	t.Run("Fail Get Short Stories", func(t *testing.T) {
		shortStoryRepoInterfaceMock.On("GetShortStoriesByIdAuthor", idAuthor, mock.Anything).Return([]shortstory.ShortStoryEntity{}, errors.New("error bro")).Once()
		entity, err := shortStoryUseCaseInterface.GetShortStoriesByIdAuthor(idAuthor)
		assert.Equal(t, errors.New("error bro"), err)
		assert.Equal(t, []shortstory.ShortStoryEntity{}, entity)
	})
}

func TestAddShortStory(t *testing.T) {
	setupGetResponseShortStories()
	t.Run("Success Add Short Story", func(t *testing.T) {
		requestAdd := shortstory.ShortStoryEntity{
			Title:      "ini judul",
			Content:    "ini content",
			AuthorID:   1,
			CategoryID: 1,
		}
		shortStoryRepoInterfaceMock.On("AddShortStory", mock.AnythingOfType("shortstory.ShortStoryEntity"), mock.Anything).Return(shortStoryResponseFromRepo, nil).Once()
		entity, err := shortStoryUseCaseInterface.AddShortStory(requestAdd)
		assert.Equal(t, nil, err)
		assert.Equal(t, shortStoryResponseFromRepo, entity)
	})

	t.Run("Fail Add Short Story", func(t *testing.T) {
		requestAdd := shortstory.ShortStoryEntity{
			Title:      "ini judul",
			Content:    "ini content",
			AuthorID:   1,
			CategoryID: 1,
		}
		shortStoryRepoInterfaceMock.On("AddShortStory", mock.AnythingOfType("shortstory.ShortStoryEntity"), mock.Anything).Return(shortstory.ShortStoryEntity{}, errors.New("error bro")).Once()
		entity, err := shortStoryUseCaseInterface.AddShortStory(requestAdd)
		assert.Equal(t, "error bro", err.Error())
		assert.Equal(t, shortstory.ShortStoryEntity{}, entity)
	})

	t.Run("Add Short Story with Title Empty", func(t *testing.T) {
		requestAdd := shortstory.ShortStoryEntity{
			Title:      "",
			Content:    "ini content",
			AuthorID:   1,
			CategoryID: 1,
		}
		entity, err := shortStoryUseCaseInterface.AddShortStory(requestAdd)
		assert.Equal(t, "title empty", err.Error())
		assert.Equal(t, shortstory.ShortStoryEntity{}, entity)
	})

	t.Run("Add Short Story with Content Empty", func(t *testing.T) {
		requestAdd := shortstory.ShortStoryEntity{
			Title:      "ini title",
			Content:    "",
			AuthorID:   1,
			CategoryID: 1,
		}
		entity, err := shortStoryUseCaseInterface.AddShortStory(requestAdd)
		assert.Equal(t, "content empty", err.Error())
		assert.Equal(t, shortstory.ShortStoryEntity{}, entity)
	})

	t.Run("Add Short Story with Category Empty", func(t *testing.T) {
		requestAdd := shortstory.ShortStoryEntity{
			Title:      "ini title",
			Content:    "ini content",
			AuthorID:   1,
			CategoryID: 0,
		}
		entity, err := shortStoryUseCaseInterface.AddShortStory(requestAdd)
		assert.Equal(t, "category empty", err.Error())
		assert.Equal(t, shortstory.ShortStoryEntity{}, entity)
	})
}

func TestEditShortStory(t *testing.T) {
	setupGetResponseShortStories()
	id := 1
	t.Run("Success Edit Short Story", func(t *testing.T) {
		requestEdit := shortstory.ShortStoryEntity{
			Title:      "ini judul",
			Content:    "ini content",
			AuthorID:   1,
			CategoryID: 1,
		}
		shortStoryRepoInterfaceMock.On("EditShortStory", id, mock.AnythingOfType("shortstory.ShortStoryEntity"), mock.Anything).Return(shortStoryResponseFromRepo, nil).Once()
		entity, err := shortStoryUseCaseInterface.EditShortStory(id, requestEdit)
		assert.Equal(t, nil, err)
		assert.Equal(t, shortStoryResponseFromRepo, entity)
	})

	t.Run("Fail Edit Short Story", func(t *testing.T) {
		requestEdit := shortstory.ShortStoryEntity{
			Title:      "ini judul",
			Content:    "ini content",
			AuthorID:   1,
			CategoryID: 1,
		}
		shortStoryRepoInterfaceMock.On("EditShortStory", id, mock.AnythingOfType("shortstory.ShortStoryEntity"), mock.Anything).Return(shortstory.ShortStoryEntity{}, errors.New("error bro")).Once()
		entity, err := shortStoryUseCaseInterface.EditShortStory(id, requestEdit)
		assert.Equal(t, "error bro", err.Error())
		assert.Equal(t, shortstory.ShortStoryEntity{}, entity)
	})

	t.Run("Edit Short Story with Title Empty", func(t *testing.T) {
		requestEdit := shortstory.ShortStoryEntity{
			Title:      "",
			Content:    "ini content",
			AuthorID:   1,
			CategoryID: 1,
		}
		entity, err := shortStoryUseCaseInterface.EditShortStory(id, requestEdit)
		assert.Equal(t, "title empty", err.Error())
		assert.Equal(t, shortstory.ShortStoryEntity{}, entity)
	})

	t.Run("Edit Short Story with Content Empty", func(t *testing.T) {
		requestEdit := shortstory.ShortStoryEntity{
			Title:      "ini title",
			Content:    "",
			AuthorID:   1,
			CategoryID: 1,
		}
		entity, err := shortStoryUseCaseInterface.EditShortStory(id, requestEdit)
		assert.Equal(t, "content empty", err.Error())
		assert.Equal(t, shortstory.ShortStoryEntity{}, entity)
	})

	t.Run("Edit Short Story with Category Empty", func(t *testing.T) {
		requestEdit := shortstory.ShortStoryEntity{
			Title:      "ini title",
			Content:    "ini content",
			AuthorID:   1,
			CategoryID: 0,
		}
		entity, err := shortStoryUseCaseInterface.EditShortStory(id, requestEdit)
		assert.Equal(t, "category empty", err.Error())
		assert.Equal(t, shortstory.ShortStoryEntity{}, entity)
	})
}

func TestDeleteShortStory(t *testing.T) {
	id := 1
	requestDelete := shortstory.ShortStoryEntity{
		AuthorID: 1,
	}
	shortStoryResponseFromRepo = shortstory.ShortStoryEntity{
		Id: uint(id),
	}
	t.Run("Success Delete Short Story", func(t *testing.T) {
		shortStoryRepoInterfaceMock.On("DeleteShortStory", id, mock.Anything).Return(shortStoryResponseFromRepo, nil).Once()
		entity, err := shortStoryUseCaseInterface.DeleteShortStory(id, requestDelete)
		assert.Equal(t, nil, err)
		assert.Equal(t, shortStoryResponseFromRepo, entity)
	})

	t.Run("Fail Delete Short Story", func(t *testing.T) {
		shortStoryRepoInterfaceMock.On("DeleteShortStory", id, mock.Anything).Return(shortstory.ShortStoryEntity{}, errors.New("error bro")).Once()
		entity, err := shortStoryUseCaseInterface.DeleteShortStory(id, requestDelete)
		assert.Equal(t, "error bro", err.Error())
		assert.Equal(t, shortstory.ShortStoryEntity{}, entity)
	})
}
