package utils

import (
	"regexp"
)

func ValidateEmail(mail string) bool {
	regex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return regex.MatchString(mail)
}

func CheckEmpty(data interface{}) bool {
	if data == nil {
		return true
	}

	switch v := data.(type) {
	case string:
		return v == ""
	case int:
		return v == 0
	case float64:
		return v == 0.0
	case bool:
		return !v
	case []interface{}:
		return len(v) == 0
	default:
		return false
	}
}
