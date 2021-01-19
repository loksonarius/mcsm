package presets

import (
	"encoding/json"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/loksonarius/mcsm/pkg/config"
)

func defaultBedrockServerProperties() *BedrockServerProperties {
	return BedrockServerPropertiesFromConfig(make(map[string]config.ConfigDict)).(*BedrockServerProperties)
}

func TestBedrockServerPropertiesValidate(t *testing.T) {
	tests := []struct {
		name string
		s    func() *BedrockServerProperties
		e    []string
	}{
		{
			name: "basic valid config",
			s:    defaultBedrockServerProperties,
			e:    []string{},
		},
		{
			name: "ports all invalid",
			s: func() *BedrockServerProperties {
				p := defaultBedrockServerProperties()
				p.ServerPort = 80001
				p.ServerPortV6 = 80002
				return p
			},
			e: []string{
				"'ServerPort' failed on the 'port'",
				"'ServerPortV6' failed on the 'port'",
			},
		},
		{
			name: "oneof-string fields all invalid",
			s: func() *BedrockServerProperties {
				p := defaultBedrockServerProperties()
				p.Gamemode = "dne"
				p.Difficulty = "dne"
				p.LevelType = "DNE"
				p.DefaultPlayerPermissionLevel = "dne"
				p.ServerAuthoritativeMovement = "dne"
				return p
			},
			e: []string{
				"'Gamemode' failed on the 'oneof'",
				"'Difficulty' failed on the 'oneof'",
				"'LevelType' failed on the 'oneof'",
				"'DefaultPlayerPermissionLevel' failed on the 'oneof'",
				"'ServerAuthoritativeMovement' failed on the 'oneof'",
			},
		},
		{
			name: "range fields all invalid",
			s: func() *BedrockServerProperties {
				p := defaultBedrockServerProperties()
				p.TickDistance = 13
				p.CompressionThreshold = 0
				p.PlayerMovementDistanceThreshold = -1.0
				return p
			},
			e: []string{
				"'TickDistance' failed on the 'lte'",
				"'CompressionThreshold' failed on the 'gte'",
				"'PlayerMovementDistanceThreshold' failed on the 'gte'",
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
