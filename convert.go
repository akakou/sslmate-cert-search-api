package sslmate_cert_search_api

import (
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"strconv"
)

func convertJsonToEntry(jsonCert *SSLMateCertJson) (*SSLMateCertEntry, error) {
	id, err := strconv.Atoi(jsonCert.Id)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	raw, err := base64.StdEncoding.DecodeString(jsonCert.RawCert)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	x509Cert, err := x509.ParseCertificate(raw)
	if err != nil {
		return nil, fmt.Errorf("%v: %v", ERROR_FAILED_TO_PARSE_CERT, err)
	}

	return &SSLMateCertEntry{
		Id:   id,
		Cert: x509Cert,
	}, nil
}
