package weatherapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"weather-app/internal/core/domain"
)

type OpenWeatherMapAdapter struct {
	apiKey  string
	baseURL string
}

func NewOpenWeatherMapAdapter(apiKey string) *OpenWeatherMapAdapter {
	return &OpenWeatherMapAdapter{
		apiKey:  apiKey,
		baseURL: "https://api.openweathermap.org/data/2.5/weather",
	}
}

func (a *OpenWeatherMapAdapter) FetchWeather(city string) (*domain.Weather, error) {
	lattitude, longitude := getLatLonForCity(city)

	url := fmt.Sprintf("%s?lat=%s&lon=%s&appid=%s&units=metric", a.baseURL, lattitude, longitude, a.apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	weather := &domain.Weather{
		City:        city,
		Temperature: result["main"].(map[string]interface{})["temp"].(float64),
		Description: result["weather"].([]interface{})[0].(map[string]interface{})["description"].(string),
	}

	return weather, nil
}

func getLatLonForCity(city string) (lattitude, longitude string) {
	// TODO: This test implementation always returns the values for the City of Karlsruhe, Germany
	return "49.006889", "8.403653"
}
