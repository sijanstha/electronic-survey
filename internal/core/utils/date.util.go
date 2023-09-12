package utils

import (
	"fmt"
	"time"
)

const (
	pollDateTimeLayout = "2006-01-02T15:04:05"
	apiDateLayout = "2006-01-02T15:01"
)

func ParseDateTime(str string) (time.Time, error) {
	return ParseDateTimeWithLayout(str, pollDateTimeLayout)
}

func ParseDateTimeWithLayout(str string, layout string) (time.Time, error) {
	datetime, err := time.Parse(layout, str)
	if err != nil {
		return datetime, fmt.Errorf("cannot parse date time: %s", datetime)
	}

	return datetime, nil
}

func Now() string {
	return time.Now().UTC().Format(apiDateLayout)
}