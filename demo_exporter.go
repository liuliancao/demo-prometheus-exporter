package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/version"
	"log"
	"net/http"
	"demo_exporter/collector"
)

const indexContent = `<html>
             <head><title>Demo Exporter</title></head>
             <body>
             <h1>Demo Exporter</h1>
             <p><a href='` + "/metrics" + `'>Metrics</a></p>
             </body>
             </html>`

func init() {
	prometheus.MustRegister(version.NewCollector("exporter"))
}

func main() {
	config := &collector.Config{
	}

	var logger log.Logger
	exporter := collector.NewExporter(logger, config)
	prometheus.MustRegister(exporter)

	// Serve metrics
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(indexContent))
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s", ":8182"), nil))
}

