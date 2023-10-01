parser grammar preludioParser;

options {
	tokenVocab = preludioLexer;
}

nl: NEWLINE;

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
varAssignStmt: IDENT ASSIGN exprCall;
varDeclStmt: IDENT DECLARE exprCall;
retStmt: RET exprCall;

pipeline: exprCall (nl funcCall)* (nl | EOF);
inlinePipeline: exprCall (BAR funcCall)*;
nestedPipeline:
	LPAREN nl* (pipeline | inlinePipeline) nl* RPAREN;

identBacktick: BACKTICK ~(NEWLINE | BACKTICK)* BACKTICK;

funcCall: IDENT funcCallParam*;

funcCallParam: namedArg | assign | multiAssign | expr;
namedArg: IDENT COLON (assign | expr);
assign: IDENT ASSIGN exprCall;
multiAssign: list ASSIGN exprCall;

exprCall: funcCall | expr;

expr:
	expr DOLLAR expr
	| expr EXP expr
	| expr (STAR | DIV | MOD) expr
	| expr (MINUS | PLUS) expr
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

exprUnary: (MINUS | PLUS | NOT) term;

literal:
	IDENT
	| NULL_
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
	| DURATION_LIT;

list:
	LBRACKET (
		nl* (assign | multiAssign | exprCall) (
			COMMA nl* (assign | multiAssign | exprCall)
		)* COMMA? nl?
	)? RBRACKET;
