package utils

import (
	"encoding/json"
	"fmt"
)

func Struct2StructByJson(s1 interface{}, s2 interface{}) {
	jsonStr := Struct2Json(s1)
	Json2Struct(jsonStr, s2)
}

func Struct2Json(s interface{}) string {
	str, err := json.Marshal(s)
	if err != nil {
		fmt.Printf("对象无法转换为JSON：%s\n", err.Error())
	}
	return string(str)
}

func Json2Struct(str string, s interface{}) {
	err := json.Unmarshal([]byte(str), s)
	if err != nil {
		fmt.Printf("JSON无法转换为对象：%s\n", err.Error())
	}
}
