package repo

import (
	"errors"
	"golang_study/blog/internal/common"
	"golang_study/blog/internal/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	GormRepository
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{GormRepository{DB: db}}
}

func (r *UserRepository) GetByEmail(email string) (*model.User, error) {
	var user model.User
	err := r.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.NewBusinessError(common.ErrUserNotFound)
		}
		return nil, common.NewBusinessError(common.ErrDatabase)
	}

	return &user, err
}

func (r *UserRepository) ListActiveUsers() ([]model.User, error) {
	var users []model.User
	err := r.DB.Where("status = 1").Find(&users).Error
	return users, err
}
