package community

import (
	"farmacare/shared"
	"farmacare/shared/models"
)

type (
	ForumCommentRepository interface {
		CreateForumComment(forumComment models.ForumComment) error 
	}

	forumCommentRepository struct {
		shared shared.Holder
	}
)

func (s *forumCommentRepository) CreateForumComment(forumComment models.ForumComment) error {
	err := s.shared.DB.Create(&forumComment).Error
	return err
}

func NewForumCommentRepository(holder shared.Holder) (ForumCommentRepository, error) {
	return &forumCommentRepository{
		shared: holder,
	}, nil
}