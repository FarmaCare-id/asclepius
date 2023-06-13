package feedback

import (
	"farmacare/shared"
	"farmacare/shared/dto"
)

type (
	Repository interface {
		CheckFeedbackCategoryExist(name string) (bool, dto.FeedbackCategory)
		CreateFeedbackCategory(feedbackCategory dto.FeedbackCategory) error
		GetFeedbackCategoryById(feedbackId uint) (dto.FeedbackCategory, error)
		GetAllFeedbackCategory(preload string) dto.FeedbackCategorySlice
		CreateFeedback(feedback dto.Feedback) error
		GetAllFeedback(preload string) dto.FeedbackSlice
	}

	repository struct {
		shared shared.Holder
	}
)

func (s *repository) CheckFeedbackCategoryExist(name string) (bool, dto.FeedbackCategory) {
	var feedbackCategory dto.FeedbackCategory
	err := s.shared.DB.First(&feedbackCategory, "name = ?", name).Error
	return err == nil, feedbackCategory
}

func (s *repository) CreateFeedbackCategory(feedbackCategory dto.FeedbackCategory) error {
	err := s.shared.DB.Create(&feedbackCategory).Error
	return err
}

func (s *repository) GetFeedbackCategoryById(feedbackId uint) (dto.FeedbackCategory, error) {
	var feedbackCategory dto.FeedbackCategory
	err := s.shared.DB.First(&feedbackCategory, feedbackId).Error
	return feedbackCategory, err
}

func (s *repository) GetAllFeedbackCategory(preload string) dto.FeedbackCategorySlice {
	var feedbackCategories []dto.FeedbackCategory
	s.shared.DB.Preload(preload).Find(&feedbackCategories)
	return feedbackCategories
}

func (s *repository) CreateFeedback(feedback dto.Feedback) error {
	err := s.shared.DB.Create(&feedback).Error
	return err
}

func (s *repository) GetAllFeedback(preload string) dto.FeedbackSlice {
	var feedbacks []dto.Feedback
	s.shared.DB.Preload(preload).Find(&feedbacks)
	return feedbacks
}

func FeedbackRepository(holder shared.Holder) (Repository, error) {
	return &repository{
		shared: holder,
	}, nil
}