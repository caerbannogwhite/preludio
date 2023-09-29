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
		"", "'func'", "'prql'", "'let'", "'ret'", "'->'", "'='", "'+'", "'-'",
		"'*'", "'**'", "'/'", "'%'", "'~'", "'=='", "'!='", "'<='", "'>='",
		"'@'", "'|'", "':'", "','", "'.'", "'$'", "'..'", "'<'", "'>'", "'['",
		"']'", "'('", "')'", "'_'", "'`'", "'\"'", "'''", "'\"\"\"'", "'''''",
		"'and'", "'or'", "'not'", "'??'", "'NA'",
	}
	staticData.symbolicNames = []string{
		"", "FUNC", "PRQL", "LET", "RET", "ARROW", "ASSIGN", "PLUS", "MINUS",
		"STAR", "POW", "DIV", "MOD", "MODEL", "EQ", "NE", "LE", "GE", "AT",
		"BAR", "COLON", "COMMA", "DOT", "DOLLAR", "RANGE", "LANG", "RANG", "LBRACKET",
		"RBRACKET", "LPAREN", "RPAREN", "UNDERSCORE", "BACKTICK", "DOUBLE_QUOTE",
		"SINGLE_QUOTE", "TRIPLE_DOUBLE_QUOTE", "TRIPLE_SINGLE_QUOTE", "AND",
		"OR", "NOT", "COALESCE", "NULL_", "BOOLEAN", "IDENT", "IDENT_START",
		"IDENT_NEXT", "WHITESPACE", "NEWLINE", "SINGLE_LINE_COMMENT", "INTEGER",
		"FLOAT", "STRING", "STRING_RAW", "REGEXP_LITERAL", "RANGE_LITERAL",
		"DATE_LITERAL", "DURATION_KIND", "DURATION_LITERAL",
	}
	staticData.ruleNames = []string{
		"nl", "program", "programIntro", "funcDef", "funcDefName", "funcDefParams",
		"funcDefParam", "typeDef", "typeTerm", "stmt", "varAssignStmt", "varDeclStmt",
		"retStmt", "pipeline", "inlinePipeline", "identBacktick", "funcCall",
		"funcCallParam", "namedArg", "assign", "multiAssign", "exprCall", "expr",
		"term", "exprUnary", "literal", "list", "nestedPipeline",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 57, 330, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 1, 0, 1, 0, 1, 1, 5, 1, 60, 8, 1, 10, 1, 12, 1, 63,
		9, 1, 1, 1, 3, 1, 66, 8, 1, 1, 1, 5, 1, 69, 8, 1, 10, 1, 12, 1, 72, 9,
		1, 1, 1, 1, 1, 1, 1, 3, 1, 77, 8, 1, 1, 1, 5, 1, 80, 8, 1, 10, 1, 12, 1,
		83, 9, 1, 5, 1, 85, 8, 1, 10, 1, 12, 1, 88, 9, 1, 1, 1, 1, 1, 1, 2, 1,
		2, 5, 2, 94, 8, 2, 10, 2, 12, 2, 97, 9, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 3, 4, 109, 8, 4, 1, 5, 5, 5, 112, 8, 5, 10,
		5, 12, 5, 115, 9, 5, 1, 6, 1, 6, 3, 6, 119, 8, 6, 1, 6, 3, 6, 122, 8, 6,
		1, 7, 1, 7, 1, 7, 1, 7, 5, 7, 128, 8, 7, 10, 7, 12, 7, 131, 9, 7, 1, 7,
		1, 7, 1, 8, 1, 8, 3, 8, 137, 8, 8, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 143, 8,
		9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12,
		1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 5, 13, 161, 8, 13, 10, 13, 12,
		13, 164, 9, 13, 1, 13, 1, 13, 3, 13, 168, 8, 13, 1, 14, 1, 14, 1, 14, 5,
		14, 173, 8, 14, 10, 14, 12, 14, 176, 9, 14, 1, 15, 1, 15, 5, 15, 180, 8,
		15, 10, 15, 12, 15, 183, 9, 15, 1, 15, 1, 15, 1, 16, 1, 16, 5, 16, 189,
		8, 16, 10, 16, 12, 16, 192, 9, 16, 1, 17, 1, 17, 1, 17, 1, 17, 3, 17, 198,
		8, 17, 1, 18, 1, 18, 1, 18, 1, 18, 3, 18, 204, 8, 18, 1, 19, 1, 19, 1,
		19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 3, 21, 216, 8, 21,
		1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 3, 22, 224, 8, 22, 1, 22, 1,
		22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22,
		1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1,
		22, 1, 22, 1, 22, 1, 22, 5, 22, 252, 8, 22, 10, 22, 12, 22, 255, 9, 22,
		1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 3, 23, 262, 8, 23, 1, 24, 1, 24, 1,
		24, 1, 24, 3, 24, 268, 8, 24, 1, 25, 1, 25, 1, 26, 1, 26, 5, 26, 274, 8,
		26, 10, 26, 12, 26, 277, 9, 26, 1, 26, 1, 26, 1, 26, 3, 26, 282, 8, 26,
		1, 26, 1, 26, 5, 26, 286, 8, 26, 10, 26, 12, 26, 289, 9, 26, 1, 26, 1,
		26, 1, 26, 3, 26, 294, 8, 26, 5, 26, 296, 8, 26, 10, 26, 12, 26, 299, 9,
		26, 1, 26, 3, 26, 302, 8, 26, 1, 26, 3, 26, 305, 8, 26, 3, 26, 307, 8,
		26, 1, 26, 1, 26, 1, 27, 1, 27, 5, 27, 313, 8, 27, 10, 27, 12, 27, 316,
		9, 27, 1, 27, 1, 27, 3, 27, 320, 8, 27, 1, 27, 5, 27, 323, 8, 27, 10, 27,
		12, 27, 326, 9, 27, 1, 27, 1, 27, 1, 27, 0, 1, 44, 28, 0, 2, 4, 6, 8, 10,
		12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46,
		48, 50, 52, 54, 0, 7, 2, 0, 32, 32, 47, 47, 2, 0, 9, 9, 11, 12, 1, 0, 7,
		8, 2, 0, 14, 17, 25, 26, 1, 0, 37, 38, 2, 0, 7, 8, 39, 39, 3, 0, 41, 43,
		49, 55, 57, 57, 356, 0, 56, 1, 0, 0, 0, 2, 61, 1, 0, 0, 0, 4, 91, 1, 0,
		0, 0, 6, 100, 1, 0, 0, 0, 8, 106, 1, 0, 0, 0, 10, 113, 1, 0, 0, 0, 12,
		118, 1, 0, 0, 0, 14, 123, 1, 0, 0, 0, 16, 134, 1, 0, 0, 0, 18, 142, 1,
		0, 0, 0, 20, 144, 1, 0, 0, 0, 22, 148, 1, 0, 0, 0, 24, 153, 1, 0, 0, 0,
		26, 156, 1, 0, 0, 0, 28, 169, 1, 0, 0, 0, 30, 177, 1, 0, 0, 0, 32, 186,
		1, 0, 0, 0, 34, 197, 1, 0, 0, 0, 36, 199, 1, 0, 0, 0, 38, 205, 1, 0, 0,
		0, 40, 209, 1, 0, 0, 0, 42, 215, 1, 0, 0, 0, 44, 223, 1, 0, 0, 0, 46, 261,
		1, 0, 0, 0, 48, 263, 1, 0, 0, 0, 50, 269, 1, 0, 0, 0, 52, 271, 1, 0, 0,
		0, 54, 310, 1, 0, 0, 0, 56, 57, 5, 47, 0, 0, 57, 1, 1, 0, 0, 0, 58, 60,
		3, 0, 0, 0, 59, 58, 1, 0, 0, 0, 60, 63, 1, 0, 0, 0, 61, 59, 1, 0, 0, 0,
		61, 62, 1, 0, 0, 0, 62, 65, 1, 0, 0, 0, 63, 61, 1, 0, 0, 0, 64, 66, 3,
		4, 2, 0, 65, 64, 1, 0, 0, 0, 65, 66, 1, 0, 0, 0, 66, 70, 1, 0, 0, 0, 67,
		69, 3, 0, 0, 0, 68, 67, 1, 0, 0, 0, 69, 72, 1, 0, 0, 0, 70, 68, 1, 0, 0,
		0, 70, 71, 1, 0, 0, 0, 71, 86, 1, 0, 0, 0, 72, 70, 1, 0, 0, 0, 73, 77,
		3, 6, 3, 0, 74, 77, 3, 18, 9, 0, 75, 77, 3, 54, 27, 0, 76, 73, 1, 0, 0,
		0, 76, 74, 1, 0, 0, 0, 76, 75, 1, 0, 0, 0, 77, 81, 1, 0, 0, 0, 78, 80,
		3, 0, 0, 0, 79, 78, 1, 0, 0, 0, 80, 83, 1, 0, 0, 0, 81, 79, 1, 0, 0, 0,
		81, 82, 1, 0, 0, 0, 82, 85, 1, 0, 0, 0, 83, 81, 1, 0, 0, 0, 84, 76, 1,
		0, 0, 0, 85, 88, 1, 0, 0, 0, 86, 84, 1, 0, 0, 0, 86, 87, 1, 0, 0, 0, 87,
		89, 1, 0, 0, 0, 88, 86, 1, 0, 0, 0, 89, 90, 5, 0, 0, 1, 90, 3, 1, 0, 0,
		0, 91, 95, 5, 2, 0, 0, 92, 94, 3, 36, 18, 0, 93, 92, 1, 0, 0, 0, 94, 97,
		1, 0, 0, 0, 95, 93, 1, 0, 0, 0, 95, 96, 1, 0, 0, 0, 96, 98, 1, 0, 0, 0,
		97, 95, 1, 0, 0, 0, 98, 99, 3, 0, 0, 0, 99, 5, 1, 0, 0, 0, 100, 101, 5,
		1, 0, 0, 101, 102, 3, 8, 4, 0, 102, 103, 3, 10, 5, 0, 103, 104, 5, 5, 0,
		0, 104, 105, 3, 44, 22, 0, 105, 7, 1, 0, 0, 0, 106, 108, 5, 43, 0, 0, 107,
		109, 3, 14, 7, 0, 108, 107, 1, 0, 0, 0, 108, 109, 1, 0, 0, 0, 109, 9, 1,
		0, 0, 0, 110, 112, 3, 12, 6, 0, 111, 110, 1, 0, 0, 0, 112, 115, 1, 0, 0,
		0, 113, 111, 1, 0, 0, 0, 113, 114, 1, 0, 0, 0, 114, 11, 1, 0, 0, 0, 115,
		113, 1, 0, 0, 0, 116, 119, 3, 36, 18, 0, 117, 119, 5, 43, 0, 0, 118, 116,
		1, 0, 0, 0, 118, 117, 1, 0, 0, 0, 119, 121, 1, 0, 0, 0, 120, 122, 3, 14,
		7, 0, 121, 120, 1, 0, 0, 0, 121, 122, 1, 0, 0, 0, 122, 13, 1, 0, 0, 0,
		123, 124, 5, 25, 0, 0, 124, 125, 3, 16, 8, 0, 125, 129, 5, 19, 0, 0, 126,
		128, 3, 16, 8, 0, 127, 126, 1, 0, 0, 0, 128, 131, 1, 0, 0, 0, 129, 127,
		1, 0, 0, 0, 129, 130, 1, 0, 0, 0, 130, 132, 1, 0, 0, 0, 131, 129, 1, 0,
		0, 0, 132, 133, 5, 26, 0, 0, 133, 15, 1, 0, 0, 0, 134, 136, 5, 43, 0, 0,
		135, 137, 3, 14, 7, 0, 136, 135, 1, 0, 0, 0, 136, 137, 1, 0, 0, 0, 137,
		17, 1, 0, 0, 0, 138, 143, 3, 20, 10, 0, 139, 143, 3, 22, 11, 0, 140, 143,
		3, 24, 12, 0, 141, 143, 3, 44, 22, 0, 142, 138, 1, 0, 0, 0, 142, 139, 1,
		0, 0, 0, 142, 140, 1, 0, 0, 0, 142, 141, 1, 0, 0, 0, 143, 19, 1, 0, 0,
		0, 144, 145, 5, 43, 0, 0, 145, 146, 5, 6, 0, 0, 146, 147, 3, 44, 22, 0,
		147, 21, 1, 0, 0, 0, 148, 149, 5, 3, 0, 0, 149, 150, 5, 43, 0, 0, 150,
		151, 5, 6, 0, 0, 151, 152, 3, 44, 22, 0, 152, 23, 1, 0, 0, 0, 153, 154,
		5, 4, 0, 0, 154, 155, 3, 44, 22, 0, 155, 25, 1, 0, 0, 0, 156, 162, 3, 42,
		21, 0, 157, 158, 3, 0, 0, 0, 158, 159, 3, 32, 16, 0, 159, 161, 1, 0, 0,
		0, 160, 157, 1, 0, 0, 0, 161, 164, 1, 0, 0, 0, 162, 160, 1, 0, 0, 0, 162,
		163, 1, 0, 0, 0, 163, 167, 1, 0, 0, 0, 164, 162, 1, 0, 0, 0, 165, 168,
		3, 0, 0, 0, 166, 168, 5, 0, 0, 1, 167, 165, 1, 0, 0, 0, 167, 166, 1, 0,
		0, 0, 168, 27, 1, 0, 0, 0, 169, 174, 3, 42, 21, 0, 170, 171, 5, 19, 0,
		0, 171, 173, 3, 32, 16, 0, 172, 170, 1, 0, 0, 0, 173, 176, 1, 0, 0, 0,
		174, 172, 1, 0, 0, 0, 174, 175, 1, 0, 0, 0, 175, 29, 1, 0, 0, 0, 176, 174,
		1, 0, 0, 0, 177, 181, 5, 32, 0, 0, 178, 180, 8, 0, 0, 0, 179, 178, 1, 0,
		0, 0, 180, 183, 1, 0, 0, 0, 181, 179, 1, 0, 0, 0, 181, 182, 1, 0, 0, 0,
		182, 184, 1, 0, 0, 0, 183, 181, 1, 0, 0, 0, 184, 185, 5, 32, 0, 0, 185,
		31, 1, 0, 0, 0, 186, 190, 5, 43, 0, 0, 187, 189, 3, 34, 17, 0, 188, 187,
		1, 0, 0, 0, 189, 192, 1, 0, 0, 0, 190, 188, 1, 0, 0, 0, 190, 191, 1, 0,
		0, 0, 191, 33, 1, 0, 0, 0, 192, 190, 1, 0, 0, 0, 193, 198, 3, 36, 18, 0,
		194, 198, 3, 38, 19, 0, 195, 198, 3, 40, 20, 0, 196, 198, 3, 44, 22, 0,
		197, 193, 1, 0, 0, 0, 197, 194, 1, 0, 0, 0, 197, 195, 1, 0, 0, 0, 197,
		196, 1, 0, 0, 0, 198, 35, 1, 0, 0, 0, 199, 200, 5, 43, 0, 0, 200, 203,
		5, 20, 0, 0, 201, 204, 3, 38, 19, 0, 202, 204, 3, 44, 22, 0, 203, 201,
		1, 0, 0, 0, 203, 202, 1, 0, 0, 0, 204, 37, 1, 0, 0, 0, 205, 206, 5, 43,
		0, 0, 206, 207, 5, 6, 0, 0, 207, 208, 3, 42, 21, 0, 208, 39, 1, 0, 0, 0,
		209, 210, 3, 52, 26, 0, 210, 211, 5, 6, 0, 0, 211, 212, 3, 42, 21, 0, 212,
		41, 1, 0, 0, 0, 213, 216, 3, 44, 22, 0, 214, 216, 3, 32, 16, 0, 215, 213,
		1, 0, 0, 0, 215, 214, 1, 0, 0, 0, 216, 43, 1, 0, 0, 0, 217, 218, 6, 22,
		-1, 0, 218, 219, 5, 29, 0, 0, 219, 220, 3, 44, 22, 0, 220, 221, 5, 30,
		0, 0, 221, 224, 1, 0, 0, 0, 222, 224, 3, 46, 23, 0, 223, 217, 1, 0, 0,
		0, 223, 222, 1, 0, 0, 0, 224, 253, 1, 0, 0, 0, 225, 226, 10, 9, 0, 0, 226,
		227, 7, 1, 0, 0, 227, 252, 3, 44, 22, 10, 228, 229, 10, 8, 0, 0, 229, 230,
		7, 2, 0, 0, 230, 252, 3, 44, 22, 9, 231, 232, 10, 7, 0, 0, 232, 233, 5,
		10, 0, 0, 233, 252, 3, 44, 22, 8, 234, 235, 10, 6, 0, 0, 235, 236, 5, 13,
		0, 0, 236, 252, 3, 44, 22, 7, 237, 238, 10, 5, 0, 0, 238, 239, 7, 3, 0,
		0, 239, 252, 3, 44, 22, 6, 240, 241, 10, 4, 0, 0, 241, 242, 5, 40, 0, 0,
		242, 252, 3, 44, 22, 5, 243, 244, 10, 3, 0, 0, 244, 245, 7, 4, 0, 0, 245,
		252, 3, 44, 22, 4, 246, 247, 10, 10, 0, 0, 247, 248, 5, 27, 0, 0, 248,
		249, 3, 44, 22, 0, 249, 250, 5, 28, 0, 0, 250, 252, 1, 0, 0, 0, 251, 225,
		1, 0, 0, 0, 251, 228, 1, 0, 0, 0, 251, 231, 1, 0, 0, 0, 251, 234, 1, 0,
		0, 0, 251, 237, 1, 0, 0, 0, 251, 240, 1, 0, 0, 0, 251, 243, 1, 0, 0, 0,
		251, 246, 1, 0, 0, 0, 252, 255, 1, 0, 0, 0, 253, 251, 1, 0, 0, 0, 253,
		254, 1, 0, 0, 0, 254, 45, 1, 0, 0, 0, 255, 253, 1, 0, 0, 0, 256, 262, 3,
		50, 25, 0, 257, 262, 3, 30, 15, 0, 258, 262, 3, 48, 24, 0, 259, 262, 3,
		52, 26, 0, 260, 262, 3, 54, 27, 0, 261, 256, 1, 0, 0, 0, 261, 257, 1, 0,
		0, 0, 261, 258, 1, 0, 0, 0, 261, 259, 1, 0, 0, 0, 261, 260, 1, 0, 0, 0,
		262, 47, 1, 0, 0, 0, 263, 267, 7, 5, 0, 0, 264, 268, 3, 54, 27, 0, 265,
		268, 3, 50, 25, 0, 266, 268, 5, 43, 0, 0, 267, 264, 1, 0, 0, 0, 267, 265,
		1, 0, 0, 0, 267, 266, 1, 0, 0, 0, 268, 49, 1, 0, 0, 0, 269, 270, 7, 6,
		0, 0, 270, 51, 1, 0, 0, 0, 271, 306, 5, 27, 0, 0, 272, 274, 3, 0, 0, 0,
		273, 272, 1, 0, 0, 0, 274, 277, 1, 0, 0, 0, 275, 273, 1, 0, 0, 0, 275,
		276, 1, 0, 0, 0, 276, 281, 1, 0, 0, 0, 277, 275, 1, 0, 0, 0, 278, 282,
		3, 38, 19, 0, 279, 282, 3, 40, 20, 0, 280, 282, 3, 42, 21, 0, 281, 278,
		1, 0, 0, 0, 281, 279, 1, 0, 0, 0, 281, 280, 1, 0, 0, 0, 282, 297, 1, 0,
		0, 0, 283, 287, 5, 21, 0, 0, 284, 286, 3, 0, 0, 0, 285, 284, 1, 0, 0, 0,
		286, 289, 1, 0, 0, 0, 287, 285, 1, 0, 0, 0, 287, 288, 1, 0, 0, 0, 288,
		293, 1, 0, 0, 0, 289, 287, 1, 0, 0, 0, 290, 294, 3, 38, 19, 0, 291, 294,
		3, 40, 20, 0, 292, 294, 3, 42, 21, 0, 293, 290, 1, 0, 0, 0, 293, 291, 1,
		0, 0, 0, 293, 292, 1, 0, 0, 0, 294, 296, 1, 0, 0, 0, 295, 283, 1, 0, 0,
		0, 296, 299, 1, 0, 0, 0, 297, 295, 1, 0, 0, 0, 297, 298, 1, 0, 0, 0, 298,
		301, 1, 0, 0, 0, 299, 297, 1, 0, 0, 0, 300, 302, 5, 21, 0, 0, 301, 300,
		1, 0, 0, 0, 301, 302, 1, 0, 0, 0, 302, 304, 1, 0, 0, 0, 303, 305, 3, 0,
		0, 0, 304, 303, 1, 0, 0, 0, 304, 305, 1, 0, 0, 0, 305, 307, 1, 0, 0, 0,
		306, 275, 1, 0, 0, 0, 306, 307, 1, 0, 0, 0, 307, 308, 1, 0, 0, 0, 308,
		309, 5, 28, 0, 0, 309, 53, 1, 0, 0, 0, 310, 314, 5, 29, 0, 0, 311, 313,
		3, 0, 0, 0, 312, 311, 1, 0, 0, 0, 313, 316, 1, 0, 0, 0, 314, 312, 1, 0,
		0, 0, 314, 315, 1, 0, 0, 0, 315, 319, 1, 0, 0, 0, 316, 314, 1, 0, 0, 0,
		317, 320, 3, 26, 13, 0, 318, 320, 3, 28, 14, 0, 319, 317, 1, 0, 0, 0, 319,
		318, 1, 0, 0, 0, 320, 324, 1, 0, 0, 0, 321, 323, 3, 0, 0, 0, 322, 321,
		1, 0, 0, 0, 323, 326, 1, 0, 0, 0, 324, 322, 1, 0, 0, 0, 324, 325, 1, 0,
		0, 0, 325, 327, 1, 0, 0, 0, 326, 324, 1, 0, 0, 0, 327, 328, 5, 30, 0, 0,
		328, 55, 1, 0, 0, 0, 38, 61, 65, 70, 76, 81, 86, 95, 108, 113, 118, 121,
		129, 136, 142, 162, 167, 174, 181, 190, 197, 203, 215, 223, 251, 253, 261,
		267, 275, 281, 287, 293, 297, 301, 304, 306, 314, 319, 324,
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
	preludioParserLET                 = 3
	preludioParserRET                 = 4
	preludioParserARROW               = 5
	preludioParserASSIGN              = 6
	preludioParserPLUS                = 7
	preludioParserMINUS               = 8
	preludioParserSTAR                = 9
	preludioParserPOW                 = 10
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
	preludioParserUNDERSCORE          = 31
	preludioParserBACKTICK            = 32
	preludioParserDOUBLE_QUOTE        = 33
	preludioParserSINGLE_QUOTE        = 34
	preludioParserTRIPLE_DOUBLE_QUOTE = 35
	preludioParserTRIPLE_SINGLE_QUOTE = 36
	preludioParserAND                 = 37
	preludioParserOR                  = 38
	preludioParserNOT                 = 39
	preludioParserCOALESCE            = 40
	preludioParserNULL_               = 41
	preludioParserBOOLEAN             = 42
	preludioParserIDENT               = 43
	preludioParserIDENT_START         = 44
	preludioParserIDENT_NEXT          = 45
	preludioParserWHITESPACE          = 46
	preludioParserNEWLINE             = 47
	preludioParserSINGLE_LINE_COMMENT = 48
	preludioParserINTEGER             = 49
	preludioParserFLOAT               = 50
	preludioParserSTRING              = 51
	preludioParserSTRING_RAW          = 52
	preludioParserREGEXP_LITERAL      = 53
	preludioParserRANGE_LITERAL       = 54
	preludioParserDATE_LITERAL        = 55
	preludioParserDURATION_KIND       = 56
	preludioParserDURATION_LITERAL    = 57
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
	preludioParserRULE_identBacktick  = 15
	preludioParserRULE_funcCall       = 16
	preludioParserRULE_funcCallParam  = 17
	preludioParserRULE_namedArg       = 18
	preludioParserRULE_assign         = 19
	preludioParserRULE_multiAssign    = 20
	preludioParserRULE_exprCall       = 21
	preludioParserRULE_expr           = 22
	preludioParserRULE_term           = 23
	preludioParserRULE_exprUnary      = 24
	preludioParserRULE_literal        = 25
	preludioParserRULE_list           = 26
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

func (s *ProgramContext) AllNestedPipeline() []INestedPipelineContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INestedPipelineContext); ok {
			len++
		}
	}

	tst := make([]INestedPipelineContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INestedPipelineContext); ok {
			tst[i] = t.(INestedPipelineContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) NestedPipeline(i int) INestedPipelineContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INestedPipelineContext); ok {
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

	return t.(INestedPipelineContext)
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
	p.SetState(86)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&215625780045021594) != 0 {
		p.SetState(76)
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
				p.Stmt()
			}

		case 3:
			{
				p.SetState(75)
				p.NestedPipeline()
			}

		}
		p.SetState(81)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == preludioParserNEWLINE {
			{
				p.SetState(78)
				p.Nl()
			}

			p.SetState(83)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

		p.SetState(88)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(89)
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
		p.SetState(91)
		p.Match(preludioParserPRQL)
	}
	p.SetState(95)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == preludioParserIDENT {
		{
			p.SetState(92)
			p.NamedArg()
		}

		p.SetState(97)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(98)
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
		p.SetState(100)
		p.Match(preludioParserFUNC)
	}
	{
		p.SetState(101)
		p.FuncDefName()
	}
	{
		p.SetState(102)
		p.FuncDefParams()
	}
	{
		p.SetState(103)
		p.Match(preludioParserARROW)
	}
	{
		p.SetState(104)
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
		p.SetState(106)
		p.Match(preludioParserIDENT)
	}
	p.SetState(108)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == preludioParserLANG {
		{
			p.SetState(107)
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
	p.SetState(113)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == preludioParserIDENT {
		{
			p.SetState(110)
			p.FuncDefParam()
		}

		p.SetState(115)
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
	p.SetState(118)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(116)
			p.NamedArg()
		}

	case 2:
		{
			p.SetState(117)
			p.Match(preludioParserIDENT)
		}

	}
	p.SetState(121)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == preludioParserLANG {
		{
			p.SetState(120)
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
		p.SetState(123)
		p.Match(preludioParserLANG)
	}
	{
		p.SetState(124)
		p.TypeTerm()
	}
	{
		p.SetState(125)
		p.Match(preludioParserBAR)
	}
	p.SetState(129)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == preludioParserIDENT {
		{
			p.SetState(126)
			p.TypeTerm()
		}

		p.SetState(131)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(132)
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
		p.SetState(134)
		p.Match(preludioParserIDENT)
	}
	p.SetState(136)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == preludioParserLANG {
		{
			p.SetState(135)
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

	p.SetState(142)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(138)
			p.VarAssignStmt()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(139)
			p.VarDeclStmt()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(140)
			p.RetStmt()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(141)
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

func (s *VarAssignStmtContext) Expr() IExprContext {
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
		p.SetState(144)
		p.Match(preludioParserIDENT)
	}
	{
		p.SetState(145)
		p.Match(preludioParserASSIGN)
	}
	{
		p.SetState(146)
		p.expr(0)
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

func (s *VarDeclStmtContext) LET() antlr.TerminalNode {
	return s.GetToken(preludioParserLET, 0)
}

func (s *VarDeclStmtContext) IDENT() antlr.TerminalNode {
	return s.GetToken(preludioParserIDENT, 0)
}

func (s *VarDeclStmtContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(preludioParserASSIGN, 0)
}

func (s *VarDeclStmtContext) Expr() IExprContext {
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
		p.SetState(148)
		p.Match(preludioParserLET)
	}
	{
		p.SetState(149)
		p.Match(preludioParserIDENT)
	}
	{
		p.SetState(150)
		p.Match(preludioParserASSIGN)
	}
	{
		p.SetState(151)
		p.expr(0)
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

func (s *RetStmtContext) Expr() IExprContext {
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
	p.EnterRule(localctx, 30, preludioParserRULE_identBacktick)
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
		p.Match(preludioParserBACKTICK)
	}
	p.SetState(181)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&288089634368389118) != 0 {
		{
			p.SetState(178)
			_la = p.GetTokenStream().LA(1)

			if _la <= 0 || _la == preludioParserBACKTICK || _la == preludioParserNEWLINE {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(183)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(184)
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
	p.EnterRule(localctx, 32, preludioParserRULE_funcCall)

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
		p.SetState(186)
		p.Match(preludioParserIDENT)
	}
	p.SetState(190)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(187)
				p.FuncCallParam()
			}

		}
		p.SetState(192)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 34, preludioParserRULE_funcCallParam)

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

	p.SetState(197)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(193)
			p.NamedArg()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(194)
			p.Assign()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(195)
			p.MultiAssign()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(196)
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
	p.EnterRule(localctx, 36, preludioParserRULE_namedArg)

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
		p.SetState(199)
		p.Match(preludioParserIDENT)
	}
	{
		p.SetState(200)
		p.Match(preludioParserCOLON)
	}
	p.SetState(203)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(201)
			p.Assign()
		}

	case 2:
		{
			p.SetState(202)
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
	p.EnterRule(localctx, 38, preludioParserRULE_assign)

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
		p.SetState(205)
		p.Match(preludioParserIDENT)
	}
	{
		p.SetState(206)
		p.Match(preludioParserASSIGN)
	}
	{
		p.SetState(207)
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
	p.EnterRule(localctx, 40, preludioParserRULE_multiAssign)

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
		p.SetState(209)
		p.List()
	}
	{
		p.SetState(210)
		p.Match(preludioParserASSIGN)
	}
	{
		p.SetState(211)
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
	p.EnterRule(localctx, 42, preludioParserRULE_exprCall)

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

	p.SetState(215)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(213)
			p.expr(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(214)
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

func (s *ExprContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(preludioParserLBRACKET, 0)
}

func (s *ExprContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(preludioParserRBRACKET, 0)
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
	_startState := 44
	p.EnterRecursionRule(localctx, 44, preludioParserRULE_expr, _p)
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
	p.SetState(223)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(218)
			p.Match(preludioParserLPAREN)
		}
		{
			p.SetState(219)
			p.expr(0)
		}
		{
			p.SetState(220)
			p.Match(preludioParserRPAREN)
		}

	case 2:
		{
			p.SetState(222)
			p.Term()
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(253)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(251)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, preludioParserRULE_expr)
				p.SetState(225)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(226)
					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&6656) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(227)
					p.expr(10)
				}

			case 2:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, preludioParserRULE_expr)
				p.SetState(228)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(229)
					_la = p.GetTokenStream().LA(1)

					if !(_la == preludioParserPLUS || _la == preludioParserMINUS) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(230)
					p.expr(9)
				}

			case 3:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, preludioParserRULE_expr)
				p.SetState(231)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(232)
					p.Match(preludioParserPOW)
				}
				{
					p.SetState(233)
					p.expr(8)
				}

			case 4:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, preludioParserRULE_expr)
				p.SetState(234)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(235)
					p.Match(preludioParserMODEL)
				}
				{
					p.SetState(236)
					p.expr(7)
				}

			case 5:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, preludioParserRULE_expr)
				p.SetState(237)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(238)
					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&100909056) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(239)
					p.expr(6)
				}

			case 6:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, preludioParserRULE_expr)
				p.SetState(240)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(241)
					p.Match(preludioParserCOALESCE)
				}
				{
					p.SetState(242)
					p.expr(5)
				}

			case 7:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, preludioParserRULE_expr)
				p.SetState(243)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(244)
					_la = p.GetTokenStream().LA(1)

					if !(_la == preludioParserAND || _la == preludioParserOR) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(245)
					p.expr(4)
				}

			case 8:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, preludioParserRULE_expr)
				p.SetState(246)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(247)
					p.Match(preludioParserLBRACKET)
				}
				{
					p.SetState(248)
					p.expr(0)
				}
				{
					p.SetState(249)
					p.Match(preludioParserRBRACKET)
				}

			}

		}
		p.SetState(255)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 46, preludioParserRULE_term)

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

	p.SetState(261)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case preludioParserNULL_, preludioParserBOOLEAN, preludioParserIDENT, preludioParserINTEGER, preludioParserFLOAT, preludioParserSTRING, preludioParserSTRING_RAW, preludioParserREGEXP_LITERAL, preludioParserRANGE_LITERAL, preludioParserDATE_LITERAL, preludioParserDURATION_LITERAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(256)
			p.Literal()
		}

	case preludioParserBACKTICK:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(257)
			p.IdentBacktick()
		}

	case preludioParserPLUS, preludioParserMINUS, preludioParserNOT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(258)
			p.ExprUnary()
		}

	case preludioParserLBRACKET:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(259)
			p.List()
		}

	case preludioParserLPAREN:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(260)
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
	p.EnterRule(localctx, 48, preludioParserRULE_exprUnary)
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
		p.SetState(263)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&549755814272) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(267)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(264)
			p.NestedPipeline()
		}

	case 2:
		{
			p.SetState(265)
			p.Literal()
		}

	case 3:
		{
			p.SetState(266)
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

func (s *LiteralContext) IDENT() antlr.TerminalNode {
	return s.GetToken(preludioParserIDENT, 0)
}

func (s *LiteralContext) NULL_() antlr.TerminalNode {
	return s.GetToken(preludioParserNULL_, 0)
}

func (s *LiteralContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(preludioParserBOOLEAN, 0)
}

func (s *LiteralContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(preludioParserINTEGER, 0)
}

func (s *LiteralContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(preludioParserFLOAT, 0)
}

func (s *LiteralContext) STRING() antlr.TerminalNode {
	return s.GetToken(preludioParserSTRING, 0)
}

func (s *LiteralContext) STRING_RAW() antlr.TerminalNode {
	return s.GetToken(preludioParserSTRING_RAW, 0)
}

func (s *LiteralContext) REGEXP_LITERAL() antlr.TerminalNode {
	return s.GetToken(preludioParserREGEXP_LITERAL, 0)
}

func (s *LiteralContext) RANGE_LITERAL() antlr.TerminalNode {
	return s.GetToken(preludioParserRANGE_LITERAL, 0)
}

func (s *LiteralContext) DATE_LITERAL() antlr.TerminalNode {
	return s.GetToken(preludioParserDATE_LITERAL, 0)
}

func (s *LiteralContext) DURATION_LITERAL() antlr.TerminalNode {
	return s.GetToken(preludioParserDURATION_LITERAL, 0)
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
	p.EnterRule(localctx, 50, preludioParserRULE_literal)
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
		p.SetState(269)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&215625225323151360) != 0) {
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
	p.EnterRule(localctx, 52, preludioParserRULE_list)
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
		p.SetState(271)
		p.Match(preludioParserLBRACKET)
	}
	p.SetState(306)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&215766517533376896) != 0 {
		p.SetState(275)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == preludioParserNEWLINE {
			{
				p.SetState(272)
				p.Nl()
			}

			p.SetState(277)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(281)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(278)
				p.Assign()
			}

		case 2:
			{
				p.SetState(279)
				p.MultiAssign()
			}

		case 3:
			{
				p.SetState(280)
				p.ExprCall()
			}

		}
		p.SetState(297)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(283)
					p.Match(preludioParserCOMMA)
				}
				p.SetState(287)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				for _la == preludioParserNEWLINE {
					{
						p.SetState(284)
						p.Nl()
					}

					p.SetState(289)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)
				}
				p.SetState(293)
				p.GetErrorHandler().Sync(p)
				switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext()) {
				case 1:
					{
						p.SetState(290)
						p.Assign()
					}

				case 2:
					{
						p.SetState(291)
						p.MultiAssign()
					}

				case 3:
					{
						p.SetState(292)
						p.ExprCall()
					}

				}

			}
			p.SetState(299)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext())
		}
		p.SetState(301)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == preludioParserCOMMA {
			{
				p.SetState(300)
				p.Match(preludioParserCOMMA)
			}

		}
		p.SetState(304)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == preludioParserNEWLINE {
			{
				p.SetState(303)
				p.Nl()
			}

		}

	}
	{
		p.SetState(308)
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
		p.SetState(310)
		p.Match(preludioParserLPAREN)
	}
	p.SetState(314)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == preludioParserNEWLINE {
		{
			p.SetState(311)
			p.Nl()
		}

		p.SetState(316)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(319)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 36, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(317)
			p.Pipeline()
		}

	case 2:
		{
			p.SetState(318)
			p.InlinePipeline()
		}

	}
	p.SetState(324)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == preludioParserNEWLINE {
		{
			p.SetState(321)
			p.Nl()
		}

		p.SetState(326)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(327)
		p.Match(preludioParserRPAREN)
	}

	return localctx
}

func (p *preludioParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 22:
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

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 10)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
