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
		"'''''", "'and'", "'or'", "'not'", "'??'", "'na'", "'@'", "'!'",
	}
	staticData.symbolicNames = []string{
		"", "FUNC", "PRQL", "RET", "ARROW", "ASSIGN", "DECLARE", "PLUS", "MINUS",
		"STAR", "EXP", "DIV", "MOD", "MODEL", "EQ", "NE", "GE", "GT", "LE",
		"LT", "BAR", "COLON", "COMMA", "DOT", "DOLLAR", "RANGE", "LBRACKET",
		"RBRACKET", "LPAREN", "RPAREN", "LBRACE", "RBRACE", "LANG", "RANG",
		"UNDERSCORE", "BACKTICK", "DOUBLE_QUOTE", "SINGLE_QUOTE", "TRIPLE_DOUBLE_QUOTE",
		"TRIPLE_SINGLE_QUOTE", "AND", "OR", "NOT", "COALESCE", "NA", "INDEXING",
		"FUNCTION_CALL", "WHITESPACE", "NEWLINE", "SINGLE_LINE_COMMENT", "BOOLEAN_LIT",
		"IDENT", "IDENT_START", "IDENT_NEXT", "INTEGER_LIT", "RANGE_LIT", "FLOAT_LIT",
		"STRING_CHAR", "STRING_LIT", "STRING_INTERP_LIT", "STRING_RAW_LIT",
		"STRING_PATH_LIT", "REGEX_LIT", "DATE_LIT", "DURATION_LIT",
	}
	staticData.ruleNames = []string{
		"nl", "program", "programIntro", "funcDef", "funcDefName", "funcDefParams",
		"funcDefParam", "typeDef", "typeTerm", "stmt", "varAssignStmt", "varDeclStmt",
		"retStmt", "exprCall", "expr", "literal", "list", "funcCall", "funcCallParam",
		"namedArg", "assign", "multiAssign", "pipeline", "inlinePipeline", "nestedPipeline",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 64, 322, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 1, 0, 1, 0, 1, 1,
		5, 1, 54, 8, 1, 10, 1, 12, 1, 57, 9, 1, 1, 1, 3, 1, 60, 8, 1, 1, 1, 5,
		1, 63, 8, 1, 10, 1, 12, 1, 66, 9, 1, 1, 1, 1, 1, 3, 1, 70, 8, 1, 1, 1,
		5, 1, 73, 8, 1, 10, 1, 12, 1, 76, 9, 1, 5, 1, 78, 8, 1, 10, 1, 12, 1, 81,
		9, 1, 1, 1, 1, 1, 1, 2, 1, 2, 5, 2, 87, 8, 2, 10, 2, 12, 2, 90, 9, 2, 1,
		2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 101, 8, 3, 10,
		3, 12, 3, 104, 9, 3, 1, 3, 1, 3, 5, 3, 108, 8, 3, 10, 3, 12, 3, 111, 9,
		3, 5, 3, 113, 8, 3, 10, 3, 12, 3, 116, 9, 3, 1, 3, 3, 3, 119, 8, 3, 1,
		4, 1, 4, 3, 4, 123, 8, 4, 1, 5, 5, 5, 126, 8, 5, 10, 5, 12, 5, 129, 9,
		5, 1, 6, 1, 6, 3, 6, 133, 8, 6, 1, 6, 3, 6, 136, 8, 6, 1, 7, 1, 7, 1, 7,
		1, 7, 5, 7, 142, 8, 7, 10, 7, 12, 7, 145, 9, 7, 1, 7, 1, 7, 1, 8, 1, 8,
		3, 8, 151, 8, 8, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 157, 8, 9, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 13, 1,
		13, 3, 13, 172, 8, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14,
		1, 14, 1, 14, 1, 14, 3, 14, 184, 8, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1,
		14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14,
		1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 5, 14, 210,
		8, 14, 10, 14, 12, 14, 213, 9, 14, 1, 15, 1, 15, 1, 16, 1, 16, 5, 16, 219,
		8, 16, 10, 16, 12, 16, 222, 9, 16, 1, 16, 1, 16, 1, 16, 3, 16, 227, 8,
		16, 1, 16, 1, 16, 5, 16, 231, 8, 16, 10, 16, 12, 16, 234, 9, 16, 1, 16,
		1, 16, 1, 16, 3, 16, 239, 8, 16, 5, 16, 241, 8, 16, 10, 16, 12, 16, 244,
		9, 16, 1, 16, 3, 16, 247, 8, 16, 1, 16, 3, 16, 250, 8, 16, 3, 16, 252,
		8, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 5, 17, 259, 8, 17, 10, 17, 12,
		17, 262, 9, 17, 1, 18, 1, 18, 1, 18, 1, 18, 3, 18, 268, 8, 18, 1, 19, 1,
		19, 1, 19, 1, 19, 3, 19, 274, 8, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 21,
		1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 5, 22, 288, 8, 22, 10,
		22, 12, 22, 291, 9, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 5, 23, 298,
		8, 23, 10, 23, 12, 23, 301, 9, 23, 1, 24, 1, 24, 5, 24, 305, 8, 24, 10,
		24, 12, 24, 308, 9, 24, 1, 24, 1, 24, 3, 24, 312, 8, 24, 1, 24, 5, 24,
		315, 8, 24, 10, 24, 12, 24, 318, 9, 24, 1, 24, 1, 24, 1, 24, 0, 1, 28,
		25, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34,
		36, 38, 40, 42, 44, 46, 48, 0, 7, 1, 0, 48, 49, 2, 0, 7, 8, 42, 42, 2,
		0, 9, 9, 11, 12, 1, 0, 7, 8, 1, 0, 14, 19, 1, 0, 40, 41, 4, 0, 44, 44,
		50, 51, 54, 56, 58, 64, 349, 0, 50, 1, 0, 0, 0, 2, 55, 1, 0, 0, 0, 4, 84,
		1, 0, 0, 0, 6, 93, 1, 0, 0, 0, 8, 120, 1, 0, 0, 0, 10, 127, 1, 0, 0, 0,
		12, 132, 1, 0, 0, 0, 14, 137, 1, 0, 0, 0, 16, 148, 1, 0, 0, 0, 18, 156,
		1, 0, 0, 0, 20, 158, 1, 0, 0, 0, 22, 162, 1, 0, 0, 0, 24, 166, 1, 0, 0,
		0, 26, 171, 1, 0, 0, 0, 28, 183, 1, 0, 0, 0, 30, 214, 1, 0, 0, 0, 32, 216,
		1, 0, 0, 0, 34, 255, 1, 0, 0, 0, 36, 267, 1, 0, 0, 0, 38, 269, 1, 0, 0,
		0, 40, 275, 1, 0, 0, 0, 42, 279, 1, 0, 0, 0, 44, 283, 1, 0, 0, 0, 46, 294,
		1, 0, 0, 0, 48, 302, 1, 0, 0, 0, 50, 51, 7, 0, 0, 0, 51, 1, 1, 0, 0, 0,
		52, 54, 3, 0, 0, 0, 53, 52, 1, 0, 0, 0, 54, 57, 1, 0, 0, 0, 55, 53, 1,
		0, 0, 0, 55, 56, 1, 0, 0, 0, 56, 59, 1, 0, 0, 0, 57, 55, 1, 0, 0, 0, 58,
		60, 3, 4, 2, 0, 59, 58, 1, 0, 0, 0, 59, 60, 1, 0, 0, 0, 60, 64, 1, 0, 0,
		0, 61, 63, 3, 0, 0, 0, 62, 61, 1, 0, 0, 0, 63, 66, 1, 0, 0, 0, 64, 62,
		1, 0, 0, 0, 64, 65, 1, 0, 0, 0, 65, 79, 1, 0, 0, 0, 66, 64, 1, 0, 0, 0,
		67, 70, 3, 6, 3, 0, 68, 70, 3, 18, 9, 0, 69, 67, 1, 0, 0, 0, 69, 68, 1,
		0, 0, 0, 70, 74, 1, 0, 0, 0, 71, 73, 3, 0, 0, 0, 72, 71, 1, 0, 0, 0, 73,
		76, 1, 0, 0, 0, 74, 72, 1, 0, 0, 0, 74, 75, 1, 0, 0, 0, 75, 78, 1, 0, 0,
		0, 76, 74, 1, 0, 0, 0, 77, 69, 1, 0, 0, 0, 78, 81, 1, 0, 0, 0, 79, 77,
		1, 0, 0, 0, 79, 80, 1, 0, 0, 0, 80, 82, 1, 0, 0, 0, 81, 79, 1, 0, 0, 0,
		82, 83, 5, 0, 0, 1, 83, 3, 1, 0, 0, 0, 84, 88, 5, 2, 0, 0, 85, 87, 3, 38,
		19, 0, 86, 85, 1, 0, 0, 0, 87, 90, 1, 0, 0, 0, 88, 86, 1, 0, 0, 0, 88,
		89, 1, 0, 0, 0, 89, 91, 1, 0, 0, 0, 90, 88, 1, 0, 0, 0, 91, 92, 3, 0, 0,
		0, 92, 5, 1, 0, 0, 0, 93, 94, 5, 1, 0, 0, 94, 95, 3, 8, 4, 0, 95, 96, 3,
		10, 5, 0, 96, 118, 5, 4, 0, 0, 97, 119, 3, 26, 13, 0, 98, 102, 5, 30, 0,
		0, 99, 101, 3, 0, 0, 0, 100, 99, 1, 0, 0, 0, 101, 104, 1, 0, 0, 0, 102,
		100, 1, 0, 0, 0, 102, 103, 1, 0, 0, 0, 103, 114, 1, 0, 0, 0, 104, 102,
		1, 0, 0, 0, 105, 109, 3, 18, 9, 0, 106, 108, 3, 0, 0, 0, 107, 106, 1, 0,
		0, 0, 108, 111, 1, 0, 0, 0, 109, 107, 1, 0, 0, 0, 109, 110, 1, 0, 0, 0,
		110, 113, 1, 0, 0, 0, 111, 109, 1, 0, 0, 0, 112, 105, 1, 0, 0, 0, 113,
		116, 1, 0, 0, 0, 114, 112, 1, 0, 0, 0, 114, 115, 1, 0, 0, 0, 115, 117,
		1, 0, 0, 0, 116, 114, 1, 0, 0, 0, 117, 119, 5, 31, 0, 0, 118, 97, 1, 0,
		0, 0, 118, 98, 1, 0, 0, 0, 119, 7, 1, 0, 0, 0, 120, 122, 5, 51, 0, 0, 121,
		123, 3, 14, 7, 0, 122, 121, 1, 0, 0, 0, 122, 123, 1, 0, 0, 0, 123, 9, 1,
		0, 0, 0, 124, 126, 3, 12, 6, 0, 125, 124, 1, 0, 0, 0, 126, 129, 1, 0, 0,
		0, 127, 125, 1, 0, 0, 0, 127, 128, 1, 0, 0, 0, 128, 11, 1, 0, 0, 0, 129,
		127, 1, 0, 0, 0, 130, 133, 3, 38, 19, 0, 131, 133, 5, 51, 0, 0, 132, 130,
		1, 0, 0, 0, 132, 131, 1, 0, 0, 0, 133, 135, 1, 0, 0, 0, 134, 136, 3, 14,
		7, 0, 135, 134, 1, 0, 0, 0, 135, 136, 1, 0, 0, 0, 136, 13, 1, 0, 0, 0,
		137, 138, 5, 32, 0, 0, 138, 139, 3, 16, 8, 0, 139, 143, 5, 20, 0, 0, 140,
		142, 3, 16, 8, 0, 141, 140, 1, 0, 0, 0, 142, 145, 1, 0, 0, 0, 143, 141,
		1, 0, 0, 0, 143, 144, 1, 0, 0, 0, 144, 146, 1, 0, 0, 0, 145, 143, 1, 0,
		0, 0, 146, 147, 5, 33, 0, 0, 147, 15, 1, 0, 0, 0, 148, 150, 5, 51, 0, 0,
		149, 151, 3, 14, 7, 0, 150, 149, 1, 0, 0, 0, 150, 151, 1, 0, 0, 0, 151,
		17, 1, 0, 0, 0, 152, 157, 3, 20, 10, 0, 153, 157, 3, 22, 11, 0, 154, 157,
		3, 24, 12, 0, 155, 157, 3, 26, 13, 0, 156, 152, 1, 0, 0, 0, 156, 153, 1,
		0, 0, 0, 156, 154, 1, 0, 0, 0, 156, 155, 1, 0, 0, 0, 157, 19, 1, 0, 0,
		0, 158, 159, 5, 51, 0, 0, 159, 160, 5, 5, 0, 0, 160, 161, 3, 26, 13, 0,
		161, 21, 1, 0, 0, 0, 162, 163, 5, 51, 0, 0, 163, 164, 5, 6, 0, 0, 164,
		165, 3, 26, 13, 0, 165, 23, 1, 0, 0, 0, 166, 167, 5, 3, 0, 0, 167, 168,
		3, 26, 13, 0, 168, 25, 1, 0, 0, 0, 169, 172, 3, 28, 14, 0, 170, 172, 3,
		34, 17, 0, 171, 169, 1, 0, 0, 0, 171, 170, 1, 0, 0, 0, 172, 27, 1, 0, 0,
		0, 173, 174, 6, 14, -1, 0, 174, 184, 3, 30, 15, 0, 175, 176, 7, 1, 0, 0,
		176, 184, 3, 28, 14, 11, 177, 178, 5, 28, 0, 0, 178, 179, 3, 28, 14, 0,
		179, 180, 5, 29, 0, 0, 180, 184, 1, 0, 0, 0, 181, 184, 3, 32, 16, 0, 182,
		184, 3, 48, 24, 0, 183, 173, 1, 0, 0, 0, 183, 175, 1, 0, 0, 0, 183, 177,
		1, 0, 0, 0, 183, 181, 1, 0, 0, 0, 183, 182, 1, 0, 0, 0, 184, 211, 1, 0,
		0, 0, 185, 186, 10, 12, 0, 0, 186, 187, 5, 45, 0, 0, 187, 210, 3, 28, 14,
		13, 188, 189, 10, 10, 0, 0, 189, 190, 5, 10, 0, 0, 190, 210, 3, 28, 14,
		11, 191, 192, 10, 9, 0, 0, 192, 193, 7, 2, 0, 0, 193, 210, 3, 28, 14, 10,
		194, 195, 10, 8, 0, 0, 195, 196, 7, 3, 0, 0, 196, 210, 3, 28, 14, 9, 197,
		198, 10, 7, 0, 0, 198, 199, 5, 13, 0, 0, 199, 210, 3, 28, 14, 8, 200, 201,
		10, 6, 0, 0, 201, 202, 7, 4, 0, 0, 202, 210, 3, 28, 14, 7, 203, 204, 10,
		5, 0, 0, 204, 205, 5, 43, 0, 0, 205, 210, 3, 28, 14, 6, 206, 207, 10, 4,
		0, 0, 207, 208, 7, 5, 0, 0, 208, 210, 3, 28, 14, 5, 209, 185, 1, 0, 0,
		0, 209, 188, 1, 0, 0, 0, 209, 191, 1, 0, 0, 0, 209, 194, 1, 0, 0, 0, 209,
		197, 1, 0, 0, 0, 209, 200, 1, 0, 0, 0, 209, 203, 1, 0, 0, 0, 209, 206,
		1, 0, 0, 0, 210, 213, 1, 0, 0, 0, 211, 209, 1, 0, 0, 0, 211, 212, 1, 0,
		0, 0, 212, 29, 1, 0, 0, 0, 213, 211, 1, 0, 0, 0, 214, 215, 7, 6, 0, 0,
		215, 31, 1, 0, 0, 0, 216, 251, 5, 26, 0, 0, 217, 219, 3, 0, 0, 0, 218,
		217, 1, 0, 0, 0, 219, 222, 1, 0, 0, 0, 220, 218, 1, 0, 0, 0, 220, 221,
		1, 0, 0, 0, 221, 226, 1, 0, 0, 0, 222, 220, 1, 0, 0, 0, 223, 227, 3, 40,
		20, 0, 224, 227, 3, 42, 21, 0, 225, 227, 3, 26, 13, 0, 226, 223, 1, 0,
		0, 0, 226, 224, 1, 0, 0, 0, 226, 225, 1, 0, 0, 0, 227, 242, 1, 0, 0, 0,
		228, 232, 5, 22, 0, 0, 229, 231, 3, 0, 0, 0, 230, 229, 1, 0, 0, 0, 231,
		234, 1, 0, 0, 0, 232, 230, 1, 0, 0, 0, 232, 233, 1, 0, 0, 0, 233, 238,
		1, 0, 0, 0, 234, 232, 1, 0, 0, 0, 235, 239, 3, 40, 20, 0, 236, 239, 3,
		42, 21, 0, 237, 239, 3, 26, 13, 0, 238, 235, 1, 0, 0, 0, 238, 236, 1, 0,
		0, 0, 238, 237, 1, 0, 0, 0, 239, 241, 1, 0, 0, 0, 240, 228, 1, 0, 0, 0,
		241, 244, 1, 0, 0, 0, 242, 240, 1, 0, 0, 0, 242, 243, 1, 0, 0, 0, 243,
		246, 1, 0, 0, 0, 244, 242, 1, 0, 0, 0, 245, 247, 5, 22, 0, 0, 246, 245,
		1, 0, 0, 0, 246, 247, 1, 0, 0, 0, 247, 249, 1, 0, 0, 0, 248, 250, 3, 0,
		0, 0, 249, 248, 1, 0, 0, 0, 249, 250, 1, 0, 0, 0, 250, 252, 1, 0, 0, 0,
		251, 220, 1, 0, 0, 0, 251, 252, 1, 0, 0, 0, 252, 253, 1, 0, 0, 0, 253,
		254, 5, 27, 0, 0, 254, 33, 1, 0, 0, 0, 255, 256, 5, 51, 0, 0, 256, 260,
		5, 46, 0, 0, 257, 259, 3, 36, 18, 0, 258, 257, 1, 0, 0, 0, 259, 262, 1,
		0, 0, 0, 260, 258, 1, 0, 0, 0, 260, 261, 1, 0, 0, 0, 261, 35, 1, 0, 0,
		0, 262, 260, 1, 0, 0, 0, 263, 268, 3, 38, 19, 0, 264, 268, 3, 40, 20, 0,
		265, 268, 3, 42, 21, 0, 266, 268, 3, 26, 13, 0, 267, 263, 1, 0, 0, 0, 267,
		264, 1, 0, 0, 0, 267, 265, 1, 0, 0, 0, 267, 266, 1, 0, 0, 0, 268, 37, 1,
		0, 0, 0, 269, 270, 5, 51, 0, 0, 270, 273, 5, 21, 0, 0, 271, 274, 3, 40,
		20, 0, 272, 274, 3, 26, 13, 0, 273, 271, 1, 0, 0, 0, 273, 272, 1, 0, 0,
		0, 274, 39, 1, 0, 0, 0, 275, 276, 5, 51, 0, 0, 276, 277, 5, 5, 0, 0, 277,
		278, 3, 26, 13, 0, 278, 41, 1, 0, 0, 0, 279, 280, 3, 32, 16, 0, 280, 281,
		5, 5, 0, 0, 281, 282, 3, 26, 13, 0, 282, 43, 1, 0, 0, 0, 283, 289, 3, 26,
		13, 0, 284, 285, 3, 0, 0, 0, 285, 286, 3, 34, 17, 0, 286, 288, 1, 0, 0,
		0, 287, 284, 1, 0, 0, 0, 288, 291, 1, 0, 0, 0, 289, 287, 1, 0, 0, 0, 289,
		290, 1, 0, 0, 0, 290, 292, 1, 0, 0, 0, 291, 289, 1, 0, 0, 0, 292, 293,
		3, 0, 0, 0, 293, 45, 1, 0, 0, 0, 294, 299, 3, 26, 13, 0, 295, 296, 5, 20,
		0, 0, 296, 298, 3, 34, 17, 0, 297, 295, 1, 0, 0, 0, 298, 301, 1, 0, 0,
		0, 299, 297, 1, 0, 0, 0, 299, 300, 1, 0, 0, 0, 300, 47, 1, 0, 0, 0, 301,
		299, 1, 0, 0, 0, 302, 306, 5, 28, 0, 0, 303, 305, 3, 0, 0, 0, 304, 303,
		1, 0, 0, 0, 305, 308, 1, 0, 0, 0, 306, 304, 1, 0, 0, 0, 306, 307, 1, 0,
		0, 0, 307, 311, 1, 0, 0, 0, 308, 306, 1, 0, 0, 0, 309, 312, 3, 44, 22,
		0, 310, 312, 3, 46, 23, 0, 311, 309, 1, 0, 0, 0, 311, 310, 1, 0, 0, 0,
		312, 316, 1, 0, 0, 0, 313, 315, 3, 0, 0, 0, 314, 313, 1, 0, 0, 0, 315,
		318, 1, 0, 0, 0, 316, 314, 1, 0, 0, 0, 316, 317, 1, 0, 0, 0, 317, 319,
		1, 0, 0, 0, 318, 316, 1, 0, 0, 0, 319, 320, 5, 29, 0, 0, 320, 49, 1, 0,
		0, 0, 38, 55, 59, 64, 69, 74, 79, 88, 102, 109, 114, 118, 122, 127, 132,
		135, 143, 150, 156, 171, 183, 209, 211, 220, 226, 232, 238, 242, 246, 249,
		251, 260, 267, 273, 289, 299, 306, 311, 316,
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
	preludioParserCOALESCE            = 43
	preludioParserNA                  = 44
	preludioParserINDEXING            = 45
	preludioParserFUNCTION_CALL       = 46
	preludioParserWHITESPACE          = 47
	preludioParserNEWLINE             = 48
	preludioParserSINGLE_LINE_COMMENT = 49
	preludioParserBOOLEAN_LIT         = 50
	preludioParserIDENT               = 51
	preludioParserIDENT_START         = 52
	preludioParserIDENT_NEXT          = 53
	preludioParserINTEGER_LIT         = 54
	preludioParserRANGE_LIT           = 55
	preludioParserFLOAT_LIT           = 56
	preludioParserSTRING_CHAR         = 57
	preludioParserSTRING_LIT          = 58
	preludioParserSTRING_INTERP_LIT   = 59
	preludioParserSTRING_RAW_LIT      = 60
	preludioParserSTRING_PATH_LIT     = 61
	preludioParserREGEX_LIT           = 62
	preludioParserDATE_LIT            = 63
	preludioParserDURATION_LIT        = 64
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
	preludioParserRULE_retStmt        = 12
	preludioParserRULE_exprCall       = 13
	preludioParserRULE_expr           = 14
	preludioParserRULE_literal        = 15
	preludioParserRULE_list           = 16
	preludioParserRULE_funcCall       = 17
	preludioParserRULE_funcCallParam  = 18
	preludioParserRULE_namedArg       = 19
	preludioParserRULE_assign         = 20
	preludioParserRULE_multiAssign    = 21
	preludioParserRULE_pipeline       = 22
	preludioParserRULE_inlinePipeline = 23
	preludioParserRULE_nestedPipeline = 24
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
		p.SetState(50)
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
	p.SetState(55)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(52)
				p.Nl()
			}

		}
		p.SetState(57)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())
	}
	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == preludioParserPRQL {
		{
			p.SetState(58)
			p.ProgramIntro()
		}

	}
	p.SetState(64)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == preludioParserNEWLINE || _la == preludioParserSINGLE_LINE_COMMENT {
		{
			p.SetState(61)
			p.Nl()
		}

		p.SetState(66)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(79)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64((_la-1)) & ^0x3f) == 0 && ((int64(1)<<(_la-1))&-79364948148354875) != 0 {
		p.SetState(69)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case preludioParserFUNC:
			{
				p.SetState(67)
				p.FuncDef()
			}

		case preludioParserRET, preludioParserPLUS, preludioParserMINUS, preludioParserLBRACKET, preludioParserLPAREN, preludioParserNOT, preludioParserNA, preludioParserBOOLEAN_LIT, preludioParserIDENT, preludioParserINTEGER_LIT, preludioParserRANGE_LIT, preludioParserFLOAT_LIT, preludioParserSTRING_LIT, preludioParserSTRING_INTERP_LIT, preludioParserSTRING_RAW_LIT, preludioParserSTRING_PATH_LIT, preludioParserREGEX_LIT, preludioParserDATE_LIT, preludioParserDURATION_LIT:
			{
				p.SetState(68)
				p.Stmt()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}
		p.SetState(74)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == preludioParserNEWLINE || _la == preludioParserSINGLE_LINE_COMMENT {
			{
				p.SetState(71)
				p.Nl()
			}

			p.SetState(76)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

		p.SetState(81)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(82)
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
		p.SetState(84)
		p.Match(preludioParserPRQL)
	}
	p.SetState(88)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == preludioParserIDENT {
		{
			p.SetState(85)
			p.NamedArg()
		}

		p.SetState(90)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(91)
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

func (s *FuncDefContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(preludioParserLBRACE, 0)
}

func (s *FuncDefContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(preludioParserRBRACE, 0)
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
		p.SetState(93)
		p.Match(preludioParserFUNC)
	}
	{
		p.SetState(94)
		p.FuncDefName()
	}
	{
		p.SetState(95)
		p.FuncDefParams()
	}
	{
		p.SetState(96)
		p.Match(preludioParserARROW)
	}
	p.SetState(118)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case preludioParserPLUS, preludioParserMINUS, preludioParserLBRACKET, preludioParserLPAREN, preludioParserNOT, preludioParserNA, preludioParserBOOLEAN_LIT, preludioParserIDENT, preludioParserINTEGER_LIT, preludioParserRANGE_LIT, preludioParserFLOAT_LIT, preludioParserSTRING_LIT, preludioParserSTRING_INTERP_LIT, preludioParserSTRING_RAW_LIT, preludioParserSTRING_PATH_LIT, preludioParserREGEX_LIT, preludioParserDATE_LIT, preludioParserDURATION_LIT:
		{
			p.SetState(97)
			p.ExprCall()
		}

	case preludioParserLBRACE:
		{
			p.SetState(98)
			p.Match(preludioParserLBRACE)
		}
		p.SetState(102)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == preludioParserNEWLINE || _la == preludioParserSINGLE_LINE_COMMENT {
			{
				p.SetState(99)
				p.Nl()
			}

			p.SetState(104)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(114)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for (int64((_la-3)) & ^0x3f) == 0 && ((int64(1)<<(_la-3))&4591844781390299185) != 0 {
			{
				p.SetState(105)
				p.Stmt()
			}
			p.SetState(109)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == preludioParserNEWLINE || _la == preludioParserSINGLE_LINE_COMMENT {
				{
					p.SetState(106)
					p.Nl()
				}

				p.SetState(111)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

			p.SetState(116)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(117)
			p.Match(preludioParserRBRACE)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
		p.SetState(120)
		p.Match(preludioParserIDENT)
	}
	p.SetState(122)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == preludioParserLANG {
		{
			p.SetState(121)
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
	p.SetState(127)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == preludioParserIDENT {
		{
			p.SetState(124)
			p.FuncDefParam()
		}

		p.SetState(129)
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
	p.SetState(132)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(130)
			p.NamedArg()
		}

	case 2:
		{
			p.SetState(131)
			p.Match(preludioParserIDENT)
		}

	}
	p.SetState(135)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == preludioParserLANG {
		{
			p.SetState(134)
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
		p.SetState(137)
		p.Match(preludioParserLANG)
	}
	{
		p.SetState(138)
		p.TypeTerm()
	}
	{
		p.SetState(139)
		p.Match(preludioParserBAR)
	}
	p.SetState(143)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == preludioParserIDENT {
		{
			p.SetState(140)
			p.TypeTerm()
		}

		p.SetState(145)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(146)
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
		p.SetState(148)
		p.Match(preludioParserIDENT)
	}
	p.SetState(150)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == preludioParserLANG {
		{
			p.SetState(149)
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

	p.SetState(156)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(152)
			p.VarAssignStmt()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(153)
			p.VarDeclStmt()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(154)
			p.RetStmt()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(155)
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
		p.SetState(158)
		p.Match(preludioParserIDENT)
	}
	{
		p.SetState(159)
		p.Match(preludioParserASSIGN)
	}
	{
		p.SetState(160)
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
		p.SetState(162)
		p.Match(preludioParserIDENT)
	}
	{
		p.SetState(163)
		p.Match(preludioParserDECLARE)
	}
	{
		p.SetState(164)
		p.ExprCall()
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
	p.EnterRule(localctx, 24, preludioParserRULE_retStmt)

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
		p.Match(preludioParserRET)
	}
	{
		p.SetState(167)
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
	p.EnterRule(localctx, 26, preludioParserRULE_exprCall)

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

	p.SetState(171)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(169)
			p.expr(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(170)
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
	_startState := 28
	p.EnterRecursionRule(localctx, 28, preludioParserRULE_expr, _p)
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
	p.SetState(183)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(174)
			p.Literal()
		}

	case 2:
		{
			p.SetState(175)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4398046511488) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(176)
			p.expr(11)
		}

	case 3:
		{
			p.SetState(177)
			p.Match(preludioParserLPAREN)
		}
		{
			p.SetState(178)
			p.expr(0)
		}
		{
			p.SetState(179)
			p.Match(preludioParserRPAREN)
		}

	case 4:
		{
			p.SetState(181)
			p.List()
		}

	case 5:
		{
			p.SetState(182)
			p.NestedPipeline()
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(211)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(209)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, preludioParserRULE_expr)
				p.SetState(185)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				{
					p.SetState(186)
					p.Match(preludioParserINDEXING)
				}
				{
					p.SetState(187)
					p.expr(13)
				}

			case 2:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, preludioParserRULE_expr)
				p.SetState(188)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(189)
					p.Match(preludioParserEXP)
				}
				{
					p.SetState(190)
					p.expr(11)
				}

			case 3:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, preludioParserRULE_expr)
				p.SetState(191)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(192)
					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&6656) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(193)
					p.expr(10)
				}

			case 4:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, preludioParserRULE_expr)
				p.SetState(194)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(195)
					_la = p.GetTokenStream().LA(1)

					if !(_la == preludioParserPLUS || _la == preludioParserMINUS) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(196)
					p.expr(9)
				}

			case 5:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, preludioParserRULE_expr)
				p.SetState(197)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(198)
					p.Match(preludioParserMODEL)
				}
				{
					p.SetState(199)
					p.expr(8)
				}

			case 6:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, preludioParserRULE_expr)
				p.SetState(200)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(201)
					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1032192) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(202)
					p.expr(7)
				}

			case 7:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, preludioParserRULE_expr)
				p.SetState(203)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(204)
					p.Match(preludioParserCOALESCE)
				}
				{
					p.SetState(205)
					p.expr(6)
				}

			case 8:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, preludioParserRULE_expr)
				p.SetState(206)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(207)
					_la = p.GetTokenStream().LA(1)

					if !(_la == preludioParserAND || _la == preludioParserOR) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(208)
					p.expr(5)
				}

			}

		}
		p.SetState(213)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 30, preludioParserRULE_literal)
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
		p.SetState(214)
		_la = p.GetTokenStream().LA(1)

		if !((int64((_la-44)) & ^0x3f) == 0 && ((int64(1)<<(_la-44))&2088129) != 0) {
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
	p.EnterRule(localctx, 32, preludioParserRULE_list)
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
		p.SetState(216)
		p.Match(preludioParserLBRACKET)
	}
	p.SetState(251)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64((_la-7)) & ^0x3f) == 0 && ((int64(1)<<(_la-7))&286996895906660355) != 0 {
		p.SetState(220)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == preludioParserNEWLINE || _la == preludioParserSINGLE_LINE_COMMENT {
			{
				p.SetState(217)
				p.Nl()
			}

			p.SetState(222)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(226)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(223)
				p.Assign()
			}

		case 2:
			{
				p.SetState(224)
				p.MultiAssign()
			}

		case 3:
			{
				p.SetState(225)
				p.ExprCall()
			}

		}
		p.SetState(242)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(228)
					p.Match(preludioParserCOMMA)
				}
				p.SetState(232)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				for _la == preludioParserNEWLINE || _la == preludioParserSINGLE_LINE_COMMENT {
					{
						p.SetState(229)
						p.Nl()
					}

					p.SetState(234)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)
				}
				p.SetState(238)
				p.GetErrorHandler().Sync(p)
				switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext()) {
				case 1:
					{
						p.SetState(235)
						p.Assign()
					}

				case 2:
					{
						p.SetState(236)
						p.MultiAssign()
					}

				case 3:
					{
						p.SetState(237)
						p.ExprCall()
					}

				}

			}
			p.SetState(244)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext())
		}
		p.SetState(246)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == preludioParserCOMMA {
			{
				p.SetState(245)
				p.Match(preludioParserCOMMA)
			}

		}
		p.SetState(249)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == preludioParserNEWLINE || _la == preludioParserSINGLE_LINE_COMMENT {
			{
				p.SetState(248)
				p.Nl()
			}

		}

	}
	{
		p.SetState(253)
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
	p.EnterRule(localctx, 34, preludioParserRULE_funcCall)

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
		p.SetState(255)
		p.Match(preludioParserIDENT)
	}
	{
		p.SetState(256)
		p.Match(preludioParserFUNCTION_CALL)
	}
	p.SetState(260)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(257)
				p.FuncCallParam()
			}

		}
		p.SetState(262)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 36, preludioParserRULE_funcCallParam)

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

	p.SetState(267)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(263)
			p.NamedArg()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(264)
			p.Assign()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(265)
			p.MultiAssign()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(266)
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
	p.EnterRule(localctx, 38, preludioParserRULE_namedArg)

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
		p.SetState(269)
		p.Match(preludioParserIDENT)
	}
	{
		p.SetState(270)
		p.Match(preludioParserCOLON)
	}
	p.SetState(273)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(271)
			p.Assign()
		}

	case 2:
		{
			p.SetState(272)
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
	p.EnterRule(localctx, 40, preludioParserRULE_assign)

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
		p.SetState(275)
		p.Match(preludioParserIDENT)
	}
	{
		p.SetState(276)
		p.Match(preludioParserASSIGN)
	}
	{
		p.SetState(277)
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
	p.EnterRule(localctx, 42, preludioParserRULE_multiAssign)

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
		p.SetState(279)
		p.List()
	}
	{
		p.SetState(280)
		p.Match(preludioParserASSIGN)
	}
	{
		p.SetState(281)
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
	p.EnterRule(localctx, 44, preludioParserRULE_pipeline)

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
		p.SetState(283)
		p.ExprCall()
	}
	p.SetState(289)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(284)
				p.Nl()
			}
			{
				p.SetState(285)
				p.FuncCall()
			}

		}
		p.SetState(291)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext())
	}
	{
		p.SetState(292)
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
	p.EnterRule(localctx, 46, preludioParserRULE_inlinePipeline)
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
		p.SetState(294)
		p.ExprCall()
	}
	p.SetState(299)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == preludioParserBAR {
		{
			p.SetState(295)
			p.Match(preludioParserBAR)
		}
		{
			p.SetState(296)
			p.FuncCall()
		}

		p.SetState(301)
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
	p.EnterRule(localctx, 48, preludioParserRULE_nestedPipeline)
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
		p.SetState(302)
		p.Match(preludioParserLPAREN)
	}
	p.SetState(306)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == preludioParserNEWLINE || _la == preludioParserSINGLE_LINE_COMMENT {
		{
			p.SetState(303)
			p.Nl()
		}

		p.SetState(308)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(311)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 36, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(309)
			p.Pipeline()
		}

	case 2:
		{
			p.SetState(310)
			p.InlinePipeline()
		}

	}
	p.SetState(316)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == preludioParserNEWLINE || _la == preludioParserSINGLE_LINE_COMMENT {
		{
			p.SetState(313)
			p.Nl()
		}

		p.SetState(318)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(319)
		p.Match(preludioParserRPAREN)
	}

	return localctx
}

func (p *preludioParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 14:
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
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
