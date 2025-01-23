package global

import (
	"fmt"
	"log"
	"os"
)

var (
	AiMetricsLabel = os.Getenv("AI_METRICS_LABEL")
	NodeLabel      = os.Getenv("NODE_LABEL")
	LokiURL        = os.Getenv("LOKI_URL")
)

func InitCheck() {
	fmt.Printf(`
====env var====
AI_METRICS_LABEL=%s
NODE_LABEL=%s
LOKI_URL=%s
====env var====

`, AiMetricsLabel, NodeLabel, LokiURL)
	if AiMetricsLabel == "" {
		log.Fatalln("[WARNING] env var AI_METRICS_LABEL  is empty. ")
	}

	if NodeLabel == "" {
		log.Fatalln("env var NODE_LABEL is empty")
	}

	if LokiURL == "" {
		log.Fatalln("env var LOKI_URL is empty")
	}

}
