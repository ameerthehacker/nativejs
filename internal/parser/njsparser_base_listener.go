// Code generated from internal/parser/NJSParser.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // NJSParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseNJSParserListener is a complete listener for a parse tree produced by NJSParser.
type BaseNJSParserListener struct{}

var _ NJSParserListener = &BaseNJSParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseNJSParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseNJSParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseNJSParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseNJSParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterTypes is called when production types is entered.
func (s *BaseNJSParserListener) EnterTypes(ctx *TypesContext) {}

// ExitTypes is called when production types is exited.
func (s *BaseNJSParserListener) ExitTypes(ctx *TypesContext) {}

// EnterTypeAnnotation is called when production typeAnnotation is entered.
func (s *BaseNJSParserListener) EnterTypeAnnotation(ctx *TypeAnnotationContext) {}

// ExitTypeAnnotation is called when production typeAnnotation is exited.
func (s *BaseNJSParserListener) ExitTypeAnnotation(ctx *TypeAnnotationContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseNJSParserListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseNJSParserListener) ExitBlock(ctx *BlockContext) {}

// EnterFunctionDeclaration is called when production functionDeclaration is entered.
func (s *BaseNJSParserListener) EnterFunctionDeclaration(ctx *FunctionDeclarationContext) {}

// ExitFunctionDeclaration is called when production functionDeclaration is exited.
func (s *BaseNJSParserListener) ExitFunctionDeclaration(ctx *FunctionDeclarationContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseNJSParserListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseNJSParserListener) ExitExpression(ctx *ExpressionContext) {}

// EnterAssignmentStatement is called when production assignmentStatement is entered.
func (s *BaseNJSParserListener) EnterAssignmentStatement(ctx *AssignmentStatementContext) {}

// ExitAssignmentStatement is called when production assignmentStatement is exited.
func (s *BaseNJSParserListener) ExitAssignmentStatement(ctx *AssignmentStatementContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseNJSParserListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseNJSParserListener) ExitStatement(ctx *StatementContext) {}

// EnterVariableDeclaration is called when production variableDeclaration is entered.
func (s *BaseNJSParserListener) EnterVariableDeclaration(ctx *VariableDeclarationContext) {}

// ExitVariableDeclaration is called when production variableDeclaration is exited.
func (s *BaseNJSParserListener) ExitVariableDeclaration(ctx *VariableDeclarationContext) {}

// EnterFunctionCall is called when production functionCall is entered.
func (s *BaseNJSParserListener) EnterFunctionCall(ctx *FunctionCallContext) {}

// ExitFunctionCall is called when production functionCall is exited.
func (s *BaseNJSParserListener) ExitFunctionCall(ctx *FunctionCallContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseNJSParserListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseNJSParserListener) ExitProgram(ctx *ProgramContext) {}
