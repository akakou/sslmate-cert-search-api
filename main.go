package sslmate_cert_search_api

import (
	"crypto/x509"
	"os"
)

const SSLMATE_API_URL = "https://api.certspotter.com/v1/issuances"

type SSLMateCertJson struct {
	Id      string `json:"id"`
	RawCert string `json:"cert_der"`
}

type SSLMateCertEntry struct {
	Id   int
	Cert *x509.Certificate
}

type Query struct {
	Domain            string
	IncludeSubdomains bool
	MatchWildcards    bool
	After             string
	Expand            string
}

type Index struct {
	First string
	Last  string
}

type Json []map[string]interface{}

type SSLMateSearchAPI struct {
	Token string
}

func New(token string) *SSLMateSearchAPI {
	return &SSLMateSearchAPI{Token: token}
}

func Default() *SSLMateSearchAPI {
	token := os.Getenv("SSLMATE_API_TOKEN")
	return &SSLMateSearchAPI{Token: token}
}
