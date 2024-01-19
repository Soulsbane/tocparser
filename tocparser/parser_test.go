package tocparser

import (
	"github.com/matryer/is"
	"testing"
)

func TestParser_EmptyLoadString(t *testing.T) {
	is := is.New(t)
	simpleToc := ""

	parser := New()
	parsed := parser.LoadString(simpleToc)
	is.Equal(parsed, false)
}

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

	is.Equal(parser.GetNumFiles(), 2)
	is.Equal(files[0], "Addon.lua")
	is.Equal(files[1], "Addon.toc")

	parser.AddEntry("Test", "Test")
	is.Equal(parser.GetEntry("Test"), "Test")

	is.Equal(parser.GetAuthor(), "")
	parser.AddEntry("Author", "Soulsbane")
	is.Equal(parser.GetAuthor(), "Soulsbane")

	is.Equal(parser.GetTitle(), "Test Addon")
	is.Equal(parser.GetInterface(), "90001")
}

func TestParser_LoadFile(t *testing.T) {
	is := is.New(t)
	parser := New()

	err := parser.LoadFile("Sample.toc")

	if err != nil {
		t.Fatal(err)
	}

	is.Equal(parser.GetTitle(), "Sample")
	is.Equal(parser.GetInterface(), "100200")
	is.Equal(parser.GetAuthor(), "Soulsbane")

	files := parser.GetFiles()

	is.Equal(parser.GetNumFiles(), 3)
	is.Equal(files[0], "Libs\\Libs.xml")
	is.Equal(files[1], "Sample.lua")
	is.Equal(files[2], "Config.lua")
}
