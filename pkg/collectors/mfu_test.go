package collectors

import (
	"fmt"
	"testing"
)

func TestQueryMFU(t *testing.T) {
	lokiURL := "http://10.129.60.70:3100"
	// only one
	//var start int64 = 1737623394122598000
	//var end int64 = 1737623429122598000

	//
	var start int64 = 1737628352591829000
	var end int64 = 1737628382591829000
	gotMfuValue, err := NewMFU().queryLoki(lokiURL, start, end)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(gotMfuValue)
}
