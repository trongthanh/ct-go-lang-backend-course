package imagebucket

import (
	"context"
	"io"
)

type ImageBucket struct {
	data map[string][]byte
}

func New() *ImageBucket {
	//TODO
	return nil
}

func (ib *ImageBucket) SaveImage(ctx context.Context, name string, r io.Reader) (string, error) {
	return "", nil
}
