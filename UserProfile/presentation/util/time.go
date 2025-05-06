package util

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func MakeCurrentTimestampMilliSecond() uint64 {
	return uint64(time.Now().UTC().UnixNano() / 1e6)
}

func ConvertTimestampToMilliSecond(timeNano int64) int64 {
	return int64(timeNano / 1e6)
}

func GetAgeFromBirthDay(birthDay string, sep string) int64 {
	dob := strings.Split(birthDay, sep)
	year, _ := strconv.ParseInt(dob[0], 10, 32)
	month, _ := strconv.ParseInt(dob[1], 10, 8)
	day, _ := strconv.ParseInt(dob[2], 10, 8)

	iBith := convertDateToInt(int(year), int(month), int(day))
	currentDateInit := convertDateToInt(time.Now().Year(), int(time.Now().Month()), time.Now().Day())
	age := (currentDateInit - iBith) / 10000
	return age
}

func ConvertDateTimeToInt(year, month, day, hour, minutes int) int64 {
	input := fmt.Sprintf("%04d%02d%02d%02d%02d", year, month, day, hour, minutes)
	intDateTime, _ := strconv.ParseInt(input, 10, 64)
	return intDateTime
}

func convertDateToInt(year, month, day int) int64 {
	input := fmt.Sprintf("%04d%02d%02d", year, month, day)
	dayToInt, _ := strconv.ParseInt(input, 10, 64)
	return dayToInt
}
