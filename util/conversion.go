package util

import (
	"encoding/json"
	"strings"
)

func UpateOneStructWithOtherWithSameKeys(updateValueStruct, targetValueStruct interface{}) (err error) {
	bValue, err := json.Marshal(updateValueStruct)
	if err != nil {
		return
	}
	err = json.Unmarshal(bValue, &targetValueStruct)
	return
}

func ConvertMapToStruct(obj map[string]interface{}, stru interface{}) (err error) {
	bData, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	return json.Unmarshal(bData, stru)
}

func AddKeyToObject(obj map[string]interface{}, key string, value interface{}) map[string]interface{} {
	output := map[string]interface{}{}
	for k, v := range obj {
		output[k] = v
	}
	output[key] = value
	return output
}

func ParseDotedKeyToNestedMap(obje interface{}) interface{} {
	obj := obje.(map[string]interface{})
	for k, v := range obj {
		keyArr := strings.Split(k, ".")
		if len(keyArr) > 1 {
			_, ok := obj[keyArr[0]]
			if ok {
				obj[keyArr[0]] = AddKeyToObject(obj[keyArr[0]].(map[string]interface{}), strings.Join(keyArr[1:], "."), v)
				delete(obj, k)
			} else {
				obj[keyArr[0]] = map[string]interface{}{
					strings.Join(keyArr[1:], "."): v,
				}
				delete(obj, k)
			}
		}

	}
	return obj
}
