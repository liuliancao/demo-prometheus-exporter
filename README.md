# A demo promethes exporter
简单说下，几个重要的点

-   在prometheus里面，[prometheus.MustRegister()](https://pkg.go.dev/github.com/prometheus/client_golang/prometheus#MustRegister)表示一定会把这个数据注册
    到采集器里面，当然你也可以MustRegister一个数据
-   在prometheus里面有多种数据类型，比如Counter, Gauge等，具体你可以去
    <https://pkg.go.dev/github.com/prometheus/client_golang/prometheus#pkg-types>
    去查看
-   为什么echo出来一个metrics 比如aaa{test\_label="xxx"} 3.5？
    也可以做，但是当我们需要采集的指标越来越多的时候，这样会些微影响可读
    性，并且无法充分利用prometheus的数据结构。
-   collector怎么编写
    collector采集器需要定义采集的数据
    collector定义：

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
    		  ), // 定义gauge类型带标签region_id, instance_id的数据结构
    	  }

collector的describe和collect方法编写，一个collector都需要这两种方法，
否则会报错。

    func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
    	 ch <- e.version
    	 e.total.Describe(ch)
     }
    
     func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
    	 ch <- prometheus.MustNewConstMetric(e.version, prometheus.GaugeValue, 0.8)
    	 e.total.Reset()
    
    	 start++
    	 e.total.With(prometheus.Labels{"region_id": "region1", "instance_id": "instance1"}).Set(float64(start))
    	 e.total.Collect(ch)
     }

我这边理解是每个采集值都要在describe和collect体现。

对于有标签，动态的gauge的，需要通过With().Set()方法更改
并且需要增加 `e.total.Collect(ch)` 这样的方法去更新。

这样执行的结果就是demo\_total这个值每次都会+1，从2开始。


