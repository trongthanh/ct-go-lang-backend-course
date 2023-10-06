package controller

import (
	"gosocial/internal/entity"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetProfiles(c echo.Context) error {

	var req entity.ProfilesRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, NewResponsePayload("error", err.Error()))
	}

	resp, err := h.uc.GetProfiles(ctx, &req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, NewResponsePayload("error", err.Error()))
	}

	return c.JSON(http.StatusOK, NewResponsePayload("success", resp))
}
