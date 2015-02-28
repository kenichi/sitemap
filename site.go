package main

import (
	"fmt"
	"net/url"
	"strings"
	"sync"
)

// init with Domain, then Map()
type Site struct {
	Domain    string
	Assets    []*Asset
	Pages     []*Page
	WaitGroup sync.WaitGroup
}

// full URL struct to start mapping from
func (site *Site) RootURL() *url.URL {
	rootURL := url.URL{Scheme: config.Scheme, Host: site.Domain, Path: config.RootPath}
	return &rootURL
}

// map the site
func (site *Site) Map() {

	// root page
	site.Pages = []*Page{&Page{URL: site.RootURL(), Site: site}}

	// page.Load() calls WaitGroup.Done()
	site.WaitGroup.Add(1)
	go LoadPage(site.Pages[0])

	site.WaitGroup.Wait()

	// after the dots
	fmt.Print("\n")
}

// register assetURL, return *Asset of existing or new
func (site *Site) AddAsset(assetURL *url.URL) *Asset {
	for _, a := range site.Assets {
		if *assetURL == *a.URL {
			a.Count++
			return a
		}
	}
	asset := &Asset{URL: assetURL, Site: site, Count: 1}
	site.Assets = append(site.Assets, asset)
	return asset
}

// register pageURL
// launch goroutine to load if new and not too deep
// return *Page of existing or new
func (site *Site) AddPage(pageURL *url.URL) *Page {
	for _, p := range site.Pages {
		if *pageURL == *p.URL {
			return p
		}
	}
	page := &Page{URL: pageURL, Site: site}
	site.Pages = append(site.Pages, page)
	if config.Depth == 0 || len(strings.Split(page.URL.Path, "/")) <= config.Depth {
		site.WaitGroup.Add(1)
		go LoadPage(page)
	}
	return page
}

// display final map of site
func (site *Site) Print() {
	for _, page := range site.Pages {
		if len(page.Assets) > 0 || len(page.Pages) > 0 {
			fmt.Println(page.URL.String())
			fmt.Println("\tstatic assets:")
			for _, asset := range page.Assets {
				fmt.Println("\t\t" + asset.URL.String())
			}
			fmt.Println("\tlinks:")
			for _, page := range page.Pages {
				fmt.Println("\t\t" + page.URL.String())
			}
		}
	}
	fmt.Println("asset dependency counts:")
	for _, asset := range site.Assets {
		fmt.Printf("%d\t%s\n", asset.Count, asset.URL.String())
	}
}

// load the page, handling error if needed
func LoadPage(page *Page) {
	if err := page.Load(); err != nil {
		Error(err)
	}
}
