package sslmate_cert_search_api

import (
	"testing"
)

func TestSearch(t *testing.T) {
	api := Default()

	query := Query{
		Domain:            "example.com",
		IncludeSubdomains: true,
		MatchWildcards:    true,
		After:             "",
		Expand:            "",
	}

	certs, err := api.Search(&query)
	if err != nil {
		t.Fatalf("failed to search certs: %v", err)
	}

	t.Logf("%v: %v %v\n", certs[0].Id, certs[0].Cert, certs[0].Cert.DNSNames)
}
