// TODO: - Some rules are silent because we don't strictly need them, but that may be too rushed
// — maybe we do actually want to know about e.g. comments in the parse tree (and potentially put
// them into SQL comments) - Need to resolve how to handle "inline pipelines"; there is a rule here
// but it's not used or tested. It's partly a language question — do those need to start with
// `from`? How do these work in the midst of an `aggregate` transform?

grammar prql;

FUNC: 'func';
PRQL: 'prql';
TABLE: 'table';

PLUS: '+';
MINUS: '-';
EQUAL: '=';

BAR: '|';
COMMA: ',';
DOT: '.';
DOLLAR: '$';
UNDERSCORE: '_';
LANG: '<';
RANG: '>';
LBRACKET: '[';
RBRACKET: ']';
LPAREN: '(';
RPAREN: ')';

NULL_: 'null';
BOOLEAN: 'true' | 'false';

fragment INT: DIGIT+;
fragment DIGIT: [0-9];
fragment LETTER: [a-zA-Z];
fragment EXP: ('E' | 'e') ('+' | '-')? INT;

NUMBER: DIGIT+ DOT DIGIT* EXP? | DIGIT+ EXP? | DOT DIGIT+ EXP?;

// Either a normal ident (starting with a letter, `$` or `_`), or any string surrounded by
// backticks. We allow `e.*`, but not just `*`, since it might conflict with multiply in some cases.
IDENT: IDENT_START (DOT IDENT_NEXT)*;
IDENT_START: (LETTER | DOLLAR | UNDERSCORE) (
		LETTER
		| DIGIT
		| UNDERSCORE
	)*;
IDENT_NEXT: IDENT_START | '*';

WHITESPACE: (' ' | '\t') -> skip;
NEWLINE: '\r'? '\n';

// Need to exclude # in strings (and maybe confirm whether this the syntax we want)
COMMENT: '#' ~('\r' | '\n')* NEWLINE;

nl: NEWLINE | COMMENT;

query:
	nl* query_def? nl* ((func_def | table | pipeline) nl*)* EOF;

query_def: PRQL named_arg* nl;

func_def: FUNC func_def_name func_def_params '->' expr;

func_def_name: IDENT type_def?;
func_def_params: func_def_param*;
func_def_param: (named_arg | IDENT) type_def?;
type_def: LANG type_term BAR type_term* RANG;
type_term: IDENT type_def?;

table: TABLE IDENT EQUAL nested_pipeline;

pipe: nl | BAR;
pipeline: expr_call (pipe expr_call)*;

// We include backticks because some DBs use them (e.g. BigQuery) and we don't, so we pass anything
// within them directly through, including otherwise invalid idents, like those with hyphens.
// Possibly we should consider applying this to expressions rather than just idents — we can adjust
// as we see more cases. ident: operator (keyword WHITESPACE) ident_start (DOT ident_next)*; //
// Either a normal ident (starting with a letter, `$` or `_`), or any string surrounded by //
// backticks. ident_start: ( (ASCII_ALPHA | DOLLAR | UNDERSCORE) ~ ( ASCII_ALPHANUMERIC | DOLLAR )*
// ) | ident_backticks+; // We allow `e.*`, but not just `*`, since it might conflict with multiply
// in some cases. ident_next: ident_start | '*'; Anything surrounded by backticks, we pass through.
ident_backticks: '`' (nl '`' .)* '`';
// For sorting
signed_ident: (PLUS | MINUS) IDENT;
keyword: PRQL | TABLE | FUNC;

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
func_call: IDENT (named_arg | assign | expr)+;

named_arg: IDENT ':' (assign | expr);
assign: IDENT '=' expr;
assign_call: IDENT '=' expr_call;

expr_call: func_call | expr;

expr:
	expr operator_mul expr
	| expr operator_add expr
	| expr operator_compare expr
	| expr operator_coalesce expr
	| expr operator_logical expr
	| LPAREN expr RPAREN
	| term;

term:
	// s_string | f_string |
	range
	| literal
	| IDENT
	| expr_unary
	| list // | jinja
	| nested_pipeline;

// expr_unary is for sorting.
expr_unary: operator_unary (nested_pipeline | literal | IDENT);
literal:
	interval
	| NUMBER
	| BOOLEAN
	| NULL_
	| STRING; // | timestamp | date | time

list:
	LBRACKET (
		nl* (assign_call | expr_call) (
			COMMA nl* (assign_call | expr_call)
		)* COMMA? nl?
	)? RBRACKET;

nested_pipeline: LPAREN nl* pipeline nl* RPAREN;

// We haven't implemented escapes — I think we can mostly pass those through to SQL, but there may
// be things we're missing. https://pest.rs/book/examples/rust/literals.html

// We need to have a non-silent rule which contains the quotes — `string` in this case — because of
// https://github.com/pest-parser/pest/issues/583. Then when converting to AST, we only keep the
// `string_inner` and discard the `string` given it contains the quotes.
// 
// TODO: I'm still a bit unclear how preceeding and trailing spaces are working -- it seems that
// inner spaces are included without an atomic operator (or with `ANY`), but prceeding & trailing
// spaces require both `ANY` _and_ an atomic operator. We have some rudimentary tests for these.

single_quote: '"' | '\'';
multi_quote: '"""' | '\'\'\'';
// opening_quote: PUSH (multi_quote) | PUSH (single_quote); PEEK refers to the opening quote; either
// `"` or `'`. string_inner: ( PEEK ANY)+; Either > 3 quotes, or just one. Currently both of those
// can be multiline.
STRING: '"' ~('"')* '"';

// We need `literal` separate from `term_simple` for things like range edges, which would infinitely
// recurse otherwise, since it'll keep trying to parse the whole span, not just the part before
// `..`.
range: literal '..' literal;

operator:
	operator_unary
	| operator_mul
	| operator_add
	| operator_compare
	| operator_logical
	| operator_coalesce;

operator_unary: '-' | '+' | '!';
operator_mul: '*' | '/' | '%';
operator_add: '+' | '-';
operator_compare: '==' | '!=' | '>=' | '<=' | '>' | '<';
operator_logical: 'and' | 'or';
operator_coalesce: '??';

// s_string: 's' opening_quote ( interpolate_string_inner | ( '{' expr_call '}') )* POP; f_string:
// 'f' opening_quote ( interpolate_string_inner | ( '{' expr_call '}') )* POP;
// interpolate_string_inner: ( ( PEEK | '{') ~ ANY)+;

interval_kind:
	'microseconds'
	| 'milliseconds'
	| 'seconds'
	| 'minutes'
	| 'hours'
	| 'days'
	| 'weeks'
	| 'months'
	| 'years';
interval: NUMBER interval_kind;

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

// We pass text between `{{` and `}}` through, so dbt can use Jinja. jinja = { ("{{" ~ (!"}}" ~
// ANY)* ~ "}}") }