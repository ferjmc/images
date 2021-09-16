package image

import (
	"context"

	"github.com/ferjmc/images/internal/models"
	uuid "github.com/satori/go.uuid"
)

type GormRepository interface {
	Create(ctx context.Context, msg *models.Image) (*models.Image, error)
	GetImageByID(ctx context.Context, imageID uuid.UUID) (*models.Image, error)
}
