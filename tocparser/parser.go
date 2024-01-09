package tocparser

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
