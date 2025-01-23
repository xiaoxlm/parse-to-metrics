package collectors

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/xiaoxlm/parse-to-metrics/pkg/loki"
	"github.com/xiaoxlm/parse-to-metrics/pkg/loki/parser"
	"time"
)

func NewMFUGaugeVec() *prometheus.GaugeVec {
	return prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "model_flops_util",        // 指标名称
		Help: "Model FLOPS Utilization", // 帮助信息
	}, []string{"node"})
}

type MfuRESP struct {
	Value  float64
	Labels map[string]string
}

func QueryMFU(lokiURL string) (mfuValue *MfuRESP, err error) {
	query := `{ai="mfu"} |= "mfu:" `

	now := time.Now()
	var start int64 = now.Add(-5 * time.Second).UnixNano()
	var end int64 = now.UnixNano()

	resp, err := loki.QueryLoki(lokiURL, query, start, end)
	if err != nil {
		return nil, err
	}

	var (
		totalMFU, mfuCount float64
		Labels             = make(map[string]string)
	)
	for _, res := range resp.Data.Result {
		Labels = res.Stream

		tmpValues := res.Values[1]
		values := tmpValues.([]interface{})[1].(string)

		mfu, err := parser.ParseMFULog(values)
		if err != nil {
			return nil, err
		}
		if !mfu.Find {
			continue
		}

		totalMFU += mfu.Value
		mfuCount++
	}

	if mfuCount == 0 {
		return &MfuRESP{Labels: Labels}, nil
	}

	return &MfuRESP{
		Value:  totalMFU / mfuCount,
		Labels: Labels,
	}, nil
}
