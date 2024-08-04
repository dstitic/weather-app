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
		baseURL: "http://api.openweathermap.org/data/3.0/onecall",
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

	current, ok := result["current"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response format")
	}

	weatherArray, ok := current["weather"].([]interface{})
	if !ok || len(weatherArray) == 0 {
		return nil, fmt.Errorf("weather information not found")
	}

	weatherInfo, ok := weatherArray[0].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid weather information format")
	}

	weather := &domain.Weather{
		City:        city,
		Temperature: current["temp"].(float64),
		Description: weatherInfo["description"].(string),
	}

	return weather, nil
}

func getLatLonForCity(city string) (lattitude, longitude string) {
	// TODO: This test implementation always returns the values for the City of Karlsruhe, Germany
	return "49.006889", "8.403653"
}
