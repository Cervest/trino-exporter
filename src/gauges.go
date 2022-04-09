package main

import "github.com/prometheus/client_golang/prometheus"

var abandonedQueries = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "queries_abandoned_count_total",
})
var cancelledQueries = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "queries_cancelled_count_total",
})
var completedQueries = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "queries_completed_count_total",
})
var failedQueries = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "queries_failed_count_total",
})
var queuedQueries = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "queries_queued_count_total",
})
var runningQueries = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "queries_running_count_total",
})
var inputBytes = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "input_bytes_total",
})
var inputRows = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "input_rows_total",
})
var cpuTime = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "cpu_time_seconds_total",
})
var execTimeAvg = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "execution_time_milliseconds_average",
})
var execTimeP95 = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "execution_time_milliseconds_p95",
})
var queuedTimeAvg = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "queued_time_milliseconds_average",
})
var queuedTimeP95 = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "queued_time_milliseconds_p95",
})
var externalFailures = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "external_failures_count_total",
})
var internalFailures = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "internal_failures_count_total",
})
var userFailures = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "user_failures_count_total",
})
var resourceFailures = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "insufficient_resources_failures_count_total",
})

var Gauges = []prometheus.Gauge{
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
	abandonedQueries.Set(m.AbandonedQueriesTotalCount)
	cancelledQueries.Set(m.CancelledQueriesTotalCount)
	completedQueries.Set(m.CompletedQueriesTotalCount)
	cpuTime.Set(m.ConsumedCpuTimeSecsTotalCount)
	execTimeAvg.Set(m.ExecutionTimeAllTimeAvg)
	execTimeP95.Set(m.ExecutionTimeAllTimeP95)
	externalFailures.Set(m.ExternalFailuresTotalCount)
	failedQueries.Set(m.FailedQueriesTotalCount)
	inputBytes.Set(m.ConsumedInputBytesTotalCount)
	inputRows.Set(m.ConsumedInputRowsTotalCount)
	internalFailures.Set(m.InternalFailuresTotalCount)
	queuedQueries.Set(m.QueuedQueries)
	queuedTimeAvg.Set(m.QueuedTimeAllTimeAvg)
	queuedTimeP95.Set(m.QueuedTimeAllTimeP95)
	resourceFailures.Set(m.InsufficientResourcesFailuresTotalCount)
	runningQueries.Set(m.RunningQueries)
	userFailures.Set(m.UserErrorFailuresTotalCount)
}
