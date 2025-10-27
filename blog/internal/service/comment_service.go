package service

import (
	"golang_study/blog/internal/model"
	"golang_study/blog/internal/repo"
)

type CommentService struct {
	commentRepository *repo.CommentRepository
}

func NewCommentService(commentRepository *repo.CommentRepository) *CommentService {
	return &CommentService{
		commentRepository: commentRepository,
	}

}

// 创建文章
func (commentService *CommentService) SaveComment(content string, postId uint, userId uint) (bool, error) {

	comment := model.Comment{
		Content: content,
		PostID:  postId,
		UserID:  userId,
	}
	success, err := commentService.commentRepository.SaveComment(&comment)

	return success, err
}

// 获取文章评论
func (commentService *CommentService) GetCommentsByPostId(id uint) ([]model.Comment, error) {
	return commentService.commentRepository.GetCommentsByPostId(id)

}
