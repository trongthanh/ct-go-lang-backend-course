package googlestorage

import (
	"context"
	"fmt"
	"gosocial/internal/entity"
	"io"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

func New(ctx context.Context, bucketName, credentialsFile string) *GoogleStorageClient {
	storageClient, err := storage.NewClient(ctx, option.WithCredentialsFile(credentialsFile))
	if err != nil {
		panic(err.Error())
	}

	return &GoogleStorageClient{
		Bucket:        bucketName,
		StorageClient: storageClient,
	}
}

type GoogleStorageClient struct {
	Bucket        string
	StorageClient *storage.Client
}

func (c *GoogleStorageClient) SaveImage(ctx context.Context, fileName string, fileReader io.Reader) (entity.Image, error) {
	bucket := c.StorageClient.Bucket(c.Bucket)
	// Create an object handle for the new file
	obj := bucket.Object(fileName)

	// Create an object writer
	writer := obj.NewWriter(ctx)

	// Copy the contents of the file reader to the object writer
	if _, err := io.Copy(writer, fileReader); err != nil {
		return entity.Image{}, err
	}

	// Close the object writer to finalize the upload
	if err := writer.Close(); err != nil {
		return entity.Image{}, err
	}

	// Set the object's ACL to make it publicly readable
	if err := obj.ACL().Set(ctx, storage.AllUsers, storage.RoleReader); err != nil {
		return entity.Image{}, err
	}

	attrs, _ := obj.Attrs(ctx)

	// store metadata
	Image := entity.Image{
		URL:      fmt.Sprintf("https://storage.googleapis.com/%s/%s", c.Bucket, fileName),
		Filename: fileName,
		Path:     fmt.Sprintf("https://storage.cloud.google.com/%s/%s", c.Bucket, fileName),
		Size:     attrs.Size,
	}

	return Image, nil
}
