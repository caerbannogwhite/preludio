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
BOOLEAN: 'true' | 'false';

IDENT: IDENT_START (DOT IDENT_NEXT)*;
IDENT_START: (LETTER | DOLLAR | UNDERSCORE) (
		LETTER
		| DIGIT
		| UNDERSCORE
	)*;
IDENT_NEXT: IDENT_START | STAR;

WHITESPACE: (' ' | '\t') -> skip;
NEWLINE: '\r'? '\n';

SINGLE_LINE_COMMENT:
	'#' ~[\r\n\u2028\u2029]* -> channel(HIDDEN);

INTEGER: DIGIT+;
FLOAT:
	DIGIT+ DOT DIGIT* EXPONENT?
	| DIGIT+ EXPONENT?
	| DOT DIGIT+ EXPONENT?;

STRING_CHAR: (ESC | ~[\\'\r\n\u2028\u2029]);

STRING: SINGLE_QUOTE STRING_CHAR*? SINGLE_QUOTE;
STRING_INTERP: STRING_INTERP_START STRING_CHAR*? SINGLE_QUOTE;
STRING_RAW: STRING_RAW_START STRING_CHAR*? SINGLE_QUOTE;
STRING_PATH: STRING_PATH_START STRING_CHAR*? SINGLE_QUOTE;

REGEXP_LITERAL:
	REGEXP_START REGEXP_FIRST_CHAR (REGEXP_CHAR | ~[\\'])*? SINGLE_QUOTE;

RANGE_LITERAL: (INTEGER | IDENT) RANGE (INTEGER | IDENT) (
		COLON (INTEGER | IDENT)
	)?;

DATE_LITERAL: DATE_START STRING_CHAR*? SINGLE_QUOTE;

DURATION_KIND:
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

DURATION_LITERAL: INTEGER COLON DURATION_KIND;

fragment DIGIT: [0-9];
fragment LETTER: [a-zA-Z];
fragment EXPONENT: ('E' | 'e') ('+' | '-')? INTEGER;

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

fragment STRING_INTERP_START: 'f\'';
fragment STRING_RAW_START: 'r\'';
fragment STRING_PATH_START: 'p\'';
fragment REGEXP_START: 'x\'';
fragment DATE_START: 'd\'';

fragment REGEXP_FIRST_CHAR:
	~[*\r\n\u2028\u2029\\/[]
	| REGEXP_BACK_SEQ
	| '[' REGEXP_CLASS_CHAR* ']';

fragment REGEXP_CHAR:
	~[\r\n\u2028\u2029\\/[]
	| REGEXP_BACK_SEQ
	| '[' REGEXP_CLASS_CHAR* ']';

fragment REGEXP_CLASS_CHAR:
	~[\r\n\u2028\u2029\]\\]
	| REGEXP_BACK_SEQ;

fragment REGEXP_BACK_SEQ: '\\' ~[\r\n\u2028\u2029];

