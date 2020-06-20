package routes_test

import (
	"encoding/json"
	"net/http"
	"outecomex/internal/weather"
	"testing"

	"github.com/stretchr/testify/require"
	"gopkg.in/go-playground/assert.v1"
)

func TestHandleWeather(t *testing.T) {

	t.Run("without city", withoutCity)
	t.Run("with invalid city", withInvalidCity)
	t.Run("with valid city", withValidCity)

}

func withoutCity(t *testing.T) {

	// Arrange

	// Act
	resp, err := http.Get(server + "/weather")

	// Assert
	require.Nil(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func withInvalidCity(t *testing.T) {

	// Arrange

	// Act
	resp, err := http.Get(server + "/weather/covid")

	// Assert
	require.Nil(t, err)
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func withValidCity(t *testing.T) {

	// Arrange

	// Act
	resp, err := http.Get(server + "/weather/sydney")

	// Assert
	require.Nil(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)

	decoder := json.NewDecoder(resp.Body)
	defer resp.Body.Close()

	var result []weather.WeatherItem
	err = decoder.Decode(&result)
	require.Nil(t, err, "deserializing observation")
	assert.Equal(t, 48, len(result))
}
