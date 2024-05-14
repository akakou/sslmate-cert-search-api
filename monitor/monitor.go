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

	monitor.Query.After = index.Last

	return result, index, nil
}

func (monitor *Monitor) nextWithCallbackAndSleep(callback Callback) {
	time.Sleep(monitor.Sleep)
	certs, last, err := monitor.Next()
	callback(certs, last, err)
}

func (monitor *Monitor) Loop(callback Callback) {
	for {
		monitor.nextWithCallbackAndSleep(callback)
	}
}

func (monitors *Monitors) Next(callback Callback) {
	for _, monitor := range monitors.Monitors {
		monitor.nextWithCallbackAndSleep(callback)
	}
}

func (monitors *Monitors) Loop(callback Callback) {
	for {
		monitors.Next(callback)
	}
}
