package main

import (
	"flag"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"time"
)

var addr = flag.String("127.0.0.1", ":9133", "The address to listen on for HTTP requests.")

var (
	opsQueued = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "our_company",
		Subsystem: "blob_storage",
		Name:      "ops_queued",
		Help:      "Number of blob storage operations waiting to be processed",
	})
)

func init() {

	prometheus.MustRegister(opsQueued)

}

func main() {
	flag.Parse()
	go func() {
		for {
			opsQueued.Add(4)
			time.Sleep(time.Second * 1)
		}
	}()
	http.Handle("/metrics", promhttp.Handler())

	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
