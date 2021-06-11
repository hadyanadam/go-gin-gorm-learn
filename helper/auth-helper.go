package helper

import (
	"fmt"
	"strings"
)


func SplitHeader(authHeader string, authType string) (string, error) {
	splitedHeader := strings.SplitAfter(authHeader, " ")
	fmt.Println(splitedHeader)
	if splitedHeader[0] == authType {
		return splitedHeader[1], fmt.Errorf("authType must be %s", authType)
	}
	return splitedHeader[1], nil
}