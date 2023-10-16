// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package bytefeeder // preludioParser
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type preludioParser struct {
	*antlr.BaseParser
}

var preludioparserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func preludioparserParserInit() {
	staticData := &preludioparserParserStaticData
	staticData.literalNames = []string{
		"", "'func'", "'prql'", "'ret'", "'->'", "'='", "':='", "'+'", "'-'",
		"'*'", "'^'", "'/'", "'%'", "'~'", "'=='", "'!='", "'>='", "'>'", "'<='",
		"'<'", "'|'", "':'", "','", "'.'", "'$'", "'..'", "'['", "']'", "'('",
		"')'", "'{'", "'}'", "", "", "'_'", "'`'", "'\"'", "'''", "'\"\"\"'",
		"'''''", "'and'", "'or'", "'not'", "'rev'", "'if'", "'do'", "'else'",
		"'for'", "'in'", "'end'", "'?'", "'??'", "'na'", "'@'", "'!'",
	}
	staticData.symbolicNames = []string{
		"", "FUNC", "PRQL", "RET", "ARROW", "ASSIGN", "DECLARE", "PLUS", "MINUS",
		"STAR", "EXP", "DIV", "MOD", "MODEL", "EQ", "NE", "GE", "GT", "LE",
		"LT", "BAR", "COLON", "COMMA", "DOT", "DOLLAR", "RANGE", "LBRACKET",
		"RBRACKET", "LPAREN", "RPAREN", "LBRACE", "RBRACE", "LANG", "RANG",
		"UNDERSCORE", "BACKTICK", "DOUBLE_QUOTE", "SINGLE_QUOTE", "TRIPLE_DOUBLE_QUOTE",
		"TRIPLE_SINGLE_QUOTE", "AND", "OR", "NOT", "REVERSE", "IF", "DO", "ELSE",
		"FOR", "IN", "END", "HELP", "COALESCE", "NA", "INDEXING", "FUNCTION_CALL",
		"WHITESPACE", "NEWLINE", "SINGLE_LINE_COMMENT", "BOOLEAN_LIT", "IDENT",
		"IDENT_START", "IDENT_NEXT", "INTEGER_LIT", "RANGE_LIT", "FLOAT_LIT",
		"STRING_CHAR", "STRING_LIT", "STRING_INTERP_LIT", "STRING_RAW_LIT",
		"STRING_PATH_LIT", "REGEX_LIT", "DATE_LIT", "DURATION_LIT",
	}
	staticData.ruleNames = []string{
		"nl", "program", "programIntro", "funcDef", "funcDefName", "funcDefParams",
		"funcDefParam", "typeDef", "typeTerm", "stmt", "varAssignStmt", "varDeclStmt",
		"ifElseStmt", "forStmt", "helpStmt", "retStmt", "exprCall", "expr",
		"literal", "list", "funcCall", "funcCallParam", "namedArg", "assign",
		"multiAssign", "pipeline", "inlinePipeline", "nestedPipeline",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 72, 415, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 1, 0, 1, 0, 1, 1, 5, 1, 60, 8, 1, 10, 1, 12, 1, 63,
		9, 1, 1, 1, 3, 1, 66, 8, 1, 1, 1, 5, 1, 69, 8, 1, 10, 1, 12, 1, 72, 9,
		1, 1, 1, 1, 1, 3, 1, 76, 8, 1, 1, 1, 5, 1, 79, 8, 1, 10, 1, 12, 1, 82,
		9, 1, 5, 1, 84, 8, 1, 10, 1, 12, 1, 87, 9, 1, 1, 1, 1, 1, 1, 2, 1, 2, 5,
		2, 93, 8, 2, 10, 2, 12, 2, 96, 9, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 5, 3, 106, 8, 3, 10, 3, 12, 3, 109, 9, 3, 1, 3, 1, 3, 5, 3,
		113, 8, 3, 10, 3, 12, 3, 116, 9, 3, 5, 3, 118, 8, 3, 10, 3, 12, 3, 121,
		9, 3, 1, 3, 3, 3, 124, 8, 3, 1, 4, 1, 4, 3, 4, 128, 8, 4, 1, 5, 5, 5, 131,
		8, 5, 10, 5, 12, 5, 134, 9, 5, 1, 6, 1, 6, 3, 6, 138, 8, 6, 1, 6, 3, 6,
		141, 8, 6, 1, 7, 1, 7, 1, 7, 1, 7, 5, 7, 147, 8, 7, 10, 7, 12, 7, 150,
		9, 7, 1, 7, 1, 7, 1, 8, 1, 8, 3, 8, 156, 8, 8, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 3, 9, 165, 8, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1,
		11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 5, 12, 180, 8, 12,
		10, 12, 12, 12, 183, 9, 12, 1, 12, 1, 12, 5, 12, 187, 8, 12, 10, 12, 12,
		12, 190, 9, 12, 5, 12, 192, 8, 12, 10, 12, 12, 12, 195, 9, 12, 1, 12, 1,
		12, 1, 12, 5, 12, 200, 8, 12, 10, 12, 12, 12, 203, 9, 12, 1, 12, 1, 12,
		5, 12, 207, 8, 12, 10, 12, 12, 12, 210, 9, 12, 5, 12, 212, 8, 12, 10, 12,
		12, 12, 215, 9, 12, 1, 12, 3, 12, 218, 8, 12, 1, 12, 3, 12, 221, 8, 12,
		3, 12, 223, 8, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 5,
		13, 232, 8, 13, 10, 13, 12, 13, 235, 9, 13, 1, 13, 1, 13, 5, 13, 239, 8,
		13, 10, 13, 12, 13, 242, 9, 13, 5, 13, 244, 8, 13, 10, 13, 12, 13, 247,
		9, 13, 1, 13, 3, 13, 250, 8, 13, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1,
		15, 1, 16, 1, 16, 3, 16, 260, 8, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17,
		1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 3, 17, 274, 8, 17, 1,
		17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17,
		1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1,
		17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 5, 17, 303, 8, 17, 10, 17, 12, 17,
		306, 9, 17, 1, 18, 1, 18, 1, 19, 1, 19, 5, 19, 312, 8, 19, 10, 19, 12,
		19, 315, 9, 19, 1, 19, 1, 19, 1, 19, 3, 19, 320, 8, 19, 1, 19, 1, 19, 5,
		19, 324, 8, 19, 10, 19, 12, 19, 327, 9, 19, 1, 19, 1, 19, 1, 19, 3, 19,
		332, 8, 19, 5, 19, 334, 8, 19, 10, 19, 12, 19, 337, 9, 19, 1, 19, 3, 19,
		340, 8, 19, 1, 19, 3, 19, 343, 8, 19, 3, 19, 345, 8, 19, 1, 19, 1, 19,
		1, 20, 1, 20, 1, 20, 5, 20, 352, 8, 20, 10, 20, 12, 20, 355, 9, 20, 1,
		21, 1, 21, 1, 21, 1, 21, 3, 21, 361, 8, 21, 1, 22, 1, 22, 1, 22, 1, 22,
		3, 22, 367, 8, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1,
		24, 1, 25, 1, 25, 1, 25, 1, 25, 5, 25, 381, 8, 25, 10, 25, 12, 25, 384,
		9, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 5, 26, 391, 8, 26, 10, 26, 12,
		26, 394, 9, 26, 1, 27, 1, 27, 5, 27, 398, 8, 27, 10, 27, 12, 27, 401, 9,
		27, 1, 27, 1, 27, 3, 27, 405, 8, 27, 1, 27, 5, 27, 408, 8, 27, 10, 27,
		12, 27, 411, 9, 27, 1, 27, 1, 27, 1, 27, 0, 1, 34, 28, 0, 2, 4, 6, 8, 10,
		12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46,
		48, 50, 52, 54, 0, 7, 1, 0, 56, 57, 2, 0, 7, 8, 42, 42, 2, 0, 9, 9, 11,
		12, 1, 0, 7, 8, 1, 0, 14, 19, 1, 0, 40, 41, 4, 0, 52, 52, 58, 59, 62, 64,
		66, 72, 457, 0, 56, 1, 0, 0, 0, 2, 61, 1, 0, 0, 0, 4, 90, 1, 0, 0, 0, 6,
		99, 1, 0, 0, 0, 8, 125, 1, 0, 0, 0, 10, 132, 1, 0, 0, 0, 12, 137, 1, 0,
		0, 0, 14, 142, 1, 0, 0, 0, 16, 153, 1, 0, 0, 0, 18, 164, 1, 0, 0, 0, 20,
		166, 1, 0, 0, 0, 22, 170, 1, 0, 0, 0, 24, 174, 1, 0, 0, 0, 26, 224, 1,
		0, 0, 0, 28, 251, 1, 0, 0, 0, 30, 254, 1, 0, 0, 0, 32, 259, 1, 0, 0, 0,
		34, 273, 1, 0, 0, 0, 36, 307, 1, 0, 0, 0, 38, 309, 1, 0, 0, 0, 40, 348,
		1, 0, 0, 0, 42, 360, 1, 0, 0, 0, 44, 362, 1, 0, 0, 0, 46, 368, 1, 0, 0,
		0, 48, 372, 1, 0, 0, 0, 50, 376, 1, 0, 0, 0, 52, 387, 1, 0, 0, 0, 54, 395,
		1, 0, 0, 0, 56, 57, 7, 0, 0, 0, 57, 1, 1, 0, 0, 0, 58, 60, 3, 0, 0, 0,
		59, 58, 1, 0, 0, 0, 60, 63, 1, 0, 0, 0, 61, 59, 1, 0, 0, 0, 61, 62, 1,
		0, 0, 0, 62, 65, 1, 0, 0, 0, 63, 61, 1, 0, 0, 0, 64, 66, 3, 4, 2, 0, 65,
		64, 1, 0, 0, 0, 65, 66, 1, 0, 0, 0, 66, 70, 1, 0, 0, 0, 67, 69, 3, 0, 0,
		0, 68, 67, 1, 0, 0, 0, 69, 72, 1, 0, 0, 0, 70, 68, 1, 0, 0, 0, 70, 71,
		1, 0, 0, 0, 71, 85, 1, 0, 0, 0, 72, 70, 1, 0, 0, 0, 73, 76, 3, 6, 3, 0,
		74, 76, 3, 18, 9, 0, 75, 73, 1, 0, 0, 0, 75, 74, 1, 0, 0, 0, 76, 80, 1,
		0, 0, 0, 77, 79, 3, 0, 0, 0, 78, 77, 1, 0, 0, 0, 79, 82, 1, 0, 0, 0, 80,
		78, 1, 0, 0, 0, 80, 81, 1, 0, 0, 0, 81, 84, 1, 0, 0, 0, 82, 80, 1, 0, 0,
		0, 83, 75, 1, 0, 0, 0, 84, 87, 1, 0, 0, 0, 85, 83, 1, 0, 0, 0, 85, 86,
		1, 0, 0, 0, 86, 88, 1, 0, 0, 0, 87, 85, 1, 0, 0, 0, 88, 89, 5, 0, 0, 1,
		89, 3, 1, 0, 0, 0, 90, 94, 5, 2, 0, 0, 91, 93, 3, 44, 22, 0, 92, 91, 1,
		0, 0, 0, 93, 96, 1, 0, 0, 0, 94, 92, 1, 0, 0, 0, 94, 95, 1, 0, 0, 0, 95,
		97, 1, 0, 0, 0, 96, 94, 1, 0, 0, 0, 97, 98, 3, 0, 0, 0, 98, 5, 1, 0, 0,
		0, 99, 100, 5, 1, 0, 0, 100, 101, 3, 8, 4, 0, 101, 102, 3, 10, 5, 0, 102,
		123, 5, 4, 0, 0, 103, 124, 3, 32, 16, 0, 104, 106, 3, 0, 0, 0, 105, 104,
		1, 0, 0, 0, 106, 109, 1, 0, 0, 0, 107, 105, 1, 0, 0, 0, 107, 108, 1, 0,
		0, 0, 108, 119, 1, 0, 0, 0, 109, 107, 1, 0, 0, 0, 110, 114, 3, 18, 9, 0,
		111, 113, 3, 0, 0, 0, 112, 111, 1, 0, 0, 0, 113, 116, 1, 0, 0, 0, 114,
		112, 1, 0, 0, 0, 114, 115, 1, 0, 0, 0, 115, 118, 1, 0, 0, 0, 116, 114,
		1, 0, 0, 0, 117, 110, 1, 0, 0, 0, 118, 121, 1, 0, 0, 0, 119, 117, 1, 0,
		0, 0, 119, 120, 1, 0, 0, 0, 120, 122, 1, 0, 0, 0, 121, 119, 1, 0, 0, 0,
		122, 124, 5, 49, 0, 0, 123, 103, 1, 0, 0, 0, 123, 107, 1, 0, 0, 0, 124,
		7, 1, 0, 0, 0, 125, 127, 5, 59, 0, 0, 126, 128, 3, 14, 7, 0, 127, 126,
		1, 0, 0, 0, 127, 128, 1, 0, 0, 0, 128, 9, 1, 0, 0, 0, 129, 131, 3, 12,
		6, 0, 130, 129, 1, 0, 0, 0, 131, 134, 1, 0, 0, 0, 132, 130, 1, 0, 0, 0,
		132, 133, 1, 0, 0, 0, 133, 11, 1, 0, 0, 0, 134, 132, 1, 0, 0, 0, 135, 138,
		3, 44, 22, 0, 136, 138, 5, 59, 0, 0, 137, 135, 1, 0, 0, 0, 137, 136, 1,
		0, 0, 0, 138, 140, 1, 0, 0, 0, 139, 141, 3, 14, 7, 0, 140, 139, 1, 0, 0,
		0, 140, 141, 1, 0, 0, 0, 141, 13, 1, 0, 0, 0, 142, 143, 5, 32, 0, 0, 143,
		144, 3, 16, 8, 0, 144, 148, 5, 20, 0, 0, 145, 147, 3, 16, 8, 0, 146, 145,
		1, 0, 0, 0, 147, 150, 1, 0, 0, 0, 148, 146, 1, 0, 0, 0, 148, 149, 1, 0,
		0, 0, 149, 151, 1, 0, 0, 0, 150, 148, 1, 0, 0, 0, 151, 152, 5, 33, 0, 0,
		152, 15, 1, 0, 0, 0, 153, 155, 5, 59, 0, 0, 154, 156, 3, 14, 7, 0, 155,
		154, 1, 0, 0, 0, 155, 156, 1, 0, 0, 0, 156, 17, 1, 0, 0, 0, 157, 165, 3,
		20, 10, 0, 158, 165, 3, 22, 11, 0, 159, 165, 3, 24, 12, 0, 160, 165, 3,
		26, 13, 0, 161, 165, 3, 28, 14, 0, 162, 165, 3, 30, 15, 0, 163, 165, 3,
		32, 16, 0, 164, 157, 1, 0, 0, 0, 164, 158, 1, 0, 0, 0, 164, 159, 1, 0,
		0, 0, 164, 160, 1, 0, 0, 0, 164, 161, 1, 0, 0, 0, 164, 162, 1, 0, 0, 0,
		164, 163, 1, 0, 0, 0, 165, 19, 1, 0, 0, 0, 166, 167, 5, 59, 0, 0, 167,
		168, 5, 5, 0, 0, 168, 169, 3, 32, 16, 0, 169, 21, 1, 0, 0, 0, 170, 171,
		5, 59, 0, 0, 171, 172, 5, 6, 0, 0, 172, 173, 3, 32, 16, 0, 173, 23, 1,
		0, 0, 0, 174, 175, 5, 44, 0, 0, 175, 176, 3, 32, 16, 0, 176, 222, 5, 45,
		0, 0, 177, 223, 3, 18, 9, 0, 178, 180, 3, 0, 0, 0, 179, 178, 1, 0, 0, 0,
		180, 183, 1, 0, 0, 0, 181, 179, 1, 0, 0, 0, 181, 182, 1, 0, 0, 0, 182,
		193, 1, 0, 0, 0, 183, 181, 1, 0, 0, 0, 184, 188, 3, 18, 9, 0, 185, 187,
		3, 0, 0, 0, 186, 185, 1, 0, 0, 0, 187, 190, 1, 0, 0, 0, 188, 186, 1, 0,
		0, 0, 188, 189, 1, 0, 0, 0, 189, 192, 1, 0, 0, 0, 190, 188, 1, 0, 0, 0,
		191, 184, 1, 0, 0, 0, 192, 195, 1, 0, 0, 0, 193, 191, 1, 0, 0, 0, 193,
		194, 1, 0, 0, 0, 194, 220, 1, 0, 0, 0, 195, 193, 1, 0, 0, 0, 196, 217,
		5, 46, 0, 0, 197, 218, 3, 18, 9, 0, 198, 200, 3, 0, 0, 0, 199, 198, 1,
		0, 0, 0, 200, 203, 1, 0, 0, 0, 201, 199, 1, 0, 0, 0, 201, 202, 1, 0, 0,
		0, 202, 213, 1, 0, 0, 0, 203, 201, 1, 0, 0, 0, 204, 208, 3, 18, 9, 0, 205,
		207, 3, 0, 0, 0, 206, 205, 1, 0, 0, 0, 207, 210, 1, 0, 0, 0, 208, 206,
		1, 0, 0, 0, 208, 209, 1, 0, 0, 0, 209, 212, 1, 0, 0, 0, 210, 208, 1, 0,
		0, 0, 211, 204, 1, 0, 0, 0, 212, 215, 1, 0, 0, 0, 213, 211, 1, 0, 0, 0,
		213, 214, 1, 0, 0, 0, 214, 216, 1, 0, 0, 0, 215, 213, 1, 0, 0, 0, 216,
		218, 5, 49, 0, 0, 217, 197, 1, 0, 0, 0, 217, 201, 1, 0, 0, 0, 218, 221,
		1, 0, 0, 0, 219, 221, 5, 49, 0, 0, 220, 196, 1, 0, 0, 0, 220, 219, 1, 0,
		0, 0, 221, 223, 1, 0, 0, 0, 222, 177, 1, 0, 0, 0, 222, 181, 1, 0, 0, 0,
		223, 25, 1, 0, 0, 0, 224, 225, 5, 47, 0, 0, 225, 226, 5, 59, 0, 0, 226,
		227, 5, 48, 0, 0, 227, 228, 3, 32, 16, 0, 228, 249, 5, 45, 0, 0, 229, 250,
		3, 18, 9, 0, 230, 232, 3, 0, 0, 0, 231, 230, 1, 0, 0, 0, 232, 235, 1, 0,
		0, 0, 233, 231, 1, 0, 0, 0, 233, 234, 1, 0, 0, 0, 234, 245, 1, 0, 0, 0,
		235, 233, 1, 0, 0, 0, 236, 240, 3, 18, 9, 0, 237, 239, 3, 0, 0, 0, 238,
		237, 1, 0, 0, 0, 239, 242, 1, 0, 0, 0, 240, 238, 1, 0, 0, 0, 240, 241,
		1, 0, 0, 0, 241, 244, 1, 0, 0, 0, 242, 240, 1, 0, 0, 0, 243, 236, 1, 0,
		0, 0, 244, 247, 1, 0, 0, 0, 245, 243, 1, 0, 0, 0, 245, 246, 1, 0, 0, 0,
		246, 248, 1, 0, 0, 0, 247, 245, 1, 0, 0, 0, 248, 250, 5, 49, 0, 0, 249,
		229, 1, 0, 0, 0, 249, 233, 1, 0, 0, 0, 250, 27, 1, 0, 0, 0, 251, 252, 5,
		50, 0, 0, 252, 253, 5, 59, 0, 0, 253, 29, 1, 0, 0, 0, 254, 255, 5, 3, 0,
		0, 255, 256, 3, 32, 16, 0, 256, 31, 1, 0, 0, 0, 257, 260, 3, 34, 17, 0,
		258, 260, 3, 40, 20, 0, 259, 257, 1, 0, 0, 0, 259, 258, 1, 0, 0, 0, 260,
		33, 1, 0, 0, 0, 261, 262, 6, 17, -1, 0, 262, 274, 3, 36, 18, 0, 263, 264,
		7, 1, 0, 0, 264, 274, 3, 34, 17, 13, 265, 266, 5, 43, 0, 0, 266, 274, 5,
		59, 0, 0, 267, 268, 5, 28, 0, 0, 268, 269, 3, 34, 17, 0, 269, 270, 5, 29,
		0, 0, 270, 274, 1, 0, 0, 0, 271, 274, 3, 38, 19, 0, 272, 274, 3, 54, 27,
		0, 273, 261, 1, 0, 0, 0, 273, 263, 1, 0, 0, 0, 273, 265, 1, 0, 0, 0, 273,
		267, 1, 0, 0, 0, 273, 271, 1, 0, 0, 0, 273, 272, 1, 0, 0, 0, 274, 304,
		1, 0, 0, 0, 275, 276, 10, 14, 0, 0, 276, 277, 5, 53, 0, 0, 277, 303, 3,
		34, 17, 15, 278, 279, 10, 11, 0, 0, 279, 280, 5, 10, 0, 0, 280, 303, 3,
		34, 17, 12, 281, 282, 10, 10, 0, 0, 282, 283, 7, 2, 0, 0, 283, 303, 3,
		34, 17, 11, 284, 285, 10, 9, 0, 0, 285, 286, 7, 3, 0, 0, 286, 303, 3, 34,
		17, 10, 287, 288, 10, 8, 0, 0, 288, 289, 5, 13, 0, 0, 289, 303, 3, 34,
		17, 9, 290, 291, 10, 7, 0, 0, 291, 292, 7, 4, 0, 0, 292, 303, 3, 34, 17,
		8, 293, 294, 10, 6, 0, 0, 294, 295, 5, 51, 0, 0, 295, 303, 3, 34, 17, 7,
		296, 297, 10, 5, 0, 0, 297, 298, 7, 5, 0, 0, 298, 303, 3, 34, 17, 6, 299,
		300, 10, 3, 0, 0, 300, 301, 5, 48, 0, 0, 301, 303, 3, 34, 17, 4, 302, 275,
		1, 0, 0, 0, 302, 278, 1, 0, 0, 0, 302, 281, 1, 0, 0, 0, 302, 284, 1, 0,
		0, 0, 302, 287, 1, 0, 0, 0, 302, 290, 1, 0, 0, 0, 302, 293, 1, 0, 0, 0,
		302, 296, 1, 0, 0, 0, 302, 299, 1, 0, 0, 0, 303, 306, 1, 0, 0, 0, 304,
		302, 1, 0, 0, 0, 304, 305, 1, 0, 0, 0, 305, 35, 1, 0, 0, 0, 306, 304, 1,
		0, 0, 0, 307, 308, 7, 6, 0, 0, 308, 37, 1, 0, 0, 0, 309, 344, 5, 26, 0,
		0, 310, 312, 3, 0, 0, 0, 311, 310, 1, 0, 0, 0, 312, 315, 1, 0, 0, 0, 313,
		311, 1, 0, 0, 0, 313, 314, 1, 0, 0, 0, 314, 319, 1, 0, 0, 0, 315, 313,
		1, 0, 0, 0, 316, 320, 3, 46, 23, 0, 317, 320, 3, 48, 24, 0, 318, 320, 3,
		32, 16, 0, 319, 316, 1, 0, 0, 0, 319, 317, 1, 0, 0, 0, 319, 318, 1, 0,
		0, 0, 320, 335, 1, 0, 0, 0, 321, 325, 5, 22, 0, 0, 322, 324, 3, 0, 0, 0,
		323, 322, 1, 0, 0, 0, 324, 327, 1, 0, 0, 0, 325, 323, 1, 0, 0, 0, 325,
		326, 1, 0, 0, 0, 326, 331, 1, 0, 0, 0, 327, 325, 1, 0, 0, 0, 328, 332,
		3, 46, 23, 0, 329, 332, 3, 48, 24, 0, 330, 332, 3, 32, 16, 0, 331, 328,
		1, 0, 0, 0, 331, 329, 1, 0, 0, 0, 331, 330, 1, 0, 0, 0, 332, 334, 1, 0,
		0, 0, 333, 321, 1, 0, 0, 0, 334, 337, 1, 0, 0, 0, 335, 333, 1, 0, 0, 0,
		335, 336, 1, 0, 0, 0, 336, 339, 1, 0, 0, 0, 337, 335, 1, 0, 0, 0, 338,
		340, 5, 22, 0, 0, 339, 338, 1, 0, 0, 0, 339, 340, 1, 0, 0, 0, 340, 342,
		1, 0, 0, 0, 341, 343, 3, 0, 0, 0, 342, 341, 1, 0, 0, 0, 342, 343, 1, 0,
		0, 0, 343, 345, 1, 0, 0, 0, 344, 313, 1, 0, 0, 0, 344, 345, 1, 0, 0, 0,
		345, 346, 1, 0, 0, 0, 346, 347, 5, 27, 0, 0, 347, 39, 1, 0, 0, 0, 348,
		349, 5, 59, 0, 0, 349, 353, 5, 54, 0, 0, 350, 352, 3, 42, 21, 0, 351, 350,
		1, 0, 0, 0, 352, 355, 1, 0, 0, 0, 353, 351, 1, 0, 0, 0, 353, 354, 1, 0,
		0, 0, 354, 41, 1, 0, 0, 0, 355, 353, 1, 0, 0, 0, 356, 361, 3, 44, 22, 0,
		357, 361, 3, 46, 23, 0, 358, 361, 3, 48, 24, 0, 359, 361, 3, 32, 16, 0,
		360, 356, 1, 0, 0, 0, 360, 357, 1, 0, 0, 0, 360, 358, 1, 0, 0, 0, 360,
		359, 1, 0, 0, 0, 361, 43, 1, 0, 0, 0, 362, 363, 5, 59, 0, 0, 363, 366,
		5, 21, 0, 0, 364, 367, 3, 46, 23, 0, 365, 367, 3, 32, 16, 0, 366, 364,
		1, 0, 0, 0, 366, 365, 1, 0, 0, 0, 367, 45, 1, 0, 0, 0, 368, 369, 5, 59,
		0, 0, 369, 370, 5, 5, 0, 0, 370, 371, 3, 32, 16, 0, 371, 47, 1, 0, 0, 0,
		372, 373, 3, 38, 19, 0, 373, 374, 5, 5, 0, 0, 374, 375, 3, 32, 16, 0, 375,
		49, 1, 0, 0, 0, 376, 382, 3, 32, 16, 0, 377, 378, 3, 0, 0, 0, 378, 379,
		3, 40, 20, 0, 379, 381, 1, 0, 0, 0, 380, 377, 1, 0, 0, 0, 381, 384, 1,
		0, 0, 0, 382, 380, 1, 0, 0, 0, 382, 383, 1, 0, 0, 0, 383, 385, 1, 0, 0,
		0, 384, 382, 1, 0, 0, 0, 385, 386, 3, 0, 0, 0, 386, 51, 1, 0, 0, 0, 387,
		392, 3, 32, 16, 0, 388, 389, 5, 20, 0, 0, 389, 391, 3, 40, 20, 0, 390,
		388, 1, 0, 0, 0, 391, 394, 1, 0, 0, 0, 392, 390, 1, 0, 0, 0, 392, 393,
		1, 0, 0, 0, 393, 53, 1, 0, 0, 0, 394, 392, 1, 0, 0, 0, 395, 399, 5, 28,
		0, 0, 396, 398, 3, 0, 0, 0, 397, 396, 1, 0, 0, 0, 398, 401, 1, 0, 0, 0,
		399, 397, 1, 0, 0, 0, 399, 400, 1, 0, 0, 0, 400, 404, 1, 0, 0, 0, 401,
		399, 1, 0, 0, 0, 402, 405, 3, 50, 25, 0, 403, 405, 3, 52, 26, 0, 404, 402,
		1, 0, 0, 0, 404, 403, 1, 0, 0, 0, 405, 409, 1, 0, 0, 0, 406, 408, 3, 0,
		0, 0, 407, 406, 1, 0, 0, 0, 408, 411, 1, 0, 0, 0, 409, 407, 1, 0, 0, 0,
		409, 410, 1, 0, 0, 0, 410, 412, 1, 0, 0, 0, 411, 409, 1, 0, 0, 0, 412,
		413, 5, 29, 0, 0, 413, 55, 1, 0, 0, 0, 51, 61, 65, 70, 75, 80, 85, 94,
		107, 114, 119, 123, 127, 132, 137, 140, 148, 155, 164, 181, 188, 193, 201,
		208, 213, 217, 220, 222, 233, 240, 245, 249, 259, 273, 302, 304, 313, 319,
		325, 331, 335, 339, 342, 344, 353, 360, 366, 382, 392, 399, 404, 409,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// preludioParserInit initializes any static state used to implement preludioParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewpreludioParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func PreludioParserInit() {
	staticData := &preludioparserParserStaticData
	staticData.once.Do(preludioparserParserInit)
}

// NewpreludioParser produces a new parser instance for the optional input antlr.TokenStream.
func NewpreludioParser(input antlr.TokenStream) *preludioParser {
	PreludioParserInit()
	this := new(preludioParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &preludioparserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "java-escape"

	return this
}

// preludioParser tokens.
const (
	preludioParserEOF                 = antlr.TokenEOF
	preludioParserFUNC                = 1
	preludioParserPRQL                = 2
	preludioParserRET                 = 3
	preludioParserARROW               = 4
	preludioParserASSIGN              = 5
	preludioParserDECLARE             = 6
	preludioParserPLUS                = 7
	preludioParserMINUS               = 8
	preludioParserSTAR                = 9
	preludioParserEXP                 = 10
	preludioParserDIV                 = 11
	preludioParserMOD                 = 12
	preludioParserMODEL               = 13
	preludioParserEQ                  = 14
	preludioParserNE                  = 15
	preludioParserGE                  = 16
	preludioParserGT                  = 17
	preludioParserLE                  = 18
	preludioParserLT                  = 19
	preludioParserBAR                 = 20
	preludioParserCOLON               = 21
	preludioParserCOMMA               = 22
	preludioParserDOT                 = 23
	preludioParserDOLLAR              = 24
	preludioParserRANGE               = 25
	preludioParserLBRACKET            = 26
	preludioParserRBRACKET            = 27
	preludioParserLPAREN              = 28
	preludioParserRPAREN              = 29
	preludioParserLBRACE              = 30
	preludioParserRBRACE              = 31
	preludioParserLANG                = 32
	preludioParserRANG                = 33
	preludioParserUNDERSCORE          = 34
	preludioParserBACKTICK            = 35
	preludioParserDOUBLE_QUOTE        = 36
	preludioParserSINGLE_QUOTE        = 37
	preludioParserTRIPLE_DOUBLE_QUOTE = 38
	preludioParserTRIPLE_SINGLE_QUOTE = 39
	preludioParserAND                 = 40
	preludioParserOR                  = 41
	preludioParserNOT                 = 42
	preludioParserREVERSE             = 43
	preludioParserIF                  = 44
	preludioParserDO                  = 45
	preludioParserELSE                = 46
	preludioParserFOR                 = 47
	preludioParserIN                  = 48
	preludioParserEND                 = 49
	preludioParserHELP                = 50
	preludioParserCOALESCE            = 51
	preludioParserNA                  = 52
	preludioParserINDEXING            = 53
	preludioParserFUNCTION_CALL       = 54
	preludioParserWHITESPACE          = 55
	preludioParserNEWLINE             = 56
	preludioParserSINGLE_LINE_COMMENT = 57
	preludioParserBOOLEAN_LIT         = 58
	preludioParserIDENT               = 59
	preludioParserIDENT_START         = 60
	preludioParserIDENT_NEXT          = 61
	preludioParserINTEGER_LIT         = 62
	preludioParserRANGE_LIT           = 63
	preludioParserFLOAT_LIT           = 64
	preludioParserSTRING_CHAR         = 65
	preludioParserSTRING_LIT          = 66
	preludioParserSTRING_INTERP_LIT   = 67
	preludioParserSTRING_RAW_LIT      = 68
	preludioParserSTRING_PATH_LIT     = 69
	preludioParserREGEX_LIT           = 70
	preludioParserDATE_LIT            = 71
	preludioParserDURATION_LIT        = 72
)

// preludioParser rules.
const (
	preludioParserRULE_nl             = 0
	preludioParserRULE_program        = 1
	preludioParserRULE_programIntro   = 2
	preludioParserRULE_funcDef        = 3
	preludioParserRULE_funcDefName    = 4
	preludioParserRULE_funcDefParams  = 5
	preludioParserRULE_funcDefParam   = 6
	preludioParserRULE_typeDef        = 7
	preludioParserRULE_typeTerm       = 8
	preludioParserRULE_stmt           = 9
	preludioParserRULE_varAssignStmt  = 10
	preludioParserRULE_varDeclStmt    = 11
	preludioParserRULE_ifElseStmt     = 12
	preludioParserRULE_forStmt        = 13
	preludioParserRULE_helpStmt       = 14
	preludioParserRULE_retStmt        = 15
	preludioParserRULE_exprCall       = 16
	preludioParserRULE_expr           = 17
	preludioParserRULE_literal        = 18
	preludioParserRULE_list           = 19
	preludioParserRULE_funcCall       = 20
	preludioParserRULE_funcCallParam  = 21
	preludioParserRULE_namedArg       = 22
	preludioParserRULE_assign         = 23
	preludioParserRULE_multiAssign    = 24
	preludioParserRULE_pipeline       = 25
	preludioParserRULE_inlinePipeline = 26
	preludioParserRULE_nestedPipeline = 27
)

// INlContext is an interface to support dynamic dispatch.
type INlContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNlContext differentiates from other interfaces.
	IsNlContext()
}

type NlContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNlContext() *NlContext {
	var p = new(NlContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = preludioParserRULE_nl
	return p
}

func (*NlContext) IsNlContext() {}

func NewNlContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NlContext {
	var p = new(NlContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = preludioParserRULE_nl

	return p
}

func (s *NlContext) GetParser() antlr.Parser { return s.parser }

func (s *NlContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(preludioParserNEWLINE, 0)
}

func (s *NlContext) SINGLE_LINE_COMMENT() antlr.TerminalNode {
	return s.GetToken(preludioParserSINGLE_LINE_COMMENT, 0)
}

func (s *NlContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NlContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NlContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioParserListener); ok {
		listenerT.EnterNl(s)
	}
}

func (s *NlContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioParserListener); ok {
		listenerT.ExitNl(s)
	}
}

func (p *preludioParser) Nl() (localctx INlContext) {
	this := p
	_ = this

	localctx = NewNlContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, preludioParserRULE_nl)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(56)
		_la = p.GetTokenStream().LA(1)

		if !(_la == preludioParserNEWLINE || _la == preludioParserSINGLE_LINE_COMMENT) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = preludioParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = preludioParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(preludioParserEOF, 0)
}

func (s *ProgramContext) AllNl() []INlContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INlContext); ok {
			len++
		}
	}

	tst := make([]INlContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INlContext); ok {
			tst[i] = t.(INlContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Nl(i int) INlContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INlContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(INlContext)
}

func (s *ProgramContext) ProgramIntro() IProgramIntroContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProgramIntroContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProgramIntroContext)
}

func (s *ProgramContext) AllFuncDef() []IFuncDefContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFuncDefContext); ok {
			len++
		}
	}

	tst := make([]IFuncDefContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFuncDefContext); ok {
			tst[i] = t.(IFuncDefContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) FuncDef(i int) IFuncDefContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncDefContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncDefContext)
}

func (s *ProgramContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioParserListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioParserListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *preludioParser) Program() (localctx IProgramContext) {
	this := p
	_ = this

	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, preludioParserRULE_program)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(58)
				p.Nl()
			}

		}
		p.SetState(63)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())
	}
	p.SetState(65)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == preludioParserPRQL {
		{
			p.SetState(64)
			p.ProgramIntro()
		}

	}
	p.SetState(70)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == preludioParserNEWLINE || _la == preludioParserSINGLE_LINE_COMMENT {
		{
			p.SetState(67)
			p.Nl()
		}

		p.SetState(72)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(85)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-3741193866288561782) != 0 || (int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&509) != 0 {
		p.SetState(75)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case preludioParserFUNC:
			{
				p.SetState(73)
				p.FuncDef()
			}

		case preludioParserRET, preludioParserPLUS, preludioParserMINUS, preludioParserLBRACKET, preludioParserLPAREN, preludioParserNOT, preludioParserREVERSE, preludioParserIF, preludioParserFOR, preludioParserHELP, preludioParserNA, preludioParserBOOLEAN_LIT, preludioParserIDENT, preludioParserINTEGER_LIT, preludioParserRANGE_LIT, preludioParserFLOAT_LIT, preludioParserSTRING_LIT, preludioParserSTRING_INTERP_LIT, preludioParserSTRING_RAW_LIT, preludioParserSTRING_PATH_LIT, preludioParserREGEX_LIT, preludioParserDATE_LIT, preludioParserDURATION_LIT:
			{
				p.SetState(74)
				p.Stmt()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}
		p.SetState(80)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == preludioParserNEWLINE || _la == preludioParserSINGLE_LINE_COMMENT {
			{
				p.SetState(77)
				p.Nl()
			}

			p.SetState(82)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

		p.SetState(87)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(88)
		p.Match(preludioParserEOF)
	}

	return localctx
}

// IProgramIntroContext is an interface to support dynamic dispatch.
type IProgramIntroContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgramIntroContext differentiates from other interfaces.
	IsProgramIntroContext()
}

type ProgramIntroContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramIntroContext() *ProgramIntroContext {
	var p = new(ProgramIntroContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = preludioParserRULE_programIntro
	return p
}

func (*ProgramIntroContext) IsProgramIntroContext() {}

func NewProgramIntroContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramIntroContext {
	var p = new(ProgramIntroContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = preludioParserRULE_programIntro

	return p
}

func (s *ProgramIntroContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramIntroContext) PRQL() antlr.TerminalNode {
	return s.GetToken(preludioParserPRQL, 0)
}

func (s *ProgramIntroContext) Nl() INlContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INlContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INlContext)
}

func (s *ProgramIntroContext) AllNamedArg() []INamedArgContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INamedArgContext); ok {
			len++
		}
	}

	tst := make([]INamedArgContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INamedArgContext); ok {
			tst[i] = t.(INamedArgContext)
			i++
		}
	}

	return tst
}

func (s *ProgramIntroContext) NamedArg(i int) INamedArgContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INamedArgContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(INamedArgContext)
}

func (s *ProgramIntroContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramIntroContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramIntroContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioParserListener); ok {
		listenerT.EnterProgramIntro(s)
	}
}

func (s *ProgramIntroContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioParserListener); ok {
		listenerT.ExitProgramIntro(s)
	}
}

func (p *preludioParser) ProgramIntro() (localctx IProgramIntroContext) {
	this := p
	_ = this

	localctx = NewProgramIntroContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, preludioParserRULE_programIntro)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(90)
		p.Match(preludioParserPRQL)
	}
	p.SetState(94)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == preludioParserIDENT {
		{
			p.SetState(91)
			p.NamedArg()
		}

		p.SetState(96)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(97)
		p.Nl()
	}

	return localctx
}

// IFuncDefContext is an interface to support dynamic dispatch.
type IFuncDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncDefContext differentiates from other interfaces.
	IsFuncDefContext()
}

type FuncDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncDefContext() *FuncDefContext {
	var p = new(FuncDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = preludioParserRULE_funcDef
	return p
}

func (*FuncDefContext) IsFuncDefContext() {}

func NewFuncDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncDefContext {
	var p = new(FuncDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = preludioParserRULE_funcDef

	return p
}

func (s *FuncDefContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncDefContext) FUNC() antlr.TerminalNode {
	return s.GetToken(preludioParserFUNC, 0)
}

func (s *FuncDefContext) FuncDefName() IFuncDefNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncDefNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncDefNameContext)
}

func (s *FuncDefContext) FuncDefParams() IFuncDefParamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncDefParamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncDefParamsContext)
}

func (s *FuncDefContext) ARROW() antlr.TerminalNode {
	return s.GetToken(preludioParserARROW, 0)
}

func (s *FuncDefContext) ExprCall() IExprCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprCallContext)
}

func (s *FuncDefContext) END() antlr.TerminalNode {
	return s.GetToken(preludioParserEND, 0)
}

func (s *FuncDefContext) AllNl() []INlContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INlContext); ok {
			len++
		}
	}

	tst := make([]INlContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INlContext); ok {
			tst[i] = t.(INlContext)
			i++
		}
	}

	return tst
}

func (s *FuncDefContext) Nl(i int) INlContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INlContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(INlContext)
}

func (s *FuncDefContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *FuncDefContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *FuncDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioParserListener); ok {
		listenerT.EnterFuncDef(s)
	}
}

func (s *FuncDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioParserListener); ok {
		listenerT.ExitFuncDef(s)
	}
}

func (p *preludioParser) FuncDef() (localctx IFuncDefContext) {
	this := p
	_ = this

	localctx = NewFuncDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, preludioParserRULE_funcDef)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(99)
		p.Match(preludioParserFUNC)
	}
	{
		p.SetState(100)
		p.FuncDefName()
	}
	{
		p.SetState(101)
		p.FuncDefParams()
	}
	{
		p.SetState(102)
		p.Match(preludioParserARROW)
	}
	p.SetState(123)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(103)
			p.ExprCall()
		}

	case 2:
		p.SetState(107)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == preludioParserNEWLINE || _la == preludioParserSINGLE_LINE_COMMENT {
			{
				p.SetState(104)
				p.Nl()
			}

			p.SetState(109)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(119)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-3741193866288561784) != 0 || (int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&509) != 0 {
			{
				p.SetState(110)
				p.Stmt()
			}
			p.SetState(114)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == preludioParserNEWLINE || _la == preludioParserSINGLE_LINE_COMMENT {
				{
					p.SetState(111)
					p.Nl()
				}

				p.SetState(116)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

			p.SetState(121)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(122)
			p.Match(preludioParserEND)
		}

	}

	return localctx
}

// IFuncDefNameContext is an interface to support dynamic dispatch.
type IFuncDefNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncDefNameContext differentiates from other interfaces.
	IsFuncDefNameContext()
}

type FuncDefNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncDefNameContext() *FuncDefNameContext {
	var p = new(FuncDefNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = preludioParserRULE_funcDefName
	return p
}

func (*FuncDefNameContext) IsFuncDefNameContext() {}

func NewFuncDefNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncDefNameContext {
	var p = new(FuncDefNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = preludioParserRULE_funcDefName

	return p
}

func (s *FuncDefNameContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncDefNameContext) IDENT() antlr.TerminalNode {
	return s.GetToken(preludioParserIDENT, 0)
}

func (s *FuncDefNameContext) TypeDef() ITypeDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeDefContext)
}

func (s *FuncDefNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncDefNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncDefNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioParserListener); ok {
		listenerT.EnterFuncDefName(s)
	}
}

func (s *FuncDefNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioParserListener); ok {
		listenerT.ExitFuncDefName(s)
	}
}

func (p *preludioParser) FuncDefName() (localctx IFuncDefNameContext) {
	this := p
	_ = this

	localctx = NewFuncDefNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, preludioParserRULE_funcDefName)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(125)
		p.Match(preludioParserIDENT)
	}
	p.SetState(127)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == preludioParserLANG {
		{
			p.SetState(126)
			p.TypeDef()
		}

	}

	return localctx
}

// IFuncDefParamsContext is an interface to support dynamic dispatch.
type IFuncDefParamsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncDefParamsContext differentiates from other interfaces.
	IsFuncDefParamsContext()
}

type FuncDefParamsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncDefParamsContext() *FuncDefParamsContext {
	var p = new(FuncDefParamsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = preludioParserRULE_funcDefParams
	return p
}

func (*FuncDefParamsContext) IsFuncDefParamsContext() {}

func NewFuncDefParamsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncDefParamsContext {
	var p = new(FuncDefParamsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = preludioParserRULE_funcDefParams

	return p
}

func (s *FuncDefParamsContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncDefParamsContext) AllFuncDefParam() []IFuncDefParamContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFuncDefParamContext); ok {
			len++
		}
	}

	tst := make([]IFuncDefParamContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFuncDefParamContext); ok {
			tst[i] = t.(IFuncDefParamContext)
			i++
		}
	}

	return tst
}

func (s *FuncDefParamsContext) FuncDefParam(i int) IFuncDefParamContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncDefParamContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncDefParamContext)
}

func (s *FuncDefParamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncDefParamsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncDefParamsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioParserListener); ok {
		listenerT.EnterFuncDefParams(s)
	}
}

func (s *FuncDefParamsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioParserListener); ok {
		listenerT.ExitFuncDefParams(s)
	}
}

func (p *preludioParser) FuncDefParams() (localctx IFuncDefParamsContext) {
	this := p
	_ = this

	localctx = NewFuncDefParamsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, preludioParserRULE_funcDefParams)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(132)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == preludioParserIDENT {
		{
			p.SetState(129)
			p.FuncDefParam()
		}

		p.SetState(134)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFuncDefParamContext is an interface to support dynamic dispatch.
type IFuncDefParamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncDefParamContext differentiates from other interfaces.
	IsFuncDefParamContext()
}

type FuncDefParamContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncDefParamContext() *FuncDefParamContext {
	var p = new(FuncDefParamContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = preludioParserRULE_funcDefParam
	return p
}

func (*FuncDefParamContext) IsFuncDefParamContext() {}

func NewFuncDefParamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncDefParamContext {
	var p = new(FuncDefParamContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = preludioParserRULE_funcDefParam

	return p
}

func (s *FuncDefParamContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncDefParamContext) NamedArg() INamedArgContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INamedArgContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INamedArgContext)
}

func (s *FuncDefParamContext) IDENT() antlr.TerminalNode {
	return s.GetToken(preludioParserIDENT, 0)
}

func (s *FuncDefParamContext) TypeDef() ITypeDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeDefContext)
}

func (s *FuncDefParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncDefParamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncDefParamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioParserListener); ok {
		listenerT.EnterFuncDefParam(s)
	}
}

func (s *FuncDefParamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioParserListener); ok {
		listenerT.ExitFuncDefParam(s)
	}
}

func (p *preludioParser) FuncDefParam() (localctx IFuncDefParamContext) {
	this := p
	_ = this

	localctx = NewFuncDefParamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, preludioParserRULE_funcDefParam)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(137)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(135)
			p.NamedArg()
		}

	case 2:
		{
			p.SetState(136)
			p.Match(preludioParserIDENT)
		}

	}
	p.SetState(140)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == preludioParserLANG {
		{
			p.SetState(139)
			p.TypeDef()
		}

	}

	return localctx
}

// ITypeDefContext is an interface to support dynamic dispatch.
type ITypeDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeDefContext differentiates from other interfaces.
	IsTypeDefContext()
}

type TypeDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeDefContext() *TypeDefContext {
	var p = new(TypeDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = preludioParserRULE_typeDef
	return p
}

func (*TypeDefContext) IsTypeDefContext() {}

func NewTypeDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeDefContext {
	var p = new(TypeDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = preludioParserRULE_typeDef

	return p
}

func (s *TypeDefContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeDefContext) LANG() antlr.TerminalNode {
	return s.GetToken(preludioParserLANG, 0)
}

func (s *TypeDefContext) AllTypeTerm() []ITypeTermContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITypeTermContext); ok {
			len++
		}
	}

	tst := make([]ITypeTermContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITypeTermContext); ok {
			tst[i] = t.(ITypeTermContext)
			i++
		}
	}

	return tst
}

func (s *TypeDefContext) TypeTerm(i int) ITypeTermContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeTermContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeTermContext)
}

func (s *TypeDefContext) BAR() antlr.TerminalNode {
	return s.GetToken(preludioParserBAR, 0)
}

func (s *TypeDefContext) RANG() antlr.TerminalNode {
	return s.GetToken(preludioParserRANG, 0)
}

func (s *TypeDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioParserListener); ok {
		listenerT.EnterTypeDef(s)
	}
}

func (s *TypeDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioParserListener); ok {
		listenerT.ExitTypeDef(s)
	}
}

func (p *preludioParser) TypeDef() (localctx ITypeDefContext) {
	this := p
	_ = this

	localctx = NewTypeDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, preludioParserRULE_typeDef)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(142)
		p.Match(preludioParserLANG)
	}
	{
		p.SetState(143)
		p.TypeTerm()
	}
	{
		p.SetState(144)
		p.Match(preludioParserBAR)
	}
	p.SetState(148)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == preludioParserIDENT {
		{
			p.SetState(145)
			p.TypeTerm()
		}

		p.SetState(150)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(151)
		p.Match(preludioParserRANG)
	}

	return localctx
}

// ITypeTermContext is an interface to support dynamic dispatch.
type ITypeTermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeTermContext differentiates from other interfaces.
	IsTypeTermContext()
}

type TypeTermContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeTermContext() *TypeTermContext {
	var p = new(TypeTermContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = preludioParserRULE_typeTerm
	return p
}

func (*TypeTermContext) IsTypeTermContext() {}

func NewTypeTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeTermContext {
	var p = new(TypeTermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = preludioParserRULE_typeTerm

	return p
}

func (s *TypeTermContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeTermContext) IDENT() antlr.TerminalNode {
	return s.GetToken(preludioParserIDENT, 0)
}

func (s *TypeTermContext) TypeDef() ITypeDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeDefContext)
}

func (s *TypeTermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeTermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeTermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioParserListener); ok {
		listenerT.EnterTypeTerm(s)
	}
}

func (s *TypeTermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioParserListener); ok {
		listenerT.ExitTypeTerm(s)
	}
}

func (p *preludioParser) TypeTerm() (localctx ITypeTermContext) {
	this := p
	_ = this

	localctx = NewTypeTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, preludioParserRULE_typeTerm)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(153)
		p.Match(preludioParserIDENT)
	}
	p.SetState(155)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == preludioParserLANG {
		{
			p.SetState(154)
			p.TypeDef()
		}

	}

	return localctx
}

// IStmtContext is an interface to support dynamic dispatch.
type IStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStmtContext differentiates from other interfaces.
	IsStmtContext()
}

type StmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStmtContext() *StmtContext {
	var p = new(StmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = preludioParserRULE_stmt
	return p
}

func (*StmtContext) IsStmtContext() {}

func NewStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StmtContext {
	var p = new(StmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = preludioParserRULE_stmt

	return p
}

func (s *StmtContext) GetParser() antlr.Parser { return s.parser }

func (s *StmtContext) VarAssignStmt() IVarAssignStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarAssignStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarAssignStmtContext)
}

func (s *StmtContext) VarDeclStmt() IVarDeclStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarDeclStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarDeclStmtContext)
}

func (s *StmtContext) IfElseStmt() IIfElseStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfElseStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfElseStmtContext)
}

func (s *StmtContext) ForStmt() IForStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForStmtContext)
}

func (s *StmtContext) HelpStmt() IHelpStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHelpStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IHelpStmtContext)
}

func (s *StmtContext) RetStmt() IRetStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRetStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRetStmtContext)
}

func (s *StmtContext) ExprCall() IExprCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprCallContext)
}

func (s *StmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioParserListener); ok {
		listenerT.EnterStmt(s)
	}
}

func (s *StmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioParserListener); ok {
		listenerT.ExitStmt(s)
	}
}

func (p *preludioParser) Stmt() (localctx IStmtContext) {
	this := p
	_ = this

	localctx = NewStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, preludioParserRULE_stmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(164)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(157)
			p.VarAssignStmt()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(158)
			p.VarDeclStmt()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(159)
			p.IfElseStmt()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(160)
			p.ForStmt()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(161)
			p.HelpStmt()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(162)
			p.RetStmt()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(163)
			p.ExprCall()
		}

	}

	return localctx
}

// IVarAssignStmtContext is an interface to support dynamic dispatch.
type IVarAssignStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVarAssignStmtContext differentiates from other interfaces.
	IsVarAssignStmtContext()
}

type VarAssignStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarAssignStmtContext() *VarAssignStmtContext {
	var p = new(VarAssignStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = preludioParserRULE_varAssignStmt
	return p
}

func (*VarAssignStmtContext) IsVarAssignStmtContext() {}

func NewVarAssignStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarAssignStmtContext {
	var p = new(VarAssignStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = preludioParserRULE_varAssignStmt

	return p
}

func (s *VarAssignStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *VarAssignStmtContext) IDENT() antlr.TerminalNode {
	return s.GetToken(preludioParserIDENT, 0)
}

func (s *VarAssignStmtContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(preludioParserASSIGN, 0)
}

func (s *VarAssignStmtContext) ExprCall() IExprCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprCallContext)
}

func (s *VarAssignStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarAssignStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarAssignStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioParserListener); ok {
		listenerT.EnterVarAssignStmt(s)
	}
}

func (s *VarAssignStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioParserListener); ok {
		listenerT.ExitVarAssignStmt(s)
	}
}

func (p *preludioParser) VarAssignStmt() (localctx IVarAssignStmtContext) {
	this := p
	_ = this

	localctx = NewVarAssignStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, preludioParserRULE_varAssignStmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(166)
		p.Match(preludioParserIDENT)
	}
	{
		p.SetState(167)
		p.Match(preludioParserASSIGN)
	}
	{
		p.SetState(168)
		p.ExprCall()
	}

	return localctx
}

// IVarDeclStmtContext is an interface to support dynamic dispatch.
type IVarDeclStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVarDeclStmtContext differentiates from other interfaces.
	IsVarDeclStmtContext()
}

type VarDeclStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarDeclStmtContext() *VarDeclStmtContext {
	var p = new(VarDeclStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = preludioParserRULE_varDeclStmt
	return p
}

func (*VarDeclStmtContext) IsVarDeclStmtContext() {}

func NewVarDeclStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarDeclStmtContext {
	var p = new(VarDeclStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = preludioParserRULE_varDeclStmt

	return p
}

func (s *VarDeclStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *VarDeclStmtContext) IDENT() antlr.TerminalNode {
	return s.GetToken(preludioParserIDENT, 0)
}

func (s *VarDeclStmtContext) DECLARE() antlr.TerminalNode {
	return s.GetToken(preludioParserDECLARE, 0)
}

func (s *VarDeclStmtContext) ExprCall() IExprCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprCallContext)
}

func (s *VarDeclStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarDeclStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarDeclStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioParserListener); ok {
		listenerT.EnterVarDeclStmt(s)
	}
}

func (s *VarDeclStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioParserListener); ok {
		listenerT.ExitVarDeclStmt(s)
	}
}

func (p *preludioParser) VarDeclStmt() (localctx IVarDeclStmtContext) {
	this := p
	_ = this

	localctx = NewVarDeclStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, preludioParserRULE_varDeclStmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(170)
		p.Match(preludioParserIDENT)
	}
	{
		p.SetState(171)
		p.Match(preludioParserDECLARE)
	}
	{
		p.SetState(172)
		p.ExprCall()
	}

	return localctx
}

// IIfElseStmtContext is an interface to support dynamic dispatch.
type IIfElseStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIfElseStmtContext differentiates from other interfaces.
	IsIfElseStmtContext()
}

type IfElseStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfElseStmtContext() *IfElseStmtContext {
	var p = new(IfElseStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = preludioParserRULE_ifElseStmt
	return p
}

func (*IfElseStmtContext) IsIfElseStmtContext() {}

func NewIfElseStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfElseStmtContext {
	var p = new(IfElseStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = preludioParserRULE_ifElseStmt

	return p
}

func (s *IfElseStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *IfElseStmtContext) IF() antlr.TerminalNode {
	return s.GetToken(preludioParserIF, 0)
}

func (s *IfElseStmtContext) ExprCall() IExprCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprCallContext)
}

func (s *IfElseStmtContext) DO() antlr.TerminalNode {
	return s.GetToken(preludioParserDO, 0)
}

func (s *IfElseStmtContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *IfElseStmtContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *IfElseStmtContext) ELSE() antlr.TerminalNode {
	return s.GetToken(preludioParserELSE, 0)
}

func (s *IfElseStmtContext) END() antlr.TerminalNode {
	return s.GetToken(preludioParserEND, 0)
}

func (s *IfElseStmtContext) AllNl() []INlContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INlContext); ok {
			len++
		}
	}

	tst := make([]INlContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INlContext); ok {
			tst[i] = t.(INlContext)
			i++
		}
	}

	return tst
}

func (s *IfElseStmtContext) Nl(i int) INlContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INlContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(INlContext)
}

func (s *IfElseStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfElseStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfElseStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioParserListener); ok {
		listenerT.EnterIfElseStmt(s)
	}
}

func (s *IfElseStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioParserListener); ok {
		listenerT.ExitIfElseStmt(s)
	}
}

func (p *preludioParser) IfElseStmt() (localctx IIfElseStmtContext) {
	this := p
	_ = this

	localctx = NewIfElseStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, preludioParserRULE_ifElseStmt)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(174)
		p.Match(preludioParserIF)
	}
	{
		p.SetState(175)
		p.ExprCall()
	}
	{
		p.SetState(176)
		p.Match(preludioParserDO)
	}
	p.SetState(222)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(177)
			p.Stmt()
		}

	case 2:
		p.SetState(181)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == preludioParserNEWLINE || _la == preludioParserSINGLE_LINE_COMMENT {
			{
				p.SetState(178)
				p.Nl()
			}

			p.SetState(183)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(193)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-3741193866288561784) != 0 || (int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&509) != 0 {
			{
				p.SetState(184)
				p.Stmt()
			}
			p.SetState(188)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == preludioParserNEWLINE || _la == preludioParserSINGLE_LINE_COMMENT {
				{
					p.SetState(185)
					p.Nl()
				}

				p.SetState(190)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

			p.SetState(195)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(220)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case preludioParserELSE:
			{
				p.SetState(196)
				p.Match(preludioParserELSE)
			}
			p.SetState(217)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext()) {
			case 1:
				{
					p.SetState(197)
					p.Stmt()
				}

			case 2:
				p.SetState(201)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				for _la == preludioParserNEWLINE || _la == preludioParserSINGLE_LINE_COMMENT {
					{
						p.SetState(198)
						p.Nl()
					}

					p.SetState(203)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)
				}
				p.SetState(213)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-3741193866288561784) != 0 || (int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&509) != 0 {
					{
						p.SetState(204)
						p.Stmt()
					}
					p.SetState(208)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)

					for _la == preludioParserNEWLINE || _la == preludioParserSINGLE_LINE_COMMENT {
						{
							p.SetState(205)
							p.Nl()
						}

						p.SetState(210)
						p.GetErrorHandler().Sync(p)
						_la = p.GetTokenStream().LA(1)
					}

					p.SetState(215)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)
				}
				{
					p.SetState(216)
					p.Match(preludioParserEND)
				}

			}

		case preludioParserEND:
			{
				p.SetState(219)
				p.Match(preludioParserEND)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

	}

	return localctx
}

// IForStmtContext is an interface to support dynamic dispatch.
type IForStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsForStmtContext differentiates from other interfaces.
	IsForStmtContext()
}

type ForStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForStmtContext() *ForStmtContext {
	var p = new(ForStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = preludioParserRULE_forStmt
	return p
}

func (*ForStmtContext) IsForStmtContext() {}

func NewForStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForStmtContext {
	var p = new(ForStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = preludioParserRULE_forStmt

	return p
}

func (s *ForStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ForStmtContext) FOR() antlr.TerminalNode {
	return s.GetToken(preludioParserFOR, 0)
}

func (s *ForStmtContext) IDENT() antlr.TerminalNode {
	return s.GetToken(preludioParserIDENT, 0)
}

func (s *ForStmtContext) IN() antlr.TerminalNode {
	return s.GetToken(preludioParserIN, 0)
}

func (s *ForStmtContext) ExprCall() IExprCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprCallContext)
}

func (s *ForStmtContext) DO() antlr.TerminalNode {
	return s.GetToken(preludioParserDO, 0)
}

func (s *ForStmtContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *ForStmtContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *ForStmtContext) END() antlr.TerminalNode {
	return s.GetToken(preludioParserEND, 0)
}

func (s *ForStmtContext) AllNl() []INlContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INlContext); ok {
			len++
		}
	}

	tst := make([]INlContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INlContext); ok {
			tst[i] = t.(INlContext)
			i++
		}
	}

	return tst
}

func (s *ForStmtContext) Nl(i int) INlContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INlContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(INlContext)
}

func (s *ForStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioParserListener); ok {
		listenerT.EnterForStmt(s)
	}
}

func (s *ForStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioParserListener); ok {
		listenerT.ExitForStmt(s)
	}
}

func (p *preludioParser) ForStmt() (localctx IForStmtContext) {
	this := p
	_ = this

	localctx = NewForStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, preludioParserRULE_forStmt)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(224)
		p.Match(preludioParserFOR)
	}
	{
		p.SetState(225)
		p.Match(preludioParserIDENT)
	}
	{
		p.SetState(226)
		p.Match(preludioParserIN)
	}
	{
		p.SetState(227)
		p.ExprCall()
	}
	{
		p.SetState(228)
		p.Match(preludioParserDO)
	}
	p.SetState(249)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(229)
			p.Stmt()
		}

	case 2:
		p.SetState(233)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == preludioParserNEWLINE || _la == preludioParserSINGLE_LINE_COMMENT {
			{
				p.SetState(230)
				p.Nl()
			}

			p.SetState(235)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(245)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-3741193866288561784) != 0 || (int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&509) != 0 {
			{
				p.SetState(236)
				p.Stmt()
			}
			p.SetState(240)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == preludioParserNEWLINE || _la == preludioParserSINGLE_LINE_COMMENT {
				{
					p.SetState(237)
					p.Nl()
				}

				p.SetState(242)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

			p.SetState(247)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(248)
			p.Match(preludioParserEND)
		}

	}

	return localctx
}

// IHelpStmtContext is an interface to support dynamic dispatch.
type IHelpStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHelpStmtContext differentiates from other interfaces.
	IsHelpStmtContext()
}

type HelpStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHelpStmtContext() *HelpStmtContext {
	var p = new(HelpStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = preludioParserRULE_helpStmt
	return p
}

func (*HelpStmtContext) IsHelpStmtContext() {}

func NewHelpStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HelpStmtContext {
	var p = new(HelpStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = preludioParserRULE_helpStmt

	return p
}

func (s *HelpStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *HelpStmtContext) HELP() antlr.TerminalNode {
	return s.GetToken(preludioParserHELP, 0)
}

func (s *HelpStmtContext) IDENT() antlr.TerminalNode {
	return s.GetToken(preludioParserIDENT, 0)
}

func (s *HelpStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HelpStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HelpStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioParserListener); ok {
		listenerT.EnterHelpStmt(s)
	}
}

func (s *HelpStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioParserListener); ok {
		listenerT.ExitHelpStmt(s)
	}
}

func (p *preludioParser) HelpStmt() (localctx IHelpStmtContext) {
	this := p
	_ = this

	localctx = NewHelpStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, preludioParserRULE_helpStmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(251)
		p.Match(preludioParserHELP)
	}
	{
		p.SetState(252)
		p.Match(preludioParserIDENT)
	}

	return localctx
}

// IRetStmtContext is an interface to support dynamic dispatch.
type IRetStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRetStmtContext differentiates from other interfaces.
	IsRetStmtContext()
}

type RetStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRetStmtContext() *RetStmtContext {
	var p = new(RetStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = preludioParserRULE_retStmt
	return p
}

func (*RetStmtContext) IsRetStmtContext() {}

func NewRetStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RetStmtContext {
	var p = new(RetStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = preludioParserRULE_retStmt

	return p
}

func (s *RetStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *RetStmtContext) RET() antlr.TerminalNode {
	return s.GetToken(preludioParserRET, 0)
}

func (s *RetStmtContext) ExprCall() IExprCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprCallContext)
}

func (s *RetStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RetStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RetStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioParserListener); ok {
		listenerT.EnterRetStmt(s)
	}
}

func (s *RetStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioParserListener); ok {
		listenerT.ExitRetStmt(s)
	}
}

func (p *preludioParser) RetStmt() (localctx IRetStmtContext) {
	this := p
	_ = this

	localctx = NewRetStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, preludioParserRULE_retStmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(254)
		p.Match(preludioParserRET)
	}
	{
		p.SetState(255)
		p.ExprCall()
	}

	return localctx
}

// IExprCallContext is an interface to support dynamic dispatch.
type IExprCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprCallContext differentiates from other interfaces.
	IsExprCallContext()
}

type ExprCallContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprCallContext() *ExprCallContext {
	var p = new(ExprCallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = preludioParserRULE_exprCall
	return p
}

func (*ExprCallContext) IsExprCallContext() {}

func NewExprCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprCallContext {
	var p = new(ExprCallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = preludioParserRULE_exprCall

	return p
}

func (s *ExprCallContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprCallContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprCallContext) FuncCall() IFuncCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncCallContext)
}

func (s *ExprCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioParserListener); ok {
		listenerT.EnterExprCall(s)
	}
}

func (s *ExprCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioParserListener); ok {
		listenerT.ExitExprCall(s)
	}
}

func (p *preludioParser) ExprCall() (localctx IExprCallContext) {
	this := p
	_ = this

	localctx = NewExprCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, preludioParserRULE_exprCall)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(259)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(257)
			p.expr(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(258)
			p.FuncCall()
		}

	}

	return localctx
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = preludioParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = preludioParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) Literal() ILiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *ExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprContext) MINUS() antlr.TerminalNode {
	return s.GetToken(preludioParserMINUS, 0)
}

func (s *ExprContext) PLUS() antlr.TerminalNode {
	return s.GetToken(preludioParserPLUS, 0)
}

func (s *ExprContext) NOT() antlr.TerminalNode {
	return s.GetToken(preludioParserNOT, 0)
}

func (s *ExprContext) REVERSE() antlr.TerminalNode {
	return s.GetToken(preludioParserREVERSE, 0)
}

func (s *ExprContext) IDENT() antlr.TerminalNode {
	return s.GetToken(preludioParserIDENT, 0)
}

func (s *ExprContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(preludioParserLPAREN, 0)
}

func (s *ExprContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(preludioParserRPAREN, 0)
}

func (s *ExprContext) List() IListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListContext)
}

func (s *ExprContext) NestedPipeline() INestedPipelineContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INestedPipelineContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INestedPipelineContext)
}

func (s *ExprContext) INDEXING() antlr.TerminalNode {
	return s.GetToken(preludioParserINDEXING, 0)
}

func (s *ExprContext) EXP() antlr.TerminalNode {
	return s.GetToken(preludioParserEXP, 0)
}

func (s *ExprContext) STAR() antlr.TerminalNode {
	return s.GetToken(preludioParserSTAR, 0)
}

func (s *ExprContext) DIV() antlr.TerminalNode {
	return s.GetToken(preludioParserDIV, 0)
}

func (s *ExprContext) MOD() antlr.TerminalNode {
	return s.GetToken(preludioParserMOD, 0)
}

func (s *ExprContext) MODEL() antlr.TerminalNode {
	return s.GetToken(preludioParserMODEL, 0)
}

func (s *ExprContext) EQ() antlr.TerminalNode {
	return s.GetToken(preludioParserEQ, 0)
}

func (s *ExprContext) NE() antlr.TerminalNode {
	return s.GetToken(preludioParserNE, 0)
}

func (s *ExprContext) GE() antlr.TerminalNode {
	return s.GetToken(preludioParserGE, 0)
}

func (s *ExprContext) GT() antlr.TerminalNode {
	return s.GetToken(preludioParserGT, 0)
}

func (s *ExprContext) LE() antlr.TerminalNode {
	return s.GetToken(preludioParserLE, 0)
}

func (s *ExprContext) LT() antlr.TerminalNode {
	return s.GetToken(preludioParserLT, 0)
}

func (s *ExprContext) COALESCE() antlr.TerminalNode {
	return s.GetToken(preludioParserCOALESCE, 0)
}

func (s *ExprContext) AND() antlr.TerminalNode {
	return s.GetToken(preludioParserAND, 0)
}

func (s *ExprContext) OR() antlr.TerminalNode {
	return s.GetToken(preludioParserOR, 0)
}

func (s *ExprContext) IN() antlr.TerminalNode {
	return s.GetToken(preludioParserIN, 0)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioParserListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioParserListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (p *preludioParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *preludioParser) expr(_p int) (localctx IExprContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 34
	p.EnterRecursionRule(localctx, 34, preludioParserRULE_expr, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(273)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(262)
			p.Literal()
		}

	case 2:
		{
			p.SetState(263)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4398046511488) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(264)
			p.expr(13)
		}

	case 3:
		{
			p.SetState(265)
			p.Match(preludioParserREVERSE)
		}
		{
			p.SetState(266)
			p.Match(preludioParserIDENT)
		}

	case 4:
		{
			p.SetState(267)
			p.Match(preludioParserLPAREN)
		}
		{
			p.SetState(268)
			p.expr(0)
		}
		{
			p.SetState(269)
			p.Match(preludioParserRPAREN)
		}

	case 5:
		{
			p.SetState(271)
			p.List()
		}

	case 6:
		{
			p.SetState(272)
			p.NestedPipeline()
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(304)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 34, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(302)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, preludioParserRULE_expr)
				p.SetState(275)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
				}
				{
					p.SetState(276)
					p.Match(preludioParserINDEXING)
				}
				{
					p.SetState(277)
					p.expr(15)
				}

			case 2:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, preludioParserRULE_expr)
				p.SetState(278)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(279)
					p.Match(preludioParserEXP)
				}
				{
					p.SetState(280)
					p.expr(12)
				}

			case 3:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, preludioParserRULE_expr)
				p.SetState(281)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(282)
					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&6656) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(283)
					p.expr(11)
				}

			case 4:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, preludioParserRULE_expr)
				p.SetState(284)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(285)
					_la = p.GetTokenStream().LA(1)

					if !(_la == preludioParserPLUS || _la == preludioParserMINUS) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(286)
					p.expr(10)
				}

			case 5:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, preludioParserRULE_expr)
				p.SetState(287)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(288)
					p.Match(preludioParserMODEL)
				}
				{
					p.SetState(289)
					p.expr(9)
				}

			case 6:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, preludioParserRULE_expr)
				p.SetState(290)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(291)
					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1032192) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(292)
					p.expr(8)
				}

			case 7:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, preludioParserRULE_expr)
				p.SetState(293)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(294)
					p.Match(preludioParserCOALESCE)
				}
				{
					p.SetState(295)
					p.expr(7)
				}

			case 8:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, preludioParserRULE_expr)
				p.SetState(296)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(297)
					_la = p.GetTokenStream().LA(1)

					if !(_la == preludioParserAND || _la == preludioParserOR) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(298)
					p.expr(6)
				}

			case 9:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, preludioParserRULE_expr)
				p.SetState(299)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(300)
					p.Match(preludioParserIN)
				}
				{
					p.SetState(301)
					p.expr(4)
				}

			}

		}
		p.SetState(306)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 34, p.GetParserRuleContext())
	}

	return localctx
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = preludioParserRULE_literal
	return p
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = preludioParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) NA() antlr.TerminalNode {
	return s.GetToken(preludioParserNA, 0)
}

func (s *LiteralContext) BOOLEAN_LIT() antlr.TerminalNode {
	return s.GetToken(preludioParserBOOLEAN_LIT, 0)
}

func (s *LiteralContext) INTEGER_LIT() antlr.TerminalNode {
	return s.GetToken(preludioParserINTEGER_LIT, 0)
}

func (s *LiteralContext) RANGE_LIT() antlr.TerminalNode {
	return s.GetToken(preludioParserRANGE_LIT, 0)
}

func (s *LiteralContext) FLOAT_LIT() antlr.TerminalNode {
	return s.GetToken(preludioParserFLOAT_LIT, 0)
}

func (s *LiteralContext) STRING_LIT() antlr.TerminalNode {
	return s.GetToken(preludioParserSTRING_LIT, 0)
}

func (s *LiteralContext) STRING_INTERP_LIT() antlr.TerminalNode {
	return s.GetToken(preludioParserSTRING_INTERP_LIT, 0)
}

func (s *LiteralContext) STRING_RAW_LIT() antlr.TerminalNode {
	return s.GetToken(preludioParserSTRING_RAW_LIT, 0)
}

func (s *LiteralContext) STRING_PATH_LIT() antlr.TerminalNode {
	return s.GetToken(preludioParserSTRING_PATH_LIT, 0)
}

func (s *LiteralContext) REGEX_LIT() antlr.TerminalNode {
	return s.GetToken(preludioParserREGEX_LIT, 0)
}

func (s *LiteralContext) DATE_LIT() antlr.TerminalNode {
	return s.GetToken(preludioParserDATE_LIT, 0)
}

func (s *LiteralContext) DURATION_LIT() antlr.TerminalNode {
	return s.GetToken(preludioParserDURATION_LIT, 0)
}

func (s *LiteralContext) IDENT() antlr.TerminalNode {
	return s.GetToken(preludioParserIDENT, 0)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioParserListener); ok {
		listenerT.EnterLiteral(s)
	}
}

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioParserListener); ok {
		listenerT.ExitLiteral(s)
	}
}

func (p *preludioParser) Literal() (localctx ILiteralContext) {
	this := p
	_ = this

	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, preludioParserRULE_literal)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(307)
		_la = p.GetTokenStream().LA(1)

		if !((int64((_la-52)) & ^0x3f) == 0 && ((int64(1)<<(_la-52))&2088129) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IListContext is an interface to support dynamic dispatch.
type IListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsListContext differentiates from other interfaces.
	IsListContext()
}

type ListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListContext() *ListContext {
	var p = new(ListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = preludioParserRULE_list
	return p
}

func (*ListContext) IsListContext() {}

func NewListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListContext {
	var p = new(ListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = preludioParserRULE_list

	return p
}

func (s *ListContext) GetParser() antlr.Parser { return s.parser }

func (s *ListContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(preludioParserLBRACKET, 0)
}

func (s *ListContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(preludioParserRBRACKET, 0)
}

func (s *ListContext) AllAssign() []IAssignContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAssignContext); ok {
			len++
		}
	}

	tst := make([]IAssignContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAssignContext); ok {
			tst[i] = t.(IAssignContext)
			i++
		}
	}

	return tst
}

func (s *ListContext) Assign(i int) IAssignContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignContext)
}

func (s *ListContext) AllMultiAssign() []IMultiAssignContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMultiAssignContext); ok {
			len++
		}
	}

	tst := make([]IMultiAssignContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMultiAssignContext); ok {
			tst[i] = t.(IMultiAssignContext)
			i++
		}
	}

	return tst
}

func (s *ListContext) MultiAssign(i int) IMultiAssignContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMultiAssignContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMultiAssignContext)
}

func (s *ListContext) AllExprCall() []IExprCallContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprCallContext); ok {
			len++
		}
	}

	tst := make([]IExprCallContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprCallContext); ok {
			tst[i] = t.(IExprCallContext)
			i++
		}
	}

	return tst
}

func (s *ListContext) ExprCall(i int) IExprCallContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprCallContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprCallContext)
}

func (s *ListContext) AllNl() []INlContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INlContext); ok {
			len++
		}
	}

	tst := make([]INlContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INlContext); ok {
			tst[i] = t.(INlContext)
			i++
		}
	}

	return tst
}

func (s *ListContext) Nl(i int) INlContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INlContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(INlContext)
}

func (s *ListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(preludioParserCOMMA)
}

func (s *ListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(preludioParserCOMMA, i)
}

func (s *ListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioParserListener); ok {
		listenerT.EnterList(s)
	}
}

func (s *ListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioParserListener); ok {
		listenerT.ExitList(s)
	}
}

func (p *preludioParser) List() (localctx IListContext) {
	this := p
	_ = this

	localctx = NewListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, preludioParserRULE_list)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(309)
		p.Match(preludioParserLBRACKET)
	}
	p.SetState(344)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-3526305313756020352) != 0 || (int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&509) != 0 {
		p.SetState(313)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == preludioParserNEWLINE || _la == preludioParserSINGLE_LINE_COMMENT {
			{
				p.SetState(310)
				p.Nl()
			}

			p.SetState(315)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(319)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 36, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(316)
				p.Assign()
			}

		case 2:
			{
				p.SetState(317)
				p.MultiAssign()
			}

		case 3:
			{
				p.SetState(318)
				p.ExprCall()
			}

		}
		p.SetState(335)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(321)
					p.Match(preludioParserCOMMA)
				}
				p.SetState(325)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				for _la == preludioParserNEWLINE || _la == preludioParserSINGLE_LINE_COMMENT {
					{
						p.SetState(322)
						p.Nl()
					}

					p.SetState(327)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)
				}
				p.SetState(331)
				p.GetErrorHandler().Sync(p)
				switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext()) {
				case 1:
					{
						p.SetState(328)
						p.Assign()
					}

				case 2:
					{
						p.SetState(329)
						p.MultiAssign()
					}

				case 3:
					{
						p.SetState(330)
						p.ExprCall()
					}

				}

			}
			p.SetState(337)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext())
		}
		p.SetState(339)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == preludioParserCOMMA {
			{
				p.SetState(338)
				p.Match(preludioParserCOMMA)
			}

		}
		p.SetState(342)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == preludioParserNEWLINE || _la == preludioParserSINGLE_LINE_COMMENT {
			{
				p.SetState(341)
				p.Nl()
			}

		}

	}
	{
		p.SetState(346)
		p.Match(preludioParserRBRACKET)
	}

	return localctx
}

// IFuncCallContext is an interface to support dynamic dispatch.
type IFuncCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncCallContext differentiates from other interfaces.
	IsFuncCallContext()
}

type FuncCallContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncCallContext() *FuncCallContext {
	var p = new(FuncCallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = preludioParserRULE_funcCall
	return p
}

func (*FuncCallContext) IsFuncCallContext() {}

func NewFuncCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncCallContext {
	var p = new(FuncCallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = preludioParserRULE_funcCall

	return p
}

func (s *FuncCallContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncCallContext) IDENT() antlr.TerminalNode {
	return s.GetToken(preludioParserIDENT, 0)
}

func (s *FuncCallContext) FUNCTION_CALL() antlr.TerminalNode {
	return s.GetToken(preludioParserFUNCTION_CALL, 0)
}

func (s *FuncCallContext) AllFuncCallParam() []IFuncCallParamContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFuncCallParamContext); ok {
			len++
		}
	}

	tst := make([]IFuncCallParamContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFuncCallParamContext); ok {
			tst[i] = t.(IFuncCallParamContext)
			i++
		}
	}

	return tst
}

func (s *FuncCallContext) FuncCallParam(i int) IFuncCallParamContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncCallParamContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncCallParamContext)
}

func (s *FuncCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioParserListener); ok {
		listenerT.EnterFuncCall(s)
	}
}

func (s *FuncCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioParserListener); ok {
		listenerT.ExitFuncCall(s)
	}
}

func (p *preludioParser) FuncCall() (localctx IFuncCallContext) {
	this := p
	_ = this

	localctx = NewFuncCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, preludioParserRULE_funcCall)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(348)
		p.Match(preludioParserIDENT)
	}
	{
		p.SetState(349)
		p.Match(preludioParserFUNCTION_CALL)
	}
	p.SetState(353)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 43, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(350)
				p.FuncCallParam()
			}

		}
		p.SetState(355)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 43, p.GetParserRuleContext())
	}

	return localctx
}

// IFuncCallParamContext is an interface to support dynamic dispatch.
type IFuncCallParamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncCallParamContext differentiates from other interfaces.
	IsFuncCallParamContext()
}

type FuncCallParamContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncCallParamContext() *FuncCallParamContext {
	var p = new(FuncCallParamContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = preludioParserRULE_funcCallParam
	return p
}

func (*FuncCallParamContext) IsFuncCallParamContext() {}

func NewFuncCallParamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncCallParamContext {
	var p = new(FuncCallParamContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = preludioParserRULE_funcCallParam

	return p
}

func (s *FuncCallParamContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncCallParamContext) NamedArg() INamedArgContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INamedArgContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INamedArgContext)
}

func (s *FuncCallParamContext) Assign() IAssignContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignContext)
}

func (s *FuncCallParamContext) MultiAssign() IMultiAssignContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMultiAssignContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMultiAssignContext)
}

func (s *FuncCallParamContext) ExprCall() IExprCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprCallContext)
}

func (s *FuncCallParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncCallParamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncCallParamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioParserListener); ok {
		listenerT.EnterFuncCallParam(s)
	}
}

func (s *FuncCallParamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioParserListener); ok {
		listenerT.ExitFuncCallParam(s)
	}
}

func (p *preludioParser) FuncCallParam() (localctx IFuncCallParamContext) {
	this := p
	_ = this

	localctx = NewFuncCallParamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, preludioParserRULE_funcCallParam)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(360)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 44, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(356)
			p.NamedArg()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(357)
			p.Assign()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(358)
			p.MultiAssign()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(359)
			p.ExprCall()
		}

	}

	return localctx
}

// INamedArgContext is an interface to support dynamic dispatch.
type INamedArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNamedArgContext differentiates from other interfaces.
	IsNamedArgContext()
}

type NamedArgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNamedArgContext() *NamedArgContext {
	var p = new(NamedArgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = preludioParserRULE_namedArg
	return p
}

func (*NamedArgContext) IsNamedArgContext() {}

func NewNamedArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NamedArgContext {
	var p = new(NamedArgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = preludioParserRULE_namedArg

	return p
}

func (s *NamedArgContext) GetParser() antlr.Parser { return s.parser }

func (s *NamedArgContext) IDENT() antlr.TerminalNode {
	return s.GetToken(preludioParserIDENT, 0)
}

func (s *NamedArgContext) COLON() antlr.TerminalNode {
	return s.GetToken(preludioParserCOLON, 0)
}

func (s *NamedArgContext) Assign() IAssignContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignContext)
}

func (s *NamedArgContext) ExprCall() IExprCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprCallContext)
}

func (s *NamedArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NamedArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NamedArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioParserListener); ok {
		listenerT.EnterNamedArg(s)
	}
}

func (s *NamedArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioParserListener); ok {
		listenerT.ExitNamedArg(s)
	}
}

func (p *preludioParser) NamedArg() (localctx INamedArgContext) {
	this := p
	_ = this

	localctx = NewNamedArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, preludioParserRULE_namedArg)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(362)
		p.Match(preludioParserIDENT)
	}
	{
		p.SetState(363)
		p.Match(preludioParserCOLON)
	}
	p.SetState(366)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 45, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(364)
			p.Assign()
		}

	case 2:
		{
			p.SetState(365)
			p.ExprCall()
		}

	}

	return localctx
}

// IAssignContext is an interface to support dynamic dispatch.
type IAssignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignContext differentiates from other interfaces.
	IsAssignContext()
}

type AssignContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignContext() *AssignContext {
	var p = new(AssignContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = preludioParserRULE_assign
	return p
}

func (*AssignContext) IsAssignContext() {}

func NewAssignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignContext {
	var p = new(AssignContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = preludioParserRULE_assign

	return p
}

func (s *AssignContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignContext) IDENT() antlr.TerminalNode {
	return s.GetToken(preludioParserIDENT, 0)
}

func (s *AssignContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(preludioParserASSIGN, 0)
}

func (s *AssignContext) ExprCall() IExprCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprCallContext)
}

func (s *AssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioParserListener); ok {
		listenerT.EnterAssign(s)
	}
}

func (s *AssignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioParserListener); ok {
		listenerT.ExitAssign(s)
	}
}

func (p *preludioParser) Assign() (localctx IAssignContext) {
	this := p
	_ = this

	localctx = NewAssignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, preludioParserRULE_assign)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(368)
		p.Match(preludioParserIDENT)
	}
	{
		p.SetState(369)
		p.Match(preludioParserASSIGN)
	}
	{
		p.SetState(370)
		p.ExprCall()
	}

	return localctx
}

// IMultiAssignContext is an interface to support dynamic dispatch.
type IMultiAssignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMultiAssignContext differentiates from other interfaces.
	IsMultiAssignContext()
}

type MultiAssignContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultiAssignContext() *MultiAssignContext {
	var p = new(MultiAssignContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = preludioParserRULE_multiAssign
	return p
}

func (*MultiAssignContext) IsMultiAssignContext() {}

func NewMultiAssignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiAssignContext {
	var p = new(MultiAssignContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = preludioParserRULE_multiAssign

	return p
}

func (s *MultiAssignContext) GetParser() antlr.Parser { return s.parser }

func (s *MultiAssignContext) List() IListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListContext)
}

func (s *MultiAssignContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(preludioParserASSIGN, 0)
}

func (s *MultiAssignContext) ExprCall() IExprCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprCallContext)
}

func (s *MultiAssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiAssignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultiAssignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioParserListener); ok {
		listenerT.EnterMultiAssign(s)
	}
}

func (s *MultiAssignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioParserListener); ok {
		listenerT.ExitMultiAssign(s)
	}
}

func (p *preludioParser) MultiAssign() (localctx IMultiAssignContext) {
	this := p
	_ = this

	localctx = NewMultiAssignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, preludioParserRULE_multiAssign)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(372)
		p.List()
	}
	{
		p.SetState(373)
		p.Match(preludioParserASSIGN)
	}
	{
		p.SetState(374)
		p.ExprCall()
	}

	return localctx
}

// IPipelineContext is an interface to support dynamic dispatch.
type IPipelineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPipelineContext differentiates from other interfaces.
	IsPipelineContext()
}

type PipelineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPipelineContext() *PipelineContext {
	var p = new(PipelineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = preludioParserRULE_pipeline
	return p
}

func (*PipelineContext) IsPipelineContext() {}

func NewPipelineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PipelineContext {
	var p = new(PipelineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = preludioParserRULE_pipeline

	return p
}

func (s *PipelineContext) GetParser() antlr.Parser { return s.parser }

func (s *PipelineContext) ExprCall() IExprCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprCallContext)
}

func (s *PipelineContext) AllNl() []INlContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INlContext); ok {
			len++
		}
	}

	tst := make([]INlContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INlContext); ok {
			tst[i] = t.(INlContext)
			i++
		}
	}

	return tst
}

func (s *PipelineContext) Nl(i int) INlContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INlContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(INlContext)
}

func (s *PipelineContext) AllFuncCall() []IFuncCallContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFuncCallContext); ok {
			len++
		}
	}

	tst := make([]IFuncCallContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFuncCallContext); ok {
			tst[i] = t.(IFuncCallContext)
			i++
		}
	}

	return tst
}

func (s *PipelineContext) FuncCall(i int) IFuncCallContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncCallContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncCallContext)
}

func (s *PipelineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PipelineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PipelineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioParserListener); ok {
		listenerT.EnterPipeline(s)
	}
}

func (s *PipelineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioParserListener); ok {
		listenerT.ExitPipeline(s)
	}
}

func (p *preludioParser) Pipeline() (localctx IPipelineContext) {
	this := p
	_ = this

	localctx = NewPipelineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, preludioParserRULE_pipeline)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(376)
		p.ExprCall()
	}
	p.SetState(382)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 46, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(377)
				p.Nl()
			}
			{
				p.SetState(378)
				p.FuncCall()
			}

		}
		p.SetState(384)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 46, p.GetParserRuleContext())
	}
	{
		p.SetState(385)
		p.Nl()
	}

	return localctx
}

// IInlinePipelineContext is an interface to support dynamic dispatch.
type IInlinePipelineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInlinePipelineContext differentiates from other interfaces.
	IsInlinePipelineContext()
}

type InlinePipelineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInlinePipelineContext() *InlinePipelineContext {
	var p = new(InlinePipelineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = preludioParserRULE_inlinePipeline
	return p
}

func (*InlinePipelineContext) IsInlinePipelineContext() {}

func NewInlinePipelineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InlinePipelineContext {
	var p = new(InlinePipelineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = preludioParserRULE_inlinePipeline

	return p
}

func (s *InlinePipelineContext) GetParser() antlr.Parser { return s.parser }

func (s *InlinePipelineContext) ExprCall() IExprCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprCallContext)
}

func (s *InlinePipelineContext) AllBAR() []antlr.TerminalNode {
	return s.GetTokens(preludioParserBAR)
}

func (s *InlinePipelineContext) BAR(i int) antlr.TerminalNode {
	return s.GetToken(preludioParserBAR, i)
}

func (s *InlinePipelineContext) AllFuncCall() []IFuncCallContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFuncCallContext); ok {
			len++
		}
	}

	tst := make([]IFuncCallContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFuncCallContext); ok {
			tst[i] = t.(IFuncCallContext)
			i++
		}
	}

	return tst
}

func (s *InlinePipelineContext) FuncCall(i int) IFuncCallContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncCallContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncCallContext)
}

func (s *InlinePipelineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InlinePipelineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InlinePipelineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioParserListener); ok {
		listenerT.EnterInlinePipeline(s)
	}
}

func (s *InlinePipelineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioParserListener); ok {
		listenerT.ExitInlinePipeline(s)
	}
}

func (p *preludioParser) InlinePipeline() (localctx IInlinePipelineContext) {
	this := p
	_ = this

	localctx = NewInlinePipelineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, preludioParserRULE_inlinePipeline)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(387)
		p.ExprCall()
	}
	p.SetState(392)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == preludioParserBAR {
		{
			p.SetState(388)
			p.Match(preludioParserBAR)
		}
		{
			p.SetState(389)
			p.FuncCall()
		}

		p.SetState(394)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// INestedPipelineContext is an interface to support dynamic dispatch.
type INestedPipelineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNestedPipelineContext differentiates from other interfaces.
	IsNestedPipelineContext()
}

type NestedPipelineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNestedPipelineContext() *NestedPipelineContext {
	var p = new(NestedPipelineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = preludioParserRULE_nestedPipeline
	return p
}

func (*NestedPipelineContext) IsNestedPipelineContext() {}

func NewNestedPipelineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NestedPipelineContext {
	var p = new(NestedPipelineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = preludioParserRULE_nestedPipeline

	return p
}

func (s *NestedPipelineContext) GetParser() antlr.Parser { return s.parser }

func (s *NestedPipelineContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(preludioParserLPAREN, 0)
}

func (s *NestedPipelineContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(preludioParserRPAREN, 0)
}

func (s *NestedPipelineContext) Pipeline() IPipelineContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPipelineContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPipelineContext)
}

func (s *NestedPipelineContext) InlinePipeline() IInlinePipelineContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInlinePipelineContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInlinePipelineContext)
}

func (s *NestedPipelineContext) AllNl() []INlContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INlContext); ok {
			len++
		}
	}

	tst := make([]INlContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INlContext); ok {
			tst[i] = t.(INlContext)
			i++
		}
	}

	return tst
}

func (s *NestedPipelineContext) Nl(i int) INlContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INlContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(INlContext)
}

func (s *NestedPipelineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NestedPipelineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NestedPipelineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioParserListener); ok {
		listenerT.EnterNestedPipeline(s)
	}
}

func (s *NestedPipelineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioParserListener); ok {
		listenerT.ExitNestedPipeline(s)
	}
}

func (p *preludioParser) NestedPipeline() (localctx INestedPipelineContext) {
	this := p
	_ = this

	localctx = NewNestedPipelineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, preludioParserRULE_nestedPipeline)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(395)
		p.Match(preludioParserLPAREN)
	}
	p.SetState(399)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == preludioParserNEWLINE || _la == preludioParserSINGLE_LINE_COMMENT {
		{
			p.SetState(396)
			p.Nl()
		}

		p.SetState(401)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(404)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 49, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(402)
			p.Pipeline()
		}

	case 2:
		{
			p.SetState(403)
			p.InlinePipeline()
		}

	}
	p.SetState(409)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == preludioParserNEWLINE || _la == preludioParserSINGLE_LINE_COMMENT {
		{
			p.SetState(406)
			p.Nl()
		}

		p.SetState(411)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(412)
		p.Match(preludioParserRPAREN)
	}

	return localctx
}

func (p *preludioParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 17:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *preludioParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
