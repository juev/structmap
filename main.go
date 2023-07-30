package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type Str struct {
	Title   string
	Body    string
	Header  string
	Numbers int16
	Fl      float32
	B       bool
}

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	str := Str{
		Title:   "title_content",
		Body:    "body_content",
		Header:  "header_content",
		Numbers: 12,
		Fl:      12.12,
		B:       true,
	}
	fmt.Printf("struct: %#v\n", str)
	m := StructToMap(str)
	fmt.Printf("map: %#v\n", m)

	fmt.Println("MapToStruct")
	s := MapToStruct(m)
	fmt.Printf("struct: %#v\n", s)
	return nil
}

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
	result := Str{}
	// m := reflect.ValueOf(mmap)
	r := reflect.ValueOf(result)
	fmt.Println(r.FieldByName("Title").CanSet())
	// r.FieldByName("Title").SetString("new title")
	// iter := m.MapRange()
	// for iter.Next() {
	// 	fmt.Printf("key: %v\nvalue: %v\n", iter.Key(), iter.Value())
	// 	fmt.Println(cases.Title(iter.Key().String()))

	// 	r.FieldByName(cases.Title(iter.Key().String())).SetString(iter.Value().String())
	// }

	return result
}
