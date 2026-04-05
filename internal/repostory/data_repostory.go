package repository

import (
	"context"
	"telegraph-clone/internal/domain"
	"telegraph-clone/internal/models"

	"gorm.io/gorm"
)

type dataRepo struct {
	db *gorm.DB
}

func NewDataRepository(db *gorm.DB) domain.DataRepository {
	return &dataRepo{db: db}
}


func (r *dataRepo) Create(ctx context.Context, data *models.Data) () {
	return r.db.WithContext(ctx).Create(data).Error
}


func (r *dataRepo) Update(ctx context.Context, data *models.Data) error {
	return r.db.WithContext(ctx).Save(data).Error
}