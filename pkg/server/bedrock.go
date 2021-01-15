package server

import (
	"regexp"
	"strings"

	"golang.org/x/net/html"

	"github.com/loksonarius/mcsm/pkg/config"
	"github.com/loksonarius/mcsm/pkg/config/presets"
)

const BedrockInstall = InstallKind("bedrock")

type BedrockServer struct {
	Definition       ServerDefinition
	ServerBinaryPath string
	Configs          []config.ConfigFile
}

func NewBedrockServer(def ServerDefinition) Server {
	return &BedrockServer{
		Definition:       def,
		ServerBinaryPath: "bedrock_server",
		Configs: []config.ConfigFile{
			presets.ServerPropertiesFromConfig(def.Configs),
			presets.EulaTxtFromConfig(def.Configs),
		},
	}
}

func init() {
	registerServer(BedrockInstall, NewBedrockServer)
}

func (bs *BedrockServer) Install() error {
	return nil
}

func (bs *BedrockServer) Configure() error {
	return nil
}

func (bs *BedrockServer) Versions() ([]string, error) {
	var versions []string
	addr := "https://minecraft.gamepedia.com/Bedrock_Edition_version_history"
	body, err := httpGet(addr)
	if err != nil {
		return versions, err
	}
	defer body.Close()

	tables := []*html.Node{}
	var crawlForVersionTables func(n *html.Node)
	crawlForVersionTables = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "table" {
			for _, a := range n.Attr {
				if a.Key == "data-description" {
					if strings.Contains(a.Val, "version history") {
						tables = append(tables, n)
					}
				}
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			crawlForVersionTables(c)
		}
	}

	var crawlForVersions func(n *html.Node)
	versionRe := regexp.MustCompile(`^\/Bedrock_Edition_(?P<version>\d+.\d+.\d+(.\d+)?)$`)
	crawlForVersions = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" && versionRe.MatchString(a.Val) {
					matches := versionRe.FindStringSubmatch(a.Val)
					if len(matches) == versionRe.NumSubexp()+1 {
						version := matches[versionRe.SubexpIndex("version")]
						versions = append(versions, version)
					}
				}
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			crawlForVersions(c)
		}
	}

	doc, err := html.Parse(body)
	if err != nil {
		return versions, err
	}

	crawlForVersionTables(doc)
	for _, t := range tables {
		crawlForVersions(t)
	}

	return versions, nil
}

func (bs *BedrockServer) Run() error {
	return nil
}
