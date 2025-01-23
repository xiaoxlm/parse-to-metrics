package parser

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

type MFUResult struct {
	Key   string
	Time  time.Time
	Value float64
	Find  bool
}

func ParseMFULog(text string) (*MFUResult, error) {
	key := "mfu"
	// 正则表达式，匹配类似 "key: value" 的格式
	mfuRegex := fmt.Sprintf(`(?i)%s:\s*([0-9.]+)`, regexp.QuoteMeta(key))
	re := regexp.MustCompile(mfuRegex)

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

//type LogData struct {
//	Mfu  float64   `json:"mfu"`
//	Time time.Time `json:"time"`
//}
//
//func ParseMFULog2(text string) {
//	// 正则表达式，匹配时间和 mfu
//	timeRegex := `(\d{2}:\d{2}:\d{2})`
//	mfuRegex := `(?i)mfu:\s*([0-9.]+)`
//
//	// 查找时间
//	timeRe := regexp.MustCompile(timeRegex)
//	timeMatch := timeRe.FindStringSubmatch(text)
//	if len(timeMatch) == 0 {
//		fmt.Println("未找到时间")
//		return
//	}
//
//	// 获取当前日期（年月日）
//	currentDate := time.Now().Format("2006-01-02")
//
//	// 拼接当前日期和提取到的时间（时分秒）
//	fullTimeStr := currentDate + " " + timeMatch[1]
//
//	// 解析时间字符串为 time.Time，设置时区为中国上海
//	location, err := time.LoadLocation("Asia/Shanghai")
//	if err != nil {
//		fmt.Println("加载时区失败:", err)
//		return
//	}
//
//	// 解析拼接后的时间字符串为 time.Time
//	parsedTime, err := time.ParseInLocation("2006-01-02 15:04:05", fullTimeStr, location)
//	if err != nil {
//		fmt.Println("解析时间失败:", err)
//		return
//	}
//
//	// 查找 mfu 的值
//	mfuRe := regexp.MustCompile(mfuRegex)
//	mfuMatch := mfuRe.FindStringSubmatch(text)
//	if len(mfuMatch) == 0 {
//		fmt.Println("未找到 mfu 的值")
//		return
//	}
//
//	// 转换 mfu 的值为浮点数
//	mfu, err := strconv.ParseFloat(mfuMatch[1], 64)
//	if err != nil {
//		fmt.Println("转换 mfu 失败:", err)
//		return
//	}
//
//	// 创建结构体对象
//	logData := LogData{
//		Mfu:  mfu,
//		Time: parsedTime,
//	}
//
//	// 输出结构体
//	fmt.Printf("%+v\n", logData)
//	fmt.Printf("%+v\n", logData.Time.Unix())
//
//}
