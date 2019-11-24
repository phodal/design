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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 49, 311,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 3, 2, 3, 2, 3, 2, 7, 2, 76, 10, 2, 12,
	2, 14, 2, 79, 11, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	5, 3, 5, 3, 6, 3, 6, 3, 6, 5, 6, 94, 10, 6, 3, 6, 3, 6, 3, 6, 5, 6, 99,
	10, 6, 3, 6, 3, 6, 5, 6, 103, 10, 6, 3, 6, 3, 6, 5, 6, 107, 10, 6, 5, 6,
	109, 10, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8,
	120, 10, 8, 3, 9, 3, 9, 3, 9, 3, 9, 7, 9, 126, 10, 9, 12, 9, 14, 9, 129,
	11, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 5, 10, 136, 10, 10, 3, 11, 3, 11,
	3, 11, 3, 11, 3, 11, 5, 11, 143, 10, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 5, 13, 155, 10, 13, 3, 13, 3, 13,
	3, 13, 5, 13, 160, 10, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 5, 15, 174, 10, 15, 3, 16, 3, 16,
	3, 17, 3, 17, 3, 18, 3, 18, 3, 19, 3, 19, 3, 20, 3, 20, 3, 21, 3, 21, 3,
	21, 3, 21, 7, 21, 190, 10, 21, 12, 21, 14, 21, 193, 11, 21, 3, 21, 3, 21,
	3, 22, 3, 22, 3, 22, 3, 22, 7, 22, 201, 10, 22, 12, 22, 14, 22, 204, 11,
	22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 3, 24,
	3, 24, 3, 24, 3, 25, 7, 25, 219, 10, 25, 12, 25, 14, 25, 222, 11, 25, 3,
	26, 3, 26, 7, 26, 226, 10, 26, 12, 26, 14, 26, 229, 11, 26, 3, 26, 3, 26,
	7, 26, 233, 10, 26, 12, 26, 14, 26, 236, 11, 26, 3, 26, 3, 26, 5, 26, 240,
	10, 26, 3, 27, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28,
	5, 28, 251, 10, 28, 3, 28, 5, 28, 254, 10, 28, 3, 29, 3, 29, 3, 29, 3,
	29, 3, 29, 3, 29, 3, 30, 3, 30, 3, 31, 3, 31, 3, 31, 7, 31, 267, 10, 31,
	12, 31, 14, 31, 270, 11, 31, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32,
	3, 33, 7, 33, 279, 10, 33, 12, 33, 14, 33, 282, 11, 33, 3, 34, 3, 34, 3,
	34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 5, 34, 294, 10, 34,
	3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 7, 35, 304, 10,
	35, 12, 35, 14, 35, 307, 11, 35, 3, 36, 3, 36, 3, 36, 2, 2, 37, 2, 4, 6,
	8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42,
	44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 2, 4, 3, 2, 3,
	5, 5, 2, 21, 21, 26, 26, 46, 46, 2, 313, 2, 77, 3, 2, 2, 2, 4, 82, 3, 2,
	2, 2, 6, 84, 3, 2, 2, 2, 8, 88, 3, 2, 2, 2, 10, 108, 3, 2, 2, 2, 12, 110,
	3, 2, 2, 2, 14, 119, 3, 2, 2, 2, 16, 121, 3, 2, 2, 2, 18, 135, 3, 2, 2,
	2, 20, 137, 3, 2, 2, 2, 22, 144, 3, 2, 2, 2, 24, 152, 3, 2, 2, 2, 26, 161,
	3, 2, 2, 2, 28, 173, 3, 2, 2, 2, 30, 175, 3, 2, 2, 2, 32, 177, 3, 2, 2,
	2, 34, 179, 3, 2, 2, 2, 36, 181, 3, 2, 2, 2, 38, 183, 3, 2, 2, 2, 40, 185,
	3, 2, 2, 2, 42, 196, 3, 2, 2, 2, 44, 207, 3, 2, 2, 2, 46, 211, 3, 2, 2,
	2, 48, 220, 3, 2, 2, 2, 50, 239, 3, 2, 2, 2, 52, 241, 3, 2, 2, 2, 54, 253,
	3, 2, 2, 2, 56, 255, 3, 2, 2, 2, 58, 261, 3, 2, 2, 2, 60, 268, 3, 2, 2,
	2, 62, 271, 3, 2, 2, 2, 64, 280, 3, 2, 2, 2, 66, 293, 3, 2, 2, 2, 68, 295,
	3, 2, 2, 2, 70, 308, 3, 2, 2, 2, 72, 76, 5, 4, 3, 2, 73, 76, 5, 6, 4, 2,
	74, 76, 5, 14, 8, 2, 75, 72, 3, 2, 2, 2, 75, 73, 3, 2, 2, 2, 75, 74, 3,
	2, 2, 2, 76, 79, 3, 2, 2, 2, 77, 75, 3, 2, 2, 2, 77, 78, 3, 2, 2, 2, 78,
	80, 3, 2, 2, 2, 79, 77, 3, 2, 2, 2, 80, 81, 7, 2, 2, 3, 81, 3, 3, 2, 2,
	2, 82, 83, 7, 45, 2, 2, 83, 5, 3, 2, 2, 2, 84, 85, 5, 8, 5, 2, 85, 86,
	7, 41, 2, 2, 86, 87, 5, 10, 6, 2, 87, 7, 3, 2, 2, 2, 88, 89, 7, 45, 2,
	2, 89, 9, 3, 2, 2, 2, 90, 93, 7, 45, 2, 2, 91, 92, 7, 43, 2, 2, 92, 94,
	7, 45, 2, 2, 93, 91, 3, 2, 2, 2, 93, 94, 3, 2, 2, 2, 94, 109, 3, 2, 2,
	2, 95, 109, 7, 26, 2, 2, 96, 98, 7, 47, 2, 2, 97, 99, 5, 12, 7, 2, 98,
	97, 3, 2, 2, 2, 98, 99, 3, 2, 2, 2, 99, 109, 3, 2, 2, 2, 100, 102, 7, 48,
	2, 2, 101, 103, 5, 12, 7, 2, 102, 101, 3, 2, 2, 2, 102, 103, 3, 2, 2, 2,
	103, 109, 3, 2, 2, 2, 104, 106, 7, 49, 2, 2, 105, 107, 5, 12, 7, 2, 106,
	105, 3, 2, 2, 2, 106, 107, 3, 2, 2, 2, 107, 109, 3, 2, 2, 2, 108, 90, 3,
	2, 2, 2, 108, 95, 3, 2, 2, 2, 108, 96, 3, 2, 2, 2, 108, 100, 3, 2, 2, 2,
	108, 104, 3, 2, 2, 2, 109, 11, 3, 2, 2, 2, 110, 111, 9, 2, 2, 2, 111, 13,
	3, 2, 2, 2, 112, 120, 5, 6, 4, 2, 113, 120, 5, 16, 9, 2, 114, 120, 5, 40,
	21, 2, 115, 120, 5, 56, 29, 2, 116, 120, 5, 42, 22, 2, 117, 120, 5, 62,
	32, 2, 118, 120, 5, 46, 24, 2, 119, 112, 3, 2, 2, 2, 119, 113, 3, 2, 2,
	2, 119, 114, 3, 2, 2, 2, 119, 115, 3, 2, 2, 2, 119, 116, 3, 2, 2, 2, 119,
	117, 3, 2, 2, 2, 119, 118, 3, 2, 2, 2, 120, 15, 3, 2, 2, 2, 121, 122, 7,
	12, 2, 2, 122, 123, 7, 45, 2, 2, 123, 127, 7, 35, 2, 2, 124, 126, 5, 18,
	10, 2, 125, 124, 3, 2, 2, 2, 126, 129, 3, 2, 2, 2, 127, 125, 3, 2, 2, 2,
	127, 128, 3, 2, 2, 2, 128, 130, 3, 2, 2, 2, 129, 127, 3, 2, 2, 2, 130,
	131, 7, 36, 2, 2, 131, 17, 3, 2, 2, 2, 132, 136, 5, 20, 11, 2, 133, 136,
	5, 22, 12, 2, 134, 136, 5, 24, 13, 2, 135, 132, 3, 2, 2, 2, 135, 133, 3,
	2, 2, 2, 135, 134, 3, 2, 2, 2, 136, 19, 3, 2, 2, 2, 137, 142, 7, 13, 2,
	2, 138, 143, 7, 45, 2, 2, 139, 140, 7, 26, 2, 2, 140, 141, 7, 42, 2, 2,
	141, 143, 5, 34, 18, 2, 142, 138, 3, 2, 2, 2, 142, 139, 3, 2, 2, 2, 143,
	21, 3, 2, 2, 2, 144, 145, 7, 14, 2, 2, 145, 146, 7, 37, 2, 2, 146, 147,
	5, 30, 16, 2, 147, 148, 7, 38, 2, 2, 148, 149, 7, 26, 2, 2, 149, 150, 7,
	42, 2, 2, 150, 151, 5, 34, 18, 2, 151, 23, 3, 2, 2, 2, 152, 154, 7, 15,
	2, 2, 153, 155, 5, 36, 19, 2, 154, 153, 3, 2, 2, 2, 154, 155, 3, 2, 2,
	2, 155, 156, 3, 2, 2, 2, 156, 157, 7, 41, 2, 2, 157, 159, 5, 28, 15, 2,
	158, 160, 5, 26, 14, 2, 159, 158, 3, 2, 2, 2, 159, 160, 3, 2, 2, 2, 160,
	25, 3, 2, 2, 2, 161, 162, 7, 16, 2, 2, 162, 163, 7, 17, 2, 2, 163, 164,
	7, 33, 2, 2, 164, 165, 5, 38, 20, 2, 165, 166, 7, 34, 2, 2, 166, 27, 3,
	2, 2, 2, 167, 168, 7, 10, 2, 2, 168, 174, 5, 34, 18, 2, 169, 170, 7, 11,
	2, 2, 170, 171, 7, 26, 2, 2, 171, 172, 7, 42, 2, 2, 172, 174, 5, 34, 18,
	2, 173, 167, 3, 2, 2, 2, 173, 169, 3, 2, 2, 2, 174, 29, 3, 2, 2, 2, 175,
	176, 7, 45, 2, 2, 176, 31, 3, 2, 2, 2, 177, 178, 7, 45, 2, 2, 178, 33,
	3, 2, 2, 2, 179, 180, 7, 45, 2, 2, 180, 35, 3, 2, 2, 2, 181, 182, 7, 45,
	2, 2, 182, 37, 3, 2, 2, 2, 183, 184, 7, 45, 2, 2, 184, 39, 3, 2, 2, 2,
	185, 186, 7, 22, 2, 2, 186, 187, 7, 45, 2, 2, 187, 191, 7, 35, 2, 2, 188,
	190, 5, 44, 23, 2, 189, 188, 3, 2, 2, 2, 190, 193, 3, 2, 2, 2, 191, 189,
	3, 2, 2, 2, 191, 192, 3, 2, 2, 2, 192, 194, 3, 2, 2, 2, 193, 191, 3, 2,
	2, 2, 194, 195, 7, 36, 2, 2, 195, 41, 3, 2, 2, 2, 196, 197, 7, 23, 2, 2,
	197, 198, 7, 45, 2, 2, 198, 202, 7, 35, 2, 2, 199, 201, 5, 44, 23, 2, 200,
	199, 3, 2, 2, 2, 201, 204, 3, 2, 2, 2, 202, 200, 3, 2, 2, 2, 202, 203,
	3, 2, 2, 2, 203, 205, 3, 2, 2, 2, 204, 202, 3, 2, 2, 2, 205, 206, 7, 36,
	2, 2, 206, 43, 3, 2, 2, 2, 207, 208, 7, 45, 2, 2, 208, 209, 7, 41, 2, 2,
	209, 210, 5, 10, 6, 2, 210, 45, 3, 2, 2, 2, 211, 212, 7, 20, 2, 2, 212,
	213, 7, 45, 2, 2, 213, 214, 7, 35, 2, 2, 214, 215, 5, 48, 25, 2, 215, 216,
	7, 36, 2, 2, 216, 47, 3, 2, 2, 2, 217, 219, 5, 50, 26, 2, 218, 217, 3,
	2, 2, 2, 219, 222, 3, 2, 2, 2, 220, 218, 3, 2, 2, 2, 220, 221, 3, 2, 2,
	2, 221, 49, 3, 2, 2, 2, 222, 220, 3, 2, 2, 2, 223, 227, 7, 6, 2, 2, 224,
	226, 7, 6, 2, 2, 225, 224, 3, 2, 2, 2, 226, 229, 3, 2, 2, 2, 227, 225,
	3, 2, 2, 2, 227, 228, 3, 2, 2, 2, 228, 240, 3, 2, 2, 2, 229, 227, 3, 2,
	2, 2, 230, 234, 5, 52, 27, 2, 231, 233, 5, 52, 27, 2, 232, 231, 3, 2, 2,
	2, 233, 236, 3, 2, 2, 2, 234, 232, 3, 2, 2, 2, 234, 235, 3, 2, 2, 2, 235,
	237, 3, 2, 2, 2, 236, 234, 3, 2, 2, 2, 237, 238, 7, 7, 2, 2, 238, 240,
	3, 2, 2, 2, 239, 223, 3, 2, 2, 2, 239, 230, 3, 2, 2, 2, 240, 51, 3, 2,
	2, 2, 241, 242, 7, 7, 2, 2, 242, 243, 5, 54, 28, 2, 243, 53, 3, 2, 2, 2,
	244, 254, 7, 48, 2, 2, 245, 254, 7, 21, 2, 2, 246, 250, 5, 34, 18, 2, 247,
	248, 7, 33, 2, 2, 248, 249, 9, 3, 2, 2, 249, 251, 7, 34, 2, 2, 250, 247,
	3, 2, 2, 2, 250, 251, 3, 2, 2, 2, 251, 254, 3, 2, 2, 2, 252, 254, 7, 26,
	2, 2, 253, 244, 3, 2, 2, 2, 253, 245, 3, 2, 2, 2, 253, 246, 3, 2, 2, 2,
	253, 252, 3, 2, 2, 2, 254, 55, 3, 2, 2, 2, 255, 256, 7, 24, 2, 2, 256,
	257, 5, 58, 30, 2, 257, 258, 7, 35, 2, 2, 258, 259, 5, 60, 31, 2, 259,
	260, 7, 36, 2, 2, 260, 57, 3, 2, 2, 2, 261, 262, 7, 45, 2, 2, 262, 59,
	3, 2, 2, 2, 263, 264, 5, 6, 4, 2, 264, 265, 7, 8, 2, 2, 265, 267, 3, 2,
	2, 2, 266, 263, 3, 2, 2, 2, 267, 270, 3, 2, 2, 2, 268, 266, 3, 2, 2, 2,
	268, 269, 3, 2, 2, 2, 269, 61, 3, 2, 2, 2, 270, 268, 3, 2, 2, 2, 271, 272,
	7, 25, 2, 2, 272, 273, 5, 70, 36, 2, 273, 274, 7, 35, 2, 2, 274, 275, 5,
	64, 33, 2, 275, 276, 7, 36, 2, 2, 276, 63, 3, 2, 2, 2, 277, 279, 5, 66,
	34, 2, 278, 277, 3, 2, 2, 2, 279, 282, 3, 2, 2, 2, 280, 278, 3, 2, 2, 2,
	280, 281, 3, 2, 2, 2, 281, 65, 3, 2, 2, 2, 282, 280, 3, 2, 2, 2, 283, 284,
	5, 8, 5, 2, 284, 285, 7, 9, 2, 2, 285, 286, 5, 10, 6, 2, 286, 287, 7, 8,
	2, 2, 287, 294, 3, 2, 2, 2, 288, 289, 7, 45, 2, 2, 289, 290, 7, 37, 2,
	2, 290, 291, 5, 68, 35, 2, 291, 292, 7, 38, 2, 2, 292, 294, 3, 2, 2, 2,
	293, 283, 3, 2, 2, 2, 293, 288, 3, 2, 2, 2, 294, 67, 3, 2, 2, 2, 295, 296,
	5, 70, 36, 2, 296, 297, 7, 42, 2, 2, 297, 305, 7, 45, 2, 2, 298, 299, 7,
	43, 2, 2, 299, 300, 5, 70, 36, 2, 300, 301, 7, 42, 2, 2, 301, 302, 7, 45,
	2, 2, 302, 304, 3, 2, 2, 2, 303, 298, 3, 2, 2, 2, 304, 307, 3, 2, 2, 2,
	305, 303, 3, 2, 2, 2, 305, 306, 3, 2, 2, 2, 306, 69, 3, 2, 2, 2, 307, 305,
	3, 2, 2, 2, 308, 309, 7, 45, 2, 2, 309, 71, 3, 2, 2, 2, 28, 75, 77, 93,
	98, 102, 106, 108, 119, 127, 135, 142, 154, 159, 173, 191, 202, 220, 227,
	234, 239, 250, 253, 268, 280, 293, 305,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'rem'", "'px'", "'em'", "'-'", "'|'", "';'", "'='", "", "", "", "",
	"", "", "", "", "'repeat'", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "'('", "')'", "'{'", "'}'", "'['", "']'", "'\"'", "'''", "':'",
	"'.'", "','",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "GOTO_KEY", "SHOW_KEY", "FLOW", "SEE",
	"DO", "REACT", "WITHTEXT", "ANIMATE", "REPEAT", "REPEAT_TIMES", "LAYOUT",
	"POSITION", "PAGE", "COMPONENT", "STYLE", "LIBRARY", "STRING_LITERAL",
	"WS", "COMMENT", "LINE_COMMENT", "EmptyLine", "Space", "NewLine", "LPAREN",
	"RPAREN", "LBRACE", "RBRACE", "LBRACK", "RBRACK", "Quote", "SingleQuote",
	"COLON", "DOT", "COMMA", "LETTER", "IDENTIFIER", "DIGITS", "DIGITS_IDENTIFIER",
	"DECIMAL_LITERAL", "FLOAT_LITERAL",
}

var ruleNames = []string{
	"start", "comment", "configDecalartion", "configKey", "configValue", "unit",
	"decalartions", "flowDecalartion", "flowBodyDecalartion", "seeDecalartion",
	"doDecalartion", "reactDecalartion", "animateDecalartion", "actionKey",
	"actionName", "componentValue", "componentName", "sceneName", "animateName",
	"pageDecalartion", "componentDecalartion", "componentBodyDecalartion",
	"layoutDecalaration", "layoutBodyDecalartion", "layoutRow", "layoutLine",
	"componentUseDeclaration", "styleDecalartion", "styleName", "styleBody",
	"libraryDecalartion", "libraryBody", "express", "libraryCall", "libraryName",
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
	DesignParserT__4              = 5
	DesignParserT__5              = 6
	DesignParserT__6              = 7
	DesignParserGOTO_KEY          = 8
	DesignParserSHOW_KEY          = 9
	DesignParserFLOW              = 10
	DesignParserSEE               = 11
	DesignParserDO                = 12
	DesignParserREACT             = 13
	DesignParserWITHTEXT          = 14
	DesignParserANIMATE           = 15
	DesignParserREPEAT            = 16
	DesignParserREPEAT_TIMES      = 17
	DesignParserLAYOUT            = 18
	DesignParserPOSITION          = 19
	DesignParserPAGE              = 20
	DesignParserCOMPONENT         = 21
	DesignParserSTYLE             = 22
	DesignParserLIBRARY           = 23
	DesignParserSTRING_LITERAL    = 24
	DesignParserWS                = 25
	DesignParserCOMMENT           = 26
	DesignParserLINE_COMMENT      = 27
	DesignParserEmptyLine         = 28
	DesignParserSpace             = 29
	DesignParserNewLine           = 30
	DesignParserLPAREN            = 31
	DesignParserRPAREN            = 32
	DesignParserLBRACE            = 33
	DesignParserRBRACE            = 34
	DesignParserLBRACK            = 35
	DesignParserRBRACK            = 36
	DesignParserQuote             = 37
	DesignParserSingleQuote       = 38
	DesignParserCOLON             = 39
	DesignParserDOT               = 40
	DesignParserCOMMA             = 41
	DesignParserLETTER            = 42
	DesignParserIDENTIFIER        = 43
	DesignParserDIGITS            = 44
	DesignParserDIGITS_IDENTIFIER = 45
	DesignParserDECIMAL_LITERAL   = 46
	DesignParserFLOAT_LITERAL     = 47
)

// DesignParser rules.
const (
	DesignParserRULE_start                    = 0
	DesignParserRULE_comment                  = 1
	DesignParserRULE_configDecalartion        = 2
	DesignParserRULE_configKey                = 3
	DesignParserRULE_configValue              = 4
	DesignParserRULE_unit                     = 5
	DesignParserRULE_decalartions             = 6
	DesignParserRULE_flowDecalartion          = 7
	DesignParserRULE_flowBodyDecalartion      = 8
	DesignParserRULE_seeDecalartion           = 9
	DesignParserRULE_doDecalartion            = 10
	DesignParserRULE_reactDecalartion         = 11
	DesignParserRULE_animateDecalartion       = 12
	DesignParserRULE_actionKey                = 13
	DesignParserRULE_actionName               = 14
	DesignParserRULE_componentValue           = 15
	DesignParserRULE_componentName            = 16
	DesignParserRULE_sceneName                = 17
	DesignParserRULE_animateName              = 18
	DesignParserRULE_pageDecalartion          = 19
	DesignParserRULE_componentDecalartion     = 20
	DesignParserRULE_componentBodyDecalartion = 21
	DesignParserRULE_layoutDecalaration       = 22
	DesignParserRULE_layoutBodyDecalartion    = 23
	DesignParserRULE_layoutRow                = 24
	DesignParserRULE_layoutLine               = 25
	DesignParserRULE_componentUseDeclaration  = 26
	DesignParserRULE_styleDecalartion         = 27
	DesignParserRULE_styleName                = 28
	DesignParserRULE_styleBody                = 29
	DesignParserRULE_libraryDecalartion       = 30
	DesignParserRULE_libraryBody              = 31
	DesignParserRULE_express                  = 32
	DesignParserRULE_libraryCall              = 33
	DesignParserRULE_libraryName              = 34
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
	p.SetState(75)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<DesignParserFLOW)|(1<<DesignParserLAYOUT)|(1<<DesignParserPAGE)|(1<<DesignParserCOMPONENT)|(1<<DesignParserSTYLE)|(1<<DesignParserLIBRARY))) != 0) || _la == DesignParserIDENTIFIER {
		p.SetState(73)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(70)
				p.Comment()
			}

		case 2:
			{
				p.SetState(71)
				p.ConfigDecalartion()
			}

		case 3:
			{
				p.SetState(72)
				p.Decalartions()
			}

		}

		p.SetState(77)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(78)
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
		p.SetState(80)
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
		p.SetState(82)
		p.ConfigKey()
	}
	{
		p.SetState(83)
		p.Match(DesignParserCOLON)
	}
	{
		p.SetState(84)
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
		p.SetState(86)
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

func (s *ConfigValueContext) Unit() IUnitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnitContext)
}

func (s *ConfigValueContext) DECIMAL_LITERAL() antlr.TerminalNode {
	return s.GetToken(DesignParserDECIMAL_LITERAL, 0)
}

func (s *ConfigValueContext) FLOAT_LITERAL() antlr.TerminalNode {
	return s.GetToken(DesignParserFLOAT_LITERAL, 0)
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

	p.SetState(106)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DesignParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(88)
			p.Match(DesignParserIDENTIFIER)
		}
		p.SetState(91)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DesignParserCOMMA {
			{
				p.SetState(89)
				p.Match(DesignParserCOMMA)
			}
			{
				p.SetState(90)
				p.Match(DesignParserIDENTIFIER)
			}

		}

	case DesignParserSTRING_LITERAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(93)
			p.Match(DesignParserSTRING_LITERAL)
		}

	case DesignParserDIGITS_IDENTIFIER:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(94)
			p.Match(DesignParserDIGITS_IDENTIFIER)
		}
		p.SetState(96)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<DesignParserT__0)|(1<<DesignParserT__1)|(1<<DesignParserT__2))) != 0 {
			{
				p.SetState(95)
				p.Unit()
			}

		}

	case DesignParserDECIMAL_LITERAL:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(98)
			p.Match(DesignParserDECIMAL_LITERAL)
		}
		p.SetState(100)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<DesignParserT__0)|(1<<DesignParserT__1)|(1<<DesignParserT__2))) != 0 {
			{
				p.SetState(99)
				p.Unit()
			}

		}

	case DesignParserFLOAT_LITERAL:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(102)
			p.Match(DesignParserFLOAT_LITERAL)
		}
		p.SetState(104)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<DesignParserT__0)|(1<<DesignParserT__1)|(1<<DesignParserT__2))) != 0 {
			{
				p.SetState(103)
				p.Unit()
			}

		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IUnitContext is an interface to support dynamic dispatch.
type IUnitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnitContext differentiates from other interfaces.
	IsUnitContext()
}

type UnitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnitContext() *UnitContext {
	var p = new(UnitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DesignParserRULE_unit
	return p
}

func (*UnitContext) IsUnitContext() {}

func NewUnitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnitContext {
	var p = new(UnitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesignParserRULE_unit

	return p
}

func (s *UnitContext) GetParser() antlr.Parser { return s.parser }
func (s *UnitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.EnterUnit(s)
	}
}

func (s *UnitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.ExitUnit(s)
	}
}

func (s *UnitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DesignVisitor:
		return t.VisitUnit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DesignParser) Unit() (localctx IUnitContext) {
	localctx = NewUnitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, DesignParserRULE_unit)
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
		p.SetState(108)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<DesignParserT__0)|(1<<DesignParserT__1)|(1<<DesignParserT__2))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
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
	p.EnterRule(localctx, 12, DesignParserRULE_decalartions)

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

	p.SetState(117)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DesignParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(110)
			p.ConfigDecalartion()
		}

	case DesignParserFLOW:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(111)
			p.FlowDecalartion()
		}

	case DesignParserPAGE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(112)
			p.PageDecalartion()
		}

	case DesignParserSTYLE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(113)
			p.StyleDecalartion()
		}

	case DesignParserCOMPONENT:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(114)
			p.ComponentDecalartion()
		}

	case DesignParserLIBRARY:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(115)
			p.LibraryDecalartion()
		}

	case DesignParserLAYOUT:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(116)
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
	p.EnterRule(localctx, 14, DesignParserRULE_flowDecalartion)
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
		p.Match(DesignParserFLOW)
	}
	{
		p.SetState(120)
		p.Match(DesignParserIDENTIFIER)
	}
	{
		p.SetState(121)
		p.Match(DesignParserLBRACE)
	}
	p.SetState(125)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<DesignParserSEE)|(1<<DesignParserDO)|(1<<DesignParserREACT))) != 0 {
		{
			p.SetState(122)
			p.FlowBodyDecalartion()
		}

		p.SetState(127)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(128)
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
	p.EnterRule(localctx, 16, DesignParserRULE_flowBodyDecalartion)

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

	p.SetState(133)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DesignParserSEE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(130)
			p.SeeDecalartion()
		}

	case DesignParserDO:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(131)
			p.DoDecalartion()
		}

	case DesignParserREACT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(132)
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
	p.EnterRule(localctx, 18, DesignParserRULE_seeDecalartion)

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
		p.Match(DesignParserSEE)
	}
	p.SetState(140)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DesignParserIDENTIFIER:
		{
			p.SetState(136)
			p.Match(DesignParserIDENTIFIER)
		}

	case DesignParserSTRING_LITERAL:
		{
			p.SetState(137)
			p.Match(DesignParserSTRING_LITERAL)
		}
		{
			p.SetState(138)
			p.Match(DesignParserDOT)
		}
		{
			p.SetState(139)
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
	p.EnterRule(localctx, 20, DesignParserRULE_doDecalartion)

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
		p.Match(DesignParserDO)
	}
	{
		p.SetState(143)
		p.Match(DesignParserLBRACK)
	}
	{
		p.SetState(144)
		p.ActionName()
	}
	{
		p.SetState(145)
		p.Match(DesignParserRBRACK)
	}
	{
		p.SetState(146)
		p.Match(DesignParserSTRING_LITERAL)
	}
	{
		p.SetState(147)
		p.Match(DesignParserDOT)
	}
	{
		p.SetState(148)
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
	p.EnterRule(localctx, 22, DesignParserRULE_reactDecalartion)
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
		p.SetState(150)
		p.Match(DesignParserREACT)
	}
	p.SetState(152)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DesignParserIDENTIFIER {
		{
			p.SetState(151)
			p.SceneName()
		}

	}
	{
		p.SetState(154)
		p.Match(DesignParserCOLON)
	}
	{
		p.SetState(155)
		p.ActionKey()
	}
	p.SetState(157)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DesignParserWITHTEXT {
		{
			p.SetState(156)
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
	p.EnterRule(localctx, 24, DesignParserRULE_animateDecalartion)

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
		p.Match(DesignParserWITHTEXT)
	}
	{
		p.SetState(160)
		p.Match(DesignParserANIMATE)
	}
	{
		p.SetState(161)
		p.Match(DesignParserLPAREN)
	}
	{
		p.SetState(162)
		p.AnimateName()
	}
	{
		p.SetState(163)
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
	p.EnterRule(localctx, 26, DesignParserRULE_actionKey)

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

	switch p.GetTokenStream().LA(1) {
	case DesignParserGOTO_KEY:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(165)
			p.Match(DesignParserGOTO_KEY)
		}
		{
			p.SetState(166)
			p.ComponentName()
		}

	case DesignParserSHOW_KEY:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(167)
			p.Match(DesignParserSHOW_KEY)
		}
		{
			p.SetState(168)
			p.Match(DesignParserSTRING_LITERAL)
		}
		{
			p.SetState(169)
			p.Match(DesignParserDOT)
		}
		{
			p.SetState(170)
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
	p.EnterRule(localctx, 28, DesignParserRULE_actionName)

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
		p.SetState(173)
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
	p.EnterRule(localctx, 30, DesignParserRULE_componentValue)

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
		p.SetState(175)
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
	p.EnterRule(localctx, 32, DesignParserRULE_componentName)

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
	p.EnterRule(localctx, 34, DesignParserRULE_sceneName)

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
		p.SetState(179)
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
	p.EnterRule(localctx, 36, DesignParserRULE_animateName)

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
		p.SetState(181)
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
	p.EnterRule(localctx, 38, DesignParserRULE_pageDecalartion)
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
		p.SetState(183)
		p.Match(DesignParserPAGE)
	}
	{
		p.SetState(184)
		p.Match(DesignParserIDENTIFIER)
	}
	{
		p.SetState(185)
		p.Match(DesignParserLBRACE)
	}
	p.SetState(189)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DesignParserIDENTIFIER {
		{
			p.SetState(186)
			p.ComponentBodyDecalartion()
		}

		p.SetState(191)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(192)
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
	p.EnterRule(localctx, 40, DesignParserRULE_componentDecalartion)
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
		p.SetState(194)
		p.Match(DesignParserCOMPONENT)
	}
	{
		p.SetState(195)
		p.Match(DesignParserIDENTIFIER)
	}
	{
		p.SetState(196)
		p.Match(DesignParserLBRACE)
	}
	p.SetState(200)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DesignParserIDENTIFIER {
		{
			p.SetState(197)
			p.ComponentBodyDecalartion()
		}

		p.SetState(202)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(203)
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
	p.EnterRule(localctx, 42, DesignParserRULE_componentBodyDecalartion)

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
		p.Match(DesignParserIDENTIFIER)
	}
	{
		p.SetState(206)
		p.Match(DesignParserCOLON)
	}
	{
		p.SetState(207)
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
	p.EnterRule(localctx, 44, DesignParserRULE_layoutDecalaration)

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
		p.Match(DesignParserLAYOUT)
	}
	{
		p.SetState(210)
		p.Match(DesignParserIDENTIFIER)
	}
	{
		p.SetState(211)
		p.Match(DesignParserLBRACE)
	}
	{
		p.SetState(212)
		p.LayoutBodyDecalartion()
	}
	{
		p.SetState(213)
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
	p.EnterRule(localctx, 46, DesignParserRULE_layoutBodyDecalartion)
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
	p.SetState(218)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DesignParserT__3 || _la == DesignParserT__4 {
		{
			p.SetState(215)
			p.LayoutRow()
		}

		p.SetState(220)
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
	p.EnterRule(localctx, 48, DesignParserRULE_layoutRow)

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

	p.SetState(237)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DesignParserT__3:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(221)
			p.Match(DesignParserT__3)
		}
		p.SetState(225)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(222)
					p.Match(DesignParserT__3)
				}

			}
			p.SetState(227)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext())
		}

	case DesignParserT__4:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(228)
			p.LayoutLine()
		}
		p.SetState(232)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(229)
					p.LayoutLine()
				}

			}
			p.SetState(234)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext())
		}
		{
			p.SetState(235)
			p.Match(DesignParserT__4)
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
	p.EnterRule(localctx, 50, DesignParserRULE_layoutLine)

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
		p.Match(DesignParserT__4)
	}
	{
		p.SetState(240)
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
	p.EnterRule(localctx, 52, DesignParserRULE_componentUseDeclaration)
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

	p.SetState(251)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DesignParserDECIMAL_LITERAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(242)
			p.Match(DesignParserDECIMAL_LITERAL)
		}

	case DesignParserPOSITION:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(243)
			p.Match(DesignParserPOSITION)
		}

	case DesignParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(244)
			p.ComponentName()
		}
		p.SetState(248)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DesignParserLPAREN {
			{
				p.SetState(245)
				p.Match(DesignParserLPAREN)
			}
			{
				p.SetState(246)
				_la = p.GetTokenStream().LA(1)

				if !(((_la-19)&-(0x1f+1)) == 0 && ((1<<uint((_la-19)))&((1<<(DesignParserPOSITION-19))|(1<<(DesignParserSTRING_LITERAL-19))|(1<<(DesignParserDIGITS-19)))) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(247)
				p.Match(DesignParserRPAREN)
			}

		}

	case DesignParserSTRING_LITERAL:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(250)
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
	p.EnterRule(localctx, 54, DesignParserRULE_styleDecalartion)

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
		p.SetState(253)
		p.Match(DesignParserSTYLE)
	}
	{
		p.SetState(254)
		p.StyleName()
	}
	{
		p.SetState(255)
		p.Match(DesignParserLBRACE)
	}
	{
		p.SetState(256)
		p.StyleBody()
	}
	{
		p.SetState(257)
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
	p.EnterRule(localctx, 56, DesignParserRULE_styleName)

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
		p.SetState(259)
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
	p.EnterRule(localctx, 58, DesignParserRULE_styleBody)
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
	p.SetState(266)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DesignParserIDENTIFIER {
		{
			p.SetState(261)
			p.ConfigDecalartion()
		}
		{
			p.SetState(262)
			p.Match(DesignParserT__5)
		}

		p.SetState(268)
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
	p.EnterRule(localctx, 60, DesignParserRULE_libraryDecalartion)

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
		p.Match(DesignParserLIBRARY)
	}
	{
		p.SetState(270)
		p.LibraryName()
	}
	{
		p.SetState(271)
		p.Match(DesignParserLBRACE)
	}
	{
		p.SetState(272)
		p.LibraryBody()
	}
	{
		p.SetState(273)
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
	p.EnterRule(localctx, 62, DesignParserRULE_libraryBody)
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
	p.SetState(278)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DesignParserIDENTIFIER {
		{
			p.SetState(275)
			p.Express()
		}

		p.SetState(280)
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
	p.EnterRule(localctx, 64, DesignParserRULE_express)

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

	p.SetState(291)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(281)
			p.ConfigKey()
		}
		{
			p.SetState(282)
			p.Match(DesignParserT__6)
		}
		{
			p.SetState(283)
			p.ConfigValue()
		}
		{
			p.SetState(284)
			p.Match(DesignParserT__5)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(286)
			p.Match(DesignParserIDENTIFIER)
		}
		{
			p.SetState(287)
			p.Match(DesignParserLBRACK)
		}
		{
			p.SetState(288)
			p.LibraryCall()
		}
		{
			p.SetState(289)
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
	p.EnterRule(localctx, 66, DesignParserRULE_libraryCall)
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
		p.SetState(293)
		p.LibraryName()
	}
	{
		p.SetState(294)
		p.Match(DesignParserDOT)
	}
	{
		p.SetState(295)
		p.Match(DesignParserIDENTIFIER)
	}
	p.SetState(303)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DesignParserCOMMA {
		{
			p.SetState(296)
			p.Match(DesignParserCOMMA)
		}
		{
			p.SetState(297)
			p.LibraryName()
		}
		{
			p.SetState(298)
			p.Match(DesignParserDOT)
		}
		{
			p.SetState(299)
			p.Match(DesignParserIDENTIFIER)
		}

		p.SetState(305)
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
	p.EnterRule(localctx, 68, DesignParserRULE_libraryName)

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
		p.SetState(306)
		p.Match(DesignParserIDENTIFIER)
	}

	return localctx
}
