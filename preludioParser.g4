// TODO: - Some rules are silent because we don't strictly need them, but that may be too rushed
// — maybe we do actually want to know about e.g. comments in the parse tree (and potentially put
// them into SQL comments) - Need to resolve how to handle "inline pipelines"; there is a rule here
// but it's not used or tested. It's partly a language question — do those need to start with
// `from`? How do these work in the midst of an `aggregate` transform?

parser grammar preludioParser;

options {
	tokenVocab = preludioLexer;
}

nl: NEWLINE;

program:
	nl* programIntro? nl* ((funcDef | stmt | nestedPipeline) nl*)* EOF;

programIntro: PRQL namedArg* nl;

funcDef: FUNC funcDefName funcDefParams ARROW expr;

funcDefName: IDENT typeDef?;
funcDefParams: funcDefParam*;
funcDefParam: (namedArg | IDENT) typeDef?;
typeDef: LANG typeTerm BAR typeTerm* RANG;
typeTerm: IDENT typeDef?;

stmt: varAssignStmt | varDeclStmt | retStmt | expr;
varAssignStmt: IDENT ASSIGN expr;
varDeclStmt: LET IDENT ASSIGN expr;
retStmt: RET expr;

pipeline: exprCall (nl funcCall)* (nl | EOF);
inlinePipeline: exprCall (BAR funcCall)*;

// We include backticks because some DBs use them (e.g. BigQuery) and we don't, so we pass anything
// within them directly through, including otherwise invalid idents, like those with hyphens.
// Possibly we should consider applying this to expressions rather than just idents — we can adjust
// as we see more cases. ident: operator (keyword WHITESPACE) ident_start (DOT ident_next)*; //
// Either a normal ident (starting with a letter, `$` or `_`), or any string surrounded by //
// backticks. ident_start: ( (ASCII_ALPHA | DOLLAR | UNDERSCORE) ~ ( ASCII_ALPHANUMERIC | DOLLAR )*
// ) | identBackticks+; // We allow `e.*`, but not just `*`, since it might conflict with multiply
// in some cases. ident_next: ident_start | '*'; Anything surrounded by backticks, we pass through.
identBacktick: BACKTICK ~(NEWLINE | BACKTICK)* BACKTICK;

// For sorting signedIdent: (PLUS | MINUS) IDENT;

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
funcCall: IDENT funcCallParam*;

funcCallParam: namedArg | assign | multiAssign | expr;
namedArg: IDENT COLON (assign | expr);
assign: IDENT ASSIGN exprCall;
multiAssign: list ASSIGN exprCall;
// assignCall: IDENT ASSIGN exprCall;

exprCall: expr | funcCall;

expr:
	expr LBRACKET expr RBRACKET
	| expr (STAR | DIV | MOD) expr
	| expr (MINUS | PLUS) expr
	| expr POW expr
	| expr MODEL expr
	| expr (EQ | NE | GE | LE | LANG | RANG) expr
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

exprUnary: (MINUS | PLUS | NOT) (
		nestedPipeline
		| literal
		| IDENT
	);

literal:
	IDENT
	| NULL_
	| BOOLEAN
	| STRING
	| STRING_RAW
	| REGEXP_LITERAL
	| INTEGER
	| FLOAT
	| RANGE_LITERAL
	| DATE_LITERAL
	| DURATION_LITERAL;

list:
	LBRACKET (
		nl* (assign | multiAssign | exprCall) (
			COMMA nl* (assign | multiAssign | exprCall)
		)* COMMA? nl?
	)? RBRACKET;

nestedPipeline:
	LPAREN nl* (pipeline | inlinePipeline) nl* RPAREN;

