package gominiprop

import (
	"testing"
)

func TestKey(t *testing.T) {
	propObj := NewMiniProp()
	propObj.SetPropInt("test", "a", 3)
	propObj.SetPropString("test", "c", "test")
	if 1 != propObj.GetPropInt("d", "b", 1) {
		t.Error("s")
	}
	propObj = NewMiniPropFromJson(propObj.ToJson())
	if 1 != propObj.GetPropInt("d", "b", 1) {
		t.Error("s")
	}

}

func TestInt(t *testing.T) {
	propObj := NewMiniProp()
	propObj.SetPropInt("test", "a", 3)
	propObj.SetPropString("test", "c", "test")
	if 3 != propObj.GetPropInt("test", "a", 0) {
		t.Error("s")
	}
	if 1 != propObj.GetPropInt("test", "b", 1) {
		t.Error("s")
	}
	if 1 != propObj.GetPropInt("test", "c", 1) {
		t.Error("s")
	}
	//
	//
	propObj = NewMiniPropFromJson(propObj.ToJson())
	if 3 != propObj.GetPropInt("test", "a", 0) {
		t.Error("s")
	}
	if 1 != propObj.GetPropInt("test", "b", 1) {
		t.Error("s")
	}
	if 1 != propObj.GetPropInt("test", "c", 1) {
		t.Error("s")
	}
}

func TestString(t *testing.T) {
	propObj := NewMiniProp()
	propObj.SetPropString("test", "a", "3")
	propObj.SetPropInt("test", "c", 1)
	if "3" != propObj.GetPropString("test", "a", "0") {
		t.Error("s")
	}
	if "1" != propObj.GetPropString("test", "b", "1") {
		t.Error("s")
	}
	if "1" != propObj.GetPropString("test", "c", "1") {
		t.Error("s")
	}
	//
	//
	propObj = NewMiniPropFromJson(propObj.ToJson())
	if "3" != propObj.GetPropString("test", "a", "0") {
		t.Error("s")
	}
	if "1" != propObj.GetPropString("test", "b", "1") {
		t.Error("s")
	}
	if "1" != propObj.GetPropString("test", "c", "1") {
		t.Error("s")
	}

}
