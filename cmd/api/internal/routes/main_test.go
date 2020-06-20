package routes_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"outecomex/cmd/api/internal/routes"
	"outecomex/internal/weather"
	"testing"
)

var server string

// TestMain runs a database for this package.
func TestMain(m *testing.M) {

	apiKey, found := os.LookupEnv("WEATHER_API_KEY")
	if !found {
		fmt.Println("Could not find weather key")
		os.Exit(1)
	}

	weatherAPI := weather.New(apiKey, http.DefaultClient)
	router := &routes.Handler{weatherAPI}
	s := httptest.NewServer(router)
	server = s.URL

	result := m.Run()

	// Cannot use defer with os.Exit
	s.Close()
	os.Exit(result)
}
