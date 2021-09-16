package repository

import (
	"context"
	"errors"

	"github.com/ferjmc/images/internal/models"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type imagePGRepository struct {
	gormDB *gorm.DB
}

func NewImagePGRepository(gormDB *gorm.DB) *imagePGRepository {
	return &imagePGRepository{gormDB: gormDB}
}

func (i *imagePGRepository) Create(ctx context.Context, msg *models.Image) (*models.Image, error) {
	res := i.gormDB.Create(&msg)
	if res.RowsAffected <= 0 {
		return nil, errors.New("error al insertar datos")
	}
	return msg, nil
}
func (i *imagePGRepository) GetImageByID(ctx context.Context, imageID uuid.UUID) (*models.Image, error) {
	var img models.Image
	res := i.gormDB.Find(&img, imageID)
	if res.RowsAffected <= 0 {
		return nil, errors.New("no se encontro imagen")
	}

	return &img, nil
}
