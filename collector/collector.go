  package collector

  import (
	  "github.com/prometheus/client_golang/prometheus"
	  "log"
  )
  type Config struct {
	  host       string
  }

  var (
	  start = 1
	  namespace = "demo"
  )


  type Exporter struct {
	  host       string

	  version   *prometheus.Desc
	  total     *prometheus.GaugeVec
	  available *prometheus.GaugeVec
	  logger    log.Logger
  }

  func NewExporter(logger log.Logger, config *Config) *Exporter {
	  return &Exporter{
		  host:       config.host,
		  logger:     logger,
		  version: prometheus.NewDesc(
			  prometheus.BuildFQName(namespace, "", "version"),
			  "demo exporter version",
			  nil,
			  nil),
		  total: prometheus.NewGaugeVec(
			  prometheus.GaugeOpts{
				  Namespace: namespace,
				  Name:      "total",
				  Help:      "demo total",
			  },
			  []string{"region_id", "instance_id"},
		  ),
	  }
  }

  func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	  ch <- e.version
	  e.total.Describe(ch)
  }

  func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	  ch <- prometheus.MustNewConstMetric(e.version, prometheus.GaugeValue, 0.8)
	  //e.total.Reset()

	  start++
	  e.total.With(prometheus.Labels{"region_id": "region1", "instance_id": "instance1"}).Set(float64(start))
	  e.total.Collect(ch)
  }

