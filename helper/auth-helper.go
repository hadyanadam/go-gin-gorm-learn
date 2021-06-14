package helper

import (
	"fmt"
	"strings"
)


func SplitHeader(authHeader string, authType string) (string, error) {
	splitedHeader := strings.SplitAfter(authHeader, " ")
	if len(splitedHeader) == 0 {
		return "", fmt.Errorf("Cannot find authorization header")
	}
	if splitedHeader[0] != (authType + " ") {
		return "", fmt.Errorf("authType must be %s", authType)
	}
	return splitedHeader[1], nil
}