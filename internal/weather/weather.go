package weather

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

// WeatherAPI defines operations you may perform to get weather.
type WeatherAPI interface {
	FetchWeather(context.Context, Location) ([]WeatherItem, error)
}

// weatherAPI is the internal structure. I explicity pass the http client
// to facilitate moving the tests from e2e to proper unit tests when I have more time.
type weatherAPI struct {
	apiKey string
	client *http.Client
}

// New creates a new api given the API key.
func New(apiKey string, client *http.Client) WeatherAPI {
	return &weatherAPI{apiKey, client}
}

// FetchWeather returns the weather for a given city.
func (wa *weatherAPI) FetchWeather(ctx context.Context, loc Location) ([]WeatherItem, error) {

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/onecall?lon=%v&lat=%v&appid=%s", loc.Lon, loc.Lat, wa.apiKey)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return []WeatherItem{}, errors.Wrap(err, "building request")
	}

	res, err := wa.client.Do(req)
	if err != nil {
		return []WeatherItem{}, errors.Wrap(err, "fetching weather")
	}

	decoder := json.NewDecoder(res.Body)
	var weatherResult weather
	if err := decoder.Decode(&weatherResult); err != nil {
		return []WeatherItem{}, errors.Wrap(err, "parsing response")
	}

	if weatherResult.Code >= 400 {
		err := fmt.Errorf("openweathermap responded with a %v and the message %q", weatherResult.Code, weatherResult.Message)
		return []WeatherItem{}, errors.Wrap(err, "fetching weather")
	}

	return parseWeather(weatherResult), nil
}

// convert from weatherAPI struct to data transfer object specific to our app.
func parseWeather(weather weather) []WeatherItem {

	var list []WeatherItem
	for _, moment := range weather.Hourly {
		var descriptions []string

		for _, w := range moment.Weather {
			descriptions = append(descriptions, w.Description)
		}

		list = append(list, WeatherItem{
			FeelsLike:   moment.FeelsLike,
			Temp:        moment.Temp,
			WindSpeed:   moment.WindSpeed,
			WindDir:     moment.WindDir,
			Description: descriptions,
		})
	}

	return list
}
