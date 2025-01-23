package parser

import (
	"fmt"
	"testing"
)

func TestParseMFULog(t *testing.T) {
	text := "15:11:58 [INFO] [Rank 0] step: 100, loss: 6.5628, tokens_per_second: 36225.40, mfu: 44.52, diloco_peers: 2"

	v, err := ParseMFULog(text)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(v)
}

func TestParseMFULog2(t *testing.T) {
	text := "15:11:58 [INFO] [Rank 0] step: 100, loss: 6.5628, tokens_per_second: 36225.40, mfu: 44.52, diloco_peers: 2"

	ParseMFULog2(text)

}
