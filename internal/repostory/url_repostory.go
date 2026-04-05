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
		Preload("Data").                      // Data relation ni yuklash
		Joins("Data").                        // GORM relation asosida join qilish
		Where("data.title = ?", title).       // title bo‘yicha filter
		First(&url).Error

	if err != nil {
		if er.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &url, nil
}