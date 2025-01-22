package parser

import (
	"fmt"
	"regexp"
	"strconv"
)

type MFUResult struct {
	Key   string
	Value float64
	Find  bool
}

func ParseMFULog(text string) (*MFUResult, error) {
	key := "mfu"
	// 正则表达式，匹配类似 "key: value" 的格式
	regex := fmt.Sprintf(`(?i)%s:\s*([0-9.]+)`, regexp.QuoteMeta(key))
	re := regexp.MustCompile(regex)

	// 查找匹配的结果
	match := re.FindStringSubmatch(text)
	if len(match) < 2 {
		return &MFUResult{}, nil
	}

	value, err := strconv.ParseFloat(match[1], 64)
	if err != nil {
		return nil, err
	}

	return &MFUResult{
		Key:   key,
		Value: value,
		Find:  true,
	}, nil
}
