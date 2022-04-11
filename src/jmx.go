package main

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type metric struct {
	ObjectName string
	Attributes *json.RawMessage
}

type metricAttribute struct {
	Name string
	// Not all attributes are float64, but the ones we’re interested in either
	// are, or can be reliably parsed to one.
	Value float64
}

func fetchMetrics() ([]metric, error) {
	client := &http.Client{
		Timeout: 2 * time.Second,
	}
	req, err := http.NewRequest("GET", Host+"/v1/jmx/mbean", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-Trino-User", "trino-exporter")
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var metrics []metric
	err = json.Unmarshal(body, &metrics)
	if err != nil {
		return nil, err
	}
	return metrics, nil
}

type JmxMetrics struct {
	AbandonedQueriesTotalCount              float64
	CancelledQueriesTotalCount              float64
	CompletedQueriesTotalCount              float64
	ConsumedCpuTimeSecsTotalCount           float64
	ConsumedInputBytesTotalCount            float64
	ConsumedInputRowsTotalCount             float64
	ExecutionTimeAllTimeAvg                 float64
	ExecutionTimeAllTimeP95                 float64
	ExternalFailuresTotalCount              float64
	FailedQueriesTotalCount                 float64
	InsufficientResourcesFailuresTotalCount float64
	InternalFailuresTotalCount              float64
	QueuedQueries                           float64
	QueuedTimeAllTimeAvg                    float64
	QueuedTimeAllTimeP95                    float64
	RunningQueries                          float64
	UserErrorFailuresTotalCount             float64
}

func ReadJmxMetrics() (*JmxMetrics, error) {
	include := map[string]bool{
		"trino.execution:name=QueryExecution": true,
		"trino.execution:name=QueryManager":   true,
	}

	metrics, err := fetchMetrics()
	if err != nil {
		return nil, err
	}
	attributes := map[string]float64{}

	for _, m := range metrics {
		if include[m.ObjectName] {
			var metricAttrs []metricAttribute
			json.Unmarshal(*m.Attributes, &metricAttrs)
			for _, attr := range metricAttrs {
				attributes[attr.Name] = attr.Value
			}
		}
	}

	return &JmxMetrics{
		AbandonedQueriesTotalCount:              attributes["AbandonedQueries.TotalCount"],
		CancelledQueriesTotalCount:              attributes["CanceledQueries.TotalCount"], // sic
		CompletedQueriesTotalCount:              attributes["CompletedQueries.TotalCount"],
		ConsumedCpuTimeSecsTotalCount:           attributes["ConsumedCpuTimeSecs.TotalCount"],
		ConsumedInputBytesTotalCount:            attributes["ConsumedInputBytes.TotalCount"],
		ConsumedInputRowsTotalCount:             attributes["ConsumedInputRows.TotalCount"],
		ExecutionTimeAllTimeAvg:                 attributes["ExecutionTime.AllTime.Avg"],
		ExecutionTimeAllTimeP95:                 attributes["ExecutionTime.AllTime.P95"],
		ExternalFailuresTotalCount:              attributes["ExternalFailures.TotalCount"],
		FailedQueriesTotalCount:                 attributes["FailedQueries.TotalCount"],
		InsufficientResourcesFailuresTotalCount: attributes["InsufficientResourcesFailures.TotalCount"],
		InternalFailuresTotalCount:              attributes["InternalFailures.TotalCount"],
		QueuedQueries:                           attributes["QueuedQueries"],
		QueuedTimeAllTimeAvg:                    attributes["QueuedTime.AllTime.Avg"],
		QueuedTimeAllTimeP95:                    attributes["QueuedTime.AllTime.P95"],
		RunningQueries:                          attributes["RunningQueries"],
		UserErrorFailuresTotalCount:             attributes["UserErrorFailures.TotalCount"],
	}, nil
}
