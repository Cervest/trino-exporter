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

func metricsHandler(w http.ResponseWriter, r *http.Request) {
	jmxMetrics, err := ReadJmxMetrics()
	if err == nil {
		UpdateGauges(*jmxMetrics)
	}
	promhttp.Handler().ServeHTTP(w, r)
}

var Host string
var Cluster string

func main() {
	Cluster = os.Args[1]
	Host = os.Args[2]

	for _, gauge := range Gauges {
		prometheus.MustRegister(gauge)
	}

	server := &http.Server{
		Addr: ":8567",
	}
	http.HandleFunc("/metrics", metricsHandler)
	http.HandleFunc("/status", statusHandler)
	log.Fatal(server.ListenAndServe())
}
