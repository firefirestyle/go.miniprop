package miniprop

import (
	"bytes"
	"testing"
)

func TestKey(t *testing.T) {
	propObj := NewMiniProp()
	propObj.SetPropInt("test", "a", 3)
	propObj.SetPropString("test", "c", "test")
	if 1 != propObj.GetPropInt("d", "b", 1) {
		t.Error("s")
	}
	//
	//
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

func TestFloat(t *testing.T) {
	propObj := NewMiniProp()
	propObj.SetPropFloat("test", "a", 3.0)
	propObj.SetPropString("test", "c", "test")
	if 3.0 != propObj.GetPropFloat64("test", "a", 0) {
		t.Error("s1")
	}
	if 1 != propObj.GetPropInt("test", "b", 1) {
		t.Error("s2")
	}
	if 1 != propObj.GetPropInt("test", "c", 1) {
		t.Error("s3")
	}
	//
	//
	propObj = NewMiniPropFromJson(propObj.ToJson())
	if 3.0 != propObj.GetPropFloat64("test", "a", 0) {
		t.Error("s1")
	}
	if 1 != propObj.GetPropInt("test", "b", 1) {
		t.Error("s5")
	}
	if 1 != propObj.GetPropInt("test", "c", 1) {
		t.Error("s6z")
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

func TestBytes(t *testing.T) {
	propObj := NewMiniProp()
	propObj.SetPropBytes("test", "a", []byte{1, 2, 3})
	propObj.SetPropInt("test", "c", 1)

	if 0 != bytes.Compare([]byte{1, 2, 3}, propObj.GetPropBytes("test", "a", []byte{1, 2})) {
		t.Error("TestBytes1")
	}
	if 0 != bytes.Compare([]byte{1, 2}, propObj.GetPropBytes("test", "b", []byte{1, 2})) {
		t.Error("TestBytes2")
	}
	if 0 != bytes.Compare([]byte{1, 2}, propObj.GetPropBytes("test", "c", []byte{1, 2})) {
		t.Error("TestBytes3")
	}
	//
	//
	propObj = NewMiniPropFromJson(propObj.ToJson())
	if 0 != bytes.Compare([]byte{1, 2, 3}, propObj.GetPropBytes("test", "a", []byte{1, 2})) {
		t.Error("TestBytes1")
	}
	if 0 != bytes.Compare([]byte{1, 2}, propObj.GetPropBytes("test", "b", []byte{1, 2})) {
		t.Error("TestBytes2")
	}
	if 0 != bytes.Compare([]byte{1, 2}, propObj.GetPropBytes("test", "c", []byte{1, 2})) {
		t.Error("TestBytes3")
	}
}
