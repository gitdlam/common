package common

import (
	"regexp"
)

func ValidBarcode8(s string) bool {
	re := regexp.MustCompile("^[A-Z0-9]{8,8}$")
	return re.MatchString(s)
}

func ValidBarcode7(s string) bool {
	re := regexp.MustCompile("^[A-Z0-9]{7,7}$")
	return re.MatchString(s)
}

func ValidCarton20(s string) bool {
	re := regexp.MustCompile("^[0-9]{20,20}$")
	return re.MatchString(s)
}
