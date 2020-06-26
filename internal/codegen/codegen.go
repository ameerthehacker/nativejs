package codegen

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"nativejs/internal/parser"
)

type njsListener struct {
	parser.BaseNJSParserListener
}

func (njsListener) EnterVariableDeclaration(cxt *parser.VariableDeclarationContext) {
	fmt.Println(cxt.NumberLiteral().GetText())
}

func GenerateIR(ctx parser.IProgramContext) {
	antlr.ParseTreeWalkerDefault.Walk(&njsListener{}, ctx)
}
