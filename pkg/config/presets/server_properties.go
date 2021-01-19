package presets

import (
	"io/ioutil"
	"math"
	"path/filepath"

	"github.com/loksonarius/mcsm/pkg/config"
	"github.com/loksonarius/mcsm/pkg/config/properties"
)

type ServerProperties struct {
	EnableJmxMonitoring            bool   `properties:"key:enable-jmx-monitoring,default:false"`
	RconPort                       uint   `properties:"key:rcon.port,default:25575"`
	LevelSeed                      string `properties:"key:level-seed"`
	Gamemode                       string `properties:"key:gamemode,default:survival"`
	EnableCommandBlock             bool   `properties:"key:enable-command-block,default:false"`
	EnableQuery                    bool   `properties:"key:enable-query,default:false"`
	GeneratorSettings              string `properties:"key:generator-settings"`
	LevelName                      string `properties:"key:level-name,default:world"`
	MOTD                           string `properties:"key:motd,default:A Minecraft Server"`
	QueryPort                      uint   `properties:"key:query.port,default:25565"`
	PVP                            bool   `properties:"key:pvp,default:true"`
	GenerateStructures             bool   `properties:"key:generate-structures,default:true"`
	Difficulty                     string `properties:"key:difficulty,default:easy"`
	NetworkCompressionThreshold    int    `properties:"key:network-compression-threshold,default:256"`
	MaxTickTime                    uint   `properties:"key:max-tick-time,default:60000"`
	MaxPlayers                     uint   `properties:"key:max-players,default:20"`
	UseNativeTransport             bool   `properties:"key:use-native-transport,default:true"`
	OnlineMode                     bool   `properties:"key:online-mode,default:true"`
	EnableStatus                   bool   `properties:"key:enable-status,default:true"`
	AllowFlight                    bool   `properties:"key:allow-flight,default:false"`
	BroadcastRconToOps             bool   `properties:"key:broadcast-rcon-to-ops,default:true"`
	ViewDistance                   uint   `properties:"key:view-distance,default:10"`
	MaxBuildHeight                 uint   `properties:"key:max-build-height,default:256"`
	ServerIP                       string `properties:"key:server-ip"`
	AllowNether                    bool   `properties:"key:allow-nether,default:true"`
	ServerPort                     uint   `properties:"key:server-port,default:25565"`
	EnableRcon                     bool   `properties:"key:enable-rcon,default:false"`
	SyncChunkWrites                bool   `properties:"key:sync-chunk-writes,default:true"`
	OpPermissionLevel              uint   `properties:"key:op-permission-level,default:4"`
	PreventProxyConnections        bool   `properties:"key:prevent-proxy-connections,default:false"`
	ResourcePack                   string `properties:"key:resource-pack"`
	EntityBroadcastRangePercentage uint   `properties:"key:entity-broadcast-range-percentage,default:100"`
	RconPassword                   string `properties:"key:rcon.password"`
	PlayerIdleTimeout              uint   `properties:"key:player-idle-timeout,default:0"`
	ForceGamemode                  bool   `properties:"key:force-gamemode,default:false"`
	RateLimit                      uint   `properties:"key:rate-limit,default:0"`
	Hardcore                       bool   `properties:"key:hardcore,default:false"`
	WhiteList                      bool   `properties:"key:white-list,default:false"`
	BroadcastConsoleToOps          bool   `properties:"key:broadcast-console-to-ops,default:true"`
	SpawnNpcs                      bool   `properties:"key:spawn-npcs,default:true"`
	SpawnAnimals                   bool   `properties:"key:spawn-animals,default:true"`
	SnooperEnabled                 bool   `properties:"key:snooper-enabled,default:true"`
	FunctionPermissionLevel        uint   `properties:"key:function-permission-level,default:2"`
	LevelType                      string `properties:"key:level-type,default:default"`
	SpawnMonsters                  bool   `properties:"key:spawn-monsters,default:true"`
	EnforceWhitelist               bool   `properties:"key:enforce-whitelist,default:false"`
	ResourcePackSha1               string `properties:"key:resource-pack-sha1"`
	SpawnProtection                uint   `properties:"key:spawn-protection,default:16"`
	MaxWorldSize                   uint   `properties:"key:max-world-size,default:29999984"`
}

func ServerPropertiesFromConfig(configs map[string]config.ConfigDict) config.ConfigFile {
	cfg := config.ConfigDict{}
	if c, ok := configs["vanilla"]; ok {
		cfg = c
	}

	var s ServerProperties
	properties.Unmarshal(cfg, &s)
	return &s
}

func (p *ServerProperties) Path() string {
	return "server.properties"
}

func (p *ServerProperties) Validate() error {
	switch gm := p.Gamemode; gm {
	case "survival", "creative", "adventure", "spectator":
	default:
		return e("gamemode %s not valid", gm)
	}

	switch d := p.Difficulty; d {
	case "peaceful", "easy", "normal", "hard":
	default:
		return e("difficulty %s not valid", d)
	}

	if p.EntityBroadcastRangePercentage < 0 ||
		p.EntityBroadcastRangePercentage > 500 {
		return e("entity broadcast range not in range [0,500]")
	}

	if p.FunctionPermissionLevel < 1 ||
		p.FunctionPermissionLevel > 4 {
		return e("function permission level not in range [1,4]")
	}

	if p.OpPermissionLevel < 1 ||
		p.OpPermissionLevel > 4 {
		return e("op permission level not in range [1,4]")
	}

	if p.MaxBuildHeight%8 != 0 {
		return e("max build height is not a multiple of 8")
	}

	if p.MaxPlayers > math.MaxInt32 {
		return e("max players greater than 2^31 - 1")
	}

	if p.MaxTickTime > math.MaxInt64 {
		return e("max tick time greater than 2^63 - 1")
	}

	if p.MaxWorldSize > 29999984 {
		return e("max world size greater than 29999984")
	}

	if p.NetworkCompressionThreshold < -1 {
		return e("network compression threshold less than 1")
	}

	if p.ViewDistance < 3 ||
		p.ViewDistance > 32 {
		return e("view distance in range [3,32]")
	}

	ports := []struct {
		n string
		v uint
	}{
		{"server-port", p.ServerPort},
		{"rcon-port", p.RconPort},
	}
	for _, p := range ports {
		if p.v < 1 || p.v > 65535 {
			return e("port %s not in range [1,65535]", p.n)
		}
	}

	return nil
}

func (p *ServerProperties) Render() []byte {
	return properties.Marshal(p)
}

func (p *ServerProperties) Write() error {
	path, err := filepath.Abs(p.Path())
	if err != nil {
		return err
	}

	return ioutil.WriteFile(path, p.Render(), 0644)
}
