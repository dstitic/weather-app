package ports

import "weather-app/internal/core/domain"

type WeatherAPI interface {
	FetchWeather(city string) (*domain.Weather, error)
}

type WeatherRepository interface {
	SaveWeather(*domain.Weather) error
	GetWeather(city string) (*domain.Weather, error)
}
