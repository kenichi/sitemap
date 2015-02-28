package main

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"strings"
	"testing"
)

func TestConfig(t *testing.T) {

	if defaultDepth != config.Depth {
		t.Error(fmt.Sprintf("depth not set %d != %d", defaultDepth, config.Depth))
	}

	if defaultDomain != config.Domain {
		t.Error(fmt.Sprintf("domain not set %s != %s", defaultDomain, config.Domain))
	}

	if config.Force {
		t.Error("force by default?")
	}

	if defaultRootPath != config.RootPath {
		t.Error(fmt.Sprintf("root path not set %s != %s", defaultRootPath, config.RootPath))
	}

	if defaultScheme != config.Scheme {
		t.Error(fmt.Sprintf("scheme not set %s != %s", defaultScheme, config.Scheme))
	}
}

func TestError(t *testing.T) {
	config.Force = true
	l := new(bytes.Buffer)
	log.SetOutput(l)
	Error(errors.New("foo"))
	if strings.HasSuffix(l.String(), "error! foo") {
		t.Error(fmt.Sprintf("Error() broked: %s", l.String()))
	}
}
