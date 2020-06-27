package codegen

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
	"nativejs/internal/parser"
)

var module *ir.Module

type njsListener struct {
	parser.BaseNJSParserListener
}

func (njsListener) EnterTypes(cxt *parser.TypesContext) {
	switch {
	case cxt.I8() != nil:
		variableType = types.I8
	case cxt.I16() != nil:
		variableType = types.I16
	case cxt.INT() != nil:
		variableType = types.I32
	}
}

func (njsListener) EnterVariableDeclaration(cxt *parser.VariableDeclarationContext) {
	GenVariableDeclaration(module, cxt)
}

func (njsListener) EnterFunctionCall(cxt *parser.FunctionCallContext) {
	println("Calling function", cxt.Identifier().GetText())
}

func GenerateIR(cxt parser.IProgramContext) string {
	module = ir.NewModule()

	antlr.ParseTreeWalkerDefault.Walk(&njsListener{}, cxt)

	return module.String()
}
