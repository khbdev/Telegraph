package repository

import (
	"context"
	"telegraph-clone/internal/domain"
	"telegraph-clone/internal/models"

	"gorm.io/gorm"
)

type urlRepo struct {
	db *gorm.DB
}

func NewURLRepository(db *gorm.DB) domain.URLRepository {
	return &urlRepo{db: db}
}

// CREATE
func (r *urlRepo) Create(ctx context.Context, url *models.URL) error {
	return r.db.WithContext(ctx).Create(url).Error
}


func (r *urlRepo) GetByTitle(ctx context.Context, title string) (*models.URL, error) {
	var url models.URL

	err := r.db.WithContext(ctx).
		Joins("JOIN data ON data.id = urls.data_id").
		Where("data.title = ?", title).
		Preload("Data").
		First(&url).Error

	if err != nil {
		return nil, err
	}

	return &url, nil
}