package main

import (
	"fmt"
	"reflect"
)

func i2s(data interface{}, out interface{}) error {
	o := reflect.ValueOf(out)

	if o.Kind() != reflect.Ptr {
		return fmt.Errorf("out value is not a pointer")
	}

	o = o.Elem()

	switch o.Kind() {
	case reflect.Bool:
		val, ok := data.(bool)
		if !ok {
			return fmt.Errorf("error converting to bool")
		}
		o.SetBool(val)
	case reflect.Int:
		val, ok := data.(float64)
		if !ok {
			return fmt.Errorf("error converting to float64")
		}
		o.SetInt(int64(val))
	case reflect.String:
		val, ok := data.(string)
		if !ok {
			return fmt.Errorf("error converting to string")
		}
		o.SetString(val)
	case reflect.Slice:
		val, ok := data.([]interface{})
		if !ok {
			return fmt.Errorf("error converting to slice")
		}

		for _, v := range val {
			e := reflect.New(o.Type().Elem())

			if err := i2s(v, e.Interface()); err != nil {
				return fmt.Errorf("error converting slice element: %s", err)
			}

			o.Set(reflect.Append(o, e.Elem()))
		}
	case reflect.Struct:
		val, ok := data.(map[string]interface{})
		if !ok {
			return fmt.Errorf("error converting to struct")
		}

		for i := 0; i < o.NumField(); i++ {
			name := o.Type().Field(i).Name

			v, ok := val[name]
			if !ok {
				return fmt.Errorf("field not found: %s", name)
			}

			if err := i2s(v, o.Field(i).Addr().Interface()); err != nil {
				return fmt.Errorf("field %s processing error: %s", name, err)
			}
		}
	default:
		return fmt.Errorf("%s - type is not supported", o.Kind())
	}

	return nil
}
