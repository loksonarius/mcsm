package presets

import (
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v3"

	"github.com/loksonarius/mcsm/pkg/config"
)

type PaperYmlUnsupportedSettings struct {
	AllowPistonDuplication                 bool   `yaml:"allow-piston-duplication"`
	AllowPermanentBlockBreakExploits       bool   `yaml:"allow-permanent-block-break-exploits"`
	AllowHeadlessPistons                   bool   `yaml:"allow-headless-pistons"`
	AllowPistonDuplicationReadme           string `yaml:"allow-piston-duplication-readme"`
	AllowPermanentBlockBreakExploitsReadme string `yaml:"allow-permanent-block-break-exploits-readme"`
	AllowHeadlessPistonsReadme             string `yaml:"allow-headless-pistons-readme"`
}

type PaperYmlWatchdog struct {
	EarlyWarningEvery int `yaml:"early-warning-every"`
	EarlyWarningDelay int `yaml:"early-warning-delay"`
}

type PaperYmlSpamLimiter struct {
	TabSpamIncrement    int `yaml:"tab-spam-increment"`
	TabSpamLimit        int `yaml:"tab-spam-limit"`
	RecipeSpamIncrement int `yaml:"recipe-spam-increment"`
	RecipeSpamLimit     int `yaml:"recipe-spam-limit"`
}

type PaperYmlBookSize struct {
	PageMax         int     `yaml:"page-max"`
	TotalMultiplier float64 `yaml:"total-multiplier"`
}

type PaperYmlAsyncChunks struct {
	Threads int `yaml:"threads"`
}

type PaperYmlVelocitySupport struct {
	Enabled    bool   `yaml:"enabled"`
	OnlineMode bool   `yaml:"online-mode"`
	Secret     string `yaml:"secret"`
}

type PaperYmlSettings struct {
	BungeeOnlineMode                         bool                        `yaml:"bungee-online-mode"`
	FixEntityPositionDesync                  bool                        `yaml:"fix-entity-position-desync"`
	ChunkTasksPerTick                        int                         `yaml:"chunk-tasks-per-tick"`
	SuggestPlayerNamesWhenNullTabCompletions bool                        `yaml:"suggest-player-names-when-null-tab-completions"`
	SaveEmptyScoreboardTeams                 bool                        `yaml:"save-empty-scoreboard-teams"`
	EnablePlayerCollisions                   bool                        `yaml:"enable-player-collisions"`
	PlayerAutoSaveRate                       int                         `yaml:"player-auto-save-rate"`
	MaxPlayerAutoSavePerTick                 int                         `yaml:"max-player-auto-save-per-tick"`
	UseAlternativeLuckFormula                bool                        `yaml:"use-alternative-luck-formula"`
	ConsoleHasAllPermissions                 bool                        `yaml:"console-has-all-permissions"`
	LoadPermissionsYmlBeforePlugins          bool                        `yaml:"load-permissions-yml-before-plugins"`
	RegionFileCacheSize                      int                         `yaml:"region-file-cache-size"`
	IncomingPacketSpamThreshold              int                         `yaml:"incoming-packet-spam-threshold"`
	MaxJoinsPerTick                          int                         `yaml:"max-joins-per-tick"`
	TrackPluginScoreboards                   bool                        `yaml:"track-plugin-scoreboards"`
	UnsupportedSettings                      PaperYmlUnsupportedSettings `yaml:"unsupported-settings"`
	Watchdog                                 PaperYmlWatchdog            `yaml:"watchdog"`
	SpamLimiter                              PaperYmlSpamLimiter         `yaml:"spam-limiter"`
	BookSize                                 PaperYmlBookSize            `yaml:"book-size"`
	AsyncChunks                              PaperYmlAsyncChunks         `yaml:"async-chunks"`
	VelocitySupport                          PaperYmlVelocitySupport     `yaml:"velocity-support"`
}

type PaperYmlTimings struct {
	Enabled             bool     `yaml:"enabled"`
	Verbose             bool     `yaml:"verbose"`
	ServerNamePrivacy   bool     `yaml:"server-name-privacy"`
	HiddenConfigEntries []string `yaml:"hidden-config-entries"`
	HistoryInterval     int      `yaml:"history-interval"`
	HistoryLength       int      `yaml:"history-length"`
	ServerName          string   `yaml:"server-name"`
}

type PaperYmlKick struct {
	AuthenticationServersDown string `yaml:"authentication-servers-down"`
	ConnectionThrottle        string `yaml:"connection-throttle"`
	FlyingPlayer              string `yaml:"flying-player"`
	FlyingVehicle             string `yaml:"flying-vehicle"`
}

type PaperYmlMessages struct {
	NoPermission string       `yaml:"no-permission"`
	Kick         PaperYmlKick `yaml:"kick"`
}

type PaperYmlSpawnDelay struct {
	PerPlayer bool `yaml:"per-player"`
	Ticks     int  `yaml:"ticks"`
}

type PaperYmlStart struct {
	PerPlayer bool `yaml:"per-player"`
	Day       int  `yaml:"day"`
}

type PaperYmlPillagerPatrols struct {
	SpawnChance float64            `yaml:"spawn-chance"`
	SpawnDelay  PaperYmlSpawnDelay `yaml:"spawn-delay"`
	Start       PaperYmlStart      `yaml:"start"`
}

type PaperYmlGameMechanics struct {
	ScanForLegacyEnderDragon                bool                    `yaml:"scan-for-legacy-ender-dragon"`
	DisableChestCatDetection                bool                    `yaml:"disable-chest-cat-detection"`
	NerfPigmenFromNetherPortals             bool                    `yaml:"nerf-pigmen-from-nether-portals"`
	DisableEndCredits                       bool                    `yaml:"disable-end-credits"`
	DisablePlayerCrits                      bool                    `yaml:"disable-player-crits"`
	DisableSprintInterruptionOnAttack       bool                    `yaml:"disable-sprint-interruption-on-attack"`
	ShieldBlockingDelay                     int                     `yaml:"shield-blocking-delay"`
	DisableUnloadedChunkEnderpearlExploit   bool                    `yaml:"disable-unloaded-chunk-enderpearl-exploit"`
	DisableRelativeProjectileVelocity       bool                    `yaml:"disable-relative-projectile-velocity"`
	DisableMobSpawnerSpawnEggTransformation bool                    `yaml:"disable-mob-spawner-spawn-egg-transformation"`
	DisablePillagerPatrols                  bool                    `yaml:"disable-pillager-patrols"`
	FixCuringZombieVillagerDiscountExploit  bool                    `yaml:"fix-curing-zombie-villager-discount-exploit"`
	PillagerPatrols                         PaperYmlPillagerPatrols `yaml:"pillager-patrols"`
}

type PaperYmlEntityPerChunkSaveLimit struct {
	Experience_orb int `yaml:"experience_orb"`
	Snowball       int `yaml:"snowball"`
	Ender_pearl    int `yaml:"ender_pearl"`
	Arrow          int `yaml:"arrow"`
}

type PaperYmlBamboo struct {
	Max int `yaml:"max"`
	Min int `yaml:"min"`
}

type PaperYmlMaxGrowthHeight struct {
	Cactus int            `yaml:"cactus"`
	Reeds  int            `yaml:"reeds"`
	Bamboo PaperYmlBamboo `yaml:"bamboo"`
}

type PaperYmlFishingTimeRange struct {
	Minimumticks int `yaml:"MinimumTicks"`
	Maximumticks int `yaml:"MaximumTicks"`
}

type PaperYmlDespawnRanges struct {
	Soft int `yaml:"soft"`
	Hard int `yaml:"hard"`
}

type PaperYmlDelay struct {
	Min int `yaml:"min"`
	Max int `yaml:"max"`
}

type PaperYmlFrostedIce struct {
	Enabled bool          `yaml:"enabled"`
	Delay   PaperYmlDelay `yaml:"delay"`
}

type PaperYmlLootables struct {
	AutoReplenish        bool   `yaml:"auto-replenish"`
	RestrictPlayerReloot bool   `yaml:"restrict-player-reloot"`
	ResetSeedOnFill      bool   `yaml:"reset-seed-on-fill"`
	MaxRefills           int    `yaml:"max-refills"`
	RefreshMin           string `yaml:"refresh-min"`
	RefreshMax           string `yaml:"refresh-max"`
}

type PaperYmlAntiXray struct {
	Enabled              bool     `yaml:"enabled"`
	EngineMode           int      `yaml:"engine-mode"`
	MaxChunkSectionIndex int      `yaml:"max-chunk-section-index"`
	UpdateRadius         int      `yaml:"update-radius"`
	LavaObscures         bool     `yaml:"lava-obscures"`
	UsePermission        bool     `yaml:"use-permission"`
	HiddenBlocks         []string `yaml:"hidden-blocks"`
	ReplacementBlocks    []string `yaml:"replacement-blocks"`
}

type PaperYmlViewdistances struct {
	NoTickViewDistance int `yaml:"no-tick-view-distance"`
}

type PaperYmlMobsCanAlwaysPickUpLoot struct {
	Zombies   bool `yaml:"zombies"`
	Skeletons bool `yaml:"skeletons"`
}

type PaperYmlItems struct {
	Cobblestone int `yaml:"COBBLESTONE"`
}

type PaperYmlAltItemDespawnRate struct {
	Enabled bool          `yaml:"enabled"`
	Items   PaperYmlItems `yaml:"items"`
}

type PaperYmlHopper struct {
	CooldownWhenFull bool `yaml:"cooldown-when-full"`
	DisableMoveEvent bool `yaml:"disable-move-event"`
}

type PaperYmlLightningStrikeDistanceLimit struct {
	Sound       int `yaml:"sound"`
	ImpactSound int `yaml:"impact-sound"`
	Flash       int `yaml:"flash"`
}

type PaperYmlWanderingTrader struct {
	SpawnMinuteLength           int `yaml:"spawn-minute-length"`
	SpawnDayLength              int `yaml:"spawn-day-length"`
	SpawnChanceFailureIncrement int `yaml:"spawn-chance-failure-increment"`
	SpawnChanceMin              int `yaml:"spawn-chance-min"`
	SpawnChanceMax              int `yaml:"spawn-chance-max"`
}

type PaperYmlDoorBreakingDifficulty struct {
	Zombie     []string `yaml:"zombie"`
	Vindicator []string `yaml:"vindicator"`
}

type PaperYmlSquidSpawnHeight struct {
	Maximum float64 `yaml:"maximum"`
}

type PaperYmlGeneratorSettings struct {
	FlatBedrock bool `yaml:"flat-bedrock"`
}

type PaperYmlDefault struct {
	RemoveCorruptTileEntities             bool                                 `yaml:"remove-corrupt-tile-entities"`
	ExperienceMergeMaxValue               int                                  `yaml:"experience-merge-max-value"`
	PreventMovingIntoUnloadedChunks       bool                                 `yaml:"prevent-moving-into-unloaded-chunks"`
	MaxAutoSaveChunksPerTick              int                                  `yaml:"max-auto-save-chunks-per-tick"`
	CountAllMobsForSpawning               bool                                 `yaml:"count-all-mobs-for-spawning"`
	PerPlayerMobSpawns                    bool                                 `yaml:"per-player-mob-spawns"`
	DelayChunkUnloadsBy                   string                               `yaml:"delay-chunk-unloads-by"`
	ShouldRemoveDragon                    bool                                 `yaml:"should-remove-dragon"`
	FallingBlockHeightNerf                int                                  `yaml:"falling-block-height-nerf"`
	TntEntityHeightNerf                   int                                  `yaml:"tnt-entity-height-nerf"`
	NonPlayerArrowDespawnRate             int                                  `yaml:"non-player-arrow-despawn-rate"`
	CreativeArrowDespawnRate              int                                  `yaml:"creative-arrow-despawn-rate"`
	AllChunksAreSlimeChunks               bool                                 `yaml:"all-chunks-are-slime-chunks"`
	SeedBasedFeatureSearch                bool                                 `yaml:"seed-based-feature-search"`
	UseFasterEigencraftRedstone           bool                                 `yaml:"use-faster-eigencraft-redstone"`
	WaterOverLavaFlowSpeed                int                                  `yaml:"water-over-lava-flow-speed"`
	GrassSpreadTickRate                   int                                  `yaml:"grass-spread-tick-rate"`
	ZombieVillagerInfectionChance         float64                              `yaml:"zombie-villager-infection-chance"`
	MobSpawnerTickRate                    int                                  `yaml:"mob-spawner-tick-rate"`
	KeepSpawnLoaded                       bool                                 `yaml:"keep-spawn-loaded"`
	ArmorStandsDoCollisionEntityLookups   bool                                 `yaml:"armor-stands-do-collision-entity-lookups"`
	DisableThunder                        bool                                 `yaml:"disable-thunder"`
	SkeletonHorseThunderSpawnChance       float64                              `yaml:"skeleton-horse-thunder-spawn-chance"`
	DisableIceAndSnow                     bool                                 `yaml:"disable-ice-and-snow"`
	KeepSpawnLoadedRange                  int                                  `yaml:"keep-spawn-loaded-range"`
	NetherCeilingVoidDamageHeight         int                                  `yaml:"nether-ceiling-void-damage-height"`
	AllowNonPlayerEntitiesOnScoreboards   bool                                 `yaml:"allow-non-player-entities-on-scoreboards"`
	PortalSearchRadius                    int                                  `yaml:"portal-search-radius"`
	PortalCreateRadius                    int                                  `yaml:"portal-create-radius"`
	PortalSearchVanillaDimensionScaling   bool                                 `yaml:"portal-search-vanilla-dimension-scaling"`
	ContainerUpdateTickRate               int                                  `yaml:"container-update-tick-rate"`
	ParrotsAreUnaffectedByPlayerMovement  bool                                 `yaml:"parrots-are-unaffected-by-player-movement"`
	DisableExplosionKnockback             bool                                 `yaml:"disable-explosion-knockback"`
	FixClimbingBypassingCrammingRule      bool                                 `yaml:"fix-climbing-bypassing-cramming-rule"`
	PreventTntFromMovingInWater           bool                                 `yaml:"prevent-tnt-from-moving-in-water"`
	IronGolemsCanSpawnInAir               bool                                 `yaml:"iron-golems-can-spawn-in-air"`
	ArmorStandsTick                       bool                                 `yaml:"armor-stands-tick"`
	EntitiesTargetWithFollowRange         bool                                 `yaml:"entities-target-with-follow-range"`
	SpawnerNerfedMobsShouldJump           bool                                 `yaml:"spawner-nerfed-mobs-should-jump"`
	ZombiesTargetTurtleEggs               bool                                 `yaml:"zombies-target-turtle-eggs"`
	EnableTreasureMaps                    bool                                 `yaml:"enable-treasure-maps"`
	TreasureMapsReturnAlreadyDiscovered   bool                                 `yaml:"treasure-maps-return-already-discovered"`
	FilterNbtDataFromSpawnEggsAndRelated  bool                                 `yaml:"filter-nbt-data-from-spawn-eggs-and-related"`
	MaxEntityCollisions                   int                                  `yaml:"max-entity-collisions"`
	DisableCreeperLingeringEffect         bool                                 `yaml:"disable-creeper-lingering-effect"`
	DuplicateUuidResolver                 string                               `yaml:"duplicate-uuid-resolver"`
	DuplicateUuidSaferegenDeleteRange     int                                  `yaml:"duplicate-uuid-saferegen-delete-range"`
	PhantomsDoNotSpawnOnCreativePlayers   bool                                 `yaml:"phantoms-do-not-spawn-on-creative-players"`
	PhantomsOnlyAttackInsomniacs          bool                                 `yaml:"phantoms-only-attack-insomniacs"`
	LightQueueSize                        int                                  `yaml:"light-queue-size"`
	AutoSaveInterval                      int                                  `yaml:"auto-save-interval"`
	BabyZombieMovementModifier            float64                              `yaml:"baby-zombie-movement-modifier"`
	OptimizeExplosions                    bool                                 `yaml:"optimize-explosions"`
	DisableTeleportationSuffocationCheck  bool                                 `yaml:"disable-teleportation-suffocation-check"`
	FixedChunkInhabitedTime               int                                  `yaml:"fixed-chunk-inhabited-time"`
	UseVanillaWorldScoreboardNameColoring bool                                 `yaml:"use-vanilla-world-scoreboard-name-coloring"`
	GameMechanics                         PaperYmlGameMechanics                `yaml:"game-mechanics"`
	EntityPerChunkSaveLimit               PaperYmlEntityPerChunkSaveLimit      `yaml:"entity-per-chunk-save-limit"`
	MaxGrowthHeight                       PaperYmlMaxGrowthHeight              `yaml:"max-growth-height"`
	FishingTimeRange                      PaperYmlFishingTimeRange             `yaml:"fishing-time-range"`
	DespawnRanges                         PaperYmlDespawnRanges                `yaml:"despawn-ranges"`
	FrostedIce                            PaperYmlFrostedIce                   `yaml:"frosted-ice"`
	Lootables                             PaperYmlLootables                    `yaml:"lootables"`
	AntiXray                              PaperYmlAntiXray                     `yaml:"anti-xray"`
	Viewdistances                         PaperYmlViewdistances                `yaml:"viewdistances"`
	MobsCanAlwaysPickUpLoot               PaperYmlMobsCanAlwaysPickUpLoot      `yaml:"mobs-can-always-pick-up-loot"`
	AltItemDespawnRate                    PaperYmlAltItemDespawnRate           `yaml:"alt-item-despawn-rate"`
	Hopper                                PaperYmlHopper                       `yaml:"hopper"`
	LightningStrikeDistanceLimit          PaperYmlLightningStrikeDistanceLimit `yaml:"lightning-strike-distance-limit"`
	WanderingTrader                       PaperYmlWanderingTrader              `yaml:"wandering-trader"`
	DoorBreakingDifficulty                PaperYmlDoorBreakingDifficulty       `yaml:"door-breaking-difficulty"`
	SquidSpawnHeight                      PaperYmlSquidSpawnHeight             `yaml:"squid-spawn-height"`
	GeneratorSettings                     PaperYmlGeneratorSettings            `yaml:"generator-settings"`
}

type PaperYmlWorldSettings struct {
	Default PaperYmlDefault `yaml:"default"`
}

type PaperYml struct {
	Verbose       bool                  `yaml:"verbose"`
	ConfigVersion int                   `yaml:"config-version"`
	Settings      PaperYmlSettings      `yaml:"settings"`
	Timings       PaperYmlTimings       `yaml:"timings"`
	Messages      PaperYmlMessages      `yaml:"messages"`
	WorldSettings PaperYmlWorldSettings `yaml:"world-settings"`
}

func defaultPaperYml() PaperYml {
	return PaperYml{
		Verbose:       false,
		ConfigVersion: 20,
		Settings: PaperYmlSettings{
			BungeeOnlineMode:                         true,
			FixEntityPositionDesync:                  true,
			ChunkTasksPerTick:                        1000,
			SuggestPlayerNamesWhenNullTabCompletions: true,
			SaveEmptyScoreboardTeams:                 false,
			EnablePlayerCollisions:                   true,
			PlayerAutoSaveRate:                       -1,
			MaxPlayerAutoSavePerTick:                 -1,
			UseAlternativeLuckFormula:                false,
			ConsoleHasAllPermissions:                 false,
			LoadPermissionsYmlBeforePlugins:          true,
			RegionFileCacheSize:                      256,
			IncomingPacketSpamThreshold:              300,
			MaxJoinsPerTick:                          3,
			TrackPluginScoreboards:                   false,
			UnsupportedSettings: PaperYmlUnsupportedSettings{
				AllowPistonDuplication:                 false,
				AllowPermanentBlockBreakExploits:       false,
				AllowHeadlessPistons:                   false,
				AllowPistonDuplicationReadme:           "This setting controls if player should be able to use TNT duplication, but this also allows duplicating carpet, rails and potentially other items",
				AllowPermanentBlockBreakExploitsReadme: "This setting controls if players should be able to break bedrock, end portals and other intended to be permanent blocks.",
				AllowHeadlessPistonsReadme:             "This setting controls if players should be able to create headless pistons.",
			},
			Watchdog: PaperYmlWatchdog{
				EarlyWarningEvery: 5000,
				EarlyWarningDelay: 10000,
			},
			SpamLimiter: PaperYmlSpamLimiter{
				TabSpamIncrement:    1,
				TabSpamLimit:        500,
				RecipeSpamIncrement: 1,
				RecipeSpamLimit:     20,
			},
			BookSize: PaperYmlBookSize{
				PageMax:         2560,
				TotalMultiplier: 0.98,
			},
			AsyncChunks: PaperYmlAsyncChunks{
				Threads: -1,
			},
			VelocitySupport: PaperYmlVelocitySupport{
				Enabled:    false,
				OnlineMode: false,
				Secret:     "",
			},
		},
		Timings: PaperYmlTimings{
			Enabled:             true,
			Verbose:             true,
			ServerNamePrivacy:   false,
			HiddenConfigEntries: []string{"database", "settings.bungeecord-addresses", "settings.velocity-support.secret"},
			HistoryInterval:     300,
			HistoryLength:       3600,
			ServerName:          "Unknown Server",
		},
		Messages: PaperYmlMessages{
			NoPermission: "&cI'm sorry, but you do not have permission to perform this command. Please contact the server administrators if you believe that this is in error.",
			Kick: PaperYmlKick{
				AuthenticationServersDown: "",
				ConnectionThrottle:        "Connection throttled! Please wait before reconnecting.",
				FlyingPlayer:              "Flying is not enabled on this server",
				FlyingVehicle:             "Flying is not enabled on this server",
			},
		},
		WorldSettings: PaperYmlWorldSettings{
			Default: PaperYmlDefault{
				RemoveCorruptTileEntities:             false,
				ExperienceMergeMaxValue:               -1,
				PreventMovingIntoUnloadedChunks:       false,
				MaxAutoSaveChunksPerTick:              24,
				CountAllMobsForSpawning:               false,
				PerPlayerMobSpawns:                    false,
				DelayChunkUnloadsBy:                   "10s",
				ShouldRemoveDragon:                    false,
				FallingBlockHeightNerf:                0,
				TntEntityHeightNerf:                   0,
				NonPlayerArrowDespawnRate:             -1,
				CreativeArrowDespawnRate:              -1,
				AllChunksAreSlimeChunks:               false,
				SeedBasedFeatureSearch:                true,
				UseFasterEigencraftRedstone:           false,
				WaterOverLavaFlowSpeed:                5,
				GrassSpreadTickRate:                   1,
				ZombieVillagerInfectionChance:         -1.0,
				MobSpawnerTickRate:                    1,
				KeepSpawnLoaded:                       true,
				ArmorStandsDoCollisionEntityLookups:   true,
				DisableThunder:                        false,
				SkeletonHorseThunderSpawnChance:       0.01,
				DisableIceAndSnow:                     false,
				KeepSpawnLoadedRange:                  10,
				NetherCeilingVoidDamageHeight:         0,
				AllowNonPlayerEntitiesOnScoreboards:   false,
				PortalSearchRadius:                    128,
				PortalCreateRadius:                    16,
				PortalSearchVanillaDimensionScaling:   true,
				ContainerUpdateTickRate:               1,
				ParrotsAreUnaffectedByPlayerMovement:  false,
				DisableExplosionKnockback:             false,
				FixClimbingBypassingCrammingRule:      false,
				PreventTntFromMovingInWater:           false,
				IronGolemsCanSpawnInAir:               false,
				ArmorStandsTick:                       true,
				EntitiesTargetWithFollowRange:         false,
				SpawnerNerfedMobsShouldJump:           false,
				ZombiesTargetTurtleEggs:               true,
				EnableTreasureMaps:                    true,
				TreasureMapsReturnAlreadyDiscovered:   false,
				FilterNbtDataFromSpawnEggsAndRelated:  true,
				MaxEntityCollisions:                   8,
				DisableCreeperLingeringEffect:         false,
				DuplicateUuidResolver:                 "saferegen",
				DuplicateUuidSaferegenDeleteRange:     32,
				PhantomsDoNotSpawnOnCreativePlayers:   true,
				PhantomsOnlyAttackInsomniacs:          true,
				LightQueueSize:                        20,
				AutoSaveInterval:                      -1,
				BabyZombieMovementModifier:            0.5,
				OptimizeExplosions:                    false,
				DisableTeleportationSuffocationCheck:  false,
				FixedChunkInhabitedTime:               -1,
				UseVanillaWorldScoreboardNameColoring: false,
				GameMechanics: PaperYmlGameMechanics{
					ScanForLegacyEnderDragon:                true,
					DisableChestCatDetection:                false,
					NerfPigmenFromNetherPortals:             false,
					DisableEndCredits:                       false,
					DisablePlayerCrits:                      false,
					DisableSprintInterruptionOnAttack:       false,
					ShieldBlockingDelay:                     5,
					DisableUnloadedChunkEnderpearlExploit:   true,
					DisableRelativeProjectileVelocity:       false,
					DisableMobSpawnerSpawnEggTransformation: false,
					DisablePillagerPatrols:                  false,
					FixCuringZombieVillagerDiscountExploit:  true,
					PillagerPatrols: PaperYmlPillagerPatrols{
						SpawnChance: 0.2,
						SpawnDelay: PaperYmlSpawnDelay{
							PerPlayer: false,
							Ticks:     12000,
						},
						Start: PaperYmlStart{
							PerPlayer: false,
							Day:       5,
						},
					},
				},
				EntityPerChunkSaveLimit: PaperYmlEntityPerChunkSaveLimit{
					Experience_orb: -1,
					Snowball:       -1,
					Ender_pearl:    -1,
					Arrow:          -1,
				},
				MaxGrowthHeight: PaperYmlMaxGrowthHeight{
					Cactus: 3,
					Reeds:  3,
					Bamboo: PaperYmlBamboo{
						Max: 16,
						Min: 11,
					},
				},
				FishingTimeRange: PaperYmlFishingTimeRange{
					Minimumticks: 100,
					Maximumticks: 600,
				},
				DespawnRanges: PaperYmlDespawnRanges{
					Soft: 32,
					Hard: 128,
				},
				FrostedIce: PaperYmlFrostedIce{
					Enabled: true,
					Delay: PaperYmlDelay{
						Min: 20,
						Max: 40,
					},
				},
				Lootables: PaperYmlLootables{
					AutoReplenish:        false,
					RestrictPlayerReloot: true,
					ResetSeedOnFill:      true,
					MaxRefills:           -1,
					RefreshMin:           "12h",
					RefreshMax:           "2d",
				},
				AntiXray: PaperYmlAntiXray{
					Enabled:              false,
					EngineMode:           1,
					MaxChunkSectionIndex: 3,
					UpdateRadius:         2,
					LavaObscures:         false,
					UsePermission:        false,
					HiddenBlocks:         []string{"gold_ore", "iron_ore", "coal_ore", "lapis_ore", "mossy_cobblestone", "obsidian", "chest", "diamond_ore", "redstone_ore", "clay", "emerald_ore", "ender_chest"},
					ReplacementBlocks:    []string{"stone", "oak_planks"},
				},
				Viewdistances: PaperYmlViewdistances{
					NoTickViewDistance: -1,
				},
				MobsCanAlwaysPickUpLoot: PaperYmlMobsCanAlwaysPickUpLoot{
					Zombies:   false,
					Skeletons: false,
				},
				AltItemDespawnRate: PaperYmlAltItemDespawnRate{
					Enabled: false,
					Items: PaperYmlItems{
						Cobblestone: 300,
					},
				},
				Hopper: PaperYmlHopper{
					CooldownWhenFull: true,
					DisableMoveEvent: false,
				},
				LightningStrikeDistanceLimit: PaperYmlLightningStrikeDistanceLimit{
					Sound:       -1,
					ImpactSound: -1,
					Flash:       -1,
				},
				WanderingTrader: PaperYmlWanderingTrader{
					SpawnMinuteLength:           1200,
					SpawnDayLength:              24000,
					SpawnChanceFailureIncrement: 25,
					SpawnChanceMin:              25,
					SpawnChanceMax:              75,
				},
				DoorBreakingDifficulty: PaperYmlDoorBreakingDifficulty{
					Zombie:     []string{"HARD"},
					Vindicator: []string{"NORMAL", "HARD"},
				},
				SquidSpawnHeight: PaperYmlSquidSpawnHeight{
					Maximum: 0.0,
				},
				GeneratorSettings: PaperYmlGeneratorSettings{
					FlatBedrock: false,
				},
			},
		},
	}
}

func PaperYmlFromConfig(configs map[string]config.ConfigDict) config.ConfigFile {
	p := defaultPaperYml()
	if c, ok := configs["paper"]; ok {
		marshalled, err := yaml.Marshal(c)
		if err != nil {
			return &p
		}

		if err := yaml.Unmarshal(marshalled, &p); err != nil {
			p = defaultPaperYml()
			return &p
		}
	}

	return &p
}

func (p *PaperYml) Path() string {
	return "paper.yml"
}

func (p *PaperYml) Validate() error {
	return nil
}
func (p *PaperYml) Render() []byte {
	out, _ := yaml.Marshal(p)
	return out
}

func (p *PaperYml) Write() error {
	path, err := filepath.Abs(p.Path())
	if err != nil {
		return err
	}

	return ioutil.WriteFile(path, p.Render(), 0644)
}
