package tocparser

import (
	"github.com/matryer/is"
	"testing"
)

func TestParser_LoadString(t *testing.T) {
	is := is.New(t)
	simpleToc := "## Interface: 90001\n## Title: Test Addon\n\nAddon.lua\nAddon.toc\n"

	parser := New()
	parsed := parser.LoadString(simpleToc)
	is.True(parsed)
}
