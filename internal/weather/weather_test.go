package weather_test

import (
	"net/http"
	"os"
	"outecomex/internal/weather"
	"testing"

	"context"

	"github.com/stretchr/testify/require"
)

// This is not a great test. When testing shared out of process dependencies mocks should be used, but I do not have time to write out a mock for the API.

func TestFetchWeather(t *testing.T) {

	// Arrange
	ctx := context.Background()
	apiKey, found := os.LookupEnv("WEATHER_API_KEY")
	require.True(t, found, "API_KEY env variable not found")
	client := weather.New(apiKey, http.DefaultClient)

	// Act
	_, err := client.FetchWeather(ctx, weather.Location{Lon: 151.209900, Lat: -33.865143})

	// Assert
	require.Nil(t, err)
}

func TestFetchWeatherWithInvalidCredentials(t *testing.T) {

	// Arrange
	ctx := context.Background()
	client := weather.New("123", http.DefaultClient)

	// Act
	_, err := client.FetchWeather(ctx, weather.Location{Lon: 151.209900, Lat: -33.865143})

	// Assert
	require.Error(t, err)
}
