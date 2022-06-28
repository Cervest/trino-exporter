package main

import (
	"encoding/json"
	"io"
	"math"
	"net/http"
	"time"
)

type metric struct {
	ObjectName string
	Attributes *json.RawMessage
}

type metricAttribute struct {
	Name  string
	Type  string
	Value interface{}
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
	ActiveNodeCount                         float64
	CancelledQueriesTotalCount              float64
	CompletedQueriesTotalCount              float64
	ConsumedCpuTimeSecsTotalCount           float64
	ConsumedInputBytesTotalCount            float64
	ConsumedInputRowsTotalCount             float64
	ExecutionTimeAllTimeAvg                 float64
	ExecutionTimeAllTimeP95                 float64
	ExternalFailuresTotalCount              float64
	FailedQueriesTotalCount                 float64
	InactiveNodeCount                       float64
	InsufficientResourcesFailuresTotalCount float64
	InternalFailuresTotalCount              float64
	QueuedQueries                           float64
	QueuedTimeAllTimeAvg                    float64
	QueuedTimeAllTimeP95                    float64
	RunningQueries                          float64
	ShuttingDownNodeCount                   float64
	UserErrorFailuresTotalCount             float64
}

func ReadJmxMetrics() (*JmxMetrics, error) {
	include := map[string]bool{
		"trino.execution:name=QueryExecution":      true,
		"trino.execution:name=QueryManager":        true,
		"trino.metadata:name=DiscoveryNodeManager": true,
	}

	metrics, err := fetchMetrics()
	if err != nil {
		return nil, err
	}
	attributes := map[string]float64{}

	for _, m := range metrics {
		if include[m.ObjectName] {
			var metricAttrs []metricAttribute
			err = json.Unmarshal(*m.Attributes, &metricAttrs)
			if err != nil {
				return nil, err
			}
			for _, attr := range metricAttrs {
				switch attr.Value.(type) {
				case int:
					attributes[attr.Name] = attr.Value.(float64)
				case float64:
					attributes[attr.Name] = attr.Value.(float64)
				default:
					// This handles two cases. The first is where the attribute is not a numeric
					// type, in which case we’re not interested in it, because it can’t be used
					// as a metric. The second is when the returned value is the string "NaN",
					// which it how it represents NaN for double attributes.
					attributes[attr.Name] = math.NaN()
				}
			}
		}
	}

	return &JmxMetrics{
		AbandonedQueriesTotalCount:              attributes["AbandonedQueries.TotalCount"],
		ActiveNodeCount:                         attributes["ActiveNodeCount"],
		CancelledQueriesTotalCount:              attributes["CanceledQueries.TotalCount"], // sic
		CompletedQueriesTotalCount:              attributes["CompletedQueries.TotalCount"],
		ConsumedCpuTimeSecsTotalCount:           attributes["ConsumedCpuTimeSecs.TotalCount"],
		ConsumedInputBytesTotalCount:            attributes["ConsumedInputBytes.TotalCount"],
		ConsumedInputRowsTotalCount:             attributes["ConsumedInputRows.TotalCount"],
		ExecutionTimeAllTimeAvg:                 attributes["ExecutionTime.AllTime.Avg"],
		ExecutionTimeAllTimeP95:                 attributes["ExecutionTime.AllTime.P95"],
		ExternalFailuresTotalCount:              attributes["ExternalFailures.TotalCount"],
		FailedQueriesTotalCount:                 attributes["FailedQueries.TotalCount"],
		InactiveNodeCount:                       attributes["InactiveNodeCount"],
		InsufficientResourcesFailuresTotalCount: attributes["InsufficientResourcesFailures.TotalCount"],
		InternalFailuresTotalCount:              attributes["InternalFailures.TotalCount"],
		QueuedQueries:                           attributes["QueuedQueries"],
		QueuedTimeAllTimeAvg:                    attributes["QueuedTime.AllTime.Avg"],
		QueuedTimeAllTimeP95:                    attributes["QueuedTime.AllTime.P95"],
		RunningQueries:                          attributes["RunningQueries"],
		ShuttingDownNodeCount:                   attributes["ShuttingDownNodeCount"],
		UserErrorFailuresTotalCount:             attributes["UserErrorFailures.TotalCount"],
	}, nil
}
