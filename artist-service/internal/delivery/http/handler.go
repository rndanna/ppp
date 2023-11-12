package http

import (
	"artist-service/internal/domain"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ArtistHandler struct {
	AUseCase domain.ArtistUseCase
}

func New(e *echo.Echo, us domain.ArtistUseCase) {
	handler := &ArtistHandler{
		AUseCase: us,
	}

	e.GET("artist/:id", handler.GetArtist)
	e.GET("artist/name", handler.GetArtistByName)
}

func (h *ArtistHandler) GetArtist(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Errorf("err GetArtist: %w", err))
	}

	if id <= 0 {
		return c.JSON(http.StatusInternalServerError, fmt.Errorf("err id<0: %w", err))
	}

	artist, err := h.AUseCase.GetArtist(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Errorf("err GetArtist: %w", err))
	}

	return c.JSON(http.StatusOK, &artist)
}

func (h *ArtistHandler) GetArtistByName(c echo.Context) error {
	name := c.QueryParam("name")
	if name != "" {
		return c.JSON(http.StatusInternalServerError, fmt.Errorf("err name is empty"))
	}

	artist, err := h.AUseCase.GetArtistByName(name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Errorf("err GetArtistByName: %w", err))
	}

	return c.JSON(http.StatusOK, &artist)
}
