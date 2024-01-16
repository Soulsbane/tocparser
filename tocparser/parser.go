package tocparser

import (
	"fmt"
	"os"
	"strings"
)

// Parser Used for parsing WoW TOC files.
type Parser struct {
	values map[string]string
	files  []string
}

// New Creates a new Parser
func New() Parser {
	var parser Parser

	parser.values = make(map[string]string)
	return parser
}

// LoadString Parses a string into key/value pairs split from the ':' character and returns true if successful.
func (parser *Parser) LoadString(content string) bool {
	if len(content) == 0 {
		return false
	}

	lines := strings.Split(content, "\n")

	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)

		if strings.HasPrefix(trimmedLine, "##") {
			values := strings.SplitN(strings.Trim(trimmedLine, "#"), ":", 2)

			// Creates a pair from this example string "## Author: Soulsbane"
			if len(values) == 2 {
				key := strings.Trim(values[0], " ")
				value := strings.Trim(values[1], " ")

				parser.AddEntry(key, value)
			}
		} else {
			// Only add files that are not empty and not commented out.
			if trimmedLine != "" && !strings.HasPrefix(line, "#") {
				parser.files = append(parser.files, line)
			}
		}
	}

	return true
}

// LoadFile Loads a TOC file's contents into a string and calls ParseString
func (parser *Parser) LoadFile(fileName string) error {
	content, err := os.ReadFile(fileName)

	if err != nil {
		return err
	}

	parser.LoadString(string(content))
	return nil
}

// AddEntry Adds a new key/value pair to the parser.
func (parser *Parser) AddEntry(key string, value string) {
	parser.values[key] = value
}

// HasEntry Check if an entry exists.
func (parser *Parser) HasEntry(name string) bool {
	if _, found := parser.values[name]; found {
		return true
	}

	return false
}

// GetEntry Get an entry. Returns an empty string if entry not found.
func (parser *Parser) GetEntry(name string) string {
	if value, found := parser.values[name]; found {
		return value
	}

	return ""
}

// GetEntryOrDefault Get an entry. Returns the default value if entry is not found.
func (parser *Parser) GetEntryOrDefault(name string, defaultValue string) string {
	if value, found := parser.values[name]; found {
		return value
	}

	return defaultValue
}

// GetTitle Gets the title of the addon or an empty string otherwise.
func (parser *Parser) GetTitle() string {
	return parser.GetEntryOrDefault("Title", "")
}

// GetAuthor Gets the author of the addon or an empty string otherwise.
func (parser *Parser) GetAuthor() string {
	return parser.GetEntryOrDefault("Author", "")
}

// GetInterface Gets the WoW interface version of the addon is written for or an empty string otherwise.
func (parser *Parser) GetInterface() string {
	return parser.GetEntryOrDefault("Interface", "")
}

// GetFiles Gets a list of files referenced in the TOC file.
func (parser *Parser) GetFiles() []string {
	return parser.files
}

// GetNumFiles Gets a list of files referenced in the TOC file.
func (parser *Parser) GetNumFiles() int {
	return len(parser.files)
}

// DumpEntries outputs the key/value pairs to stdout.
func (parser *Parser) DumpEntries() {
	for key, value := range parser.values {
		fmt.Printf("%s => %s\n", key, value)
	}
}
