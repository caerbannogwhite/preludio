lexer grammar preludioLexer;

SINGLE_LINE_COMMENT:
	'#' ~[\r\n\u2028\u2029]* -> channel(HIDDEN);
REGEXP_LITERAL: 'r/' REGEXP_FIRST_CHAR REGEXP_CHAR* '/';

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
NULL_: 'null';
BOOLEAN: 'true' | 'false';

INTEGER: DIGIT+;
FLOAT: DIGIT+ DOT DIGIT* EXP? | DIGIT+ EXP? | DOT DIGIT+ EXP?;

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
COMMENT: '#' ~('\r' | '\n')* NEWLINE;

INTERVAL_KIND:
	'microseconds'
	| 'milliseconds'
	| 'seconds'
	| 'minutes'
	| 'hours'
	| 'days'
	| 'weeks'
	| 'months'
	| 'years';

RANGELIT: (INTEGER | IDENT) RANGE (INTEGER | IDENT) (
		COLON (INTEGER | IDENT)
	)?;

STRING: '"' (ESC | ~[\\"])*? '"' | '\'' (ESC | ~[\\'])*? '\'';

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

// date: '@' date_inner end_expr; time: '@' time_inner end_expr; timestamp: '@' timestamp_inner
// end_expr;

// We use the `inner` types as containing the data that we want to retain in the AST. date_inner =
// ${ ASCII_DIGIT{4} ~ "-" ~ ASCII_DIGIT{2} ~ "-" ~ ASCII_DIGIT{2} } Times are liberally defined
// atm, we could make this more robust. time_inner = ${ ASCII_DIGIT{2} ~ (( ":" | "." ) ~
// ASCII_DIGIT* )* ~ ((( "+" | "-" ) ~ (ASCII_DIGIT | ":" )*) | "Z")? } timestamp_inner = ${
// date_inner ~ "T" ~ time_inner }

// We can use this when want to ensure something is ending, like a date, so `@20-01-0` isn't treated
// like a time `@20-01` `-` (minus) `0`. (Not sure whether `..` should be here or in the items that
// allow it; feel free to demote it to those items if `end_expr` is used somewhere where it's not
// supported) end_expr: ',' | ')' | ']' | nl | '..';
