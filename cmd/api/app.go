package main

import (
	"expvar"
	"fmt"
	"log"
	"net/http"
	"os"
	"outecomex/cmd/api/internal/routes"
	"outecomex/internal/weather"

	"github.com/ardanlabs/conf"
	"github.com/pkg/errors"
)

var build = "develop"
var version = "v0.0.1"

// main is just a pass through function for run. I do this because
// main is not allowed to have the type signature that I want it to.
func main() {
	if err := run(); err != nil {
		fmt.Println("error :", err)
		os.Exit(1)
	}
}

// run is like main except it returns an error.
func run() error {

	// =============================================== //
	// Read Configuration
	// =============================================== //
	var cfg struct {
		APIHost string `conf:"default:0.0.0.0:3000"`
		APIKey  string `conf:"noprint"`
	}

	if err := conf.Parse(os.Args[1:], "Weather", &cfg); err != nil {
		fmt.Println(err)
		if err == conf.ErrHelpWanted {
			usage, err := conf.Usage("Weather", &cfg)
			if err != nil {
				return errors.Wrap(err, "generating config usage")
			}
			fmt.Println(usage)
			return nil
		}
		return errors.Wrap(err, "parsing config")
	}

	// =============================================== //
	// Report Build Parameters
	// =============================================== //
	expvar.NewString("build").Set(build)
	log.Printf("main : Started : Application initializing : version %q %q", build, version)
	defer log.Println("main : Completed")

	out, err := conf.String(&cfg)
	if err != nil {
		return errors.Wrap(err, "generating config for output")
	}
	log.Printf("main : Config :\n%v\n", out)

	// =============================================== //
	// Starting API
	// =============================================== //
	log.Println("main : Started : Initializing API support")

	if cfg.APIKey == "" {
		return errors.New("No API Key Found")
	}

	weatherAPI := weather.New(cfg.APIKey, http.DefaultClient)

	server := &http.Server{
		Addr:    cfg.APIHost,
		Handler: &routes.Handler{Wa: weatherAPI},
	}

	return server.ListenAndServe()
}
