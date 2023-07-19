package usecase

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"time"

	"thanhtran-s04-2/entity"
	"thanhtran-s04-2/pkg/auth"
)

type UserStore interface {
	Save(info entity.UserInfo) error
	Get(username string) (entity.UserInfo, error)
}

type ImageBucket interface {
	SaveImage(ctx context.Context, name string, r io.Reader) (string, error)
}

func New(userStore UserStore, imageBucket ImageBucket) *ucImplement {
	return &ucImplement{
		store:     userStore,
		imgBucket: imageBucket,
	}
}

type ucImplement struct {
	store     UserStore
	imgBucket ImageBucket
}

func (uc *ucImplement) Register(ctx context.Context, req *entity.RegisterRequest) (*entity.RegisterResponse, error) {
	if err := uc.store.Save(entity.UserInfo{
		Username: req.Username,
		Password: req.Password,
		FullName: req.FullName,
		Address:  req.Address,
	}); err != nil {
		return nil, err
	}

	return &entity.RegisterResponse{Username: req.Username}, nil
}

func (uc *ucImplement) Login(ctx context.Context, req *entity.LoginRequest) (*entity.LoginResponse, error) {

	user, err := uc.store.Get(req.Username)
	if err != nil {
		return nil, ErrInvalidUserOrPassword
	}

	if user.Password != req.Password {
		return nil, ErrInvalidUserOrPassword
	}

	token, err := auth.GenerateToken(user.Username, 24*time.Hour)

	if err != nil {
		return nil, ErrGenerateToken
	}

	return &entity.LoginResponse{Token: token}, nil
}

func (uc *ucImplement) Self(ctx context.Context, req *entity.SelfRequest) (*entity.SelfResponse, error) {

	user, err := uc.store.Get(req.Username)

	if err != nil {
		return nil, err
	}

	return &entity.SelfResponse{
		Username: user.Username,
		FullName: user.FullName,
		Address:  user.Address,
	}, nil
}

func (uc *ucImplement) UploadImage(ctx context.Context, req *entity.UploadImageRequest) (*entity.UploadImageResponse, error) {
	file := req.File

	src, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create(getPublicFolder() + "/images/" + file.Filename)
	if err != nil {
		return nil, err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return nil, err
	}

	url := "http://localhost:8090/images/" + file.Filename

	return &entity.UploadImageResponse{URL: url}, nil
}

var ErrInvalidUserOrPassword = errors.New("invalid username or password")
var ErrGenerateToken = errors.New("generate token failed")

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
