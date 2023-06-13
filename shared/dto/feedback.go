package dto

import (
	"time"

	"gorm.io/gorm"
)

type (
	FeedbackCategory struct {
		gorm.Model
		ID uint `gorm:"primaryKey;autoIncrement"`
		Name string `gorm:"column:name"`
		Description string `gorm:"column:description"`
		Feedbacks []Feedback `gorm:"foreignKey:FeedbackCategoryID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	}

	FeedbackCategorySlice []FeedbackCategory

	// CreateFeedbackCategoryRequest CreateFeedbackCategoryRequest
	CreateFeedbackCategoryRequest struct {
		Name string `json:"name" validate:"required"`
		Description string `json:"description" validate:"required"`
	}

	// CreateFeedbackCategoryResponse CreateFeedbackCategoryResponse
	CreateFeedbackCategoryResponse struct {
		Name string `json:"name"`
		Description string `json:"description"`
	}

	Feedback struct {
		gorm.Model
		ID uint `gorm:"primaryKey;autoIncrement"`
		Issue string `gorm:"column:issue"`
		Detail string `gorm:"column:detail"`
		FeedbackCategoryID uint
		FeedbackCategory FeedbackCategory `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
		UserID uint
		User User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
		CreatedAt time.Time
		UpdatedAt time.Time
		DeletedAt gorm.DeletedAt
	}

	FeedbackSlices []Feedback

	// CreateFeedbackRequest CreateFeedbackRequest
	CreateFeedbackRequest struct {
		Issue string `json:"issue" validate:"required"`
		Detail string `json:"detail" validate:"required"`
		FeedbackCategoryID uint `json:"feedback_category_id" validate:"required"`
	}

	// CreateFeedbackResponse CreateFeedbackResponse
	CreateFeedbackResponse struct {
		Issue string `json:"issue"`
		Detail string `json:"detail"`
		FeedbackCategoryID uint `json:"feedback_category_id" validate:"required"`
	}
)

func (r *CreateFeedbackCategoryRequest) TransformToFeedbackCategoryModel() FeedbackCategory {
	return FeedbackCategory{
		Name: r.Name,
		Description: r.Description,
	}
}

func (r *CreateFeedbackRequest) TransformToFeedbackModel(uid uint) Feedback {
	return Feedback{
		Issue: r.Issue,
		Detail: r.Detail,
		FeedbackCategoryID: r.FeedbackCategoryID,
		UserID: uid,
	}
}
