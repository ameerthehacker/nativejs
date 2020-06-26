lexer grammar NJSLexer;

MultiLineComment :       '/*' .*? '*/' -> channel(HIDDEN);
SingleLineComment:       '//' ~[\r\n\u2028\u2029]* -> channel(HIDDEN);
Comma:                   ',';
Colon:                   ':';
EqualSign:               '=';
Let:                     'let';
Semicolon:               ';';
Const:                   'const';
Identifier:              [a-zA-Z]+;
NumberLiteral:           [0-9]+;
// types
INT:                     'number';
I8:                      'i8';
I16:                     'i16';
// skip whitespces and new lines
WhiteSpaces:             [\t\u000B\u000C\u0020\u00A0]+ -> channel(HIDDEN);
LineTerminator:                 [\r\n\u2028\u2029] -> channel(HIDDEN);
