package presets

import (
	"io/ioutil"
	"path/filepath"

	"github.com/loksonarius/mcsm/pkg/config"
	"github.com/loksonarius/mcsm/pkg/config/types/properties"
)

type BedrockServerProperties struct {
	ServerName                          string  `properties:"key:server-name,default:Dedicated Server"`
	Gamemode                            string  `properties:"key:gamemode,default:survival" validate:"oneof=survival creative adventure"`
	Difficulty                          string  `properties:"key:difficulty,default:easy" validate:"oneof=peaceful easy normal hard"`
	AllowCheats                         bool    `properties:"key:allow-cheats,default:false"`
	MaxPlayers                          uint    `properties:"key:max-players,default:10"`
	OnlineMode                          bool    `properties:"key:online-mode,default:true"`
	Whitelist                           bool    `properties:"key:white-list,default:false"`
	ServerPort                          uint    `properties:"key:server-port,default:19132" validate:"port"`
	ServerPortV6                        uint    `properties:"key:server-portv6,default:19133" validate:"port"`
	ViewDistance                        uint    `properties:"key:view-distance,default:32"`
	TickDistance                        uint    `properties:"key:tick-distance,default:4" validate:"gte=4,lte=12"`
	PlayerIdleTimeout                   uint    `properties:"key:player-idle-timeout,default:30"`
	MaxThreads                          uint    `properties:"key:max-threads,default:8"`
	LevelName                           string  `properties:"key:level-name,default:Bedrock level"`
	LevelSeed                           string  `properties:"key:level-seed,default:"`
	LevelType                           string  `properties:"key:level-type,default:DEFAULT" validate:"oneof=FLAT LEGACY DEFAULT"`
	DefaultPlayerPermissionLevel        string  `properties:"key:default-player-permission-level,default:member" validate:"oneof=visitor member operator"`
	TexturepackRequired                 bool    `properties:"key:texturepack-required,default:false"`
	ContentLogFileEnabled               bool    `properties:"key:content-log-file-enabled,default:false"`
	CompressionThreshold                uint    `properties:"key:compression-threshold,default:1" validate:"gte=1,lte=65535"`
	ServerAuthoritativeMovement         string  `properties:"key:server-authoritative-movement,default:server-auth" validate:"oneof=client-auth server-auth"`
	PlayerMovementScoreThreshold        uint    `properties:"key:player-movement-score-threshold,default:20"`
	PlayerMovementDistanceThreshold     float64 `properties:"key:player-movement-distance-threshold,default:0.3" validate:"gte=0"`
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
	return validate.Struct(p)
}

func (p *BedrockServerProperties) Render() []byte {
	out, _ := properties.Marshal(p)
	return out
}

func (p *BedrockServerProperties) Write() error {
	path, err := filepath.Abs(p.Path())
	if err != nil {
		return err
	}

	return ioutil.WriteFile(path, p.Render(), 0644)
}
