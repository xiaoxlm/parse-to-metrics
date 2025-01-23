package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/xiaoxlm/parse-to-metrics/global"
	pkgPrometheus "github.com/xiaoxlm/parse-to-metrics/pkg/prometheus"
	"github.com/xiaoxlm/parse-to-metrics/pkg/prometheus/collectors"
	"log"
	"net/http"
	"time"
)

var mfuGaugeVec *prometheus.GaugeVec

func init() {
	global.InitCheck()

	//os.Environ()
	mfuGaugeVec = collectors.NewMFUGaugeVec()
	http.Handle("/metrics", promhttp.HandlerFor(pkgPrometheus.NewMetricsRegistry(map[string]string{
		"service":    "parse-to-metrics",
		"ai_metrics": global.AiMetricsLabel,
	}, mfuGaugeVec), promhttp.HandlerOpts{}))
}

func main() {
	go func() {
		for range time.Tick(1 * time.Second) {
			if err := setGaugeVecValue(); err != nil {
				fmt.Printf("[ERROR] setGaugeVecValue error. err:%v \n", err)
			}
		}
	}()
	fmt.Println("Starting server at :9133")
	if err := http.ListenAndServe(":9133", nil); err != nil {
		log.Fatalln(err)
	}
}

func setGaugeVecValue() error {

	mfuValue, err := collectors.QueryMFU(global.LokiURL)
	if err != nil {
		return err
	}

	mfuGaugeVec.WithLabelValues(global.NodeLabel).Set(mfuValue.Value)
	return nil
}
