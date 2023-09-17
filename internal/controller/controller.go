package controller

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"gosocial/internal/entity"

	"github.com/labstack/echo/v4"
)

type ImageBucket interface {
	SaveImage(ctx context.Context, name string, r io.Reader) (string, error)
}

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

func (h *Handler) Register(c echo.Context) error {
	var req entity.RegisterRequest
	if err := c.Bind(&req); err != nil {
		// return fmt.Errorf("bind: %w", err)
		return echo.NewHTTPError(http.StatusBadRequest, NewResponsePayload("error", err.Error()))
	}

	if err := c.Validate(req); err != nil {
		return err
	}

	// use case
	resp, err := h.uc.Register(ctx, &req)

	if err != nil {
		// return fmt.Errorf("uc.Register: %w", err)
		return echo.NewHTTPError(http.StatusBadRequest, NewResponsePayload("error", err.Error()))
	}

	return c.JSON(http.StatusOK, NewResponsePayload("success", resp))
}

func (h *Handler) Login(c echo.Context) error {

	loginReq := new(entity.LoginRequest)
	if err := c.Bind(loginReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, NewResponsePayload("error", err.Error()))
	}

	resp, err := h.uc.Login(ctx, loginReq)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, NewResponsePayload("error", err.Error()))
	}

	return c.JSON(http.StatusOK, NewResponsePayload("success", resp))
}

func (h *Handler) Self(c echo.Context) error {

	userid := c.Get("userid").(string)
	selfReq := &entity.SelfRequest{Userid: userid}

	resp, err := h.uc.Self(ctx, selfReq)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, NewResponsePayload("error", err.Error()))
	}

	return c.JSON(http.StatusOK, NewResponsePayload("success", resp))
}

func (h *Handler) UploadImage(c echo.Context) error {
	// userid from jwt
	userid := c.Get("userid").(string)
	// Source
	fileHeader, err := c.FormFile("file")
	if err != nil {
		return err
	}

	file, err := fileHeader.Open()
	defer file.Close()
	if err != nil {
		return err
	}

	uploadImageReq := &entity.UploadImageRequest{Userid: userid, Filename: fileHeader.Filename, File: file}

	resp, err := h.uc.UploadImage(ctx, uploadImageReq)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, NewResponsePayload("error", err.Error()))
	}

	return c.JSON(http.StatusOK, NewResponsePayload("success", resp))

}

func (h *Handler) ChangePassword(c echo.Context) error {
	// userid from jwt
	userid := c.Get("userid").(string)

	changePasswordReq := new(entity.ChangePasswordRequest)
	if err := c.Bind(changePasswordReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, NewResponsePayload("error", err.Error()))
	}

	// validate new password
	if err := c.Validate(changePasswordReq); err != nil {
		return err
	}

	changePasswordReq.Userid = userid

	resp, err := h.uc.ChangePassword(ctx, changePasswordReq)

	fmt.Println("error", err)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, NewResponsePayload("error", err.Error()))
	}

	return c.JSON(http.StatusOK, NewResponsePayload("success", resp))
}
