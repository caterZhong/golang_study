package comment_dto

type CommentDTO struct {
	PostId  uint   `json:"postId" form:"postId" binding:"required"`
	Content string `json:"content" form:"content" binding:"required"`
}

type GetPostCommentDTO struct {
	PostId uint `json:"post_id" form:"postId" binding:"required"`
}
