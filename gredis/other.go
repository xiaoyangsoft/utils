package gredis

import (
	"github.com/mitchellh/mapstructure"
	"reflect"
)

// StructToMapJson
/**
 *  @Description: 结构体转Json
 *  @param obj
 *  @return map[string]interface{}
 */
func StructToMapJson(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		jsonKey := t.Field(i).Tag.Get("json")
		if jsonKey != "-" {
			data[jsonKey] = v.Field(i).Interface()
		}
	}
	return data
}

// MapToStruct
/**
 *  @Description: Map转结构体
 *  @param m
 *  @param data
 *  @return err
 */
func MapToStruct(m map[string]interface{}, data interface{}) (err error) {
	err = mapstructure.Decode(m, &data)
	return
}
