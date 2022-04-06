grammar json;

// Grammar Rules
// =============
start
  : Value EOF
  ;

Value
  : Dict
  | List
  | String
  | Number
  | Boolean
  | Null
  ;

Dict
  : CURLY_START Pair (COMMA Pair)* CURLY_END
  | CURLY_START CURLY_END
  ;

List
  : SQUARE_START Value (COMMA Value)* SQUARE_END
  | SQUARE_START SQUARE_END
  ;

String
  : STRING
  ;

Number
  : NUMBER
  ;

Boolean
  : TRUE_SYM
  | FALSE_SYM
  ;

Null
  : NULL_SYM
  ;

Pair
  : STRING COLON Value
  ;

// Lexer Rules
// =============

TRUE_SYM         : 'true';
FALSE_SYM        : 'false';
NULL_SYM         : 'null';

CURLY_START      : '{';
CURLY_END        : '}';
SQUARE_START     : '[';
SQUARE_END       : ']';
COLON            : ':';
COMMA            : ',';

fragment INT     : '0' | [1-9] [0-9]*;
fragment EXP     : [Ee] [+\-]? INT;
fragment ESC     : '\\' (["\\/bfnrt] | UNICODE);
fragment UNICODE : 'u' HEX HEX HEX HEX;
fragment HEX     : [0-9a-fA-F];

NUMBER           : '-'? INT '.' [0-9]+ EXP?
                 | '-'? INT EXP?
                 ;

STRING           : '"' (ESC | ~["\\])* '"'
                 ;

WHITESPACE       : [ \t\n\r]+ -> skip;