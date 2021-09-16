package repository

import (
	"bytes"
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	uuid "github.com/satori/go.uuid"
)

const (
	imagesBucket = "images"
)

type imageAWSRepository struct {
	s3 *s3.S3
}

func NewImageAWSRepository(s3 *s3.S3) *imageAWSRepository {
	return &imageAWSRepository{s3: s3}
}

func (i *imageAWSRepository) PutObject(ctx context.Context, data []byte, fileType string) (string, error) {
	newFilename := uuid.NewV4().String()
	key := i.getFileKey(newFilename, fileType)

	object, err := i.s3.PutObjectWithContext(ctx, &s3.PutObjectInput{
		Body:   bytes.NewReader(data),
		Bucket: aws.String(imagesBucket),
		Key:    aws.String(key),
		ACL:    aws.String(s3.BucketCannedACLPublicRead),
	})
	if err != nil {
		return "", fmt.Errorf("s3.PutObjectWithContext %w", err)
	}

	log.Printf("object : %-v", object)

	return i.getFilePublicURL(key), err
}

func (i *imageAWSRepository) GetObject(ctx context.Context, key string) (*s3.GetObjectOutput, error) {
	obj, err := i.s3.GetObjectWithContext(ctx, &s3.GetObjectInput{
		Bucket: aws.String(imagesBucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, fmt.Errorf("s3.GetObjectWithContext %w", err)
	}

	return obj, nil
}

func (i *imageAWSRepository) DeleteObject(ctx context.Context, key string) error {

	_, err := i.s3.DeleteObjectWithContext(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(imagesBucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return fmt.Errorf("s3.DeleteObjectWithContext %w", err)
	}

	return nil
}

func (i *imageAWSRepository) getFileKey(fileID string, fileType string) string {
	return fmt.Sprintf("%s.%s", fileID, fileType)
}

func (i *imageAWSRepository) getFilePublicURL(key string) string {
	return "/" + imagesBucket + "/" + key
}
