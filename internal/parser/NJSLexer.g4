lexer grammar NJSLexer;

MultiLineComment :       '/*' .*? '*/' -> channel(HIDDEN);
SingleLineComment:       '//' ~[\r\n\u2028\u2029]* -> channel(HIDDEN);
Comma:                   ',';
Colon:                   ':';
EqualSign:               '=';
Semicolon:               ';';
OPEN_PARENTHESIS:        '(';
CLOSED_PARENTHESIS:      ')';
OPEN_CURLY:              '{';
CLOSED_CURLY:            '}';
// keywords
Let:                     'let';
Const:                   'const';
Function:                'function';
INT:                     'number';
I8:                      'i8';
I16:                     'i16';
// Literals
NumberLiteral:           [0-9]+;
Identifier:              [a-zA-Z]+;
// skip whitespces and new lines
WhiteSpaces:             [\t\u000B\u000C\u0020\u00A0]+ -> channel(HIDDEN);
LineTerminator:                 [\r\n\u2028\u2029] -> channel(HIDDEN);
