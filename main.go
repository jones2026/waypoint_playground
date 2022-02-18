package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	metrics "github.com/slok/go-http-metrics/metrics/prometheus"
	slokmdlw "github.com/slok/go-http-metrics/middleware"
	"github.com/slok/go-http-metrics/middleware/std"
)

func main() {
	log.Println("Starting app...")
	r := chi.NewRouter()

	prometheusMiddleware := slokmdlw.New(slokmdlw.Config{
		Recorder: metrics.NewRecorder(metrics.Config{}),
	})

	r.Use(std.HandlerProvider("", prometheusMiddleware))
	r.Use(middleware.Heartbeat("/ping"))
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", Hello)
	r.Handle("/metrics", promhttp.Handler())

	port := ":8080"
	log.Println("Listening on port:", port)
	log.Fatalln(http.ListenAndServe(port, r))
}

func Hello(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello World!"))
	if err != nil {
		log.Fatalf(err.Error())
	}
}
