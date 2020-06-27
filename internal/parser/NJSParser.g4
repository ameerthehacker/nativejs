parser grammar NJSParser;

options {
  tokenVocab=NJSLexer;
}

types: INT | I8 | I16;
typeAnnotation: ':' types;

expression: Identifier | NumericLiteral;

variableDeclaration: 'let' Identifier typeAnnotation? '=' NumberLiteral ';';

functionCall: Identifier '(' expression ')' ';';

program: (variableDeclaration | functionCall)*;
