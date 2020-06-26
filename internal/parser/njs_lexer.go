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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 17, 111,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 3, 2, 3, 2, 3,
	2, 3, 2, 7, 2, 38, 10, 2, 12, 2, 14, 2, 41, 11, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 52, 10, 3, 12, 3, 14, 3, 55, 11,
	3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3,
	7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 6, 10, 78, 10,
	10, 13, 10, 14, 10, 79, 3, 11, 6, 11, 83, 10, 11, 13, 11, 14, 11, 84, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 14,
	3, 14, 3, 14, 3, 14, 3, 15, 6, 15, 102, 10, 15, 13, 15, 14, 15, 103, 3,
	15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 39, 2, 17, 3, 3, 5, 4, 7, 5,
	9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27,
	15, 29, 16, 31, 17, 3, 2, 6, 5, 2, 12, 12, 15, 15, 8234, 8235, 4, 2, 67,
	92, 99, 124, 3, 2, 50, 59, 6, 2, 11, 11, 13, 14, 34, 34, 162, 162, 2, 115,
	2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2,
	2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2,
	2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2,
	2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 3, 33, 3,
	2, 2, 2, 5, 47, 3, 2, 2, 2, 7, 58, 3, 2, 2, 2, 9, 60, 3, 2, 2, 2, 11, 62,
	3, 2, 2, 2, 13, 64, 3, 2, 2, 2, 15, 68, 3, 2, 2, 2, 17, 70, 3, 2, 2, 2,
	19, 77, 3, 2, 2, 2, 21, 82, 3, 2, 2, 2, 23, 86, 3, 2, 2, 2, 25, 93, 3,
	2, 2, 2, 27, 96, 3, 2, 2, 2, 29, 101, 3, 2, 2, 2, 31, 107, 3, 2, 2, 2,
	33, 34, 7, 49, 2, 2, 34, 35, 7, 44, 2, 2, 35, 39, 3, 2, 2, 2, 36, 38, 11,
	2, 2, 2, 37, 36, 3, 2, 2, 2, 38, 41, 3, 2, 2, 2, 39, 40, 3, 2, 2, 2, 39,
	37, 3, 2, 2, 2, 40, 42, 3, 2, 2, 2, 41, 39, 3, 2, 2, 2, 42, 43, 7, 44,
	2, 2, 43, 44, 7, 49, 2, 2, 44, 45, 3, 2, 2, 2, 45, 46, 8, 2, 2, 2, 46,
	4, 3, 2, 2, 2, 47, 48, 7, 49, 2, 2, 48, 49, 7, 49, 2, 2, 49, 53, 3, 2,
	2, 2, 50, 52, 10, 2, 2, 2, 51, 50, 3, 2, 2, 2, 52, 55, 3, 2, 2, 2, 53,
	51, 3, 2, 2, 2, 53, 54, 3, 2, 2, 2, 54, 56, 3, 2, 2, 2, 55, 53, 3, 2, 2,
	2, 56, 57, 8, 3, 2, 2, 57, 6, 3, 2, 2, 2, 58, 59, 7, 46, 2, 2, 59, 8, 3,
	2, 2, 2, 60, 61, 7, 60, 2, 2, 61, 10, 3, 2, 2, 2, 62, 63, 7, 63, 2, 2,
	63, 12, 3, 2, 2, 2, 64, 65, 7, 110, 2, 2, 65, 66, 7, 103, 2, 2, 66, 67,
	7, 118, 2, 2, 67, 14, 3, 2, 2, 2, 68, 69, 7, 61, 2, 2, 69, 16, 3, 2, 2,
	2, 70, 71, 7, 101, 2, 2, 71, 72, 7, 113, 2, 2, 72, 73, 7, 112, 2, 2, 73,
	74, 7, 117, 2, 2, 74, 75, 7, 118, 2, 2, 75, 18, 3, 2, 2, 2, 76, 78, 9,
	3, 2, 2, 77, 76, 3, 2, 2, 2, 78, 79, 3, 2, 2, 2, 79, 77, 3, 2, 2, 2, 79,
	80, 3, 2, 2, 2, 80, 20, 3, 2, 2, 2, 81, 83, 9, 4, 2, 2, 82, 81, 3, 2, 2,
	2, 83, 84, 3, 2, 2, 2, 84, 82, 3, 2, 2, 2, 84, 85, 3, 2, 2, 2, 85, 22,
	3, 2, 2, 2, 86, 87, 7, 112, 2, 2, 87, 88, 7, 119, 2, 2, 88, 89, 7, 111,
	2, 2, 89, 90, 7, 100, 2, 2, 90, 91, 7, 103, 2, 2, 91, 92, 7, 116, 2, 2,
	92, 24, 3, 2, 2, 2, 93, 94, 7, 107, 2, 2, 94, 95, 7, 58, 2, 2, 95, 26,
	3, 2, 2, 2, 96, 97, 7, 107, 2, 2, 97, 98, 7, 51, 2, 2, 98, 99, 7, 56, 2,
	2, 99, 28, 3, 2, 2, 2, 100, 102, 9, 5, 2, 2, 101, 100, 3, 2, 2, 2, 102,
	103, 3, 2, 2, 2, 103, 101, 3, 2, 2, 2, 103, 104, 3, 2, 2, 2, 104, 105,
	3, 2, 2, 2, 105, 106, 8, 15, 2, 2, 106, 30, 3, 2, 2, 2, 107, 108, 9, 2,
	2, 2, 108, 109, 3, 2, 2, 2, 109, 110, 8, 16, 2, 2, 110, 32, 3, 2, 2, 2,
	8, 2, 39, 53, 79, 84, 103, 3, 2, 3, 2,
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
	"", "", "", "','", "':'", "'='", "'let'", "';'", "'const'", "", "", "'number'",
	"'i8'", "'i16'",
}

var lexerSymbolicNames = []string{
	"", "MultiLineComment", "SingleLineComment", "Comma", "Colon", "EqualSign",
	"Let", "Semicolon", "Const", "Identifier", "NumberLiteral", "INT", "I8",
	"I16", "WhiteSpaces", "LineTerminator",
}

var lexerRuleNames = []string{
	"MultiLineComment", "SingleLineComment", "Comma", "Colon", "EqualSign",
	"Let", "Semicolon", "Const", "Identifier", "NumberLiteral", "INT", "I8",
	"I16", "WhiteSpaces", "LineTerminator",
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
	NJSLexerMultiLineComment  = 1
	NJSLexerSingleLineComment = 2
	NJSLexerComma             = 3
	NJSLexerColon             = 4
	NJSLexerEqualSign         = 5
	NJSLexerLet               = 6
	NJSLexerSemicolon         = 7
	NJSLexerConst             = 8
	NJSLexerIdentifier        = 9
	NJSLexerNumberLiteral     = 10
	NJSLexerINT               = 11
	NJSLexerI8                = 12
	NJSLexerI16               = 13
	NJSLexerWhiteSpaces       = 14
	NJSLexerLineTerminator    = 15
)
