package reflection

import (
	"fmt"
	"reflect"
	"strconv"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Clothes struct {
		Color string `json:"color"`
	}
}

func RecursivePrintlnTagName(in interface{}) {
	rt := reflect.TypeOf(in)
	rv := reflect.ValueOf(in)
	if rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
		rv = rv.Elem()
	}
	for i := 0; i < rt.NumField(); i++ {
		filed := rt.Field(i)
		if filed.Type.Kind() == reflect.Ptr {
			RecursivePrintlnTagName(filed.Type.Elem())
			continue
		}
		if filed.Type.Kind() == reflect.Struct {
			RecursivePrintlnTagName(filed.Type)
			continue
		}
		jsonTagName, ok := rt.Field(i).Tag.Lookup("json")
		if ok {
			fmt.Println(jsonTagName)
		}
	}
}

func Display(name string, x interface{}) {
	fmt.Printf("Display %s (%T):\n", name, x)
	display(name, reflect.ValueOf(x))
}

func display(path string, v reflect.Value) {
	switch v.Kind() {
	case reflect.Invalid:
		fmt.Printf("%s = invalid\n", path)
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			display(fmt.Sprintf("%s[%d]", path, i), v.Index(i))
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			fieldPath := fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name)
			display(fieldPath, v.Field(i))
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			display(fmt.Sprintf("%s[%s]", path,
				strconv.FormatUint(uint64(v.MapIndex(key).Pointer()), 16)),
				v.MapIndex(key))
		}
	case reflect.Ptr:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			display(fmt.Sprintf("(*%s)", path), v.Elem())
		}
	case reflect.Interface:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			fmt.Printf("%s.type = %s\n", path, v.Elem().Type())
			display(path+".value", v.Elem())
		}
	default:
		fmt.Printf("%s = %s\n", path,
			strconv.FormatUint(uint64(v.Pointer()), 16))
	}
}
