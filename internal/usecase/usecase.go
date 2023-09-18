package usecase

import (
	"context"
	"io"

	"gosocial/config"
	"gosocial/internal/entity"
)

type ImageStore interface {
	Save(info entity.Image) error
	Get(username string) ([]entity.Image, error)
}

type ImageBucket interface {
	SaveImage(ctx context.Context, name string, r io.Reader) (entity.Image, error)
}

func New(config config.Config, userStore UserStore, profileStore ProfileStore, imageStore ImageStore, imageBucket ImageBucket) *ucImplement {
	return &ucImplement{
		config:       config,
		userStore:    userStore,
		profileStore: profileStore,
		imageStore:   imageStore,
		imgBucket:    imageBucket,
	}
}

type ucImplement struct {
	config       config.Config
	userStore    UserStore
	profileStore ProfileStore
	imageStore   ImageStore
	imgBucket    ImageBucket
}

func (uc *ucImplement) UploadImage(ctx context.Context, req *entity.UploadImageRequest) (*entity.UploadImageResponse, error) {

	// store image to bucket
	imgInfo, err := uc.imgBucket.SaveImage(ctx, req.Filename, req.File)
	if err != nil {
		return nil, err
	}
	imgInfo.URL = "/images/" + imgInfo.Filename
	// store image info to db and associate with user
	uc.imageStore.Save(imgInfo)

	return &entity.UploadImageResponse{URL: uc.GetFullURL(imgInfo.URL)}, nil
}

// GetFullURL return full url from config
func (uc *ucImplement) GetFullURL(url string) string {
	return uc.config.Scheme + uc.config.Host + ":" + uc.config.Port + url
}
