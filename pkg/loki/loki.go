package loki

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	APIPath        = "/loki/api/v1/query_range"
	QueryParamsKey = "query"
	StartParamsKey = "start"
	EndParamsKey   = "end"
)

func QueryLoki(lokiURL, query string, start, end int64) (*LokiRESP, error) {
	var (
		respBody []byte
		err      error
	)
	{
		u, err := url.Parse(lokiURL)
		if err != nil {
			return nil, err
		}
		u.Path = APIPath
		values := url.Values{} //拼接query参数
		values.Add(QueryParamsKey, query)
		if start > 0 {
			values.Add(StartParamsKey, fmt.Sprintf("%d", start))
		}
		if end > 0 {
			values.Add(EndParamsKey, fmt.Sprintf("%d", end))
		}
		u.RawQuery = values.Encode()

		// 发送 HTTP GET 请求
		resp, err := http.Get(u.String())
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		// 读取响应
		respBody, err = io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
	}

	var RESP = &LokiRESP{}
	if err = json.Unmarshal(respBody, RESP); err != nil {
		return nil, err
	}

	return RESP, nil
}

type LokiRESP struct {
	Status string   `json:"status"`
	Data   LokiData `json:"data"`
}

type LokiData struct {
	ResultType string       `json:"resultType"`
	Result     []LokiResult `json:"result"`
	Stats      any          `json:"stats"`
}

type LokiResult struct {
	Stream map[string]string `json:"stream"`
	Values LokiValues        `json:"values"`
}

type LokiValues []any
