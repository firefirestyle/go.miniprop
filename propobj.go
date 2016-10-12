package gominiprop

import "encoding/json"

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

func (obj *MiniProp) GetPropString(category string, key string, defaultValue string) string {
	v := obj.GetProp(category, key, defaultValue)
	switch v.(type) {
	case string:
		return obj.GetProp(category, key, defaultValue).(string)
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

func (obj *MiniProp) SetPropString(category string, key string, value string) {
	obj.SetProp(category, key, value)
}

func (obj *MiniProp) ToJson() []byte {
	v, e := json.Marshal(obj.prop)
	if e != nil {
		return []byte("{}")
	}
	return v
}
