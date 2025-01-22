package collectors

import (
	"fmt"
	"testing"
)

func TestQueryMFU(t *testing.T) {
	lokiURL := "http://127.0.0.1:3100"
	gotMfuValue, err := QueryMFU(lokiURL)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(gotMfuValue)
}
