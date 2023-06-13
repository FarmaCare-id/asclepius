package feedback

import (
	"farmacare/shared"
	"farmacare/shared/models"
)

type (
	Repository interface {
		CheckFeedbackCategoryExist(name string) (bool, models.FeedbackCategory)
		CreateFeedbackCategory(feedbackCategory models.FeedbackCategory) error
		GetFeedbackCategoryById(feedbackId uint) (models.FeedbackCategory, error)
		GetAllFeedbackCategory(preload string) []models.FeedbackCategory
		CreateFeedback(feedback models.Feedback) error
		GetAllFeedback(preload string) []models.Feedback
	}

	repository struct {
		shared shared.Holder
	}
)

func (s *repository) CheckFeedbackCategoryExist(name string) (bool, models.FeedbackCategory) {
	var feedbackCategory models.FeedbackCategory
	err := s.shared.DB.First(&feedbackCategory, "name = ?", name).Error
	return err == nil, feedbackCategory
}

func (s *repository) CreateFeedbackCategory(feedbackCategory models.FeedbackCategory) error {
	err := s.shared.DB.Create(&feedbackCategory).Error
	return err
}

func (s *repository) GetFeedbackCategoryById(feedbackId uint) (models.FeedbackCategory, error) {
	var feedbackCategory models.FeedbackCategory
	err := s.shared.DB.First(&feedbackCategory, feedbackId).Error
	return feedbackCategory, err
}

func (s *repository) GetAllFeedbackCategory(preload string) []models.FeedbackCategory {
	var feedbackCategories []models.FeedbackCategory
	s.shared.DB.Preload(preload).Find(&feedbackCategories)
	return feedbackCategories
}

func (s *repository) CreateFeedback(feedback models.Feedback) error {
	err := s.shared.DB.Create(&feedback).Error
	return err
}

func (s *repository) GetAllFeedback(preload string) []models.Feedback {
	var feedbacks []models.Feedback
	s.shared.DB.Preload(preload).Find(&feedbacks)
	return feedbacks
}

func FeedbackRepository(holder shared.Holder) (Repository, error) {
	return &repository{
		shared: holder,
	}, nil
}