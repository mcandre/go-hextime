package hextime

import (
	"fmt"
	"time"
)

func Hex(t time.Time) string {
	tUTC := t.UTC()
	secondsPastMidnight := tUTC.Hour()*3600 + tUTC.Minute()*60 + tUTC.Second()

	hexHour := secondsPastMidnight / 4096
	hexMinute := (secondsPastMidnight % 4096) / 16
	hexSecond := secondsPastMidnight % 16

	return fmt.Sprintf("%x_%02x_%x", hexHour, hexMinute, hexSecond)
}

func Hextime() string {
	return Hex(time.Now())
}
