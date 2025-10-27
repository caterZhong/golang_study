package controller

import (
	"golang_study/blog/internal/common"
	"golang_study/blog/internal/dto"
	"golang_study/blog/internal/dto/post_dto"
	"golang_study/blog/internal/service"

	"github.com/gin-gonic/gin"
)

type PostController struct {
	postService *service.PostService
}

func NewPostController(postService *service.PostService) *PostController {
	return &PostController{
		postService: postService,
	}
}

// 发布文章
func (p *PostController) CreatePost(c *gin.Context) {
	userIdValue, _ := c.Get("userID")

	var postDTO post_dto.PostDTO
	if err := c.ShouldBind(&postDTO); err != nil {
		common.BadRequest(c, common.ErrInvalidParams.Message)
		return
	}
	// 类型断言
	userId, ok := userIdValue.(uint)
	if !ok {
		common.BadRequest(c, common.ErrInvalidParams.Message)
		return
	}
	p.postService.SavePost(postDTO.Title, postDTO.Content, uint(userId))
	common.Success(c, dto.True)
	return
}

// 删除文章
func (p *PostController) DeletePost(c *gin.Context) {
	userIdValue, _ := c.Get("userID")

	var postDTO post_dto.DeletePostDTO
	if err := c.ShouldBind(&postDTO); err != nil {
		common.BadRequest(c, common.ErrInvalidParams.Message)
		return
	}
	// 类型断言
	userId, ok := userIdValue.(uint)
	if !ok {
		common.BadRequest(c, common.ErrInvalidParams.Message)
		return
	}
	postDTO.UserId = userId
	if err := p.postService.DeletePost(postDTO); err != nil {
		common.BadRequest(c, "删除文章失败")
		return
	}
	common.Success(c, dto.True)
	return

}

// 获取文章
func (p *PostController) GetPost(c *gin.Context) {
	return
}

// 获取文章
func (p *PostController) PagePost(c *gin.Context) {
	var pageDTO dto.Pagination
	if err := c.ShouldBindQuery(&pageDTO); err != nil {
		common.BadRequest(c, common.ErrInvalidParams.Message)
		return
	}
	// 类型断言
	// userId, ok := userIdValue.(uint)
	// if !ok {
	// 	common.BadRequest(c, common.ErrInvalidParams.Message)
	// 	return
	// }
	// 业务逻辑
	posts, total, err := p.postService.PagePost(pageDTO)

	if err != nil {
		common.BadRequest(c, err.Error())
		return
	}

	// 返回分页响应
	response := dto.NewPageResponse(posts, total, pageDTO)
	common.Success(c, response)
	return
}

// 修改文章
func (p *PostController) UpdatePost(c *gin.Context) {
	userIdValue, _ := c.Get("userID")

	var postDTO post_dto.UpdatePostDTO
	if err := c.ShouldBind(&postDTO); err != nil {
		common.BadRequest(c, common.ErrInvalidParams.Message)
		return
	}
	// 类型断言
	userId, ok := userIdValue.(uint)
	if !ok {
		common.BadRequest(c, common.ErrInvalidParams.Message)
		return
	}
	postDTO.UserId = userId
	if result, err := p.postService.UpdatePost(postDTO); !result || err != nil {
		common.BadRequest(c, err.Error())
		return
	}
	common.Success(c, dto.True)
	return

}
