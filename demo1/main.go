package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string `json:"name" label:"姓名"`
	Age  int    `json:"age"`
}

func main() {
	t := reflect.TypeOf(Person{})

	field := t.Field(0)
	fmt.Println(field.Tag.Get("json"))
	fmt.Println(field.Tag.Get("label"))

}
