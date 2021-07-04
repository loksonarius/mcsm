package presets

import (
	"io/ioutil"
	"path/filepath"

	"github.com/loksonarius/mcsm/pkg/config"
	"github.com/loksonarius/mcsm/pkg/config/types/yaml"
)

type spigotMessages struct {
	Whitelist      string `yaml:"whitelist"`
	UnknownCommand string `yaml:"unknown-command"`
	ServerFull     string `yaml:"server-full"`
	OutdatedClient string `yaml:"outdated-client"`
	OutdatedServer string `yaml:"outdated-server"`
	Restart        string `yaml:"restart"`
}

type spigotMaxhealth struct {
	Max float64 `yaml:"max"`
}

type spigotMovementspeed struct {
	Max float64 `yaml:"max"`
}

type spigotAttackdamage struct {
	Max float64 `yaml:"max"`
}

type spigotAttribute struct {
	Maxhealth     spigotMaxhealth     `yaml:"maxHealth"`
	Movementspeed spigotMovementspeed `yaml:"movementSpeed"`
	Attackdamage  spigotAttackdamage  `yaml:"attackDamage"`
}

type spigotSettings struct {
	SaveUserCacheOnStopOnly   bool            `yaml:"save-user-cache-on-stop-only"`
	MovedWronglyThreshold     float64         `yaml:"moved-wrongly-threshold"`
	MovedTooQuicklyMultiplier float64         `yaml:"moved-too-quickly-multiplier"`
	LogVillagerDeaths         bool            `yaml:"log-villager-deaths"`
	Bungeecord                bool            `yaml:"bungeecord"`
	SampleCount               uint            `yaml:"sample-count"`
	PlayerShuffle             uint            `yaml:"player-shuffle"`
	UserCacheSize             uint            `yaml:"user-cache-size"`
	TimeoutTime               uint            `yaml:"timeout-time"`
	RestartOnCrash            bool            `yaml:"restart-on-crash"`
	RestartScript             string          `yaml:"restart-script"`
	NettyThreads              uint            `yaml:"netty-threads"`
	Debug                     bool            `yaml:"debug"`
	Attribute                 spigotAttribute `yaml:"attribute"`
}

type spigotAdvancements struct {
	DisableSaving bool     `yaml:"disable-saving"`
	Disabled      []string `yaml:"disabled"`
}

type spigotStats struct {
	DisableSaving bool              `yaml:"disable-saving"`
	ForcedStats   map[string]string `yaml:"forced-stats"`
}

type spigotPlayers struct {
	DisableSaving bool `yaml:"disable-saving"`
}

type spigotCommands struct {
	Log                       bool     `yaml:"log"`
	TabComplete               uint     `yaml:"tab-complete"`
	SendNamespaced            bool     `yaml:"send-namespaced"`
	SpamExclusions            []string `yaml:"spam-exclusions"`
	SilentCommandblockConsole bool     `yaml:"silent-commandblock-console"`
	ReplaceCommands           []string `yaml:"replace-commands"`
}

type spigotEntityTrackingRange struct {
	Players  uint `yaml:"players"`
	Animals  uint `yaml:"animals"`
	Monsters uint `yaml:"monsters"`
	Misc     uint `yaml:"misc"`
	Other    uint `yaml:"other"`
}

type spigotMergeRadius struct {
	Item float64 `yaml:"item"`
	Exp  float64 `yaml:"exp"`
}

type spigotHunger struct {
	JumpWalkExhaustion   float64 `yaml:"jump-walk-exhaustion"`
	JumpSprintExhaustion float64 `yaml:"jump-sprint-exhaustion"`
	CombatExhaustion     float64 `yaml:"combat-exhaustion"`
	RegenExhaustion      float64 `yaml:"regen-exhaustion"`
	SwimMultiplier       float64 `yaml:"swim-multiplier"`
	SprintMultiplier     float64 `yaml:"sprint-multiplier"`
	OtherMultiplier      float64 `yaml:"other-multiplier"`
}

type spigotMaxTickTime struct {
	Tile   uint `yaml:"tile"`
	Entity uint `yaml:"entity"`
}

type spigotSquidSpawnRange struct {
	Min float64 `yaml:"min"`
}

type spigotGrowth struct {
	CactusModifier     uint `yaml:"cactus-modifier"`
	CaneModifier       uint `yaml:"cane-modifier"`
	MelonModifier      uint `yaml:"melon-modifier"`
	MushroomModifier   uint `yaml:"mushroom-modifier"`
	PumpkinModifier    uint `yaml:"pumpkin-modifier"`
	SaplingModifier    uint `yaml:"sapling-modifier"`
	BeetrootModifier   uint `yaml:"beetroot-modifier"`
	CarrotModifier     uint `yaml:"carrot-modifier"`
	PotatoModifier     uint `yaml:"potato-modifier"`
	WheatModifier      uint `yaml:"wheat-modifier"`
	NetherwartModifier uint `yaml:"netherwart-modifier"`
	VineModifier       uint `yaml:"vine-modifier"`
	CocoaModifier      uint `yaml:"cocoa-modifier"`
	BambooModifier     uint `yaml:"bamboo-modifier"`
	SweetberryModifier uint `yaml:"sweetberry-modifier"`
	KelpModifier       uint `yaml:"kelp-modifier"`
}

type spigotWakeUpInactive struct {
	AnimalsMaxPerTick        uint `yaml:"animals-max-per-tick"`
	AnimalsEvery             uint `yaml:"animals-every"`
	AnimalsFor               uint `yaml:"animals-for"`
	MonstersMaxPerTick       uint `yaml:"monsters-max-per-tick"`
	MonstersEvery            uint `yaml:"monsters-every"`
	MonstersFor              uint `yaml:"monsters-for"`
	VillagersMaxPerTick      uint `yaml:"villagers-max-per-tick"`
	VillagersEvery           uint `yaml:"villagers-every"`
	VillagersFor             uint `yaml:"villagers-for"`
	FlyingMonstersMaxPerTick uint `yaml:"flying-monsters-max-per-tick"`
	FlyingMonstersEvery      uint `yaml:"flying-monsters-every"`
	FlyingMonstersFor        uint `yaml:"flying-monsters-for"`
}

type spigotEntityActivationRange struct {
	Animals                    uint                 `yaml:"animals"`
	Monsters                   uint                 `yaml:"monsters"`
	Raiders                    uint                 `yaml:"raiders"`
	Misc                       uint                 `yaml:"misc"`
	Water                      uint                 `yaml:"water"`
	Villagers                  uint                 `yaml:"villagers"`
	FlyingMonsters             uint                 `yaml:"flying-monsters"`
	VillagersWorkImmunityAfter uint                 `yaml:"villagers-work-immunity-after"`
	VillagersWorkImmunityFor   uint                 `yaml:"villagers-work-immunity-for"`
	VillagersActiveForPanic    bool                 `yaml:"villagers-active-for-panic"`
	TickInactiveVillagers      bool                 `yaml:"tick-inactive-villagers"`
	WakeUpInactive             spigotWakeUpInactive `yaml:"wake-up-inactive"`
}

type spigotTicksPer struct {
	HopperTransfer uint `yaml:"hopper-transfer"`
	HopperCheck    uint `yaml:"hopper-check"`
}

type spigotDefault struct {
	Verbose                         bool                        `yaml:"verbose"`
	WitherSpawnSoundRadius          uint                        `yaml:"wither-spawn-sound-radius"`
	EnableZombiePigmenPortalSpawns  bool                        `yaml:"enable-zombie-pigmen-portal-spawns"`
	ViewDistance                    string                      `yaml:"view-distance" validate:"oneof=default 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15"` // honestly, this is subtly genius
	ItemDespawnRate                 uint                        `yaml:"item-despawn-rate"`
	ArrowDespawnRate                uint                        `yaml:"arrow-despawn-rate"`
	TridentDespawnRate              uint                        `yaml:"trident-despawn-rate"`
	ZombieAggressiveTowardsVillager bool                        `yaml:"zombie-aggressive-towards-villager"`
	HangingTickFrequency            uint                        `yaml:"hanging-tick-frequency"`
	NerfSpawnerMobs                 bool                        `yaml:"nerf-spawner-mobs"`
	MobSpawnRange                   uint                        `yaml:"mob-spawn-range"`
	MaxTntPerTick                   uint                        `yaml:"max-tnt-per-tick"`
	EndPortalSoundRadius            uint                        `yaml:"end-portal-sound-radius"`
	HopperAmount                    uint                        `yaml:"hopper-amount"`
	DragonDeathSoundRadius          uint                        `yaml:"dragon-death-sound-radius"`
	SeedVillage                     uint                        `yaml:"seed-village"`
	SeedDesert                      uint                        `yaml:"seed-desert"`
	SeedIgloo                       uint                        `yaml:"seed-igloo"`
	SeedJungle                      uint                        `yaml:"seed-jungle"`
	SeedSwamp                       uint                        `yaml:"seed-swamp"`
	SeedMonument                    uint                        `yaml:"seed-monument"`
	SeedShipwreck                   uint                        `yaml:"seed-shipwreck"`
	SeedOcean                       uint                        `yaml:"seed-ocean"`
	SeedOutpost                     uint                        `yaml:"seed-outpost"`
	SeedEndcity                     uint                        `yaml:"seed-endcity"`
	SeedSlime                       uint                        `yaml:"seed-slime"`
	SeedBastion                     uint                        `yaml:"seed-bastion"`
	SeedFortress                    uint                        `yaml:"seed-fortress"`
	SeedMansion                     uint                        `yaml:"seed-mansion"`
	SeedFossil                      uint                        `yaml:"seed-fossil"`
	SeedPortal                      uint                        `yaml:"seed-portal"`
	MaxEntityCollisions             uint                        `yaml:"max-entity-collisions"`
	EntityTrackingRange             spigotEntityTrackingRange   `yaml:"entity-tracking-range"`
	MergeRadius                     spigotMergeRadius           `yaml:"merge-radius"`
	Hunger                          spigotHunger                `yaml:"hunger"`
	MaxTickTime                     spigotMaxTickTime           `yaml:"max-tick-time"`
	SquidSpawnRange                 spigotSquidSpawnRange       `yaml:"squid-spawn-range"`
	Growth                          spigotGrowth                `yaml:"growth"`
	EntityActivationRange           spigotEntityActivationRange `yaml:"entity-activation-range"`
	TicksPer                        spigotTicksPer              `yaml:"ticks-per"`
}

type spigotWorldSettings struct {
	Default spigotDefault `yaml:"default"`
}

type SpigotYml struct {
	ConfigVersion uint                `yaml:"config-version"`
	Messages      spigotMessages      `yaml:"messages"`
	Settings      spigotSettings      `yaml:"settings"`
	Advancements  spigotAdvancements  `yaml:"advancements"`
	Stats         spigotStats         `yaml:"stats"`
	Players       spigotPlayers       `yaml:"players"`
	Commands      spigotCommands      `yaml:"commands"`
	WorldSettings spigotWorldSettings `yaml:"world-settings"`
}

func defaultSpigotYml() SpigotYml {
	return SpigotYml{
		ConfigVersion: 12,
		Messages: spigotMessages{
			Whitelist:      "You are not whitelisted on this server!",
			UnknownCommand: "Unknown command. Type \"/help\" for help.",
			ServerFull:     "The server is full!",
			OutdatedClient: "Outdated client! Please use {0}",
			OutdatedServer: "Outdated server! I'm still on {0}",
			Restart:        "Server is restarting",
		},
		Settings: spigotSettings{
			SaveUserCacheOnStopOnly:   false,
			MovedWronglyThreshold:     0.0625,
			MovedTooQuicklyMultiplier: 10.0,
			LogVillagerDeaths:         true,
			Bungeecord:                false,
			SampleCount:               12,
			PlayerShuffle:             0,
			UserCacheSize:             1000,
			TimeoutTime:               60,
			RestartOnCrash:            true,
			RestartScript:             "./start.sh",
			NettyThreads:              4,
			Debug:                     false,
			Attribute: spigotAttribute{
				Maxhealth: spigotMaxhealth{
					Max: 2048.0,
				},
				Movementspeed: spigotMovementspeed{
					Max: 2048.0,
				},
				Attackdamage: spigotAttackdamage{
					Max: 2048.0,
				},
			},
		},
		Advancements: spigotAdvancements{
			DisableSaving: false,
			Disabled:      []string{"minecraft:story/disabled"},
		},
		Stats: spigotStats{
			DisableSaving: false,
			ForcedStats:   map[string]string{},
		},
		Players: spigotPlayers{
			DisableSaving: false,
		},
		Commands: spigotCommands{
			Log:                       true,
			TabComplete:               0,
			SendNamespaced:            true,
			SpamExclusions:            []string{"/skill"},
			SilentCommandblockConsole: false,
			ReplaceCommands:           []string{"setblock", "summon", "testforblock", "tellraw"},
		},
		WorldSettings: spigotWorldSettings{
			Default: spigotDefault{
				Verbose:                         false,
				WitherSpawnSoundRadius:          0,
				EnableZombiePigmenPortalSpawns:  true,
				ViewDistance:                    "default",
				ItemDespawnRate:                 6000,
				ArrowDespawnRate:                1200,
				TridentDespawnRate:              1200,
				ZombieAggressiveTowardsVillager: true,
				HangingTickFrequency:            100,
				NerfSpawnerMobs:                 false,
				MobSpawnRange:                   8,
				MaxTntPerTick:                   100,
				EndPortalSoundRadius:            0,
				HopperAmount:                    1,
				DragonDeathSoundRadius:          0,
				SeedVillage:                     10387312,
				SeedDesert:                      14357617,
				SeedIgloo:                       14357618,
				SeedJungle:                      14357619,
				SeedSwamp:                       14357620,
				SeedMonument:                    10387313,
				SeedShipwreck:                   165745295,
				SeedOcean:                       14357621,
				SeedOutpost:                     165745296,
				SeedEndcity:                     10387313,
				SeedSlime:                       987234911,
				SeedBastion:                     30084232,
				SeedFortress:                    30084232,
				SeedMansion:                     10387319,
				SeedFossil:                      14357921,
				SeedPortal:                      34222645,
				MaxEntityCollisions:             8,
				EntityTrackingRange: spigotEntityTrackingRange{
					Players:  48,
					Animals:  48,
					Monsters: 48,
					Misc:     32,
					Other:    64,
				},
				MergeRadius: spigotMergeRadius{
					Item: 2.5,
					Exp:  3.0,
				},
				Hunger: spigotHunger{
					JumpWalkExhaustion:   0.05,
					JumpSprintExhaustion: 0.2,
					CombatExhaustion:     0.1,
					RegenExhaustion:      6.0,
					SwimMultiplier:       0.01,
					SprintMultiplier:     0.1,
					OtherMultiplier:      0.0,
				},
				MaxTickTime: spigotMaxTickTime{
					Tile:   50,
					Entity: 50,
				},
				SquidSpawnRange: spigotSquidSpawnRange{
					Min: 45.0,
				},
				Growth: spigotGrowth{
					CactusModifier:     100,
					CaneModifier:       100,
					MelonModifier:      100,
					MushroomModifier:   100,
					PumpkinModifier:    100,
					SaplingModifier:    100,
					BeetrootModifier:   100,
					CarrotModifier:     100,
					PotatoModifier:     100,
					WheatModifier:      100,
					NetherwartModifier: 100,
					VineModifier:       100,
					CocoaModifier:      100,
					BambooModifier:     100,
					SweetberryModifier: 100,
					KelpModifier:       100,
				},
				EntityActivationRange: spigotEntityActivationRange{
					Animals:                    32,
					Monsters:                   32,
					Raiders:                    48,
					Misc:                       16,
					Water:                      16,
					Villagers:                  32,
					FlyingMonsters:             32,
					VillagersWorkImmunityAfter: 100,
					VillagersWorkImmunityFor:   20,
					VillagersActiveForPanic:    true,
					TickInactiveVillagers:      true,
					WakeUpInactive: spigotWakeUpInactive{
						AnimalsMaxPerTick:        4,
						AnimalsEvery:             1200,
						AnimalsFor:               100,
						MonstersMaxPerTick:       8,
						MonstersEvery:            400,
						MonstersFor:              100,
						VillagersMaxPerTick:      4,
						VillagersEvery:           600,
						VillagersFor:             100,
						FlyingMonstersMaxPerTick: 8,
						FlyingMonstersEvery:      200,
						FlyingMonstersFor:        100,
					},
				},
				TicksPer: spigotTicksPer{
					HopperTransfer: 8,
					HopperCheck:    1,
				},
			},
		},
	}

}

func SpigotYmlFromConfig(configs map[string]config.ConfigDict) config.ConfigFile {
	s := defaultSpigotYml()
	if c, ok := configs["spigot"]; ok {
		if err := yaml.Unmarshal(c, &s); err != nil {
			s = defaultSpigotYml()
			return &s
		}
	}

	return &s
}

func (s *SpigotYml) Path() string {
	return "spigot.yml"
}

func (s *SpigotYml) Validate() error {
	return validate.Struct(s)
}
func (s *SpigotYml) Render() []byte {
	out, _ := yaml.Marshal(s)
	return out
}

func (s *SpigotYml) Write() error {
	path, err := filepath.Abs(s.Path())
	if err != nil {
		return err
	}

	return ioutil.WriteFile(path, s.Render(), 0644)
}
