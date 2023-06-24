package community

import (
	"farmacare/repository"
	"farmacare/shared"
	"farmacare/shared/dto"
	"farmacare/shared/models"
)

type (
	ViewService interface {
		CreateForum(ctx dto.SessionContext, req dto.CreateForumRequest) (dto.CreateForumResponse, error)
		GetAllForum() ([]dto.GetForumResponse, error)
		CreateForumComment(ctx dto.SessionContext, req dto.CreateForumCommentRequest) (dto.CreateForumCommentResponse, error)
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

func (v *viewService) CreateForumComment(ctx dto.SessionContext, req dto.CreateForumCommentRequest) (dto.CreateForumCommentResponse, error) {
	var (
		res dto.CreateForumCommentResponse
	)

	forumComment := models.ForumComment {
		ForumID: req.ForumID,
		Comment: req.Comment,
		UserID: ctx.User.ID,
	}

	err := v.repository.ForumCommentRepository.CreateForumComment(forumComment)
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


func NewViewService(repository repository.Holder, shared shared.Holder) ViewService {
	return &viewService{
		repository: repository,
		shared:      shared,
	}
}