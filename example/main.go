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
	m2 := monitor.DefaultMonitor("ochano.co")
	m2.Sleep = DEBUG_SLEEP

	monitors := monitor.Monitors{
		Monitors: []monitor.Monitor{*m2},
	}

	monitors.Loop(func(certs []x509.Certificate, index *api.Index, err error) {
		if err != nil {
			log.Fatalf("failed to get certs	: %v", err)
		}
		fmt.Printf("Getting (%v ~ %v)\n", index.First, index.Last)
		log.Printf("%v", certs)
	})
}
