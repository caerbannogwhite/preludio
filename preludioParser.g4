parser grammar preludioParser;

options {
	tokenVocab = preludioLexer;
}

nl: NEWLINE | SINGLE_LINE_COMMENT;

program:
	nl* programIntro? nl* (
		(funcDef | pipeline | inlinePipeline | stmt) nl*
	)* EOF;

programIntro: PRQL namedArg* nl;

funcDef: FUNC funcDefName funcDefParams ARROW expr;

funcDefName: IDENT typeDef?;
funcDefParams: funcDefParam*;
funcDefParam: (namedArg | IDENT) typeDef?;

typeDef: LANG typeTerm BAR typeTerm* RANG;
typeTerm: IDENT typeDef?;

stmt: varAssignStmt | varDeclStmt | retStmt | expr;
varAssignStmt: IDENT ASSIGN expr;
varDeclStmt: IDENT DECLARE expr;
retStmt: RET expr;

expr:
	expr AT expr
	| expr EXP expr
	| expr (STAR | DIV | MOD) expr
	| expr (MINUS | PLUS) expr
	| expr MODEL expr
	| expr (EQ | NE | GE | LE | LANG | RANG) expr
	| expr COALESCE expr
	| expr (AND | OR) expr
	| LPAREN expr RPAREN
	| (MINUS | PLUS | NOT) expr
	| list
	| funcCall
	| nestedPipeline
	| literal;

literal:
	NULL_
	| BOOLEAN_LIT
	| INTEGER_LIT
	| RANGE_LIT
	| FLOAT_LIT
	| STRING_LIT
	| STRING_INTERP_LIT
	| STRING_RAW_LIT
	| STRING_PATH_LIT
	| REGEX_LIT
	| DATE_LIT
	| DURATION_LIT
	| IDENT;

list:
	LBRACKET (
		nl* (assign | multiAssign | expr) (
			COMMA nl* (assign | multiAssign | expr)
		)* COMMA? nl?
	)? RBRACKET;

funcCall: IDENT DOLLAR funcCallParam*;
funcCallParam: namedArg | assign | multiAssign | expr;
namedArg: IDENT COLON (assign | expr);
assign: IDENT ASSIGN expr;
multiAssign: list ASSIGN expr;

pipeline: expr (nl funcCall)* (nl | EOF);
inlinePipeline: expr (BAR funcCall)*;
nestedPipeline:
	LPAREN nl* (pipeline | inlinePipeline) nl* RPAREN;