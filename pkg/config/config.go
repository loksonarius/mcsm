package config

type ConfigDict map[interface{}]interface{}

type ConfigFile interface {
	Path() string
	Validate() error
	Render() []byte
	Write() error
}
