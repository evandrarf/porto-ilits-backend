package repository

import (
	"github.com/evandrarf/porto-ilits-backend/internal/pkg/meta"
	"gorm.io/gorm"
)

type Repository[T any] interface {
	Create(tx *gorm.DB, entity *T) error
	Update(tx *gorm.DB, entity *T) error
	Delete(tx *gorm.DB, entity *T) error
	CountById(tx *gorm.DB, id any) (int64, error)
	FindById(tx *gorm.DB, entity *T, id any) error
	Exists(tx *gorm.DB, id any) bool
	All(tx *gorm.DB, entity *[]T, filter *meta.Filter, paginate *meta.Pagination, columns ...string) error
	BulkExists(tx *gorm.DB, entity *T, ids []int) bool
	BulkCreate(tx *gorm.DB, entities []T) error
	BulkDelete(tx *gorm.DB, entity *T, ids []int) error
	Save(tx *gorm.DB, entity *T) error
}

type repository[T any] struct {
	DB *gorm.DB
}

func (r *repository[T]) Create(tx *gorm.DB, entity *T) error {
	if tx == nil {
		tx = r.DB
	}
	return tx.Create(entity).Error
}

func (r *repository[T]) Update(tx *gorm.DB, entity *T) error {
	if tx == nil {
		tx = r.DB
	}
	return tx.Debug().Updates(entity).Error
}

func (r *repository[T]) Delete(tx *gorm.DB, entity *T) error {
	if tx == nil {
		tx = r.DB
	}
	return tx.Delete(entity).Error
}

func (r *repository[T]) CountById(tx *gorm.DB, id any) (int64, error) {
	if tx == nil {
		tx = r.DB
	}
	var total int64
	err := tx.Model(new(T)).Where("id = ?", id).Count(&total).Error
	return total, err
}

func (r *repository[T]) FindById(tx *gorm.DB, entity *T, id any) error {
	if tx == nil {
		tx = r.DB
	}
	return tx.Where("id = ?", id).Take(entity).Error
}

func (r *repository[T]) Exists(tx *gorm.DB, id any) bool {
	if tx == nil {
		tx = r.DB
	}
	count, err := r.CountById(tx, id)
	if err != nil {
		return false
	}
	return count > 0
}

func (r *repository[T]) All(tx *gorm.DB, entity *[]T, filter *meta.Filter, pagination *meta.Pagination, columns ...string) error {
	if tx == nil {
		tx = r.DB
	}

	query := tx.Model(entity)

	if filter != nil {
		query = filter.Filterize(query)
		if query.Error != nil {
			return query.Error
		}
	}

	if len(columns) > 0 {
		query = query.Select(columns)
	}

	if pagination != nil {
		query := pagination.Paginate(query)
		if query.Error != nil {
			return query.Error
		}
	}

	return query.Find(entity).Error
}

func (r *repository[T]) BulkExists(tx *gorm.DB, entity *T, ids []int) bool {
	if tx == nil {
		tx = r.DB
	}
	var count int64
	err := tx.Model(entity).
		Where("id IN ?", ids).
		Count(&count).Error
	if err != nil {
		return false
	}
	return count == int64(len(ids))
}

func (r *repository[T]) BulkCreate(tx *gorm.DB, entities []T) error {
	if tx == nil {
		tx = r.DB
	}
	return tx.Create(&entities).Error
}

func (r *repository[T]) BulkDelete(tx *gorm.DB, entity *T, ids []int) error {
	if tx == nil {
		tx = r.DB
	}
	return tx.Where("id IN ?", ids).Delete(entity).Error
}

func (r *repository[T]) Save(tx *gorm.DB, entity *T) error {
	if tx == nil {
		tx = r.DB
	}
	return tx.Save(entity).Error
}
