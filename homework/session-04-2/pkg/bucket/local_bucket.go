package localbucket

import (
	"context"
	"fmt"
	"io"
	"os"
	"sync"
	"thanhtran-s04-2/entity"
)

type LocalBucket struct {
	mu   sync.Mutex
	data map[string]entity.ImageInfo
}

func New() *LocalBucket {
	return &LocalBucket{
		data: make(map[string]entity.ImageInfo),
	}
}

func (lb *LocalBucket) SaveImage(ctx context.Context, name string, r io.Reader) (entity.ImageInfo, error) {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	localPath := getPublicFolder() + "/images/" + name
	// Destination
	dst, err := os.Create(localPath)
	if err != nil {
		return entity.ImageInfo{}, err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, r); err != nil {
		return entity.ImageInfo{}, err
	}

	// store metadata
	imageInfo := entity.ImageInfo{
		Name: name,
		Path: localPath,
		// Question: not sure if this is the right place to store URL
		URL: "/images/" + name,
	}
	lb.data[name] = imageInfo

	fmt.Println("Image saved to: " + imageInfo.Path)

	return imageInfo, nil
}

func getPublicFolder() string {
	// Get the current working directory
	wd, err := os.Getwd()
	if err != nil {
		return ""
	}

	fmt.Println("Current working directory: " + wd)

	// Append the public folder
	return wd + "/public"
}
