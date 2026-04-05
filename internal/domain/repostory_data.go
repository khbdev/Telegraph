package domain

import (
	"context"
	"telegraph-clone/internal/models"
)

type DataRepository interface {
	Create(ctx context.Context, data *models.Data) (*)
	Update(ctx context.Context, data *models.Data) error
}