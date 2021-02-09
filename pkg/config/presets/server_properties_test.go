package presets

import (
	"encoding/json"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/loksonarius/mcsm/pkg/config"
)

func defaultServerProperties() *ServerProperties {
	return ServerPropertiesFromConfig(make(map[string]config.ConfigDict)).(*ServerProperties)
}

func TestServerPropertiesValidate(t *testing.T) {
	tests := []struct {
		name string
		s    func() *ServerProperties
		e    []string
	}{
		{
			name: "basic valid config",
			s:    defaultServerProperties,
			e:    []string{},
		},
		{
			name: "ports all invalid",
			s: func() *ServerProperties {
				p := defaultServerProperties()
				p.RconPort = 80001
				p.ServerPort = 80002
				p.QueryPort = 80003
				return p
			},
			e: []string{
				"'RconPort' failed on the 'port'",
				"'ServerPort' failed on the 'port'",
				"'QueryPort' failed on the 'port'",
			},
		},
		{
			name: "oneof-string fields all invalid",
			s: func() *ServerProperties {
				p := defaultServerProperties()
				p.Gamemode = "dne"
				p.Difficulty = "dne"
				return p
			},
			e: []string{
				"'Gamemode' failed on the 'oneof'",
				"'Difficulty' failed on the 'oneof'",
			},
		},
		{
			name: "range fields all invalid",
			s: func() *ServerProperties {
				p := defaultServerProperties()
				p.MaxPlayers = -1
				p.ViewDistance = 0
				p.OpPermissionLevel = 0
				p.EntityBroadcastRangePercentage = 501
				p.FunctionPermissionLevel = 5
				p.MaxWorldSize = 29999985
				return p
			},
			e: []string{
				"'MaxPlayers' failed on the 'gte'",
				"'ViewDistance' failed on the 'gte'",
				"'OpPermissionLevel' failed on the 'gte'",
				"'EntityBroadcastRangePercentage' failed on the 'lte'",
				"'FunctionPermissionLevel' failed on the 'lte'",
				"'MaxWorldSize' failed on the 'lte'",
			},
		},
		{
			name: "divby fields all invalid",
			s: func() *ServerProperties {
				p := defaultServerProperties()
				p.MaxBuildHeight = 9
				return p
			},
			e: []string{
				"'MaxBuildHeight' failed on the 'divby'",
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
