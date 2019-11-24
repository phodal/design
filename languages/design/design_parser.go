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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 45, 297,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 3, 2, 3, 2, 3, 2, 7, 2, 74, 10, 2, 12, 2, 14, 2, 77,
	11, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6,
	3, 6, 3, 6, 5, 6, 92, 10, 6, 3, 6, 3, 6, 3, 6, 5, 6, 97, 10, 6, 3, 7, 3,
	7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 106, 10, 7, 3, 8, 3, 8, 3, 8, 3,
	8, 7, 8, 112, 10, 8, 12, 8, 14, 8, 115, 11, 8, 3, 8, 3, 8, 3, 9, 3, 9,
	3, 9, 5, 9, 122, 10, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 129,
	10, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12,
	3, 12, 5, 12, 141, 10, 12, 3, 12, 3, 12, 3, 12, 5, 12, 146, 10, 12, 3,
	13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14,
	3, 14, 5, 14, 160, 10, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3,
	18, 3, 18, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 20, 7, 20, 176, 10, 20,
	12, 20, 14, 20, 179, 11, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 21,
	7, 21, 187, 10, 21, 12, 21, 14, 21, 190, 11, 21, 3, 21, 3, 21, 3, 22, 3,
	22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 24, 7, 24,
	205, 10, 24, 12, 24, 14, 24, 208, 11, 24, 3, 25, 3, 25, 7, 25, 212, 10,
	25, 12, 25, 14, 25, 215, 11, 25, 3, 25, 3, 25, 7, 25, 219, 10, 25, 12,
	25, 14, 25, 222, 11, 25, 3, 25, 3, 25, 5, 25, 226, 10, 25, 3, 26, 3, 26,
	3, 26, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 5, 27, 237, 10, 27, 3,
	27, 5, 27, 240, 10, 27, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 29,
	3, 29, 3, 30, 3, 30, 3, 30, 7, 30, 253, 10, 30, 12, 30, 14, 30, 256, 11,
	30, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 32, 7, 32, 265, 10, 32,
	12, 32, 14, 32, 268, 11, 32, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33,
	3, 33, 3, 33, 3, 33, 3, 33, 5, 33, 280, 10, 33, 3, 34, 3, 34, 3, 34, 3,
	34, 3, 34, 3, 34, 3, 34, 3, 34, 7, 34, 290, 10, 34, 12, 34, 14, 34, 293,
	11, 34, 3, 35, 3, 35, 3, 35, 2, 2, 36, 2, 4, 6, 8, 10, 12, 14, 16, 18,
	20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54,
	56, 58, 60, 62, 64, 66, 68, 2, 3, 5, 2, 18, 18, 23, 23, 43, 43, 2, 296,
	2, 75, 3, 2, 2, 2, 4, 80, 3, 2, 2, 2, 6, 82, 3, 2, 2, 2, 8, 86, 3, 2, 2,
	2, 10, 96, 3, 2, 2, 2, 12, 105, 3, 2, 2, 2, 14, 107, 3, 2, 2, 2, 16, 121,
	3, 2, 2, 2, 18, 123, 3, 2, 2, 2, 20, 130, 3, 2, 2, 2, 22, 138, 3, 2, 2,
	2, 24, 147, 3, 2, 2, 2, 26, 159, 3, 2, 2, 2, 28, 161, 3, 2, 2, 2, 30, 163,
	3, 2, 2, 2, 32, 165, 3, 2, 2, 2, 34, 167, 3, 2, 2, 2, 36, 169, 3, 2, 2,
	2, 38, 171, 3, 2, 2, 2, 40, 182, 3, 2, 2, 2, 42, 193, 3, 2, 2, 2, 44, 197,
	3, 2, 2, 2, 46, 206, 3, 2, 2, 2, 48, 225, 3, 2, 2, 2, 50, 227, 3, 2, 2,
	2, 52, 239, 3, 2, 2, 2, 54, 241, 3, 2, 2, 2, 56, 247, 3, 2, 2, 2, 58, 254,
	3, 2, 2, 2, 60, 257, 3, 2, 2, 2, 62, 266, 3, 2, 2, 2, 64, 279, 3, 2, 2,
	2, 66, 281, 3, 2, 2, 2, 68, 294, 3, 2, 2, 2, 70, 74, 5, 4, 3, 2, 71, 74,
	5, 6, 4, 2, 72, 74, 5, 12, 7, 2, 73, 70, 3, 2, 2, 2, 73, 71, 3, 2, 2, 2,
	73, 72, 3, 2, 2, 2, 74, 77, 3, 2, 2, 2, 75, 73, 3, 2, 2, 2, 75, 76, 3,
	2, 2, 2, 76, 78, 3, 2, 2, 2, 77, 75, 3, 2, 2, 2, 78, 79, 7, 2, 2, 3, 79,
	3, 3, 2, 2, 2, 80, 81, 7, 42, 2, 2, 81, 5, 3, 2, 2, 2, 82, 83, 5, 8, 5,
	2, 83, 84, 7, 38, 2, 2, 84, 85, 5, 10, 6, 2, 85, 7, 3, 2, 2, 2, 86, 87,
	7, 42, 2, 2, 87, 9, 3, 2, 2, 2, 88, 91, 7, 42, 2, 2, 89, 90, 7, 40, 2,
	2, 90, 92, 7, 42, 2, 2, 91, 89, 3, 2, 2, 2, 91, 92, 3, 2, 2, 2, 92, 97,
	3, 2, 2, 2, 93, 97, 7, 23, 2, 2, 94, 97, 7, 44, 2, 2, 95, 97, 7, 45, 2,
	2, 96, 88, 3, 2, 2, 2, 96, 93, 3, 2, 2, 2, 96, 94, 3, 2, 2, 2, 96, 95,
	3, 2, 2, 2, 97, 11, 3, 2, 2, 2, 98, 106, 5, 6, 4, 2, 99, 106, 5, 14, 8,
	2, 100, 106, 5, 38, 20, 2, 101, 106, 5, 54, 28, 2, 102, 106, 5, 40, 21,
	2, 103, 106, 5, 60, 31, 2, 104, 106, 5, 44, 23, 2, 105, 98, 3, 2, 2, 2,
	105, 99, 3, 2, 2, 2, 105, 100, 3, 2, 2, 2, 105, 101, 3, 2, 2, 2, 105, 102,
	3, 2, 2, 2, 105, 103, 3, 2, 2, 2, 105, 104, 3, 2, 2, 2, 106, 13, 3, 2,
	2, 2, 107, 108, 7, 9, 2, 2, 108, 109, 7, 42, 2, 2, 109, 113, 7, 32, 2,
	2, 110, 112, 5, 16, 9, 2, 111, 110, 3, 2, 2, 2, 112, 115, 3, 2, 2, 2, 113,
	111, 3, 2, 2, 2, 113, 114, 3, 2, 2, 2, 114, 116, 3, 2, 2, 2, 115, 113,
	3, 2, 2, 2, 116, 117, 7, 33, 2, 2, 117, 15, 3, 2, 2, 2, 118, 122, 5, 18,
	10, 2, 119, 122, 5, 20, 11, 2, 120, 122, 5, 22, 12, 2, 121, 118, 3, 2,
	2, 2, 121, 119, 3, 2, 2, 2, 121, 120, 3, 2, 2, 2, 122, 17, 3, 2, 2, 2,
	123, 128, 7, 10, 2, 2, 124, 129, 7, 42, 2, 2, 125, 126, 7, 23, 2, 2, 126,
	127, 7, 39, 2, 2, 127, 129, 5, 32, 17, 2, 128, 124, 3, 2, 2, 2, 128, 125,
	3, 2, 2, 2, 129, 19, 3, 2, 2, 2, 130, 131, 7, 11, 2, 2, 131, 132, 7, 34,
	2, 2, 132, 133, 5, 28, 15, 2, 133, 134, 7, 35, 2, 2, 134, 135, 7, 23, 2,
	2, 135, 136, 7, 39, 2, 2, 136, 137, 5, 32, 17, 2, 137, 21, 3, 2, 2, 2,
	138, 140, 7, 12, 2, 2, 139, 141, 5, 34, 18, 2, 140, 139, 3, 2, 2, 2, 140,
	141, 3, 2, 2, 2, 141, 142, 3, 2, 2, 2, 142, 143, 7, 38, 2, 2, 143, 145,
	5, 26, 14, 2, 144, 146, 5, 24, 13, 2, 145, 144, 3, 2, 2, 2, 145, 146, 3,
	2, 2, 2, 146, 23, 3, 2, 2, 2, 147, 148, 7, 13, 2, 2, 148, 149, 7, 14, 2,
	2, 149, 150, 7, 30, 2, 2, 150, 151, 5, 36, 19, 2, 151, 152, 7, 31, 2, 2,
	152, 25, 3, 2, 2, 2, 153, 154, 7, 7, 2, 2, 154, 160, 5, 32, 17, 2, 155,
	156, 7, 8, 2, 2, 156, 157, 7, 23, 2, 2, 157, 158, 7, 39, 2, 2, 158, 160,
	5, 32, 17, 2, 159, 153, 3, 2, 2, 2, 159, 155, 3, 2, 2, 2, 160, 27, 3, 2,
	2, 2, 161, 162, 7, 42, 2, 2, 162, 29, 3, 2, 2, 2, 163, 164, 7, 42, 2, 2,
	164, 31, 3, 2, 2, 2, 165, 166, 7, 42, 2, 2, 166, 33, 3, 2, 2, 2, 167, 168,
	7, 42, 2, 2, 168, 35, 3, 2, 2, 2, 169, 170, 7, 42, 2, 2, 170, 37, 3, 2,
	2, 2, 171, 172, 7, 19, 2, 2, 172, 173, 7, 42, 2, 2, 173, 177, 7, 32, 2,
	2, 174, 176, 5, 42, 22, 2, 175, 174, 3, 2, 2, 2, 176, 179, 3, 2, 2, 2,
	177, 175, 3, 2, 2, 2, 177, 178, 3, 2, 2, 2, 178, 180, 3, 2, 2, 2, 179,
	177, 3, 2, 2, 2, 180, 181, 7, 33, 2, 2, 181, 39, 3, 2, 2, 2, 182, 183,
	7, 20, 2, 2, 183, 184, 7, 42, 2, 2, 184, 188, 7, 32, 2, 2, 185, 187, 5,
	42, 22, 2, 186, 185, 3, 2, 2, 2, 187, 190, 3, 2, 2, 2, 188, 186, 3, 2,
	2, 2, 188, 189, 3, 2, 2, 2, 189, 191, 3, 2, 2, 2, 190, 188, 3, 2, 2, 2,
	191, 192, 7, 33, 2, 2, 192, 41, 3, 2, 2, 2, 193, 194, 7, 42, 2, 2, 194,
	195, 7, 38, 2, 2, 195, 196, 5, 10, 6, 2, 196, 43, 3, 2, 2, 2, 197, 198,
	7, 17, 2, 2, 198, 199, 7, 42, 2, 2, 199, 200, 7, 32, 2, 2, 200, 201, 5,
	46, 24, 2, 201, 202, 7, 33, 2, 2, 202, 45, 3, 2, 2, 2, 203, 205, 5, 48,
	25, 2, 204, 203, 3, 2, 2, 2, 205, 208, 3, 2, 2, 2, 206, 204, 3, 2, 2, 2,
	206, 207, 3, 2, 2, 2, 207, 47, 3, 2, 2, 2, 208, 206, 3, 2, 2, 2, 209, 213,
	7, 3, 2, 2, 210, 212, 7, 3, 2, 2, 211, 210, 3, 2, 2, 2, 212, 215, 3, 2,
	2, 2, 213, 211, 3, 2, 2, 2, 213, 214, 3, 2, 2, 2, 214, 226, 3, 2, 2, 2,
	215, 213, 3, 2, 2, 2, 216, 220, 5, 50, 26, 2, 217, 219, 5, 50, 26, 2, 218,
	217, 3, 2, 2, 2, 219, 222, 3, 2, 2, 2, 220, 218, 3, 2, 2, 2, 220, 221,
	3, 2, 2, 2, 221, 223, 3, 2, 2, 2, 222, 220, 3, 2, 2, 2, 223, 224, 7, 4,
	2, 2, 224, 226, 3, 2, 2, 2, 225, 209, 3, 2, 2, 2, 225, 216, 3, 2, 2, 2,
	226, 49, 3, 2, 2, 2, 227, 228, 7, 4, 2, 2, 228, 229, 5, 52, 27, 2, 229,
	51, 3, 2, 2, 2, 230, 240, 7, 45, 2, 2, 231, 240, 7, 18, 2, 2, 232, 236,
	5, 32, 17, 2, 233, 234, 7, 30, 2, 2, 234, 235, 9, 2, 2, 2, 235, 237, 7,
	31, 2, 2, 236, 233, 3, 2, 2, 2, 236, 237, 3, 2, 2, 2, 237, 240, 3, 2, 2,
	2, 238, 240, 7, 23, 2, 2, 239, 230, 3, 2, 2, 2, 239, 231, 3, 2, 2, 2, 239,
	232, 3, 2, 2, 2, 239, 238, 3, 2, 2, 2, 240, 53, 3, 2, 2, 2, 241, 242, 7,
	21, 2, 2, 242, 243, 5, 56, 29, 2, 243, 244, 7, 32, 2, 2, 244, 245, 5, 58,
	30, 2, 245, 246, 7, 33, 2, 2, 246, 55, 3, 2, 2, 2, 247, 248, 7, 42, 2,
	2, 248, 57, 3, 2, 2, 2, 249, 250, 5, 6, 4, 2, 250, 251, 7, 5, 2, 2, 251,
	253, 3, 2, 2, 2, 252, 249, 3, 2, 2, 2, 253, 256, 3, 2, 2, 2, 254, 252,
	3, 2, 2, 2, 254, 255, 3, 2, 2, 2, 255, 59, 3, 2, 2, 2, 256, 254, 3, 2,
	2, 2, 257, 258, 7, 22, 2, 2, 258, 259, 5, 68, 35, 2, 259, 260, 7, 32, 2,
	2, 260, 261, 5, 62, 32, 2, 261, 262, 7, 33, 2, 2, 262, 61, 3, 2, 2, 2,
	263, 265, 5, 64, 33, 2, 264, 263, 3, 2, 2, 2, 265, 268, 3, 2, 2, 2, 266,
	264, 3, 2, 2, 2, 266, 267, 3, 2, 2, 2, 267, 63, 3, 2, 2, 2, 268, 266, 3,
	2, 2, 2, 269, 270, 5, 8, 5, 2, 270, 271, 7, 6, 2, 2, 271, 272, 5, 10, 6,
	2, 272, 273, 7, 5, 2, 2, 273, 280, 3, 2, 2, 2, 274, 275, 7, 42, 2, 2, 275,
	276, 7, 34, 2, 2, 276, 277, 5, 66, 34, 2, 277, 278, 7, 35, 2, 2, 278, 280,
	3, 2, 2, 2, 279, 269, 3, 2, 2, 2, 279, 274, 3, 2, 2, 2, 280, 65, 3, 2,
	2, 2, 281, 282, 5, 68, 35, 2, 282, 283, 7, 39, 2, 2, 283, 291, 7, 42, 2,
	2, 284, 285, 7, 40, 2, 2, 285, 286, 5, 68, 35, 2, 286, 287, 7, 39, 2, 2,
	287, 288, 7, 42, 2, 2, 288, 290, 3, 2, 2, 2, 289, 284, 3, 2, 2, 2, 290,
	293, 3, 2, 2, 2, 291, 289, 3, 2, 2, 2, 291, 292, 3, 2, 2, 2, 292, 67, 3,
	2, 2, 2, 293, 291, 3, 2, 2, 2, 294, 295, 7, 42, 2, 2, 295, 69, 3, 2, 2,
	2, 25, 73, 75, 91, 96, 105, 113, 121, 128, 140, 145, 159, 177, 188, 206,
	213, 220, 225, 236, 239, 254, 266, 279, 291,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'-'", "'|'", "';'", "'='", "", "", "", "", "", "", "", "", "'repeat'",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "'('", "')'", "'{'",
	"'}'", "'['", "']'", "'\"'", "'''", "':'", "'.'", "','",
}
var symbolicNames = []string{
	"", "", "", "", "", "GOTO_KEY", "SHOW_KEY", "FLOW", "SEE", "DO", "REACT",
	"WITHTEXT", "ANIMATE", "REPEAT", "REPEAT_TIMES", "LAYOUT", "POSITION",
	"PAGE", "COMPONENT", "STYLE", "LIBRARY", "STRING_LITERAL", "WS", "COMMENT",
	"LINE_COMMENT", "EmptyLine", "Space", "NewLine", "LPAREN", "RPAREN", "LBRACE",
	"RBRACE", "LBRACK", "RBRACK", "Quote", "SingleQuote", "COLON", "DOT", "COMMA",
	"LETTER", "IDENTIFIER", "DIGITS", "DIGITS_IDENTIFIER", "DECIMAL_LITERAL",
}

var ruleNames = []string{
	"start", "comment", "configDecalartion", "configKey", "configValue", "decalartions",
	"flowDecalartion", "flowBodyDecalartion", "seeDecalartion", "doDecalartion",
	"reactDecalartion", "animateDecalartion", "actionKey", "actionName", "componentValue",
	"componentName", "sceneName", "animateName", "pageDecalartion", "componentDecalartion",
	"componentBodyDecalartion", "layoutDecalaration", "layoutBodyDecalartion",
	"layoutRow", "layoutLine", "componentUseDeclaration", "styleDecalartion",
	"styleName", "styleBody", "libraryDecalartion", "libraryBody", "express",
	"libraryCall", "libraryName",
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
	DesignParserT__3              = 4
	DesignParserGOTO_KEY          = 5
	DesignParserSHOW_KEY          = 6
	DesignParserFLOW              = 7
	DesignParserSEE               = 8
	DesignParserDO                = 9
	DesignParserREACT             = 10
	DesignParserWITHTEXT          = 11
	DesignParserANIMATE           = 12
	DesignParserREPEAT            = 13
	DesignParserREPEAT_TIMES      = 14
	DesignParserLAYOUT            = 15
	DesignParserPOSITION          = 16
	DesignParserPAGE              = 17
	DesignParserCOMPONENT         = 18
	DesignParserSTYLE             = 19
	DesignParserLIBRARY           = 20
	DesignParserSTRING_LITERAL    = 21
	DesignParserWS                = 22
	DesignParserCOMMENT           = 23
	DesignParserLINE_COMMENT      = 24
	DesignParserEmptyLine         = 25
	DesignParserSpace             = 26
	DesignParserNewLine           = 27
	DesignParserLPAREN            = 28
	DesignParserRPAREN            = 29
	DesignParserLBRACE            = 30
	DesignParserRBRACE            = 31
	DesignParserLBRACK            = 32
	DesignParserRBRACK            = 33
	DesignParserQuote             = 34
	DesignParserSingleQuote       = 35
	DesignParserCOLON             = 36
	DesignParserDOT               = 37
	DesignParserCOMMA             = 38
	DesignParserLETTER            = 39
	DesignParserIDENTIFIER        = 40
	DesignParserDIGITS            = 41
	DesignParserDIGITS_IDENTIFIER = 42
	DesignParserDECIMAL_LITERAL   = 43
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
	DesignParserRULE_layoutRow                = 23
	DesignParserRULE_layoutLine               = 24
	DesignParserRULE_componentUseDeclaration  = 25
	DesignParserRULE_styleDecalartion         = 26
	DesignParserRULE_styleName                = 27
	DesignParserRULE_styleBody                = 28
	DesignParserRULE_libraryDecalartion       = 29
	DesignParserRULE_libraryBody              = 30
	DesignParserRULE_express                  = 31
	DesignParserRULE_libraryCall              = 32
	DesignParserRULE_libraryName              = 33
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
	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<DesignParserFLOW)|(1<<DesignParserLAYOUT)|(1<<DesignParserPAGE)|(1<<DesignParserCOMPONENT)|(1<<DesignParserSTYLE)|(1<<DesignParserLIBRARY))) != 0) || _la == DesignParserIDENTIFIER {
		p.SetState(71)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(68)
				p.Comment()
			}

		case 2:
			{
				p.SetState(69)
				p.ConfigDecalartion()
			}

		case 3:
			{
				p.SetState(70)
				p.Decalartions()
			}

		}

		p.SetState(75)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(76)
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
		p.SetState(78)
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
		p.SetState(80)
		p.ConfigKey()
	}
	{
		p.SetState(81)
		p.Match(DesignParserCOLON)
	}
	{
		p.SetState(82)
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
		p.SetState(84)
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

func (s *ConfigValueContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(DesignParserSTRING_LITERAL, 0)
}

func (s *ConfigValueContext) DIGITS_IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(DesignParserDIGITS_IDENTIFIER, 0)
}

func (s *ConfigValueContext) DECIMAL_LITERAL() antlr.TerminalNode {
	return s.GetToken(DesignParserDECIMAL_LITERAL, 0)
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

	p.SetState(94)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DesignParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(86)
			p.Match(DesignParserIDENTIFIER)
		}
		p.SetState(89)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DesignParserCOMMA {
			{
				p.SetState(87)
				p.Match(DesignParserCOMMA)
			}
			{
				p.SetState(88)
				p.Match(DesignParserIDENTIFIER)
			}

		}

	case DesignParserSTRING_LITERAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(91)
			p.Match(DesignParserSTRING_LITERAL)
		}

	case DesignParserDIGITS_IDENTIFIER:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(92)
			p.Match(DesignParserDIGITS_IDENTIFIER)
		}

	case DesignParserDECIMAL_LITERAL:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(93)
			p.Match(DesignParserDECIMAL_LITERAL)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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

func (s *DecalartionsContext) LayoutDecalaration() ILayoutDecalarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILayoutDecalarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILayoutDecalarationContext)
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

	p.SetState(103)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DesignParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(96)
			p.ConfigDecalartion()
		}

	case DesignParserFLOW:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(97)
			p.FlowDecalartion()
		}

	case DesignParserPAGE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(98)
			p.PageDecalartion()
		}

	case DesignParserSTYLE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(99)
			p.StyleDecalartion()
		}

	case DesignParserCOMPONENT:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(100)
			p.ComponentDecalartion()
		}

	case DesignParserLIBRARY:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(101)
			p.LibraryDecalartion()
		}

	case DesignParserLAYOUT:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(102)
			p.LayoutDecalaration()
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
		p.SetState(105)
		p.Match(DesignParserFLOW)
	}
	{
		p.SetState(106)
		p.Match(DesignParserIDENTIFIER)
	}
	{
		p.SetState(107)
		p.Match(DesignParserLBRACE)
	}
	p.SetState(111)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<DesignParserSEE)|(1<<DesignParserDO)|(1<<DesignParserREACT))) != 0 {
		{
			p.SetState(108)
			p.FlowBodyDecalartion()
		}

		p.SetState(113)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(114)
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

	p.SetState(119)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DesignParserSEE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(116)
			p.SeeDecalartion()
		}

	case DesignParserDO:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(117)
			p.DoDecalartion()
		}

	case DesignParserREACT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(118)
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
		p.SetState(121)
		p.Match(DesignParserSEE)
	}
	p.SetState(126)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DesignParserIDENTIFIER:
		{
			p.SetState(122)
			p.Match(DesignParserIDENTIFIER)
		}

	case DesignParserSTRING_LITERAL:
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
		p.SetState(128)
		p.Match(DesignParserDO)
	}
	{
		p.SetState(129)
		p.Match(DesignParserLBRACK)
	}
	{
		p.SetState(130)
		p.ActionName()
	}
	{
		p.SetState(131)
		p.Match(DesignParserRBRACK)
	}
	{
		p.SetState(132)
		p.Match(DesignParserSTRING_LITERAL)
	}
	{
		p.SetState(133)
		p.Match(DesignParserDOT)
	}
	{
		p.SetState(134)
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
		p.SetState(136)
		p.Match(DesignParserREACT)
	}
	p.SetState(138)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DesignParserIDENTIFIER {
		{
			p.SetState(137)
			p.SceneName()
		}

	}
	{
		p.SetState(140)
		p.Match(DesignParserCOLON)
	}
	{
		p.SetState(141)
		p.ActionKey()
	}
	p.SetState(143)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DesignParserWITHTEXT {
		{
			p.SetState(142)
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
		p.SetState(145)
		p.Match(DesignParserWITHTEXT)
	}
	{
		p.SetState(146)
		p.Match(DesignParserANIMATE)
	}
	{
		p.SetState(147)
		p.Match(DesignParserLPAREN)
	}
	{
		p.SetState(148)
		p.AnimateName()
	}
	{
		p.SetState(149)
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

	p.SetState(157)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DesignParserGOTO_KEY:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(151)
			p.Match(DesignParserGOTO_KEY)
		}
		{
			p.SetState(152)
			p.ComponentName()
		}

	case DesignParserSHOW_KEY:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(153)
			p.Match(DesignParserSHOW_KEY)
		}
		{
			p.SetState(154)
			p.Match(DesignParserSTRING_LITERAL)
		}
		{
			p.SetState(155)
			p.Match(DesignParserDOT)
		}
		{
			p.SetState(156)
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
		p.SetState(159)
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
		p.SetState(161)
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
		p.SetState(163)
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
		p.SetState(165)
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
		p.SetState(167)
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
		p.SetState(169)
		p.Match(DesignParserPAGE)
	}
	{
		p.SetState(170)
		p.Match(DesignParserIDENTIFIER)
	}
	{
		p.SetState(171)
		p.Match(DesignParserLBRACE)
	}
	p.SetState(175)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DesignParserIDENTIFIER {
		{
			p.SetState(172)
			p.ComponentBodyDecalartion()
		}

		p.SetState(177)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(178)
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
		p.SetState(180)
		p.Match(DesignParserCOMPONENT)
	}
	{
		p.SetState(181)
		p.Match(DesignParserIDENTIFIER)
	}
	{
		p.SetState(182)
		p.Match(DesignParserLBRACE)
	}
	p.SetState(186)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DesignParserIDENTIFIER {
		{
			p.SetState(183)
			p.ComponentBodyDecalartion()
		}

		p.SetState(188)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(189)
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
		p.SetState(191)
		p.Match(DesignParserIDENTIFIER)
	}
	{
		p.SetState(192)
		p.Match(DesignParserCOLON)
	}
	{
		p.SetState(193)
		p.ConfigValue()
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

func (s *LayoutDecalarationContext) LAYOUT() antlr.TerminalNode {
	return s.GetToken(DesignParserLAYOUT, 0)
}

func (s *LayoutDecalarationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(DesignParserIDENTIFIER, 0)
}

func (s *LayoutDecalarationContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(DesignParserLBRACE, 0)
}

func (s *LayoutDecalarationContext) LayoutBodyDecalartion() ILayoutBodyDecalartionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILayoutBodyDecalartionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILayoutBodyDecalartionContext)
}

func (s *LayoutDecalarationContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(DesignParserRBRACE, 0)
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
		p.SetState(195)
		p.Match(DesignParserLAYOUT)
	}
	{
		p.SetState(196)
		p.Match(DesignParserIDENTIFIER)
	}
	{
		p.SetState(197)
		p.Match(DesignParserLBRACE)
	}
	{
		p.SetState(198)
		p.LayoutBodyDecalartion()
	}
	{
		p.SetState(199)
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

func (s *LayoutBodyDecalartionContext) AllLayoutRow() []ILayoutRowContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILayoutRowContext)(nil)).Elem())
	var tst = make([]ILayoutRowContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILayoutRowContext)
		}
	}

	return tst
}

func (s *LayoutBodyDecalartionContext) LayoutRow(i int) ILayoutRowContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILayoutRowContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILayoutRowContext)
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
	p.SetState(204)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DesignParserT__0 || _la == DesignParserT__1 {
		{
			p.SetState(201)
			p.LayoutRow()
		}

		p.SetState(206)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ILayoutRowContext is an interface to support dynamic dispatch.
type ILayoutRowContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLayoutRowContext differentiates from other interfaces.
	IsLayoutRowContext()
}

type LayoutRowContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLayoutRowContext() *LayoutRowContext {
	var p = new(LayoutRowContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DesignParserRULE_layoutRow
	return p
}

func (*LayoutRowContext) IsLayoutRowContext() {}

func NewLayoutRowContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LayoutRowContext {
	var p = new(LayoutRowContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesignParserRULE_layoutRow

	return p
}

func (s *LayoutRowContext) GetParser() antlr.Parser { return s.parser }

func (s *LayoutRowContext) AllLayoutLine() []ILayoutLineContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILayoutLineContext)(nil)).Elem())
	var tst = make([]ILayoutLineContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILayoutLineContext)
		}
	}

	return tst
}

func (s *LayoutRowContext) LayoutLine(i int) ILayoutLineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILayoutLineContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILayoutLineContext)
}

func (s *LayoutRowContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LayoutRowContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LayoutRowContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.EnterLayoutRow(s)
	}
}

func (s *LayoutRowContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.ExitLayoutRow(s)
	}
}

func (s *LayoutRowContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DesignVisitor:
		return t.VisitLayoutRow(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DesignParser) LayoutRow() (localctx ILayoutRowContext) {
	localctx = NewLayoutRowContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, DesignParserRULE_layoutRow)

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

	p.SetState(223)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DesignParserT__0:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(207)
			p.Match(DesignParserT__0)
		}
		p.SetState(211)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(208)
					p.Match(DesignParserT__0)
				}

			}
			p.SetState(213)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext())
		}

	case DesignParserT__1:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(214)
			p.LayoutLine()
		}
		p.SetState(218)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(215)
					p.LayoutLine()
				}

			}
			p.SetState(220)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())
		}
		{
			p.SetState(221)
			p.Match(DesignParserT__1)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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

func (s *LayoutLineContext) ComponentUseDeclaration() IComponentUseDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComponentUseDeclarationContext)(nil)).Elem(), 0)

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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(225)
		p.Match(DesignParserT__1)
	}
	{
		p.SetState(226)
		p.ComponentUseDeclaration()
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

func (s *ComponentUseDeclarationContext) DECIMAL_LITERAL() antlr.TerminalNode {
	return s.GetToken(DesignParserDECIMAL_LITERAL, 0)
}

func (s *ComponentUseDeclarationContext) POSITION() antlr.TerminalNode {
	return s.GetToken(DesignParserPOSITION, 0)
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

func (s *ComponentUseDeclarationContext) DIGITS() antlr.TerminalNode {
	return s.GetToken(DesignParserDIGITS, 0)
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

	p.SetState(237)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DesignParserDECIMAL_LITERAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(228)
			p.Match(DesignParserDECIMAL_LITERAL)
		}

	case DesignParserPOSITION:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(229)
			p.Match(DesignParserPOSITION)
		}

	case DesignParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(230)
			p.ComponentName()
		}
		p.SetState(234)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DesignParserLPAREN {
			{
				p.SetState(231)
				p.Match(DesignParserLPAREN)
			}
			{
				p.SetState(232)
				_la = p.GetTokenStream().LA(1)

				if !(((_la-16)&-(0x1f+1)) == 0 && ((1<<uint((_la-16)))&((1<<(DesignParserPOSITION-16))|(1<<(DesignParserSTRING_LITERAL-16))|(1<<(DesignParserDIGITS-16)))) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(233)
				p.Match(DesignParserRPAREN)
			}

		}

	case DesignParserSTRING_LITERAL:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(236)
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
		p.SetState(239)
		p.Match(DesignParserSTYLE)
	}
	{
		p.SetState(240)
		p.StyleName()
	}
	{
		p.SetState(241)
		p.Match(DesignParserLBRACE)
	}
	{
		p.SetState(242)
		p.StyleBody()
	}
	{
		p.SetState(243)
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
		p.SetState(245)
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
	p.SetState(252)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DesignParserIDENTIFIER {
		{
			p.SetState(247)
			p.ConfigDecalartion()
		}
		{
			p.SetState(248)
			p.Match(DesignParserT__2)
		}

		p.SetState(254)
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

func (s *LibraryDecalartionContext) LIBRARY() antlr.TerminalNode {
	return s.GetToken(DesignParserLIBRARY, 0)
}

func (s *LibraryDecalartionContext) LibraryName() ILibraryNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILibraryNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILibraryNameContext)
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
		p.SetState(255)
		p.Match(DesignParserLIBRARY)
	}
	{
		p.SetState(256)
		p.LibraryName()
	}
	{
		p.SetState(257)
		p.Match(DesignParserLBRACE)
	}
	{
		p.SetState(258)
		p.LibraryBody()
	}
	{
		p.SetState(259)
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

func (s *LibraryBodyContext) AllExpress() []IExpressContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressContext)(nil)).Elem())
	var tst = make([]IExpressContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressContext)
		}
	}

	return tst
}

func (s *LibraryBodyContext) Express(i int) IExpressContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressContext)
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
	p.SetState(264)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DesignParserIDENTIFIER {
		{
			p.SetState(261)
			p.Express()
		}

		p.SetState(266)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IExpressContext is an interface to support dynamic dispatch.
type IExpressContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressContext differentiates from other interfaces.
	IsExpressContext()
}

type ExpressContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressContext() *ExpressContext {
	var p = new(ExpressContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DesignParserRULE_express
	return p
}

func (*ExpressContext) IsExpressContext() {}

func NewExpressContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressContext {
	var p = new(ExpressContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesignParserRULE_express

	return p
}

func (s *ExpressContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressContext) ConfigKey() IConfigKeyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConfigKeyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConfigKeyContext)
}

func (s *ExpressContext) ConfigValue() IConfigValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConfigValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConfigValueContext)
}

func (s *ExpressContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(DesignParserIDENTIFIER, 0)
}

func (s *ExpressContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(DesignParserLBRACK, 0)
}

func (s *ExpressContext) LibraryCall() ILibraryCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILibraryCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILibraryCallContext)
}

func (s *ExpressContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(DesignParserRBRACK, 0)
}

func (s *ExpressContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.EnterExpress(s)
	}
}

func (s *ExpressContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.ExitExpress(s)
	}
}

func (s *ExpressContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DesignVisitor:
		return t.VisitExpress(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DesignParser) Express() (localctx IExpressContext) {
	localctx = NewExpressContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, DesignParserRULE_express)

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

	p.SetState(277)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(267)
			p.ConfigKey()
		}
		{
			p.SetState(268)
			p.Match(DesignParserT__3)
		}
		{
			p.SetState(269)
			p.ConfigValue()
		}
		{
			p.SetState(270)
			p.Match(DesignParserT__2)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(272)
			p.Match(DesignParserIDENTIFIER)
		}
		{
			p.SetState(273)
			p.Match(DesignParserLBRACK)
		}
		{
			p.SetState(274)
			p.LibraryCall()
		}
		{
			p.SetState(275)
			p.Match(DesignParserRBRACK)
		}

	}

	return localctx
}

// ILibraryCallContext is an interface to support dynamic dispatch.
type ILibraryCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLibraryCallContext differentiates from other interfaces.
	IsLibraryCallContext()
}

type LibraryCallContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLibraryCallContext() *LibraryCallContext {
	var p = new(LibraryCallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DesignParserRULE_libraryCall
	return p
}

func (*LibraryCallContext) IsLibraryCallContext() {}

func NewLibraryCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LibraryCallContext {
	var p = new(LibraryCallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesignParserRULE_libraryCall

	return p
}

func (s *LibraryCallContext) GetParser() antlr.Parser { return s.parser }

func (s *LibraryCallContext) AllLibraryName() []ILibraryNameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILibraryNameContext)(nil)).Elem())
	var tst = make([]ILibraryNameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILibraryNameContext)
		}
	}

	return tst
}

func (s *LibraryCallContext) LibraryName(i int) ILibraryNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILibraryNameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILibraryNameContext)
}

func (s *LibraryCallContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(DesignParserDOT)
}

func (s *LibraryCallContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(DesignParserDOT, i)
}

func (s *LibraryCallContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(DesignParserIDENTIFIER)
}

func (s *LibraryCallContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(DesignParserIDENTIFIER, i)
}

func (s *LibraryCallContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(DesignParserCOMMA)
}

func (s *LibraryCallContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(DesignParserCOMMA, i)
}

func (s *LibraryCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LibraryCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LibraryCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.EnterLibraryCall(s)
	}
}

func (s *LibraryCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.ExitLibraryCall(s)
	}
}

func (s *LibraryCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DesignVisitor:
		return t.VisitLibraryCall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DesignParser) LibraryCall() (localctx ILibraryCallContext) {
	localctx = NewLibraryCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, DesignParserRULE_libraryCall)
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
		p.SetState(279)
		p.LibraryName()
	}
	{
		p.SetState(280)
		p.Match(DesignParserDOT)
	}
	{
		p.SetState(281)
		p.Match(DesignParserIDENTIFIER)
	}
	p.SetState(289)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DesignParserCOMMA {
		{
			p.SetState(282)
			p.Match(DesignParserCOMMA)
		}
		{
			p.SetState(283)
			p.LibraryName()
		}
		{
			p.SetState(284)
			p.Match(DesignParserDOT)
		}
		{
			p.SetState(285)
			p.Match(DesignParserIDENTIFIER)
		}

		p.SetState(291)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ILibraryNameContext is an interface to support dynamic dispatch.
type ILibraryNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLibraryNameContext differentiates from other interfaces.
	IsLibraryNameContext()
}

type LibraryNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLibraryNameContext() *LibraryNameContext {
	var p = new(LibraryNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DesignParserRULE_libraryName
	return p
}

func (*LibraryNameContext) IsLibraryNameContext() {}

func NewLibraryNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LibraryNameContext {
	var p = new(LibraryNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesignParserRULE_libraryName

	return p
}

func (s *LibraryNameContext) GetParser() antlr.Parser { return s.parser }

func (s *LibraryNameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(DesignParserIDENTIFIER, 0)
}

func (s *LibraryNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LibraryNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LibraryNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.EnterLibraryName(s)
	}
}

func (s *LibraryNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.ExitLibraryName(s)
	}
}

func (s *LibraryNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DesignVisitor:
		return t.VisitLibraryName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DesignParser) LibraryName() (localctx ILibraryNameContext) {
	localctx = NewLibraryNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, DesignParserRULE_libraryName)

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
		p.SetState(292)
		p.Match(DesignParserIDENTIFIER)
	}

	return localctx
}
