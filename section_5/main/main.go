package main

import (
	"fmt"
	"reflect"
)

func main() {
	s := "hello world"
	Print(s)
}

func Print(i interface{}) {
	validateInterface(i)
	t := reflect.TypeOf(i)
	if t.Kind() == reflect.Struct {
		printStructType(t)
	} else {
		printNonStructType(t)
	}
}

func validateInterface(i interface{}) {
	iv := reflect.ValueOf(i)
	if !iv.IsValid() {
		fmt.Errorf("given interface is invalid")
	}
	if iv.IsNil() {
		fmt.Errorf("given interface is nil")
	}
}

func printStructType(t reflect.Type) {
	fmt.Println("Type:{", t)
	for i := 0; i < t.NumField(); i++ {
		fmt.Print("{\n")
		field := t.Field(i)
		fmt.Printf("Name: %+v\n", field.Name)
		fmt.Printf("Type: %+v\n", field.Type)
		v := reflect.ValueOf(field)
		printValues(v)
		fmt.Print("}\n")
	}
	fmt.Println("}")
}

func printNonStructType(t reflect.Type) {
	fmt.Println("Type:", t)
	v := reflect.ValueOf(t)
	printValues(v)
}

func printValues(v reflect.Value) {
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	for i := 0; i < v.NumField(); i++ {
		fmt.Printf("Value: %+v\n", v.Field(i))
	}
}
