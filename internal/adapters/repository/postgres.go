package repository

import (
	"database/sql"
	"fmt"
	"sync"
	"weather-app/internal/core/domain"
	"weather-app/internal/core/ports"
)

type PostgresAdapter struct {
	db *sql.DB
}

func NewPostgresAdapter(connString string) (*PostgresAdapter, error) {
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}
	return &PostgresAdapter{
		db: db,
	}, nil

}

func (a *PostgresAdapter) SaveWeather(w *domain.Weather) error {
	_, err := a.db.Exec("INSERT INTO weather (city, temperature, description) VALUES ($1, $2, $3)",
		w.City, w.Temperature, w.Description)

	return err
}

func (a *PostgresAdapter) GetWeather(city string) (*domain.Weather, error) {
	var w domain.Weather
	err := a.db.QueryRow("SELECT city, temperature, description FROM weather WHERE city = $1", city).
		Scan(&w.City, &w.Temperature, &w.Description)
	if err != nil {
		return nil, err
	}
	return &w, err
}

var _ ports.WeatherRepository = (*MockWeatherRepository)(nil)

// MockWeatherRepository is an in-memory implementation of WeatherRepository
type MockWeatherRepository struct {
	mu      sync.RWMutex
	weather map[string]*domain.Weather
}

// NewMockWeatherRepository creates a new instance of MockWeatherRepository
func NewMockWeatherRepository(dummyConnString string) *MockWeatherRepository {
	return &MockWeatherRepository{
		weather: make(map[string]*domain.Weather),
	}
}

// SaveWeather saves weather data to the in-memory store
func (m *MockWeatherRepository) SaveWeather(w *domain.Weather) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.weather[w.City] = w
	return nil
}

// GetWeather retrieves weather data from the in-memory store
func (m *MockWeatherRepository) GetWeather(city string) (*domain.Weather, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	w, ok := m.weather[city]
	if !ok {
		return nil, fmt.Errorf("weather data not found for city: %s", city)
	}
	return w, nil
}
