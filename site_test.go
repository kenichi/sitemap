package main

import (
	"net/url"
	"testing"
)

func TestSiteRootURL(t *testing.T) {
	const (
		d        = "example.com"
		expected = "http://example.com/"
	)
	site := Site{Domain: d}
	srus := site.RootURL().String()
	if srus != expected {
		t.Error("expected: "+expected, "got: "+srus)
	}
}

func TestSiteAddAsset(t *testing.T) {
	u1, _ := url.Parse("/foo")
	u2, _ := url.Parse("/bar")
	site := Site{Domain: "example.com"}

	if len(site.Assets) > 0 {
		t.Error("site.Assets initialized incorrectly!")
	}

	ret1 := site.AddAsset(u1)
	if len(site.Assets) != 1 {
		t.Error("site.AddAsset did not add correctly!")
	}
	if ret1.Count != 1 {
		t.Error("site.AddAsset inited count incorrectly!")
	}

	ret2 := site.AddAsset(u1)
	if len(site.Assets) != 1 {
		t.Error("site.AddAsset added a dupe!")
	}
	if ret2 != ret1 {
		t.Error("site.AddAsset did not return already added *Asset!")
	}
	if ret2.Count != 2 {
		t.Error("site.AddAsset did not increment count!")
	}

	_ = site.AddAsset(u2)
	if len(site.Assets) != 2 {
		t.Error("site.AddAsset did not add correctly!")
	}
}

func TestSiteAddPage(t *testing.T) {
	p1, _ := url.Parse("/foo")
	p2, _ := url.Parse("/bar")
	site := Site{Domain: "example.com"}

	// prevent descent (site.go:64)
	config.Depth = -1

	if len(site.Pages) > 0 {
		t.Error("site.pages initialized incorrectly!")
	}

	ret1 := site.AddPage(p1)
	if len(site.Pages) != 1 {
		t.Error("site.AddPage did not add correctly!")
	}

	ret2 := site.AddPage(p1)
	if len(site.Pages) != 1 {
		t.Error("site.AddPage added a dupe!")
	}
	if ret2 != ret1 {
		t.Error("site.AddPage did not return already added *Page!")
	}

	_ = site.AddPage(p2)
	if len(site.Pages) != 2 {
		t.Error("site.AddPage did not add correctly!")
	}

	// reset
	config.Depth = defaultDepth
}
