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

	hello := parser.GetEntryOrDefault("Helloworld", "Hello")
	is.Equal(hello, "Hello")

	files := parser.GetFiles()
	is.Equal(len(files), 2)
	is.Equal(parser.GetNumFiles(), 2)

	parser.AddEntry("Test", "Test")
	is.Equal(parser.GetEntry("Test"), "Test")
}
