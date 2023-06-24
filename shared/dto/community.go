package dto

type (
	// CreateForumRequest CreateForumRequest
	CreateForumRequest struct {
		Title string `json:"title"`
		Description string `json:"description"`
	}
		
	// CreateForumResponse CreateForumResponse
	CreateForumResponse struct {
		ID uint `json:"id"`
		Title string `json:"title"`
		Description string `json:"description"`
		Vote uint `json:"vote"`
		UserID uint `json:"user_id"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}

	// GetForumResponse GetForumResponse
	GetForumResponse struct {
		ID uint `json:"id"`
		Title string `json:"title"`
		Description string `json:"description"`
		Vote uint `json:"vote"`
		UserID uint `json:"user_id"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}

	// CreateForumCommentRequest CreateForumCommentRequest
	CreateForumCommentRequest struct {
		Comment string `json:"comment"`
	}

	// CreateForumCommentResponse CreateForumCommentResponse
	CreateForumCommentResponse struct {
		ID uint `json:"id"`
		ForumID uint `json:"forum_id"`
		Comment string `json:"comment"`
		UserID uint `json:"user_id"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}

	// GetForumCommentResponse GetForumCommentResponse
	GetForumCommentResponse struct {
		ID uint `json:"id"`
		ForumID uint `json:"forum_id"`
		Comment string `json:"comment"`
		UserID uint `json:"user_id"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}
)