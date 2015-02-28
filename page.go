package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"net/http"
	"net/url"
)

type Page struct {
	URL    *url.URL
	Pages  []*Page
	Assets []*Asset
	Site   *Site
}

// pass body reader from Fetch() to Parse()
func (page *Page) Load() (e error) {
	fmt.Print(".")
	if body, err := page.Fetch(); err != nil {
		e = err
	} else {
		if body != nil {
			page.Parse(body)
		}
	}
	page.Site.WaitGroup.Done()
	return
}

// return reader of html for page
func (page *Page) Fetch() (body io.ReadCloser, e error) {
	if resp, err := http.Get(page.URL.String()); err != nil {
		e = err
	} else {
		if resp.StatusCode == 200 {
			body = resp.Body
		}
	}
	return
}

// parse up the html from the reader
func (page *Page) Parse(body io.ReadCloser) (e error) {
	doc, err := html.Parse(body)
	body.Close()
	if err != nil {
		e = err
	} else {
		page.Node(doc)
	}
	return
}

// iterate element nodes
// https://godoc.org/golang.org/x/net/html
func (page *Page) Node(n *html.Node) {

	// handler for assets (css, js, images)
	var handleAsset func(url *url.URL) = func(url *url.URL) {
		asset := page.Site.AddAsset(url)
		page.AddAsset(asset)
	}

	if n.Type == html.ElementNode {
		switch n.Data {
		case "a":
			// handler for links
			page.HandleNode(n, "href", func(url *url.URL) {
				link := page.Site.AddPage(url)
				page.AddPage(link)
			})
		case "link":
			page.HandleNode(n, "href", handleAsset)
		case "img", "script":
			page.HandleNode(n, "src", handleAsset)
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		page.Node(c)
	}
}

// pull out value for node attr key, build/sanitize url, call handler if valid
func (page *Page) HandleNode(n *html.Node, key string, handler func(url *url.URL)) {
	for _, attr := range n.Attr {
		if attr.Key == key {

			if url, err := url.Parse(attr.Val); err != nil {
				fmt.Println("error handling "+key+" value:", err)
			} else {
				if url.IsAbs() {

					// only this domain
					if url.Host == page.Site.Domain {
						url.RawQuery = ""
						url.Fragment = ""

						handler(url)
					}

				} else {

					// qualify
					url.Host = page.Site.Domain
					url.Scheme = defaultScheme
					url.RawQuery = ""
					url.Fragment = ""

					// don't dupe /
					if url.Path == "" {
						url.Path = "/"
					}

					handler(url)

				}
			}

		}
	}
}

// add dep on asset, true if added, false if exists
func (page *Page) AddAsset(asset *Asset) bool {
	for _, a := range page.Assets {
		if a == asset {
			return false
		}
	}
	page.Assets = append(page.Assets, asset)
	return true
}

// add link to page, true if added, false if exists
func (page *Page) AddPage(link *Page) bool {
	for _, l := range page.Pages {
		if l == link {
			return false
		}
	}
	page.Pages = append(page.Pages, link)
	return true
}
