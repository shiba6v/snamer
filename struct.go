package snamer

import (
	"fmt"
	"reflect"
)

func innerPascalStructToCamel(v reflect.Value) interface{} {
	// https://qiita.com/nirasan/items/b6b89f8c61c35b563e8c
	// https://qiita.com/tsubaki_dev/items/a8ffd28d4513e8750355

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
			result[pascalStringToCamel(key)] = innerPascalStructToCamel(val)
		}
		// fmt.Printf("%#v\n", result)
		return result
	case reflect.Array, reflect.Slice:
		result := make([]interface{}, v.Len())
		for i := 0; i < v.Len(); i++ {
			result[i] = innerPascalStructToCamel(v.Index(i))
		}
		return result
	case reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr, reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128:
		// primitive type
		return v.Interface()
	case reflect.Map:
		result := make(map[string]interface{})
		for _, k := range v.MapKeys() {
			result[pascalStringToCamel(fmt.Sprintf("%s", k.Interface()))] = innerPascalStructToCamel(v.MapIndex(k))
		}
		return result
	case reflect.String:
		return v.Interface()
	default:
		return nil
	}
}

// a
func PascalStructToCamel(input interface{}) (interface{}, error) {
	// camelCase
	v := reflect.ValueOf(input)

	result := innerPascalStructToCamel(v)
	// fmt.Printf("%#v\n", v)
	return result, nil
}

// TODO: Snake, Chain, Kebab case and general key conversion function
