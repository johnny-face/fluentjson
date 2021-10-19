package fluentjson

import "encoding/json"

type Array []interface{}

func NewArray() *Array {
	return &Array{}
}

func (arr *Array) Add(v interface{}) *Array {
	*arr = append(*arr, v)
	return arr
}

func (arr *Array) AddNil() *Array {
	var v interface{}
	*arr = append(*arr, v)
	return arr
}

func (arr *Array) Set(idx int, v interface{}) *Array {
	//check idx
	return arr
}

func (arr *Array) SetNil(idx int) *Array {
	//check idx
	return arr
}

func (arr *Array) Remove(idx int) *Array {
	//check idx
	return arr
}

func (arr *Array) Copy() *Array {
	ret := Array{}
	copy(ret, *arr)
	return &ret
}

func (arr *Array) getValue(idx int) (interface{}, error) {
	return arr, nil
}

func (arr *Array) getString(idx int) (interface{}, error) {
	return arr, nil
}

func (arr *Array) getInt(idx int) (interface{}, error) {
	return arr, nil
}

func (arr *Array) getFloat(idx int) (interface{}, error) {
	return arr, nil
}

func (arr *Array) getDouble(idx int) (interface{}, error) {
	return arr, nil
}

func (arr *Array) getBool(idx int) (interface{}, error) {
	return arr, nil
}

func (arr *Array) getTime(idx int) (interface{}, error) {
	return arr, nil
}

func (arr *Array) getObject(idx int) (interface{}, error) {
	return arr, nil
}

func (arr *Array) getArray(idx int) (interface{}, error) {
	return arr, nil
}

func (arr *Array) Size() int {
	return len(*arr)
}

func (arr *Array) IsEmpty() bool {
	return arr.Size() == 0
}

func (arr *Array) Values() []interface{} {
	return *arr
}

func (arr *Array) String() string {
	b, _ := json.Marshal(*arr)
	return string(b)
}