package parser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// func Parse parses the source code and returns the parser
func NewParser(source string) *NJSParser {
	inputStream := antlr.NewInputStream(source)

	// Create the Lexer
	lexer := NewNJSLexer(inputStream)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	return NewNJSParser(stream)
}
