package authors

import "errors"

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
	author, err := usecase.repo.Login(author)
	if err != nil {
		return AuthorEntity{}, err
	}
	return author, nil
}

func (usecase *AuthorUseCase) GetAllAuthors() ([]AuthorEntity, error) {
	return []AuthorEntity{}, nil
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
	if author.Profile == "" {
		return AuthorEntity{}, errors.New("profile empty")
	}
	author, err := usecase.repo.AddAuthor(author)
	if err != nil {
		return AuthorEntity{}, err
	}
	return author, nil
}
