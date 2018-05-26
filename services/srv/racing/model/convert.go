package model

import (
	"reflect"
)

func genericMap(arr interface{}, mapFunc interface{}) interface{} {
	funcValue := reflect.ValueOf(mapFunc)
	arrValue := reflect.ValueOf(arr)

	// Retrieve the type, and check if it is one of the array or slice.
	arrType := arrValue.Type()
	arrElemType := arrType.Elem()
	if arrType.Kind() != reflect.Array && arrType.Kind() != reflect.Slice {
		panic("Array parameter's type is neither array nor slice.")
	}

	funcType := funcValue.Type()

	// Checking whether the second argument is function or not.
	// And also checking whether its signature is func ({type A}) {type B}.
	if funcType.Kind() != reflect.Func || funcType.NumIn() != 1 || funcType.NumOut() != 1 {
		panic("Second argument must be map function.")
	}

	// Checking whether element type is convertible to function's first argument's type.
	if !arrElemType.ConvertibleTo(funcType.In(0)) {
		panic("Map function's argument is not compatible with type of array.")
	}

	// Get slice type corresponding to function's return value's type.
	resultSliceType := reflect.SliceOf(funcType.Out(0))

	// MakeSlice takes a slice kind type, and makes a slice.
	resultSlice := reflect.MakeSlice(resultSliceType, 0, arrValue.Len())

	for i := 0; i < arrValue.Len(); i++ {
		resultSlice = reflect.Append(resultSlice, funcValue.Call([]reflect.Value{arrValue.Index(i)})[0])
	}

	// Convering resulting slice back to generic interface.
	return resultSlice.Interface()
}
