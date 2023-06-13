package feedback

import (
	"farmacare/shared"
	"farmacare/shared/dto"
)

type (
	Service interface {
		CheckFeedbackCategoryExist(name string) (bool, dto.FeedbackCategory)
		CreateFeedbackCategory(feedbackCategory dto.FeedbackCategory) error
		GetFeedbackCategoryById(feedbackId uint) (dto.FeedbackCategory, error)
		GetAllFeedbackCategory(preload string) dto.FeedbackCategorySlice
		CreateFeedback(feedback dto.Feedback) error
		GetAllFeedback(preload string) dto.FeedbackSlice
	}

	service struct {
		shared shared.Holder
	}
)

func (s *service) CheckFeedbackCategoryExist(name string) (bool, dto.FeedbackCategory) {
	var feedbackCategory dto.FeedbackCategory
	err := s.shared.DB.First(&feedbackCategory, "name = ?", name).Error
	return err == nil, feedbackCategory
}

func (s *service) CreateFeedbackCategory(feedbackCategory dto.FeedbackCategory) error {
	err := s.shared.DB.Create(&feedbackCategory).Error
	return err
}

func (s *service) GetFeedbackCategoryById(feedbackId uint) (dto.FeedbackCategory, error) {
	var feedbackCategory dto.FeedbackCategory
	err := s.shared.DB.First(&feedbackCategory, feedbackId).Error
	return feedbackCategory, err
}

func (s *service) GetAllFeedbackCategory(preload string) dto.FeedbackCategorySlice {
	var feedbackCategories []dto.FeedbackCategory
	s.shared.DB.Preload(preload).Find(&feedbackCategories)
	return feedbackCategories
}

func (s *service) CreateFeedback(feedback dto.Feedback) error {
	err := s.shared.DB.Create(&feedback).Error
	return err
}

func (s *service) GetAllFeedback(preload string) dto.FeedbackSlice {
	var feedbacks []dto.Feedback
	s.shared.DB.Preload(preload).Find(&feedbacks)
	return feedbacks
}

func NewFeedbackService(holder shared.Holder) (Service, error) {
	return &service{
		shared: holder,
	}, nil
}