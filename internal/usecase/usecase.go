package usecase

import (
	"context"
	"errors"
	"fmt"
	"io"
	"time"

	"gosocial/config"
	"gosocial/internal/entity"
	mongostore "gosocial/internal/store/mongo"
	"gosocial/pkg/auth"
	"gosocial/pkg/hashpass"

	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserStore interface {
	Save(info entity.User) (primitive.ObjectID, error)
	Get(id string) (mongostore.UserDoc, error)
	GetByEmail(email string) (mongostore.UserDoc, error)
	Update(id string, info entity.User) error
}

type ImageStore interface {
	Save(info entity.Image) error
	Get(username string) ([]entity.Image, error)
}

type ImageBucket interface {
	SaveImage(ctx context.Context, name string, r io.Reader) (entity.Image, error)
}

func New(config config.Config, userStore UserStore, imageStore ImageStore, imageBucket ImageBucket) *ucImplement {
	return &ucImplement{
		config:     config,
		userStore:  userStore,
		imageStore: imageStore,
		imgBucket:  imageBucket,
	}
}

type ucImplement struct {
	config     config.Config
	userStore  UserStore
	imageStore ImageStore
	imgBucket  ImageBucket
}

func (uc *ucImplement) Register(ctx context.Context, req *entity.RegisterRequest) (*entity.RegisterResponse, error) {

	userId, err := uc.userStore.Save(entity.User{
		Email:          req.Email,
		HashedPassword: hashpass.HashPassword(req.Password),
		Active:         true,
	})
	if err != nil {
		log.Error(err)
		// fmt.Println(err)
		return nil, err //fmt.Errorf("save user failed")
	}

	// TODO: create Profile with userId

	return &entity.RegisterResponse{
		Id: userId.Hex(),
	}, nil
}

func (uc *ucImplement) Login(ctx context.Context, req *entity.LoginRequest) (*entity.LoginResponse, error) {

	user, err := uc.userStore.GetByEmail(req.Email)
	if err != nil {
		return nil, ErrInvalidUserOrPassword
	}

	hashedPassword := hashpass.HashPasswordLogin(req.Password, user.HashedPassword)

	fmt.Println("hashedPassword", hashedPassword)

	if user.HashedPassword != hashedPassword {
		return nil, ErrInvalidUserOrPassword
	}

	token, err := auth.GenerateToken(user.Email, 24*time.Hour)

	if err != nil {
		return nil, ErrGenerateToken
	}

	return &entity.LoginResponse{Token: token}, nil
}

func (uc *ucImplement) Self(ctx context.Context, req *entity.SelfRequest) (*entity.SelfResponse, error) {

	user, err := uc.userStore.Get(req.Userid)
	// images, _ := uc.imageStore.Get(req.Username)

	if err != nil {
		return nil, err
	}

	// drop path and internal data
	// var imgRes []entity.ImageResponse
	// for _, img := range images {
	// 	imgRes = append(imgRes, entity.ImageResponse{
	// 		FileName: img.FileName,
	// 		URL:      uc.GetFullURL(img.URL),
	// 	})
	// }

	return &entity.SelfResponse{
		Email: user.Email,
		// FullName: user.FullName,
		// Address:  user.Address,
		// Images:   imgRes,
	}, nil
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

func (uc *ucImplement) ChangePassword(ctx context.Context, req *entity.ChangePasswordRequest) (*entity.ChangePasswordResponse, error) {
	user, err := uc.userStore.Get(req.Userid)
	if err != nil {
		return nil, err
	}
	fmt.Println(req)

	if req.RepeatPassword != req.NewPassword {
		return nil, ErrRepeatPassword
	}

	if req.NewPassword == req.CurrentPassword {
		return nil, ErrSamePassword
	}

	loginHashedPassword := hashpass.HashPasswordLogin(req.CurrentPassword, user.HashedPassword)

	if user.HashedPassword != loginHashedPassword {
		return nil, ErrInvalidUserOrPassword
	}

	// assign new Password
	user.HashedPassword = hashpass.HashPasswordLogin(req.NewPassword, user.HashedPassword)

	// fmt.Println("save user", user)

	if err := uc.userStore.Update(req.Userid, entity.User{
		Email:          user.Email,
		HashedPassword: user.HashedPassword,
		Active:         user.Active,
	}); err != nil {
		return nil, err
	}

	return &entity.ChangePasswordResponse{Success: true}, nil
}

var ErrSamePassword = errors.New("new password is same as current password")
var ErrRepeatPassword = errors.New("repeat password does not match")
var ErrInvalidUserOrPassword = errors.New("invalid username or password")
var ErrGenerateToken = errors.New("generate token failed")
