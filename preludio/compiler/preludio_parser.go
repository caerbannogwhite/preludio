// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package compiler // preludio
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

var preludioParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func preludioParserInit() {
	staticData := &preludioParserStaticData
	staticData.literalNames = []string{
		"", "'func'", "'prql'", "'let'", "'->'", "'='", "'+'", "'-'", "'*'",
		"'**'", "'/'", "'%'", "'~'", "'=='", "'!='", "'<='", "'>='", "'|'",
		"':'", "','", "'.'", "'$'", "'..'", "'<'", "'>'", "'['", "']'", "'('",
		"')'", "'_'", "'`'", "'\"'", "'''", "'\"\"\"'", "'''''", "'and'", "'or'",
		"'not'", "'??'", "'null'",
	}
	staticData.symbolicNames = []string{
		"", "FUNC", "PRQL", "LET", "ARROW", "ASSIGN", "PLUS", "MINUS", "STAR",
		"POW", "DIV", "MOD", "MODEL", "EQ", "NE", "LE", "GE", "BAR", "COLON",
		"COMMA", "DOT", "DOLLAR", "RANGE", "LANG", "RANG", "LBRACKET", "RBRACKET",
		"LPAREN", "RPAREN", "UNDERSCORE", "BACKTICK", "DOUBLE_QUOTE", "SINGLE_QUOTE",
		"TRIPLE_DOUBLE_QUOTE", "TRIPLE_SINGLE_QUOTE", "AND", "OR", "NOT", "COALESCE",
		"NULL_", "BOOLEAN", "INTEGER", "FLOAT", "IDENT", "IDENT_START", "IDENT_NEXT",
		"WHITESPACE", "NEWLINE", "COMMENT", "INTERVAL_KIND", "STRING",
	}
	staticData.ruleNames = []string{
		"nl", "program", "programIntro", "funcDef", "funcDefName", "funcDefParams",
		"funcDefParam", "typeDef", "typeTerm", "stmt", "assignStmt", "pipeline",
		"inlinePipeline", "identBacktick", "funcCall", "funcCallParam", "namedArg",
		"assign", "multiAssign", "exprCall", "expr", "term", "exprUnary", "literal",
		"list", "nestedPipeline",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 50, 321, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 1, 0,
		1, 0, 1, 1, 5, 1, 56, 8, 1, 10, 1, 12, 1, 59, 9, 1, 1, 1, 3, 1, 62, 8,
		1, 1, 1, 5, 1, 65, 8, 1, 10, 1, 12, 1, 68, 9, 1, 1, 1, 1, 1, 1, 1, 3, 1,
		73, 8, 1, 1, 1, 5, 1, 76, 8, 1, 10, 1, 12, 1, 79, 9, 1, 5, 1, 81, 8, 1,
		10, 1, 12, 1, 84, 9, 1, 1, 1, 1, 1, 1, 2, 1, 2, 5, 2, 90, 8, 2, 10, 2,
		12, 2, 93, 9, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4,
		1, 4, 3, 4, 105, 8, 4, 1, 5, 5, 5, 108, 8, 5, 10, 5, 12, 5, 111, 9, 5,
		1, 6, 1, 6, 3, 6, 115, 8, 6, 1, 6, 3, 6, 118, 8, 6, 1, 7, 1, 7, 1, 7, 1,
		7, 5, 7, 124, 8, 7, 10, 7, 12, 7, 127, 9, 7, 1, 7, 1, 7, 1, 8, 1, 8, 3,
		8, 133, 8, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1,
		11, 1, 11, 1, 11, 5, 11, 146, 8, 11, 10, 11, 12, 11, 149, 9, 11, 1, 11,
		1, 11, 3, 11, 153, 8, 11, 1, 12, 1, 12, 1, 12, 5, 12, 158, 8, 12, 10, 12,
		12, 12, 161, 9, 12, 1, 13, 1, 13, 5, 13, 165, 8, 13, 10, 13, 12, 13, 168,
		9, 13, 1, 13, 1, 13, 1, 14, 1, 14, 5, 14, 174, 8, 14, 10, 14, 12, 14, 177,
		9, 14, 1, 15, 1, 15, 1, 15, 1, 15, 3, 15, 183, 8, 15, 1, 16, 1, 16, 1,
		16, 1, 16, 3, 16, 189, 8, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18,
		1, 18, 1, 18, 1, 19, 1, 19, 3, 19, 201, 8, 19, 1, 20, 1, 20, 1, 20, 1,
		20, 1, 20, 1, 20, 3, 20, 209, 8, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20,
		1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1,
		20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 5, 20, 232, 8, 20, 10, 20, 12, 20,
		235, 9, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 3, 21, 242, 8, 21, 1, 22,
		1, 22, 1, 22, 1, 22, 3, 22, 248, 8, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1,
		23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 3, 23, 261, 8, 23, 1, 24,
		1, 24, 5, 24, 265, 8, 24, 10, 24, 12, 24, 268, 9, 24, 1, 24, 1, 24, 1,
		24, 3, 24, 273, 8, 24, 1, 24, 1, 24, 5, 24, 277, 8, 24, 10, 24, 12, 24,
		280, 9, 24, 1, 24, 1, 24, 1, 24, 3, 24, 285, 8, 24, 5, 24, 287, 8, 24,
		10, 24, 12, 24, 290, 9, 24, 1, 24, 3, 24, 293, 8, 24, 1, 24, 3, 24, 296,
		8, 24, 3, 24, 298, 8, 24, 1, 24, 1, 24, 1, 25, 1, 25, 5, 25, 304, 8, 25,
		10, 25, 12, 25, 307, 9, 25, 1, 25, 1, 25, 3, 25, 311, 8, 25, 1, 25, 5,
		25, 314, 8, 25, 10, 25, 12, 25, 317, 9, 25, 1, 25, 1, 25, 1, 25, 0, 1,
		40, 26, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32,
		34, 36, 38, 40, 42, 44, 46, 48, 50, 0, 9, 1, 0, 47, 48, 2, 0, 30, 30, 47,
		47, 2, 0, 8, 8, 10, 11, 1, 0, 6, 7, 2, 0, 13, 16, 23, 24, 1, 0, 35, 36,
		2, 0, 6, 7, 37, 37, 1, 0, 41, 42, 1, 0, 41, 43, 352, 0, 52, 1, 0, 0, 0,
		2, 57, 1, 0, 0, 0, 4, 87, 1, 0, 0, 0, 6, 96, 1, 0, 0, 0, 8, 102, 1, 0,
		0, 0, 10, 109, 1, 0, 0, 0, 12, 114, 1, 0, 0, 0, 14, 119, 1, 0, 0, 0, 16,
		130, 1, 0, 0, 0, 18, 134, 1, 0, 0, 0, 20, 136, 1, 0, 0, 0, 22, 141, 1,
		0, 0, 0, 24, 154, 1, 0, 0, 0, 26, 162, 1, 0, 0, 0, 28, 171, 1, 0, 0, 0,
		30, 182, 1, 0, 0, 0, 32, 184, 1, 0, 0, 0, 34, 190, 1, 0, 0, 0, 36, 194,
		1, 0, 0, 0, 38, 200, 1, 0, 0, 0, 40, 208, 1, 0, 0, 0, 42, 241, 1, 0, 0,
		0, 44, 243, 1, 0, 0, 0, 46, 260, 1, 0, 0, 0, 48, 262, 1, 0, 0, 0, 50, 301,
		1, 0, 0, 0, 52, 53, 7, 0, 0, 0, 53, 1, 1, 0, 0, 0, 54, 56, 3, 0, 0, 0,
		55, 54, 1, 0, 0, 0, 56, 59, 1, 0, 0, 0, 57, 55, 1, 0, 0, 0, 57, 58, 1,
		0, 0, 0, 58, 61, 1, 0, 0, 0, 59, 57, 1, 0, 0, 0, 60, 62, 3, 4, 2, 0, 61,
		60, 1, 0, 0, 0, 61, 62, 1, 0, 0, 0, 62, 66, 1, 0, 0, 0, 63, 65, 3, 0, 0,
		0, 64, 63, 1, 0, 0, 0, 65, 68, 1, 0, 0, 0, 66, 64, 1, 0, 0, 0, 66, 67,
		1, 0, 0, 0, 67, 82, 1, 0, 0, 0, 68, 66, 1, 0, 0, 0, 69, 73, 3, 6, 3, 0,
		70, 73, 3, 18, 9, 0, 71, 73, 3, 22, 11, 0, 72, 69, 1, 0, 0, 0, 72, 70,
		1, 0, 0, 0, 72, 71, 1, 0, 0, 0, 73, 77, 1, 0, 0, 0, 74, 76, 3, 0, 0, 0,
		75, 74, 1, 0, 0, 0, 76, 79, 1, 0, 0, 0, 77, 75, 1, 0, 0, 0, 77, 78, 1,
		0, 0, 0, 78, 81, 1, 0, 0, 0, 79, 77, 1, 0, 0, 0, 80, 72, 1, 0, 0, 0, 81,
		84, 1, 0, 0, 0, 82, 80, 1, 0, 0, 0, 82, 83, 1, 0, 0, 0, 83, 85, 1, 0, 0,
		0, 84, 82, 1, 0, 0, 0, 85, 86, 5, 0, 0, 1, 86, 3, 1, 0, 0, 0, 87, 91, 5,
		2, 0, 0, 88, 90, 3, 32, 16, 0, 89, 88, 1, 0, 0, 0, 90, 93, 1, 0, 0, 0,
		91, 89, 1, 0, 0, 0, 91, 92, 1, 0, 0, 0, 92, 94, 1, 0, 0, 0, 93, 91, 1,
		0, 0, 0, 94, 95, 3, 0, 0, 0, 95, 5, 1, 0, 0, 0, 96, 97, 5, 1, 0, 0, 97,
		98, 3, 8, 4, 0, 98, 99, 3, 10, 5, 0, 99, 100, 5, 4, 0, 0, 100, 101, 3,
		40, 20, 0, 101, 7, 1, 0, 0, 0, 102, 104, 5, 43, 0, 0, 103, 105, 3, 14,
		7, 0, 104, 103, 1, 0, 0, 0, 104, 105, 1, 0, 0, 0, 105, 9, 1, 0, 0, 0, 106,
		108, 3, 12, 6, 0, 107, 106, 1, 0, 0, 0, 108, 111, 1, 0, 0, 0, 109, 107,
		1, 0, 0, 0, 109, 110, 1, 0, 0, 0, 110, 11, 1, 0, 0, 0, 111, 109, 1, 0,
		0, 0, 112, 115, 3, 32, 16, 0, 113, 115, 5, 43, 0, 0, 114, 112, 1, 0, 0,
		0, 114, 113, 1, 0, 0, 0, 115, 117, 1, 0, 0, 0, 116, 118, 3, 14, 7, 0, 117,
		116, 1, 0, 0, 0, 117, 118, 1, 0, 0, 0, 118, 13, 1, 0, 0, 0, 119, 120, 5,
		23, 0, 0, 120, 121, 3, 16, 8, 0, 121, 125, 5, 17, 0, 0, 122, 124, 3, 16,
		8, 0, 123, 122, 1, 0, 0, 0, 124, 127, 1, 0, 0, 0, 125, 123, 1, 0, 0, 0,
		125, 126, 1, 0, 0, 0, 126, 128, 1, 0, 0, 0, 127, 125, 1, 0, 0, 0, 128,
		129, 5, 24, 0, 0, 129, 15, 1, 0, 0, 0, 130, 132, 5, 43, 0, 0, 131, 133,
		3, 14, 7, 0, 132, 131, 1, 0, 0, 0, 132, 133, 1, 0, 0, 0, 133, 17, 1, 0,
		0, 0, 134, 135, 3, 20, 10, 0, 135, 19, 1, 0, 0, 0, 136, 137, 5, 3, 0, 0,
		137, 138, 5, 43, 0, 0, 138, 139, 5, 5, 0, 0, 139, 140, 3, 40, 20, 0, 140,
		21, 1, 0, 0, 0, 141, 147, 3, 38, 19, 0, 142, 143, 3, 0, 0, 0, 143, 144,
		3, 28, 14, 0, 144, 146, 1, 0, 0, 0, 145, 142, 1, 0, 0, 0, 146, 149, 1,
		0, 0, 0, 147, 145, 1, 0, 0, 0, 147, 148, 1, 0, 0, 0, 148, 152, 1, 0, 0,
		0, 149, 147, 1, 0, 0, 0, 150, 153, 3, 0, 0, 0, 151, 153, 5, 0, 0, 1, 152,
		150, 1, 0, 0, 0, 152, 151, 1, 0, 0, 0, 153, 23, 1, 0, 0, 0, 154, 159, 3,
		38, 19, 0, 155, 156, 5, 17, 0, 0, 156, 158, 3, 28, 14, 0, 157, 155, 1,
		0, 0, 0, 158, 161, 1, 0, 0, 0, 159, 157, 1, 0, 0, 0, 159, 160, 1, 0, 0,
		0, 160, 25, 1, 0, 0, 0, 161, 159, 1, 0, 0, 0, 162, 166, 5, 30, 0, 0, 163,
		165, 8, 1, 0, 0, 164, 163, 1, 0, 0, 0, 165, 168, 1, 0, 0, 0, 166, 164,
		1, 0, 0, 0, 166, 167, 1, 0, 0, 0, 167, 169, 1, 0, 0, 0, 168, 166, 1, 0,
		0, 0, 169, 170, 5, 30, 0, 0, 170, 27, 1, 0, 0, 0, 171, 175, 5, 43, 0, 0,
		172, 174, 3, 30, 15, 0, 173, 172, 1, 0, 0, 0, 174, 177, 1, 0, 0, 0, 175,
		173, 1, 0, 0, 0, 175, 176, 1, 0, 0, 0, 176, 29, 1, 0, 0, 0, 177, 175, 1,
		0, 0, 0, 178, 183, 3, 32, 16, 0, 179, 183, 3, 34, 17, 0, 180, 183, 3, 36,
		18, 0, 181, 183, 3, 40, 20, 0, 182, 178, 1, 0, 0, 0, 182, 179, 1, 0, 0,
		0, 182, 180, 1, 0, 0, 0, 182, 181, 1, 0, 0, 0, 183, 31, 1, 0, 0, 0, 184,
		185, 5, 43, 0, 0, 185, 188, 5, 18, 0, 0, 186, 189, 3, 34, 17, 0, 187, 189,
		3, 40, 20, 0, 188, 186, 1, 0, 0, 0, 188, 187, 1, 0, 0, 0, 189, 33, 1, 0,
		0, 0, 190, 191, 5, 43, 0, 0, 191, 192, 5, 5, 0, 0, 192, 193, 3, 38, 19,
		0, 193, 35, 1, 0, 0, 0, 194, 195, 3, 48, 24, 0, 195, 196, 5, 5, 0, 0, 196,
		197, 3, 38, 19, 0, 197, 37, 1, 0, 0, 0, 198, 201, 3, 40, 20, 0, 199, 201,
		3, 28, 14, 0, 200, 198, 1, 0, 0, 0, 200, 199, 1, 0, 0, 0, 201, 39, 1, 0,
		0, 0, 202, 203, 6, 20, -1, 0, 203, 204, 5, 27, 0, 0, 204, 205, 3, 40, 20,
		0, 205, 206, 5, 28, 0, 0, 206, 209, 1, 0, 0, 0, 207, 209, 3, 42, 21, 0,
		208, 202, 1, 0, 0, 0, 208, 207, 1, 0, 0, 0, 209, 233, 1, 0, 0, 0, 210,
		211, 10, 9, 0, 0, 211, 212, 7, 2, 0, 0, 212, 232, 3, 40, 20, 10, 213, 214,
		10, 8, 0, 0, 214, 215, 7, 3, 0, 0, 215, 232, 3, 40, 20, 9, 216, 217, 10,
		7, 0, 0, 217, 218, 5, 9, 0, 0, 218, 232, 3, 40, 20, 8, 219, 220, 10, 6,
		0, 0, 220, 221, 5, 12, 0, 0, 221, 232, 3, 40, 20, 7, 222, 223, 10, 5, 0,
		0, 223, 224, 7, 4, 0, 0, 224, 232, 3, 40, 20, 6, 225, 226, 10, 4, 0, 0,
		226, 227, 5, 38, 0, 0, 227, 232, 3, 40, 20, 5, 228, 229, 10, 3, 0, 0, 229,
		230, 7, 5, 0, 0, 230, 232, 3, 40, 20, 4, 231, 210, 1, 0, 0, 0, 231, 213,
		1, 0, 0, 0, 231, 216, 1, 0, 0, 0, 231, 219, 1, 0, 0, 0, 231, 222, 1, 0,
		0, 0, 231, 225, 1, 0, 0, 0, 231, 228, 1, 0, 0, 0, 232, 235, 1, 0, 0, 0,
		233, 231, 1, 0, 0, 0, 233, 234, 1, 0, 0, 0, 234, 41, 1, 0, 0, 0, 235, 233,
		1, 0, 0, 0, 236, 242, 3, 46, 23, 0, 237, 242, 3, 26, 13, 0, 238, 242, 3,
		44, 22, 0, 239, 242, 3, 48, 24, 0, 240, 242, 3, 50, 25, 0, 241, 236, 1,
		0, 0, 0, 241, 237, 1, 0, 0, 0, 241, 238, 1, 0, 0, 0, 241, 239, 1, 0, 0,
		0, 241, 240, 1, 0, 0, 0, 242, 43, 1, 0, 0, 0, 243, 247, 7, 6, 0, 0, 244,
		248, 3, 50, 25, 0, 245, 248, 3, 46, 23, 0, 246, 248, 5, 43, 0, 0, 247,
		244, 1, 0, 0, 0, 247, 245, 1, 0, 0, 0, 247, 246, 1, 0, 0, 0, 248, 45, 1,
		0, 0, 0, 249, 261, 5, 43, 0, 0, 250, 261, 5, 39, 0, 0, 251, 261, 5, 40,
		0, 0, 252, 261, 5, 50, 0, 0, 253, 261, 5, 41, 0, 0, 254, 261, 5, 42, 0,
		0, 255, 256, 7, 7, 0, 0, 256, 261, 5, 49, 0, 0, 257, 258, 7, 8, 0, 0, 258,
		259, 5, 22, 0, 0, 259, 261, 7, 8, 0, 0, 260, 249, 1, 0, 0, 0, 260, 250,
		1, 0, 0, 0, 260, 251, 1, 0, 0, 0, 260, 252, 1, 0, 0, 0, 260, 253, 1, 0,
		0, 0, 260, 254, 1, 0, 0, 0, 260, 255, 1, 0, 0, 0, 260, 257, 1, 0, 0, 0,
		261, 47, 1, 0, 0, 0, 262, 297, 5, 25, 0, 0, 263, 265, 3, 0, 0, 0, 264,
		263, 1, 0, 0, 0, 265, 268, 1, 0, 0, 0, 266, 264, 1, 0, 0, 0, 266, 267,
		1, 0, 0, 0, 267, 272, 1, 0, 0, 0, 268, 266, 1, 0, 0, 0, 269, 273, 3, 34,
		17, 0, 270, 273, 3, 36, 18, 0, 271, 273, 3, 38, 19, 0, 272, 269, 1, 0,
		0, 0, 272, 270, 1, 0, 0, 0, 272, 271, 1, 0, 0, 0, 273, 288, 1, 0, 0, 0,
		274, 278, 5, 19, 0, 0, 275, 277, 3, 0, 0, 0, 276, 275, 1, 0, 0, 0, 277,
		280, 1, 0, 0, 0, 278, 276, 1, 0, 0, 0, 278, 279, 1, 0, 0, 0, 279, 284,
		1, 0, 0, 0, 280, 278, 1, 0, 0, 0, 281, 285, 3, 34, 17, 0, 282, 285, 3,
		36, 18, 0, 283, 285, 3, 38, 19, 0, 284, 281, 1, 0, 0, 0, 284, 282, 1, 0,
		0, 0, 284, 283, 1, 0, 0, 0, 285, 287, 1, 0, 0, 0, 286, 274, 1, 0, 0, 0,
		287, 290, 1, 0, 0, 0, 288, 286, 1, 0, 0, 0, 288, 289, 1, 0, 0, 0, 289,
		292, 1, 0, 0, 0, 290, 288, 1, 0, 0, 0, 291, 293, 5, 19, 0, 0, 292, 291,
		1, 0, 0, 0, 292, 293, 1, 0, 0, 0, 293, 295, 1, 0, 0, 0, 294, 296, 3, 0,
		0, 0, 295, 294, 1, 0, 0, 0, 295, 296, 1, 0, 0, 0, 296, 298, 1, 0, 0, 0,
		297, 266, 1, 0, 0, 0, 297, 298, 1, 0, 0, 0, 298, 299, 1, 0, 0, 0, 299,
		300, 5, 26, 0, 0, 300, 49, 1, 0, 0, 0, 301, 305, 5, 27, 0, 0, 302, 304,
		3, 0, 0, 0, 303, 302, 1, 0, 0, 0, 304, 307, 1, 0, 0, 0, 305, 303, 1, 0,
		0, 0, 305, 306, 1, 0, 0, 0, 306, 310, 1, 0, 0, 0, 307, 305, 1, 0, 0, 0,
		308, 311, 3, 22, 11, 0, 309, 311, 3, 24, 12, 0, 310, 308, 1, 0, 0, 0, 310,
		309, 1, 0, 0, 0, 311, 315, 1, 0, 0, 0, 312, 314, 3, 0, 0, 0, 313, 312,
		1, 0, 0, 0, 314, 317, 1, 0, 0, 0, 315, 313, 1, 0, 0, 0, 315, 316, 1, 0,
		0, 0, 316, 318, 1, 0, 0, 0, 317, 315, 1, 0, 0, 0, 318, 319, 5, 28, 0, 0,
		319, 51, 1, 0, 0, 0, 38, 57, 61, 66, 72, 77, 82, 91, 104, 109, 114, 117,
		125, 132, 147, 152, 159, 166, 175, 182, 188, 200, 208, 231, 233, 241, 247,
		260, 266, 272, 278, 284, 288, 292, 295, 297, 305, 310, 315,
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
	staticData := &preludioParserStaticData
	staticData.once.Do(preludioParserInit)
}

// NewpreludioParser produces a new parser instance for the optional input antlr.TokenStream.
func NewpreludioParser(input antlr.TokenStream) *preludioParser {
	PreludioParserInit()
	this := new(preludioParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &preludioParserStaticData
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
	preludioParserLET                 = 3
	preludioParserARROW               = 4
	preludioParserASSIGN              = 5
	preludioParserPLUS                = 6
	preludioParserMINUS               = 7
	preludioParserSTAR                = 8
	preludioParserPOW                 = 9
	preludioParserDIV                 = 10
	preludioParserMOD                 = 11
	preludioParserMODEL               = 12
	preludioParserEQ                  = 13
	preludioParserNE                  = 14
	preludioParserLE                  = 15
	preludioParserGE                  = 16
	preludioParserBAR                 = 17
	preludioParserCOLON               = 18
	preludioParserCOMMA               = 19
	preludioParserDOT                 = 20
	preludioParserDOLLAR              = 21
	preludioParserRANGE               = 22
	preludioParserLANG                = 23
	preludioParserRANG                = 24
	preludioParserLBRACKET            = 25
	preludioParserRBRACKET            = 26
	preludioParserLPAREN              = 27
	preludioParserRPAREN              = 28
	preludioParserUNDERSCORE          = 29
	preludioParserBACKTICK            = 30
	preludioParserDOUBLE_QUOTE        = 31
	preludioParserSINGLE_QUOTE        = 32
	preludioParserTRIPLE_DOUBLE_QUOTE = 33
	preludioParserTRIPLE_SINGLE_QUOTE = 34
	preludioParserAND                 = 35
	preludioParserOR                  = 36
	preludioParserNOT                 = 37
	preludioParserCOALESCE            = 38
	preludioParserNULL_               = 39
	preludioParserBOOLEAN             = 40
	preludioParserINTEGER             = 41
	preludioParserFLOAT               = 42
	preludioParserIDENT               = 43
	preludioParserIDENT_START         = 44
	preludioParserIDENT_NEXT          = 45
	preludioParserWHITESPACE          = 46
	preludioParserNEWLINE             = 47
	preludioParserCOMMENT             = 48
	preludioParserINTERVAL_KIND       = 49
	preludioParserSTRING              = 50
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
	preludioParserRULE_assignStmt     = 10
	preludioParserRULE_pipeline       = 11
	preludioParserRULE_inlinePipeline = 12
	preludioParserRULE_identBacktick  = 13
	preludioParserRULE_funcCall       = 14
	preludioParserRULE_funcCallParam  = 15
	preludioParserRULE_namedArg       = 16
	preludioParserRULE_assign         = 17
	preludioParserRULE_multiAssign    = 18
	preludioParserRULE_exprCall       = 19
	preludioParserRULE_expr           = 20
	preludioParserRULE_term           = 21
	preludioParserRULE_exprUnary      = 22
	preludioParserRULE_literal        = 23
	preludioParserRULE_list           = 24
	preludioParserRULE_nestedPipeline = 25
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

func (s *NlContext) COMMENT() antlr.TerminalNode {
	return s.GetToken(preludioParserCOMMENT, 0)
}

func (s *NlContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NlContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NlContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioListener); ok {
		listenerT.EnterNl(s)
	}
}

func (s *NlContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioListener); ok {
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
		p.SetState(52)
		_la = p.GetTokenStream().LA(1)

		if !(_la == preludioParserNEWLINE || _la == preludioParserCOMMENT) {
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

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioListener); ok {
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
	p.SetState(57)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(54)
				p.Nl()
			}

		}
		p.SetState(59)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())
	}
	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == preludioParserPRQL {
		{
			p.SetState(60)
			p.ProgramIntro()
		}

	}
	p.SetState(66)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == preludioParserNEWLINE || _la == preludioParserCOMMENT {
		{
			p.SetState(63)
			p.Nl()
		}

		p.SetState(68)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(82)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1143081017540810) != 0 {
		p.SetState(72)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case preludioParserFUNC:
			{
				p.SetState(69)
				p.FuncDef()
			}

		case preludioParserLET:
			{
				p.SetState(70)
				p.Stmt()
			}

		case preludioParserPLUS, preludioParserMINUS, preludioParserLBRACKET, preludioParserLPAREN, preludioParserBACKTICK, preludioParserNOT, preludioParserNULL_, preludioParserBOOLEAN, preludioParserINTEGER, preludioParserFLOAT, preludioParserIDENT, preludioParserSTRING:
			{
				p.SetState(71)
				p.Pipeline()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}
		p.SetState(77)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == preludioParserNEWLINE || _la == preludioParserCOMMENT {
			{
				p.SetState(74)
				p.Nl()
			}

			p.SetState(79)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

		p.SetState(84)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(85)
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
	if listenerT, ok := listener.(preludioListener); ok {
		listenerT.EnterProgramIntro(s)
	}
}

func (s *ProgramIntroContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioListener); ok {
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
		p.SetState(87)
		p.Match(preludioParserPRQL)
	}
	p.SetState(91)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == preludioParserIDENT {
		{
			p.SetState(88)
			p.NamedArg()
		}

		p.SetState(93)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(94)
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
	if listenerT, ok := listener.(preludioListener); ok {
		listenerT.EnterFuncDef(s)
	}
}

func (s *FuncDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioListener); ok {
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
		p.SetState(96)
		p.Match(preludioParserFUNC)
	}
	{
		p.SetState(97)
		p.FuncDefName()
	}
	{
		p.SetState(98)
		p.FuncDefParams()
	}
	{
		p.SetState(99)
		p.Match(preludioParserARROW)
	}
	{
		p.SetState(100)
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
	if listenerT, ok := listener.(preludioListener); ok {
		listenerT.EnterFuncDefName(s)
	}
}

func (s *FuncDefNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioListener); ok {
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
		p.SetState(102)
		p.Match(preludioParserIDENT)
	}
	p.SetState(104)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == preludioParserLANG {
		{
			p.SetState(103)
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
	if listenerT, ok := listener.(preludioListener); ok {
		listenerT.EnterFuncDefParams(s)
	}
}

func (s *FuncDefParamsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioListener); ok {
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
	p.SetState(109)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == preludioParserIDENT {
		{
			p.SetState(106)
			p.FuncDefParam()
		}

		p.SetState(111)
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
	if listenerT, ok := listener.(preludioListener); ok {
		listenerT.EnterFuncDefParam(s)
	}
}

func (s *FuncDefParamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioListener); ok {
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
	p.SetState(114)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(112)
			p.NamedArg()
		}

	case 2:
		{
			p.SetState(113)
			p.Match(preludioParserIDENT)
		}

	}
	p.SetState(117)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == preludioParserLANG {
		{
			p.SetState(116)
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
	if listenerT, ok := listener.(preludioListener); ok {
		listenerT.EnterTypeDef(s)
	}
}

func (s *TypeDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioListener); ok {
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
		p.SetState(119)
		p.Match(preludioParserLANG)
	}
	{
		p.SetState(120)
		p.TypeTerm()
	}
	{
		p.SetState(121)
		p.Match(preludioParserBAR)
	}
	p.SetState(125)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == preludioParserIDENT {
		{
			p.SetState(122)
			p.TypeTerm()
		}

		p.SetState(127)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(128)
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
	if listenerT, ok := listener.(preludioListener); ok {
		listenerT.EnterTypeTerm(s)
	}
}

func (s *TypeTermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioListener); ok {
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
		p.SetState(130)
		p.Match(preludioParserIDENT)
	}
	p.SetState(132)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == preludioParserLANG {
		{
			p.SetState(131)
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

func (s *StmtContext) AssignStmt() IAssignStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignStmtContext)
}

func (s *StmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioListener); ok {
		listenerT.EnterStmt(s)
	}
}

func (s *StmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioListener); ok {
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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(134)
		p.AssignStmt()
	}

	return localctx
}

// IAssignStmtContext is an interface to support dynamic dispatch.
type IAssignStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignStmtContext differentiates from other interfaces.
	IsAssignStmtContext()
}

type AssignStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignStmtContext() *AssignStmtContext {
	var p = new(AssignStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = preludioParserRULE_assignStmt
	return p
}

func (*AssignStmtContext) IsAssignStmtContext() {}

func NewAssignStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignStmtContext {
	var p = new(AssignStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = preludioParserRULE_assignStmt

	return p
}

func (s *AssignStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignStmtContext) LET() antlr.TerminalNode {
	return s.GetToken(preludioParserLET, 0)
}

func (s *AssignStmtContext) IDENT() antlr.TerminalNode {
	return s.GetToken(preludioParserIDENT, 0)
}

func (s *AssignStmtContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(preludioParserASSIGN, 0)
}

func (s *AssignStmtContext) Expr() IExprContext {
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

func (s *AssignStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioListener); ok {
		listenerT.EnterAssignStmt(s)
	}
}

func (s *AssignStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioListener); ok {
		listenerT.ExitAssignStmt(s)
	}
}

func (p *preludioParser) AssignStmt() (localctx IAssignStmtContext) {
	this := p
	_ = this

	localctx = NewAssignStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, preludioParserRULE_assignStmt)

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
		p.SetState(136)
		p.Match(preludioParserLET)
	}
	{
		p.SetState(137)
		p.Match(preludioParserIDENT)
	}
	{
		p.SetState(138)
		p.Match(preludioParserASSIGN)
	}
	{
		p.SetState(139)
		p.expr(0)
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
	if listenerT, ok := listener.(preludioListener); ok {
		listenerT.EnterPipeline(s)
	}
}

func (s *PipelineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioListener); ok {
		listenerT.ExitPipeline(s)
	}
}

func (p *preludioParser) Pipeline() (localctx IPipelineContext) {
	this := p
	_ = this

	localctx = NewPipelineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, preludioParserRULE_pipeline)

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
		p.SetState(141)
		p.ExprCall()
	}
	p.SetState(147)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(142)
				p.Nl()
			}
			{
				p.SetState(143)
				p.FuncCall()
			}

		}
		p.SetState(149)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())
	}
	p.SetState(152)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case preludioParserNEWLINE, preludioParserCOMMENT:
		{
			p.SetState(150)
			p.Nl()
		}

	case preludioParserEOF:
		{
			p.SetState(151)
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
	if listenerT, ok := listener.(preludioListener); ok {
		listenerT.EnterInlinePipeline(s)
	}
}

func (s *InlinePipelineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioListener); ok {
		listenerT.ExitInlinePipeline(s)
	}
}

func (p *preludioParser) InlinePipeline() (localctx IInlinePipelineContext) {
	this := p
	_ = this

	localctx = NewInlinePipelineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, preludioParserRULE_inlinePipeline)
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
		p.SetState(154)
		p.ExprCall()
	}
	p.SetState(159)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == preludioParserBAR {
		{
			p.SetState(155)
			p.Match(preludioParserBAR)
		}
		{
			p.SetState(156)
			p.FuncCall()
		}

		p.SetState(161)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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
	if listenerT, ok := listener.(preludioListener); ok {
		listenerT.EnterIdentBacktick(s)
	}
}

func (s *IdentBacktickContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioListener); ok {
		listenerT.ExitIdentBacktick(s)
	}
}

func (p *preludioParser) IdentBacktick() (localctx IIdentBacktickContext) {
	this := p
	_ = this

	localctx = NewIdentBacktickContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, preludioParserRULE_identBacktick)
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
		p.SetState(162)
		p.Match(preludioParserBACKTICK)
	}
	p.SetState(166)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2111061251588094) != 0 {
		{
			p.SetState(163)
			_la = p.GetTokenStream().LA(1)

			if _la <= 0 || _la == preludioParserBACKTICK || _la == preludioParserNEWLINE {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(168)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(169)
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
	if listenerT, ok := listener.(preludioListener); ok {
		listenerT.EnterFuncCall(s)
	}
}

func (s *FuncCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioListener); ok {
		listenerT.ExitFuncCall(s)
	}
}

func (p *preludioParser) FuncCall() (localctx IFuncCallContext) {
	this := p
	_ = this

	localctx = NewFuncCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, preludioParserRULE_funcCall)

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
		p.SetState(171)
		p.Match(preludioParserIDENT)
	}
	p.SetState(175)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(172)
				p.FuncCallParam()
			}

		}
		p.SetState(177)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext())
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
	if listenerT, ok := listener.(preludioListener); ok {
		listenerT.EnterFuncCallParam(s)
	}
}

func (s *FuncCallParamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioListener); ok {
		listenerT.ExitFuncCallParam(s)
	}
}

func (p *preludioParser) FuncCallParam() (localctx IFuncCallParamContext) {
	this := p
	_ = this

	localctx = NewFuncCallParamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, preludioParserRULE_funcCallParam)

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

	p.SetState(182)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(178)
			p.NamedArg()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(179)
			p.Assign()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(180)
			p.MultiAssign()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(181)
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
	if listenerT, ok := listener.(preludioListener); ok {
		listenerT.EnterNamedArg(s)
	}
}

func (s *NamedArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioListener); ok {
		listenerT.ExitNamedArg(s)
	}
}

func (p *preludioParser) NamedArg() (localctx INamedArgContext) {
	this := p
	_ = this

	localctx = NewNamedArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, preludioParserRULE_namedArg)

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
		p.SetState(184)
		p.Match(preludioParserIDENT)
	}
	{
		p.SetState(185)
		p.Match(preludioParserCOLON)
	}
	p.SetState(188)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(186)
			p.Assign()
		}

	case 2:
		{
			p.SetState(187)
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
	if listenerT, ok := listener.(preludioListener); ok {
		listenerT.EnterAssign(s)
	}
}

func (s *AssignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioListener); ok {
		listenerT.ExitAssign(s)
	}
}

func (p *preludioParser) Assign() (localctx IAssignContext) {
	this := p
	_ = this

	localctx = NewAssignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, preludioParserRULE_assign)

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
		p.SetState(190)
		p.Match(preludioParserIDENT)
	}
	{
		p.SetState(191)
		p.Match(preludioParserASSIGN)
	}
	{
		p.SetState(192)
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
	if listenerT, ok := listener.(preludioListener); ok {
		listenerT.EnterMultiAssign(s)
	}
}

func (s *MultiAssignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioListener); ok {
		listenerT.ExitMultiAssign(s)
	}
}

func (p *preludioParser) MultiAssign() (localctx IMultiAssignContext) {
	this := p
	_ = this

	localctx = NewMultiAssignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, preludioParserRULE_multiAssign)

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
		p.SetState(194)
		p.List()
	}
	{
		p.SetState(195)
		p.Match(preludioParserASSIGN)
	}
	{
		p.SetState(196)
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
	if listenerT, ok := listener.(preludioListener); ok {
		listenerT.EnterExprCall(s)
	}
}

func (s *ExprCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioListener); ok {
		listenerT.ExitExprCall(s)
	}
}

func (p *preludioParser) ExprCall() (localctx IExprCallContext) {
	this := p
	_ = this

	localctx = NewExprCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, preludioParserRULE_exprCall)

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

	p.SetState(200)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(198)
			p.expr(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(199)
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

func (s *ExprContext) POW() antlr.TerminalNode {
	return s.GetToken(preludioParserPOW, 0)
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
	if listenerT, ok := listener.(preludioListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioListener); ok {
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
	_startState := 40
	p.EnterRecursionRule(localctx, 40, preludioParserRULE_expr, _p)
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
	p.SetState(208)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(203)
			p.Match(preludioParserLPAREN)
		}
		{
			p.SetState(204)
			p.expr(0)
		}
		{
			p.SetState(205)
			p.Match(preludioParserRPAREN)
		}

	case 2:
		{
			p.SetState(207)
			p.Term()
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(233)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(231)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, preludioParserRULE_expr)
				p.SetState(210)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(211)
					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3328) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(212)
					p.expr(10)
				}

			case 2:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, preludioParserRULE_expr)
				p.SetState(213)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(214)
					_la = p.GetTokenStream().LA(1)

					if !(_la == preludioParserPLUS || _la == preludioParserMINUS) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(215)
					p.expr(9)
				}

			case 3:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, preludioParserRULE_expr)
				p.SetState(216)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(217)
					p.Match(preludioParserPOW)
				}
				{
					p.SetState(218)
					p.expr(8)
				}

			case 4:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, preludioParserRULE_expr)
				p.SetState(219)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(220)
					p.Match(preludioParserMODEL)
				}
				{
					p.SetState(221)
					p.expr(7)
				}

			case 5:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, preludioParserRULE_expr)
				p.SetState(222)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(223)
					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&25288704) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(224)
					p.expr(6)
				}

			case 6:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, preludioParserRULE_expr)
				p.SetState(225)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(226)
					p.Match(preludioParserCOALESCE)
				}
				{
					p.SetState(227)
					p.expr(5)
				}

			case 7:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, preludioParserRULE_expr)
				p.SetState(228)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(229)
					_la = p.GetTokenStream().LA(1)

					if !(_la == preludioParserAND || _la == preludioParserOR) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(230)
					p.expr(4)
				}

			}

		}
		p.SetState(235)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext())
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
	if listenerT, ok := listener.(preludioListener); ok {
		listenerT.EnterTerm(s)
	}
}

func (s *TermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioListener); ok {
		listenerT.ExitTerm(s)
	}
}

func (p *preludioParser) Term() (localctx ITermContext) {
	this := p
	_ = this

	localctx = NewTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, preludioParserRULE_term)

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

	p.SetState(241)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case preludioParserNULL_, preludioParserBOOLEAN, preludioParserINTEGER, preludioParserFLOAT, preludioParserIDENT, preludioParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(236)
			p.Literal()
		}

	case preludioParserBACKTICK:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(237)
			p.IdentBacktick()
		}

	case preludioParserPLUS, preludioParserMINUS, preludioParserNOT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(238)
			p.ExprUnary()
		}

	case preludioParserLBRACKET:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(239)
			p.List()
		}

	case preludioParserLPAREN:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(240)
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

func (s *ExprUnaryContext) MINUS() antlr.TerminalNode {
	return s.GetToken(preludioParserMINUS, 0)
}

func (s *ExprUnaryContext) PLUS() antlr.TerminalNode {
	return s.GetToken(preludioParserPLUS, 0)
}

func (s *ExprUnaryContext) NOT() antlr.TerminalNode {
	return s.GetToken(preludioParserNOT, 0)
}

func (s *ExprUnaryContext) NestedPipeline() INestedPipelineContext {
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

func (s *ExprUnaryContext) Literal() ILiteralContext {
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

func (s *ExprUnaryContext) IDENT() antlr.TerminalNode {
	return s.GetToken(preludioParserIDENT, 0)
}

func (s *ExprUnaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprUnaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprUnaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioListener); ok {
		listenerT.EnterExprUnary(s)
	}
}

func (s *ExprUnaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioListener); ok {
		listenerT.ExitExprUnary(s)
	}
}

func (p *preludioParser) ExprUnary() (localctx IExprUnaryContext) {
	this := p
	_ = this

	localctx = NewExprUnaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, preludioParserRULE_exprUnary)
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
		p.SetState(243)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&137438953664) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(247)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(244)
			p.NestedPipeline()
		}

	case 2:
		{
			p.SetState(245)
			p.Literal()
		}

	case 3:
		{
			p.SetState(246)
			p.Match(preludioParserIDENT)
		}

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

func (s *LiteralContext) AllIDENT() []antlr.TerminalNode {
	return s.GetTokens(preludioParserIDENT)
}

func (s *LiteralContext) IDENT(i int) antlr.TerminalNode {
	return s.GetToken(preludioParserIDENT, i)
}

func (s *LiteralContext) NULL_() antlr.TerminalNode {
	return s.GetToken(preludioParserNULL_, 0)
}

func (s *LiteralContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(preludioParserBOOLEAN, 0)
}

func (s *LiteralContext) STRING() antlr.TerminalNode {
	return s.GetToken(preludioParserSTRING, 0)
}

func (s *LiteralContext) AllINTEGER() []antlr.TerminalNode {
	return s.GetTokens(preludioParserINTEGER)
}

func (s *LiteralContext) INTEGER(i int) antlr.TerminalNode {
	return s.GetToken(preludioParserINTEGER, i)
}

func (s *LiteralContext) AllFLOAT() []antlr.TerminalNode {
	return s.GetTokens(preludioParserFLOAT)
}

func (s *LiteralContext) FLOAT(i int) antlr.TerminalNode {
	return s.GetToken(preludioParserFLOAT, i)
}

func (s *LiteralContext) INTERVAL_KIND() antlr.TerminalNode {
	return s.GetToken(preludioParserINTERVAL_KIND, 0)
}

func (s *LiteralContext) RANGE() antlr.TerminalNode {
	return s.GetToken(preludioParserRANGE, 0)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioListener); ok {
		listenerT.EnterLiteral(s)
	}
}

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioListener); ok {
		listenerT.ExitLiteral(s)
	}
}

func (p *preludioParser) Literal() (localctx ILiteralContext) {
	this := p
	_ = this

	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, preludioParserRULE_literal)
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

	p.SetState(260)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(249)
			p.Match(preludioParserIDENT)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(250)
			p.Match(preludioParserNULL_)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(251)
			p.Match(preludioParserBOOLEAN)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(252)
			p.Match(preludioParserSTRING)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(253)
			p.Match(preludioParserINTEGER)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(254)
			p.Match(preludioParserFLOAT)
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(255)
			_la = p.GetTokenStream().LA(1)

			if !(_la == preludioParserINTEGER || _la == preludioParserFLOAT) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(256)
			p.Match(preludioParserINTERVAL_KIND)
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(257)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&15393162788864) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(258)
			p.Match(preludioParserRANGE)
		}
		{
			p.SetState(259)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&15393162788864) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
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
	if listenerT, ok := listener.(preludioListener); ok {
		listenerT.EnterList(s)
	}
}

func (s *ListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioListener); ok {
		listenerT.ExitList(s)
	}
}

func (p *preludioParser) List() (localctx IListContext) {
	this := p
	_ = this

	localctx = NewListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, preludioParserRULE_list)
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
		p.SetState(262)
		p.Match(preludioParserLBRACKET)
	}
	p.SetState(297)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1565293482606784) != 0 {
		p.SetState(266)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == preludioParserNEWLINE || _la == preludioParserCOMMENT {
			{
				p.SetState(263)
				p.Nl()
			}

			p.SetState(268)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(272)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(269)
				p.Assign()
			}

		case 2:
			{
				p.SetState(270)
				p.MultiAssign()
			}

		case 3:
			{
				p.SetState(271)
				p.ExprCall()
			}

		}
		p.SetState(288)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(274)
					p.Match(preludioParserCOMMA)
				}
				p.SetState(278)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				for _la == preludioParserNEWLINE || _la == preludioParserCOMMENT {
					{
						p.SetState(275)
						p.Nl()
					}

					p.SetState(280)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)
				}
				p.SetState(284)
				p.GetErrorHandler().Sync(p)
				switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext()) {
				case 1:
					{
						p.SetState(281)
						p.Assign()
					}

				case 2:
					{
						p.SetState(282)
						p.MultiAssign()
					}

				case 3:
					{
						p.SetState(283)
						p.ExprCall()
					}

				}

			}
			p.SetState(290)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext())
		}
		p.SetState(292)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == preludioParserCOMMA {
			{
				p.SetState(291)
				p.Match(preludioParserCOMMA)
			}

		}
		p.SetState(295)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == preludioParserNEWLINE || _la == preludioParserCOMMENT {
			{
				p.SetState(294)
				p.Nl()
			}

		}

	}
	{
		p.SetState(299)
		p.Match(preludioParserRBRACKET)
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
	if listenerT, ok := listener.(preludioListener); ok {
		listenerT.EnterNestedPipeline(s)
	}
}

func (s *NestedPipelineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(preludioListener); ok {
		listenerT.ExitNestedPipeline(s)
	}
}

func (p *preludioParser) NestedPipeline() (localctx INestedPipelineContext) {
	this := p
	_ = this

	localctx = NewNestedPipelineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, preludioParserRULE_nestedPipeline)
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
		p.SetState(301)
		p.Match(preludioParserLPAREN)
	}
	p.SetState(305)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == preludioParserNEWLINE || _la == preludioParserCOMMENT {
		{
			p.SetState(302)
			p.Nl()
		}

		p.SetState(307)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(310)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 36, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(308)
			p.Pipeline()
		}

	case 2:
		{
			p.SetState(309)
			p.InlinePipeline()
		}

	}
	p.SetState(315)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == preludioParserNEWLINE || _la == preludioParserCOMMENT {
		{
			p.SetState(312)
			p.Nl()
		}

		p.SetState(317)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(318)
		p.Match(preludioParserRPAREN)
	}

	return localctx
}

func (p *preludioParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 20:
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
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
