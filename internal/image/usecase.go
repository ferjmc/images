package image

import (
	"context"

	uuid "github.com/satori/go.uuid"
	"github.com/streadway/amqp"

	"github.com/ferjmc/images/internal/models"
)

type UseCase interface {
	ResizeImage(ctx context.Context, delivery amqp.Delivery) error
	ProcessHotelImage(ctx context.Context, delivery amqp.Delivery) error
	Create(ctx context.Context, delivery amqp.Delivery) error
	GetImageByID(ctx context.Context, imageID uuid.UUID) (*models.Image, error)
}
