package codegen

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/llir/llvm/ir"
	constant2 "github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"nativejs/internal"
	"nativejs/internal/parser"
)

var variableType types.Type

type variableDeclarationListener struct {
	parser.BaseNJSParserListener
}

func (variableDeclarationListener) EnterTypes(cxt *parser.TypesContext) {
	switch {
	case cxt.I8() != nil:
		variableType = types.I8
	case cxt.I16() != nil:
		variableType = types.I16
	case cxt.INT() != nil:
		variableType = types.I32
	}
}

func GenVariableDeclaration(module *ir.Module, cxt *parser.VariableDeclarationContext) {
	variableName := cxt.Identifier().GetText()
	initializer := cxt.NumberLiteral().GetText()

	// walk to find the variable type
	antlr.ParseTreeWalkerDefault.Walk(&variableDeclarationListener{}, cxt)

	switch variableType {
	case types.I8, types.I16, types.I32:
		constant, err := constant2.NewIntFromString(new(types.IntType), initializer)

		internal.Must(err)

		module.NewGlobalDef(variableName, constant)
	}
}
