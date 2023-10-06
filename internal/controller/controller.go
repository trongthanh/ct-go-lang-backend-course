package controller

import (
	"context"
	"gosocial/internal/entity"
	"io"
)

type UseCase interface {
	Login(ctx context.Context, req *entity.LoginRequest) (*entity.LoginResponse, error)
	Register(ctx context.Context, req *entity.RegisterRequest) (*entity.RegisterResponse, error)
	Self(ctx context.Context, req *entity.SelfRequest) (*entity.ProfileResponse, error)
	// UploadImage(ctx context.Context, req *entity.UploadImageRequest) (*entity.UploadImageResponse, error)
	// ChangePassword(ctx context.Context, req *entity.ChangePasswordRequest) (*entity.ChangePasswordResponse, error)
	GetProfiles(ctx context.Context, req *entity.ProfilesRequest) (*entity.ProfilesResponse, error)
	GetPosts(ctx context.Context, req *entity.PostsRequest) (*entity.PostsResponse, error)
	GetPostsByUser(ctx context.Context, req *entity.PostsByUserRequest) (*entity.PostsResponse, error)
	CreatePost(ctx context.Context, req *entity.CreatePostRequest) (*entity.CreatePostResponse, error)
	DeletePost(ctx context.Context, req *entity.DeletePostRequest) (*entity.DeletePostResponse, error)
	LikePost(ctx context.Context, req *entity.LikePostRequest) (*entity.LikePostResponse, error)
}

type ImageBucket interface {
	SaveImage(ctx context.Context, name string, r io.Reader) (string, error)
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

// func (h *Handler) UploadImage(c echo.Context) error {
// 	// userid from jwt
// 	userid := c.Get("userid").(string)
// 	// Source
// 	fileHeader, err := c.FormFile("file")
// 	if err != nil {
// 		return err
// 	}
//
// 	file, err := fileHeader.Open()
// 	defer file.Close()
// 	if err != nil {
// 		return err
// 	}
//
// 	uploadImageReq := &entity.UploadImageRequest{Userid: userid, Filename: fileHeader.Filename, File: file}
//
// 	resp, err := h.uc.UploadImage(ctx, uploadImageReq)
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, NewResponsePayload("error", err.Error()))
// 	}
//
// 	return c.JSON(http.StatusOK, NewResponsePayload("success", resp))
//
// }
