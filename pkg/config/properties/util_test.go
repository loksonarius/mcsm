package properties

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/loksonarius/mcsm/pkg/config"
)

func TestAsStringMap(t *testing.T) {
	tests := []struct {
		name string
		c    config.ConfigDict
		m    map[string]interface{}
	}{
		{
			name: "drops non-string keys",
			c: config.ConfigDict{
				"foo1": "bar",
				"foo2": 2,
				42:     4200,
				false:  []string{},
			},
			m: map[string]interface{}{
				"foo1": "bar",
				"foo2": 2,
			},
		},
		{
			name: "handles empty configs",
			c:    emptyConfig,
			m:    map[string]interface{}{},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			out := asStringMap(tc.c)
			if !reflect.DeepEqual(out, tc.m) {
				g, _ := json.Marshal(out)
				e, _ := json.Marshal(tc.m)
				t.Errorf("got %s, expected %s", g, e)
			}
		})
	}
}

func TestToPropertiesKey(t *testing.T) {
	tests := []struct {
		name string
		s    string
		e    string
	}{
		{
			name: "leaves valid keys as-is",
			s:    "foo",
			e:    "foo",
		},
		{
			name: "downcases and separates words",
			s:    "F1AbClS",
			e:    "f1-ab-cl-s",
		},
		{
			name: "returns empty key",
			s:    "",
			e:    "",
		},
		{
			// technically, this shouldn't matter if we only process struct
			// fields, but...
			name: "ignores non-digit and non-alphas",
			s:    "*$!@#%^&*",
			e:    "*$!@#%^&*",
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			g := toPropertiesKey(tc.s)
			if g != tc.e {
				t.Errorf("got %s, expected %s", g, tc.e)
			}
		})
	}
}

func TestGetFieldKeyAndDefault(t *testing.T) {
	tests := []struct {
		name string
		f    reflect.StructField
		ek   string
		ed   string
	}{
		{
			name: "parses empty tag",
			f: reflect.StructField{
				Name: "FieldNameHere",
				Tag:  reflect.StructTag(``),
			},
			ek: "field-name-here",
			ed: "",
		},
		{
			name: "returns premade key when declared",
			f: reflect.StructField{
				Name: "FieldNameHere",
				Tag:  reflect.StructTag(`properties:"key:foeybar"`),
			},
			ek: "foeybar",
			ed: "",
		},
		{
			name: "returns automatic key when undeclared",
			f: reflect.StructField{
				Name: "FieldNameHere",
				Tag:  reflect.StructTag(`properties:""`),
			},
			ek: "field-name-here",
			ed: "",
		},
		{
			name: "returns empty for undeclared default",
			f: reflect.StructField{
				Name: "FieldNameHere",
				Tag:  reflect.StructTag(`properties:"key:foo"`),
			},
			ek: "foo",
			ed: "",
		},
		{
			name: "returns default when declared",
			f: reflect.StructField{
				Name: "FieldNameHere",
				Tag:  reflect.StructTag(`properties:"default:hoho"`),
			},
			ek: "field-name-here",
			ed: "hoho",
		},
		{
			name: "ignores extra tag fields",
			f: reflect.StructField{
				Name: "FieldNameHere",
				Tag:  reflect.StructTag(`properties:"lolwat:hoho,key:bar,tartarsauce"`),
			},
			ek: "bar",
			ed: "",
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			gk, gd := getFieldKeyAndDefault(tc.f)
			if gk != tc.ek {
				t.Errorf("got %s, expected %s", gk, tc.ek)
			}

			if gd != tc.ed {
				t.Errorf("got %s, expected %s", gd, tc.ed)
			}
		})
	}
}
