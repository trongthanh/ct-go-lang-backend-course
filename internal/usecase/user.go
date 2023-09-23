package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	"gosocial/internal/entity"
	mongostore "gosocial/internal/store/mongo"
	"gosocial/pkg/auth"
	"gosocial/pkg/hashpass"

	"github.com/labstack/gommon/log"
)

type UserStore interface {
	Save(info entity.User) (mongostore.UserDoc, error)
	Get(id string) (mongostore.UserDoc, error)
	GetByEmail(email string) (mongostore.UserDoc, error)
	Update(id string, info entity.User) error
}

type ProfileStore interface {
	Save(info entity.Profile) (mongostore.ProfileDoc, error)
	Get(userid string) (mongostore.ProfileDoc, error)
	Update(userid string, profile entity.Profile) error
}

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
		Id: user.Id.Hex(),
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

	token, tokenErr := auth.GenerateToken(user.Id.Hex(), 24*time.Hour)

	if tokenErr != nil {
		return nil, ErrGenerateToken
	}

	profileDoc, profileErr := uc.profileStore.Get(user.Id.Hex())
	var profile entity.Profile
	// if this is the first time login, create profileDoc
	if profileErr != nil {
		log.Info("First login, create profile for user")
		newProfile, newProfileErr := uc.profileStore.Save(entity.Profile{
			Userid: user.Id.Hex(),
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

func (uc *ucImplement) Self(ctx context.Context, req *entity.SelfRequest) (*entity.SelfResponse, error) {

	profile, profileErr := uc.profileStore.Get(req.Userid)

	if profileErr != nil {
		return nil, profileErr
	}

	return &entity.SelfResponse{
		Profile: profile.ToProfile(),
	}, nil
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
