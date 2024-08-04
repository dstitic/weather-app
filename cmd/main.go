package main

import (
	"log"
	bhttp "net/http"
	"weather-app/internal/adapters/http"
	"weather-app/internal/adapters/repository"
	"weather-app/internal/adapters/weatherapi"
	"weather-app/internal/core/service"
)

func main() {
	weatherAPI := weatherapi.NewOpenWeatherMapAdapter("your-api-code")
	/*weatherRepo, err := repository.NewPostgresAdapter("foo")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}*/
	weatherRepo := repository.NewMockWeatherRepository("dummy-connection-string")

	weatherService := service.NewWeatherService(weatherAPI, weatherRepo)

	weatherHandler := http.NewWeatherHandler(weatherService)

	bhttp.HandleFunc("/weather", weatherHandler.GetWeather)

	log.Println("Server starting on :8080")
	log.Fatal(bhttp.ListenAndServe(":8080", nil))

}
