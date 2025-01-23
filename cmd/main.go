package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"github.com/xiaoxlm/parse-to-metrics/global"
	"github.com/xiaoxlm/parse-to-metrics/pkg/collectors"
	pkgPrometheus "github.com/xiaoxlm/parse-to-metrics/pkg/prometheus"
	"log"
	"net/http"
	"time"
)

var mfuCollector *collectors.MFU

func init() {
	global.InitCheck()
	mfuCollector = collectors.NewMFU()

	//os.Environ()
	mfuGaugeVec := mfuCollector.GetGaugeVec()
	http.Handle("/metrics", promhttp.HandlerFor(pkgPrometheus.NewMetricsRegistry(map[string]string{
		"service":    "parse-to-metrics",
		"ai_metrics": global.AiMetricsLabel,
	}, mfuGaugeVec), promhttp.HandlerOpts{}))
}

func main() {
	go func() {
		for {
			if err := mfuCollector.SetGaugeVecValue(); err != nil {
				logrus.Errorf("setGaugeVecValue error. err:%v ", err)
			}
			time.Sleep(2 * time.Second)
		}
	}()

	fmt.Println("Starting server at :9133")
	if err := http.ListenAndServe(":9133", nil); err != nil {
		log.Fatalln(err)
	}
}
