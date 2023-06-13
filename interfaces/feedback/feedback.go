package feedback

import (
	"errors"
	"farmacare/application"
	"farmacare/shared"
	"farmacare/shared/dto"
	"strconv"
)

type (
	ViewService interface {
		CreateFeedbackCategory(req dto.CreateFeedbackCategoryRequest) (dto.CreateFeedbackCategoryResponse, error)
		GetFeedbackCategoryById(feedbackId string) (dto.FeedbackCategory, error)
		ListFeedbackCategory() (dto.FeedbackCategorySlice, error)
		CreateFeedback(ctx dto.SessionContext, req dto.CreateFeedbackRequest) (dto.CreateFeedbackResponse, error)
		ListFeedback() ([]dto.GetAllFeedbackResponse, error)
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

func (v *viewService) GetFeedbackCategoryById(feedbackId string) (dto.FeedbackCategory, error) {
	var (
		feedbackCategory dto.FeedbackCategory
	)

	cid, err := strconv.Atoi(feedbackId)
	if err != nil {
		return feedbackCategory, err
	}

	feedbackCategory, err = v.application.FeedbackService.GetFeedbackCategoryById(uint(cid))
	if feedbackCategory.ID == 0 {
		return feedbackCategory, err
	}

	return feedbackCategory, nil
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

func (v *viewService) ListFeedback() ([]dto.GetAllFeedbackResponse, error) {
	var (
		res dto.FeedbackSlice
	)

	feedbacks := v.application.FeedbackService.ListFeedback("")
	for _, feedback := range feedbacks {
		res = append(res, feedback)
	}
	v.shared.Logger.Infof("Found %d feedbacks", len(feedbacks))

	feedbackLists := make([]dto.GetAllFeedbackResponse, 0)
	for _, feedback := range feedbacks {
		feedbackCategory, err := v.application.FeedbackService.GetFeedbackCategoryById(feedback.FeedbackCategoryID)
		if err != nil {
			return nil, err
		}
		feedbackUser := v.application.AuthService.GetUserContext(feedback.UserID)
		feedbackList := dto.GetAllFeedbackResponse{
			ID: feedback.ID,
			Issue: feedback.Issue,
			Detail: feedback.Detail,
			Category: feedbackCategory.Name,
			UserFullname: feedbackUser.Fullname,
			UserRole: feedbackUser.Role,
			CreatedAt: feedback.CreatedAt,
		}
		feedbackLists = append(feedbackLists, feedbackList)
	}

	return feedbackLists, nil
}

func NewViewService(application application.Holder, shared shared.Holder) ViewService {
	return &viewService{
		application: application,
		shared:      shared,
	}
}