package model

import "strings"

type Payload struct {
	Status string `json:"status"`
	Data   any    `json:"data"`
	Error  any    `json:"error"`
}

func IsPayloadMsg(s string) bool {
	for _, ch := range s {
		if (ch >= 'a' && ch <= 'z') || ch == ' ' {
			return false
		}
	}
	return true
}

func ToPayloadStatus(s string) string {
	return strings.ReplaceAll(strings.ToUpper(s), " ", "_")
}
