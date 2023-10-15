parser grammar preludioParser;

options {
	tokenVocab = preludioLexer;
}

nl: NEWLINE | SINGLE_LINE_COMMENT;

program: nl* programIntro? nl* ((funcDef | stmt) nl*)* EOF;

programIntro: PRQL namedArg* nl;

funcDef:
	FUNC funcDefName funcDefParams ARROW (
		exprCall
		| nl* (stmt nl*)* END
	);

funcDefName: IDENT typeDef?;
funcDefParams: funcDefParam*;
funcDefParam: (namedArg | IDENT) typeDef?;

typeDef: LANG typeTerm BAR typeTerm* RANG;
typeTerm: IDENT typeDef?;

stmt:
	varAssignStmt
	| varDeclStmt
	| ifElseStmt
	| forStmt
	| helpStmt
	| retStmt
	| exprCall;
varAssignStmt: IDENT ASSIGN exprCall;
varDeclStmt: IDENT DECLARE exprCall;
ifElseStmt:
	IF exprCall DO (
		stmt
		| nl* (stmt nl*)* (
			ELSE (stmt | nl* (stmt nl*)* END)
			| END
		)
	);
forStmt: FOR IDENT IN exprCall DO (stmt | nl* (stmt nl*)* END);
helpStmt: HELP IDENT;
retStmt: RET exprCall;

exprCall: expr | funcCall;

expr:
	literal
	| expr INDEXING expr
	| (MINUS | PLUS | NOT) expr
	| REVERSE IDENT
	| expr EXP expr
	| expr (STAR | DIV | MOD) expr
	| expr (MINUS | PLUS) expr
	| expr MODEL expr
	| expr (EQ | NE | GE | GT | LE | LT) expr
	| expr COALESCE expr
	| expr (AND | OR) expr
	| LPAREN expr RPAREN
	| expr IN expr
	| list
	| nestedPipeline;

literal:
	NA
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
		nl* (assign | multiAssign | exprCall) (
			COMMA nl* (assign | multiAssign | exprCall)
		)* COMMA? nl?
	)? RBRACKET;

funcCall: IDENT FUNCTION_CALL funcCallParam*;
funcCallParam: namedArg | assign | multiAssign | exprCall;
namedArg: IDENT COLON (assign | exprCall);
assign: IDENT ASSIGN exprCall;
multiAssign: list ASSIGN exprCall;

pipeline: exprCall (nl funcCall)* nl;
inlinePipeline: exprCall (BAR funcCall)*;
nestedPipeline:
	LPAREN nl* (pipeline | inlinePipeline) nl* RPAREN;