package localbucket

import (
	"context"
	"fmt"
	"gosocial/internal/entity"
	"io"
	"os"
	"sync"
)

type LocalBucket struct {
	mu   sync.Mutex
	data map[string]entity.Image
}

func New() *LocalBucket {
	return &LocalBucket{
		data: make(map[string]entity.Image),
	}
}

func (lb *LocalBucket) SaveImage(ctx context.Context, fileName string, fileReader io.Reader) (entity.Image, error) {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	localPath := getPublicFolder() + "/" + fileName
	// Destination
	dst, err := os.Create(localPath)
	if err != nil {
		return entity.Image{}, err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, fileReader); err != nil {
		return entity.Image{}, err
	}

	// store metadata
	Image := entity.Image{
		URL:      "http://localhost:8090/" + fileName,
		Filename: fileName,
		Path:     localPath,
	}
	lb.data[fileName] = Image

	fmt.Println("Image saved to: " + Image.Path)

	return Image, nil
}

func getPublicFolder() string {
	// Get the current working directory
	wd, err := os.Getwd()
	if err != nil {
		return ""
	}

	fmt.Println("Current working directory: " + wd)

	// Append the public folder
	return wd + "/www/public"
}
