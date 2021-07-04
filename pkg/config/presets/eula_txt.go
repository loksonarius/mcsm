package presets

import (
	"io/ioutil"
	"path/filepath"

	"github.com/loksonarius/mcsm/pkg/config"
	"github.com/loksonarius/mcsm/pkg/config/types/properties"
)

type EulaTxt struct {
	// look, honestly, I've no clue /why/ someone would set this to false
	Accepted bool `properties:"key:eula,default:true"`
}

func EulaTxtFromConfig(configs map[string]config.ConfigDict) config.ConfigFile {
	cfg := config.ConfigDict{}
	if c, ok := configs["eula"]; ok {
		cfg = c
	}

	var e EulaTxt
	properties.Unmarshal(cfg, &e)
	return &e
}

func (p *EulaTxt) Path() string {
	return "eula.txt"
}

func (p *EulaTxt) Validate() error {
	// nothing can go wrong with type checking!!!
	return nil
}

func (p *EulaTxt) Render() []byte {
	out, _ := properties.Marshal(p)
	return out
}

func (p *EulaTxt) Write() error {
	path, err := filepath.Abs(p.Path())
	if err != nil {
		return err
	}

	return ioutil.WriteFile(path, p.Render(), 0644)
}
