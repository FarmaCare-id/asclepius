package community

import (
	"farmacare/shared"
	"farmacare/shared/models"
)

type (
	ForumRepository interface {
		CreateForum(forum models.Forum) error 
	}

	forumRepository struct {
		shared shared.Holder
	}
)

func (s *forumRepository) CreateForum(forum models.Forum) error {
	err := s.shared.DB.Create(&forum).Error
	return err
}

func NewForumRepository(holder shared.Holder) (ForumRepository, error) {
	return &forumRepository{
		shared: holder,
	}, nil
}