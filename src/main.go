package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func statusHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "I'm well, thank you.")
}

func metricsHandler(w http.ResponseWriter, r *http.Request) {
	jmxMetrics, err := ReadJmxMetrics()
	if err == nil {
		UpdateGauges(*jmxMetrics)
	} else {
		log.Error().Msg(err.Error())
	}
	apiQueries, err := FetchQueryInfo()
	if err == nil {
		for _, q := range apiQueries {
			log.Log().
				Str("catalog", q.Catalog).
				Time("time", q.CreateTime).
				Int64("duration_ms", q.DurationMs).
				Str("id", q.QueryId).
				Int("size_bytes", q.QuerySizeBytes).
				Str("type", q.QueryType).
				Str("state", q.State).
				Str("user", q.User).
				Str("user_agent", q.UserAgent).
				Msg("")
		}
	} else {
		log.Error().Msg(err.Error())
	}
	promhttp.Handler().ServeHTTP(w, r)
}

var Host string
var Cluster string

func main() {
	Cluster = os.Args[1]
	Host = os.Args[2]

	zerolog.TimestampFieldName = "exporter_log_time"

	for _, gauge := range Gauges {
		prometheus.MustRegister(gauge)
	}

	server := &http.Server{
		Addr: ":5885",
	}
	http.HandleFunc("/metrics", metricsHandler)
	http.HandleFunc("/status", statusHandler)
	server.ListenAndServe()
}
