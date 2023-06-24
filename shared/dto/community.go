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
)