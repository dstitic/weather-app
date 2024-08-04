package service

import (
	"weather-app/internal/core/domain"
	"weather-app/internal/core/ports"
)

type WeatherService struct {
	api  ports.WeatherAPI
	repo ports.WeatherRepository
}

func NewWeatherService(api ports.WeatherAPI, repo ports.WeatherRepository) *WeatherService {
	return &WeatherService{api: api, repo: repo}
}

func (s *WeatherService) GetWeather(city string) (*domain.Weather, error) {
	weather, err := s.repo.GetWeather(city)
	if err == nil {
err = s.repo.SaveWeather(weather)
	if err != nil {
		return nil, err
	}
		return weather, nil
	}


	weather, err = s.api.FetchWeather(city)
	if err != nil {
		return weather, err
	}

	

	return weather, nil
}
