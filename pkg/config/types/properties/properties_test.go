package properties

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/loksonarius/mcsm/pkg/config"
)

type BasicStruct struct {
	A string  `properties:"key:basica"`
	B int     `properties:"-"`
	C bool    // ignored field
	D uint64  `properties:"-"`
	E bool    `properties:"-"`
	F []int   `properties:"-"`
	G int32   `properties:"default:16"`
	H string  `properties:"default:hhh"`
	I []bool  `properties:"key:impossible?"`
	J bool    `properties:"default:true"`
	K float32 `properties:"default:0.32"`
	L float64 `properties:"default:0.64"`
}

var emptyConfig = config.ConfigDict{}
var basicConfig = config.ConfigDict{
	"basica": "foo",
	"b":      42,
	"d":      86,
	"e":      true,
	"l":      0.86,
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
				B: 42,
				C: true,
				L: 0.86,
			},
			e: "basica=foo\nb=42\nd=0\ne=false\ng=0\nh=\nj=false\nk=0.0000\nl=0.8600\n",
		},
		{
			name: "marshals pointer to basic structs",
			i: &BasicStruct{
				A: "foo",
				B: 42,
				C: true,
				L: 0.86,
			},
			e: "basica=foo\nb=42\nd=0\ne=false\ng=0\nh=\nj=false\nk=0.0000\nl=0.8600\n",
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
				B: 42,
				E: true,
				D: 86,
				G: 16,
				H: "hhh",
				J: true,
				K: 0.32,
				L: 0.86,
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
				t.Errorf("got %s, expected %s", g, e)
			}
		})
	}
}

func TestUnmarshalEdgeCases(t *testing.T) {
	t.Run("doesn't panic unmarshalling non-structs", func(t *testing.T) {
		c := config.ConfigDict{
			"impossible?": []bool{true},
		}
		g := true
		e := true
		Unmarshal(c, &g)
		if g != e { // the real test is if we panci
			t.Errorf("got %t, expected %t", g, e)
		}
	})

	t.Run("ignores non-struct fields", func(t *testing.T) {
		c := config.ConfigDict{
			"impossible?": []bool{true},
		}
		var g BasicStruct
		e := BasicStruct{
			G: 16,
			H: "hhh",
			J: true,
			K: 0.32,
			L: 0.64,
		}
		Unmarshal(c, &g)
		if !reflect.DeepEqual(g, e) {
			g, _ := json.Marshal(g)
			e, _ := json.Marshal(e)
			t.Errorf("got %s, expected %s", g, e)
		}
	})

	t.Run("skips non-editable fields on a non-ptr", func(t *testing.T) {
		c := config.ConfigDict{
			"basica": "foo",
			"g":      8,
		}
		var g BasicStruct
		e := BasicStruct{}
		Unmarshal(c, g)
		if !reflect.DeepEqual(g, e) {
			g, _ := json.Marshal(g)
			e, _ := json.Marshal(e)
			t.Errorf("got %s, expected %s", g, e)
		}
	})

	t.Run("handles different type defaults", func(t *testing.T) {
		type StructOfDefaults struct {
			S  string `properties:"default:foo"`
			I  int    `properties:"default:-2"`
			II int64  `properties:"default:-3"`
			U  uint   `properties:"default:2"`
			UU uint64 `properties:"default:3"`
			B  bool   `properties:"default:true"`
		}

		c := config.ConfigDict{}
		var g StructOfDefaults
		e := StructOfDefaults{
			S:  "foo",
			I:  -2,
			II: -3,
			U:  2,
			UU: 3,
			B:  true,
		}
		Unmarshal(c, &g)
		if !reflect.DeepEqual(g, e) {
			g, _ := json.Marshal(g)
			e, _ := json.Marshal(e)
			t.Errorf("got %s, expected %s", g, e)
		}
	})

	t.Run("handles non-uint literals for uint fields", func(t *testing.T) {
		type UintStruct struct {
			A uint   `properties:"default:1"`
			B uint   `properties:"default:2"`
			C uint64 `properties:"default:3"`
		}

		c := config.ConfigDict{
			"a": uint(1),
			"b": int(2),
			"c": int64(3),
		}
		var g UintStruct
		e := UintStruct{
			A: 1,
			B: 2,
			C: 3,
		}
		Unmarshal(c, &g)
		if !reflect.DeepEqual(g, e) {
			g, _ := json.Marshal(g)
			e, _ := json.Marshal(e)
			t.Errorf("got %s, expected %s", g, e)
		}
	})
}
