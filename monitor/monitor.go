package monitor

import (
	"crypto/x509"
	"time"

	"github.com/akakou/sslmate-cert-search-api/api"
)

type Callback func([]x509.Certificate, *api.Index, error)

func (monitor *Monitor) Next() ([]x509.Certificate, *api.Index, error) {
	result := []x509.Certificate{}

	certs, index, err := monitor.Api.Search(monitor.Query)
	if err != nil {
		return nil, index, err
	}

	result = append(result, certs...)

	time.Sleep(monitor.Sleep)

	return result, index, nil
}

func (monitor *Monitor) run(callback Callback) {
	certs, last, err := monitor.Next()
	callback(certs, last, err)
}

func (monitor *Monitor) Run(callback Callback) {
	for {
		monitor.run(callback)
	}
}

func (monitors Monitors) Run(callback Callback) {
	for _, monitor := range monitors {
		monitor.Run(callback)
	}
}
