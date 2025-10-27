package controller

import (
	"golang_study/blog/internal/common"
	"golang_study/blog/internal/dto"
	"golang_study/blog/internal/dto/comment_dto"
	"golang_study/blog/internal/service"

	"github.com/gin-gonic/gin"
)

type CommentController struct {
	commentService *service.CommentService
}

func NewCommentController(commnetService *service.CommentService) *CommentController {
	return &CommentController{
		commentService: commnetService,
	}
}

// 发布评论
func (com *CommentController) CreateComment(c *gin.Context) {
	userIdValue, _ := c.Get("userID")

	var commentDTO comment_dto.CommentDTO
	if err := c.ShouldBind(&commentDTO); err != nil {
		common.BadRequest(c, common.ErrInvalidParams.Message)
		return
	}
	// 类型断言
	userId, ok := userIdValue.(uint)
	if !ok {
		common.BadRequest(c, common.ErrInvalidParams.Message)
		return
	}

	if result, err := com.commentService.SaveComment(commentDTO.Content, commentDTO.PostId, uint(userId)); !result || err != nil {
		common.BadRequest(c, "创建评论失败")
		return
	}
	common.Success(c, dto.True)
	return

}

// 获取评论列表
func (com *CommentController) GetPostComment(c *gin.Context) {
	var dto comment_dto.GetPostCommentDTO
	if err := c.ShouldBindQuery(&dto); err != nil {
		common.BadRequest(c, common.ErrInvalidParams.Message)
		return
	}

	// 业务逻辑
	posts, err := com.commentService.GetCommentsByPostId(dto.PostId)

	if err != nil {
		common.BadRequest(c, err.Error())
		return
	}

	// 返回
	if err != nil {
		common.BadRequest(c, "获取评论失败")
		return
	}
	common.Success(c, posts)
	return
}
