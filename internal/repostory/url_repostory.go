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


func (r *urlRepo) GetByURL(ctx context.Context, urlStr string) (*models.URL, error) {
    var url models.URL
    err := r.db.WithContext(ctx).
        Preload("Data").              // Data ob'ektini yuklash
        Where("url = ?", urlStr).     // URL bo‘yicha qidiruv
        First(&url).Error
    if err != nil {
        return nil, err
    }
    return &url, nil
}