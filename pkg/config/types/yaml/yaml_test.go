package yaml

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/loksonarius/mcsm/pkg/config"
)

type AuxStruct struct {
	Field1 string
	Field2 bool
}

type BasicStruct struct {
	A string `yaml:"aaa"`
	B string // should still be parsed
	C int    `yaml:"-"`
	D []int
	E AuxStruct `yaml:"e"`
}

var emptyConfig = config.ConfigDict{}
var basicConfig = config.ConfigDict{
	"aaa": "foo",
	"b":   "bar",
	"c":   "ignoreme",
	"d":   []int{0, 1, 2},
	"e": map[string]interface{}{
		"field1": "value1",
		"field2": true,
		"field3": "ignored",
	},
}

func TestMarshal(t *testing.T) {
	tests := []struct {
		name string
		i    interface{}
		e    string
	}{
		{
			name: "marshals basic structs fine",
			i: BasicStruct{
				A: "foo",
				B: "bar",
				C: 1,
				D: []int{1, 2},
				E: AuxStruct{
					Field2: true,
				},
			},
			e: "aaa: foo\nb: bar\nd:\n    - 1\n    - 2\ne:\n    field1: \"\"\n    field2: true\n",
		},
		{
			name: "marshals pointer to basic structs",
			i: &BasicStruct{
				A: "foo",
				B: "bar",
				C: 1,
				D: []int{1, 2},
				E: AuxStruct{
					Field2: true,
				},
			},
			e: "aaa: foo\nb: bar\nd:\n    - 1\n    - 2\ne:\n    field1: \"\"\n    field2: true\n",
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			o, err := Marshal(tc.i)
			if err != nil {
				t.Errorf("unexpected error unmarshalling: %v", err)
			}
			g := string(o)
			if g != tc.e {
				t.Errorf("got:\n%s\nexpected:\n%s", g, tc.e)
			}
		})
	}
}

func TestUnmarshal(t *testing.T) {
	tests := []struct {
		name string
		c    config.ConfigDict
		g    BasicStruct
		e    BasicStruct
	}{
		{
			name: "unmarshals basic structs fine",
			c:    basicConfig,
			g:    BasicStruct{},
			e: BasicStruct{
				A: "foo",
				B: "bar",
				D: []int{0, 1, 2},
				E: AuxStruct{
					Field1: "value1",
					Field2: true,
				},
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			Unmarshal(tc.c, &tc.g)
			if !reflect.DeepEqual(tc.g, tc.e) {
				g, _ := json.Marshal(tc.g)
				e, _ := json.Marshal(tc.e)
				t.Errorf("got:\n%s\nexpected:\n%s", g, e)
			}
		})
	}
}
