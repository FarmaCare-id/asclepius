package feedback

import (
	"farmacare/shared"
	"farmacare/shared/dto"
)

type (
	Feedback interface {
		CheckFeedbackCategoryExist(name string) (bool, dto.FeedbackCategory)
		CreateFeedbackCategory(feedbackCategory dto.FeedbackCategory) error
		GetFeedbackCategoryById(feedbackId uint) (dto.FeedbackCategory, error)
		GetAllFeedbackCategory(preload string) dto.FeedbackCategorySlice
		CreateFeedback(feedback dto.Feedback) error
		GetAllFeedback(preload string) dto.FeedbackSlice
	}

	feedback struct {
		shared shared.Holder
	}
)

func (s *feedback) CheckFeedbackCategoryExist(name string) (bool, dto.FeedbackCategory) {
	var feedbackCategory dto.FeedbackCategory
	err := s.shared.DB.First(&feedbackCategory, "name = ?", name).Error
	return err == nil, feedbackCategory
}

func (s *feedback) CreateFeedbackCategory(feedbackCategory dto.FeedbackCategory) error {
	err := s.shared.DB.Create(&feedbackCategory).Error
	return err
}

func (s *feedback) GetFeedbackCategoryById(feedbackId uint) (dto.FeedbackCategory, error) {
	var feedbackCategory dto.FeedbackCategory
	err := s.shared.DB.First(&feedbackCategory, feedbackId).Error
	return feedbackCategory, err
}

func (s *feedback) GetAllFeedbackCategory(preload string) dto.FeedbackCategorySlice {
	var feedbackCategories []dto.FeedbackCategory
	s.shared.DB.Preload(preload).Find(&feedbackCategories)
	return feedbackCategories
}

func (s *feedback) CreateFeedback(feedback dto.Feedback) error {
	err := s.shared.DB.Create(&feedback).Error
	return err
}

func (s *feedback) GetAllFeedback(preload string) dto.FeedbackSlice {
	var feedbacks []dto.Feedback
	s.shared.DB.Preload(preload).Find(&feedbacks)
	return feedbacks
}

func FeedbackRepository(holder shared.Holder) (Feedback, error) {
	return &feedback{
		shared: holder,
	}, nil
}