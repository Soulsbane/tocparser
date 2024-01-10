package tocparser

import "strings"

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

	lines := strings.Split(string(content), "\n")

	for _, line := range lines {
		line := strings.TrimSpace(line)

		if strings.HasPrefix(line, "##") {
			line := strings.Trim(line, "#")
			values := strings.SplitN(line, ":", 2)

			// Creates a pair from this example string "## Author: Soulsbane"
			if len(values) == 2 {
				key := strings.Trim(values[0], " ")
				value := strings.Trim((values[1]), " ")
				parser.values[key] = value
			}
			// Line is a comment
		} else if len(line) == 0 || (strings.HasPrefix(line, "#") && !strings.Contains(line, ":")) {
			continue
			// Line is a empty or a filename. If blank ignore.
		} else {
			if strings.TrimSpace(line) != "" {
				parser.files = append(parser.files, line)
			}
		}
	}

	return true
}

// HasEntry Check if an entry exists.
func (parser Parser) HasEntry(name string) bool {
	if _, found := parser.values[name]; found {
		return true
	}

	return false
}

// GetEntry Get an entry. Returns an empty string if entry not found.
func (parser Parser) GetEntry(name string) string {
	if value, found := parser.values[name]; found {
		return value
	}

	return ""
}

// GetEntry Get an entry. Returns the default value if entry is not found.
func (parser Parser) GetEntryOrDefault(name string, defaultValue string) string {
	if value, found := parser.values[name]; found {
		return value
	}

	return defaultValue
}
