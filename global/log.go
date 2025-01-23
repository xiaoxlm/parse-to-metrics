package global

import "github.com/xiaoxlm/parse-to-metrics/pkg/log"

func init() {
	(&log.Log{}).SetDefaults().Build()
}
