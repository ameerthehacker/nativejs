package codegen

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
	"nativejs/internal/parser"
)

var module *ir.Module
var currentType types.Type

type njsListener struct {
	parser.BaseNJSParserListener
}

func (njsListener) EnterTypes(cxt *parser.TypesContext) {
	switch {
	case cxt.I8() != nil:
		currentType = types.I8
	case cxt.I16() != nil:
		currentType = types.I16
	case cxt.INT() != nil:
		currentType = types.I32
	}
}

func (njsListener) EnterVariableDeclaration(cxt *parser.VariableDeclarationContext) {
	variableName := cxt.Identifier().GetText()

	module.NewGlobal(variableName, currentType)
}

func GenerateIR(ctx parser.IProgramContext) {
	module = ir.NewModule()

	antlr.ParseTreeWalkerDefault.Walk(&njsListener{}, ctx)
}

func GetIR() string {
	return module.String()
}
