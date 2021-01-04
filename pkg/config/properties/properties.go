package properties

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/loksonarius/mcsm/pkg/config"
)

const tagName = "properties"

func Marshal(i interface{}) []byte {
	out := ""

	v := reflect.ValueOf(i)
	if v.Kind() == reflect.Ptr {
		v = reflect.Indirect(v)
	}

	for i := 0; i < v.NumField(); i++ {
		typeField := v.Type().Field(i)
		kind := typeField.Type.Kind()
		tag := typeField.Tag.Get(tagName)
		if tag == "" || tag == "-" {
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
		case reflect.Bool:
			value = fmt.Sprintf("%t", field.Bool())
		default:
			continue
		}

		out += fmt.Sprintf("%s=%s\n", key, value)
	}

	return []byte(out)
}

func Unmarshal(dict config.ConfigDict, target interface{}) {
	properties := asStringMap(dict)

	v := reflect.ValueOf(target)
	if v.Kind() == reflect.Ptr {
		v = reflect.Indirect(v)
	}

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if !field.IsValid() {
			continue
		}

		typeField := v.Type().Field(i)
		key, def := getFieldKeyAndDefault(typeField)

		var value interface{}
		value, valueFound := properties[key]

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
			field.SetInt(value.(int64))
		case reflect.Uint, reflect.Uint8, reflect.Uint16,
			reflect.Uint32, reflect.Uint64:
			if !valueFound {
				if v, err := strconv.ParseUint(def, 10, 64); err == nil {
					value = v
				}
			}
			field.SetUint(value.(uint64))
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
}
