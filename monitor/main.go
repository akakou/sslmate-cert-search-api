package monitor

import api "github.com/akakou/sslmate_cert_search_api"

var DefaultQuery = api.Query{
	IncludeSubdomains: true,
	MatchWildcards:    true,
	After:             "",
}

type Monitor struct {
	Domains   []string
	Api       *api.SSLMateSearchAPI
	BaseQuery *api.Query
}

func New(domains []string, api *api.SSLMateSearchAPI, base *api.Query) *Monitor {
	return &Monitor{
		Domains:   domains,
		Api:       api,
		BaseQuery: base,
	}
}

func Default(domains []string) *Monitor {
	return New(
		domains,
		api.Default(),
		&DefaultQuery,
	)
}
