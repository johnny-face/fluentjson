package fluentjson

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

type Object map[string]interface{}

func (obj Object) Copy() Object {
	ret := Object{}
	for k, v := range obj {
		ret[k] = v
	}
	return ret
}

func (obj Object) PutValue(k string, v interface{}) Object {
	// change struct to map
	obj[k] = v
	return obj
}

func (obj Object) Remove(k string) Object {
	delete(obj, k)
	return obj
}

func (obj Object) getValue(k string) (interface{}, error) {
	v, ok := obj[k]
	if !ok {
		return nil, nil
	}
	return v, nil
}

func (obj Object) getString(k string) (interface{}, error) {
	return obj, nil
}

func (obj Object) getInt(k string) (interface{}, error) {
	return obj, nil
}

func (obj Object) getFloat(k string) (interface{}, error) {
	return obj, nil
}

func (obj Object) getDouble(k string) (interface{}, error) {
	return obj, nil
}

func (obj Object) getBool(k string) (interface{}, error) {
	return obj, nil
}

func (obj Object) getTime(k string) (interface{}, error) {
	return obj, nil
}

func (obj Object) getObject(k string) (Object, error) {
	v, ok := obj[k]
	if !ok {
		return Object{}, nil
	}
	switch v.(type) {
	case Object:
		return v.(Object), nil
	case map[string]interface{}:
		return v.(map[string]interface{}), nil
	}
	return nil, errors.New(fmt.Sprintf("Casting error. Interface is %s, not fluntjson.object", reflect.TypeOf(v)))
}

func (obj Object) getArray(k string) (interface{}, error) {
	return obj, nil
}

func (obj Object) Size() int {
	return len(obj)
}

func (obj Object) IsEmpty() bool {
	return obj.Size() == 0
}

func (obj Object) Contains(k string) bool {
	_, ok := obj[k]
	return ok
}

func (obj Object) Values() []interface{} {
	return nil
}

func (obj Object) String() string {
	b, _ := json.Marshal(obj)
	return string(b)
}
