package main

import (
	"fmt"
	"github.com/rcrowley/go-metrics"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	hitCounter := metrics.NewMeter()
	metrics.Register("hits", hitCounter)

	adminHandler := http.NewServeMux()
	adminHandler.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		hitCounter.Mark(1)
	})
	adminHandler.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		hitCounter.Mark(1)
		fmt.Fprint(w, ",\"metrics\":")
		metrics.WriteJSONOnce(metrics.DefaultRegistry, w)
		fmt.Fprint(w, "}")
	})
	admin := &http.Server{
		Addr:    ":" + port,
		Handler: adminHandler,
	}
	log.Fatal(admin.ListenAndServe())
}
