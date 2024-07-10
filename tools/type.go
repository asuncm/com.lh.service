package tools

import (
	"errors"
	"reflect"
)

func IsIntEmtpy(num int) bool {
	return num == 0
}

func IsInt8Emtpy(num int8) bool {
	return num == 0
}

func IsInt16Emtpy(num int16) bool {
	return num == 0
}

func IsInt32Emtpy(num int32) bool {
	return num == 0
}

func IsInt64Emtpy(num int64) bool {
	return num == 0
}

func IsStringEmtpy(value string) bool {
	return value == ""
}

func ToInt64(val interface{}) (int64, error) {
	vType := reflect.TypeOf(val)
	if vType.Kind() == reflect.Ptr {
		vType = vType.Elem()
	}
	if vType.Kind() != reflect.Int64 {
		return 0, errors.New("value is not of type int64")
	} else {
		return reflect.ValueOf(val).Int(), nil
	}
}
