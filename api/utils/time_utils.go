package utils

import "time"

const (
	// TimeFormatterSECOND TODO
	TimeFormatterSECOND string = "2006-01-02 15:04:05"
	// TimeFormatterMINUTE TODO
	TimeFormatterMINUTE string = "2006-01-02 15:04"
	// TimeFormatterHOUR TODO
	TimeFormatterHOUR string = "2006-01-02 15"
	// TimeFormatterDAY TODO
	TimeFormatterDAY string = "2006-01-02"
	// TimeFormatterMONTH TODO
	TimeFormatterMONTH string = "2006-01"
	// TimeFormatterYEAR TODO
	TimeFormatterYEAR string = "2006"
)

// SecondFormatToString TODO
func SecondFormatToString(second int64, format string) string {
	t := time.Unix(second, 0)
	return t.Format(format)
}

// MillisecondFormatToString TODO
func MillisecondFormatToString(milli int64, format string) string {
	t := time.Unix(milli/1000, 0)
	return t.Format(format)
}

// MicrosecondFormatToString TODO
func MicrosecondFormatToString(micro int64, format string) string {
	t := time.Unix(micro/1000000, 0)
	return t.Format(format)
}

// NanosecondFormatToString TODO
func NanosecondFormatToString(nano int64, format string) string {
	t := time.Unix(nano/1000000000, 0)
	return t.Format(format)
}
