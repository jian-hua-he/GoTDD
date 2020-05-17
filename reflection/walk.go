package main

import (
	"reflect"
)

func Walk(x interface{}, fn func(input string)) {
	val := getValue(x)
	numOfVal := 0
	var getField func(int) reflect.Value

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		numOfVal = val.NumField()
		getField = val.Field
	case reflect.Slice, reflect.Array:
		numOfVal = val.Len()
		getField = val.Index
	}

	for i := 0; i < numOfVal; i += 1 {
		Walk(getField(i).Interface(), fn)
	}
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	return val
}
