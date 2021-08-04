package main

import (
	"fmt"
	"reflect"
)

type UserInfo struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func HandleStructToMap(in interface{}, tagName string) (map[string]interface{}, error) {
	out := make(map[string]interface{})

	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("only accepts struct or struct pointer but got %T", v)
	}

	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		fi := t.Field(i)
		if tagValue := fi.Tag.Get(tagName); tagValue != "" {
			out[tagValue] = v.Field(i).Interface()
		}
	}
	return out, nil
}

func main() {
  u1 := UserInfo{Name: "leesanel", Age: 18}
	m2, _ := HandleStructToMap(&u1, "json")
	for k, v := range m2 {
		fmt.Printf("key:%v || value:%v || type:%T\n", k, v, v)
	}
}
