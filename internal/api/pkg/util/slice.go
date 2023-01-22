package util

import "reflect"

func ContainsType(target interface{}, slice interface{}) bool {
	valueOfSlice := reflect.ValueOf(slice)

	for i := 0; i < valueOfSlice.Len(); i++ {
		switch target.(type) {
		case int:
			if target.(int) == valueOfSlice.Index(i).Interface().(int) {
				return true
			}
		case string:
			if target.(string) == valueOfSlice.Index(i).Interface().(string) {
				return true
			}
		}
	}
	return false
}

func UniqType(slice interface{}) bool {
	valueOfSlice := reflect.ValueOf(slice)
	if valueOfSlice.Len() == 0 {
		return true
	}

	switch slice.(type) {
	case []string:
		encountered := map[string]bool{}
		for i := 0; i < valueOfSlice.Len(); i++ {
			value := valueOfSlice.Index(i).Interface().(string)
			if !encountered[value] {
				encountered[value] = true
			} else {
				return false
			}
		}
		return true
	case []int:
		encountered := map[int]bool{}
		for i := 0; i < valueOfSlice.Len(); i++ {
			value := valueOfSlice.Index(i).Interface().(int)
			if !encountered[value] {
				encountered[value] = true
			} else {
				return false
			}
		}
		return true
	default:
		return false
	}
}
