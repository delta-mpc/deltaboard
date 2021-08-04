package reflect_model

import (
	"errors"
	"fmt"
	"reflect"
)

func GetValueByFieldName(filedName string, model interface{}) (val reflect.Value, err error) {
	t := reflect.TypeOf(model)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		err = errors.New("reflect_model.GetValueByFieldName: model is not struct")
		return
	}
	v := reflect.ValueOf(model)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	val = v.FieldByName(filedName)
	if !val.IsValid() {
		err = fmt.Errorf("no filed %v in model %v", filedName, t)
	}
	return
}

func SetValueByFiledName(fieldName string, model, value interface{}) (err error) {

	t := reflect.TypeOf(model)
	if t.Kind() != reflect.Ptr {
		err = errors.New("reflect_model.SetValueByFiledName: model is not ptr")
		return
	}
	if t.Elem().Kind() != reflect.Struct {
		err = errors.New("reflect_model.SetValueByFiledName: model is not ptr of struct")
		return
	}
	v := reflect.ValueOf(model).Elem()
	f := v.FieldByName(fieldName)
	if f.CanSet() {
		if f.Type() == reflect.TypeOf(value) {
			f.Set(reflect.ValueOf(value))
		} else {
			err = fmt.Errorf("reflect_model.SetValueByFiledName: filed type %v is not equal the value type %v",
				f.Type(), reflect.TypeOf(value))
		}
	} else {
		err = fmt.Errorf("reflect_model.SetValueByFiledName: field %v cannot set", fieldName)
	}
	return
}
