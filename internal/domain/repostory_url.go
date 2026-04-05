package domain

import (
	"context"
	"telegraph-clone/internal/models"
)

type URLRepository interface {
	Create(ctx context.Context, url *models.URL) error
	GetByTitle(ctx context.Context, title string) (*models.URL, error)
}