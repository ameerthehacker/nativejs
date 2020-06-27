parser grammar NJSParser;

options {
  tokenVocab=NJSLexer;
}

types: INT | I8 | I16;
typeAnnotation: ':' types;

block: '{'statement*'}';

functionDeclaration: 'function' Identifier '('')' block;

expression:  Identifier | NumberLiteral | functionCall;

assignmentStatement: Identifier '=' expression ';';

statement: assignmentStatement | expression;

variableDeclaration: 'let' Identifier typeAnnotation? '=' NumberLiteral ';';

functionCall: Identifier '(' expression* ')' ';';

program: (variableDeclaration | statement | functionCall | functionDeclaration)*;
