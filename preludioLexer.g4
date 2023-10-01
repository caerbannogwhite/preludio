lexer grammar preludioLexer;

FUNC: 'func';
PRQL: 'prql';
RET: 'ret';
ARROW: '->';
ASSIGN: '=';
DECLARE: ':=';

PLUS: '+';
MINUS: '-';
STAR: '*';
EXP: '^';
DIV: '/';
MOD: '%';
MODEL: '~';

EQ: '==';
NE: '!=';
LE: '<=';
GE: '>=';

AT: '@';
BAR: '|';
COLON: ':';
COMMA: ',';
DOT: '.';
DOLLAR: '$';
RANGE: '..';
LANG: '<';
RANG: '>';
LBRACKET: '[';
RBRACKET: ']';
LPAREN: '(';
RPAREN: ')';
LBRACE: '{';
RBRACE: '}';
UNDERSCORE: '_';

BACKTICK: '`';
DOUBLE_QUOTE: '"';
SINGLE_QUOTE: '\'';
TRIPLE_DOUBLE_QUOTE: '"""';
TRIPLE_SINGLE_QUOTE: '\'\'\'';

AND: 'and';
OR: 'or';
NOT: 'not';
COALESCE: '??';
NULL_: 'na';

IDENT: IDENT_START (DOT IDENT_NEXT)*;
IDENT_START: (LETTER | UNDERSCORE) (LETTER | DIGIT | UNDERSCORE)*;
IDENT_NEXT: IDENT_START | STAR;

WHITESPACE: (' ' | '\t') -> skip;
NEWLINE: '\r'? '\n';

SINGLE_LINE_COMMENT:
	'#' ~[\r\n\u2028\u2029]* -> channel(HIDDEN);

// Literals
BOOL_LIT: 'true' | 'false';

INT_LIT: DIGIT+;

RNG_LIT: (INT_LIT | IDENT) RANGE (INT_LIT | IDENT) (
		COLON (INT_LIT | IDENT)
	)?;

FLT_LIT:
	DIGIT+ DOT DIGIT* EXPONENT?
	| DIGIT+ EXPONENT?
	| DOT DIGIT+ EXPONENT?;

STR_CHAR: (ESC | ~[\\'\r\n\u2028\u2029]);
STR_RAW_CHAR: ~[\\'\r\n\u2028\u2029];

STR_LIT: SINGLE_QUOTE STR_CHAR*? SINGLE_QUOTE;
STR_INTERP: STR_INTERP_START STR_CHAR*? SINGLE_QUOTE;
STR_RAW: STR_RAW_START STR_RAW_CHAR*? SINGLE_QUOTE;
STR_PATH: STR_PATH_START STR_CHAR*? SINGLE_QUOTE;

RXP_LIT:
	RXP_START RXP_FIRST_CHAR (RXP_CHAR | ~[\\'])*? SINGLE_QUOTE;

DAT_LIT: DATE_START STR_CHAR*? SINGLE_QUOTE;

DUR_LIT: INT_LIT COLON DUR_KIND;

fragment DIGIT: [0-9];
fragment LETTER: [a-zA-Z];
fragment EXPONENT: ('E' | 'e') ('+' | '-')? INT_LIT;

fragment ESC:
	'\\' [abtnfrv"'\\]
	| UNICODE_ESCAPE
	| HEX_ESCAPE
	| OCTAL_ESCAPE;

fragment UNICODE_ESCAPE:
	'\\' 'u' HEXDIGIT HEXDIGIT HEXDIGIT HEXDIGIT
	| '\\' 'u' '{' HEXDIGIT HEXDIGIT HEXDIGIT HEXDIGIT '}';

fragment OCTAL_ESCAPE:
	'\\' [0-3] [0-7] [0-7]
	| '\\' [0-7] [0-7]
	| '\\' [0-7];

fragment HEX_ESCAPE: '\\' HEXDIGIT HEXDIGIT?;

fragment HEXDIGIT: ('0' ..'9' | 'a' ..'f' | 'A' ..'F');

fragment STR_INTERP_START: 'f\'';
fragment STR_RAW_START: 'r\'';
fragment STR_PATH_START: 'p\'';
fragment RXP_START: 'x\'';
fragment DATE_START: 'd\'';

fragment RXP_FIRST_CHAR:
	~[*\r\n\u2028\u2029\\/[]
	| RXP_BACK_SEQ
	| '[' RXP_CLASS_CHAR* ']';

fragment RXP_CHAR:
	~[\r\n\u2028\u2029\\/[]
	| RXP_BACK_SEQ
	| '[' RXP_CLASS_CHAR* ']';

fragment RXP_CLASS_CHAR: ~[\r\n\u2028\u2029\]\\] | RXP_BACK_SEQ;

fragment RXP_BACK_SEQ: '\\' ~[\r\n\u2028\u2029];

fragment DUR_KIND:
	'microseconds'
	| 'milliseconds'
	| 'seconds'
	| 'minutes'
	| 'hours'
	| 'days'
	| 'weeks'
	| 'months'
	| 'years'
	| 'us'
	| 'ms'
	| 's'
	| 'm'
	| 'h'
	| 'd'
	| 'w'
	| 'M'
	| 'y';
