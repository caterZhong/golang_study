package repo

import (
	"errors"
	"golang_study/blog/internal/common"
	"golang_study/blog/internal/dto"
	"golang_study/blog/internal/model"

	"gorm.io/gorm"
)

type PostRepository struct {
	GormRepository
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{GormRepository{DB: db}}
}

func (r *PostRepository) SavePost(post *model.Post) (bool, error) {
	err := r.DB.Create(post).Error
	if err != nil {
		// 需要记录日志，同时返回堆栈
		return false, common.NewBusinessError(common.ErrDatabaseBiz)
	}

	return true, nil
}

func (r *PostRepository) GetById(id uint) (*model.Post, error) {
	var post model.Post
	err := r.DB.First(&post, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.NewBusinessError(common.ErrPostNotFound)
		}
		return nil, common.NewBusinessError(common.ErrDatabase)
	}
	return &post, err
}

func (r *PostRepository) Page(pagine dto.Pagination) ([]model.Post, int64, error) {
	var post model.Post

	var count int64
	err := r.DB.Model(&post).Count(&count).Error
	if err != nil {
		return []model.Post{}, 0, common.NewBusinessError(common.ErrDatabase)
	}

	var posts []model.Post
	listErr := r.DB.Model(&post).Order("id desc").Limit(pagine.GetLimit()).Offset(pagine.GetOffset()).Find(&posts).Error
	if listErr != nil {
		return []model.Post{}, 0, common.NewBusinessError(common.ErrDatabase)
	}
	return posts, count, nil
}

func (r *PostRepository) Update(post model.Post) (bool, error) {
	if post.ID == 0 {
		return false, nil
	}
	err := r.DB.Model(post).
		Where("id = ?", post.ID).
		Updates(post).Error

	if err != nil {
		return false, common.NewBusinessError(common.ErrDatabase)
	}
	return true, nil
}
