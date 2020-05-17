package main

import (
	"reflect"
)

func Walk(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x)

	for i := 0; i < val.NumField(); i += 1 {
		field := val.Field(i)
		fn(field.String())
	}
}
