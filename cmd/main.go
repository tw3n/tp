package main

import (
	"log"
	"net/http"

	"github.com/alexflint/go-arg"
	owm "github.com/briandowns/openweathermap"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type Config struct {
	APIKey  string `arg:"env"`
	Country string `arg:"env"`
}

func main() {
	var cfg Config
	arg.MustParse(&cfg)

	cwd, err := owm.NewCurrent("C", cfg.Country, cfg.APIKey)
	if err != nil {
		log.Fatalln(err)
	}

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Get("/weather", func(w http.ResponseWriter, r *http.Request) {
		if err := cwd.CurrentByZipcode(r.URL.Query().Get("zip"), cfg.Country); err != nil {
			render.JSON(w, r, errorResponse{
				Cause: err.Error(),
				Code:  "not_found",
			})
			render.Status(r, 404)
			return
		}

		render.JSON(w, r, &weatherResponse{
			Current: weatherCurrentResponse{
				FeelsLike: cwd.Main.FeelsLike,
				Temp:      cwd.Main.Temp,
			},
			Latitude:  cwd.GeoPos.Latitude,
			Longitude: cwd.GeoPos.Longitude,
			Name:      cwd.Name,
		})
	})

	if err := http.ListenAndServe(":4000", r); err != nil {
		log.Fatalln(err)
	}
}

type errorResponse struct {
	Cause string `json:"cause"`
	Code  string `json:"code"`
}

type weatherCurrentResponse struct {
	FeelsLike float64 `json:"feelsLike"`
	Temp      float64 `json:"temp"`
}

type weatherResponse struct {
	Current   weatherCurrentResponse `json:"current"`
	Latitude  float64                `json:"lat"`
	Longitude float64                `json:"lon"`
	Name      string                 `json:"name"`
}
