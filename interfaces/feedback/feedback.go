package feedback

import (
	"errors"
	"farmacare/application"
	"farmacare/shared"
	"farmacare/shared/dto"
)

type (
	ViewService interface {
		CreateFeedbackCategory(req dto.CreateFeedbackCategoryRequest) (dto.CreateFeedbackCategoryResponse, error)
		ListFeedbackCategory() (dto.FeedbackCategorySlice, error)
		CreateFeedback(ctx dto.SessionContext, req dto.CreateFeedbackRequest) (dto.CreateFeedbackResponse, error)
	}

	viewService struct {
		application application.Holder
		shared      shared.Holder
	}
)

func (v *viewService) CreateFeedbackCategory(req dto.CreateFeedbackCategoryRequest) (dto.CreateFeedbackCategoryResponse, error) {
	var (
		res dto.CreateFeedbackCategoryResponse
	)

	isFeedbackCategoryExist, _ := v.application.FeedbackService.CheckFeedbackCategoryExist(req.Name)
	if isFeedbackCategoryExist {
		return res, errors.New("Feedback category already exist")
	}

	err := v.application.FeedbackService.CreateFeedbackCategory(req.TransformToFeedbackCategoryModel())
	if err != nil {
		return res, err
	}

	res = dto.CreateFeedbackCategoryResponse{
		Name: req.Name,
		Description: req.Description,
	}

	return res, nil
}

func (v *viewService) ListFeedbackCategory() (dto.FeedbackCategorySlice, error) {
	var (
		res dto.FeedbackCategorySlice
	)

	feedbackCategories := v.application.FeedbackService.ListFeedbackCategory("")
	for _, feedbackCategory := range feedbackCategories {
		res = append(res, feedbackCategory)
	}

	return res, nil
}

func (v *viewService) CreateFeedback(ctx dto.SessionContext, req dto.CreateFeedbackRequest) (dto.CreateFeedbackResponse, error) {
	var (
		res dto.CreateFeedbackResponse
	)

	feedback := dto.Feedback{
		UserID:    ctx.User.ID,
		Issue:     req.Issue,
		Detail:    req.Detail,
		FeedbackCategoryID: req.FeedbackCategoryID,
	}

	err := v.application.FeedbackService.CreateFeedback(feedback)
	if err != nil {
		return res, err
	}

	res = dto.CreateFeedbackResponse{
		Issue: req.Issue,
		Detail: req.Detail,
		FeedbackCategoryID: req.FeedbackCategoryID,
	}

	return res, nil
}

func NewViewService(application application.Holder, shared shared.Holder) ViewService {
	return &viewService{
		application: application,
		shared:      shared,
	}
}