// Package hextime provides conversion functions for formatting time in hexadecimal.
package hextime

import (
	"fmt"
	"time"
)

// Version is semver.
var Version = "0.0.2"

// Hex formats the given time into a hexadecimal notation.
func Hex(t time.Time) string {
	tUTC := t.UTC()
	secondsPastMidnight := tUTC.Hour()*3600 + tUTC.Minute()*60 + tUTC.Second()

	hexHour := secondsPastMidnight / 4096
	hexMinute := (secondsPastMidnight % 4096) / 16
	hexSecond := secondsPastMidnight % 16

	return fmt.Sprintf("%x_%02x_%x", hexHour, hexMinute, hexSecond)
}

// Hextime formats the current time into a hexadecimal notation.
func Hextime() string {
	return Hex(time.Now())
}
