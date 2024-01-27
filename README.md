# Description

A simple Golang library to read World of Warcraft TOC files.

## Usage

```go
is := is.New(t) // Using github.com/matryer/is testing framework
simpleToc := "## Interface: 90001\n## Title: Test Addon\n\nAddon.lua\nAddon.toc\n"

parser := New()
// We can load a file using LoadFile also
parsed := parser.LoadString(simpleToc)

is.True(parsed)
is.True(parser.HasEntry("Interface"))
is.Equal(parser.HasEntry("Helloworld"), false)

interfaceVersion := parser.GetEntry("Interface")
is.Equal(interfaceVersion, "90001")

emptyEntry := parser.GetEntry("Helloworld")
is.Equal(emptyEntry, "")

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
```

