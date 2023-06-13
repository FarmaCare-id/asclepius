package dto

import (
	"farmacare/shared/models"
	"time"
)

type (
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

	// GetAllFeedbackResponse GetAllFeedbackResponseID: 
	GetAllFeedbackResponse struct {
		ID uint `json:"id"`
		Issue string `json:"issue"`
		Detail string `json:"detail"`
		Category string `json:"category"`
		UserFullname string `json:"user_fullname"`
		UserRole models.UserRole `json:"user_role"`
		CreatedAt time.Time `json:"created_at"`
	}
)

func (r *CreateFeedbackCategoryRequest) TransformToFeedbackCategoryModel() models.FeedbackCategory {
	return models.FeedbackCategory{
		Name: r.Name,
		Description: r.Description,
	}
}

func (r *CreateFeedbackRequest) TransformToFeedbackModel(uid uint) models.Feedback {
	return models.Feedback{
		Issue: r.Issue,
		Detail: r.Detail,
		FeedbackCategoryID: r.FeedbackCategoryID,
		UserID: uid,
	}
}
