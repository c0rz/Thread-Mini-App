package models

type RegisterAccount struct {
	Name     string `json:"name" validate:"required" binding:"required"`
	Email    string `json:"email" validate:"required" binding:"required,email"`
	Password string `json:"password" validate:"required" binding:"required"`
}

type LoginInput struct {
	Email    string `json:"email" validate:"required" binding:"required,email"`
	Password string `json:"password" validate:"required" binding:"required"`
}

type BlogInput struct {
	Title      string `json:"title" validate:"required" binding:"required"`
	Text       string `json:"text" validate:"required" binding:"required"`
	CategoryID string `json:"category_id" validate:"required" binding:"required"`
}

type CategoryInput struct {
	Name string `json:"name" validate:"required" binding:"required"`
}

type CommentInput struct {
	IDPost  int    `json:"id_post" validate:"required" binding:"required"`
	Comment string `json:"comment" validate:"required" binding:"required"`
}
