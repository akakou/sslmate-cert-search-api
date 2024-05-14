package main

import (
	"crypto/x509"
	"fmt"
	"log"
	"time"

	"github.com/akakou/sslmate-cert-search-api/api"
	"github.com/akakou/sslmate-cert-search-api/monitor"
)

const DEBUG_SLEEP = time.Second * 20

func main() {
	m1 := monitor.DefaultMonitor("example.com")
	m1.Sleep = DEBUG_SLEEP

	m2 := monitor.DefaultMonitor("ochanoco.com")
	m2.Sleep = DEBUG_SLEEP

	monitors := monitor.Monitors([]monitor.Monitor{*m1, *m2})

	monitors.Run(func(certs []x509.Certificate, index *api.Index, err error) {
		if err != nil {
			log.Fatalf("failed to get certs	: %v", err)
		}
		fmt.Printf("Getting (%v ~ %v)\n", index.First, index.Last)
		log.Printf("%v", certs)
	})
}
