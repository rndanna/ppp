package http

import (
	"fmt"
	"net/http"
	"strconv"
	"track-service/internal/domain"

	"github.com/labstack/echo/v4"
)

type TrackHandler struct {
	TUseCase domain.TrackUseCase
}

func New(e *echo.Echo, us domain.TrackUseCase) {
	handler := &TrackHandler{
		TUseCase: us,
	}

	e.GET("track/:id", handler.GetTrack)
	e.GET("track/tag/:id", handler.GetTrackByTag)
	e.GET("track/artist/:id", handler.GetTrackByArtist)
	e.GET("track/chart", handler.GetTrackChart)
}

func (h *TrackHandler) GetTrack(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Errorf("err GetTrack: %w", err))
	}

	if id <= 0 {
		return c.JSON(http.StatusInternalServerError, fmt.Errorf("err id<0: %w", err))
	}

	track, err := h.TUseCase.GetTrack(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Errorf("err GetTrack: %w", err))
	}

	return c.JSON(http.StatusOK, &track)
}

func (h *TrackHandler) GetTrackByTag(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, fmt.Errorf("err GetTrackByTag: %w", err))
	}

	if id <= 0 {
		fmt.Println(err)

		return c.JSON(http.StatusInternalServerError, fmt.Errorf("err id<0: %w", err))
	}

	track, err := h.TUseCase.GetTrackByTag(id)
	if err != nil {
		fmt.Println(err)

		return c.JSON(http.StatusInternalServerError, fmt.Errorf("err GetTrackByTag: %w", err))
	}

	return c.JSON(http.StatusOK, &track)
}

func (h *TrackHandler) GetTrackByArtist(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, fmt.Errorf("err GetTrackByArtist: %w", err))
	}

	if id <= 0 {
		fmt.Println(err)

		return c.JSON(http.StatusInternalServerError, fmt.Errorf("err id<0: %w", err))
	}

	track, err := h.TUseCase.GetTrackByArtist(id)
	if err != nil {
		fmt.Println(err)

		return c.JSON(http.StatusInternalServerError, fmt.Errorf("err GetTrackByArtist: %w", err))
	}

	return c.JSON(http.StatusOK, &track)
}

func (h *TrackHandler) GetTrackChart(c echo.Context) error {
	var (
		tracks []*domain.Track
		err    error
	)

	switch c.QueryParam("chart") {
	case "playcount":
		tracks, err = h.TUseCase.GetTopTrackByPlayCount()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, fmt.Errorf("err GetTrackByArtist: %w", err))
		}
	case "listeners":
		tracks, err = h.TUseCase.GetTopTrackByListeners()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, fmt.Errorf("err GetTrackByArtist: %w", err))
		}
	default:
		return c.JSON(http.StatusInternalServerError, fmt.Errorf("err GetTrackByArtist: %w", err))
	}

	return c.JSON(http.StatusOK, &tracks)
}
