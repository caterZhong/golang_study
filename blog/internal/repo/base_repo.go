package repo

import (
	"gorm.io/gorm"
)

type BaseRepository interface {
	Create(model interface{}) error
	Update(model interface{}) error
	Delete(model interface{}, id uint) error
	GetByID(model interface{}, id uint) error
}

type GormRepository struct {
	DB *gorm.DB
}

func (r *GormRepository) Create(model interface{}) error {
	return r.DB.Create(model).Error
}

func (r *GormRepository) Update(model interface{}) error {
	return r.DB.Save(model).Error
}

func (r *GormRepository) Delete(model interface{}, id uint) error {
	return r.DB.Delete(model, id).Error
}

func (r *GormRepository) GetByID(model interface{}, id uint) error {
	return r.DB.First(model, id).Error
}
