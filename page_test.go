package main

import (
	"testing"
)

func TestAddAsset(t *testing.T) {
	p := Page{}
	a1 := &Asset{}
	a2 := &Asset{}

	if len(p.Assets) != 0 {
		t.Error("page.Assets initialized incorrectly!")
	}

	ret1 := p.AddAsset(a1)
	if len(p.Assets) != 1 {
		t.Error("page.AddAssets did not add correctly!")
	}
	if !ret1 {
		t.Error("page.AddAssets did not return correctly!")
	}

	ret2 := p.AddAsset(a1)
	if len(p.Assets) != 1 {
		t.Error("page.AddAssets added a dupe!")
	}
	if ret2 {
		t.Error("page.AddAssets did not return correctly!")
	}

	_ = p.AddAsset(a2)
	if len(p.Assets) != 2 {
		t.Error("page.AddAssets did not add correctly!")
	}
}

func TestAddPage(t *testing.T) {
	p := Page{}
	p1 := &Page{}
	p2 := &Page{}

	if len(p.Pages) != 0 {
		t.Error("page.Pages initialized incorrectly!")
	}

	ret1 := p.AddPage(p1)
	if len(p.Pages) != 1 {
		t.Error("page.AddPages did not add correctly!")
	}
	if !ret1 {
		t.Error("page.AddPages did not return correctly!")
	}

	ret2 := p.AddPage(p1)
	if len(p.Pages) != 1 {
		t.Error("page.AddPages added a dupe!")
	}
	if ret2 {
		t.Error("page.AddPages did not return correctly!")
	}

	_ = p.AddPage(p2)
	if len(p.Pages) != 2 {
		t.Error("page.AddPages did not add correctly!")
	}
}
