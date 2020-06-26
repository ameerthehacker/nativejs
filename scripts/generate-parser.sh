rm -rf internal/parser/grammar/antlr

antlr -Dlanguage=Go internal/parser/NJSLexer.g4 internal/parser/NJSParser.g4
