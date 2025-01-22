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

var (
	aiMetricsLabel = os.Getenv("AI_METRICS_LABEL")
	nodeLabel      = os.Getenv("NODE_LABEL")
	lokiURL        = os.Getenv("LOKI_URL")
)

func initCheck() {
	if aiMetricsLabel == "" {
		log.Fatalln("[WARNING] env var AI_METRICS_LABEL  is empty. ")
	}

	if nodeLabel == "" {
		log.Fatalln("env var NODE_LABEL is empty")
	}

	if lokiURL == "" {
		log.Fatalln("env var LOKI_URL is empty")
	}
}

var mfuGaugeVec *prometheus.GaugeVec

func init() {
	initCheck()

	//os.Environ()
	mfuGaugeVec = collectors.NewMFUGaugeVec()
	http.Handle("/metrics", promhttp.HandlerFor(pkgPrometheus.NewMetricsRegistry(map[string]string{
		"service":    "parse-to-metrics",
		"ai_metrics": aiMetricsLabel,
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

	mfuValue, err := collectors.QueryMFU(lokiURL)
	if err != nil {
		return err
	}

	mfuGaugeVec.WithLabelValues(nodeLabel).Set(mfuValue.Value)
	return nil
}
