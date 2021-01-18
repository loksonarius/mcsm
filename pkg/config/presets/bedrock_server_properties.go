package presets

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/loksonarius/mcsm/pkg/config"
	"github.com/loksonarius/mcsm/pkg/config/properties"
)

type BedrockServerProperties struct {
	ServerName                          string  `properties:"key:server-name,default:Dedicated Server"`
	Gamemode                            string  `properties:"key:gamemode,default:survival"`
	Difficulty                          string  `properties:"key:difficulty,default:easy"`
	AllowCheats                         bool    `properties:"key:allow-cheats,default:false"`
	MaxPlayers                          uint    `properties:"key:max-players,default:10"`
	OnlineMode                          bool    `properties:"key:online-mode,default:true"`
	Whitelist                           bool    `properties:"key:white-list,default:false"`
	ServerPort                          uint    `properties:"key:server-port,default:19132"`
	ServerPortV6                        uint    `properties:"key:server-portv6,default:19133"`
	ViewDistance                        uint    `properties:"key:view-distance,default:32"`
	TickDistance                        uint    `properties:"key:tick-distance,default:4"`
	PlayerIdleTimeout                   uint    `properties:"key:player-idle-timeout,default:30"`
	MaxThreads                          uint    `properties:"key:max-threads,default:8"`
	LevelName                           string  `properties:"key:level-name,default:Bedrock level"`
	LevelSeed                           string  `properties:"key:level-seed,default:"`
	DefaultPlayerPermissionLevel        string  `properties:"key:default-player-permission-level,default:member"`
	TexturepackRequired                 bool    `properties:"key:texturepack-required,default:false"`
	ContentLogFileEnabled               bool    `properties:"key:content-log-file-enabled,default:false"`
	CompressionThreshold                uint    `properties:"key:compression-threshold,default:1"`
	ServerAuthoritativeMovement         string  `properties:"key:server-authoritative-movement,default:server-auth"`
	PlayerMovementScoreThreshold        uint    `properties:"key:player-movement-score-threshold,default:20"`
	PlayerMovementDistanceThreshold     float64 `properties:"key:player-movement-distance-threshold,default:0.3"`
	PlayerMovementDurationThresholdInMs uint    `properties:"key:player-movement-duration-threshold-in-ms,default:500"`
	CorrectPlayerMovement               bool    `properties:"key:correct-player-movement,default:false"`
}

func BedrockServerPropertiesFromConfig(configs map[string]config.ConfigDict) config.ConfigFile {
	cfg := config.ConfigDict{}
	if c, ok := configs["bedrock"]; ok {
		cfg = c
	}

	var s BedrockServerProperties
	properties.Unmarshal(cfg, &s)
	return &s
}

func (p *BedrockServerProperties) Path() string {
	return "server.properties"
}

func (p *BedrockServerProperties) Validate() error {
	e := func(s string, v ...interface{}) error {
		return fmt.Errorf(s, v...)
	}

	switch gm := p.Gamemode; gm {
	case "survival", "creative", "adventure":
	default:
		return e("gamemode %s not valid", gm)
	}

	switch d := p.Difficulty; d {
	case "peaceful", "easy", "normal", "hard":
	default:
		return e("difficulty %s not valid", d)
	}

	ports := []struct {
		n string
		v uint
	}{
		{"server-port", p.ServerPort},
		{"server-portv6", p.ServerPortV6},
	}
	for _, p := range ports {
		if p.v < 1 || p.v > 65535 {
			return e("port %s not in range [1,65535]", p.n)
		}
	}

	if p.TickDistance < 4 || p.TickDistance > 12 {
		return e("tick distance outside of range [4,12]")
	}

	if p.CompressionThreshold > 65535 {
		return e("compression threshold outside of range [0,65535]")
	}

	switch pl := p.DefaultPlayerPermissionLevel; pl {
	case "visitor", "member", "operator":
	default:
		return e("default player permission level %s not valid", pl)
	}

	switch sam := p.ServerAuthoritativeMovement; sam {
	case "client-auth", "server-auth":
	default:
		return e("server auth movement %s not valid", sam)
	}

	return nil
}

func (p *BedrockServerProperties) Render() []byte {
	return properties.Marshal(p)
}

func (p *BedrockServerProperties) Write() error {
	path, err := filepath.Abs(p.Path())
	if err != nil {
		return err
	}

	return ioutil.WriteFile(path, p.Render(), 0644)
}
