package main

import (
	"fmt"

	"github.com/juev/structmap"
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
	m := structmap.StructToMap(str)
	fmt.Printf("map: %#v\n", m)

	s := structmap.MapToStruct(m, Str{})
	fmt.Printf("struct: %#v\n", s)
	return nil
}
