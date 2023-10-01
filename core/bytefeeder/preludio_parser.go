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
		"'*'", "'^'", "'/'", "'%'", "'~'", "'=='", "'!='", "'<='", "'>='", "'@'",
		"'|'", "':'", "','", "'.'", "'$'", "'..'", "'<'", "'>'", "'['", "']'",
		"'('", "')'", "'{'", "'}'", "'_'", "'`'", "'\"'", "'''", "'\"\"\"'",
		"'''''", "'and'", "'or'", "'not'", "'??'", "'na'",
	}
	staticData.symbolicNames = []string{
		"", "FUNC", "PRQL", "RET", "ARROW", "ASSIGN", "DECLARE", "PLUS", "MINUS",
		"STAR", "EXP", "DIV", "MOD", "MODEL", "EQ", "NE", "LE", "GE", "AT",
		"BAR", "COLON", "COMMA", "DOT", "DOLLAR", "RANGE", "LANG", "RANG", "LBRACKET",
		"RBRACKET", "LPAREN", "RPAREN", "LBRACE", "RBRACE", "UNDERSCORE", "BACKTICK",
		"DOUBLE_QUOTE", "SINGLE_QUOTE", "TRIPLE_DOUBLE_QUOTE", "TRIPLE_SINGLE_QUOTE",
		"AND", "OR", "NOT", "COALESCE", "NULL_", "IDENT", "IDENT_START", "IDENT_NEXT",
		"WHITESPACE", "NEWLINE", "SINGLE_LINE_COMMENT", "BOOL_LIT", "INT_LIT",
		"RNG_LIT", "FLT_LIT", "STR_CHAR", "STR_RAW_CHAR", "STR_LIT", "STR_INTERP",
		"STR_RAW", "STR_PATH", "RXP_LIT", "DAT_LIT", "DUR_LIT",
	}
	staticData.ruleNames = []string{
		"nl", "program", "programIntro", "funcDef", "funcDefName", "funcDefParams",
		"funcDefParam", "typeDef", "typeTerm", "stmt", "varAssignStmt", "varDeclStmt",
		"retStmt", "pipeline", "inlinePipeline", "nestedPipeline", "identBacktick",
		"funcCall", "funcCallParam", "namedArg", "assign", "multiAssign", "exprCall",
		"expr", "term", "exprUnary", "literal", "list",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 62, 325, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 1, 0, 1, 0, 1, 1, 5, 1, 60, 8, 1, 10, 1, 12, 1, 63,
		9, 1, 1, 1, 3, 1, 66, 8, 1, 1, 1, 5, 1, 69, 8, 1, 10, 1, 12, 1, 72, 9,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 78, 8, 1, 1, 1, 5, 1, 81, 8, 1, 10, 1,
		12, 1, 84, 9, 1, 5, 1, 86, 8, 1, 10, 1, 12, 1, 89, 9, 1, 1, 1, 1, 1, 1,
		2, 1, 2, 5, 2, 95, 8, 2, 10, 2, 12, 2, 98, 9, 2, 1, 2, 1, 2, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 3, 4, 110, 8, 4, 1, 5, 5, 5, 113, 8,
		5, 10, 5, 12, 5, 116, 9, 5, 1, 6, 1, 6, 3, 6, 120, 8, 6, 1, 6, 3, 6, 123,
		8, 6, 1, 7, 1, 7, 1, 7, 1, 7, 5, 7, 129, 8, 7, 10, 7, 12, 7, 132, 9, 7,
		1, 7, 1, 7, 1, 8, 1, 8, 3, 8, 138, 8, 8, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9,
		144, 8, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12,
		1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 5, 13, 161, 8, 13, 10, 13, 12,
		13, 164, 9, 13, 1, 13, 1, 13, 3, 13, 168, 8, 13, 1, 14, 1, 14, 1, 14, 5,
		14, 173, 8, 14, 10, 14, 12, 14, 176, 9, 14, 1, 15, 1, 15, 5, 15, 180, 8,
		15, 10, 15, 12, 15, 183, 9, 15, 1, 15, 1, 15, 3, 15, 187, 8, 15, 1, 15,
		5, 15, 190, 8, 15, 10, 15, 12, 15, 193, 9, 15, 1, 15, 1, 15, 1, 16, 1,
		16, 5, 16, 199, 8, 16, 10, 16, 12, 16, 202, 9, 16, 1, 16, 1, 16, 1, 17,
		1, 17, 5, 17, 208, 8, 17, 10, 17, 12, 17, 211, 9, 17, 1, 18, 1, 18, 1,
		18, 1, 18, 3, 18, 217, 8, 18, 1, 19, 1, 19, 1, 19, 1, 19, 3, 19, 223, 8,
		19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22,
		3, 22, 235, 8, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 3, 23, 243,
		8, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1,
		23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23,
		1, 23, 1, 23, 1, 23, 1, 23, 5, 23, 269, 8, 23, 10, 23, 12, 23, 272, 9,
		23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 3, 24, 279, 8, 24, 1, 25, 1, 25,
		1, 25, 1, 26, 1, 26, 1, 27, 1, 27, 5, 27, 288, 8, 27, 10, 27, 12, 27, 291,
		9, 27, 1, 27, 1, 27, 1, 27, 3, 27, 296, 8, 27, 1, 27, 1, 27, 5, 27, 300,
		8, 27, 10, 27, 12, 27, 303, 9, 27, 1, 27, 1, 27, 1, 27, 3, 27, 308, 8,
		27, 5, 27, 310, 8, 27, 10, 27, 12, 27, 313, 9, 27, 1, 27, 3, 27, 316, 8,
		27, 1, 27, 3, 27, 319, 8, 27, 3, 27, 321, 8, 27, 1, 27, 1, 27, 1, 27, 0,
		1, 46, 28, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32,
		34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 0, 7, 2, 0, 34, 34, 48, 48,
		2, 0, 9, 9, 11, 12, 1, 0, 7, 8, 2, 0, 14, 17, 25, 26, 1, 0, 39, 40, 2,
		0, 7, 8, 41, 41, 3, 0, 43, 44, 50, 53, 56, 62, 350, 0, 56, 1, 0, 0, 0,
		2, 61, 1, 0, 0, 0, 4, 92, 1, 0, 0, 0, 6, 101, 1, 0, 0, 0, 8, 107, 1, 0,
		0, 0, 10, 114, 1, 0, 0, 0, 12, 119, 1, 0, 0, 0, 14, 124, 1, 0, 0, 0, 16,
		135, 1, 0, 0, 0, 18, 143, 1, 0, 0, 0, 20, 145, 1, 0, 0, 0, 22, 149, 1,
		0, 0, 0, 24, 153, 1, 0, 0, 0, 26, 156, 1, 0, 0, 0, 28, 169, 1, 0, 0, 0,
		30, 177, 1, 0, 0, 0, 32, 196, 1, 0, 0, 0, 34, 205, 1, 0, 0, 0, 36, 216,
		1, 0, 0, 0, 38, 218, 1, 0, 0, 0, 40, 224, 1, 0, 0, 0, 42, 228, 1, 0, 0,
		0, 44, 234, 1, 0, 0, 0, 46, 242, 1, 0, 0, 0, 48, 278, 1, 0, 0, 0, 50, 280,
		1, 0, 0, 0, 52, 283, 1, 0, 0, 0, 54, 285, 1, 0, 0, 0, 56, 57, 5, 48, 0,
		0, 57, 1, 1, 0, 0, 0, 58, 60, 3, 0, 0, 0, 59, 58, 1, 0, 0, 0, 60, 63, 1,
		0, 0, 0, 61, 59, 1, 0, 0, 0, 61, 62, 1, 0, 0, 0, 62, 65, 1, 0, 0, 0, 63,
		61, 1, 0, 0, 0, 64, 66, 3, 4, 2, 0, 65, 64, 1, 0, 0, 0, 65, 66, 1, 0, 0,
		0, 66, 70, 1, 0, 0, 0, 67, 69, 3, 0, 0, 0, 68, 67, 1, 0, 0, 0, 69, 72,
		1, 0, 0, 0, 70, 68, 1, 0, 0, 0, 70, 71, 1, 0, 0, 0, 71, 87, 1, 0, 0, 0,
		72, 70, 1, 0, 0, 0, 73, 78, 3, 6, 3, 0, 74, 78, 3, 26, 13, 0, 75, 78, 3,
		28, 14, 0, 76, 78, 3, 18, 9, 0, 77, 73, 1, 0, 0, 0, 77, 74, 1, 0, 0, 0,
		77, 75, 1, 0, 0, 0, 77, 76, 1, 0, 0, 0, 78, 82, 1, 0, 0, 0, 79, 81, 3,
		0, 0, 0, 80, 79, 1, 0, 0, 0, 81, 84, 1, 0, 0, 0, 82, 80, 1, 0, 0, 0, 82,
		83, 1, 0, 0, 0, 83, 86, 1, 0, 0, 0, 84, 82, 1, 0, 0, 0, 85, 77, 1, 0, 0,
		0, 86, 89, 1, 0, 0, 0, 87, 85, 1, 0, 0, 0, 87, 88, 1, 0, 0, 0, 88, 90,
		1, 0, 0, 0, 89, 87, 1, 0, 0, 0, 90, 91, 5, 0, 0, 1, 91, 3, 1, 0, 0, 0,
		92, 96, 5, 2, 0, 0, 93, 95, 3, 38, 19, 0, 94, 93, 1, 0, 0, 0, 95, 98, 1,
		0, 0, 0, 96, 94, 1, 0, 0, 0, 96, 97, 1, 0, 0, 0, 97, 99, 1, 0, 0, 0, 98,
		96, 1, 0, 0, 0, 99, 100, 3, 0, 0, 0, 100, 5, 1, 0, 0, 0, 101, 102, 5, 1,
		0, 0, 102, 103, 3, 8, 4, 0, 103, 104, 3, 10, 5, 0, 104, 105, 5, 4, 0, 0,
		105, 106, 3, 46, 23, 0, 106, 7, 1, 0, 0, 0, 107, 109, 5, 44, 0, 0, 108,
		110, 3, 14, 7, 0, 109, 108, 1, 0, 0, 0, 109, 110, 1, 0, 0, 0, 110, 9, 1,
		0, 0, 0, 111, 113, 3, 12, 6, 0, 112, 111, 1, 0, 0, 0, 113, 116, 1, 0, 0,
		0, 114, 112, 1, 0, 0, 0, 114, 115, 1, 0, 0, 0, 115, 11, 1, 0, 0, 0, 116,
		114, 1, 0, 0, 0, 117, 120, 3, 38, 19, 0, 118, 120, 5, 44, 0, 0, 119, 117,
		1, 0, 0, 0, 119, 118, 1, 0, 0, 0, 120, 122, 1, 0, 0, 0, 121, 123, 3, 14,
		7, 0, 122, 121, 1, 0, 0, 0, 122, 123, 1, 0, 0, 0, 123, 13, 1, 0, 0, 0,
		124, 125, 5, 25, 0, 0, 125, 126, 3, 16, 8, 0, 126, 130, 5, 19, 0, 0, 127,
		129, 3, 16, 8, 0, 128, 127, 1, 0, 0, 0, 129, 132, 1, 0, 0, 0, 130, 128,
		1, 0, 0, 0, 130, 131, 1, 0, 0, 0, 131, 133, 1, 0, 0, 0, 132, 130, 1, 0,
		0, 0, 133, 134, 5, 26, 0, 0, 134, 15, 1, 0, 0, 0, 135, 137, 5, 44, 0, 0,
		136, 138, 3, 14, 7, 0, 137, 136, 1, 0, 0, 0, 137, 138, 1, 0, 0, 0, 138,
		17, 1, 0, 0, 0, 139, 144, 3, 20, 10, 0, 140, 144, 3, 22, 11, 0, 141, 144,
		3, 24, 12, 0, 142, 144, 3, 46, 23, 0, 143, 139, 1, 0, 0, 0, 143, 140, 1,
		0, 0, 0, 143, 141, 1, 0, 0, 0, 143, 142, 1, 0, 0, 0, 144, 19, 1, 0, 0,
		0, 145, 146, 5, 44, 0, 0, 146, 147, 5, 5, 0, 0, 147, 148, 3, 44, 22, 0,
		148, 21, 1, 0, 0, 0, 149, 150, 5, 44, 0, 0, 150, 151, 5, 6, 0, 0, 151,
		152, 3, 44, 22, 0, 152, 23, 1, 0, 0, 0, 153, 154, 5, 3, 0, 0, 154, 155,
		3, 44, 22, 0, 155, 25, 1, 0, 0, 0, 156, 162, 3, 44, 22, 0, 157, 158, 3,
		0, 0, 0, 158, 159, 3, 34, 17, 0, 159, 161, 1, 0, 0, 0, 160, 157, 1, 0,
		0, 0, 161, 164, 1, 0, 0, 0, 162, 160, 1, 0, 0, 0, 162, 163, 1, 0, 0, 0,
		163, 167, 1, 0, 0, 0, 164, 162, 1, 0, 0, 0, 165, 168, 3, 0, 0, 0, 166,
		168, 5, 0, 0, 1, 167, 165, 1, 0, 0, 0, 167, 166, 1, 0, 0, 0, 168, 27, 1,
		0, 0, 0, 169, 174, 3, 44, 22, 0, 170, 171, 5, 19, 0, 0, 171, 173, 3, 34,
		17, 0, 172, 170, 1, 0, 0, 0, 173, 176, 1, 0, 0, 0, 174, 172, 1, 0, 0, 0,
		174, 175, 1, 0, 0, 0, 175, 29, 1, 0, 0, 0, 176, 174, 1, 0, 0, 0, 177, 181,
		5, 29, 0, 0, 178, 180, 3, 0, 0, 0, 179, 178, 1, 0, 0, 0, 180, 183, 1, 0,
		0, 0, 181, 179, 1, 0, 0, 0, 181, 182, 1, 0, 0, 0, 182, 186, 1, 0, 0, 0,
		183, 181, 1, 0, 0, 0, 184, 187, 3, 26, 13, 0, 185, 187, 3, 28, 14, 0, 186,
		184, 1, 0, 0, 0, 186, 185, 1, 0, 0, 0, 187, 191, 1, 0, 0, 0, 188, 190,
		3, 0, 0, 0, 189, 188, 1, 0, 0, 0, 190, 193, 1, 0, 0, 0, 191, 189, 1, 0,
		0, 0, 191, 192, 1, 0, 0, 0, 192, 194, 1, 0, 0, 0, 193, 191, 1, 0, 0, 0,
		194, 195, 5, 30, 0, 0, 195, 31, 1, 0, 0, 0, 196, 200, 5, 34, 0, 0, 197,
		199, 8, 0, 0, 0, 198, 197, 1, 0, 0, 0, 199, 202, 1, 0, 0, 0, 200, 198,
		1, 0, 0, 0, 200, 201, 1, 0, 0, 0, 201, 203, 1, 0, 0, 0, 202, 200, 1, 0,
		0, 0, 203, 204, 5, 34, 0, 0, 204, 33, 1, 0, 0, 0, 205, 209, 5, 44, 0, 0,
		206, 208, 3, 36, 18, 0, 207, 206, 1, 0, 0, 0, 208, 211, 1, 0, 0, 0, 209,
		207, 1, 0, 0, 0, 209, 210, 1, 0, 0, 0, 210, 35, 1, 0, 0, 0, 211, 209, 1,
		0, 0, 0, 212, 217, 3, 38, 19, 0, 213, 217, 3, 40, 20, 0, 214, 217, 3, 42,
		21, 0, 215, 217, 3, 46, 23, 0, 216, 212, 1, 0, 0, 0, 216, 213, 1, 0, 0,
		0, 216, 214, 1, 0, 0, 0, 216, 215, 1, 0, 0, 0, 217, 37, 1, 0, 0, 0, 218,
		219, 5, 44, 0, 0, 219, 222, 5, 20, 0, 0, 220, 223, 3, 40, 20, 0, 221, 223,
		3, 46, 23, 0, 222, 220, 1, 0, 0, 0, 222, 221, 1, 0, 0, 0, 223, 39, 1, 0,
		0, 0, 224, 225, 5, 44, 0, 0, 225, 226, 5, 5, 0, 0, 226, 227, 3, 44, 22,
		0, 227, 41, 1, 0, 0, 0, 228, 229, 3, 54, 27, 0, 229, 230, 5, 5, 0, 0, 230,
		231, 3, 44, 22, 0, 231, 43, 1, 0, 0, 0, 232, 235, 3, 34, 17, 0, 233, 235,
		3, 46, 23, 0, 234, 232, 1, 0, 0, 0, 234, 233, 1, 0, 0, 0, 235, 45, 1, 0,
		0, 0, 236, 237, 6, 23, -1, 0, 237, 238, 5, 29, 0, 0, 238, 239, 3, 46, 23,
		0, 239, 240, 5, 30, 0, 0, 240, 243, 1, 0, 0, 0, 241, 243, 3, 48, 24, 0,
		242, 236, 1, 0, 0, 0, 242, 241, 1, 0, 0, 0, 243, 270, 1, 0, 0, 0, 244,
		245, 10, 10, 0, 0, 245, 246, 5, 23, 0, 0, 246, 269, 3, 46, 23, 11, 247,
		248, 10, 9, 0, 0, 248, 249, 5, 10, 0, 0, 249, 269, 3, 46, 23, 10, 250,
		251, 10, 8, 0, 0, 251, 252, 7, 1, 0, 0, 252, 269, 3, 46, 23, 9, 253, 254,
		10, 7, 0, 0, 254, 255, 7, 2, 0, 0, 255, 269, 3, 46, 23, 8, 256, 257, 10,
		6, 0, 0, 257, 258, 5, 13, 0, 0, 258, 269, 3, 46, 23, 7, 259, 260, 10, 5,
		0, 0, 260, 261, 7, 3, 0, 0, 261, 269, 3, 46, 23, 6, 262, 263, 10, 4, 0,
		0, 263, 264, 5, 42, 0, 0, 264, 269, 3, 46, 23, 5, 265, 266, 10, 3, 0, 0,
		266, 267, 7, 4, 0, 0, 267, 269, 3, 46, 23, 4, 268, 244, 1, 0, 0, 0, 268,
		247, 1, 0, 0, 0, 268, 250, 1, 0, 0, 0, 268, 253, 1, 0, 0, 0, 268, 256,
		1, 0, 0, 0, 268, 259, 1, 0, 0, 0, 268, 262, 1, 0, 0, 0, 268, 265, 1, 0,
		0, 0, 269, 272, 1, 0, 0, 0, 270, 268, 1, 0, 0, 0, 270, 271, 1, 0, 0, 0,
		271, 47, 1, 0, 0, 0, 272, 270, 1, 0, 0, 0, 273, 279, 3, 52, 26, 0, 274,
		279, 3, 32, 16, 0, 275, 279, 3, 50, 25, 0, 276, 279, 3, 54, 27, 0, 277,
		279, 3, 30, 15, 0, 278, 273, 1, 0, 0, 0, 278, 274, 1, 0, 0, 0, 278, 275,
		1, 0, 0, 0, 278, 276, 1, 0, 0, 0, 278, 277, 1, 0, 0, 0, 279, 49, 1, 0,
		0, 0, 280, 281, 7, 5, 0, 0, 281, 282, 3, 48, 24, 0, 282, 51, 1, 0, 0, 0,
		283, 284, 7, 6, 0, 0, 284, 53, 1, 0, 0, 0, 285, 320, 5, 27, 0, 0, 286,
		288, 3, 0, 0, 0, 287, 286, 1, 0, 0, 0, 288, 291, 1, 0, 0, 0, 289, 287,
		1, 0, 0, 0, 289, 290, 1, 0, 0, 0, 290, 295, 1, 0, 0, 0, 291, 289, 1, 0,
		0, 0, 292, 296, 3, 40, 20, 0, 293, 296, 3, 42, 21, 0, 294, 296, 3, 44,
		22, 0, 295, 292, 1, 0, 0, 0, 295, 293, 1, 0, 0, 0, 295, 294, 1, 0, 0, 0,
		296, 311, 1, 0, 0, 0, 297, 301, 5, 21, 0, 0, 298, 300, 3, 0, 0, 0, 299,
		298, 1, 0, 0, 0, 300, 303, 1, 0, 0, 0, 301, 299, 1, 0, 0, 0, 301, 302,
		1, 0, 0, 0, 302, 307, 1, 0, 0, 0, 303, 301, 1, 0, 0, 0, 304, 308, 3, 40,
		20, 0, 305, 308, 3, 42, 21, 0, 306, 308, 3, 44, 22, 0, 307, 304, 1, 0,
		0, 0, 307, 305, 1, 0, 0, 0, 307, 306, 1, 0, 0, 0, 308, 310, 1, 0, 0, 0,
		309, 297, 1, 0, 0, 0, 310, 313, 1, 0, 0, 0, 311, 309, 1, 0, 0, 0, 311,
		312, 1, 0, 0, 0, 312, 315, 1, 0, 0, 0, 313, 311, 1, 0, 0, 0, 314, 316,
		5, 21, 0, 0, 315, 314, 1, 0, 0, 0, 315, 316, 1, 0, 0, 0, 316, 318, 1, 0,
		0, 0, 317, 319, 3, 0, 0, 0, 318, 317, 1, 0, 0, 0, 318, 319, 1, 0, 0, 0,
		319, 321, 1, 0, 0, 0, 320, 289, 1, 0, 0, 0, 320, 321, 1, 0, 0, 0, 321,
		322, 1, 0, 0, 0, 322, 323, 5, 28, 0, 0, 323, 55, 1, 0, 0, 0, 37, 61, 65,
		70, 77, 82, 87, 96, 109, 114, 119, 122, 130, 137, 143, 162, 167, 174, 181,
		186, 191, 200, 209, 216, 222, 234, 242, 268, 270, 278, 289, 295, 301, 307,
		311, 315, 318, 320,
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
	preludioParserLE                  = 16
	preludioParserGE                  = 17
	preludioParserAT                  = 18
	preludioParserBAR                 = 19
	preludioParserCOLON               = 20
	preludioParserCOMMA               = 21
	preludioParserDOT                 = 22
	preludioParserDOLLAR              = 23
	preludioParserRANGE               = 24
	preludioParserLANG                = 25
	preludioParserRANG                = 26
	preludioParserLBRACKET            = 27
	preludioParserRBRACKET            = 28
	preludioParserLPAREN              = 29
	preludioParserRPAREN              = 30
	preludioParserLBRACE              = 31
	preludioParserRBRACE              = 32
	preludioParserUNDERSCORE          = 33
	preludioParserBACKTICK            = 34
	preludioParserDOUBLE_QUOTE        = 35
	preludioParserSINGLE_QUOTE        = 36
	preludioParserTRIPLE_DOUBLE_QUOTE = 37
	preludioParserTRIPLE_SINGLE_QUOTE = 38
	preludioParserAND                 = 39
	preludioParserOR                  = 40
	preludioParserNOT                 = 41
	preludioParserCOALESCE            = 42
	preludioParserNULL_               = 43
	preludioParserIDENT               = 44
	preludioParserIDENT_START         = 45
	preludioParserIDENT_NEXT          = 46
	preludioParserWHITESPACE          = 47
	preludioParserNEWLINE             = 48
	preludioParserSINGLE_LINE_COMMENT = 49
	preludioParserBOOL_LIT            = 50
	preludioParserINT_LIT             = 51
	preludioParserRNG_LIT             = 52
	preludioParserFLT_LIT             = 53
	preludioParserSTR_CHAR            = 54
	preludioParserSTR_RAW_CHAR        = 55
	preludioParserSTR_LIT             = 56
	preludioParserSTR_INTERP          = 57
	preludioParserSTR_RAW             = 58
	preludioParserSTR_PATH            = 59
	preludioParserRXP_LIT             = 60
	preludioParserDAT_LIT             = 61
	preludioParserDUR_LIT             = 62
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
	preludioParserRULE_pipeline       = 13
	preludioParserRULE_inlinePipeline = 14
	preludioParserRULE_nestedPipeline = 15
	preludioParserRULE_identBacktick  = 16
	preludioParserRULE_funcCall       = 17
	preludioParserRULE_funcCallParam  = 18
	preludioParserRULE_namedArg       = 19
	preludioParserRULE_assign         = 20
	preludioParserRULE_multiAssign    = 21
	preludioParserRULE_exprCall       = 22
	preludioParserRULE_expr           = 23
	preludioParserRULE_term           = 24
	preludioParserRULE_exprUnary      = 25
	preludioParserRULE_literal        = 26
	preludioParserRULE_list           = 27
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
		p.Match(preludioParserNEWLINE)
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

func (s *ProgramContext) AllPipeline() []IPipelineContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPipelineContext); ok {
			len++
		}
	}

	tst := make([]IPipelineContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPipelineContext); ok {
			tst[i] = t.(IPipelineContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Pipeline(i int) IPipelineContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPipelineContext); ok {
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

	return t.(IPipelineContext)
}

func (s *ProgramContext) AllInlinePipeline() []IInlinePipelineContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IInlinePipelineContext); ok {
			len++
		}
	}

	tst := make([]IInlinePipelineContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IInlinePipelineContext); ok {
			tst[i] = t.(IInlinePipelineContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) InlinePipeline(i int) IInlinePipelineContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInlinePipelineContext); ok {
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

	return t.(IInlinePipelineContext)
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

	for _la == preludioParserNEWLINE {
		{
			p.SetState(67)
			p.Nl()
		}

		p.SetState(72)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&9168231546572767626) != 0 {
		p.SetState(77)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(73)
				p.FuncDef()
			}

		case 2:
			{
				p.SetState(74)
				p.Pipeline()
			}

		case 3:
			{
				p.SetState(75)
				p.InlinePipeline()
			}

		case 4:
			{
				p.SetState(76)
				p.Stmt()
			}

		}
		p.SetState(82)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == preludioParserNEWLINE {
			{
				p.SetState(79)
				p.Nl()
			}

			p.SetState(84)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

		p.SetState(89)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(90)
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
		p.SetState(92)
		p.Match(preludioParserPRQL)
	}
	p.SetState(96)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == preludioParserIDENT {
		{
			p.SetState(93)
			p.NamedArg()
		}

		p.SetState(98)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(99)
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

func (s *FuncDefContext) Expr() IExprContext {
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
		p.SetState(101)
		p.Match(preludioParserFUNC)
	}
	{
		p.SetState(102)
		p.FuncDefName()
	}
	{
		p.SetState(103)
		p.FuncDefParams()
	}
	{
		p.SetState(104)
		p.Match(preludioParserARROW)
	}
	{
		p.SetState(105)
		p.expr(0)
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
		p.SetState(107)
		p.Match(preludioParserIDENT)
	}
	p.SetState(109)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == preludioParserLANG {
		{
			p.SetState(108)
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
	p.SetState(114)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == preludioParserIDENT {
		{
			p.SetState(111)
			p.FuncDefParam()
		}

		p.SetState(116)
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
	p.SetState(119)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(117)
			p.NamedArg()
		}

	case 2:
		{
			p.SetState(118)
			p.Match(preludioParserIDENT)
		}

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
		p.SetState(124)
		p.Match(preludioParserLANG)
	}
	{
		p.SetState(125)
		p.TypeTerm()
	}
	{
		p.SetState(126)
		p.Match(preludioParserBAR)
	}
	p.SetState(130)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == preludioParserIDENT {
		{
			p.SetState(127)
			p.TypeTerm()
		}

		p.SetState(132)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(133)
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
		p.SetState(135)
		p.Match(preludioParserIDENT)
	}
	p.SetState(137)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == preludioParserLANG {
		{
			p.SetState(136)
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

func (s *StmtContext) Expr() IExprContext {
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

	p.SetState(143)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(139)
			p.VarAssignStmt()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(140)
			p.VarDeclStmt()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(141)
			p.RetStmt()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(142)
			p.expr(0)
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
		p.SetState(145)
		p.Match(preludioParserIDENT)
	}
	{
		p.SetState(146)
		p.Match(preludioParserASSIGN)
	}
	{
		p.SetState(147)
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
		p.SetState(149)
		p.Match(preludioParserIDENT)
	}
	{
		p.SetState(150)
		p.Match(preludioParserDECLARE)
	}
	{
		p.SetState(151)
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
		p.SetState(153)
		p.Match(preludioParserRET)
	}
	{
		p.SetState(154)
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

func (s *PipelineContext) EOF() antlr.TerminalNode {
	return s.GetToken(preludioParserEOF, 0)
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
	p.EnterRule(localctx, 26, preludioParserRULE_pipeline)

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
		p.SetState(156)
		p.ExprCall()
	}
	p.SetState(162)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(157)
				p.Nl()
			}
			{
				p.SetState(158)
				p.FuncCall()
			}

		}
		p.SetState(164)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext())
	}
	p.SetState(167)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case preludioParserNEWLINE:
		{
			p.SetState(165)
			p.Nl()
		}

	case preludioParserEOF:
		{
			p.SetState(166)
			p.Match(preludioParserEOF)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
	p.EnterRule(localctx, 28, preludioParserRULE_inlinePipeline)
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
		p.SetState(169)
		p.ExprCall()
	}
	p.SetState(174)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == preludioParserBAR {
		{
			p.SetState(170)
			p.Match(preludioParserBAR)
		}
		{
			p.SetState(171)
			p.FuncCall()
		}

		p.SetState(176)
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
	p.EnterRule(localctx, 30, preludioParserRULE_nestedPipeline)
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
		p.SetState(177)
		p.Match(preludioParserLPAREN)
	}
	p.SetState(181)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == preludioParserNEWLINE {
		{
			p.SetState(178)
			p.Nl()
		}

		p.SetState(183)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(186)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(184)
			p.Pipeline()
		}

	case 2:
		{
			p.SetState(185)
			p.InlinePipeline()
		}

	}
	p.SetState(191)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == preludioParserNEWLINE {
		{
			p.SetState(188)
			p.Nl()
		}

		p.SetState(193)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(194)
		p.Match(preludioParserRPAREN)
	}

	return localctx
}

// IIdentBacktickContext is an interface to support dynamic dispatch.
type IIdentBacktickContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdentBacktickContext differentiates from other interfaces.
	IsIdentBacktickContext()
}

type IdentBacktickContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentBacktickContext() *IdentBacktickContext {
	var p = new(IdentBacktickContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = preludioParserRULE_identBacktick
	return p
}

func (*IdentBacktickContext) IsIdentBacktickContext() {}

func NewIdentBacktickContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentBacktickContext {
	var p = new(IdentBacktickContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = preludioParserRULE_identBacktick

	return p
}

func (s *IdentBacktickContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentBacktickContext) AllBACKTICK() []antlr.TerminalNode {
	return s.GetTokens(preludioParserBACKTICK)
}

func (s *IdentBacktickContext) BACKTICK(i int) antlr.TerminalNode {
	return s.GetToken(preludioParserBACKTICK, i)
}

func (s *IdentBacktickContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(preludioParserNEWLINE)
}

func (s *IdentBacktickContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(preludioParserNEWLINE, i)
}

func (s *IdentBacktickContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentBacktickContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentBacktickContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioParserListener); ok {
		listenerT.EnterIdentBacktick(s)
	}
}

func (s *IdentBacktickContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioParserListener); ok {
		listenerT.ExitIdentBacktick(s)
	}
}

func (p *preludioParser) IdentBacktick() (localctx IIdentBacktickContext) {
	this := p
	_ = this

	localctx = NewIdentBacktickContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, preludioParserRULE_identBacktick)
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
		p.SetState(196)
		p.Match(preludioParserBACKTICK)
	}
	p.SetState(200)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&9223090544698195966) != 0 {
		{
			p.SetState(197)
			_la = p.GetTokenStream().LA(1)

			if _la <= 0 || _la == preludioParserBACKTICK || _la == preludioParserNEWLINE {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(202)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(203)
		p.Match(preludioParserBACKTICK)
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
		p.SetState(205)
		p.Match(preludioParserIDENT)
	}
	p.SetState(209)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(206)
				p.FuncCallParam()
			}

		}
		p.SetState(211)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext())
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

func (s *FuncCallParamContext) Expr() IExprContext {
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

	p.SetState(216)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(212)
			p.NamedArg()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(213)
			p.Assign()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(214)
			p.MultiAssign()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(215)
			p.expr(0)
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

func (s *NamedArgContext) Expr() IExprContext {
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
		p.SetState(218)
		p.Match(preludioParserIDENT)
	}
	{
		p.SetState(219)
		p.Match(preludioParserCOLON)
	}
	p.SetState(222)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(220)
			p.Assign()
		}

	case 2:
		{
			p.SetState(221)
			p.expr(0)
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
		p.SetState(224)
		p.Match(preludioParserIDENT)
	}
	{
		p.SetState(225)
		p.Match(preludioParserASSIGN)
	}
	{
		p.SetState(226)
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
		p.SetState(228)
		p.List()
	}
	{
		p.SetState(229)
		p.Match(preludioParserASSIGN)
	}
	{
		p.SetState(230)
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
	p.EnterRule(localctx, 44, preludioParserRULE_exprCall)

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

	p.SetState(234)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(232)
			p.FuncCall()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(233)
			p.expr(0)
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

func (s *ExprContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(preludioParserLPAREN, 0)
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

func (s *ExprContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(preludioParserRPAREN, 0)
}

func (s *ExprContext) Term() ITermContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITermContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *ExprContext) DOLLAR() antlr.TerminalNode {
	return s.GetToken(preludioParserDOLLAR, 0)
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

func (s *ExprContext) MINUS() antlr.TerminalNode {
	return s.GetToken(preludioParserMINUS, 0)
}

func (s *ExprContext) PLUS() antlr.TerminalNode {
	return s.GetToken(preludioParserPLUS, 0)
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

func (s *ExprContext) LE() antlr.TerminalNode {
	return s.GetToken(preludioParserLE, 0)
}

func (s *ExprContext) LANG() antlr.TerminalNode {
	return s.GetToken(preludioParserLANG, 0)
}

func (s *ExprContext) RANG() antlr.TerminalNode {
	return s.GetToken(preludioParserRANG, 0)
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
	_startState := 46
	p.EnterRecursionRule(localctx, 46, preludioParserRULE_expr, _p)
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
	p.SetState(242)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(237)
			p.Match(preludioParserLPAREN)
		}
		{
			p.SetState(238)
			p.expr(0)
		}
		{
			p.SetState(239)
			p.Match(preludioParserRPAREN)
		}

	case 2:
		{
			p.SetState(241)
			p.Term()
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(270)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(268)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, preludioParserRULE_expr)
				p.SetState(244)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(245)
					p.Match(preludioParserDOLLAR)
				}
				{
					p.SetState(246)
					p.expr(11)
				}

			case 2:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, preludioParserRULE_expr)
				p.SetState(247)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(248)
					p.Match(preludioParserEXP)
				}
				{
					p.SetState(249)
					p.expr(10)
				}

			case 3:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, preludioParserRULE_expr)
				p.SetState(250)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(251)
					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&6656) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(252)
					p.expr(9)
				}

			case 4:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, preludioParserRULE_expr)
				p.SetState(253)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(254)
					_la = p.GetTokenStream().LA(1)

					if !(_la == preludioParserPLUS || _la == preludioParserMINUS) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(255)
					p.expr(8)
				}

			case 5:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, preludioParserRULE_expr)
				p.SetState(256)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(257)
					p.Match(preludioParserMODEL)
				}
				{
					p.SetState(258)
					p.expr(7)
				}

			case 6:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, preludioParserRULE_expr)
				p.SetState(259)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(260)
					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&100909056) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(261)
					p.expr(6)
				}

			case 7:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, preludioParserRULE_expr)
				p.SetState(262)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(263)
					p.Match(preludioParserCOALESCE)
				}
				{
					p.SetState(264)
					p.expr(5)
				}

			case 8:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, preludioParserRULE_expr)
				p.SetState(265)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(266)
					_la = p.GetTokenStream().LA(1)

					if !(_la == preludioParserAND || _la == preludioParserOR) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(267)
					p.expr(4)
				}

			}

		}
		p.SetState(272)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext())
	}

	return localctx
}

// ITermContext is an interface to support dynamic dispatch.
type ITermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTermContext differentiates from other interfaces.
	IsTermContext()
}

type TermContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTermContext() *TermContext {
	var p = new(TermContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = preludioParserRULE_term
	return p
}

func (*TermContext) IsTermContext() {}

func NewTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermContext {
	var p = new(TermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = preludioParserRULE_term

	return p
}

func (s *TermContext) GetParser() antlr.Parser { return s.parser }

func (s *TermContext) Literal() ILiteralContext {
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

func (s *TermContext) IdentBacktick() IIdentBacktickContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentBacktickContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentBacktickContext)
}

func (s *TermContext) ExprUnary() IExprUnaryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprUnaryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprUnaryContext)
}

func (s *TermContext) List() IListContext {
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

func (s *TermContext) NestedPipeline() INestedPipelineContext {
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

func (s *TermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioParserListener); ok {
		listenerT.EnterTerm(s)
	}
}

func (s *TermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioParserListener); ok {
		listenerT.ExitTerm(s)
	}
}

func (p *preludioParser) Term() (localctx ITermContext) {
	this := p
	_ = this

	localctx = NewTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, preludioParserRULE_term)

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

	p.SetState(278)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case preludioParserNULL_, preludioParserIDENT, preludioParserBOOL_LIT, preludioParserINT_LIT, preludioParserRNG_LIT, preludioParserFLT_LIT, preludioParserSTR_LIT, preludioParserSTR_INTERP, preludioParserSTR_RAW, preludioParserSTR_PATH, preludioParserRXP_LIT, preludioParserDAT_LIT, preludioParserDUR_LIT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(273)
			p.Literal()
		}

	case preludioParserBACKTICK:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(274)
			p.IdentBacktick()
		}

	case preludioParserPLUS, preludioParserMINUS, preludioParserNOT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(275)
			p.ExprUnary()
		}

	case preludioParserLBRACKET:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(276)
			p.List()
		}

	case preludioParserLPAREN:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(277)
			p.NestedPipeline()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IExprUnaryContext is an interface to support dynamic dispatch.
type IExprUnaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprUnaryContext differentiates from other interfaces.
	IsExprUnaryContext()
}

type ExprUnaryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprUnaryContext() *ExprUnaryContext {
	var p = new(ExprUnaryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = preludioParserRULE_exprUnary
	return p
}

func (*ExprUnaryContext) IsExprUnaryContext() {}

func NewExprUnaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprUnaryContext {
	var p = new(ExprUnaryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = preludioParserRULE_exprUnary

	return p
}

func (s *ExprUnaryContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprUnaryContext) Term() ITermContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITermContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *ExprUnaryContext) MINUS() antlr.TerminalNode {
	return s.GetToken(preludioParserMINUS, 0)
}

func (s *ExprUnaryContext) PLUS() antlr.TerminalNode {
	return s.GetToken(preludioParserPLUS, 0)
}

func (s *ExprUnaryContext) NOT() antlr.TerminalNode {
	return s.GetToken(preludioParserNOT, 0)
}

func (s *ExprUnaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprUnaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprUnaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioParserListener); ok {
		listenerT.EnterExprUnary(s)
	}
}

func (s *ExprUnaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioParserListener); ok {
		listenerT.ExitExprUnary(s)
	}
}

func (p *preludioParser) ExprUnary() (localctx IExprUnaryContext) {
	this := p
	_ = this

	localctx = NewExprUnaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, preludioParserRULE_exprUnary)
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
		p.SetState(280)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2199023255936) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(281)
		p.Term()
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

func (s *LiteralContext) IDENT() antlr.TerminalNode {
	return s.GetToken(preludioParserIDENT, 0)
}

func (s *LiteralContext) NULL_() antlr.TerminalNode {
	return s.GetToken(preludioParserNULL_, 0)
}

func (s *LiteralContext) BOOL_LIT() antlr.TerminalNode {
	return s.GetToken(preludioParserBOOL_LIT, 0)
}

func (s *LiteralContext) INT_LIT() antlr.TerminalNode {
	return s.GetToken(preludioParserINT_LIT, 0)
}

func (s *LiteralContext) RNG_LIT() antlr.TerminalNode {
	return s.GetToken(preludioParserRNG_LIT, 0)
}

func (s *LiteralContext) FLT_LIT() antlr.TerminalNode {
	return s.GetToken(preludioParserFLT_LIT, 0)
}

func (s *LiteralContext) STR_LIT() antlr.TerminalNode {
	return s.GetToken(preludioParserSTR_LIT, 0)
}

func (s *LiteralContext) STR_INTERP() antlr.TerminalNode {
	return s.GetToken(preludioParserSTR_INTERP, 0)
}

func (s *LiteralContext) STR_RAW() antlr.TerminalNode {
	return s.GetToken(preludioParserSTR_RAW, 0)
}

func (s *LiteralContext) STR_PATH() antlr.TerminalNode {
	return s.GetToken(preludioParserSTR_PATH, 0)
}

func (s *LiteralContext) RXP_LIT() antlr.TerminalNode {
	return s.GetToken(preludioParserRXP_LIT, 0)
}

func (s *LiteralContext) DAT_LIT() antlr.TerminalNode {
	return s.GetToken(preludioParserDAT_LIT, 0)
}

func (s *LiteralContext) DUR_LIT() antlr.TerminalNode {
	return s.GetToken(preludioParserDUR_LIT, 0)
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
	p.EnterRule(localctx, 52, preludioParserRULE_literal)
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
		p.SetState(283)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&9168229329698553856) != 0) {
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
	p.EnterRule(localctx, 54, preludioParserRULE_list)
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
		p.SetState(285)
		p.Match(preludioParserLBRACKET)
	}
	p.SetState(320)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&9168513021549478272) != 0 {
		p.SetState(289)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == preludioParserNEWLINE {
			{
				p.SetState(286)
				p.Nl()
			}

			p.SetState(291)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(295)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(292)
				p.Assign()
			}

		case 2:
			{
				p.SetState(293)
				p.MultiAssign()
			}

		case 3:
			{
				p.SetState(294)
				p.ExprCall()
			}

		}
		p.SetState(311)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(297)
					p.Match(preludioParserCOMMA)
				}
				p.SetState(301)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				for _la == preludioParserNEWLINE {
					{
						p.SetState(298)
						p.Nl()
					}

					p.SetState(303)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)
				}
				p.SetState(307)
				p.GetErrorHandler().Sync(p)
				switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext()) {
				case 1:
					{
						p.SetState(304)
						p.Assign()
					}

				case 2:
					{
						p.SetState(305)
						p.MultiAssign()
					}

				case 3:
					{
						p.SetState(306)
						p.ExprCall()
					}

				}

			}
			p.SetState(313)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext())
		}
		p.SetState(315)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == preludioParserCOMMA {
			{
				p.SetState(314)
				p.Match(preludioParserCOMMA)
			}

		}
		p.SetState(318)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == preludioParserNEWLINE {
			{
				p.SetState(317)
				p.Nl()
			}

		}

	}
	{
		p.SetState(322)
		p.Match(preludioParserRBRACKET)
	}

	return localctx
}

func (p *preludioParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 23:
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
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
