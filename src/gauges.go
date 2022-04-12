package main

import "github.com/prometheus/client_golang/prometheus"

var metricClusterKey = "trino-cluster"

var abandonedQueries = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Namespace: "trino",
	Name:      "queries_abandoned_count_total",
}, []string{metricClusterKey})

var cancelledQueries = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Namespace: "trino",
	Name:      "queries_cancelled_count_total",
}, []string{metricClusterKey})

var completedQueries = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Namespace: "trino",
	Name:      "queries_completed_count_total",
}, []string{metricClusterKey})

var failedQueries = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Namespace: "trino",
	Name:      "queries_failed_count_total",
}, []string{metricClusterKey})

var queuedQueries = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Namespace: "trino",
	Name:      "queries_queued_count_total",
}, []string{metricClusterKey})

var runningQueries = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Namespace: "trino",
	Name:      "queries_running_count_total",
}, []string{metricClusterKey})

var inputBytes = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Namespace: "trino",
	Name:      "input_bytes_total",
}, []string{metricClusterKey})

var inputRows = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Namespace: "trino",
	Name:      "input_rows_total",
}, []string{metricClusterKey})

var cpuTime = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Namespace: "trino",
	Name:      "cpu_time_seconds_total",
}, []string{metricClusterKey})

var execTimeAvg = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Namespace: "trino",
	Name:      "execution_time_milliseconds_average",
}, []string{metricClusterKey})

var execTimeP95 = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Namespace: "trino",
	Name:      "execution_time_milliseconds_p95",
}, []string{metricClusterKey})

var queuedTimeAvg = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Namespace: "trino",
	Name:      "queued_time_milliseconds_average",
}, []string{metricClusterKey})

var queuedTimeP95 = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Namespace: "trino",
	Name:      "queued_time_milliseconds_p95",
}, []string{metricClusterKey})

var externalFailures = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Namespace: "trino",
	Name:      "failures_external_count_total",
}, []string{metricClusterKey})

var internalFailures = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Namespace: "trino",
	Name:      "failures_internal_count_total",
}, []string{metricClusterKey})

var userFailures = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Namespace: "trino",
	Name:      "failures_user_count_total",
}, []string{metricClusterKey})

var resourceFailures = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Namespace: "trino",
	Name:      "failures_insufficient_resources_count_total",
}, []string{metricClusterKey})

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
	abandonedQueries.With(prometheus.Labels{metricClusterKey: Cluster}).Set(m.AbandonedQueriesTotalCount)
	abandonedQueries.With(prometheus.Labels{metricClusterKey: Cluster}).Set(m.AbandonedQueriesTotalCount)
	cancelledQueries.With(prometheus.Labels{metricClusterKey: Cluster}).Set(m.CancelledQueriesTotalCount)
	completedQueries.With(prometheus.Labels{metricClusterKey: Cluster}).Set(m.CompletedQueriesTotalCount)
	cpuTime.With(prometheus.Labels{metricClusterKey: Cluster}).Set(m.ConsumedCpuTimeSecsTotalCount)
	execTimeAvg.With(prometheus.Labels{metricClusterKey: Cluster}).Set(m.ExecutionTimeAllTimeAvg)
	execTimeP95.With(prometheus.Labels{metricClusterKey: Cluster}).Set(m.ExecutionTimeAllTimeP95)
	externalFailures.With(prometheus.Labels{metricClusterKey: Cluster}).Set(m.ExternalFailuresTotalCount)
	failedQueries.With(prometheus.Labels{metricClusterKey: Cluster}).Set(m.FailedQueriesTotalCount)
	inputBytes.With(prometheus.Labels{metricClusterKey: Cluster}).Set(m.ConsumedInputBytesTotalCount)
	inputRows.With(prometheus.Labels{metricClusterKey: Cluster}).Set(m.ConsumedInputRowsTotalCount)
	internalFailures.With(prometheus.Labels{metricClusterKey: Cluster}).Set(m.InternalFailuresTotalCount)
	queuedQueries.With(prometheus.Labels{metricClusterKey: Cluster}).Set(m.QueuedQueries)
	queuedTimeAvg.With(prometheus.Labels{metricClusterKey: Cluster}).Set(m.QueuedTimeAllTimeAvg)
	queuedTimeP95.With(prometheus.Labels{metricClusterKey: Cluster}).Set(m.QueuedTimeAllTimeP95)
	resourceFailures.With(prometheus.Labels{metricClusterKey: Cluster}).Set(m.InsufficientResourcesFailuresTotalCount)
	runningQueries.With(prometheus.Labels{metricClusterKey: Cluster}).Set(m.RunningQueries)
	userFailures.With(prometheus.Labels{metricClusterKey: Cluster}).Set(m.UserErrorFailuresTotalCount)
}
