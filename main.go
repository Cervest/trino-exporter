package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

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
	fmt.Println(ReadJmxMetrics())
	promhttp.Handler().ServeHTTP(w, r)
}

var host string

func main() {
	host = os.Args[1]
	prometheus.MustRegister(counter)

	server := &http.Server{
		Addr: ":8567",
	}
	http.HandleFunc("/metrics", metricsHandler)
	http.HandleFunc("/status", statusHandler)
	log.Fatal(server.ListenAndServe())
}
