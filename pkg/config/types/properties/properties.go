package properties

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/loksonarius/mcsm/pkg/config"
)

const tagName = "properties"

func Marshal(i interface{}) ([]byte, error) {
	out := ""

	v := reflect.ValueOf(i)
	if v.Kind() == reflect.Ptr {
		v = reflect.Indirect(v)
	}

	for i := 0; i < v.NumField(); i++ {
		typeField := v.Type().Field(i)
		kind := typeField.Type.Kind()
		tag := typeField.Tag.Get(tagName)
		if tag == "" {
			continue
		}

		key, _ := getFieldKeyAndDefault(typeField)
		value := fmt.Sprintf("%v", typeField)
		field := v.Field(i)
		switch kind {
		case reflect.String:
			value = field.String()
		case reflect.Int, reflect.Int8, reflect.Int16,
			reflect.Int32, reflect.Int64:
			value = fmt.Sprintf("%d", field.Int())
		case reflect.Uint, reflect.Uint8, reflect.Uint16,
			reflect.Uint32, reflect.Uint64:
			value = fmt.Sprintf("%d", field.Uint())
		case reflect.Float32, reflect.Float64:
			value = fmt.Sprintf("%.4f", field.Float())
		case reflect.Bool:
			value = fmt.Sprintf("%t", field.Bool())
		default:
			continue
		}

		out += fmt.Sprintf("%s=%s\n", key, value)
	}

	return []byte(out), nil
}

func Unmarshal(dict config.ConfigDict, target interface{}) error {
	v := reflect.ValueOf(target)
	for v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return fmt.Errorf("unable to unmarshal non-struct values as properties")
	}
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if !field.IsValid() || !field.CanSet() {
			continue
		}

		typeField := v.Type().Field(i)
		key, def := getFieldKeyAndDefault(typeField)
		tag := typeField.Tag.Get(tagName)
		if tag == "" {
			continue
		}

		var value interface{}
		value, valueFound := dict[key]

		// no default, no value, just leave it alone
		if !valueFound && field.Kind() != reflect.String && def == "" {
			continue
		}

		switch field.Kind() {
		case reflect.String:
			if !valueFound {
				value = def
			}
			field.SetString(value.(string))
		case reflect.Int, reflect.Int8, reflect.Int16,
			reflect.Int32, reflect.Int64:
			if !valueFound {
				if v, err := strconv.ParseInt(def, 10, 64); err == nil {
					value = v
				}
			}
			switch value.(type) {
			case int64:
				field.SetInt(value.(int64))
			default:
				field.SetInt(int64(value.(int)))
			}
		case reflect.Uint, reflect.Uint8, reflect.Uint16,
			reflect.Uint32, reflect.Uint64:
			if !valueFound {
				if v, err := strconv.ParseUint(def, 10, 64); err == nil {
					value = v
				}
			}
			switch value.(type) {
			case uint64:
				field.SetUint(value.(uint64))
			case int, int8, int16, int32:
				field.SetUint(uint64(value.(int)))
			case int64:
				field.SetUint(uint64(value.(int64)))
			default:
				field.SetUint(uint64(value.(uint)))
			}
		case reflect.Float32, reflect.Float64:
			if !valueFound {
				if v, err := strconv.ParseFloat(def, 64); err == nil {
					value = v
				}
			}
			switch value.(type) {
			case float64:
				field.SetFloat(float64(value.(float64)))
			default:
				field.SetFloat(float64(value.(float32)))
			}
		case reflect.Bool:
			if !valueFound {
				if v, err := strconv.ParseBool(def); err == nil {
					value = v
				}
			}
			field.SetBool(value.(bool))
		default:
			continue
		}
	}

	return nil
}
