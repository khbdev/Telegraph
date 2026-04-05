package domain

import (
	"context"

)

type DataRepository interface {
	Create(ctx context.Context, data *models.Data) (*models.Data, error)
	Update(ctx context.Context, data *models.Data) (*models.Data, error)
}