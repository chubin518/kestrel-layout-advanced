package utils

import "time"

const (
	// yyyy-MM-dd
	PATTERN_NORM_DATE = "2006-01-02"
	// yyyyMMdd
	PATTERN_PURE_DATE = "20060102"
	// HH:mm:ss
	PATTERN_NORM_TIME = "15:04:05"
	// HHmmss
	PATTERN_PURE_TIME = "150405"
	// yyyy-MM-dd HH:mm:ss
	PATTERN_NORM_DATETIME = "2006-01-02 15:04:05"
	// yyyyMMddHHmmss
	PATTERN_PURE_DATETIME = "20060102150405"
	// yyyy-MM-dd HH:mm:ss.SSS
	PATTERN_NORM_DATETIME_MS = "2006-01-02 15:04:05.999"
	// yyyyMMddHHmmss.SSS
	PATTERN_PURE_DATETIME_MS = "20060102150405.999"
)

var (
	// NORM_DATE format yyyy-MM-dd
	NORM_DATE = NewFormat(PATTERN_NORM_DATE)
	// PURE_DATE format yyyyMMdd
	PURE_DATE = NewFormat(PATTERN_PURE_DATE)
	// NORM_TIME format HH:mm:ss
	NORM_TIME = NewFormat(PATTERN_NORM_TIME)
	// PURE_TIME format HHmmss
	PURE_TIME = NewFormat(PATTERN_PURE_TIME)
	// NORM_DATETIME format yyyy-MM-dd HH:mm:ss
	NORM_DATETIME = NewFormat(PATTERN_NORM_DATETIME)
	// PURE_DATETIME format yyyyMMddHHmmss
	PURE_DATETIME = NewFormat(PATTERN_PURE_DATETIME)
	// NORM_DATETIME_MS format yyyy-MM-dd HH:mm:ss.SSS
	NORM_DATETIME_MS = NewFormat(PATTERN_NORM_DATETIME_MS)
	// PURE_DATETIME_MS format yyyyMMddHHmmss.SSS
	PURE_DATETIME_MS = NewFormat(PATTERN_PURE_DATETIME_MS)
)

type TimeFormat struct {
	pattern string
}

func NewFormat(pattern string) *TimeFormat {
	return &TimeFormat{
		pattern: pattern,
	}
}

// 格式化
func (tf *TimeFormat) String(t time.Time) string {
	if tf.pattern == "" {
		return t.String()
	}
	return t.Format(tf.pattern)
}
