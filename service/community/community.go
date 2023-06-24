package community

import (
	"farmacare/repository"
	"farmacare/shared"
	"farmacare/shared/dto"
	"farmacare/shared/models"
	"strconv"
)

type (
	ViewService interface {
		CreateForum(ctx dto.SessionContext, req dto.CreateForumRequest) (dto.CreateForumResponse, error)
		GetAllForum() ([]dto.GetForumResponse, error)
		CreateForumComment(ctx dto.SessionContext, req dto.CreateForumCommentRequest, forumID string) (dto.CreateForumCommentResponse, error)
		GetAllForumCommentByForumID(forumID string) ([]dto.GetForumCommentResponse, error)
	}

	viewService struct {
		repository  repository.Holder
		shared      shared.Holder
	}
)

func (v *viewService) CreateForum(ctx dto.SessionContext, req dto.CreateForumRequest) (dto.CreateForumResponse, error) {
	var (
		res dto.CreateForumResponse
	)

	forum := models.Forum {
		Title: req.Title,
		Description: req.Description,
		Vote: 0,
		UserID: ctx.User.ID,
	}

	err := v.repository.ForumRepository.CreateForum(forum)
	if err != nil {
		return res, err
	}

	res = dto.CreateForumResponse{
		ID: forum.ID,
		Title: forum.Title,
		Description: forum.Description,
		Vote: forum.Vote,
		UserID: forum.UserID,
	}

	return res, nil
}

func (v *viewService) GetAllForum() ([]dto.GetForumResponse, error) {
	var (
		res []dto.GetForumResponse
	)

	forums := v.repository.ForumRepository.GetAllForum("")

	for _, forum := range forums {
		res = append(res, dto.GetForumResponse{
			ID: forum.ID,
			Title: forum.Title,
			Description: forum.Description,
			Vote: forum.Vote,
			UserID: forum.UserID,
		})
	}

	return res, nil
}

func (v *viewService) CreateForumComment(ctx dto.SessionContext, req dto.CreateForumCommentRequest, forumID string) (dto.CreateForumCommentResponse, error) {
	var (
		res dto.CreateForumCommentResponse
	)

	// convert with atoi to uint
	fid, err := strconv.Atoi(forumID)
	if err != nil {
		return res, err
	}

	forumComment := models.ForumComment {
		ForumID: uint(fid),
		Comment: req.Comment,
		UserID: ctx.User.ID,
	}

	err = v.repository.ForumCommentRepository.CreateForumComment(forumComment)
	if err != nil {
		return res, err
	}

	res = dto.CreateForumCommentResponse{
		ID: forumComment.ID,
		ForumID: forumComment.ForumID,
		Comment: forumComment.Comment,
		UserID: forumComment.UserID,
	}

	return res, nil
}

func (v *viewService) GetAllForumCommentByForumID(forumID string) ([]dto.GetForumCommentResponse, error) {
	var (
		res []dto.GetForumCommentResponse
	)

	// convert with atoi to uint
	cid, err := strconv.Atoi(forumID)
	if err != nil {
		return res, err
	}

	forumComments := v.repository.ForumCommentRepository.GetAllForumCommentByForumID(uint(cid))

	for _, forumComment := range forumComments {
		res = append(res, dto.GetForumCommentResponse{
			ID: forumComment.ID,
			ForumID: forumComment.ForumID,
			Comment: forumComment.Comment,
			UserID: forumComment.UserID,
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