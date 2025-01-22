package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	pkgPrometheus "github.com/xiaoxlm/parse-to-metrics/pkg/prometheus"
	"github.com/xiaoxlm/parse-to-metrics/pkg/prometheus/collectors"
	"log"
	"net/http"
	"os"
	"time"
)

var mfuGaugeVec *prometheus.GaugeVec

func init() {
	aiMetrics := os.Getenv("AI_METRICS")
	if aiMetrics == "" {
		fmt.Println("[WARNING] AI_METRICS env is empty. set default value 'mfu'")
		aiMetrics = "mfu"
	}

	//os.Environ()
	mfuGaugeVec = collectors.NewMFUGaugeVec()
	http.Handle("/metrics", promhttp.HandlerFor(pkgPrometheus.NewMetricsRegistry(map[string]string{
		"service":    "parse-to-metrics",
		"ai_metrics": aiMetrics,
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
	lokiURL := os.Getenv("LOKI_URL")
	if lokiURL == "" {
		return fmt.Errorf("loki url env is empty")
	}

	mfuValue, err := collectors.QueryMFU(lokiURL)
	if err != nil {
		return err
	}

	nodeValue, ok := mfuValue.Labels["ai"]
	if !ok {
		nodeValue = "localhost"
	}

	mfuGaugeVec.WithLabelValues(nodeValue).Set(mfuValue.Value)
	return nil
}
