package util

import (
	"github.com/fatih/structs"
	"github.com/satori/go.uuid"
	"strconv"
)

// UUID 获取UUID
func UUID() string {
	return uuid.Must(uuid.NewV4(), nil).String()
}

// StructToMap 转换struct为map
func StructToMap(s interface{}) map[string]interface{} {
	return structs.Map(s)
}

// StringToInt 字符串转数值
func StringToInt(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}
