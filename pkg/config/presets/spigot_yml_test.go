package presets

import (
	"encoding/json"
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestSpigotYmlValidate(t *testing.T) {
	tests := []struct {
		name string
		s    func() *SpigotYml
		e    []string
	}{
		{
			name: "basic valid config",
			s: func() *SpigotYml {
				s := defaultSpigotYml()
				return &s
			},
			e: []string{},
		},
		{
			name: "invalid view distance",
			s: func() *SpigotYml {
				s := defaultSpigotYml()
				s.WorldSettings.Default.ViewDistance = "notoneof"
				return &s
			},
			e: []string{
				"'ViewDistance' failed on the 'oneof'",
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			err := tc.s().Validate()
			errs, ok := err.(validator.ValidationErrors)
			if !ok {
				errs = validator.ValidationErrors{}
			}
			got, unexpected, missing := getDiff(errs, tc.e)
			if len(unexpected) > 0 || len(missing) > 0 {
				g, _ := json.Marshal(got)
				u, _ := json.Marshal(unexpected)
				m, _ := json.Marshal(missing)
				e, _ := json.Marshal(tc.e)
				t.Errorf(
					"got %s, unexpected %s, missing %s, expected %s",
					string(g), string(u), string(m), string(e))
			}
		})
	}
}
