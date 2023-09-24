package repository

import (
	"gorm.io/gorm"
)

type repository[T any] struct {
	db *gorm.DB
}

func (r *repository[T]) List() ([]T, error) {
	var entries []T
	if err := r.db.Find(&entries).Error; err != nil {
		return nil, err
	}
	return entries, nil
}

func (r *repository[T]) Add(e *T) error {
	return r.db.Create(e).Error
}

func (r *repository[T]) Remove(e *T) error {
	return r.db.Delete(e).Error
}
