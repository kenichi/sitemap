package main

import (
	"flag"
	"log"
	"net/url"
	"os"
)

const (
	defaultDepth    = 2
	defaultDomain   = "digitalocean.com"
	defaultRootPath = "/"
	defaultScheme   = "http"
	defaultTesting  = false
)

type Config struct {
	Depth    int
	Domain   string
	Force    bool
	RootPath string
	Scheme   string
}

type Asset struct {
	URL   *url.URL
	Site  *Site
	Count int
}

var config Config

func init() {

	// cmd line opts / defaults
	config = Config{}
	flag.IntVar(&config.Depth, "n", defaultDepth, "depth (-1 for infinite)")
	flag.StringVar(&config.Domain, "d", defaultDomain, "domain")
	flag.BoolVar(&config.Force, "f", false, "force")
	flag.StringVar(&config.RootPath, "p", defaultRootPath, "root path")
	flag.StringVar(&config.Scheme, "s", defaultScheme, "scheme")

	// log to
	log.SetOutput(os.Stderr)
}

func main() {

	flag.Parse()

	// len(Split("/foo/", "/")) == 3, but depth 2 meaning 2 dirs
	config.Depth = config.Depth + 1

	site := Site{Domain: config.Domain}
	site.Map()
	site.Print()
}

func Error(err error) {
	log.Println("error!", err)
	if !config.Force {
		os.Exit(1)
	}
}
