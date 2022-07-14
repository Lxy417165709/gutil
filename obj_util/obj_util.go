package obj_util

import (
	"encoding/json"

	"github.com/Lxy417165709/gutil/string_util"
)

// MustFormatObjectJson 格式化展示结构体，失败 panic。
func MustFormatObjectJson(obj interface{}) []byte {
	data, err := FormatObjectJson(obj)
	if err != nil {
		panic(err)
	}
	return string_util.MustFormatJson(data)
}

// FormatObjectJson 格式化展示结构体。
func FormatObjectJson(obj interface{}) ([]byte, error) {
	data, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	return string_util.FormatJson(data)
}
