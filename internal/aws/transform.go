package aws

import (
	"reflect"

	"github.com/tupyy/aws-lua/internal/lua"
)

func toLua(t any) lua.Object {
	return walk(reflect.ValueOf(t))
}

func walk(v reflect.Value) lua.Object {
	o := lua.Object{}

	if v.Type().Kind() != reflect.Struct {
		return o
	}

	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		if f.IsZero() {
			continue
		}

		if f.Type().Kind() == reflect.Pointer {
			f = f.Elem()
		}

		name := reflect.Indirect(v).Type().Field(i).Name

		switch f.Type().Kind() {
		case reflect.Struct:
			o[name] = walk(f)
		case reflect.String:
			o[name] = f.String()
		case reflect.Bool:
			o[name] = f.Bool()
		case reflect.Int:
			o[name] = f.Int()
		case reflect.Array, reflect.Slice:
			if f.Len() != 0 {
				arr := make([]interface{}, 0, f.Len())
				for i := 0; i < f.Len(); i++ {
					if f.Index(i).Type().Kind() == reflect.Struct {
						arr = append(arr, walk(f.Index(i)))
					} else {
						arr = append(arr, f.Index(i).Interface())
					}
				}
				o[name] = arr
			}
		}
	}

	return o
}
