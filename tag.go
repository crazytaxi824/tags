package tags

import (
	"errors"
	"reflect"
	"strings"
)

// GetTag 映射结构体 tag 字段
func GetTag(s interface{}, fromTag, toTag string, filterFields []string) ([]string, error) {
	typeof := reflect.TypeOf(s)
	numberField := typeof.NumField()
	if numberField == 0 {
		return nil, errors.New("模型没有fields")
	}

	var resultSlice []string

	for _, v := range filterFields {
		match := false

		for i := 0; i < numberField; i++ {
			field := typeof.Field(i)

			fromVSlice := strings.Split(field.Tag.Get(fromTag), ",")
			fromValue := strings.TrimSpace(fromVSlice[0])

			toVSlice := strings.Split(field.Tag.Get(toTag), ",")
			toValue := strings.TrimSpace(toVSlice[0])

			if v == fromValue {
				resultSlice = append(resultSlice, toValue)
				match = true
			}
		}
		if !match {
			return nil, errors.New(v + " 不存在")
		}
	}

	return resultSlice, nil
}
