package community

import (
	"farmacare/shared"
	"farmacare/shared/models"
)

type (
	ForumRepository interface {
		CreateForum(forum models.Forum) error 
		GetAllForum(preload string) []models.Forum
	}

	forumRepository struct {
		shared shared.Holder
	}
)

func (s *forumRepository) CreateForum(forum models.Forum) error {
	err := s.shared.DB.Create(&forum).Error
	return err
}

func (s *forumRepository) GetAllForum(preload string) []models.Forum {
	var forums []models.Forum
	s.shared.DB.Preload(preload).Find(&forums)
	return forums
}

func NewForumRepository(holder shared.Holder) (ForumRepository, error) {
	return &forumRepository{
		shared: holder,
	}, nil
}