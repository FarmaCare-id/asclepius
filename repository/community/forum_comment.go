package community

import (
	"farmacare/shared"
	"farmacare/shared/models"
)

type (
	ForumCommentRepository interface {
		CreateForumComment(forumComment models.ForumComment) error
		GetAllForumCommentByForumID(forumID uint) []models.ForumComment 
	}

	forumCommentRepository struct {
		shared shared.Holder
	}
)

func (s *forumCommentRepository) CreateForumComment(forumComment models.ForumComment) error {
	err := s.shared.DB.Create(&forumComment).Error
	return err
}

func (s *forumCommentRepository) GetAllForumCommentByForumID(forumID uint) []models.ForumComment {
	var forumComments []models.ForumComment
	s.shared.DB.Where("forum_id = ?", forumID).Find(&forumComments)
	return forumComments
}

func NewForumCommentRepository(holder shared.Holder) (ForumCommentRepository, error) {
	return &forumCommentRepository{
		shared: holder,
	}, nil
}