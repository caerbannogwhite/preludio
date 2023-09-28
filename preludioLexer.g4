lexer grammar preludioLexer;

FUNC: 'func';
PRQL: 'prql';
LET: 'let';
RET: 'ret';
ARROW: '->';
ASSIGN: '=';

PLUS: '+';
MINUS: '-';
STAR: '*';
POW: '**';
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
NULL_: 'NA';
BOOLEAN: 'true' | 'false';

// Either a normal ident (starting with a letter, `$` or `_`), or any string surrounded by
// backticks. We allow `e.*`, but not just `*`, since it might conflict with multiply in some cases.
IDENT: IDENT_START (DOT IDENT_NEXT)*;
IDENT_START: (LETTER | DOLLAR | UNDERSCORE) (
		LETTER
		| DIGIT
		| UNDERSCORE
	)*;
IDENT_NEXT: IDENT_START | STAR;

WHITESPACE: (' ' | '\t') -> skip;
NEWLINE: '\r'? '\n';

// Need to exclude # in strings (and maybe confirm whether this the syntax we want)

// COMMENT: '#' ~('\r' | '\n')* NEWLINE;
SINGLE_LINE_COMMENT:
	'#' ~[\r\n\u2028\u2029]* -> channel(HIDDEN);

INTEGER: DIGIT+;
FLOAT: DIGIT+ DOT DIGIT* EXP? | DIGIT+ EXP? | DOT DIGIT+ EXP?;

STRING: SINGLE_QUOTE (ESC | ~[\\'])*? SINGLE_QUOTE;

STRING_RAW: 'r' SINGLE_QUOTE (ESC | ~[\\'])*? SINGLE_QUOTE;

REGEXP_LITERAL:
	'x' SINGLE_QUOTE REGEXP_FIRST_CHAR REGEXP_CHAR* SINGLE_QUOTE;

RANGE_LITERAL: (INTEGER | IDENT) RANGE (INTEGER | IDENT) (
		COLON (INTEGER | IDENT)
	)?;

DATE_LITERAL: 'd' SINGLE_QUOTE (ESC | ~[\\'])*? SINGLE_QUOTE;

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
fragment EXP: ('E' | 'e') ('+' | '-')? INTEGER;

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

