package controller

import (
	"context"
	"io"
)

type ImageBucket interface {
	SaveImage(ctx context.Context, name string, r io.Reader) (string, error)
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
