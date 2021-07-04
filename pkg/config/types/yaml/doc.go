// Package yaml exists to internally provide any YAML parsing and rendering for
// configs. It currently just re-exports modified Unmarshal and Marshal
// functions from gopkg.in/yaml.v3 module. If any handling has to be customized
// when rendering config files or parsing config dictionaries, then it will be
// implemented here.
package yaml
