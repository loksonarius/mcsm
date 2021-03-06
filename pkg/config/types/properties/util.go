package properties

import (
	"fmt"
	"reflect"
	"strings"
	"unicode"
)

func toPropertiesKey(s string) string {
	result := ""
	for i, c := range s {
		if unicode.IsUpper(c) && i > 0 {
			result += "-"
		}
		result += string(c)
	}

	return strings.ToLower(result)
}

func getFieldKeyAndDefault(f reflect.StructField) (string, string) {
	key := toPropertiesKey(f.Name)
	def := ""

	tag := f.Tag.Get(tagName)
	if tag == "" || tag == "-" {
		return key, ""
	}

	args := strings.Split(tag, ",")
	for _, a := range args {
		parts := strings.SplitN(a, ":", 2)
		if len(parts) < 2 {
			continue
		}

		switch parts[0] {
		case "key":
			key = fmt.Sprintf("%v", parts[1])
		case "default":
			def = fmt.Sprintf("%v", parts[1])
		default:
			continue
		}
	}

	return key, def
}
