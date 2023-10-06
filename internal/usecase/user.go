package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	"gosocial/internal/entity"
	"gosocial/pkg/auth"
	"gosocial/pkg/hashpass"

	"github.com/labstack/gommon/log"
)

func (uc *ucImplement) Register(ctx context.Context, req *entity.RegisterRequest) (*entity.RegisterResponse, error) {

	user, err := uc.userStore.Save(entity.User{
		Email:          req.Email,
		HashedPassword: hashpass.HashPassword(req.Password),
		Active:         true,
	})
	if err != nil {
		log.Error(err)
		// fmt.Println(err)
		return nil, err //fmt.Errorf("save user failed")
	}

	return &entity.RegisterResponse{
		Id: user.DocId.Hex(),
	}, nil
}

func (uc *ucImplement) Login(ctx context.Context, req *entity.LoginRequest) (*entity.LoginResponse, error) {

	user, userErr := uc.userStore.GetByEmail(req.Email)
	if userErr != nil {
		return nil, ErrInvalidUserOrPassword
	}

	hashedPassword := hashpass.HashPasswordLogin(req.Password, user.HashedPassword)

	fmt.Println("hashedPassword", hashedPassword)

	if user.HashedPassword != hashedPassword {
		return nil, ErrInvalidUserOrPassword
	}

	token, tokenErr := auth.GenerateToken(user.DocId.Hex(), 24*time.Hour)

	if tokenErr != nil {
		return nil, ErrGenerateToken
	}

	profileDoc, profileErr := uc.profileStore.Get(user.DocId.Hex())
	var profile entity.Profile
	// if this is the first time login, create profileDoc
	if profileErr != nil {
		log.Info("First login, create profile for user")
		newProfile, newProfileErr := uc.profileStore.Save(entity.Profile{
			Userid: user.DocId.Hex(),
		})
		if newProfileErr != nil {
			log.Error(newProfileErr)
			return nil, newProfileErr
		}
		profile = newProfile.ToProfile()
	} else {
		profile = profileDoc.ToProfile()
	}

	return &entity.LoginResponse{Token: token, Profile: profile}, nil
}

func (uc *ucImplement) Self(ctx context.Context, req *entity.SelfRequest) (*entity.ProfileResponse, error) {

	profile, profileErr := uc.profileStore.Get(req.Userid)

	if profileErr != nil {
		return nil, profileErr
	}

	return &entity.ProfileResponse{
		Profile: profile.ToProfile(),
	}, nil
}

var ErrRepeatPassword = errors.New("repeat password does not match")
var ErrInvalidUserOrPassword = errors.New("invalid username or password")
var ErrGenerateToken = errors.New("generate token failed")
