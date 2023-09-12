package utils

import "strconv"

// ToInt Convert a String to an int, returning a default value if the conversion fails.
func ToInt(str string, defaultValue int) int {
	if str == "" {
		return defaultValue
	}
	i, err := strconv.Atoi(str)
	if err != nil {
		return defaultValue
	}
	return i
}

func ToUInt32(str string, defaultValue uint32) uint32 {
	if str == "" {
		return defaultValue
	}
	u, err := strconv.ParseUint(str, 10, 32)
	if err != nil {
		return defaultValue
	}
	return uint32(u)
}

// ToInt64 Convert a String to an int64, returning a default value if the conversion fails.
func ToInt64(str string, defaultValue int64) int64 {
	if str == "" {
		return defaultValue
	}
	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return defaultValue
	}
	return i
}

// ToFloat32 Convert a String to an float32, returning a default value if the conversion fails.
func ToFloat32(str string, defaultValue float32) float32 {
	if str == "" {
		return defaultValue
	}
	f, err := strconv.ParseFloat(str, 32)
	if err != nil {
		return defaultValue
	}
	return float32(f)
}

// ToFloat64 Convert a String to an float64, returning a default value if the conversion fails.
func ToFloat64(str string, defaultValue float64) float64 {
	if str == "" {
		return defaultValue
	}
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return defaultValue
	}
	return f
}
