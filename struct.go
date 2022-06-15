/*
Package snamer renames fields of a struct.
It is useful to create snake_case json from a struct.
In Golang, an initial character of a field must be uppercase to export the field.
*/
package snamer

import (
	"errors"
	"fmt"
	"reflect"
)

// convert field names recursively
func convertFields(v reflect.Value, fieldConversionFunc func(s string) string) (interface{}, error) {
	// https://qiita.com/nirasan/items/b6b89f8c61c35b563e8c
	// https://qiita.com/tsubaki_dev/items/a8ffd28d4513e8750355
	if !v.CanInterface() {
		return nil, nil
	}

	if v.Kind() == reflect.Ptr {
		// pointer dereference (Pointer of Struct to Struct)
		v = v.Elem()
	}

	// TODO: PtrとInterfaceの順番これで大丈夫か？ Pointer of Interfaceは？
	if v.Kind() == reflect.Interface {
		// interface to concrete type
		v = v.Elem()
	}

	switch v.Kind() {
	case reflect.Struct:
		result := make(map[string]interface{})
		for i := 0; i < v.NumField(); i++ {
			val := v.Field(i)
			key := v.Type().Field(i).Name
			val2, err := convertFields(val, fieldConversionFunc)
			if err != nil {
				return nil, err
			}
			result[fieldConversionFunc(key)] = val2
		}
		// fmt.Printf("%#v\n", result)
		return result, nil
	case reflect.Array, reflect.Slice:
		result := make([]interface{}, v.Len())
		for i := 0; i < v.Len(); i++ {
			val2, err := convertFields(v.Index(i), fieldConversionFunc)
			if err != nil {
				return nil, err
			}
			result[i] = val2
		}
		return result, nil
	case reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr, reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128:
		// primitive type
		return v.Interface(), nil
	case reflect.Map:
		result := make(map[string]interface{})
		for _, k := range v.MapKeys() {
			val2, err := convertFields(v.MapIndex(k), fieldConversionFunc)
			if err != nil {
				return nil, err
			}
			result[fieldConversionFunc(fmt.Sprintf("%s", k.Interface()))] = val2
		}
		return result, nil
	case reflect.String:
		return v.Interface(), nil
	default:
		return nil, errors.New("method convertKeys is not implemented for kind: " + v.Kind().String())
	}
}

// Convert a struct that has PascalCase fields to a complex of Slice and map[string]interface{} that has camelCase fields
func PascalStructToCamel(input interface{}) (interface{}, error) {
	return AnyStructToAny(input, pascalStringToCamel)
}

// Convert a struct that has PascalCase fields to a complex of Slice and map[string]interface{} that has camelCase fields
func PascalStructToSnake(input interface{}) (interface{}, error) {
	return AnyStructToAny(input, pascalStringToSnake)
}

// generalized version of PascalStructToCamel.
func AnyStructToAny(input interface{}, fieldConversionFunc func(s string) string) (interface{}, error) {
	// camelCase
	v := reflect.ValueOf(input)

	result, err := convertFields(v, fieldConversionFunc)
	if err != nil {
		return nil, err
	}
	// fmt.Printf("%#v\n", v)
	return result, nil
}

// TODO: Chain, Kebab case and general key conversion function
