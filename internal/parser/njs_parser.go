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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 21, 44, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 3,
	2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 5, 5, 25, 10,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3,
	7, 7, 7, 39, 10, 7, 12, 7, 14, 7, 42, 11, 7, 3, 7, 2, 2, 8, 2, 4, 6, 8,
	10, 12, 2, 4, 3, 2, 16, 18, 4, 2, 11, 11, 21, 21, 2, 40, 2, 14, 3, 2, 2,
	2, 4, 16, 3, 2, 2, 2, 6, 19, 3, 2, 2, 2, 8, 21, 3, 2, 2, 2, 10, 30, 3,
	2, 2, 2, 12, 40, 3, 2, 2, 2, 14, 15, 9, 2, 2, 2, 15, 3, 3, 2, 2, 2, 16,
	17, 7, 6, 2, 2, 17, 18, 5, 2, 2, 2, 18, 5, 3, 2, 2, 2, 19, 20, 9, 3, 2,
	2, 20, 7, 3, 2, 2, 2, 21, 22, 7, 8, 2, 2, 22, 24, 7, 11, 2, 2, 23, 25,
	5, 4, 3, 2, 24, 23, 3, 2, 2, 2, 24, 25, 3, 2, 2, 2, 25, 26, 3, 2, 2, 2,
	26, 27, 7, 7, 2, 2, 27, 28, 7, 12, 2, 2, 28, 29, 7, 9, 2, 2, 29, 9, 3,
	2, 2, 2, 30, 31, 7, 11, 2, 2, 31, 32, 7, 14, 2, 2, 32, 33, 5, 6, 4, 2,
	33, 34, 7, 15, 2, 2, 34, 35, 7, 9, 2, 2, 35, 11, 3, 2, 2, 2, 36, 39, 5,
	8, 5, 2, 37, 39, 5, 10, 6, 2, 38, 36, 3, 2, 2, 2, 38, 37, 3, 2, 2, 2, 39,
	42, 3, 2, 2, 2, 40, 38, 3, 2, 2, 2, 40, 41, 3, 2, 2, 2, 41, 13, 3, 2, 2,
	2, 42, 40, 3, 2, 2, 2, 5, 24, 38, 40,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "", "", "','", "':'", "'='", "'let'", "';'", "'const'", "", "", "'function'",
	"'('", "')'", "'number'", "'i8'", "'i16'",
}
var symbolicNames = []string{
	"", "MultiLineComment", "SingleLineComment", "Comma", "Colon", "EqualSign",
	"Let", "Semicolon", "Const", "Identifier", "NumberLiteral", "Function",
	"OPEN_PARENTHESIS", "CLOSED_PARENTHESIS", "INT", "I8", "I16", "WhiteSpaces",
	"LineTerminator", "NumericLiteral",
}

var ruleNames = []string{
	"types", "typeAnnotation", "expression", "variableDeclaration", "functionCall",
	"program",
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
	NJSParserEOF                = antlr.TokenEOF
	NJSParserMultiLineComment   = 1
	NJSParserSingleLineComment  = 2
	NJSParserComma              = 3
	NJSParserColon              = 4
	NJSParserEqualSign          = 5
	NJSParserLet                = 6
	NJSParserSemicolon          = 7
	NJSParserConst              = 8
	NJSParserIdentifier         = 9
	NJSParserNumberLiteral      = 10
	NJSParserFunction           = 11
	NJSParserOPEN_PARENTHESIS   = 12
	NJSParserCLOSED_PARENTHESIS = 13
	NJSParserINT                = 14
	NJSParserI8                 = 15
	NJSParserI16                = 16
	NJSParserWhiteSpaces        = 17
	NJSParserLineTerminator     = 18
	NJSParserNumericLiteral     = 19
)

// NJSParser rules.
const (
	NJSParserRULE_types               = 0
	NJSParserRULE_typeAnnotation      = 1
	NJSParserRULE_expression          = 2
	NJSParserRULE_variableDeclaration = 3
	NJSParserRULE_functionCall        = 4
	NJSParserRULE_program             = 5
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

func (s *TypesContext) I16() antlr.TerminalNode {
	return s.GetToken(NJSParserI16, 0)
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
		p.SetState(12)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NJSParserINT)|(1<<NJSParserI8)|(1<<NJSParserI16))) != 0) {
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
		p.SetState(14)
		p.Match(NJSParserColon)
	}
	{
		p.SetState(15)
		p.Types()
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NJSParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NJSParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Identifier() antlr.TerminalNode {
	return s.GetToken(NJSParserIdentifier, 0)
}

func (s *ExpressionContext) NumericLiteral() antlr.TerminalNode {
	return s.GetToken(NJSParserNumericLiteral, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NJSParserListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NJSParserListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *NJSParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, NJSParserRULE_expression)
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
		p.SetState(17)
		_la = p.GetTokenStream().LA(1)

		if !(_la == NJSParserIdentifier || _la == NJSParserNumericLiteral) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
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
	p.EnterRule(localctx, 6, NJSParserRULE_variableDeclaration)
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
		p.SetState(19)
		p.Match(NJSParserLet)
	}
	{
		p.SetState(20)
		p.Match(NJSParserIdentifier)
	}
	p.SetState(22)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NJSParserColon {
		{
			p.SetState(21)
			p.TypeAnnotation()
		}

	}
	{
		p.SetState(24)
		p.Match(NJSParserEqualSign)
	}
	{
		p.SetState(25)
		p.Match(NJSParserNumberLiteral)
	}
	{
		p.SetState(26)
		p.Match(NJSParserSemicolon)
	}

	return localctx
}

// IFunctionCallContext is an interface to support dynamic dispatch.
type IFunctionCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionCallContext differentiates from other interfaces.
	IsFunctionCallContext()
}

type FunctionCallContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionCallContext() *FunctionCallContext {
	var p = new(FunctionCallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NJSParserRULE_functionCall
	return p
}

func (*FunctionCallContext) IsFunctionCallContext() {}

func NewFunctionCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionCallContext {
	var p = new(FunctionCallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NJSParserRULE_functionCall

	return p
}

func (s *FunctionCallContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionCallContext) Identifier() antlr.TerminalNode {
	return s.GetToken(NJSParserIdentifier, 0)
}

func (s *FunctionCallContext) OPEN_PARENTHESIS() antlr.TerminalNode {
	return s.GetToken(NJSParserOPEN_PARENTHESIS, 0)
}

func (s *FunctionCallContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *FunctionCallContext) CLOSED_PARENTHESIS() antlr.TerminalNode {
	return s.GetToken(NJSParserCLOSED_PARENTHESIS, 0)
}

func (s *FunctionCallContext) Semicolon() antlr.TerminalNode {
	return s.GetToken(NJSParserSemicolon, 0)
}

func (s *FunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NJSParserListener); ok {
		listenerT.EnterFunctionCall(s)
	}
}

func (s *FunctionCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NJSParserListener); ok {
		listenerT.ExitFunctionCall(s)
	}
}

func (p *NJSParser) FunctionCall() (localctx IFunctionCallContext) {
	localctx = NewFunctionCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, NJSParserRULE_functionCall)

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
		p.SetState(28)
		p.Match(NJSParserIdentifier)
	}
	{
		p.SetState(29)
		p.Match(NJSParserOPEN_PARENTHESIS)
	}
	{
		p.SetState(30)
		p.Expression()
	}
	{
		p.SetState(31)
		p.Match(NJSParserCLOSED_PARENTHESIS)
	}
	{
		p.SetState(32)
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

func (s *ProgramContext) AllVariableDeclaration() []IVariableDeclarationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVariableDeclarationContext)(nil)).Elem())
	var tst = make([]IVariableDeclarationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVariableDeclarationContext)
		}
	}

	return tst
}

func (s *ProgramContext) VariableDeclaration(i int) IVariableDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableDeclarationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVariableDeclarationContext)
}

func (s *ProgramContext) AllFunctionCall() []IFunctionCallContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFunctionCallContext)(nil)).Elem())
	var tst = make([]IFunctionCallContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFunctionCallContext)
		}
	}

	return tst
}

func (s *ProgramContext) FunctionCall(i int) IFunctionCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionCallContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
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
	p.EnterRule(localctx, 10, NJSParserRULE_program)
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
	p.SetState(38)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == NJSParserLet || _la == NJSParserIdentifier {
		p.SetState(36)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case NJSParserLet:
			{
				p.SetState(34)
				p.VariableDeclaration()
			}

		case NJSParserIdentifier:
			{
				p.SetState(35)
				p.FunctionCall()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(40)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}
