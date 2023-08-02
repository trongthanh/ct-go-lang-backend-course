package usecase

import (
	"bytes"
	"context"
	"io"
	"testing"
	"thanhtran/constants"
	"thanhtran/entity"

	"gopkg.in/go-playground/assert.v1"
)

// Mock Image Bucket for testing
type mockImageBucket struct{}

func (m *mockImageBucket) SaveImage(ctx context.Context, name string, r io.Reader) (entity.ImageInfo, error) {
	// Mock implementation to return a dummy image URL
	return entity.ImageInfo{
		Name: name,
		Path: "/os/images/" + name,
		URL:  "/images/" + name,
	}, nil
}

func TestUcImplement_UploadImage(t *testing.T) {
	// Create a test instance of ucImplement with the mock Image Bucket
	uc := &ucImplement{
		imgBucket: &mockImageBucket{},
	}
	// Mock file data (you can provide your own image data)
	mockImageData := []byte{ /* Your mock image data here */ }

	// Create a bytes.Reader as a mock file io.Reader
	mockFile := bytes.NewReader(mockImageData)
	ctx := context.Background()
	req := &entity.UploadImageRequest{
		Filename: "dummy.jpg",
		File:     mockFile,
	}

	// Call the UploadImage method
	resp, err := uc.UploadImage(ctx, req)

	// Assertions
	assert.Equal(t, err, nil)
	assert.NotEqual(t, resp, nil)
	assert.Equal(t, resp.URL, constants.Host+"/images/dummy.jpg")
}
