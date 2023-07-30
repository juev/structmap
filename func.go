package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func StructToMap(str any) map[string]string {
	val := reflect.ValueOf(str)
	typ := reflect.TypeOf(str)
	result := make(map[string]string)
	for i := 0; i < typ.NumField(); i++ {
		name := strings.ToLower(typ.Field(i).Name)
		switch val.Field(i).Kind() {
		case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			result[name] = strconv.FormatInt(val.Field(i).Int(), 10)
		case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			result[name] = strconv.FormatUint(val.Field(i).Uint(), 10)
		case reflect.Float32:
			result[name] = strconv.FormatFloat(val.Field(i).Float(), 'f', -1, 32)
		case reflect.Float64:
			result[name] = strconv.FormatFloat(val.Field(i).Float(), 'f', -1, 64)
		case reflect.Bool:
			result[name] = strconv.FormatBool(val.Field(i).Bool())
		case reflect.String:
			result[name] = val.Field(i).String()
		}
	}
	return result
}

func MapToStruct(mmap map[string]string) any {
	value := reflect.ValueOf(mmap)
	result := reflect.New(reflect.TypeOf(Str{})).Elem()
	iter := value.MapRange()
	for iter.Next() {
		fmt.Printf("key: %v\nvalue: %v\n", iter.Key(), iter.Value())
		key := toTitle(iter.Key().String())
		switch result.FieldByName(key).Kind() {
		case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			value, _ := strconv.Atoi(iter.Value().String())
			result.FieldByName(key).SetInt(int64(value))
		case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			value, _ := strconv.Atoi(iter.Value().String())
			result.FieldByName(key).SetUint(uint64(value))
		case reflect.Float32, reflect.Float64:
			value, _ := strconv.ParseFloat(iter.Value().String(), 64)
			result.FieldByName(key).SetFloat(value)
		case reflect.Bool:
			value, _ := strconv.ParseBool(iter.Value().String())
			result.FieldByName(key).SetBool(value)
		case reflect.String:
			result.FieldByName(key).SetString(iter.Value().String())
		}

	}

	return result
}

func toTitle(str string) (result string) {
	if len(str) == 1 {
		return strings.ToUpper(string(str[0]))
	}

	if len(str) > 1 {
		return strings.ToUpper(string(str[0])) + str[1:]
	}

	return
}
