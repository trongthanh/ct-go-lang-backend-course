package controller

import (
	"context"
	"gosocial/internal/entity"
)

type UseCase interface {
	Login(ctx context.Context, req *entity.LoginRequest) (*entity.LoginResponse, error)
	Register(ctx context.Context, req *entity.RegisterRequest) (*entity.RegisterResponse, error)
	Self(ctx context.Context, req *entity.SelfRequest) (*entity.SelfResponse, error)
	UploadImage(ctx context.Context, req *entity.UploadImageRequest) (*entity.UploadImageResponse, error)
	ChangePassword(ctx context.Context, req *entity.ChangePasswordRequest) (*entity.ChangePasswordResponse, error)
}

func New(uc UseCase) *Handler {
	return &Handler{uc: uc}
}

type Handler struct {
	uc UseCase
}

func NewResponsePayload(status string, data interface{}) *entity.Response {
	return &entity.Response{
		Status: status,
		Data:   data,
	}
}

var ctx = context.TODO()
