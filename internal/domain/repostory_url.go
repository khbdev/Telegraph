package domain

import (
	"context"
	"telegraph-clone/internal/domain"
)

type URLRepository interface {
	Create(ctx context.Context, url *domain.URL) (*URL, error)
}