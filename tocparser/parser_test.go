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
	is.True(parser.HasEntry("Interface"))
	is.Equal(parser.HasEntry("Helloworld"), false)
	interfaceVersion := parser.GetEntry("Interface")
	is.Equal(interfaceVersion, "90001")
}
