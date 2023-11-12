package http

import (
	"album-service/internal/domain"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AlbumHandler struct {
	AUseCase domain.AlbumUseCase
}

func New(e *echo.Echo, us domain.AlbumUseCase) {
	handler := &AlbumHandler{
		AUseCase: us,
	}

	e.GET("artist/name", handler.GetAlbumByTitle)
}

func (h *AlbumHandler) GetAlbumByTitle(c echo.Context) error {
	title := c.QueryParam("title")
	if title != "" {
		return c.JSON(http.StatusInternalServerError, fmt.Errorf("err title is empty"))
	}

	album, err := h.AUseCase.GetAlbumByTitle(title)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Errorf("err GetAlbumByTitle: %w", err))
	}

	return c.JSON(http.StatusOK, &album)
}
