package collectors

import (
	"fmt"
	"testing"
)

func TestQueryMFU(t *testing.T) {
	lokiURL := "http://10.129.60.70:3100"
	gotMfuValue, err := NewMFU().queryLoki(lokiURL)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(gotMfuValue)
}
