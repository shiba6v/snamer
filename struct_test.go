package snamer_test

import (
	"fmt"
	"reflect"
	"testing"

	"shiba6v.com/snamer"
)

type TestPascalStructToCamelExample struct {
	Name     string
	Input    interface{}
	Expected interface{}
}

type User struct {
	UserId int
}

type Post1 struct {
	Text string
	User User
}

type Post2 struct {
	Text string
	User *User
}

type Primitive1 struct {
	Uint uint
}
type Primitive2 struct {
	Bool   bool
	Int    int
	Int8   int8
	Int16  int16
	Int32  int32
	Int64  int64
	Uint   uint
	Uint8  uint8
	Uint16 uint16
	Uint32 uint32
	Uint64 uint64
	// Uintptr uintptr
	Float32    float32
	Float64    float64
	Complex64  complex64
	Complex128 complex128
}

func TestPascalStructToCamel(t *testing.T) {
	fmt.Print("TestPascalStructToCamel\n")
	examples := []TestPascalStructToCamelExample{
		{
			Name:  "1: basic",
			Input: User{1},
			Expected: map[string]interface{}{
				"userId": 1,
			},
		},
		{
			Name:  "2: pointer input",
			Input: &User{1},
			Expected: map[string]interface{}{
				"userId": 1,
			},
		},
		{
			Name:  "3: nested",
			Input: Post1{"a", User{1}},
			Expected: map[string]interface{}{
				"text": "a",
				"user": map[string]interface{}{
					"userId": 1,
				},
			},
		},
		{
			Name:  "4: nested pointer",
			Input: Post2{"a", &User{1}},
			Expected: map[string]interface{}{
				"text": "a",
				"user": map[string]interface{}{
					"userId": 1,
				},
			},
		},
		{
			Name:     "5: primitive1",
			Input:    Primitive1{Uint: 0},
			Expected: map[string]interface{}{"uint": uint(0)},
		},
		{
			Name:     "6: primitive2",
			Input:    Primitive2{Bool: false, Complex128: (0 + 0i), Complex64: (0 + 0i), Float32: 0.0, Float64: 0.0, Int: 0, Int16: 0, Int32: 0, Int64: 0, Int8: 0, Uint: 0, Uint16: 0, Uint32: 0, Uint64: 0, Uint8: 0},
			Expected: map[string]interface{}{"bool": false, "complex128": complex128(0 + 0i), "complex64": complex64(0 + 0i), "float32": float32(0), "float64": float64(0), "int": int(0), "int16": int16(0), "int32": int32(0), "int64": int64(0), "int8": int8(0), "uint": uint(0), "uint16": uint16(0), "uint32": uint32(0), "uint64": uint64(0), "uint8": uint8(0)},
		},
		{
			Name:     "7: map",
			Input:    map[string]interface{}{"u": 1},
			Expected: map[string]interface{}{"u": 1},
		},
		{
			Name:     "8: array",
			Input:    [2]User{{1}, {2}},
			Expected: []interface{}{map[string]interface{}{"userId": 1}, map[string]interface{}{"userId": 2}},
		},
		{
			Name:     "9: array",
			Input:    map[string]interface{}{"u": [2]User{{1}, {2}}},
			Expected: map[string]interface{}{"u": []interface{}{map[string]interface{}{"userId": 1}, map[string]interface{}{"userId": 2}}},
		},
	}
	for _, ex := range examples {
		t.Run(
			"TestPascalStructToCamel_"+ex.Name, func(t *testing.T) {

				inputType := reflect.TypeOf(ex.Input)

				v, _ := snamer.PascalStructToCamel(ex.Input)

				// 入力がpointerの場合にpointerから勝手に書き変わらないことを一応チェック
				if inputType != reflect.TypeOf(ex.Input) {
					t.Errorf("Input Type changed by side effect")
				}

				if !reflect.DeepEqual(v, ex.Expected) {
					t.Errorf("Error: input:%#v\n  result  : %#v,\n  expected: %#v", ex.Input, v, ex.Expected)
				}
			})
	}
}

func ExamplePascalStructToCamel() {
	user := User{UserId: 1}
	result, err := snamer.PascalStructToCamel(user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result) // map[string]interface{}{"userId":1}
}
