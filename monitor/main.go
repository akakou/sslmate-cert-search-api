package monitor

import (
	"time"

	"github.com/akakou/sslmate-cert-search-api/api"
)

const DEFAULT_SLEEP = time.Second * 20

var DefaultQuery = api.Query{
	IncludeSubdomains: true,
	MatchWildcards:    true,
	After:             "",
}

type Monitor struct {
	Domains string
	Api     *api.SSLMateSearchAPI
	Query   *api.Query
	Sleep   time.Duration
}

type Monitors []Monitor

func newMonitor(base *api.Query, api *api.SSLMateSearchAPI, sleep time.Duration) *Monitor {
	return &Monitor{
		Query: base,
		Api:   api,
		Sleep: sleep,
	}
}

func DefaultMonitor(domain string) *Monitor {
	query := DefaultQuery
	query.Domain = domain

	return newMonitor(
		&query,
		api.Default(),
		DEFAULT_SLEEP,
	)
}
