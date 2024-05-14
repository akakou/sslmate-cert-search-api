package monitor

import (
	"time"

	"github.com/akakou/sslmate-cert-search-api/api"
)

const DEFAULT_SLEEP = time.Minute * 5

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

func New(base *api.Query, api *api.SSLMateSearchAPI) *Monitor {
	return &Monitor{
		Query: base,
		Api:   api,
		Sleep: DEFAULT_SLEEP,
	}
}

func Default(domain string) *Monitor {
	query := DefaultQuery
	query.Domain = domain

	return New(
		&query,
		api.Default(),
	)
}
