package monitor

import "crypto/x509"

type Callback func([]x509.Certificate, error)

func (monitor *Monitor) Next() ([]x509.Certificate, error) {
	result := []x509.Certificate{}

	for _, domain := range monitor.Domains {
		query := *monitor.BaseQuery
		query.Domain = domain

		certs, err := monitor.Api.Search(&query)
		if err != nil {
			return nil, err
		}

		result = append(result, certs...)
	}

	return result, nil
}

func (monitor *Monitor) Run(callback Callback) error {
	for {
		certs, err := monitor.Next()
		callback(certs, err)
	}
}
