package http

import (
	"encoding/json"
	"net/http"
	"weather-app/internal/core/service"
)

type WeatherHandler struct {
	service *service.WeatherService
}

func NewWeatherHandler(s *service.WeatherService) *WeatherHandler {
	return &WeatherHandler{
		service: s,
	}
}

func (h *WeatherHandler) GetWeather(w http.ResponseWriter, r *http.Request) {
	city := r.URL.Query().Get("city")
	if city == "" {
		http.Error(w, "City parameter is required", http.StatusBadRequest)
		return
	}

	weather, err := h.service.GetWeather(city)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(weather)
}
