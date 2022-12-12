// TODO: - Some rules are silent because we don't strictly need them, but that may be too rushed
// — maybe we do actually want to know about e.g. comments in the parse tree (and potentially put
// them into SQL comments) - Need to resolve how to handle "inline pipelines"; there is a rule here
// but it's not used or tested. It's partly a language question — do those need to start with
// `from`? How do these work in the midst of an `aggregate` transform?

grammar prql;

FUNC: 'func';
PRQL: 'prql';
LET: 'let';
ARROW: '->';
ASSIGN: '=';

PLUS: '+';
MINUS: '-';
STAR: '*';
POW: '**';
DIV: '/';
MOD: '%';
MODEL: '~';

EQ: 'eq';
NE: 'ne';
LE: 'le';
LT: 'lt';
GE: 'ge';
GT: 'gt';

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

fragment DIGIT: [0-9];
fragment LETTER: [a-zA-Z];
fragment EXP: ('E' | 'e') ('+' | '-')? INTEGER;

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

nl: NEWLINE | COMMENT;

program:
	nl* programIntro? nl* ((funcDef | stmt | pipeline) nl*)* EOF;

programIntro: PRQL namedArg* nl;

funcDef: FUNC funcDefName funcDefParams ARROW expr;

funcDefName: IDENT typeDef?;
funcDefParams: funcDefParam*;
funcDefParam: (namedArg | IDENT) typeDef?;
typeDef: LANG typeTerm BAR typeTerm* RANG;
typeTerm: IDENT typeDef?;

stmt: assignStmt;
assignStmt: LET IDENT ASSIGN expr;

pipe: nl | BAR;
pipeline: exprCall (pipe exprCall)*;

// We include backticks because some DBs use them (e.g. BigQuery) and we don't, so we pass anything
// within them directly through, including otherwise invalid idents, like those with hyphens.
// Possibly we should consider applying this to expressions rather than just idents — we can adjust
// as we see more cases. ident: operator (keyword WHITESPACE) ident_start (DOT ident_next)*; //
// Either a normal ident (starting with a letter, `$` or `_`), or any string surrounded by //
// backticks. ident_start: ( (ASCII_ALPHA | DOLLAR | UNDERSCORE) ~ ( ASCII_ALPHANUMERIC | DOLLAR )*
// ) | identBackticks+; // We allow `e.*`, but not just `*`, since it might conflict with multiply
// in some cases. ident_next: ident_start | '*'; Anything surrounded by backticks, we pass through.
identBacktick: BACKTICK ~(NEWLINE | BACKTICK)* BACKTICK;

// For sorting
signedIdent: (PLUS | MINUS) IDENT;

// A central issue around the terms vs expr is that we want to be able to parse: [foo bar + 1, 2]
// as: - foo bar + 1 - foo bar - foo - bar - + - 1 - 2 So this requires two non-silent rules: - A
// notion of list item that contains anything, including operators (but not commas); e.g. `foo bar +
// 1`. - A notion of expr that aggregates things between operators, e.g. foo bar. So we call the
// list item `expr`, and the things between separators `terms`.
// 
// We could have them be the same, but then we need logic in the parser to account for where the
// list item is in this parse tree - foo bar - foo - bar - + - 1 - 2

// whitespace is required to prevent matching s"string". Forbid `operator` so `a - b` can't parse as
// `a` & `-b`.
funcCall: IDENT funcCallParam+;

funcCallParam: namedArg | assign | expr;
namedArg: IDENT COLON (assign | expr);
assign: IDENT ASSIGN expr;
assignCall: IDENT ASSIGN exprCall;

exprCall: expr | funcCall;

expr:
	expr (STAR | DIV | MOD) expr
	| expr (MINUS | PLUS) expr
	| expr POW expr
	| expr MODEL expr
	| expr (EQ | NE | GE | LE | LT | GT) expr
	| expr COALESCE expr
	| expr (AND | OR) expr
	| LPAREN expr RPAREN
	| term;

term:
	literal
	| identBacktick
	| exprUnary
	| list
	| nestedPipeline;

// exprUnary is for sorting.
exprUnary: (MINUS | PLUS | NOT) (
		nestedPipeline
		| literal
		| IDENT
	);

literal:
	IDENT
	| NULL_
	| BOOLEAN
	| STRING // | timestamp | date | time | s_string | f_string |
	| INTEGER
	| FLOAT
	| FLOAT INTERVAL_KIND
	| (FLOAT | IDENT) RANGE (FLOAT | IDENT);

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

list:
	LBRACKET (
		nl* (assignCall | exprCall) (
			COMMA nl* (assignCall | exprCall)
		)* COMMA? nl?
	)? RBRACKET;

nestedPipeline: LPAREN nl* pipeline nl* RPAREN;

// We haven't implemented escapes — I think we can mostly pass those through to SQL, but there may
// be things we're missing. https://pest.rs/book/examples/rust/literals.html

// We need to have a non-silent rule which contains the quotes — `string` in this case — because of
// https://github.com/pest-parser/pest/issues/583. Then when converting to AST, we only keep the
// `string_inner` and discard the `string` given it contains the quotes.
// 
// TODO: I'm still a bit unclear how preceeding and trailing spaces are working -- it seems that
// inner spaces are included without an atomic operator (or with `ANY`), but prceeding & trailing
// spaces require both `ANY` _and_ an atomic operator. We have some rudimentary tests for these.

STRING: '"' (ESC | ~[\\"])*? '"' | '\'' (ESC | ~[\\'])*? '\'';

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
