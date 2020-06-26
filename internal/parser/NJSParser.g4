parser grammar NJSParser;

options {
  tokenVocab=NJSLexer;
}

types: INT | I8 | I16;
typeAnnotation: ':' types;

variableDeclaration: 'let' Identifier typeAnnotation? '=' NumberLiteral ';';

program: variableDeclaration;