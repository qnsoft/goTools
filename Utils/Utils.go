package Utils

import (
	"crypto/md5"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"
	"time"
)

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func GetClientIP(RemoteAddr string) (string, error) {
	var clientIP string
	var err error = nil
	index := strings.Index(RemoteAddr, ":")
	if index > 0 {
		clientIP = string([]rune(RemoteAddr)[:index])
	} else {
		err = errors.New("get clientIP fail!")
	}
	return clientIP, err
}

func MD5(str string) string {
	has := md5.Sum([]byte(str))
	md5str := fmt.Sprintf("%x", has)
	return md5str
}

func Base64_encode(src string) string {
	return base64.StdEncoding.EncodeToString([]byte(src))
}

func Base64_decode(dec string) string {
	buf, err := base64.StdEncoding.DecodeString(dec)
	if err == nil {
		return string(buf)
	} else {
		return ""
	}
}

func TodayTimestamp() int64 {
	now := time.Now()
	timestamp := now.Unix() - int64(now.Second()) - int64((60 * now.Minute())) - int64((3600 * now.Hour()))
	return timestamp
}
func HourTimestamp() int64 {
	now := time.Now()
	timestamp := now.Unix() - int64(now.Second()) - int64((60 * now.Minute()))
	return timestamp
}

func GetDateTimeStamp(date string) int64 {
	loc, _ := time.LoadLocation("Local")
	layout := ""
	firstFalg := strings.Count(date, "-")
	if firstFalg > 1 {
		secondFalg := strings.Count(date, ":")
		if secondFalg > 1 {
			layout = "2006-01-02 15:04:05"
		} else if secondFalg > 0 {
			layout = "2006-01-02 15:04"
		} else {
			thirdFalg := strings.Count(date, " ")
			if thirdFalg > 0 {
				layout = "2006-01-02 15"
			} else {
				layout = "2006-01-02"
			}
		}
	} else if firstFalg > 0 {
		layout = "2006-01"
	} else {
		layout = "2006"
	}
	tc, err := time.ParseInLocation(layout, date, loc)
	if err == nil {
		return tc.Unix()
	} else {
		fmt.Println(err)
		return 0
	}
}

func GetLastMonthStart() int64 {
	now := time.Now()
	month := int(now.Month())
	year := now.Year()
	var dateStr string = "%d-%d"
	if month > 1 {
		month -= 1
		if month < 10 {
			dateStr = "%d-0%d"
		}
	} else {
		month = 12
		year -= 1
	}
	lastMonthStr := fmt.Sprintf(dateStr, year, month)
	return GetDateTimeStamp(lastMonthStr)
}

func GetThisMonthStart() int64 {
	now := time.Now()
	month := int(now.Month())
	year := now.Year()
	var dateStr string = "%d-%d"
	if month < 10 {
		dateStr = "%d-0%d"
	}
	thisMonthStr := fmt.Sprintf(dateStr, year, month)
	return GetDateTimeStamp(thisMonthStr)
}

func GetLastYearStart() int64 {
	now := time.Now()
	year := now.Year() - 1
	lastYearStr := fmt.Sprintf("%d", year)
	return GetDateTimeStamp(lastYearStr)
}

func GetThisYearStart() int64 {
	now := time.Now()
	year := now.Year()
	thisYearStr := fmt.Sprintf("%d", year)
	return GetDateTimeStamp(thisYearStr)
}
