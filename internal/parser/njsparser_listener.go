// Code generated from internal/parser/NJSParser.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // NJSParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// NJSParserListener is a complete listener for a parse tree produced by NJSParser.
type NJSParserListener interface {
	antlr.ParseTreeListener

	// EnterTypes is called when entering the types production.
	EnterTypes(c *TypesContext)

	// EnterTypeAnnotation is called when entering the typeAnnotation production.
	EnterTypeAnnotation(c *TypeAnnotationContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterVariableDeclaration is called when entering the variableDeclaration production.
	EnterVariableDeclaration(c *VariableDeclarationContext)

	// EnterFunctionCall is called when entering the functionCall production.
	EnterFunctionCall(c *FunctionCallContext)

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// ExitTypes is called when exiting the types production.
	ExitTypes(c *TypesContext)

	// ExitTypeAnnotation is called when exiting the typeAnnotation production.
	ExitTypeAnnotation(c *TypeAnnotationContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitVariableDeclaration is called when exiting the variableDeclaration production.
	ExitVariableDeclaration(c *VariableDeclarationContext)

	// ExitFunctionCall is called when exiting the functionCall production.
	ExitFunctionCall(c *FunctionCallContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)
}
