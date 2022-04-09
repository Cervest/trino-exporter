package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func statusHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "I'm well, thank you.")
}

var counter = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "counter",
	Help: "Help!",
})

func metricsHandler(w http.ResponseWriter, r *http.Request) {
	counter.Inc()
	promhttp.Handler().ServeHTTP(w, r)
}

func main() {
	prometheus.MustRegister(counter)

	server := &http.Server{
		Addr: ":8567",
	}
	http.HandleFunc("/metrics", metricsHandler)
	http.HandleFunc("/status", statusHandler)
	log.Fatal(server.ListenAndServe())
}
