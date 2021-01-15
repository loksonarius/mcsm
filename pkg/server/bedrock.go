package server

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
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
	versions, err := bs.Versions()
	if err != nil {
		return err
	}

	// assume we want latest version
	targetVersion := versions[len(versions)-1]
	reqVersion := bs.Definition.Install.Version
	if reqVersion != "latest" {
		found := false
		for _, v := range versions {
			if v == reqVersion {
				targetVersion = v
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("no bedrock server version %s", reqVersion)
		}
	}

	versionURL := fmt.Sprintf(
		"https://minecraft.azureedge.net/bin-linux/bedrock-server-%s.zip",
		targetVersion,
	)
	fmt.Printf("downloading from %s\n", versionURL)

	zipDir := filepath.Dir(bs.ServerBinaryPath)
	zipPath := filepath.Join(zipDir, "bedrock.zip")
	if err := downloadFileToPath(versionURL, zipPath); err != nil {
		return err
	}

	r, err := zip.OpenReader(zipPath)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		path, err := filepath.Abs(filepath.Join(zipDir, f.Name))
		if err != nil {
			return err
		}

		out, err := os.Create(path)
		if err != nil {
			return err
		}
		defer out.Close()

		_, err = io.Copy(out, rc)
		if err != nil {
			return err
		}
	}

	return nil
}

func (bs *BedrockServer) Configure() error {
	return nil
}

func (bs *BedrockServer) getGamepediaListedVersions() ([]string, error) {
	var versions []string
	baseAddr := "https://minecraft.gamepedia.com"
	addr := baseAddr + "/Bedrock_Edition_version_history"
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

	urlRe := regexp.MustCompile(`^\/Bedrock_Edition_(?P<version>\d+\.\d+\.\d+(\.\d+)?)$`)
	var versionUrls []string
	var crawlForVersionUrls func(n *html.Node)
	crawlForVersionUrls = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" && urlRe.MatchString(a.Val) {
					versionUrls = append(versionUrls, baseAddr+a.Val)
				}
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			crawlForVersionUrls(c)
		}
	}

	vmap := make(map[string]bool)
	versionRe := regexp.MustCompile(`^\d+\.\d+\.\d+\.\d+$`)
	var crawlForServerVersion func(n *html.Node)
	crawlForServerVersion = func(n *html.Node) {
		d := strings.TrimSpace(n.Data)
		if n.Type == html.TextNode && versionRe.MatchString(d) {
			vmap[d] = true
			return
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			crawlForServerVersion(c)
		}
	}

	doc, err := html.Parse(body)
	if err != nil {
		return versions, err
	}

	crawlForVersionTables(doc)
	for _, t := range tables {
		crawlForVersionUrls(t)
	}

	for _, u := range versionUrls {
		body, err := httpGet(u)
		if err != nil {
			return versions, err
		}
		defer body.Close()

		if doc, err := html.Parse(body); err == nil {
			crawlForServerVersion(doc)
		}
	}

	for k := range vmap {
		versions = append(versions, k)
	}

	return versions, nil
}

func (bs *BedrockServer) getMojangLatestVersion() (string, error) {
	var version string
	addr := "https://www.minecraft.net/en-us/download/server/bedrock"
	body, err := httpGet(addr)
	if err != nil {
		return version, err
	}
	defer body.Close()

	var downloadBtn *html.Node
	var crawl func(n *html.Node)
	crawl = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "data-platform" && a.Val == "serverBedrockLinux" {
					downloadBtn = n
					return
				}
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			crawl(c)
		}
	}

	doc, err := html.Parse(body)
	if err != nil {
		return version, err
	}

	crawl(doc)

	if downloadBtn == nil {
		return version, fmt.Errorf("failed finding latest bedrock version")
	}

	urlRe := regexp.MustCompile(`^.*bedrock-server-(?P<version>\d+.\d+.\d+(.\d+)?)`)
	for _, a := range downloadBtn.Attr {
		if a.Key != "href" || !urlRe.MatchString(a.Val) {
			continue
		}

		matches := urlRe.FindStringSubmatch(a.Val)
		if len(matches) == urlRe.NumSubexp()+1 {
			version = matches[urlRe.SubexpIndex("version")]
			return version, nil
		}
	}

	return version, fmt.Errorf("failed to find latest bedrock server version")
}

func (bs *BedrockServer) Versions() ([]string, error) {
	versions := []string{}

	versions, err := bs.getGamepediaListedVersions()
	if err != nil {
		return versions, err
	}

	// latest, err := bs.getMojangLatestVersion()
	// if err != nil {
	// 	return versions, err
	// }
	//
	// skip := true
	// for _, v := range gpVersions {
	// 	fmt.Printf("looking at %s vs %s\n", v, latest)
	// 	if strings.Contains(latest, v) {
	// 		skip = false
	// 	}
	//
	// 	if skip {
	// 		continue
	// 	}
	//
	// 	versions = append(versions, v)
	// }
	//
	if len(versions) == 0 {
		return versions, fmt.Errorf("failed to find released bedrock versions")
	}

	sort.Slice(versions, func(i, j int) bool {
		a := strings.Split(versions[i], ".")
		b := strings.Split(versions[j], ".")

		if len(b) > len(a) {
			c := a
			a = b
			b = c
		}

		for k := range a {
			if k >= len(b) {
				return true
			}

			ad, aerr := strconv.Atoi(a[k])
			bd, berr := strconv.Atoi(b[k])

			if aerr != nil || berr != nil {
				continue
			}

			if ad == bd {
				continue
			}

			return ad < bd
		}

		return false
	})

	return versions, nil
}

func (bs *BedrockServer) Run() error {
	return nil
}
