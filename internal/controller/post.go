package controller

import (
	"fmt"
	"gosocial/internal/entity"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) CreatePost(c echo.Context) error {
	userid := c.Get("userid").(string)

	// Parse the form data
	// Retrieve the text input field "caption" from the form data
	caption := c.FormValue("caption")
	// image
	imageHeader, imageErr := c.FormFile("image")
	if imageErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, NewResponsePayload("error", "An image attachment is required"))
	}
	contentType := imageHeader.Header.Get("Content-Type")
	// Ensure that the uploaded file is of type
	if contentType != "image/png" &&
		contentType != "image/jpeg" &&
		contentType != "image/gif" {
		return echo.NewHTTPError(http.StatusBadRequest, NewResponsePayload("error", "Invalid file type. Please upload an image."))
	}

	file, fileErr := imageHeader.Open()
	defer file.Close()
	if fileErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, NewResponsePayload("error", fileErr.Error()))
	}

	createPostRequest := &entity.CreatePostRequest{
		Post: entity.Post{
			Userid:  userid,
			Caption: caption,
		},
		Filename: imageHeader.Filename,
		File:     file,
	}

	fmt.Println("CreatePostRequest:", *createPostRequest)

	resp, err := h.uc.CreatePost(ctx, createPostRequest)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, NewResponsePayload("error", err.Error()))
	}

	return c.JSON(http.StatusOK, NewResponsePayload("success", resp))
}

func (h *Handler) GetPosts(c echo.Context) error {

	var req entity.PostsRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, NewResponsePayload("error", err.Error()))
	}

	posts, err := h.uc.GetPosts(ctx, &req)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, NewResponsePayload("error", err.Error()))
	}

	return c.JSON(http.StatusOK, NewResponsePayload("success", posts))
}

func (h *Handler) GetPostsByUser(c echo.Context) error {

	var req entity.PostsByUserRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, NewResponsePayload("error", err.Error()))
	}

	posts, err := h.uc.GetPostsByUser(ctx, &req)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, NewResponsePayload("error", err.Error()))
	}

	return c.JSON(http.StatusOK, NewResponsePayload("success", posts))
}

func (h *Handler) DeletePost(c echo.Context) error {

	var req entity.DeletePostRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, NewResponsePayload("error", err.Error()))
	}

	resp, err := h.uc.DeletePost(ctx, &req)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, NewResponsePayload("error", err.Error()))
	}

	return c.JSON(http.StatusOK, NewResponsePayload("success", resp))
}

func (h *Handler) LikePost(c echo.Context) error {

	var req entity.LikePostRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, NewResponsePayload("error", err.Error()))
	}

	resp, err := h.uc.LikePost(ctx, &req)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, NewResponsePayload("error", err.Error()))
	}

	return c.JSON(http.StatusOK, NewResponsePayload("success", resp))
}
