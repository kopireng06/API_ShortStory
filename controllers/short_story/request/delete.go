package request

import shortstory "api_short_story/business/short_story"

type DeleteShortStory struct {
	AuthorID int
}

func (deleteStory DeleteShortStory) ToShortStoryEntity() shortstory.ShortStoryEntity {
	return shortstory.ShortStoryEntity{
		AuthorID: uint(deleteStory.AuthorID),
	}
}
