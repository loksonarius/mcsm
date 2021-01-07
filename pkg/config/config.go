package config

type ConfigDict map[string]interface{}

type ConfigFile interface {
	Path() string
	Validate() error
	Render() []byte
	Write() error
}
