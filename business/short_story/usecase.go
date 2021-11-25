package shortstory

import "errors"

type ShortStoryUseCase struct {
	repo ShortStoryRepoInterface
}

func NewUseCase(shortStoryRepo ShortStoryRepoInterface) ShortStoryUseCaseInterface {
	return &ShortStoryUseCase{
		repo: shortStoryRepo,
	}
}

func (usecase *ShortStoryUseCase) GetShortStories(offset int, limit int) ([]ShortStoryEntity, error) {
	shortStories, err := usecase.repo.GetShortStories(offset, limit)
	if err != nil {
		return shortStories, err
	}
	return shortStories, nil
}

func (usecase *ShortStoryUseCase) GetShortStoryById(id int) (ShortStoryEntity, error) {
	shortStories, err := usecase.repo.GetShortStoryById(id)
	if err != nil {
		return shortStories, err
	}
	return shortStories, nil
}

func (usecase *ShortStoryUseCase) AddShortStory(shortStory ShortStoryEntity) (ShortStoryEntity, error) {
	if shortStory.Title == "" {
		return ShortStoryEntity{}, errors.New("title empty")
	}
	if shortStory.Content == "" {
		return ShortStoryEntity{}, errors.New("content empty")
	}
	shortStory, err := usecase.repo.AddShortStory(shortStory)
	if err != nil {
		return ShortStoryEntity{}, err
	}
	return shortStory, nil
}

func (usecase *ShortStoryUseCase) EditShortStory(id int, shortStory ShortStoryEntity) (ShortStoryEntity, error) {
	if shortStory.Title == "" {
		return ShortStoryEntity{}, errors.New("title empty")
	}
	if shortStory.Content == "" {
		return ShortStoryEntity{}, errors.New("content empty")
	}
	shortStory, err := usecase.repo.EditShortStory(id, shortStory)
	if err != nil {
		return ShortStoryEntity{}, err
	}
	return shortStory, nil
}

func (usecase *ShortStoryUseCase) DeleteShortStory(id int) (ShortStoryEntity, error) {
	story, err := usecase.repo.DeleteShortStory(id)
	if err != nil {
		return ShortStoryEntity{}, err
	}
	return story, nil
}
