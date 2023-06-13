package feedback

import (
	"errors"
	"farmacare/repository"
	"farmacare/shared"
	"farmacare/shared/dto"
	"strconv"
)

type (
	ViewService interface {
		CreateFeedbackCategory(req dto.CreateFeedbackCategoryRequest) (dto.CreateFeedbackCategoryResponse, error)
		GetFeedbackCategoryById(feedbackId string) (dto.FeedbackCategory, error)
		GetAllFeedbackCategory() (dto.FeedbackCategorySlice, error)
		CreateFeedback(ctx dto.SessionContext, req dto.CreateFeedbackRequest) (dto.CreateFeedbackResponse, error)
		GetAllFeedback() ([]dto.GetAllFeedbackResponse, error)
	}

	viewService struct {
		repository repository.Holder
		shared      shared.Holder
	}
)

func (v *viewService) CreateFeedbackCategory(req dto.CreateFeedbackCategoryRequest) (dto.CreateFeedbackCategoryResponse, error) {
	var (
		res dto.CreateFeedbackCategoryResponse
	)

	isFeedbackCategoryExist, _ := v.repository.FeedbackService.CheckFeedbackCategoryExist(req.Name)
	if isFeedbackCategoryExist {
		return res, errors.New("Feedback category already exist")
	}

	err := v.repository.FeedbackService.CreateFeedbackCategory(req.TransformToFeedbackCategoryModel())
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

	feedbackCategory, err = v.repository.FeedbackService.GetFeedbackCategoryById(uint(cid))
	if feedbackCategory.ID == 0 {
		return feedbackCategory, err
	}

	return feedbackCategory, nil
}

func (v *viewService) GetAllFeedbackCategory() (dto.FeedbackCategorySlice, error) {
	var (
		res dto.FeedbackCategorySlice
	)

	feedbackCategories := v.repository.FeedbackService.GetAllFeedbackCategory("")
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

	err := v.repository.FeedbackService.CreateFeedback(feedback)
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

func (v *viewService) GetAllFeedback() ([]dto.GetAllFeedbackResponse, error) {
	var (
		res dto.FeedbackSlice
	)

	feedbacks := v.repository.FeedbackService.GetAllFeedback("")
	for _, feedback := range feedbacks {
		res = append(res, feedback)
	}
	v.shared.Logger.Infof("Found %d feedbacks", len(feedbacks))

	feedbackLists := make([]dto.GetAllFeedbackResponse, 0)
	for _, feedback := range feedbacks {
		feedbackCategory, err := v.repository.FeedbackService.GetFeedbackCategoryById(feedback.FeedbackCategoryID)
		if err != nil {
			return nil, err
		}
		feedbackUser := v.repository.AuthService.GetUserContext(feedback.UserID)
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

func NewViewService(repository repository.Holder, shared shared.Holder) ViewService {
	return &viewService{
		repository: repository,
		shared:      shared,
	}
}