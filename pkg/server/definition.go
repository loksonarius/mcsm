package server

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strconv"

	"gopkg.in/yaml.v3"

	"github.com/loksonarius/mcsm/pkg/config"
)

type Memory uint64

const (
	Byte     = Memory(1 << iota)
	Kilobyte = Memory(1 << (iota * 10))
	Megabyte = Memory(1 << (iota * 10))
	Gigabyte = Memory(1 << (iota * 10))
)

func (m Memory) String() string {
	if m > Gigabyte && m%Gigabyte == 0 {
		return fmt.Sprintf("%dg", m/Gigabyte)
	}

	if m > Megabyte && m%Megabyte == 0 {
		return fmt.Sprintf("%dm", m/Megabyte)
	}

	if m > Kilobyte && m%Kilobyte == 0 {
		return fmt.Sprintf("%dk", m/Kilobyte)
	}

	return fmt.Sprintf("%db", m)
}

func (m *Memory) UnmarshalString(value string) error {
	badLiteral := fmt.Errorf("bad memory literal: %s", value)

	re := regexp.MustCompile(`^(?P<quantity>\d+)(?P<unit>G|g|M|m|K|k|B|b)$`)
	if !re.Match([]byte(value)) {
		return badLiteral
	}

	matches := re.FindStringSubmatch(value)
	if len(matches) != 3 {
		return badLiteral
	}

	quantity, err := strconv.ParseInt(matches[1], 10, 64)
	if err != nil {
		return badLiteral
	}
	unit := matches[2]

	switch unit {
	case "B", "b":
		*m = Memory(quantity) * Byte
	case "K", "k":
		*m = Memory(quantity) * Kilobyte
	case "M", "m":
		*m = Memory(quantity) * Megabyte
	case "G", "g":
		*m = Memory(quantity) * Gigabyte
	}

	return nil
}

func (m *Memory) UnmarshalYAML(value *yaml.Node) error {
	if err := m.UnmarshalString(value.Value); err != nil {
		return &yaml.TypeError{
			Errors: []string{err.Error()},
		}
	}

	return nil
}

func (m *Memory) MarshalJSON() ([]byte, error) {
	return []byte(m.String()), nil
}

type RuntimeOpts struct {
	InitialMemory Memory
	MaxMemory     Memory
	DebugGC       bool
}

func NewRuntimeOpts() RuntimeOpts {
	return RuntimeOpts{
		InitialMemory: 256 * Megabyte,
		MaxMemory:     2 * Gigabyte,
		DebugGC:       false,
	}
}

type InstallKind string

type InstallOpts struct {
	Kind    InstallKind
	Version string
	Mods    []JarSource
	Plugins []JarSource
}

func NewInstallOpts() InstallOpts {
	return InstallOpts{
		Kind:    VanillaInstall,
		Version: "latest",
		Mods:    []JarSource{},
		Plugins: []JarSource{},
	}
}

type ServerDefinition struct {
	root    string
	Install InstallOpts
	Run     RuntimeOpts
	Configs map[string]config.ConfigDict
}

func NewServerDefinition() ServerDefinition {
	return ServerDefinition{
		root:    "./",
		Install: NewInstallOpts(),
		Run:     NewRuntimeOpts(),
		Configs: make(map[string]config.ConfigDict),
	}
}

func DefinitionFromPath(path string) (ServerDefinition, error) {
	defaultServerDefinition := NewServerDefinition()
	serverDef := defaultServerDefinition

	path, err := filepath.Abs(path)
	if err != nil {
		return serverDef, err
	}
	serverDef.root = path

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return serverDef, err
	}

	if err := yaml.Unmarshal(data, &serverDef); err != nil {
		return serverDef, err
	}

	if serverDef.Install.Mods == nil {
		serverDef.Install.Mods = defaultServerDefinition.Install.Mods
	}

	if serverDef.Install.Plugins == nil {
		serverDef.Install.Plugins = defaultServerDefinition.Install.Plugins
	}

	if serverDef.Configs == nil {
		fmt.Println("hoho")
		serverDef.Configs = defaultServerDefinition.Configs
	}

	return serverDef, nil
}
