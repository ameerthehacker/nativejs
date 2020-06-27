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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 20, 130,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 3, 2, 3, 2, 3, 2, 3, 2, 7, 2, 44, 10, 2, 12, 2,
	14, 2, 47, 11, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3,
	7, 3, 58, 10, 3, 12, 3, 14, 3, 61, 11, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5,
	3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 10, 6, 10, 84, 10, 10, 13, 10, 14, 10, 85, 3, 11,
	6, 11, 89, 10, 11, 13, 11, 14, 11, 90, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3,
	15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17,
	3, 17, 3, 18, 6, 18, 121, 10, 18, 13, 18, 14, 18, 122, 3, 18, 3, 18, 3,
	19, 3, 19, 3, 19, 3, 19, 3, 45, 2, 20, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13,
	8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17,
	33, 18, 35, 19, 37, 20, 3, 2, 6, 5, 2, 12, 12, 15, 15, 8234, 8235, 4, 2,
	67, 92, 99, 124, 3, 2, 50, 59, 6, 2, 11, 11, 13, 14, 34, 34, 162, 162,
	2, 134, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3,
	2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17,
	3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2,
	25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2,
	2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 3, 39, 3, 2, 2,
	2, 5, 53, 3, 2, 2, 2, 7, 64, 3, 2, 2, 2, 9, 66, 3, 2, 2, 2, 11, 68, 3,
	2, 2, 2, 13, 70, 3, 2, 2, 2, 15, 74, 3, 2, 2, 2, 17, 76, 3, 2, 2, 2, 19,
	83, 3, 2, 2, 2, 21, 88, 3, 2, 2, 2, 23, 92, 3, 2, 2, 2, 25, 101, 3, 2,
	2, 2, 27, 103, 3, 2, 2, 2, 29, 105, 3, 2, 2, 2, 31, 112, 3, 2, 2, 2, 33,
	115, 3, 2, 2, 2, 35, 120, 3, 2, 2, 2, 37, 126, 3, 2, 2, 2, 39, 40, 7, 49,
	2, 2, 40, 41, 7, 44, 2, 2, 41, 45, 3, 2, 2, 2, 42, 44, 11, 2, 2, 2, 43,
	42, 3, 2, 2, 2, 44, 47, 3, 2, 2, 2, 45, 46, 3, 2, 2, 2, 45, 43, 3, 2, 2,
	2, 46, 48, 3, 2, 2, 2, 47, 45, 3, 2, 2, 2, 48, 49, 7, 44, 2, 2, 49, 50,
	7, 49, 2, 2, 50, 51, 3, 2, 2, 2, 51, 52, 8, 2, 2, 2, 52, 4, 3, 2, 2, 2,
	53, 54, 7, 49, 2, 2, 54, 55, 7, 49, 2, 2, 55, 59, 3, 2, 2, 2, 56, 58, 10,
	2, 2, 2, 57, 56, 3, 2, 2, 2, 58, 61, 3, 2, 2, 2, 59, 57, 3, 2, 2, 2, 59,
	60, 3, 2, 2, 2, 60, 62, 3, 2, 2, 2, 61, 59, 3, 2, 2, 2, 62, 63, 8, 3, 2,
	2, 63, 6, 3, 2, 2, 2, 64, 65, 7, 46, 2, 2, 65, 8, 3, 2, 2, 2, 66, 67, 7,
	60, 2, 2, 67, 10, 3, 2, 2, 2, 68, 69, 7, 63, 2, 2, 69, 12, 3, 2, 2, 2,
	70, 71, 7, 110, 2, 2, 71, 72, 7, 103, 2, 2, 72, 73, 7, 118, 2, 2, 73, 14,
	3, 2, 2, 2, 74, 75, 7, 61, 2, 2, 75, 16, 3, 2, 2, 2, 76, 77, 7, 101, 2,
	2, 77, 78, 7, 113, 2, 2, 78, 79, 7, 112, 2, 2, 79, 80, 7, 117, 2, 2, 80,
	81, 7, 118, 2, 2, 81, 18, 3, 2, 2, 2, 82, 84, 9, 3, 2, 2, 83, 82, 3, 2,
	2, 2, 84, 85, 3, 2, 2, 2, 85, 83, 3, 2, 2, 2, 85, 86, 3, 2, 2, 2, 86, 20,
	3, 2, 2, 2, 87, 89, 9, 4, 2, 2, 88, 87, 3, 2, 2, 2, 89, 90, 3, 2, 2, 2,
	90, 88, 3, 2, 2, 2, 90, 91, 3, 2, 2, 2, 91, 22, 3, 2, 2, 2, 92, 93, 7,
	104, 2, 2, 93, 94, 7, 119, 2, 2, 94, 95, 7, 112, 2, 2, 95, 96, 7, 101,
	2, 2, 96, 97, 7, 118, 2, 2, 97, 98, 7, 107, 2, 2, 98, 99, 7, 113, 2, 2,
	99, 100, 7, 112, 2, 2, 100, 24, 3, 2, 2, 2, 101, 102, 7, 42, 2, 2, 102,
	26, 3, 2, 2, 2, 103, 104, 7, 43, 2, 2, 104, 28, 3, 2, 2, 2, 105, 106, 7,
	112, 2, 2, 106, 107, 7, 119, 2, 2, 107, 108, 7, 111, 2, 2, 108, 109, 7,
	100, 2, 2, 109, 110, 7, 103, 2, 2, 110, 111, 7, 116, 2, 2, 111, 30, 3,
	2, 2, 2, 112, 113, 7, 107, 2, 2, 113, 114, 7, 58, 2, 2, 114, 32, 3, 2,
	2, 2, 115, 116, 7, 107, 2, 2, 116, 117, 7, 51, 2, 2, 117, 118, 7, 56, 2,
	2, 118, 34, 3, 2, 2, 2, 119, 121, 9, 5, 2, 2, 120, 119, 3, 2, 2, 2, 121,
	122, 3, 2, 2, 2, 122, 120, 3, 2, 2, 2, 122, 123, 3, 2, 2, 2, 123, 124,
	3, 2, 2, 2, 124, 125, 8, 18, 2, 2, 125, 36, 3, 2, 2, 2, 126, 127, 9, 2,
	2, 2, 127, 128, 3, 2, 2, 2, 128, 129, 8, 19, 2, 2, 129, 38, 3, 2, 2, 2,
	8, 2, 45, 59, 85, 90, 122, 3, 2, 3, 2,
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
	"", "", "", "','", "':'", "'='", "'let'", "';'", "'const'", "", "", "'function'",
	"'('", "')'", "'number'", "'i8'", "'i16'",
}

var lexerSymbolicNames = []string{
	"", "MultiLineComment", "SingleLineComment", "Comma", "Colon", "EqualSign",
	"Let", "Semicolon", "Const", "Identifier", "NumberLiteral", "Function",
	"OPEN_PARENTHESIS", "CLOSED_PARENTHESIS", "INT", "I8", "I16", "WhiteSpaces",
	"LineTerminator",
}

var lexerRuleNames = []string{
	"MultiLineComment", "SingleLineComment", "Comma", "Colon", "EqualSign",
	"Let", "Semicolon", "Const", "Identifier", "NumberLiteral", "Function",
	"OPEN_PARENTHESIS", "CLOSED_PARENTHESIS", "INT", "I8", "I16", "WhiteSpaces",
	"LineTerminator",
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
	NJSLexerLet                = 6
	NJSLexerSemicolon          = 7
	NJSLexerConst              = 8
	NJSLexerIdentifier         = 9
	NJSLexerNumberLiteral      = 10
	NJSLexerFunction           = 11
	NJSLexerOPEN_PARENTHESIS   = 12
	NJSLexerCLOSED_PARENTHESIS = 13
	NJSLexerINT                = 14
	NJSLexerI8                 = 15
	NJSLexerI16                = 16
	NJSLexerWhiteSpaces        = 17
	NJSLexerLineTerminator     = 18
)
