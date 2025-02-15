package LghTool

import (
	"fmt"
	"strings"
	"time"
)

// Format: "2006-01-02 15:04:05"
func GetPreday(timeS string) (string, error) {
	parsedTime, err := time.Parse("2006-01-02 15:04:05", timeS)
	if err != nil {
		return "", fmt.Errorf("时间解析失败: %v，要求格式: %s", err, "2006-01-02 15:04:05")
	}
	str1 := parsedTime.Add(-1 * 24 * time.Hour).String()
	if strings.HasSuffix(str1, "UTC") {
		arr := strings.Split(str1, " +")
		return arr[0], nil
	} else {
		return str1, nil
	}
}

// Format: "2006-01-02 15:04:05"
func GetNextDay(timeS string, needReduceOneSec bool) string {
	parsedTime, err := time.Parse("2006-01-02 15:04:05", timeS)
	if err != nil {
		fmt.Println("时间解析失败:", err)
		return ""
	}
	var str1 string
	if needReduceOneSec {
		str1 = parsedTime.Add(24*time.Hour - time.Second).String()
	} else {
		str1 = parsedTime.Add(24 * time.Hour).String()
	}
	if strings.HasSuffix(str1, "UTC") {
		arr := strings.Split(str1, " +")
		return arr[0]
	} else {
		return str1
	}
}

// Format: "2006-01-02"
func BeforeDay(day1, day2 string) bool {
	parsedTime1, _ := time.Parse("2006-01-02", strings.Split(day1, " ")[0])
	parsedTime2, _ := time.Parse("2006-01-02", strings.Split(day2, " ")[0])
	return parsedTime1.Unix() < parsedTime2.Unix()
}

// Format: "2006-01-02 15:04:05"
func BetweenDay(startTime, endTime, target string) bool {
	parsedTime1, _ := time.Parse("2006-01-02 15:04:05", startTime)
	parsedTime2, _ := time.Parse("2006-01-02 15:04:05", endTime)
	parsedTime3, _ := time.Parse("2006-01-02 15:04:05", target)
	return parsedTime3.Unix() >= parsedTime1.Unix() && parsedTime3.Unix() <= parsedTime2.Unix()
}
