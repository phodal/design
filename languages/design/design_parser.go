// Code generated from Design.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Design

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 43, 252,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 3, 2, 3, 2, 3, 2,
	7, 2, 68, 10, 2, 12, 2, 14, 2, 71, 11, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 87, 10, 6,
	5, 6, 89, 10, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 97, 10, 7, 3,
	8, 3, 8, 3, 8, 3, 8, 7, 8, 103, 10, 8, 12, 8, 14, 8, 106, 11, 8, 3, 8,
	3, 8, 3, 9, 3, 9, 3, 9, 5, 9, 113, 10, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	10, 5, 10, 120, 10, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11,
	3, 11, 3, 12, 3, 12, 5, 12, 132, 10, 12, 3, 12, 3, 12, 3, 12, 5, 12, 137,
	10, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14,
	3, 14, 3, 14, 3, 14, 5, 14, 151, 10, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3,
	17, 3, 17, 3, 18, 3, 18, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 20, 7, 20,
	167, 10, 20, 12, 20, 14, 20, 170, 11, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3,
	21, 3, 21, 7, 21, 178, 10, 21, 12, 21, 14, 21, 181, 11, 21, 3, 21, 3, 21,
	3, 22, 3, 22, 3, 22, 3, 22, 5, 22, 189, 10, 22, 3, 23, 3, 23, 7, 23, 193,
	10, 23, 12, 23, 14, 23, 196, 11, 23, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24,
	5, 24, 203, 10, 24, 3, 25, 3, 25, 7, 25, 207, 10, 25, 12, 25, 14, 25, 210,
	11, 25, 3, 26, 3, 26, 7, 26, 214, 10, 26, 12, 26, 14, 26, 217, 11, 26,
	3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 5, 27, 224, 10, 27, 3, 27, 5, 27, 227,
	10, 27, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 29, 3, 29, 3, 30,
	3, 30, 3, 30, 7, 30, 240, 10, 30, 12, 30, 14, 30, 243, 11, 30, 3, 31, 3,
	31, 3, 31, 3, 31, 3, 31, 3, 32, 3, 32, 3, 32, 2, 2, 33, 2, 4, 6, 8, 10,
	12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46,
	48, 50, 52, 54, 56, 58, 60, 62, 2, 4, 3, 2, 3, 4, 4, 2, 14, 14, 20, 20,
	2, 249, 2, 69, 3, 2, 2, 2, 4, 74, 3, 2, 2, 2, 6, 76, 3, 2, 2, 2, 8, 80,
	3, 2, 2, 2, 10, 88, 3, 2, 2, 2, 12, 96, 3, 2, 2, 2, 14, 98, 3, 2, 2, 2,
	16, 112, 3, 2, 2, 2, 18, 114, 3, 2, 2, 2, 20, 121, 3, 2, 2, 2, 22, 129,
	3, 2, 2, 2, 24, 138, 3, 2, 2, 2, 26, 150, 3, 2, 2, 2, 28, 152, 3, 2, 2,
	2, 30, 154, 3, 2, 2, 2, 32, 156, 3, 2, 2, 2, 34, 158, 3, 2, 2, 2, 36, 160,
	3, 2, 2, 2, 38, 162, 3, 2, 2, 2, 40, 173, 3, 2, 2, 2, 42, 184, 3, 2, 2,
	2, 44, 190, 3, 2, 2, 2, 46, 199, 3, 2, 2, 2, 48, 204, 3, 2, 2, 2, 50, 215,
	3, 2, 2, 2, 52, 226, 3, 2, 2, 2, 54, 228, 3, 2, 2, 2, 56, 234, 3, 2, 2,
	2, 58, 241, 3, 2, 2, 2, 60, 244, 3, 2, 2, 2, 62, 249, 3, 2, 2, 2, 64, 68,
	5, 4, 3, 2, 65, 68, 5, 6, 4, 2, 66, 68, 5, 12, 7, 2, 67, 64, 3, 2, 2, 2,
	67, 65, 3, 2, 2, 2, 67, 66, 3, 2, 2, 2, 68, 71, 3, 2, 2, 2, 69, 67, 3,
	2, 2, 2, 69, 70, 3, 2, 2, 2, 70, 72, 3, 2, 2, 2, 71, 69, 3, 2, 2, 2, 72,
	73, 7, 2, 2, 3, 73, 3, 3, 2, 2, 2, 74, 75, 7, 40, 2, 2, 75, 5, 3, 2, 2,
	2, 76, 77, 5, 8, 5, 2, 77, 78, 7, 35, 2, 2, 78, 79, 5, 10, 6, 2, 79, 7,
	3, 2, 2, 2, 80, 81, 7, 40, 2, 2, 81, 9, 3, 2, 2, 2, 82, 89, 7, 40, 2, 2,
	83, 86, 7, 40, 2, 2, 84, 85, 7, 37, 2, 2, 85, 87, 7, 40, 2, 2, 86, 84,
	3, 2, 2, 2, 86, 87, 3, 2, 2, 2, 87, 89, 3, 2, 2, 2, 88, 82, 3, 2, 2, 2,
	88, 83, 3, 2, 2, 2, 89, 11, 3, 2, 2, 2, 90, 97, 5, 6, 4, 2, 91, 97, 5,
	14, 8, 2, 92, 97, 5, 38, 20, 2, 93, 97, 5, 54, 28, 2, 94, 97, 5, 40, 21,
	2, 95, 97, 5, 60, 31, 2, 96, 90, 3, 2, 2, 2, 96, 91, 3, 2, 2, 2, 96, 92,
	3, 2, 2, 2, 96, 93, 3, 2, 2, 2, 96, 94, 3, 2, 2, 2, 96, 95, 3, 2, 2, 2,
	97, 13, 3, 2, 2, 2, 98, 99, 7, 8, 2, 2, 99, 100, 7, 40, 2, 2, 100, 104,
	7, 29, 2, 2, 101, 103, 5, 16, 9, 2, 102, 101, 3, 2, 2, 2, 103, 106, 3,
	2, 2, 2, 104, 102, 3, 2, 2, 2, 104, 105, 3, 2, 2, 2, 105, 107, 3, 2, 2,
	2, 106, 104, 3, 2, 2, 2, 107, 108, 7, 30, 2, 2, 108, 15, 3, 2, 2, 2, 109,
	113, 5, 18, 10, 2, 110, 113, 5, 20, 11, 2, 111, 113, 5, 22, 12, 2, 112,
	109, 3, 2, 2, 2, 112, 110, 3, 2, 2, 2, 112, 111, 3, 2, 2, 2, 113, 17, 3,
	2, 2, 2, 114, 119, 7, 9, 2, 2, 115, 120, 7, 40, 2, 2, 116, 117, 7, 20,
	2, 2, 117, 118, 7, 36, 2, 2, 118, 120, 5, 32, 17, 2, 119, 115, 3, 2, 2,
	2, 119, 116, 3, 2, 2, 2, 120, 19, 3, 2, 2, 2, 121, 122, 7, 10, 2, 2, 122,
	123, 7, 31, 2, 2, 123, 124, 5, 28, 15, 2, 124, 125, 7, 32, 2, 2, 125, 126,
	7, 20, 2, 2, 126, 127, 7, 36, 2, 2, 127, 128, 5, 32, 17, 2, 128, 21, 3,
	2, 2, 2, 129, 131, 7, 11, 2, 2, 130, 132, 5, 34, 18, 2, 131, 130, 3, 2,
	2, 2, 131, 132, 3, 2, 2, 2, 132, 133, 3, 2, 2, 2, 133, 134, 7, 35, 2, 2,
	134, 136, 5, 26, 14, 2, 135, 137, 5, 24, 13, 2, 136, 135, 3, 2, 2, 2, 136,
	137, 3, 2, 2, 2, 137, 23, 3, 2, 2, 2, 138, 139, 7, 12, 2, 2, 139, 140,
	7, 13, 2, 2, 140, 141, 7, 27, 2, 2, 141, 142, 5, 36, 19, 2, 142, 143, 7,
	28, 2, 2, 143, 25, 3, 2, 2, 2, 144, 145, 7, 6, 2, 2, 145, 151, 5, 32, 17,
	2, 146, 147, 7, 7, 2, 2, 147, 148, 7, 20, 2, 2, 148, 149, 7, 36, 2, 2,
	149, 151, 5, 32, 17, 2, 150, 144, 3, 2, 2, 2, 150, 146, 3, 2, 2, 2, 151,
	27, 3, 2, 2, 2, 152, 153, 7, 40, 2, 2, 153, 29, 3, 2, 2, 2, 154, 155, 7,
	40, 2, 2, 155, 31, 3, 2, 2, 2, 156, 157, 7, 40, 2, 2, 157, 33, 3, 2, 2,
	2, 158, 159, 7, 40, 2, 2, 159, 35, 3, 2, 2, 2, 160, 161, 7, 40, 2, 2, 161,
	37, 3, 2, 2, 2, 162, 163, 7, 16, 2, 2, 163, 164, 7, 40, 2, 2, 164, 168,
	7, 29, 2, 2, 165, 167, 5, 42, 22, 2, 166, 165, 3, 2, 2, 2, 167, 170, 3,
	2, 2, 2, 168, 166, 3, 2, 2, 2, 168, 169, 3, 2, 2, 2, 169, 171, 3, 2, 2,
	2, 170, 168, 3, 2, 2, 2, 171, 172, 7, 30, 2, 2, 172, 39, 3, 2, 2, 2, 173,
	174, 7, 17, 2, 2, 174, 175, 7, 40, 2, 2, 175, 179, 7, 29, 2, 2, 176, 178,
	5, 42, 22, 2, 177, 176, 3, 2, 2, 2, 178, 181, 3, 2, 2, 2, 179, 177, 3,
	2, 2, 2, 179, 180, 3, 2, 2, 2, 180, 182, 3, 2, 2, 2, 181, 179, 3, 2, 2,
	2, 182, 183, 7, 30, 2, 2, 183, 41, 3, 2, 2, 2, 184, 188, 7, 40, 2, 2, 185,
	186, 7, 35, 2, 2, 186, 189, 5, 10, 6, 2, 187, 189, 5, 44, 23, 2, 188, 185,
	3, 2, 2, 2, 188, 187, 3, 2, 2, 2, 189, 43, 3, 2, 2, 2, 190, 194, 7, 29,
	2, 2, 191, 193, 5, 46, 24, 2, 192, 191, 3, 2, 2, 2, 193, 196, 3, 2, 2,
	2, 194, 192, 3, 2, 2, 2, 194, 195, 3, 2, 2, 2, 195, 197, 3, 2, 2, 2, 196,
	194, 3, 2, 2, 2, 197, 198, 7, 30, 2, 2, 198, 45, 3, 2, 2, 2, 199, 202,
	7, 3, 2, 2, 200, 203, 5, 48, 25, 2, 201, 203, 5, 50, 26, 2, 202, 200, 3,
	2, 2, 2, 202, 201, 3, 2, 2, 2, 203, 47, 3, 2, 2, 2, 204, 208, 7, 4, 2,
	2, 205, 207, 9, 2, 2, 2, 206, 205, 3, 2, 2, 2, 207, 210, 3, 2, 2, 2, 208,
	206, 3, 2, 2, 2, 208, 209, 3, 2, 2, 2, 209, 49, 3, 2, 2, 2, 210, 208, 3,
	2, 2, 2, 211, 214, 7, 3, 2, 2, 212, 214, 5, 52, 27, 2, 213, 211, 3, 2,
	2, 2, 213, 212, 3, 2, 2, 2, 214, 217, 3, 2, 2, 2, 215, 213, 3, 2, 2, 2,
	215, 216, 3, 2, 2, 2, 216, 51, 3, 2, 2, 2, 217, 215, 3, 2, 2, 2, 218, 227,
	7, 14, 2, 2, 219, 223, 5, 32, 17, 2, 220, 221, 7, 27, 2, 2, 221, 222, 9,
	3, 2, 2, 222, 224, 7, 28, 2, 2, 223, 220, 3, 2, 2, 2, 223, 224, 3, 2, 2,
	2, 224, 227, 3, 2, 2, 2, 225, 227, 7, 20, 2, 2, 226, 218, 3, 2, 2, 2, 226,
	219, 3, 2, 2, 2, 226, 225, 3, 2, 2, 2, 227, 53, 3, 2, 2, 2, 228, 229, 7,
	18, 2, 2, 229, 230, 5, 56, 29, 2, 230, 231, 7, 29, 2, 2, 231, 232, 5, 58,
	30, 2, 232, 233, 7, 30, 2, 2, 233, 55, 3, 2, 2, 2, 234, 235, 7, 40, 2,
	2, 235, 57, 3, 2, 2, 2, 236, 237, 5, 6, 4, 2, 237, 238, 7, 5, 2, 2, 238,
	240, 3, 2, 2, 2, 239, 236, 3, 2, 2, 2, 240, 243, 3, 2, 2, 2, 241, 239,
	3, 2, 2, 2, 241, 242, 3, 2, 2, 2, 242, 59, 3, 2, 2, 2, 243, 241, 3, 2,
	2, 2, 244, 245, 7, 19, 2, 2, 245, 246, 7, 29, 2, 2, 246, 247, 5, 62, 32,
	2, 247, 248, 7, 30, 2, 2, 248, 61, 3, 2, 2, 2, 249, 250, 7, 20, 2, 2, 250,
	63, 3, 2, 2, 2, 24, 67, 69, 86, 88, 96, 104, 112, 119, 131, 136, 150, 168,
	179, 188, 194, 202, 208, 213, 215, 223, 226, 241,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'|'", "'-'", "';'", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "'('", "')'", "'{'", "'}'", "'['",
	"']'", "'\"'", "'''", "':'", "'.'", "','",
}
var symbolicNames = []string{
	"", "", "", "", "GOTO_KEY", "SHOW_KEY", "FLOW", "SEE", "DO", "REACT", "WITHTEXT",
	"ANIMATE", "GridSize", "POSITION", "PAGE", "COMPONENT", "STYLE", "LIBRARAY",
	"STRING_LITERAL", "WS", "COMMENT", "LINE_COMMENT", "EmptyLine", "Space",
	"NewLine", "LPAREN", "RPAREN", "LBRACE", "RBRACE", "LBRACK", "RBRACK",
	"Quote", "SingleQuote", "COLON", "DOT", "COMMA", "LETTER", "DIGITS", "IDENTIFIER",
	"DIGITS_IDENTIFIER", "CONFIG_VALUE", "DECIMAL_LITERAL",
}

var ruleNames = []string{
	"start", "comment", "configDecalartion", "configKey", "configValue", "decalartions",
	"flowDecalartion", "flowBodyDecalartion", "seeDecalartion", "doDecalartion",
	"reactDecalartion", "animateDecalartion", "actionKey", "actionName", "componentValue",
	"componentName", "sceneName", "animateName", "pageDecalartion", "componentDecalartion",
	"componentBodyDecalartion", "layoutDecalaration", "layoutBodyDecalartion",
	"emptyLine", "layoutLine", "componentUseDeclaration", "styleDecalartion",
	"styleName", "styleBody", "libraryDecalartion", "libraryBody",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type DesignParser struct {
	*antlr.BaseParser
}

func NewDesignParser(input antlr.TokenStream) *DesignParser {
	this := new(DesignParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Design.g4"

	return this
}

// DesignParser tokens.
const (
	DesignParserEOF               = antlr.TokenEOF
	DesignParserT__0              = 1
	DesignParserT__1              = 2
	DesignParserT__2              = 3
	DesignParserGOTO_KEY          = 4
	DesignParserSHOW_KEY          = 5
	DesignParserFLOW              = 6
	DesignParserSEE               = 7
	DesignParserDO                = 8
	DesignParserREACT             = 9
	DesignParserWITHTEXT          = 10
	DesignParserANIMATE           = 11
	DesignParserGridSize          = 12
	DesignParserPOSITION          = 13
	DesignParserPAGE              = 14
	DesignParserCOMPONENT         = 15
	DesignParserSTYLE             = 16
	DesignParserLIBRARAY          = 17
	DesignParserSTRING_LITERAL    = 18
	DesignParserWS                = 19
	DesignParserCOMMENT           = 20
	DesignParserLINE_COMMENT      = 21
	DesignParserEmptyLine         = 22
	DesignParserSpace             = 23
	DesignParserNewLine           = 24
	DesignParserLPAREN            = 25
	DesignParserRPAREN            = 26
	DesignParserLBRACE            = 27
	DesignParserRBRACE            = 28
	DesignParserLBRACK            = 29
	DesignParserRBRACK            = 30
	DesignParserQuote             = 31
	DesignParserSingleQuote       = 32
	DesignParserCOLON             = 33
	DesignParserDOT               = 34
	DesignParserCOMMA             = 35
	DesignParserLETTER            = 36
	DesignParserDIGITS            = 37
	DesignParserIDENTIFIER        = 38
	DesignParserDIGITS_IDENTIFIER = 39
	DesignParserCONFIG_VALUE      = 40
	DesignParserDECIMAL_LITERAL   = 41
)

// DesignParser rules.
const (
	DesignParserRULE_start                    = 0
	DesignParserRULE_comment                  = 1
	DesignParserRULE_configDecalartion        = 2
	DesignParserRULE_configKey                = 3
	DesignParserRULE_configValue              = 4
	DesignParserRULE_decalartions             = 5
	DesignParserRULE_flowDecalartion          = 6
	DesignParserRULE_flowBodyDecalartion      = 7
	DesignParserRULE_seeDecalartion           = 8
	DesignParserRULE_doDecalartion            = 9
	DesignParserRULE_reactDecalartion         = 10
	DesignParserRULE_animateDecalartion       = 11
	DesignParserRULE_actionKey                = 12
	DesignParserRULE_actionName               = 13
	DesignParserRULE_componentValue           = 14
	DesignParserRULE_componentName            = 15
	DesignParserRULE_sceneName                = 16
	DesignParserRULE_animateName              = 17
	DesignParserRULE_pageDecalartion          = 18
	DesignParserRULE_componentDecalartion     = 19
	DesignParserRULE_componentBodyDecalartion = 20
	DesignParserRULE_layoutDecalaration       = 21
	DesignParserRULE_layoutBodyDecalartion    = 22
	DesignParserRULE_emptyLine                = 23
	DesignParserRULE_layoutLine               = 24
	DesignParserRULE_componentUseDeclaration  = 25
	DesignParserRULE_styleDecalartion         = 26
	DesignParserRULE_styleName                = 27
	DesignParserRULE_styleBody                = 28
	DesignParserRULE_libraryDecalartion       = 29
	DesignParserRULE_libraryBody              = 30
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DesignParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesignParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) EOF() antlr.TerminalNode {
	return s.GetToken(DesignParserEOF, 0)
}

func (s *StartContext) AllComment() []ICommentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICommentContext)(nil)).Elem())
	var tst = make([]ICommentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICommentContext)
		}
	}

	return tst
}

func (s *StartContext) Comment(i int) ICommentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICommentContext)
}

func (s *StartContext) AllConfigDecalartion() []IConfigDecalartionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IConfigDecalartionContext)(nil)).Elem())
	var tst = make([]IConfigDecalartionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IConfigDecalartionContext)
		}
	}

	return tst
}

func (s *StartContext) ConfigDecalartion(i int) IConfigDecalartionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConfigDecalartionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IConfigDecalartionContext)
}

func (s *StartContext) AllDecalartions() []IDecalartionsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDecalartionsContext)(nil)).Elem())
	var tst = make([]IDecalartionsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDecalartionsContext)
		}
	}

	return tst
}

func (s *StartContext) Decalartions(i int) IDecalartionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDecalartionsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDecalartionsContext)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.ExitStart(s)
	}
}

func (s *StartContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DesignVisitor:
		return t.VisitStart(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DesignParser) Start() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, DesignParserRULE_start)
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
	p.SetState(67)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<DesignParserFLOW)|(1<<DesignParserPAGE)|(1<<DesignParserCOMPONENT)|(1<<DesignParserSTYLE)|(1<<DesignParserLIBRARAY))) != 0) || _la == DesignParserIDENTIFIER {
		p.SetState(65)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(62)
				p.Comment()
			}

		case 2:
			{
				p.SetState(63)
				p.ConfigDecalartion()
			}

		case 3:
			{
				p.SetState(64)
				p.Decalartions()
			}

		}

		p.SetState(69)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(70)
		p.Match(DesignParserEOF)
	}

	return localctx
}

// ICommentContext is an interface to support dynamic dispatch.
type ICommentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCommentContext differentiates from other interfaces.
	IsCommentContext()
}

type CommentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommentContext() *CommentContext {
	var p = new(CommentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DesignParserRULE_comment
	return p
}

func (*CommentContext) IsCommentContext() {}

func NewCommentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommentContext {
	var p = new(CommentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesignParserRULE_comment

	return p
}

func (s *CommentContext) GetParser() antlr.Parser { return s.parser }

func (s *CommentContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(DesignParserIDENTIFIER, 0)
}

func (s *CommentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.EnterComment(s)
	}
}

func (s *CommentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.ExitComment(s)
	}
}

func (s *CommentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DesignVisitor:
		return t.VisitComment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DesignParser) Comment() (localctx ICommentContext) {
	localctx = NewCommentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, DesignParserRULE_comment)

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
		p.SetState(72)
		p.Match(DesignParserIDENTIFIER)
	}

	return localctx
}

// IConfigDecalartionContext is an interface to support dynamic dispatch.
type IConfigDecalartionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConfigDecalartionContext differentiates from other interfaces.
	IsConfigDecalartionContext()
}

type ConfigDecalartionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConfigDecalartionContext() *ConfigDecalartionContext {
	var p = new(ConfigDecalartionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DesignParserRULE_configDecalartion
	return p
}

func (*ConfigDecalartionContext) IsConfigDecalartionContext() {}

func NewConfigDecalartionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConfigDecalartionContext {
	var p = new(ConfigDecalartionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesignParserRULE_configDecalartion

	return p
}

func (s *ConfigDecalartionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConfigDecalartionContext) ConfigKey() IConfigKeyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConfigKeyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConfigKeyContext)
}

func (s *ConfigDecalartionContext) COLON() antlr.TerminalNode {
	return s.GetToken(DesignParserCOLON, 0)
}

func (s *ConfigDecalartionContext) ConfigValue() IConfigValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConfigValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConfigValueContext)
}

func (s *ConfigDecalartionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConfigDecalartionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConfigDecalartionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.EnterConfigDecalartion(s)
	}
}

func (s *ConfigDecalartionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.ExitConfigDecalartion(s)
	}
}

func (s *ConfigDecalartionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DesignVisitor:
		return t.VisitConfigDecalartion(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DesignParser) ConfigDecalartion() (localctx IConfigDecalartionContext) {
	localctx = NewConfigDecalartionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, DesignParserRULE_configDecalartion)

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
		p.SetState(74)
		p.ConfigKey()
	}
	{
		p.SetState(75)
		p.Match(DesignParserCOLON)
	}
	{
		p.SetState(76)
		p.ConfigValue()
	}

	return localctx
}

// IConfigKeyContext is an interface to support dynamic dispatch.
type IConfigKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConfigKeyContext differentiates from other interfaces.
	IsConfigKeyContext()
}

type ConfigKeyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConfigKeyContext() *ConfigKeyContext {
	var p = new(ConfigKeyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DesignParserRULE_configKey
	return p
}

func (*ConfigKeyContext) IsConfigKeyContext() {}

func NewConfigKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConfigKeyContext {
	var p = new(ConfigKeyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesignParserRULE_configKey

	return p
}

func (s *ConfigKeyContext) GetParser() antlr.Parser { return s.parser }

func (s *ConfigKeyContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(DesignParserIDENTIFIER, 0)
}

func (s *ConfigKeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConfigKeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConfigKeyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.EnterConfigKey(s)
	}
}

func (s *ConfigKeyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.ExitConfigKey(s)
	}
}

func (s *ConfigKeyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DesignVisitor:
		return t.VisitConfigKey(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DesignParser) ConfigKey() (localctx IConfigKeyContext) {
	localctx = NewConfigKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, DesignParserRULE_configKey)

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
		p.SetState(78)
		p.Match(DesignParserIDENTIFIER)
	}

	return localctx
}

// IConfigValueContext is an interface to support dynamic dispatch.
type IConfigValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConfigValueContext differentiates from other interfaces.
	IsConfigValueContext()
}

type ConfigValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConfigValueContext() *ConfigValueContext {
	var p = new(ConfigValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DesignParserRULE_configValue
	return p
}

func (*ConfigValueContext) IsConfigValueContext() {}

func NewConfigValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConfigValueContext {
	var p = new(ConfigValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesignParserRULE_configValue

	return p
}

func (s *ConfigValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ConfigValueContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(DesignParserIDENTIFIER)
}

func (s *ConfigValueContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(DesignParserIDENTIFIER, i)
}

func (s *ConfigValueContext) COMMA() antlr.TerminalNode {
	return s.GetToken(DesignParserCOMMA, 0)
}

func (s *ConfigValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConfigValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConfigValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.EnterConfigValue(s)
	}
}

func (s *ConfigValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.ExitConfigValue(s)
	}
}

func (s *ConfigValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DesignVisitor:
		return t.VisitConfigValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DesignParser) ConfigValue() (localctx IConfigValueContext) {
	localctx = NewConfigValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, DesignParserRULE_configValue)
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

	p.SetState(86)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(80)
			p.Match(DesignParserIDENTIFIER)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(81)
			p.Match(DesignParserIDENTIFIER)
		}
		p.SetState(84)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DesignParserCOMMA {
			{
				p.SetState(82)
				p.Match(DesignParserCOMMA)
			}
			{
				p.SetState(83)
				p.Match(DesignParserIDENTIFIER)
			}

		}

	}

	return localctx
}

// IDecalartionsContext is an interface to support dynamic dispatch.
type IDecalartionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDecalartionsContext differentiates from other interfaces.
	IsDecalartionsContext()
}

type DecalartionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDecalartionsContext() *DecalartionsContext {
	var p = new(DecalartionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DesignParserRULE_decalartions
	return p
}

func (*DecalartionsContext) IsDecalartionsContext() {}

func NewDecalartionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DecalartionsContext {
	var p = new(DecalartionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesignParserRULE_decalartions

	return p
}

func (s *DecalartionsContext) GetParser() antlr.Parser { return s.parser }

func (s *DecalartionsContext) ConfigDecalartion() IConfigDecalartionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConfigDecalartionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConfigDecalartionContext)
}

func (s *DecalartionsContext) FlowDecalartion() IFlowDecalartionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFlowDecalartionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFlowDecalartionContext)
}

func (s *DecalartionsContext) PageDecalartion() IPageDecalartionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPageDecalartionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPageDecalartionContext)
}

func (s *DecalartionsContext) StyleDecalartion() IStyleDecalartionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStyleDecalartionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStyleDecalartionContext)
}

func (s *DecalartionsContext) ComponentDecalartion() IComponentDecalartionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComponentDecalartionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComponentDecalartionContext)
}

func (s *DecalartionsContext) LibraryDecalartion() ILibraryDecalartionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILibraryDecalartionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILibraryDecalartionContext)
}

func (s *DecalartionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DecalartionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DecalartionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.EnterDecalartions(s)
	}
}

func (s *DecalartionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.ExitDecalartions(s)
	}
}

func (s *DecalartionsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DesignVisitor:
		return t.VisitDecalartions(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DesignParser) Decalartions() (localctx IDecalartionsContext) {
	localctx = NewDecalartionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, DesignParserRULE_decalartions)

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

	p.SetState(94)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DesignParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(88)
			p.ConfigDecalartion()
		}

	case DesignParserFLOW:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(89)
			p.FlowDecalartion()
		}

	case DesignParserPAGE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(90)
			p.PageDecalartion()
		}

	case DesignParserSTYLE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(91)
			p.StyleDecalartion()
		}

	case DesignParserCOMPONENT:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(92)
			p.ComponentDecalartion()
		}

	case DesignParserLIBRARAY:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(93)
			p.LibraryDecalartion()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFlowDecalartionContext is an interface to support dynamic dispatch.
type IFlowDecalartionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFlowDecalartionContext differentiates from other interfaces.
	IsFlowDecalartionContext()
}

type FlowDecalartionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFlowDecalartionContext() *FlowDecalartionContext {
	var p = new(FlowDecalartionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DesignParserRULE_flowDecalartion
	return p
}

func (*FlowDecalartionContext) IsFlowDecalartionContext() {}

func NewFlowDecalartionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FlowDecalartionContext {
	var p = new(FlowDecalartionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesignParserRULE_flowDecalartion

	return p
}

func (s *FlowDecalartionContext) GetParser() antlr.Parser { return s.parser }

func (s *FlowDecalartionContext) FLOW() antlr.TerminalNode {
	return s.GetToken(DesignParserFLOW, 0)
}

func (s *FlowDecalartionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(DesignParserIDENTIFIER, 0)
}

func (s *FlowDecalartionContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(DesignParserLBRACE, 0)
}

func (s *FlowDecalartionContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(DesignParserRBRACE, 0)
}

func (s *FlowDecalartionContext) AllFlowBodyDecalartion() []IFlowBodyDecalartionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFlowBodyDecalartionContext)(nil)).Elem())
	var tst = make([]IFlowBodyDecalartionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFlowBodyDecalartionContext)
		}
	}

	return tst
}

func (s *FlowDecalartionContext) FlowBodyDecalartion(i int) IFlowBodyDecalartionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFlowBodyDecalartionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFlowBodyDecalartionContext)
}

func (s *FlowDecalartionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FlowDecalartionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FlowDecalartionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.EnterFlowDecalartion(s)
	}
}

func (s *FlowDecalartionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.ExitFlowDecalartion(s)
	}
}

func (s *FlowDecalartionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DesignVisitor:
		return t.VisitFlowDecalartion(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DesignParser) FlowDecalartion() (localctx IFlowDecalartionContext) {
	localctx = NewFlowDecalartionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, DesignParserRULE_flowDecalartion)
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
		p.SetState(96)
		p.Match(DesignParserFLOW)
	}
	{
		p.SetState(97)
		p.Match(DesignParserIDENTIFIER)
	}
	{
		p.SetState(98)
		p.Match(DesignParserLBRACE)
	}
	p.SetState(102)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<DesignParserSEE)|(1<<DesignParserDO)|(1<<DesignParserREACT))) != 0 {
		{
			p.SetState(99)
			p.FlowBodyDecalartion()
		}

		p.SetState(104)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(105)
		p.Match(DesignParserRBRACE)
	}

	return localctx
}

// IFlowBodyDecalartionContext is an interface to support dynamic dispatch.
type IFlowBodyDecalartionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFlowBodyDecalartionContext differentiates from other interfaces.
	IsFlowBodyDecalartionContext()
}

type FlowBodyDecalartionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFlowBodyDecalartionContext() *FlowBodyDecalartionContext {
	var p = new(FlowBodyDecalartionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DesignParserRULE_flowBodyDecalartion
	return p
}

func (*FlowBodyDecalartionContext) IsFlowBodyDecalartionContext() {}

func NewFlowBodyDecalartionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FlowBodyDecalartionContext {
	var p = new(FlowBodyDecalartionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesignParserRULE_flowBodyDecalartion

	return p
}

func (s *FlowBodyDecalartionContext) GetParser() antlr.Parser { return s.parser }

func (s *FlowBodyDecalartionContext) SeeDecalartion() ISeeDecalartionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISeeDecalartionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISeeDecalartionContext)
}

func (s *FlowBodyDecalartionContext) DoDecalartion() IDoDecalartionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDoDecalartionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDoDecalartionContext)
}

func (s *FlowBodyDecalartionContext) ReactDecalartion() IReactDecalartionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReactDecalartionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReactDecalartionContext)
}

func (s *FlowBodyDecalartionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FlowBodyDecalartionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FlowBodyDecalartionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.EnterFlowBodyDecalartion(s)
	}
}

func (s *FlowBodyDecalartionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.ExitFlowBodyDecalartion(s)
	}
}

func (s *FlowBodyDecalartionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DesignVisitor:
		return t.VisitFlowBodyDecalartion(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DesignParser) FlowBodyDecalartion() (localctx IFlowBodyDecalartionContext) {
	localctx = NewFlowBodyDecalartionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, DesignParserRULE_flowBodyDecalartion)

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

	p.SetState(110)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DesignParserSEE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(107)
			p.SeeDecalartion()
		}

	case DesignParserDO:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(108)
			p.DoDecalartion()
		}

	case DesignParserREACT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(109)
			p.ReactDecalartion()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ISeeDecalartionContext is an interface to support dynamic dispatch.
type ISeeDecalartionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSeeDecalartionContext differentiates from other interfaces.
	IsSeeDecalartionContext()
}

type SeeDecalartionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySeeDecalartionContext() *SeeDecalartionContext {
	var p = new(SeeDecalartionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DesignParserRULE_seeDecalartion
	return p
}

func (*SeeDecalartionContext) IsSeeDecalartionContext() {}

func NewSeeDecalartionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SeeDecalartionContext {
	var p = new(SeeDecalartionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesignParserRULE_seeDecalartion

	return p
}

func (s *SeeDecalartionContext) GetParser() antlr.Parser { return s.parser }

func (s *SeeDecalartionContext) SEE() antlr.TerminalNode {
	return s.GetToken(DesignParserSEE, 0)
}

func (s *SeeDecalartionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(DesignParserIDENTIFIER, 0)
}

func (s *SeeDecalartionContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(DesignParserSTRING_LITERAL, 0)
}

func (s *SeeDecalartionContext) DOT() antlr.TerminalNode {
	return s.GetToken(DesignParserDOT, 0)
}

func (s *SeeDecalartionContext) ComponentName() IComponentNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComponentNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComponentNameContext)
}

func (s *SeeDecalartionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SeeDecalartionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SeeDecalartionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.EnterSeeDecalartion(s)
	}
}

func (s *SeeDecalartionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.ExitSeeDecalartion(s)
	}
}

func (s *SeeDecalartionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DesignVisitor:
		return t.VisitSeeDecalartion(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DesignParser) SeeDecalartion() (localctx ISeeDecalartionContext) {
	localctx = NewSeeDecalartionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, DesignParserRULE_seeDecalartion)

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
		p.SetState(112)
		p.Match(DesignParserSEE)
	}
	p.SetState(117)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DesignParserIDENTIFIER:
		{
			p.SetState(113)
			p.Match(DesignParserIDENTIFIER)
		}

	case DesignParserSTRING_LITERAL:
		{
			p.SetState(114)
			p.Match(DesignParserSTRING_LITERAL)
		}
		{
			p.SetState(115)
			p.Match(DesignParserDOT)
		}
		{
			p.SetState(116)
			p.ComponentName()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IDoDecalartionContext is an interface to support dynamic dispatch.
type IDoDecalartionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDoDecalartionContext differentiates from other interfaces.
	IsDoDecalartionContext()
}

type DoDecalartionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDoDecalartionContext() *DoDecalartionContext {
	var p = new(DoDecalartionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DesignParserRULE_doDecalartion
	return p
}

func (*DoDecalartionContext) IsDoDecalartionContext() {}

func NewDoDecalartionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DoDecalartionContext {
	var p = new(DoDecalartionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesignParserRULE_doDecalartion

	return p
}

func (s *DoDecalartionContext) GetParser() antlr.Parser { return s.parser }

func (s *DoDecalartionContext) DO() antlr.TerminalNode {
	return s.GetToken(DesignParserDO, 0)
}

func (s *DoDecalartionContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(DesignParserLBRACK, 0)
}

func (s *DoDecalartionContext) ActionName() IActionNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IActionNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IActionNameContext)
}

func (s *DoDecalartionContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(DesignParserRBRACK, 0)
}

func (s *DoDecalartionContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(DesignParserSTRING_LITERAL, 0)
}

func (s *DoDecalartionContext) DOT() antlr.TerminalNode {
	return s.GetToken(DesignParserDOT, 0)
}

func (s *DoDecalartionContext) ComponentName() IComponentNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComponentNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComponentNameContext)
}

func (s *DoDecalartionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DoDecalartionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DoDecalartionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.EnterDoDecalartion(s)
	}
}

func (s *DoDecalartionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.ExitDoDecalartion(s)
	}
}

func (s *DoDecalartionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DesignVisitor:
		return t.VisitDoDecalartion(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DesignParser) DoDecalartion() (localctx IDoDecalartionContext) {
	localctx = NewDoDecalartionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, DesignParserRULE_doDecalartion)

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
		p.Match(DesignParserDO)
	}
	{
		p.SetState(120)
		p.Match(DesignParserLBRACK)
	}
	{
		p.SetState(121)
		p.ActionName()
	}
	{
		p.SetState(122)
		p.Match(DesignParserRBRACK)
	}
	{
		p.SetState(123)
		p.Match(DesignParserSTRING_LITERAL)
	}
	{
		p.SetState(124)
		p.Match(DesignParserDOT)
	}
	{
		p.SetState(125)
		p.ComponentName()
	}

	return localctx
}

// IReactDecalartionContext is an interface to support dynamic dispatch.
type IReactDecalartionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReactDecalartionContext differentiates from other interfaces.
	IsReactDecalartionContext()
}

type ReactDecalartionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReactDecalartionContext() *ReactDecalartionContext {
	var p = new(ReactDecalartionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DesignParserRULE_reactDecalartion
	return p
}

func (*ReactDecalartionContext) IsReactDecalartionContext() {}

func NewReactDecalartionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReactDecalartionContext {
	var p = new(ReactDecalartionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesignParserRULE_reactDecalartion

	return p
}

func (s *ReactDecalartionContext) GetParser() antlr.Parser { return s.parser }

func (s *ReactDecalartionContext) REACT() antlr.TerminalNode {
	return s.GetToken(DesignParserREACT, 0)
}

func (s *ReactDecalartionContext) COLON() antlr.TerminalNode {
	return s.GetToken(DesignParserCOLON, 0)
}

func (s *ReactDecalartionContext) ActionKey() IActionKeyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IActionKeyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IActionKeyContext)
}

func (s *ReactDecalartionContext) SceneName() ISceneNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISceneNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISceneNameContext)
}

func (s *ReactDecalartionContext) AnimateDecalartion() IAnimateDecalartionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnimateDecalartionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAnimateDecalartionContext)
}

func (s *ReactDecalartionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReactDecalartionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReactDecalartionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.EnterReactDecalartion(s)
	}
}

func (s *ReactDecalartionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.ExitReactDecalartion(s)
	}
}

func (s *ReactDecalartionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DesignVisitor:
		return t.VisitReactDecalartion(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DesignParser) ReactDecalartion() (localctx IReactDecalartionContext) {
	localctx = NewReactDecalartionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, DesignParserRULE_reactDecalartion)
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
		p.SetState(127)
		p.Match(DesignParserREACT)
	}
	p.SetState(129)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DesignParserIDENTIFIER {
		{
			p.SetState(128)
			p.SceneName()
		}

	}
	{
		p.SetState(131)
		p.Match(DesignParserCOLON)
	}
	{
		p.SetState(132)
		p.ActionKey()
	}
	p.SetState(134)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DesignParserWITHTEXT {
		{
			p.SetState(133)
			p.AnimateDecalartion()
		}

	}

	return localctx
}

// IAnimateDecalartionContext is an interface to support dynamic dispatch.
type IAnimateDecalartionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAnimateDecalartionContext differentiates from other interfaces.
	IsAnimateDecalartionContext()
}

type AnimateDecalartionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAnimateDecalartionContext() *AnimateDecalartionContext {
	var p = new(AnimateDecalartionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DesignParserRULE_animateDecalartion
	return p
}

func (*AnimateDecalartionContext) IsAnimateDecalartionContext() {}

func NewAnimateDecalartionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AnimateDecalartionContext {
	var p = new(AnimateDecalartionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesignParserRULE_animateDecalartion

	return p
}

func (s *AnimateDecalartionContext) GetParser() antlr.Parser { return s.parser }

func (s *AnimateDecalartionContext) WITHTEXT() antlr.TerminalNode {
	return s.GetToken(DesignParserWITHTEXT, 0)
}

func (s *AnimateDecalartionContext) ANIMATE() antlr.TerminalNode {
	return s.GetToken(DesignParserANIMATE, 0)
}

func (s *AnimateDecalartionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(DesignParserLPAREN, 0)
}

func (s *AnimateDecalartionContext) AnimateName() IAnimateNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnimateNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAnimateNameContext)
}

func (s *AnimateDecalartionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(DesignParserRPAREN, 0)
}

func (s *AnimateDecalartionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AnimateDecalartionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AnimateDecalartionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.EnterAnimateDecalartion(s)
	}
}

func (s *AnimateDecalartionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.ExitAnimateDecalartion(s)
	}
}

func (s *AnimateDecalartionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DesignVisitor:
		return t.VisitAnimateDecalartion(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DesignParser) AnimateDecalartion() (localctx IAnimateDecalartionContext) {
	localctx = NewAnimateDecalartionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, DesignParserRULE_animateDecalartion)

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
		p.Match(DesignParserWITHTEXT)
	}
	{
		p.SetState(137)
		p.Match(DesignParserANIMATE)
	}
	{
		p.SetState(138)
		p.Match(DesignParserLPAREN)
	}
	{
		p.SetState(139)
		p.AnimateName()
	}
	{
		p.SetState(140)
		p.Match(DesignParserRPAREN)
	}

	return localctx
}

// IActionKeyContext is an interface to support dynamic dispatch.
type IActionKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsActionKeyContext differentiates from other interfaces.
	IsActionKeyContext()
}

type ActionKeyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyActionKeyContext() *ActionKeyContext {
	var p = new(ActionKeyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DesignParserRULE_actionKey
	return p
}

func (*ActionKeyContext) IsActionKeyContext() {}

func NewActionKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ActionKeyContext {
	var p = new(ActionKeyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesignParserRULE_actionKey

	return p
}

func (s *ActionKeyContext) GetParser() antlr.Parser { return s.parser }

func (s *ActionKeyContext) GOTO_KEY() antlr.TerminalNode {
	return s.GetToken(DesignParserGOTO_KEY, 0)
}

func (s *ActionKeyContext) ComponentName() IComponentNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComponentNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComponentNameContext)
}

func (s *ActionKeyContext) SHOW_KEY() antlr.TerminalNode {
	return s.GetToken(DesignParserSHOW_KEY, 0)
}

func (s *ActionKeyContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(DesignParserSTRING_LITERAL, 0)
}

func (s *ActionKeyContext) DOT() antlr.TerminalNode {
	return s.GetToken(DesignParserDOT, 0)
}

func (s *ActionKeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ActionKeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ActionKeyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.EnterActionKey(s)
	}
}

func (s *ActionKeyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.ExitActionKey(s)
	}
}

func (s *ActionKeyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DesignVisitor:
		return t.VisitActionKey(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DesignParser) ActionKey() (localctx IActionKeyContext) {
	localctx = NewActionKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, DesignParserRULE_actionKey)

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

	p.SetState(148)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DesignParserGOTO_KEY:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(142)
			p.Match(DesignParserGOTO_KEY)
		}
		{
			p.SetState(143)
			p.ComponentName()
		}

	case DesignParserSHOW_KEY:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(144)
			p.Match(DesignParserSHOW_KEY)
		}
		{
			p.SetState(145)
			p.Match(DesignParserSTRING_LITERAL)
		}
		{
			p.SetState(146)
			p.Match(DesignParserDOT)
		}
		{
			p.SetState(147)
			p.ComponentName()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IActionNameContext is an interface to support dynamic dispatch.
type IActionNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsActionNameContext differentiates from other interfaces.
	IsActionNameContext()
}

type ActionNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyActionNameContext() *ActionNameContext {
	var p = new(ActionNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DesignParserRULE_actionName
	return p
}

func (*ActionNameContext) IsActionNameContext() {}

func NewActionNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ActionNameContext {
	var p = new(ActionNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesignParserRULE_actionName

	return p
}

func (s *ActionNameContext) GetParser() antlr.Parser { return s.parser }

func (s *ActionNameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(DesignParserIDENTIFIER, 0)
}

func (s *ActionNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ActionNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ActionNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.EnterActionName(s)
	}
}

func (s *ActionNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.ExitActionName(s)
	}
}

func (s *ActionNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DesignVisitor:
		return t.VisitActionName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DesignParser) ActionName() (localctx IActionNameContext) {
	localctx = NewActionNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, DesignParserRULE_actionName)

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
		p.SetState(150)
		p.Match(DesignParserIDENTIFIER)
	}

	return localctx
}

// IComponentValueContext is an interface to support dynamic dispatch.
type IComponentValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsComponentValueContext differentiates from other interfaces.
	IsComponentValueContext()
}

type ComponentValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComponentValueContext() *ComponentValueContext {
	var p = new(ComponentValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DesignParserRULE_componentValue
	return p
}

func (*ComponentValueContext) IsComponentValueContext() {}

func NewComponentValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComponentValueContext {
	var p = new(ComponentValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesignParserRULE_componentValue

	return p
}

func (s *ComponentValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ComponentValueContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(DesignParserIDENTIFIER, 0)
}

func (s *ComponentValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComponentValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComponentValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.EnterComponentValue(s)
	}
}

func (s *ComponentValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.ExitComponentValue(s)
	}
}

func (s *ComponentValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DesignVisitor:
		return t.VisitComponentValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DesignParser) ComponentValue() (localctx IComponentValueContext) {
	localctx = NewComponentValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, DesignParserRULE_componentValue)

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
		p.SetState(152)
		p.Match(DesignParserIDENTIFIER)
	}

	return localctx
}

// IComponentNameContext is an interface to support dynamic dispatch.
type IComponentNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsComponentNameContext differentiates from other interfaces.
	IsComponentNameContext()
}

type ComponentNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComponentNameContext() *ComponentNameContext {
	var p = new(ComponentNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DesignParserRULE_componentName
	return p
}

func (*ComponentNameContext) IsComponentNameContext() {}

func NewComponentNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComponentNameContext {
	var p = new(ComponentNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesignParserRULE_componentName

	return p
}

func (s *ComponentNameContext) GetParser() antlr.Parser { return s.parser }

func (s *ComponentNameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(DesignParserIDENTIFIER, 0)
}

func (s *ComponentNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComponentNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComponentNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.EnterComponentName(s)
	}
}

func (s *ComponentNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.ExitComponentName(s)
	}
}

func (s *ComponentNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DesignVisitor:
		return t.VisitComponentName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DesignParser) ComponentName() (localctx IComponentNameContext) {
	localctx = NewComponentNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, DesignParserRULE_componentName)

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
		p.Match(DesignParserIDENTIFIER)
	}

	return localctx
}

// ISceneNameContext is an interface to support dynamic dispatch.
type ISceneNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSceneNameContext differentiates from other interfaces.
	IsSceneNameContext()
}

type SceneNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySceneNameContext() *SceneNameContext {
	var p = new(SceneNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DesignParserRULE_sceneName
	return p
}

func (*SceneNameContext) IsSceneNameContext() {}

func NewSceneNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SceneNameContext {
	var p = new(SceneNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesignParserRULE_sceneName

	return p
}

func (s *SceneNameContext) GetParser() antlr.Parser { return s.parser }

func (s *SceneNameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(DesignParserIDENTIFIER, 0)
}

func (s *SceneNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SceneNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SceneNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.EnterSceneName(s)
	}
}

func (s *SceneNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.ExitSceneName(s)
	}
}

func (s *SceneNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DesignVisitor:
		return t.VisitSceneName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DesignParser) SceneName() (localctx ISceneNameContext) {
	localctx = NewSceneNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, DesignParserRULE_sceneName)

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
		p.SetState(156)
		p.Match(DesignParserIDENTIFIER)
	}

	return localctx
}

// IAnimateNameContext is an interface to support dynamic dispatch.
type IAnimateNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAnimateNameContext differentiates from other interfaces.
	IsAnimateNameContext()
}

type AnimateNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAnimateNameContext() *AnimateNameContext {
	var p = new(AnimateNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DesignParserRULE_animateName
	return p
}

func (*AnimateNameContext) IsAnimateNameContext() {}

func NewAnimateNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AnimateNameContext {
	var p = new(AnimateNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesignParserRULE_animateName

	return p
}

func (s *AnimateNameContext) GetParser() antlr.Parser { return s.parser }

func (s *AnimateNameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(DesignParserIDENTIFIER, 0)
}

func (s *AnimateNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AnimateNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AnimateNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.EnterAnimateName(s)
	}
}

func (s *AnimateNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.ExitAnimateName(s)
	}
}

func (s *AnimateNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DesignVisitor:
		return t.VisitAnimateName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DesignParser) AnimateName() (localctx IAnimateNameContext) {
	localctx = NewAnimateNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, DesignParserRULE_animateName)

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
		p.Match(DesignParserIDENTIFIER)
	}

	return localctx
}

// IPageDecalartionContext is an interface to support dynamic dispatch.
type IPageDecalartionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPageDecalartionContext differentiates from other interfaces.
	IsPageDecalartionContext()
}

type PageDecalartionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPageDecalartionContext() *PageDecalartionContext {
	var p = new(PageDecalartionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DesignParserRULE_pageDecalartion
	return p
}

func (*PageDecalartionContext) IsPageDecalartionContext() {}

func NewPageDecalartionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PageDecalartionContext {
	var p = new(PageDecalartionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesignParserRULE_pageDecalartion

	return p
}

func (s *PageDecalartionContext) GetParser() antlr.Parser { return s.parser }

func (s *PageDecalartionContext) PAGE() antlr.TerminalNode {
	return s.GetToken(DesignParserPAGE, 0)
}

func (s *PageDecalartionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(DesignParserIDENTIFIER, 0)
}

func (s *PageDecalartionContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(DesignParserLBRACE, 0)
}

func (s *PageDecalartionContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(DesignParserRBRACE, 0)
}

func (s *PageDecalartionContext) AllComponentBodyDecalartion() []IComponentBodyDecalartionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IComponentBodyDecalartionContext)(nil)).Elem())
	var tst = make([]IComponentBodyDecalartionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IComponentBodyDecalartionContext)
		}
	}

	return tst
}

func (s *PageDecalartionContext) ComponentBodyDecalartion(i int) IComponentBodyDecalartionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComponentBodyDecalartionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IComponentBodyDecalartionContext)
}

func (s *PageDecalartionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PageDecalartionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PageDecalartionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.EnterPageDecalartion(s)
	}
}

func (s *PageDecalartionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.ExitPageDecalartion(s)
	}
}

func (s *PageDecalartionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DesignVisitor:
		return t.VisitPageDecalartion(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DesignParser) PageDecalartion() (localctx IPageDecalartionContext) {
	localctx = NewPageDecalartionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, DesignParserRULE_pageDecalartion)
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
		p.SetState(160)
		p.Match(DesignParserPAGE)
	}
	{
		p.SetState(161)
		p.Match(DesignParserIDENTIFIER)
	}
	{
		p.SetState(162)
		p.Match(DesignParserLBRACE)
	}
	p.SetState(166)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DesignParserIDENTIFIER {
		{
			p.SetState(163)
			p.ComponentBodyDecalartion()
		}

		p.SetState(168)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(169)
		p.Match(DesignParserRBRACE)
	}

	return localctx
}

// IComponentDecalartionContext is an interface to support dynamic dispatch.
type IComponentDecalartionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsComponentDecalartionContext differentiates from other interfaces.
	IsComponentDecalartionContext()
}

type ComponentDecalartionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComponentDecalartionContext() *ComponentDecalartionContext {
	var p = new(ComponentDecalartionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DesignParserRULE_componentDecalartion
	return p
}

func (*ComponentDecalartionContext) IsComponentDecalartionContext() {}

func NewComponentDecalartionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComponentDecalartionContext {
	var p = new(ComponentDecalartionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesignParserRULE_componentDecalartion

	return p
}

func (s *ComponentDecalartionContext) GetParser() antlr.Parser { return s.parser }

func (s *ComponentDecalartionContext) COMPONENT() antlr.TerminalNode {
	return s.GetToken(DesignParserCOMPONENT, 0)
}

func (s *ComponentDecalartionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(DesignParserIDENTIFIER, 0)
}

func (s *ComponentDecalartionContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(DesignParserLBRACE, 0)
}

func (s *ComponentDecalartionContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(DesignParserRBRACE, 0)
}

func (s *ComponentDecalartionContext) AllComponentBodyDecalartion() []IComponentBodyDecalartionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IComponentBodyDecalartionContext)(nil)).Elem())
	var tst = make([]IComponentBodyDecalartionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IComponentBodyDecalartionContext)
		}
	}

	return tst
}

func (s *ComponentDecalartionContext) ComponentBodyDecalartion(i int) IComponentBodyDecalartionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComponentBodyDecalartionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IComponentBodyDecalartionContext)
}

func (s *ComponentDecalartionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComponentDecalartionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComponentDecalartionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.EnterComponentDecalartion(s)
	}
}

func (s *ComponentDecalartionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.ExitComponentDecalartion(s)
	}
}

func (s *ComponentDecalartionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DesignVisitor:
		return t.VisitComponentDecalartion(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DesignParser) ComponentDecalartion() (localctx IComponentDecalartionContext) {
	localctx = NewComponentDecalartionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, DesignParserRULE_componentDecalartion)
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
		p.SetState(171)
		p.Match(DesignParserCOMPONENT)
	}
	{
		p.SetState(172)
		p.Match(DesignParserIDENTIFIER)
	}
	{
		p.SetState(173)
		p.Match(DesignParserLBRACE)
	}
	p.SetState(177)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DesignParserIDENTIFIER {
		{
			p.SetState(174)
			p.ComponentBodyDecalartion()
		}

		p.SetState(179)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(180)
		p.Match(DesignParserRBRACE)
	}

	return localctx
}

// IComponentBodyDecalartionContext is an interface to support dynamic dispatch.
type IComponentBodyDecalartionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsComponentBodyDecalartionContext differentiates from other interfaces.
	IsComponentBodyDecalartionContext()
}

type ComponentBodyDecalartionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComponentBodyDecalartionContext() *ComponentBodyDecalartionContext {
	var p = new(ComponentBodyDecalartionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DesignParserRULE_componentBodyDecalartion
	return p
}

func (*ComponentBodyDecalartionContext) IsComponentBodyDecalartionContext() {}

func NewComponentBodyDecalartionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComponentBodyDecalartionContext {
	var p = new(ComponentBodyDecalartionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesignParserRULE_componentBodyDecalartion

	return p
}

func (s *ComponentBodyDecalartionContext) GetParser() antlr.Parser { return s.parser }

func (s *ComponentBodyDecalartionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(DesignParserIDENTIFIER, 0)
}

func (s *ComponentBodyDecalartionContext) COLON() antlr.TerminalNode {
	return s.GetToken(DesignParserCOLON, 0)
}

func (s *ComponentBodyDecalartionContext) ConfigValue() IConfigValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConfigValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConfigValueContext)
}

func (s *ComponentBodyDecalartionContext) LayoutDecalaration() ILayoutDecalarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILayoutDecalarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILayoutDecalarationContext)
}

func (s *ComponentBodyDecalartionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComponentBodyDecalartionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComponentBodyDecalartionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.EnterComponentBodyDecalartion(s)
	}
}

func (s *ComponentBodyDecalartionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.ExitComponentBodyDecalartion(s)
	}
}

func (s *ComponentBodyDecalartionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DesignVisitor:
		return t.VisitComponentBodyDecalartion(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DesignParser) ComponentBodyDecalartion() (localctx IComponentBodyDecalartionContext) {
	localctx = NewComponentBodyDecalartionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, DesignParserRULE_componentBodyDecalartion)

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
		p.SetState(182)
		p.Match(DesignParserIDENTIFIER)
	}
	p.SetState(186)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DesignParserCOLON:
		{
			p.SetState(183)
			p.Match(DesignParserCOLON)
		}
		{
			p.SetState(184)
			p.ConfigValue()
		}

	case DesignParserLBRACE:
		{
			p.SetState(185)
			p.LayoutDecalaration()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ILayoutDecalarationContext is an interface to support dynamic dispatch.
type ILayoutDecalarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLayoutDecalarationContext differentiates from other interfaces.
	IsLayoutDecalarationContext()
}

type LayoutDecalarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLayoutDecalarationContext() *LayoutDecalarationContext {
	var p = new(LayoutDecalarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DesignParserRULE_layoutDecalaration
	return p
}

func (*LayoutDecalarationContext) IsLayoutDecalarationContext() {}

func NewLayoutDecalarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LayoutDecalarationContext {
	var p = new(LayoutDecalarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesignParserRULE_layoutDecalaration

	return p
}

func (s *LayoutDecalarationContext) GetParser() antlr.Parser { return s.parser }

func (s *LayoutDecalarationContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(DesignParserLBRACE, 0)
}

func (s *LayoutDecalarationContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(DesignParserRBRACE, 0)
}

func (s *LayoutDecalarationContext) AllLayoutBodyDecalartion() []ILayoutBodyDecalartionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILayoutBodyDecalartionContext)(nil)).Elem())
	var tst = make([]ILayoutBodyDecalartionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILayoutBodyDecalartionContext)
		}
	}

	return tst
}

func (s *LayoutDecalarationContext) LayoutBodyDecalartion(i int) ILayoutBodyDecalartionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILayoutBodyDecalartionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILayoutBodyDecalartionContext)
}

func (s *LayoutDecalarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LayoutDecalarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LayoutDecalarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.EnterLayoutDecalaration(s)
	}
}

func (s *LayoutDecalarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.ExitLayoutDecalaration(s)
	}
}

func (s *LayoutDecalarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DesignVisitor:
		return t.VisitLayoutDecalaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DesignParser) LayoutDecalaration() (localctx ILayoutDecalarationContext) {
	localctx = NewLayoutDecalarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, DesignParserRULE_layoutDecalaration)
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
		p.SetState(188)
		p.Match(DesignParserLBRACE)
	}
	p.SetState(192)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DesignParserT__0 {
		{
			p.SetState(189)
			p.LayoutBodyDecalartion()
		}

		p.SetState(194)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(195)
		p.Match(DesignParserRBRACE)
	}

	return localctx
}

// ILayoutBodyDecalartionContext is an interface to support dynamic dispatch.
type ILayoutBodyDecalartionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLayoutBodyDecalartionContext differentiates from other interfaces.
	IsLayoutBodyDecalartionContext()
}

type LayoutBodyDecalartionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLayoutBodyDecalartionContext() *LayoutBodyDecalartionContext {
	var p = new(LayoutBodyDecalartionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DesignParserRULE_layoutBodyDecalartion
	return p
}

func (*LayoutBodyDecalartionContext) IsLayoutBodyDecalartionContext() {}

func NewLayoutBodyDecalartionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LayoutBodyDecalartionContext {
	var p = new(LayoutBodyDecalartionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesignParserRULE_layoutBodyDecalartion

	return p
}

func (s *LayoutBodyDecalartionContext) GetParser() antlr.Parser { return s.parser }

func (s *LayoutBodyDecalartionContext) EmptyLine() IEmptyLineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEmptyLineContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEmptyLineContext)
}

func (s *LayoutBodyDecalartionContext) LayoutLine() ILayoutLineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILayoutLineContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILayoutLineContext)
}

func (s *LayoutBodyDecalartionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LayoutBodyDecalartionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LayoutBodyDecalartionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.EnterLayoutBodyDecalartion(s)
	}
}

func (s *LayoutBodyDecalartionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.ExitLayoutBodyDecalartion(s)
	}
}

func (s *LayoutBodyDecalartionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DesignVisitor:
		return t.VisitLayoutBodyDecalartion(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DesignParser) LayoutBodyDecalartion() (localctx ILayoutBodyDecalartionContext) {
	localctx = NewLayoutBodyDecalartionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, DesignParserRULE_layoutBodyDecalartion)

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
		p.SetState(197)
		p.Match(DesignParserT__0)
	}
	p.SetState(200)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DesignParserT__1:
		{
			p.SetState(198)
			p.EmptyLine()
		}

	case DesignParserT__0, DesignParserGridSize, DesignParserSTRING_LITERAL, DesignParserRBRACE, DesignParserIDENTIFIER:
		{
			p.SetState(199)
			p.LayoutLine()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IEmptyLineContext is an interface to support dynamic dispatch.
type IEmptyLineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEmptyLineContext differentiates from other interfaces.
	IsEmptyLineContext()
}

type EmptyLineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEmptyLineContext() *EmptyLineContext {
	var p = new(EmptyLineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DesignParserRULE_emptyLine
	return p
}

func (*EmptyLineContext) IsEmptyLineContext() {}

func NewEmptyLineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EmptyLineContext {
	var p = new(EmptyLineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesignParserRULE_emptyLine

	return p
}

func (s *EmptyLineContext) GetParser() antlr.Parser { return s.parser }
func (s *EmptyLineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EmptyLineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EmptyLineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.EnterEmptyLine(s)
	}
}

func (s *EmptyLineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.ExitEmptyLine(s)
	}
}

func (s *EmptyLineContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DesignVisitor:
		return t.VisitEmptyLine(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DesignParser) EmptyLine() (localctx IEmptyLineContext) {
	localctx = NewEmptyLineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, DesignParserRULE_emptyLine)
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
		p.SetState(202)
		p.Match(DesignParserT__1)
	}
	p.SetState(206)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(203)
				_la = p.GetTokenStream().LA(1)

				if !(_la == DesignParserT__0 || _la == DesignParserT__1) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		p.SetState(208)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext())
	}

	return localctx
}

// ILayoutLineContext is an interface to support dynamic dispatch.
type ILayoutLineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLayoutLineContext differentiates from other interfaces.
	IsLayoutLineContext()
}

type LayoutLineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLayoutLineContext() *LayoutLineContext {
	var p = new(LayoutLineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DesignParserRULE_layoutLine
	return p
}

func (*LayoutLineContext) IsLayoutLineContext() {}

func NewLayoutLineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LayoutLineContext {
	var p = new(LayoutLineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesignParserRULE_layoutLine

	return p
}

func (s *LayoutLineContext) GetParser() antlr.Parser { return s.parser }

func (s *LayoutLineContext) AllComponentUseDeclaration() []IComponentUseDeclarationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IComponentUseDeclarationContext)(nil)).Elem())
	var tst = make([]IComponentUseDeclarationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IComponentUseDeclarationContext)
		}
	}

	return tst
}

func (s *LayoutLineContext) ComponentUseDeclaration(i int) IComponentUseDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComponentUseDeclarationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IComponentUseDeclarationContext)
}

func (s *LayoutLineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LayoutLineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LayoutLineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.EnterLayoutLine(s)
	}
}

func (s *LayoutLineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.ExitLayoutLine(s)
	}
}

func (s *LayoutLineContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DesignVisitor:
		return t.VisitLayoutLine(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DesignParser) LayoutLine() (localctx ILayoutLineContext) {
	localctx = NewLayoutLineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, DesignParserRULE_layoutLine)

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
	p.SetState(213)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(211)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case DesignParserT__0:
				{
					p.SetState(209)
					p.Match(DesignParserT__0)
				}

			case DesignParserGridSize, DesignParserSTRING_LITERAL, DesignParserIDENTIFIER:
				{
					p.SetState(210)
					p.ComponentUseDeclaration()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

		}
		p.SetState(215)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext())
	}

	return localctx
}

// IComponentUseDeclarationContext is an interface to support dynamic dispatch.
type IComponentUseDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsComponentUseDeclarationContext differentiates from other interfaces.
	IsComponentUseDeclarationContext()
}

type ComponentUseDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComponentUseDeclarationContext() *ComponentUseDeclarationContext {
	var p = new(ComponentUseDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DesignParserRULE_componentUseDeclaration
	return p
}

func (*ComponentUseDeclarationContext) IsComponentUseDeclarationContext() {}

func NewComponentUseDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComponentUseDeclarationContext {
	var p = new(ComponentUseDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesignParserRULE_componentUseDeclaration

	return p
}

func (s *ComponentUseDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *ComponentUseDeclarationContext) GridSize() antlr.TerminalNode {
	return s.GetToken(DesignParserGridSize, 0)
}

func (s *ComponentUseDeclarationContext) ComponentName() IComponentNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComponentNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComponentNameContext)
}

func (s *ComponentUseDeclarationContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(DesignParserLPAREN, 0)
}

func (s *ComponentUseDeclarationContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(DesignParserRPAREN, 0)
}

func (s *ComponentUseDeclarationContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(DesignParserSTRING_LITERAL, 0)
}

func (s *ComponentUseDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComponentUseDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComponentUseDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.EnterComponentUseDeclaration(s)
	}
}

func (s *ComponentUseDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.ExitComponentUseDeclaration(s)
	}
}

func (s *ComponentUseDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DesignVisitor:
		return t.VisitComponentUseDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DesignParser) ComponentUseDeclaration() (localctx IComponentUseDeclarationContext) {
	localctx = NewComponentUseDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, DesignParserRULE_componentUseDeclaration)
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

	p.SetState(224)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DesignParserGridSize:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(216)
			p.Match(DesignParserGridSize)
		}

	case DesignParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(217)
			p.ComponentName()
		}
		p.SetState(221)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DesignParserLPAREN {
			{
				p.SetState(218)
				p.Match(DesignParserLPAREN)
			}
			{
				p.SetState(219)
				_la = p.GetTokenStream().LA(1)

				if !(_la == DesignParserGridSize || _la == DesignParserSTRING_LITERAL) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(220)
				p.Match(DesignParserRPAREN)
			}

		}

	case DesignParserSTRING_LITERAL:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(223)
			p.Match(DesignParserSTRING_LITERAL)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IStyleDecalartionContext is an interface to support dynamic dispatch.
type IStyleDecalartionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStyleDecalartionContext differentiates from other interfaces.
	IsStyleDecalartionContext()
}

type StyleDecalartionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStyleDecalartionContext() *StyleDecalartionContext {
	var p = new(StyleDecalartionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DesignParserRULE_styleDecalartion
	return p
}

func (*StyleDecalartionContext) IsStyleDecalartionContext() {}

func NewStyleDecalartionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StyleDecalartionContext {
	var p = new(StyleDecalartionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesignParserRULE_styleDecalartion

	return p
}

func (s *StyleDecalartionContext) GetParser() antlr.Parser { return s.parser }

func (s *StyleDecalartionContext) STYLE() antlr.TerminalNode {
	return s.GetToken(DesignParserSTYLE, 0)
}

func (s *StyleDecalartionContext) StyleName() IStyleNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStyleNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStyleNameContext)
}

func (s *StyleDecalartionContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(DesignParserLBRACE, 0)
}

func (s *StyleDecalartionContext) StyleBody() IStyleBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStyleBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStyleBodyContext)
}

func (s *StyleDecalartionContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(DesignParserRBRACE, 0)
}

func (s *StyleDecalartionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StyleDecalartionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StyleDecalartionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.EnterStyleDecalartion(s)
	}
}

func (s *StyleDecalartionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.ExitStyleDecalartion(s)
	}
}

func (s *StyleDecalartionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DesignVisitor:
		return t.VisitStyleDecalartion(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DesignParser) StyleDecalartion() (localctx IStyleDecalartionContext) {
	localctx = NewStyleDecalartionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, DesignParserRULE_styleDecalartion)

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
		p.SetState(226)
		p.Match(DesignParserSTYLE)
	}
	{
		p.SetState(227)
		p.StyleName()
	}
	{
		p.SetState(228)
		p.Match(DesignParserLBRACE)
	}
	{
		p.SetState(229)
		p.StyleBody()
	}
	{
		p.SetState(230)
		p.Match(DesignParserRBRACE)
	}

	return localctx
}

// IStyleNameContext is an interface to support dynamic dispatch.
type IStyleNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStyleNameContext differentiates from other interfaces.
	IsStyleNameContext()
}

type StyleNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStyleNameContext() *StyleNameContext {
	var p = new(StyleNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DesignParserRULE_styleName
	return p
}

func (*StyleNameContext) IsStyleNameContext() {}

func NewStyleNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StyleNameContext {
	var p = new(StyleNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesignParserRULE_styleName

	return p
}

func (s *StyleNameContext) GetParser() antlr.Parser { return s.parser }

func (s *StyleNameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(DesignParserIDENTIFIER, 0)
}

func (s *StyleNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StyleNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StyleNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.EnterStyleName(s)
	}
}

func (s *StyleNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.ExitStyleName(s)
	}
}

func (s *StyleNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DesignVisitor:
		return t.VisitStyleName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DesignParser) StyleName() (localctx IStyleNameContext) {
	localctx = NewStyleNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, DesignParserRULE_styleName)

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
		p.SetState(232)
		p.Match(DesignParserIDENTIFIER)
	}

	return localctx
}

// IStyleBodyContext is an interface to support dynamic dispatch.
type IStyleBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStyleBodyContext differentiates from other interfaces.
	IsStyleBodyContext()
}

type StyleBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStyleBodyContext() *StyleBodyContext {
	var p = new(StyleBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DesignParserRULE_styleBody
	return p
}

func (*StyleBodyContext) IsStyleBodyContext() {}

func NewStyleBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StyleBodyContext {
	var p = new(StyleBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesignParserRULE_styleBody

	return p
}

func (s *StyleBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *StyleBodyContext) AllConfigDecalartion() []IConfigDecalartionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IConfigDecalartionContext)(nil)).Elem())
	var tst = make([]IConfigDecalartionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IConfigDecalartionContext)
		}
	}

	return tst
}

func (s *StyleBodyContext) ConfigDecalartion(i int) IConfigDecalartionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConfigDecalartionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IConfigDecalartionContext)
}

func (s *StyleBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StyleBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StyleBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.EnterStyleBody(s)
	}
}

func (s *StyleBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.ExitStyleBody(s)
	}
}

func (s *StyleBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DesignVisitor:
		return t.VisitStyleBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DesignParser) StyleBody() (localctx IStyleBodyContext) {
	localctx = NewStyleBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, DesignParserRULE_styleBody)
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
	p.SetState(239)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DesignParserIDENTIFIER {
		{
			p.SetState(234)
			p.ConfigDecalartion()
		}
		{
			p.SetState(235)
			p.Match(DesignParserT__2)
		}

		p.SetState(241)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ILibraryDecalartionContext is an interface to support dynamic dispatch.
type ILibraryDecalartionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLibraryDecalartionContext differentiates from other interfaces.
	IsLibraryDecalartionContext()
}

type LibraryDecalartionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLibraryDecalartionContext() *LibraryDecalartionContext {
	var p = new(LibraryDecalartionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DesignParserRULE_libraryDecalartion
	return p
}

func (*LibraryDecalartionContext) IsLibraryDecalartionContext() {}

func NewLibraryDecalartionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LibraryDecalartionContext {
	var p = new(LibraryDecalartionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesignParserRULE_libraryDecalartion

	return p
}

func (s *LibraryDecalartionContext) GetParser() antlr.Parser { return s.parser }

func (s *LibraryDecalartionContext) LIBRARAY() antlr.TerminalNode {
	return s.GetToken(DesignParserLIBRARAY, 0)
}

func (s *LibraryDecalartionContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(DesignParserLBRACE, 0)
}

func (s *LibraryDecalartionContext) LibraryBody() ILibraryBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILibraryBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILibraryBodyContext)
}

func (s *LibraryDecalartionContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(DesignParserRBRACE, 0)
}

func (s *LibraryDecalartionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LibraryDecalartionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LibraryDecalartionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.EnterLibraryDecalartion(s)
	}
}

func (s *LibraryDecalartionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.ExitLibraryDecalartion(s)
	}
}

func (s *LibraryDecalartionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DesignVisitor:
		return t.VisitLibraryDecalartion(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DesignParser) LibraryDecalartion() (localctx ILibraryDecalartionContext) {
	localctx = NewLibraryDecalartionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, DesignParserRULE_libraryDecalartion)

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
		p.SetState(242)
		p.Match(DesignParserLIBRARAY)
	}
	{
		p.SetState(243)
		p.Match(DesignParserLBRACE)
	}
	{
		p.SetState(244)
		p.LibraryBody()
	}
	{
		p.SetState(245)
		p.Match(DesignParserRBRACE)
	}

	return localctx
}

// ILibraryBodyContext is an interface to support dynamic dispatch.
type ILibraryBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLibraryBodyContext differentiates from other interfaces.
	IsLibraryBodyContext()
}

type LibraryBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLibraryBodyContext() *LibraryBodyContext {
	var p = new(LibraryBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DesignParserRULE_libraryBody
	return p
}

func (*LibraryBodyContext) IsLibraryBodyContext() {}

func NewLibraryBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LibraryBodyContext {
	var p = new(LibraryBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesignParserRULE_libraryBody

	return p
}

func (s *LibraryBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *LibraryBodyContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(DesignParserSTRING_LITERAL, 0)
}

func (s *LibraryBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LibraryBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LibraryBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.EnterLibraryBody(s)
	}
}

func (s *LibraryBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.ExitLibraryBody(s)
	}
}

func (s *LibraryBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DesignVisitor:
		return t.VisitLibraryBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DesignParser) LibraryBody() (localctx ILibraryBodyContext) {
	localctx = NewLibraryBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, DesignParserRULE_libraryBody)

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
		p.SetState(247)
		p.Match(DesignParserSTRING_LITERAL)
	}

	return localctx
}
