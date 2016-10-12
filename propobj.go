package gominiprop

import "encoding/json"

import "encoding/base64"

//import "io"
//import "io/ioutil"

type MiniProp struct {
	prop map[string]interface{}
}

func NewMiniPropFromJson(source []byte) *MiniProp {
	propObj := new(MiniProp)
	propObj.prop = make(map[string]interface{})
	json.Unmarshal(source, &propObj.prop)
	return propObj
}

func NewMiniProp() *MiniProp {
	propObj := new(MiniProp)
	propObj.prop = make(map[string]interface{})
	return propObj
}

func (obj *MiniProp) GetProps(category string, defaultValue interface{}) interface{} {
	v := obj.prop[category]
	if v == nil {
		return defaultValue
	} else {
		return v
	}
}

func (obj *MiniProp) GetProp(category string, key string, defaultValue interface{}) interface{} {
	v := obj.prop[category]
	if v == nil {
		return defaultValue
	}
	if obj.prop[category].(map[string]interface{})[key] == nil {
		return defaultValue
	}
	return obj.prop[category].(map[string]interface{})[key]
}

func (obj *MiniProp) GetPropInt(category string, key string, defaultValue int) int {
	v := obj.GetProp(category, key, defaultValue)
	switch v.(type) {
	case int:
		return obj.GetProp(category, key, defaultValue).(int)
	case float64:
		return int(obj.GetProp(category, key, defaultValue).(float64))
	}
	return defaultValue
}
func (obj *MiniProp) GetPropFloat64(category string, key string, defaultValue float64) float64 {
	v := obj.GetProp(category, key, defaultValue)
	switch v.(type) {
	case float64:
		return obj.GetProp(category, key, defaultValue).(float64)
	}
	return defaultValue
}

func (obj *MiniProp) GetPropString(category string, key string, defaultValue string) string {
	v := obj.GetProp(category, key, defaultValue)
	switch v.(type) {
	case string:
		return obj.GetProp(category, key, defaultValue).(string)
	}
	return defaultValue
}

func (obj *MiniProp) GetPropBytes(category string, key string, defaultValue []byte) []byte {
	v := obj.GetProp(category, key, defaultValue)
	switch v.(type) {
	case string:
		va, er := base64.StdEncoding.DecodeString(v.(string))
		if er == nil {
			return va
		} else {
			return defaultValue
		}
	}
	return defaultValue
}

func (obj *MiniProp) SetProp(category string, key string, value interface{}) {
	v := obj.prop[category]
	if v == nil {
		obj.prop[category] = make(map[string]interface{})
	}
	obj.prop[category].(map[string]interface{})[key] = value
}

func (obj *MiniProp) SetPropInt(category string, key string, value int) {
	obj.SetProp(category, key, value)
}

func (obj *MiniProp) SetPropFloat(category string, key string, value float64) {
	obj.SetProp(category, key, value)
}

func (obj *MiniProp) SetPropString(category string, key string, value string) {
	obj.SetProp(category, key, value)
}

func (obj *MiniProp) SetPropBytes(category string, key string, value []byte) {
	obj.SetProp(category, key, base64.StdEncoding.EncodeToString([]byte(value)))
}

func (obj *MiniProp) ToJson() []byte {
	v, e := json.Marshal(obj.prop)
	if e != nil {
		return []byte("{}")
	}
	return v
}
