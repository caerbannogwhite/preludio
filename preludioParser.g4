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
varDeclStmt: IDENT DECLARE expr;
retStmt: RET expr;

pipeline: exprCall (nl funcCall)* (nl | EOF);
inlinePipeline: exprCall (BAR funcCall)*;

identBacktick: BACKTICK ~(NEWLINE | BACKTICK)* BACKTICK;

funcCall: IDENT WHITESPACE funcCallParam*;

funcCallParam: namedArg | assign | multiAssign | expr;
namedArg: IDENT COLON (assign | expr);
assign: IDENT ASSIGN exprCall;
multiAssign: list ASSIGN exprCall;
// assignCall: IDENT ASSIGN exprCall;

exprCall: expr | funcCall;

expr:
	expr LBRACKET expr RBRACKET
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
	IDENT				# ident
	| NULL_				# null
	| BOOLEAN			# boolean
	| INTEGER			# integer
	| FLOAT				# float
	| STRING			# string
	| STRING_RAW		# stringRaw
	| REGEXP_LITERAL	# regexp
	| RANGE_LITERAL		# range
	| DATE_LITERAL		# date
	| DURATION_LITERAL	# duration;

list:
	LBRACKET (
		nl* (assign | multiAssign | exprCall) (
			COMMA nl* (assign | multiAssign | exprCall)
		)* COMMA? nl?
	)? RBRACKET;

nestedPipeline:
	LPAREN nl* (pipeline | inlinePipeline) nl* RPAREN;

