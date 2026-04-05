package domain

import (
	"context"
	"telegraph-clone/internal/models"
)

type DataRepository interface {
	Create(ctx context.Context, data *models.Data) (*models.Data, error)
	Update(ctx context.Context, data *models.Data) (*models.Data, error)
}