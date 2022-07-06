package main

import (
	"encoding/json"
	"io"
	"net/http"
	"sort"
	"time"
)

type session struct {
	Catalog   string
	User      string
	UserAgent string
}

type queryStats struct {
	CreateTime time.Time
	EndTime    time.Time
}

type rawQueryInfo struct {
	Query      string
	QueryId    string
	QueryStats queryStats
	QueryType  string
	Session    session
	State      string
}

type QueryInfo struct {
	Catalog        string
	CreateTime     time.Time
	DurationMs     int64
	QueryId        string
	QuerySizeBytes int
	QueryType      string
	State          string
	User           string
	UserAgent      string
}

var lastProcessed = time.Now()

func FetchQueryInfo() ([]QueryInfo, error) {
	client := &http.Client{
		Timeout: 2 * time.Second,
	}
	req, err := http.NewRequest("GET", Host+"/v1/query", nil)
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
	var raw []rawQueryInfo
	err = json.Unmarshal(body, &raw)
	if err != nil {
		return nil, err
	}
	// Sort queries by end time, oldest first.
	sort.Slice(raw, func(i, j int) bool {
		return raw[i].QueryStats.EndTime.Before(raw[j].QueryStats.EndTime)
	})

	queries := []QueryInfo{}
	for _, q := range raw {
		if q.QueryStats.EndTime.IsZero() {
			// the query is still running, skip it
			continue
		}

		if q.QueryStats.EndTime.After(lastProcessed) {
			queries = append(queries, QueryInfo{
				Catalog:        q.Session.Catalog,
				CreateTime:     q.QueryStats.CreateTime,
				DurationMs:     q.QueryStats.EndTime.Sub(q.QueryStats.CreateTime).Milliseconds(),
				QueryId:        q.QueryId,
				QuerySizeBytes: len(q.Query),
				QueryType:      q.QueryType,
				State:          q.State,
				User:           q.Session.User,
				UserAgent:      q.Session.UserAgent,
			})
		}
		if q.QueryStats.EndTime.After(lastProcessed) {
			lastProcessed = q.QueryStats.EndTime
		}
	}
	return queries, nil
}
