package main

import "github.com/prometheus/client_golang/prometheus"

var abandonedQueries = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Namespace: "trino",
	Name:      "queries_abandoned_count_total",
}, []string{"cluster"})

var cancelledQueries = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Namespace: "trino",
	Name:      "queries_cancelled_count_total",
}, []string{"cluster"})

var completedQueries = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Namespace: "trino",
	Name:      "queries_completed_count_total",
}, []string{"cluster"})

var failedQueries = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Namespace: "trino",
	Name:      "queries_failed_count_total",
}, []string{"cluster"})

var queuedQueries = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Namespace: "trino",
	Name:      "queries_queued_count_total",
}, []string{"cluster"})

var runningQueries = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Namespace: "trino",
	Name:      "queries_running_count_total",
}, []string{"cluster"})

var inputBytes = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Namespace: "trino",
	Name:      "input_bytes_total",
}, []string{"cluster"})

var inputRows = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Namespace: "trino",
	Name:      "input_rows_total",
}, []string{"cluster"})

var cpuTime = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Namespace: "trino",
	Name:      "cpu_time_seconds_total",
}, []string{"cluster"})

var execTimeAvg = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Namespace: "trino",
	Name:      "execution_time_milliseconds_average",
}, []string{"cluster"})

var execTimeP95 = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Namespace: "trino",
	Name:      "execution_time_milliseconds_p95",
}, []string{"cluster"})

var queuedTimeAvg = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Namespace: "trino",
	Name:      "queued_time_milliseconds_average",
}, []string{"cluster"})

var queuedTimeP95 = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Namespace: "trino",
	Name:      "queued_time_milliseconds_p95",
}, []string{"cluster"})

var externalFailures = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Namespace: "trino",
	Name:      "failures_external_count_total",
}, []string{"cluster"})

var internalFailures = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Namespace: "trino",
	Name:      "failures_internal_count_total",
}, []string{"cluster"})

var userFailures = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Namespace: "trino",
	Name:      "failures_user_count_total",
}, []string{"cluster"})

var resourceFailures = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Namespace: "trino",
	Name:      "failures_insufficient_resources_count_total",
}, []string{"cluster"})

var Gauges = []*prometheus.GaugeVec{
	abandonedQueries,
	cancelledQueries,
	completedQueries,
	cpuTime,
	execTimeAvg,
	execTimeP95,
	externalFailures,
	failedQueries,
	inputBytes,
	inputRows,
	internalFailures,
	queuedQueries,
	queuedTimeAvg,
	queuedTimeP95,
	resourceFailures,
	runningQueries,
	userFailures,
}

func UpdateGauges(m JmxMetrics) {
	abandonedQueries.With(prometheus.Labels{"cluster": Cluster}).Set(m.AbandonedQueriesTotalCount)
	abandonedQueries.With(prometheus.Labels{"cluster": Cluster}).Set(m.AbandonedQueriesTotalCount)
	cancelledQueries.With(prometheus.Labels{"cluster": Cluster}).Set(m.CancelledQueriesTotalCount)
	completedQueries.With(prometheus.Labels{"cluster": Cluster}).Set(m.CompletedQueriesTotalCount)
	cpuTime.With(prometheus.Labels{"cluster": Cluster}).Set(m.ConsumedCpuTimeSecsTotalCount)
	execTimeAvg.With(prometheus.Labels{"cluster": Cluster}).Set(m.ExecutionTimeAllTimeAvg)
	execTimeP95.With(prometheus.Labels{"cluster": Cluster}).Set(m.ExecutionTimeAllTimeP95)
	externalFailures.With(prometheus.Labels{"cluster": Cluster}).Set(m.ExternalFailuresTotalCount)
	failedQueries.With(prometheus.Labels{"cluster": Cluster}).Set(m.FailedQueriesTotalCount)
	inputBytes.With(prometheus.Labels{"cluster": Cluster}).Set(m.ConsumedInputBytesTotalCount)
	inputRows.With(prometheus.Labels{"cluster": Cluster}).Set(m.ConsumedInputRowsTotalCount)
	internalFailures.With(prometheus.Labels{"cluster": Cluster}).Set(m.InternalFailuresTotalCount)
	queuedQueries.With(prometheus.Labels{"cluster": Cluster}).Set(m.QueuedQueries)
	queuedTimeAvg.With(prometheus.Labels{"cluster": Cluster}).Set(m.QueuedTimeAllTimeAvg)
	queuedTimeP95.With(prometheus.Labels{"cluster": Cluster}).Set(m.QueuedTimeAllTimeP95)
	resourceFailures.With(prometheus.Labels{"cluster": Cluster}).Set(m.InsufficientResourcesFailuresTotalCount)
	runningQueries.With(prometheus.Labels{"cluster": Cluster}).Set(m.RunningQueries)
	userFailures.With(prometheus.Labels{"cluster": Cluster}).Set(m.UserErrorFailuresTotalCount)
}
