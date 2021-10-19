package fluentjson

import "testing"

func TestObject_Size(t *testing.T) {
	obj := Object{
		"first": "first",
	}
	if obj.Size() != 1 {
		t.Errorf("except size %d", obj.Size())
	}

	obj.PutValue("second", "second")
	if obj.Size() != 2 {
		t.Errorf("except size %d", obj.Size())
	}

	obj.PutValue("second", "second-1")
	if obj.Size() != 2 {
		t.Errorf("except size %d", obj.Size())
	}

	obj.Remove("second")
	if obj.Size() != 1 {
		t.Errorf("except size %d", obj.Size())
	}

}
