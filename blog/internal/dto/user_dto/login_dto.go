package user_dto

type LoginDTO struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	UserID      uint
	Username    string
	Email       string
	AccessToken string
	ExpiresAt   int64
	TokenType   string
}
