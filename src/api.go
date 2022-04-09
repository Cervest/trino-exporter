package main

import (
	"encoding/json"
	"io"
	"net/http"
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

var lastProcessed, _ = time.Parse("2006-01-02", "1970-01-01")

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
	var raw []rawQueryInfo
	err = json.Unmarshal(body, &raw)
	if err != nil {
		return nil, err
	}

	queries := []QueryInfo{}
	latestQuery := lastProcessed
	for _, q := range raw {
		if q.QueryStats.CreateTime.After(lastProcessed) {
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
		if q.QueryStats.CreateTime.After(latestQuery) {
			latestQuery = q.QueryStats.CreateTime
		}
	}
	lastProcessed = latestQuery
	return queries, nil
}
