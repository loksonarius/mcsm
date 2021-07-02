package presets

import (
	"io/ioutil"
	"path/filepath"

	"github.com/loksonarius/mcsm/pkg/config"
	"github.com/loksonarius/mcsm/pkg/config/types/yaml"
)

type bukkitSettings struct {
	AllowEnd           bool   `yaml:"allow-end"`
	WarnOnOverload     bool   `yaml:"warn-on-overload"`
	PermissionsFile    string `yaml:"permissions-file"`
	UpdateFolder       string `yaml:"update-folder"`
	PluginProfiling    bool   `yaml:"plugin-profiling"`
	ConnectionThrottle uint   `yaml:"connection-throttle"`
	QueryPlugins       bool   `yaml:"query-plugins"`
	DeprecatedVerbose  string `yaml:"deprecated-verbose"`
	ShutdownMessage    string `yaml:"shutdown-message"`
	MinimumApi         string `yaml:"minimum-api"`
}

type bukkitSpawnLimits struct {
	Monsters     uint `yaml:"monsters"`
	Animals      uint `yaml:"animals"`
	WaterAnimals uint `yaml:"water-animals"`
	WaterAmbient uint `yaml:"water-ambient"`
	Ambient      uint `yaml:"ambient"`
}

type bukkitChunkGC struct {
	PeriodInTicks uint `yaml:"period-in-ticks"`
}

type bukkitTicksPer struct {
	AnimalSpawns       uint `yaml:"animal-spawns"`
	MonsterSpawns      uint `yaml:"monster-spawns"`
	WaterSpawns        uint `yaml:"water-spawns"`
	WaterAmbientSpawns uint `yaml:"water-ambient-spawns"`
	AmbientSpawns      uint `yaml:"ambient-spawns"`
	Autosave           uint `yaml:"autosave"`
}

type BukkitYml struct {
	Settings    bukkitSettings    `yaml:"settings"`
	SpawnLimits bukkitSpawnLimits `yaml:"spawn-limits"`
	ChunkGC     bukkitChunkGC     `yaml:"chunk-gc"`
	TicksPer    bukkitTicksPer    `yaml:"ticks-per"`
	Aliases     string            `yaml:"aliases"`
}

func defaultBukkitYml() BukkitYml {
	return BukkitYml{
		Settings: bukkitSettings{
			AllowEnd:           true,
			WarnOnOverload:     true,
			PermissionsFile:    "permissions.yml",
			UpdateFolder:       "update",
			PluginProfiling:    false,
			ConnectionThrottle: 4000,
			QueryPlugins:       true,
			DeprecatedVerbose:  "default",
			ShutdownMessage:    "Server closed",
			MinimumApi:         "none",
		},
		SpawnLimits: bukkitSpawnLimits{
			Monsters:     70,
			Animals:      10,
			WaterAnimals: 15,
			WaterAmbient: 20,
			Ambient:      15,
		},
		ChunkGC: bukkitChunkGC{
			PeriodInTicks: 600,
		},
		TicksPer: bukkitTicksPer{
			AnimalSpawns:       400,
			MonsterSpawns:      1,
			WaterSpawns:        1,
			WaterAmbientSpawns: 1,
			AmbientSpawns:      1,
			Autosave:           6000,
		},
		Aliases: "now-in-commands.yml",
	}
}

func BukkitYmlFromConfig(configs map[string]config.ConfigDict) config.ConfigFile {
	b := defaultBukkitYml()
	if c, ok := configs["bukkit"]; ok {
		if err := yaml.Unmarshal(c, &b); err != nil {
			b = defaultBukkitYml()
			return &b
		}
	}

	return &b
}

func (b *BukkitYml) Path() string {
	return "bukkit.yml"
}

func (b *BukkitYml) Validate() error {
	return nil
}
func (b *BukkitYml) Render() []byte {
	out, _ := yaml.Marshal(b)
	return out
}

func (b *BukkitYml) Write() error {
	path, err := filepath.Abs(b.Path())
	if err != nil {
		return err
	}

	return ioutil.WriteFile(path, b.Render(), 0644)
}
