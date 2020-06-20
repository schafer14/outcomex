package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"outecomex/internal/weather"
	"strings"

	"github.com/pkg/errors"
)

// Handler is structure capable of serving http for the full website.
// It would be better to not expose the weather API and instantiate it
// with a new function, but time.
type Handler struct {
	Wa weather.WeatherAPI
}

// ServeHTTP handles all incoming requets. This does _very_ basic routing.
// With a larger app it would be worth structing this out a little bit better and
// using a better router such as Gorilla, Chi or any of the other billion routers.
//
// The router routes between the go api under the /weather endpoint and if not that
// delegates to the react app. This means the front-end is responsible for 404's and what not.
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// Break the url into parts
	parts := strings.Split(r.URL.Path, "/")
	switch parts[1] {
	case "weather":
		h.handleWeather(w, r)
		return
	default:
		http.StripPrefix("/", http.FileServer(http.Dir("./front-end/build"))).ServeHTTP(w, r)
		return
	}
}

// In a reall application it would probably be better to put this into a configuration file
// like json to make it more human read/editable.
var cityMap = map[string]weather.Location{
	"sydney":    {Lon: 151.209900, Lat: -33.865143},
	"melbourne": {Lon: 144.96332, Lat: -37.814},
	"adelaide":  {Lon: 138.59863, Lat: -34.92866},
}

// handleWeather responds to a weather API request.
//
// This function is particularly verbose because it is doing a lot of what would normally be
// boilerplate code. Since there is only one api endpoint it does not make sense to abstract
// this boilerplate code at this point.
//
// The boilerplate that would be abstracted is
// - Handling a request url
// - Converting a response to JSON
// - Writing the response to the ResponseWriter
func (h *Handler) handleWeather(w http.ResponseWriter, r *http.Request) {

	// We expect the route to have the shape of "/weather/sydney"
	// so there should be three parts separated by '/'
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {

		// In a real application I would probably serve the list of cities from this endpoint.
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("City not found"))
		return
	}

	location, found := cityMap[parts[2]]
	if !found {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Unknown city"))
		return
	}

	res, err := h.Wa.FetchWeather(r.Context(), location)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not fetch weather"))
		return
	}

	jsonData, err := json.Marshal(res)
	if err != nil {
		log.Println(errors.Wrap(err, "parsing json data"))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not fetch weather"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(jsonData); err != nil {
		log.Println(errors.Wrap(err, "sending response"))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not fetch weather"))
	}
	return
}
