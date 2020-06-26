// Code generated from internal/parser/NJSParser.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // NJSParser

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 16, 27, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 3, 2, 3, 2, 3, 3, 3, 3, 3,
	3, 3, 4, 3, 4, 3, 4, 5, 4, 19, 10, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3,
	5, 3, 5, 2, 2, 6, 2, 4, 6, 8, 2, 3, 3, 2, 13, 14, 2, 23, 2, 10, 3, 2, 2,
	2, 4, 12, 3, 2, 2, 2, 6, 15, 3, 2, 2, 2, 8, 24, 3, 2, 2, 2, 10, 11, 9,
	2, 2, 2, 11, 3, 3, 2, 2, 2, 12, 13, 7, 6, 2, 2, 13, 14, 5, 2, 2, 2, 14,
	5, 3, 2, 2, 2, 15, 16, 7, 8, 2, 2, 16, 18, 7, 11, 2, 2, 17, 19, 5, 4, 3,
	2, 18, 17, 3, 2, 2, 2, 18, 19, 3, 2, 2, 2, 19, 20, 3, 2, 2, 2, 20, 21,
	7, 7, 2, 2, 21, 22, 7, 12, 2, 2, 22, 23, 7, 9, 2, 2, 23, 7, 3, 2, 2, 2,
	24, 25, 5, 6, 4, 2, 25, 9, 3, 2, 2, 2, 3, 18,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "", "", "','", "':'", "'='", "'let'", "';'", "'const'", "", "", "'number'",
	"'i8'",
}
var symbolicNames = []string{
	"", "MultiLineComment", "SingleLineComment", "Comma", "Colon", "EqualSign",
	"Let", "Semicolon", "Const", "Identifier", "NumberLiteral", "INT", "I8",
	"WhiteSpaces", "LineTerminator",
}

var ruleNames = []string{
	"types", "typeAnnotation", "variableDeclaration", "program",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type NJSParser struct {
	*antlr.BaseParser
}

func NewNJSParser(input antlr.TokenStream) *NJSParser {
	this := new(NJSParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "NJSParser.g4"

	return this
}

// NJSParser tokens.
const (
	NJSParserEOF               = antlr.TokenEOF
	NJSParserMultiLineComment  = 1
	NJSParserSingleLineComment = 2
	NJSParserComma             = 3
	NJSParserColon             = 4
	NJSParserEqualSign         = 5
	NJSParserLet               = 6
	NJSParserSemicolon         = 7
	NJSParserConst             = 8
	NJSParserIdentifier        = 9
	NJSParserNumberLiteral     = 10
	NJSParserINT               = 11
	NJSParserI8                = 12
	NJSParserWhiteSpaces       = 13
	NJSParserLineTerminator    = 14
)

// NJSParser rules.
const (
	NJSParserRULE_types               = 0
	NJSParserRULE_typeAnnotation      = 1
	NJSParserRULE_variableDeclaration = 2
	NJSParserRULE_program             = 3
)

// ITypesContext is an interface to support dynamic dispatch.
type ITypesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypesContext differentiates from other interfaces.
	IsTypesContext()
}

type TypesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypesContext() *TypesContext {
	var p = new(TypesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NJSParserRULE_types
	return p
}

func (*TypesContext) IsTypesContext() {}

func NewTypesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypesContext {
	var p = new(TypesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NJSParserRULE_types

	return p
}

func (s *TypesContext) GetParser() antlr.Parser { return s.parser }

func (s *TypesContext) INT() antlr.TerminalNode {
	return s.GetToken(NJSParserINT, 0)
}

func (s *TypesContext) I8() antlr.TerminalNode {
	return s.GetToken(NJSParserI8, 0)
}

func (s *TypesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NJSParserListener); ok {
		listenerT.EnterTypes(s)
	}
}

func (s *TypesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NJSParserListener); ok {
		listenerT.ExitTypes(s)
	}
}

func (p *NJSParser) Types() (localctx ITypesContext) {
	localctx = NewTypesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, NJSParserRULE_types)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(8)
		_la = p.GetTokenStream().LA(1)

		if !(_la == NJSParserINT || _la == NJSParserI8) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ITypeAnnotationContext is an interface to support dynamic dispatch.
type ITypeAnnotationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeAnnotationContext differentiates from other interfaces.
	IsTypeAnnotationContext()
}

type TypeAnnotationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeAnnotationContext() *TypeAnnotationContext {
	var p = new(TypeAnnotationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NJSParserRULE_typeAnnotation
	return p
}

func (*TypeAnnotationContext) IsTypeAnnotationContext() {}

func NewTypeAnnotationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeAnnotationContext {
	var p = new(TypeAnnotationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NJSParserRULE_typeAnnotation

	return p
}

func (s *TypeAnnotationContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeAnnotationContext) Colon() antlr.TerminalNode {
	return s.GetToken(NJSParserColon, 0)
}

func (s *TypeAnnotationContext) Types() ITypesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypesContext)
}

func (s *TypeAnnotationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeAnnotationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeAnnotationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NJSParserListener); ok {
		listenerT.EnterTypeAnnotation(s)
	}
}

func (s *TypeAnnotationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NJSParserListener); ok {
		listenerT.ExitTypeAnnotation(s)
	}
}

func (p *NJSParser) TypeAnnotation() (localctx ITypeAnnotationContext) {
	localctx = NewTypeAnnotationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, NJSParserRULE_typeAnnotation)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(10)
		p.Match(NJSParserColon)
	}
	{
		p.SetState(11)
		p.Types()
	}

	return localctx
}

// IVariableDeclarationContext is an interface to support dynamic dispatch.
type IVariableDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableDeclarationContext differentiates from other interfaces.
	IsVariableDeclarationContext()
}

type VariableDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableDeclarationContext() *VariableDeclarationContext {
	var p = new(VariableDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NJSParserRULE_variableDeclaration
	return p
}

func (*VariableDeclarationContext) IsVariableDeclarationContext() {}

func NewVariableDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDeclarationContext {
	var p = new(VariableDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NJSParserRULE_variableDeclaration

	return p
}

func (s *VariableDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableDeclarationContext) Let() antlr.TerminalNode {
	return s.GetToken(NJSParserLet, 0)
}

func (s *VariableDeclarationContext) Identifier() antlr.TerminalNode {
	return s.GetToken(NJSParserIdentifier, 0)
}

func (s *VariableDeclarationContext) EqualSign() antlr.TerminalNode {
	return s.GetToken(NJSParserEqualSign, 0)
}

func (s *VariableDeclarationContext) NumberLiteral() antlr.TerminalNode {
	return s.GetToken(NJSParserNumberLiteral, 0)
}

func (s *VariableDeclarationContext) Semicolon() antlr.TerminalNode {
	return s.GetToken(NJSParserSemicolon, 0)
}

func (s *VariableDeclarationContext) TypeAnnotation() ITypeAnnotationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeAnnotationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeAnnotationContext)
}

func (s *VariableDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NJSParserListener); ok {
		listenerT.EnterVariableDeclaration(s)
	}
}

func (s *VariableDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NJSParserListener); ok {
		listenerT.ExitVariableDeclaration(s)
	}
}

func (p *NJSParser) VariableDeclaration() (localctx IVariableDeclarationContext) {
	localctx = NewVariableDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, NJSParserRULE_variableDeclaration)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(13)
		p.Match(NJSParserLet)
	}
	{
		p.SetState(14)
		p.Match(NJSParserIdentifier)
	}
	p.SetState(16)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NJSParserColon {
		{
			p.SetState(15)
			p.TypeAnnotation()
		}

	}
	{
		p.SetState(18)
		p.Match(NJSParserEqualSign)
	}
	{
		p.SetState(19)
		p.Match(NJSParserNumberLiteral)
	}
	{
		p.SetState(20)
		p.Match(NJSParserSemicolon)
	}

	return localctx
}

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NJSParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NJSParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) VariableDeclaration() IVariableDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableDeclarationContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NJSParserListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NJSParserListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *NJSParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, NJSParserRULE_program)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(22)
		p.VariableDeclaration()
	}

	return localctx
}
