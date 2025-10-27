package post_dto

type PostDTO struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type UpdatePostDTO struct {
	PostDTO
	Id     uint `json:"id" binding:"required"`
	UserId uint
}

type DeletePostDTO struct {
	Id     uint `json:"id" binding:"required"`
	UserId uint
}
