package string_util

import (
	"bytes"
	"encoding/json"
)

// MustFormatJson 格式化为 json，失败 panic。
func MustFormatJson(jsonBytes []byte) []byte {
	bts, err := FormatJson(jsonBytes)
	if err != nil {
		panic(err)
	}
	return bts
}

// FormatJson 格式化为 json。
func FormatJson(jsonBytes []byte) ([]byte, error) {
	var buffer bytes.Buffer
	if err := json.Indent(&buffer, jsonBytes, "", "\t"); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}
