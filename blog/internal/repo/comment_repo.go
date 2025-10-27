package repo

import (
	"errors"
	"golang_study/blog/internal/common"
	"golang_study/blog/internal/model"

	"gorm.io/gorm"
)

type CommentRepository struct {
	GormRepository
}

func NewCommentRepository(db *gorm.DB) *CommentRepository {
	return &CommentRepository{GormRepository{DB: db}}
}

func (r *CommentRepository) SaveComment(comment *model.Comment) (bool, error) {
	err := r.DB.Create(comment).Error
	if err != nil {
		// 需要记录日志，同时返回堆栈
		return false, common.NewBusinessError(common.ErrDatabaseBiz)
	}

	return true, nil
}

func (r *CommentRepository) GetCommentsByPostId(postId uint) ([]model.Comment, error) {
	var comments []model.Comment
	err := r.DB.First(&comments, "post_id = ?", postId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 如果没有找到记录，返回空数组和nil错误
			return []model.Comment{}, nil
		}
		// 如果查询出错，返回空数组和错误
		return []model.Comment{}, err
	}
	return comments, err
}
