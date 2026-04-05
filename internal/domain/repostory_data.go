package domain

import (
	"context"
	"your_project/internal/models"
)

type DataRepository interface {
	Create(ctx context.Context, data *models.Data) (*models.Data, error)
	Update(ctx context.Context, data *models.Data) (*models.Data, error)
}