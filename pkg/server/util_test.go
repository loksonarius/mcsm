package server

import (
	"fmt"
	"sort"
	"strings"
	"testing"
)

func TestParseJavaVersion(t *testing.T) {
	tests := []struct {
		name string
		in   string
		e    string
	}{
		{
			name: "parses 1.8 output",
			in: `
openjdk version "1.8.0_212"
OpenJDK Runtime Environment (IcedTea 3.12.0) (Alpine 8.212.04-r0)
OpenJDK 64-Bit Server VM (build 25.252-b09, mixed mode)
			`,
			e: "1.8.0",
		},
		{
			name: "parses 1.7 output",
			in: `
java version "1.7.0_211"
OpenJDK Runtime Environment (IcedTea 2.6.17) (Alpine 7.211.2.6.17-r0)
OpenJDK 64-Bit Server VM (build 24.211-b02, mixed mode)
			`,
			e: "1.7.0",
		},
		{
			name: "parses oracle image 1.14 output",
			in: `
openjdk version "14.0.2" 2020-07-14
OpenJDK Runtime Environment (build 14.0.2+12-46)
OpenJDK 64-Bit Server VM (build 14.0.2+12-46, mixed mode, sharing)
			`,
			e: "14.0.2",
		},
		{
			name: "parses partial oracle image 1.14 output",
			in: `
openjdk version "14.0.0"
			`,
			e: "14.0.0",
		},
		{
			name: "returns 0.0.0 when missing version line",
			in: `
OpenJDK Runtime Environment (build 14.0.2+12-46)
OpenJDK 64-Bit Server VM (build 14.0.2+12-46, mixed mode, sharing)
			`,
			e: "0.0.0",
		},
		{
			name: "parses alpine image 1.14 output",
			in: `
openjdk version "14-ea" 2020-03-17
OpenJDK Runtime Environment (build 14-ea+33)
OpenJDK 64-Bit Server VM (build 14-ea+33, mixed mode, sharing)
			`,
			e: "1.14.0",
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			g := parseJavaVersion([]byte(tc.in)).String()
			if g != tc.e {
				t.Errorf("got %s, expected %s", g, tc.e)
			}
		})
	}
}

var (
	basicJarPath  = "./foo.jar"
	basicJarFlags = []string{
		"-jar",
		"./foo.jar",
		"nogui",
	}

	smallRuntimeOpts = RuntimeOpts{
		InitialMemory: Megabyte,
		MaxMemory:     Megabyte,
	}
	bigRuntimeOpts = RuntimeOpts{
		InitialMemory: 16 * Gigabyte,
		MaxMemory:     16 * Gigabyte,
	}
	debugRuntimeOpts = RuntimeOpts{
		InitialMemory: 16 * Gigabyte,
		MaxMemory:     16 * Gigabyte,
		DebugGC:       true,
	}
	smallMemArgs = []string{
		"-Xms1024k",
		"-Xmx1024k",
	}
	bigMemArgs = []string{
		"-Xms16g",
		"-Xmx16g",
	}

	baseAikarFlags = []string{
		"-XX:+UseG1GC",
		"-XX:+ParallelRefProcEnabled",
		"-XX:MaxGCPauseMillis=200",
		"-XX:+UnlockExperimentalVMOptions",
		"-XX:+DisableExplicitGC",
		"-XX:+AlwaysPreTouch",
		"-XX:G1HeapWastePercent=5",
		"-XX:G1MixedGCCountTarget=4",
		"-XX:G1MixedGCLiveThresholdPercent=90",
		"-XX:G1RSetUpdatingPauseTimePercent=5",
		"-XX:SurvivorRatio=32",
		"-XX:+PerfDisableSharedMem",
		"-XX:MaxTenuringThreshold=1",
		"-Dusing.aikars.flags=https://mcflags.emc.gs",
		"-Daikars.new.flags=true",
	}
	pre110DebugGCFlags = []string{
		"-Xloggc:gc.log",
		"-verbose:gc",
		"-XX:+PrintGCDetails",
		"-XX:+PrintGCDateStamps",
		"-XX:+PrintGCTimeStamps",
		"-XX:+UseGCLogFileRotation",
		"-XX:NumberOfGCLogFiles=5",
		"-XX:GCLogFileSize=1M",
	}
	post110DebugGCFlags = []string{
		"-Xlog:gc*:logs/gc.log:time,uptime:filecount=5,filesize=1M",
	}
	bigMemFlags = []string{
		"-XX:G1NewSizePercent=40",
		"-XX:G1MaxNewSizePercent=50",
		"-XX:G1HeapRegionSize=16M",
		"-XX:G1ReservePercent=15",
		"-XX:InitiatingHeapOccupancyPercent=20",
	}
	smallMemFlags = []string{
		"-XX:G1NewSizePercent=30",
		"-XX:G1MaxNewSizePercent=40",
		"-XX:G1HeapRegionSize=8M",
		"-XX:G1ReservePercent=20",
		"-XX:InitiatingHeapOccupancyPercent=15",
	}
)

func TestJavaArgs(t *testing.T) {
	diffArr := func(a, b []string) []string {
		diff := make([]string, 0, 0)
		m := make(map[string]bool)

		for _, item := range b {
			m[item] = true
		}

		for _, item := range a {
			if _, ok := m[item]; !ok {
				diff = append(diff, item)
			}
		}
		return diff
	}

	tests := []struct {
		name string
		p    string
		ro   RuntimeOpts
		v    javaVersion
		e    [][]string
	}{
		{
			name: "low mem runtime opts",
			p:    basicJarPath,
			ro:   smallRuntimeOpts,
			v:    javaVersion{1, 8, 0},
			e: [][]string{
				smallMemArgs,
				baseAikarFlags,
				smallMemFlags,
				basicJarFlags,
			},
		},
		{
			name: "high mem runtime opts",
			p:    basicJarPath,
			ro:   bigRuntimeOpts,
			v:    javaVersion{1, 8, 0},
			e: [][]string{
				bigMemArgs,
				baseAikarFlags,
				bigMemFlags,
				basicJarFlags,
			},
		},
		{
			name: "debug 1.9 runtime opts",
			p:    basicJarPath,
			ro:   debugRuntimeOpts,
			v:    javaVersion{1, 9, 0},
			e: [][]string{
				bigMemArgs,
				baseAikarFlags,
				bigMemFlags,
				pre110DebugGCFlags,
				basicJarFlags,
			},
		},
		{
			name: "debug 2.0 runtime opts",
			p:    basicJarPath,
			ro:   debugRuntimeOpts,
			v:    javaVersion{2, 0, 0},
			e: [][]string{
				bigMemArgs,
				baseAikarFlags,
				bigMemFlags,
				post110DebugGCFlags,
				basicJarFlags,
			},
		},
		{
			name: "debug 1.6 runtime opts",
			p:    basicJarPath,
			ro:   debugRuntimeOpts,
			v:    javaVersion{1, 6, 0},
			e: [][]string{
				bigMemArgs,
				baseAikarFlags,
				bigMemFlags,
				basicJarFlags,
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			var e []string
			for _, s := range tc.e {
				e = append(e, s...)
			}
			g := javaArgs(tc.p, tc.ro, tc.v)

			d := diffArr(e, g)
			sort.Strings(d)
			if len(d) > 0 {
				fmt.Printf("got: %s\n", strings.Join(g, ","))
				t.Errorf("missing %d expected flags: %s", len(d), d)
			}
		})
	}
}
