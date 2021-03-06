// Code generated from internal/parser/NJSLexer.g4 by ANTLR 4.8. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 22, 138,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 3, 2, 3, 2, 3, 2,
	3, 2, 7, 2, 48, 10, 2, 12, 2, 14, 2, 51, 11, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 62, 10, 3, 12, 3, 14, 3, 65, 11, 3,
	3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8,
	3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3,
	13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14,
	3, 14, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3,
	15, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 18, 6, 18, 119,
	10, 18, 13, 18, 14, 18, 120, 3, 19, 6, 19, 124, 10, 19, 13, 19, 14, 19,
	125, 3, 20, 6, 20, 129, 10, 20, 13, 20, 14, 20, 130, 3, 20, 3, 20, 3, 21,
	3, 21, 3, 21, 3, 21, 3, 49, 2, 22, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8,
	15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17,
	33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 3, 2, 6, 5, 2, 12, 12, 15, 15,
	8234, 8235, 3, 2, 50, 59, 4, 2, 67, 92, 99, 124, 6, 2, 11, 11, 13, 14,
	34, 34, 162, 162, 2, 142, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3,
	2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15,
	3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2,
	23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2,
	2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2,
	2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 3, 43, 3, 2, 2, 2, 5, 57, 3, 2,
	2, 2, 7, 68, 3, 2, 2, 2, 9, 70, 3, 2, 2, 2, 11, 72, 3, 2, 2, 2, 13, 74,
	3, 2, 2, 2, 15, 76, 3, 2, 2, 2, 17, 78, 3, 2, 2, 2, 19, 80, 3, 2, 2, 2,
	21, 82, 3, 2, 2, 2, 23, 84, 3, 2, 2, 2, 25, 88, 3, 2, 2, 2, 27, 94, 3,
	2, 2, 2, 29, 103, 3, 2, 2, 2, 31, 110, 3, 2, 2, 2, 33, 113, 3, 2, 2, 2,
	35, 118, 3, 2, 2, 2, 37, 123, 3, 2, 2, 2, 39, 128, 3, 2, 2, 2, 41, 134,
	3, 2, 2, 2, 43, 44, 7, 49, 2, 2, 44, 45, 7, 44, 2, 2, 45, 49, 3, 2, 2,
	2, 46, 48, 11, 2, 2, 2, 47, 46, 3, 2, 2, 2, 48, 51, 3, 2, 2, 2, 49, 50,
	3, 2, 2, 2, 49, 47, 3, 2, 2, 2, 50, 52, 3, 2, 2, 2, 51, 49, 3, 2, 2, 2,
	52, 53, 7, 44, 2, 2, 53, 54, 7, 49, 2, 2, 54, 55, 3, 2, 2, 2, 55, 56, 8,
	2, 2, 2, 56, 4, 3, 2, 2, 2, 57, 58, 7, 49, 2, 2, 58, 59, 7, 49, 2, 2, 59,
	63, 3, 2, 2, 2, 60, 62, 10, 2, 2, 2, 61, 60, 3, 2, 2, 2, 62, 65, 3, 2,
	2, 2, 63, 61, 3, 2, 2, 2, 63, 64, 3, 2, 2, 2, 64, 66, 3, 2, 2, 2, 65, 63,
	3, 2, 2, 2, 66, 67, 8, 3, 2, 2, 67, 6, 3, 2, 2, 2, 68, 69, 7, 46, 2, 2,
	69, 8, 3, 2, 2, 2, 70, 71, 7, 60, 2, 2, 71, 10, 3, 2, 2, 2, 72, 73, 7,
	63, 2, 2, 73, 12, 3, 2, 2, 2, 74, 75, 7, 61, 2, 2, 75, 14, 3, 2, 2, 2,
	76, 77, 7, 42, 2, 2, 77, 16, 3, 2, 2, 2, 78, 79, 7, 43, 2, 2, 79, 18, 3,
	2, 2, 2, 80, 81, 7, 125, 2, 2, 81, 20, 3, 2, 2, 2, 82, 83, 7, 127, 2, 2,
	83, 22, 3, 2, 2, 2, 84, 85, 7, 110, 2, 2, 85, 86, 7, 103, 2, 2, 86, 87,
	7, 118, 2, 2, 87, 24, 3, 2, 2, 2, 88, 89, 7, 101, 2, 2, 89, 90, 7, 113,
	2, 2, 90, 91, 7, 112, 2, 2, 91, 92, 7, 117, 2, 2, 92, 93, 7, 118, 2, 2,
	93, 26, 3, 2, 2, 2, 94, 95, 7, 104, 2, 2, 95, 96, 7, 119, 2, 2, 96, 97,
	7, 112, 2, 2, 97, 98, 7, 101, 2, 2, 98, 99, 7, 118, 2, 2, 99, 100, 7, 107,
	2, 2, 100, 101, 7, 113, 2, 2, 101, 102, 7, 112, 2, 2, 102, 28, 3, 2, 2,
	2, 103, 104, 7, 112, 2, 2, 104, 105, 7, 119, 2, 2, 105, 106, 7, 111, 2,
	2, 106, 107, 7, 100, 2, 2, 107, 108, 7, 103, 2, 2, 108, 109, 7, 116, 2,
	2, 109, 30, 3, 2, 2, 2, 110, 111, 7, 107, 2, 2, 111, 112, 7, 58, 2, 2,
	112, 32, 3, 2, 2, 2, 113, 114, 7, 107, 2, 2, 114, 115, 7, 51, 2, 2, 115,
	116, 7, 56, 2, 2, 116, 34, 3, 2, 2, 2, 117, 119, 9, 3, 2, 2, 118, 117,
	3, 2, 2, 2, 119, 120, 3, 2, 2, 2, 120, 118, 3, 2, 2, 2, 120, 121, 3, 2,
	2, 2, 121, 36, 3, 2, 2, 2, 122, 124, 9, 4, 2, 2, 123, 122, 3, 2, 2, 2,
	124, 125, 3, 2, 2, 2, 125, 123, 3, 2, 2, 2, 125, 126, 3, 2, 2, 2, 126,
	38, 3, 2, 2, 2, 127, 129, 9, 5, 2, 2, 128, 127, 3, 2, 2, 2, 129, 130, 3,
	2, 2, 2, 130, 128, 3, 2, 2, 2, 130, 131, 3, 2, 2, 2, 131, 132, 3, 2, 2,
	2, 132, 133, 8, 20, 2, 2, 133, 40, 3, 2, 2, 2, 134, 135, 9, 2, 2, 2, 135,
	136, 3, 2, 2, 2, 136, 137, 8, 21, 2, 2, 137, 42, 3, 2, 2, 2, 8, 2, 49,
	63, 120, 125, 130, 3, 2, 3, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "", "", "','", "':'", "'='", "';'", "'('", "')'", "'{'", "'}'", "'let'",
	"'const'", "'function'", "'number'", "'i8'", "'i16'",
}

var lexerSymbolicNames = []string{
	"", "MultiLineComment", "SingleLineComment", "Comma", "Colon", "EqualSign",
	"Semicolon", "OPEN_PARENTHESIS", "CLOSED_PARENTHESIS", "OPEN_CURLY", "CLOSED_CURLY",
	"Let", "Const", "Function", "INT", "I8", "I16", "NumberLiteral", "Identifier",
	"WhiteSpaces", "LineTerminator",
}

var lexerRuleNames = []string{
	"MultiLineComment", "SingleLineComment", "Comma", "Colon", "EqualSign",
	"Semicolon", "OPEN_PARENTHESIS", "CLOSED_PARENTHESIS", "OPEN_CURLY", "CLOSED_CURLY",
	"Let", "Const", "Function", "INT", "I8", "I16", "NumberLiteral", "Identifier",
	"WhiteSpaces", "LineTerminator",
}

type NJSLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewNJSLexer(input antlr.CharStream) *NJSLexer {

	l := new(NJSLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "NJSLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// NJSLexer tokens.
const (
	NJSLexerMultiLineComment   = 1
	NJSLexerSingleLineComment  = 2
	NJSLexerComma              = 3
	NJSLexerColon              = 4
	NJSLexerEqualSign          = 5
	NJSLexerSemicolon          = 6
	NJSLexerOPEN_PARENTHESIS   = 7
	NJSLexerCLOSED_PARENTHESIS = 8
	NJSLexerOPEN_CURLY         = 9
	NJSLexerCLOSED_CURLY       = 10
	NJSLexerLet                = 11
	NJSLexerConst              = 12
	NJSLexerFunction           = 13
	NJSLexerINT                = 14
	NJSLexerI8                 = 15
	NJSLexerI16                = 16
	NJSLexerNumberLiteral      = 17
	NJSLexerIdentifier         = 18
	NJSLexerWhiteSpaces        = 19
	NJSLexerLineTerminator     = 20
)
