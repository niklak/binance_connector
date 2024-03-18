package helpers

import (
	"strings"
	"time"
)

func CurrentTimestamp() int64 {
	return FormatTimestamp(time.Now())
}

// FormatTimestamp formats a time into Unix timestamp in milliseconds, as requested by Binance.
func FormatTimestamp(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}

func StringifyStringSlice(s []string) string {
	buf := strings.Builder{}
	buf.WriteString("[")
	for i, v := range s {
		if i != 0 {
			buf.WriteString(`,`)
		}
		buf.WriteString(`"`)
		buf.WriteString(v)
		buf.WriteString(`"`)
	}
	buf.WriteString("]")

	return buf.String()
}
