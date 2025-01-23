package collectors

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
	"github.com/xiaoxlm/parse-to-metrics/global"
	"github.com/xiaoxlm/parse-to-metrics/pkg/loki"
	"github.com/xiaoxlm/parse-to-metrics/pkg/loki/parser"
	"time"
)

type MFU struct {
	query         string
	zeroTimestamp *time.Time
	gaugeVec      *prometheus.GaugeVec
}

func NewMFU() *MFU {
	return &MFU{
		query:         `{ai="mfu"} |= "mfu:" `,
		zeroTimestamp: nil,
		gaugeVec: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: "model_flops_util",        // 指标名称
			Help: "Model FLOPS Utilization", // 帮助信息
		}, []string{"node"}),
	}
}

func (mfu *MFU) GetGaugeVec() *prometheus.GaugeVec {
	return mfu.gaugeVec
}

func (mfu *MFU) SetGaugeVecValue() error {
	mfuValue, err := mfu.queryLoki(global.LokiURL)
	if err != nil {
		return err
	}

	if !mfuValue.GetRecords { // 只有没有数据的时候才设置成0
		mfu.gaugeVec.WithLabelValues(global.NodeLabel).Set(0)
		return nil
	}

	if mfuValue.Value != 0 {
		mfu.gaugeVec.WithLabelValues(global.NodeLabel).Set(mfuValue.Value)
	}

	return nil
}

func (mfu *MFU) queryLoki(lokiURL string) (mfuValue *MfuRESP, err error) {
	now := time.Now()
	var start int64 = now.Add(-35 * time.Second).UnixNano()
	var end int64 = now.UnixNano()

	resp, err := loki.QueryLoki(lokiURL, mfu.query, start, end)
	if err != nil {
		return nil, err
	}

	if len(resp.Data.Result) == 0 { // 查询不到日志数据
		logrus.Warning("no records from loki")
		return &MfuRESP{}, nil
	}

	var (
		totalMFU, mfuCount float64
		Labels             = make(map[string]string)
	)
	for _, res := range resp.Data.Result {
		Labels = res.Stream

		if len(res.Values) < 1 {
			continue
		}
		if len(res.Values) < 2 {
			logrus.Warningf("only get one value, content is %v", res.Values[0])
			continue
		}

		tmpValues := res.Values[1]
		values := tmpValues.([]interface{})[1].(string)

		parseValue, err := parser.ParseMFULog(values)
		if err != nil {
			return nil, err
		}
		if !parseValue.Find {
			continue
		}

		totalMFU += parseValue.Value
		mfuCount++
	}

	if mfuCount == 0 {
		return &MfuRESP{GetRecords: true, Labels: Labels}, nil
	}

	return &MfuRESP{
		GetRecords: true,
		Value:      totalMFU / mfuCount,
		Labels:     Labels,
	}, nil
}

type MfuRESP struct {
	GetRecords bool
	Value      float64
	Labels     map[string]string
}
