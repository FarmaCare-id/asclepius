package feedback

import (
	"farmacare/repository"
	"farmacare/shared"
	"farmacare/shared/dto"
	"farmacare/shared/exception"
	"farmacare/shared/models"
	"strconv"
)

type (
	ViewService interface {
		CreateFeedbackCategory(req dto.CreateFeedbackCategoryRequest) (dto.CreateFeedbackCategoryResponse, error)
		GetFeedbackCategoryById(feedbackId string) (models.FeedbackCategory, error)
		GetAllFeedbackCategory() ([]dto.GetAllFeedbackCategoryResponse, error)
		CreateFeedback(ctx dto.SessionContext, req dto.CreateFeedbackRequest) (dto.CreateFeedbackResponse, error)
		GetAllFeedback() ([]dto.GetAllFeedbackResponse, error)
		GetAllFeedbackByUserId(userId string) ([]dto.GetAllFeedbackByUserIdResponse, error)
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

	isFeedbackCategoryExist, _ := v.repository.FeedbackRepository.CheckFeedbackCategoryExist(req.Name)
	if isFeedbackCategoryExist {
		return res, exception.FeedbackCategoryAlreadyExist()
	}

	err := v.repository.FeedbackRepository.CreateFeedbackCategory(req.TransformToFeedbackCategoryModel())
	if err != nil {
		return res, err
	}

	res = dto.CreateFeedbackCategoryResponse{
		Name: req.Name,
		Description: req.Description,
	}

	return res, nil
}

func (v *viewService) GetFeedbackCategoryById(feedbackId string) (models.FeedbackCategory, error) {
	var (
		feedbackCategory models.FeedbackCategory
	)

	cid, err := strconv.Atoi(feedbackId)
	if err != nil {
		return feedbackCategory, err
	}

	feedbackCategory, err = v.repository.FeedbackRepository.GetFeedbackCategoryById(uint(cid))
	if feedbackCategory.ID == 0 {
		return feedbackCategory, err
	}

	return feedbackCategory, nil
}

func (v *viewService) GetAllFeedbackCategory() ([]dto.GetAllFeedbackCategoryResponse, error) {
	var (
		res []dto.GetAllFeedbackCategoryResponse
	)

	feedbackCategories := v.repository.FeedbackRepository.GetAllFeedbackCategory("")
	for _, feedbackCategory := range feedbackCategories {
		res = append(res, dto.GetAllFeedbackCategoryResponse{
			ID: feedbackCategory.ID,
			Name: feedbackCategory.Name,
			Description: feedbackCategory.Description,
		})
	}

	return res, nil
}

func (v *viewService) CreateFeedback(ctx dto.SessionContext, req dto.CreateFeedbackRequest) (dto.CreateFeedbackResponse, error) {
	var (
		res dto.CreateFeedbackResponse
	)

	feedback := models.Feedback{
		UserID:    ctx.User.ID,
		Issue:     req.Issue,
		Detail:    req.Detail,
		FeedbackCategoryID: req.FeedbackCategoryID,
	}

	err := v.repository.FeedbackRepository.CreateFeedback(feedback)
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
		res []models.Feedback
	)

	feedbacks := v.repository.FeedbackRepository.GetAllFeedback("")
	for _, feedback := range feedbacks {
		res = append(res, feedback)
	}
	v.shared.Logger.Infof("Found %d feedbacks", len(feedbacks))

	feedbackLists := make([]dto.GetAllFeedbackResponse, 0)
	for _, feedback := range feedbacks {
		feedbackCategory, err := v.repository.FeedbackRepository.GetFeedbackCategoryById(feedback.FeedbackCategoryID)
		if err != nil {
			return nil, err
		}
		feedbackUser := v.repository.UserRepository.GetUserContext(feedback.UserID)
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

func (v *viewService) GetAllFeedbackByUserId(userId string) ([]dto.GetAllFeedbackByUserIdResponse, error) {
	var (
		res []dto.GetAllFeedbackByUserIdResponse
	)

	cid, err := strconv.Atoi(userId)
	if err != nil {
		return res, err
	}

	feedbacks, err := v.repository.FeedbackRepository.GetAllFeedbackByUserId(uint(cid))
	if err != nil {
		return res, err
	}

	for _, feedback := range feedbacks {
		res = append(res, dto.GetAllFeedbackByUserIdResponse{
			ID: feedback.ID,
			Issue: feedback.Issue,
			Detail: feedback.Detail,
			Category: feedback.FeedbackCategory.Name,
			CreatedAt: feedback.CreatedAt,
		})
	}

	return res, nil
}

func NewViewService(repository repository.Holder, shared shared.Holder) ViewService {
	return &viewService{
		repository: repository,
		shared:      shared,
	}
}