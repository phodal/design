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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 49, 334,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 3, 2, 3, 2, 3, 2, 7, 2, 82, 10, 2, 12, 2, 14, 2, 85, 11, 2, 3, 2, 3,
	2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 5, 6, 99,
	10, 6, 3, 6, 3, 6, 5, 6, 103, 10, 6, 3, 6, 3, 6, 5, 6, 107, 10, 6, 3, 6,
	3, 6, 3, 6, 5, 6, 112, 10, 6, 3, 6, 5, 6, 115, 10, 6, 3, 7, 3, 7, 3, 8,
	3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 126, 10, 8, 3, 9, 3, 9, 3, 9,
	3, 9, 7, 9, 132, 10, 9, 12, 9, 14, 9, 135, 11, 9, 3, 9, 3, 9, 3, 10, 3,
	10, 3, 10, 5, 10, 142, 10, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 5, 11,
	149, 10, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3,
	13, 3, 13, 5, 13, 161, 10, 13, 3, 13, 3, 13, 3, 13, 5, 13, 166, 10, 13,
	3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 5, 15, 176, 10,
	15, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18,
	3, 19, 3, 19, 3, 20, 3, 20, 3, 21, 3, 21, 3, 22, 3, 22, 3, 23, 3, 23, 3,
	23, 3, 23, 7, 23, 200, 10, 23, 12, 23, 14, 23, 203, 11, 23, 3, 23, 3, 23,
	3, 24, 3, 24, 3, 24, 3, 24, 7, 24, 211, 10, 24, 12, 24, 14, 24, 214, 11,
	24, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 7, 25, 221, 10, 25, 12, 25, 14,
	25, 224, 11, 25, 3, 25, 3, 25, 3, 25, 3, 25, 5, 25, 230, 10, 25, 3, 26,
	3, 26, 3, 26, 3, 26, 7, 26, 236, 10, 26, 12, 26, 14, 26, 239, 11, 26, 3,
	26, 3, 26, 3, 27, 3, 27, 7, 27, 245, 10, 27, 12, 27, 14, 27, 248, 11, 27,
	3, 27, 3, 27, 3, 27, 5, 27, 253, 10, 27, 3, 28, 3, 28, 7, 28, 257, 10,
	28, 12, 28, 14, 28, 260, 11, 28, 3, 29, 3, 29, 3, 29, 3, 30, 3, 30, 3,
	30, 3, 30, 3, 30, 3, 30, 3, 30, 5, 30, 272, 10, 30, 3, 30, 5, 30, 275,
	10, 30, 3, 31, 3, 31, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 33,
	3, 33, 3, 34, 3, 34, 3, 34, 7, 34, 290, 10, 34, 12, 34, 14, 34, 293, 11,
	34, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 36, 7, 36, 302, 10, 36,
	12, 36, 14, 36, 305, 11, 36, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37,
	3, 37, 3, 37, 3, 37, 3, 37, 5, 37, 317, 10, 37, 3, 38, 3, 38, 3, 38, 3,
	38, 3, 38, 3, 38, 3, 38, 3, 38, 7, 38, 327, 10, 38, 12, 38, 14, 38, 330,
	11, 38, 3, 39, 3, 39, 3, 39, 2, 2, 40, 2, 4, 6, 8, 10, 12, 14, 16, 18,
	20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54,
	56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 2, 4, 3, 2, 3, 5, 5, 2, 21,
	21, 26, 26, 47, 47, 2, 335, 2, 83, 3, 2, 2, 2, 4, 88, 3, 2, 2, 2, 6, 90,
	3, 2, 2, 2, 8, 94, 3, 2, 2, 2, 10, 114, 3, 2, 2, 2, 12, 116, 3, 2, 2, 2,
	14, 125, 3, 2, 2, 2, 16, 127, 3, 2, 2, 2, 18, 141, 3, 2, 2, 2, 20, 143,
	3, 2, 2, 2, 22, 150, 3, 2, 2, 2, 24, 158, 3, 2, 2, 2, 26, 167, 3, 2, 2,
	2, 28, 175, 3, 2, 2, 2, 30, 177, 3, 2, 2, 2, 32, 180, 3, 2, 2, 2, 34, 185,
	3, 2, 2, 2, 36, 187, 3, 2, 2, 2, 38, 189, 3, 2, 2, 2, 40, 191, 3, 2, 2,
	2, 42, 193, 3, 2, 2, 2, 44, 195, 3, 2, 2, 2, 46, 206, 3, 2, 2, 2, 48, 229,
	3, 2, 2, 2, 50, 231, 3, 2, 2, 2, 52, 252, 3, 2, 2, 2, 54, 254, 3, 2, 2,
	2, 56, 261, 3, 2, 2, 2, 58, 274, 3, 2, 2, 2, 60, 276, 3, 2, 2, 2, 62, 278,
	3, 2, 2, 2, 64, 284, 3, 2, 2, 2, 66, 291, 3, 2, 2, 2, 68, 294, 3, 2, 2,
	2, 70, 303, 3, 2, 2, 2, 72, 316, 3, 2, 2, 2, 74, 318, 3, 2, 2, 2, 76, 331,
	3, 2, 2, 2, 78, 82, 5, 4, 3, 2, 79, 82, 5, 6, 4, 2, 80, 82, 5, 14, 8, 2,
	81, 78, 3, 2, 2, 2, 81, 79, 3, 2, 2, 2, 81, 80, 3, 2, 2, 2, 82, 85, 3,
	2, 2, 2, 83, 81, 3, 2, 2, 2, 83, 84, 3, 2, 2, 2, 84, 86, 3, 2, 2, 2, 85,
	83, 3, 2, 2, 2, 86, 87, 7, 2, 2, 3, 87, 3, 3, 2, 2, 2, 88, 89, 7, 45, 2,
	2, 89, 5, 3, 2, 2, 2, 90, 91, 5, 8, 5, 2, 91, 92, 7, 41, 2, 2, 92, 93,
	5, 10, 6, 2, 93, 7, 3, 2, 2, 2, 94, 95, 7, 45, 2, 2, 95, 9, 3, 2, 2, 2,
	96, 98, 7, 47, 2, 2, 97, 99, 5, 12, 7, 2, 98, 97, 3, 2, 2, 2, 98, 99, 3,
	2, 2, 2, 99, 115, 3, 2, 2, 2, 100, 102, 7, 48, 2, 2, 101, 103, 5, 12, 7,
	2, 102, 101, 3, 2, 2, 2, 102, 103, 3, 2, 2, 2, 103, 115, 3, 2, 2, 2, 104,
	106, 7, 49, 2, 2, 105, 107, 5, 12, 7, 2, 106, 105, 3, 2, 2, 2, 106, 107,
	3, 2, 2, 2, 107, 115, 3, 2, 2, 2, 108, 111, 7, 45, 2, 2, 109, 110, 7, 43,
	2, 2, 110, 112, 7, 45, 2, 2, 111, 109, 3, 2, 2, 2, 111, 112, 3, 2, 2, 2,
	112, 115, 3, 2, 2, 2, 113, 115, 7, 26, 2, 2, 114, 96, 3, 2, 2, 2, 114,
	100, 3, 2, 2, 2, 114, 104, 3, 2, 2, 2, 114, 108, 3, 2, 2, 2, 114, 113,
	3, 2, 2, 2, 115, 11, 3, 2, 2, 2, 116, 117, 9, 2, 2, 2, 117, 13, 3, 2, 2,
	2, 118, 126, 5, 6, 4, 2, 119, 126, 5, 16, 9, 2, 120, 126, 5, 44, 23, 2,
	121, 126, 5, 62, 32, 2, 122, 126, 5, 46, 24, 2, 123, 126, 5, 68, 35, 2,
	124, 126, 5, 50, 26, 2, 125, 118, 3, 2, 2, 2, 125, 119, 3, 2, 2, 2, 125,
	120, 3, 2, 2, 2, 125, 121, 3, 2, 2, 2, 125, 122, 3, 2, 2, 2, 125, 123,
	3, 2, 2, 2, 125, 124, 3, 2, 2, 2, 126, 15, 3, 2, 2, 2, 127, 128, 7, 12,
	2, 2, 128, 129, 7, 45, 2, 2, 129, 133, 7, 35, 2, 2, 130, 132, 5, 18, 10,
	2, 131, 130, 3, 2, 2, 2, 132, 135, 3, 2, 2, 2, 133, 131, 3, 2, 2, 2, 133,
	134, 3, 2, 2, 2, 134, 136, 3, 2, 2, 2, 135, 133, 3, 2, 2, 2, 136, 137,
	7, 36, 2, 2, 137, 17, 3, 2, 2, 2, 138, 142, 5, 20, 11, 2, 139, 142, 5,
	22, 12, 2, 140, 142, 5, 24, 13, 2, 141, 138, 3, 2, 2, 2, 141, 139, 3, 2,
	2, 2, 141, 140, 3, 2, 2, 2, 142, 19, 3, 2, 2, 2, 143, 148, 7, 13, 2, 2,
	144, 149, 7, 45, 2, 2, 145, 146, 7, 26, 2, 2, 146, 147, 7, 42, 2, 2, 147,
	149, 5, 38, 20, 2, 148, 144, 3, 2, 2, 2, 148, 145, 3, 2, 2, 2, 149, 21,
	3, 2, 2, 2, 150, 151, 7, 14, 2, 2, 151, 152, 7, 37, 2, 2, 152, 153, 5,
	34, 18, 2, 153, 154, 7, 38, 2, 2, 154, 155, 7, 26, 2, 2, 155, 156, 7, 42,
	2, 2, 156, 157, 5, 38, 20, 2, 157, 23, 3, 2, 2, 2, 158, 160, 7, 15, 2,
	2, 159, 161, 5, 40, 21, 2, 160, 159, 3, 2, 2, 2, 160, 161, 3, 2, 2, 2,
	161, 162, 3, 2, 2, 2, 162, 163, 7, 41, 2, 2, 163, 165, 5, 28, 15, 2, 164,
	166, 5, 26, 14, 2, 165, 164, 3, 2, 2, 2, 165, 166, 3, 2, 2, 2, 166, 25,
	3, 2, 2, 2, 167, 168, 7, 16, 2, 2, 168, 169, 7, 17, 2, 2, 169, 170, 7,
	33, 2, 2, 170, 171, 5, 42, 22, 2, 171, 172, 7, 34, 2, 2, 172, 27, 3, 2,
	2, 2, 173, 176, 5, 30, 16, 2, 174, 176, 5, 32, 17, 2, 175, 173, 3, 2, 2,
	2, 175, 174, 3, 2, 2, 2, 176, 29, 3, 2, 2, 2, 177, 178, 7, 10, 2, 2, 178,
	179, 5, 38, 20, 2, 179, 31, 3, 2, 2, 2, 180, 181, 7, 11, 2, 2, 181, 182,
	7, 26, 2, 2, 182, 183, 7, 42, 2, 2, 183, 184, 5, 38, 20, 2, 184, 33, 3,
	2, 2, 2, 185, 186, 7, 45, 2, 2, 186, 35, 3, 2, 2, 2, 187, 188, 7, 45, 2,
	2, 188, 37, 3, 2, 2, 2, 189, 190, 7, 45, 2, 2, 190, 39, 3, 2, 2, 2, 191,
	192, 7, 45, 2, 2, 192, 41, 3, 2, 2, 2, 193, 194, 7, 45, 2, 2, 194, 43,
	3, 2, 2, 2, 195, 196, 7, 22, 2, 2, 196, 197, 7, 45, 2, 2, 197, 201, 7,
	35, 2, 2, 198, 200, 5, 48, 25, 2, 199, 198, 3, 2, 2, 2, 200, 203, 3, 2,
	2, 2, 201, 199, 3, 2, 2, 2, 201, 202, 3, 2, 2, 2, 202, 204, 3, 2, 2, 2,
	203, 201, 3, 2, 2, 2, 204, 205, 7, 36, 2, 2, 205, 45, 3, 2, 2, 2, 206,
	207, 7, 23, 2, 2, 207, 208, 7, 45, 2, 2, 208, 212, 7, 35, 2, 2, 209, 211,
	5, 48, 25, 2, 210, 209, 3, 2, 2, 2, 211, 214, 3, 2, 2, 2, 212, 210, 3,
	2, 2, 2, 212, 213, 3, 2, 2, 2, 213, 215, 3, 2, 2, 2, 214, 212, 3, 2, 2,
	2, 215, 216, 7, 36, 2, 2, 216, 47, 3, 2, 2, 2, 217, 222, 5, 38, 20, 2,
	218, 219, 7, 43, 2, 2, 219, 221, 5, 38, 20, 2, 220, 218, 3, 2, 2, 2, 221,
	224, 3, 2, 2, 2, 222, 220, 3, 2, 2, 2, 222, 223, 3, 2, 2, 2, 223, 230,
	3, 2, 2, 2, 224, 222, 3, 2, 2, 2, 225, 226, 5, 8, 5, 2, 226, 227, 7, 41,
	2, 2, 227, 228, 5, 10, 6, 2, 228, 230, 3, 2, 2, 2, 229, 217, 3, 2, 2, 2,
	229, 225, 3, 2, 2, 2, 230, 49, 3, 2, 2, 2, 231, 232, 7, 20, 2, 2, 232,
	233, 7, 45, 2, 2, 233, 237, 7, 35, 2, 2, 234, 236, 5, 52, 27, 2, 235, 234,
	3, 2, 2, 2, 236, 239, 3, 2, 2, 2, 237, 235, 3, 2, 2, 2, 237, 238, 3, 2,
	2, 2, 238, 240, 3, 2, 2, 2, 239, 237, 3, 2, 2, 2, 240, 241, 7, 36, 2, 2,
	241, 51, 3, 2, 2, 2, 242, 246, 7, 6, 2, 2, 243, 245, 7, 6, 2, 2, 244, 243,
	3, 2, 2, 2, 245, 248, 3, 2, 2, 2, 246, 244, 3, 2, 2, 2, 246, 247, 3, 2,
	2, 2, 247, 253, 3, 2, 2, 2, 248, 246, 3, 2, 2, 2, 249, 250, 5, 54, 28,
	2, 250, 251, 7, 7, 2, 2, 251, 253, 3, 2, 2, 2, 252, 242, 3, 2, 2, 2, 252,
	249, 3, 2, 2, 2, 253, 53, 3, 2, 2, 2, 254, 258, 5, 56, 29, 2, 255, 257,
	5, 56, 29, 2, 256, 255, 3, 2, 2, 2, 257, 260, 3, 2, 2, 2, 258, 256, 3,
	2, 2, 2, 258, 259, 3, 2, 2, 2, 259, 55, 3, 2, 2, 2, 260, 258, 3, 2, 2,
	2, 261, 262, 7, 7, 2, 2, 262, 263, 5, 58, 30, 2, 263, 57, 3, 2, 2, 2, 264,
	275, 7, 48, 2, 2, 265, 275, 7, 21, 2, 2, 266, 271, 5, 38, 20, 2, 267, 268,
	7, 33, 2, 2, 268, 269, 5, 60, 31, 2, 269, 270, 7, 34, 2, 2, 270, 272, 3,
	2, 2, 2, 271, 267, 3, 2, 2, 2, 271, 272, 3, 2, 2, 2, 272, 275, 3, 2, 2,
	2, 273, 275, 7, 26, 2, 2, 274, 264, 3, 2, 2, 2, 274, 265, 3, 2, 2, 2, 274,
	266, 3, 2, 2, 2, 274, 273, 3, 2, 2, 2, 275, 59, 3, 2, 2, 2, 276, 277, 9,
	3, 2, 2, 277, 61, 3, 2, 2, 2, 278, 279, 7, 24, 2, 2, 279, 280, 5, 64, 33,
	2, 280, 281, 7, 35, 2, 2, 281, 282, 5, 66, 34, 2, 282, 283, 7, 36, 2, 2,
	283, 63, 3, 2, 2, 2, 284, 285, 7, 45, 2, 2, 285, 65, 3, 2, 2, 2, 286, 287,
	5, 6, 4, 2, 287, 288, 7, 8, 2, 2, 288, 290, 3, 2, 2, 2, 289, 286, 3, 2,
	2, 2, 290, 293, 3, 2, 2, 2, 291, 289, 3, 2, 2, 2, 291, 292, 3, 2, 2, 2,
	292, 67, 3, 2, 2, 2, 293, 291, 3, 2, 2, 2, 294, 295, 7, 25, 2, 2, 295,
	296, 5, 76, 39, 2, 296, 297, 7, 35, 2, 2, 297, 298, 5, 70, 36, 2, 298,
	299, 7, 36, 2, 2, 299, 69, 3, 2, 2, 2, 300, 302, 5, 72, 37, 2, 301, 300,
	3, 2, 2, 2, 302, 305, 3, 2, 2, 2, 303, 301, 3, 2, 2, 2, 303, 304, 3, 2,
	2, 2, 304, 71, 3, 2, 2, 2, 305, 303, 3, 2, 2, 2, 306, 307, 5, 8, 5, 2,
	307, 308, 7, 9, 2, 2, 308, 309, 5, 10, 6, 2, 309, 310, 7, 8, 2, 2, 310,
	317, 3, 2, 2, 2, 311, 312, 7, 45, 2, 2, 312, 313, 7, 37, 2, 2, 313, 314,
	5, 74, 38, 2, 314, 315, 7, 38, 2, 2, 315, 317, 3, 2, 2, 2, 316, 306, 3,
	2, 2, 2, 316, 311, 3, 2, 2, 2, 317, 73, 3, 2, 2, 2, 318, 319, 5, 76, 39,
	2, 319, 320, 7, 42, 2, 2, 320, 328, 7, 45, 2, 2, 321, 322, 7, 43, 2, 2,
	322, 323, 5, 76, 39, 2, 323, 324, 7, 42, 2, 2, 324, 325, 7, 45, 2, 2, 325,
	327, 3, 2, 2, 2, 326, 321, 3, 2, 2, 2, 327, 330, 3, 2, 2, 2, 328, 326,
	3, 2, 2, 2, 328, 329, 3, 2, 2, 2, 329, 75, 3, 2, 2, 2, 330, 328, 3, 2,
	2, 2, 331, 332, 7, 45, 2, 2, 332, 77, 3, 2, 2, 2, 30, 81, 83, 98, 102,
	106, 111, 114, 125, 133, 141, 148, 160, 165, 175, 201, 212, 222, 229, 237,
	246, 252, 258, 271, 274, 291, 303, 316, 328,
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
	"start", "comment", "configDeclaration", "configKey", "configValue", "unit",
	"decalartions", "flowDeclaration", "interactionDeclaration", "seeDeclaration",
	"doDeclaration", "reactDeclaration", "animateDeclaration", "reactAction",
	"gotoAction", "showAction", "actionName", "componentValue", "componentName",
	"sceneName", "animateName", "pageDeclaration", "componentDeclaration",
	"componentBodyDeclaration", "layoutDeclaration", "layoutRow", "layoutLines",
	"layoutLine", "componentUseDeclaration", "componentLayoutValue", "styleDeclaration",
	"styleName", "styleBody", "libraryDeclaration", "libraryBody", "express",
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
	DesignParserRULE_configDeclaration        = 2
	DesignParserRULE_configKey                = 3
	DesignParserRULE_configValue              = 4
	DesignParserRULE_unit                     = 5
	DesignParserRULE_decalartions             = 6
	DesignParserRULE_flowDeclaration          = 7
	DesignParserRULE_interactionDeclaration   = 8
	DesignParserRULE_seeDeclaration           = 9
	DesignParserRULE_doDeclaration            = 10
	DesignParserRULE_reactDeclaration         = 11
	DesignParserRULE_animateDeclaration       = 12
	DesignParserRULE_reactAction              = 13
	DesignParserRULE_gotoAction               = 14
	DesignParserRULE_showAction               = 15
	DesignParserRULE_actionName               = 16
	DesignParserRULE_componentValue           = 17
	DesignParserRULE_componentName            = 18
	DesignParserRULE_sceneName                = 19
	DesignParserRULE_animateName              = 20
	DesignParserRULE_pageDeclaration          = 21
	DesignParserRULE_componentDeclaration     = 22
	DesignParserRULE_componentBodyDeclaration = 23
	DesignParserRULE_layoutDeclaration        = 24
	DesignParserRULE_layoutRow                = 25
	DesignParserRULE_layoutLines              = 26
	DesignParserRULE_layoutLine               = 27
	DesignParserRULE_componentUseDeclaration  = 28
	DesignParserRULE_componentLayoutValue     = 29
	DesignParserRULE_styleDeclaration         = 30
	DesignParserRULE_styleName                = 31
	DesignParserRULE_styleBody                = 32
	DesignParserRULE_libraryDeclaration       = 33
	DesignParserRULE_libraryBody              = 34
	DesignParserRULE_express                  = 35
	DesignParserRULE_libraryCall              = 36
	DesignParserRULE_libraryName              = 37
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

func (s *StartContext) AllConfigDeclaration() []IConfigDeclarationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IConfigDeclarationContext)(nil)).Elem())
	var tst = make([]IConfigDeclarationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IConfigDeclarationContext)
		}
	}

	return tst
}

func (s *StartContext) ConfigDeclaration(i int) IConfigDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConfigDeclarationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IConfigDeclarationContext)
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
	p.SetState(81)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<DesignParserFLOW)|(1<<DesignParserLAYOUT)|(1<<DesignParserPAGE)|(1<<DesignParserCOMPONENT)|(1<<DesignParserSTYLE)|(1<<DesignParserLIBRARY))) != 0) || _la == DesignParserIDENTIFIER {
		p.SetState(79)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(76)
				p.Comment()
			}

		case 2:
			{
				p.SetState(77)
				p.ConfigDeclaration()
			}

		case 3:
			{
				p.SetState(78)
				p.Decalartions()
			}

		}

		p.SetState(83)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(84)
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
		p.SetState(86)
		p.Match(DesignParserIDENTIFIER)
	}

	return localctx
}

// IConfigDeclarationContext is an interface to support dynamic dispatch.
type IConfigDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConfigDeclarationContext differentiates from other interfaces.
	IsConfigDeclarationContext()
}

type ConfigDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConfigDeclarationContext() *ConfigDeclarationContext {
	var p = new(ConfigDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DesignParserRULE_configDeclaration
	return p
}

func (*ConfigDeclarationContext) IsConfigDeclarationContext() {}

func NewConfigDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConfigDeclarationContext {
	var p = new(ConfigDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesignParserRULE_configDeclaration

	return p
}

func (s *ConfigDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *ConfigDeclarationContext) ConfigKey() IConfigKeyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConfigKeyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConfigKeyContext)
}

func (s *ConfigDeclarationContext) COLON() antlr.TerminalNode {
	return s.GetToken(DesignParserCOLON, 0)
}

func (s *ConfigDeclarationContext) ConfigValue() IConfigValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConfigValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConfigValueContext)
}

func (s *ConfigDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConfigDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConfigDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.EnterConfigDeclaration(s)
	}
}

func (s *ConfigDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.ExitConfigDeclaration(s)
	}
}

func (s *ConfigDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DesignVisitor:
		return t.VisitConfigDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DesignParser) ConfigDeclaration() (localctx IConfigDeclarationContext) {
	localctx = NewConfigDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, DesignParserRULE_configDeclaration)

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
		p.SetState(88)
		p.ConfigKey()
	}
	{
		p.SetState(89)
		p.Match(DesignParserCOLON)
	}
	{
		p.SetState(90)
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
		p.SetState(92)
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

	p.SetState(112)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DesignParserDIGITS_IDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
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
		p.EnterOuterAlt(localctx, 2)
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
		p.EnterOuterAlt(localctx, 3)
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

	case DesignParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(106)
			p.Match(DesignParserIDENTIFIER)
		}
		p.SetState(109)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DesignParserCOMMA {
			{
				p.SetState(107)
				p.Match(DesignParserCOMMA)
			}
			{
				p.SetState(108)
				p.Match(DesignParserIDENTIFIER)
			}

		}

	case DesignParserSTRING_LITERAL:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(111)
			p.Match(DesignParserSTRING_LITERAL)
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
		p.SetState(114)
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

func (s *DecalartionsContext) ConfigDeclaration() IConfigDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConfigDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConfigDeclarationContext)
}

func (s *DecalartionsContext) FlowDeclaration() IFlowDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFlowDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFlowDeclarationContext)
}

func (s *DecalartionsContext) PageDeclaration() IPageDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPageDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPageDeclarationContext)
}

func (s *DecalartionsContext) StyleDeclaration() IStyleDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStyleDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStyleDeclarationContext)
}

func (s *DecalartionsContext) ComponentDeclaration() IComponentDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComponentDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComponentDeclarationContext)
}

func (s *DecalartionsContext) LibraryDeclaration() ILibraryDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILibraryDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILibraryDeclarationContext)
}

func (s *DecalartionsContext) LayoutDeclaration() ILayoutDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILayoutDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILayoutDeclarationContext)
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

	p.SetState(123)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DesignParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(116)
			p.ConfigDeclaration()
		}

	case DesignParserFLOW:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(117)
			p.FlowDeclaration()
		}

	case DesignParserPAGE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(118)
			p.PageDeclaration()
		}

	case DesignParserSTYLE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(119)
			p.StyleDeclaration()
		}

	case DesignParserCOMPONENT:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(120)
			p.ComponentDeclaration()
		}

	case DesignParserLIBRARY:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(121)
			p.LibraryDeclaration()
		}

	case DesignParserLAYOUT:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(122)
			p.LayoutDeclaration()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFlowDeclarationContext is an interface to support dynamic dispatch.
type IFlowDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFlowDeclarationContext differentiates from other interfaces.
	IsFlowDeclarationContext()
}

type FlowDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFlowDeclarationContext() *FlowDeclarationContext {
	var p = new(FlowDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DesignParserRULE_flowDeclaration
	return p
}

func (*FlowDeclarationContext) IsFlowDeclarationContext() {}

func NewFlowDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FlowDeclarationContext {
	var p = new(FlowDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesignParserRULE_flowDeclaration

	return p
}

func (s *FlowDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *FlowDeclarationContext) FLOW() antlr.TerminalNode {
	return s.GetToken(DesignParserFLOW, 0)
}

func (s *FlowDeclarationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(DesignParserIDENTIFIER, 0)
}

func (s *FlowDeclarationContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(DesignParserLBRACE, 0)
}

func (s *FlowDeclarationContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(DesignParserRBRACE, 0)
}

func (s *FlowDeclarationContext) AllInteractionDeclaration() []IInteractionDeclarationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInteractionDeclarationContext)(nil)).Elem())
	var tst = make([]IInteractionDeclarationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInteractionDeclarationContext)
		}
	}

	return tst
}

func (s *FlowDeclarationContext) InteractionDeclaration(i int) IInteractionDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInteractionDeclarationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInteractionDeclarationContext)
}

func (s *FlowDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FlowDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FlowDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.EnterFlowDeclaration(s)
	}
}

func (s *FlowDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.ExitFlowDeclaration(s)
	}
}

func (s *FlowDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DesignVisitor:
		return t.VisitFlowDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DesignParser) FlowDeclaration() (localctx IFlowDeclarationContext) {
	localctx = NewFlowDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, DesignParserRULE_flowDeclaration)
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
		p.Match(DesignParserFLOW)
	}
	{
		p.SetState(126)
		p.Match(DesignParserIDENTIFIER)
	}
	{
		p.SetState(127)
		p.Match(DesignParserLBRACE)
	}
	p.SetState(131)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<DesignParserSEE)|(1<<DesignParserDO)|(1<<DesignParserREACT))) != 0 {
		{
			p.SetState(128)
			p.InteractionDeclaration()
		}

		p.SetState(133)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(134)
		p.Match(DesignParserRBRACE)
	}

	return localctx
}

// IInteractionDeclarationContext is an interface to support dynamic dispatch.
type IInteractionDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInteractionDeclarationContext differentiates from other interfaces.
	IsInteractionDeclarationContext()
}

type InteractionDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInteractionDeclarationContext() *InteractionDeclarationContext {
	var p = new(InteractionDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DesignParserRULE_interactionDeclaration
	return p
}

func (*InteractionDeclarationContext) IsInteractionDeclarationContext() {}

func NewInteractionDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InteractionDeclarationContext {
	var p = new(InteractionDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesignParserRULE_interactionDeclaration

	return p
}

func (s *InteractionDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *InteractionDeclarationContext) SeeDeclaration() ISeeDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISeeDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISeeDeclarationContext)
}

func (s *InteractionDeclarationContext) DoDeclaration() IDoDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDoDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDoDeclarationContext)
}

func (s *InteractionDeclarationContext) ReactDeclaration() IReactDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReactDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReactDeclarationContext)
}

func (s *InteractionDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InteractionDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InteractionDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.EnterInteractionDeclaration(s)
	}
}

func (s *InteractionDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.ExitInteractionDeclaration(s)
	}
}

func (s *InteractionDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DesignVisitor:
		return t.VisitInteractionDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DesignParser) InteractionDeclaration() (localctx IInteractionDeclarationContext) {
	localctx = NewInteractionDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, DesignParserRULE_interactionDeclaration)

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

	p.SetState(139)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DesignParserSEE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(136)
			p.SeeDeclaration()
		}

	case DesignParserDO:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(137)
			p.DoDeclaration()
		}

	case DesignParserREACT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(138)
			p.ReactDeclaration()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ISeeDeclarationContext is an interface to support dynamic dispatch.
type ISeeDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSeeDeclarationContext differentiates from other interfaces.
	IsSeeDeclarationContext()
}

type SeeDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySeeDeclarationContext() *SeeDeclarationContext {
	var p = new(SeeDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DesignParserRULE_seeDeclaration
	return p
}

func (*SeeDeclarationContext) IsSeeDeclarationContext() {}

func NewSeeDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SeeDeclarationContext {
	var p = new(SeeDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesignParserRULE_seeDeclaration

	return p
}

func (s *SeeDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *SeeDeclarationContext) SEE() antlr.TerminalNode {
	return s.GetToken(DesignParserSEE, 0)
}

func (s *SeeDeclarationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(DesignParserIDENTIFIER, 0)
}

func (s *SeeDeclarationContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(DesignParserSTRING_LITERAL, 0)
}

func (s *SeeDeclarationContext) DOT() antlr.TerminalNode {
	return s.GetToken(DesignParserDOT, 0)
}

func (s *SeeDeclarationContext) ComponentName() IComponentNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComponentNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComponentNameContext)
}

func (s *SeeDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SeeDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SeeDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.EnterSeeDeclaration(s)
	}
}

func (s *SeeDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.ExitSeeDeclaration(s)
	}
}

func (s *SeeDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DesignVisitor:
		return t.VisitSeeDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DesignParser) SeeDeclaration() (localctx ISeeDeclarationContext) {
	localctx = NewSeeDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, DesignParserRULE_seeDeclaration)

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
		p.SetState(141)
		p.Match(DesignParserSEE)
	}
	p.SetState(146)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DesignParserIDENTIFIER:
		{
			p.SetState(142)
			p.Match(DesignParserIDENTIFIER)
		}

	case DesignParserSTRING_LITERAL:
		{
			p.SetState(143)
			p.Match(DesignParserSTRING_LITERAL)
		}
		{
			p.SetState(144)
			p.Match(DesignParserDOT)
		}
		{
			p.SetState(145)
			p.ComponentName()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IDoDeclarationContext is an interface to support dynamic dispatch.
type IDoDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDoDeclarationContext differentiates from other interfaces.
	IsDoDeclarationContext()
}

type DoDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDoDeclarationContext() *DoDeclarationContext {
	var p = new(DoDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DesignParserRULE_doDeclaration
	return p
}

func (*DoDeclarationContext) IsDoDeclarationContext() {}

func NewDoDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DoDeclarationContext {
	var p = new(DoDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesignParserRULE_doDeclaration

	return p
}

func (s *DoDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *DoDeclarationContext) DO() antlr.TerminalNode {
	return s.GetToken(DesignParserDO, 0)
}

func (s *DoDeclarationContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(DesignParserLBRACK, 0)
}

func (s *DoDeclarationContext) ActionName() IActionNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IActionNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IActionNameContext)
}

func (s *DoDeclarationContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(DesignParserRBRACK, 0)
}

func (s *DoDeclarationContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(DesignParserSTRING_LITERAL, 0)
}

func (s *DoDeclarationContext) DOT() antlr.TerminalNode {
	return s.GetToken(DesignParserDOT, 0)
}

func (s *DoDeclarationContext) ComponentName() IComponentNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComponentNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComponentNameContext)
}

func (s *DoDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DoDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DoDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.EnterDoDeclaration(s)
	}
}

func (s *DoDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.ExitDoDeclaration(s)
	}
}

func (s *DoDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DesignVisitor:
		return t.VisitDoDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DesignParser) DoDeclaration() (localctx IDoDeclarationContext) {
	localctx = NewDoDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, DesignParserRULE_doDeclaration)

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
		p.Match(DesignParserDO)
	}
	{
		p.SetState(149)
		p.Match(DesignParserLBRACK)
	}
	{
		p.SetState(150)
		p.ActionName()
	}
	{
		p.SetState(151)
		p.Match(DesignParserRBRACK)
	}
	{
		p.SetState(152)
		p.Match(DesignParserSTRING_LITERAL)
	}
	{
		p.SetState(153)
		p.Match(DesignParserDOT)
	}
	{
		p.SetState(154)
		p.ComponentName()
	}

	return localctx
}

// IReactDeclarationContext is an interface to support dynamic dispatch.
type IReactDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReactDeclarationContext differentiates from other interfaces.
	IsReactDeclarationContext()
}

type ReactDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReactDeclarationContext() *ReactDeclarationContext {
	var p = new(ReactDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DesignParserRULE_reactDeclaration
	return p
}

func (*ReactDeclarationContext) IsReactDeclarationContext() {}

func NewReactDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReactDeclarationContext {
	var p = new(ReactDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesignParserRULE_reactDeclaration

	return p
}

func (s *ReactDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *ReactDeclarationContext) REACT() antlr.TerminalNode {
	return s.GetToken(DesignParserREACT, 0)
}

func (s *ReactDeclarationContext) COLON() antlr.TerminalNode {
	return s.GetToken(DesignParserCOLON, 0)
}

func (s *ReactDeclarationContext) ReactAction() IReactActionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReactActionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReactActionContext)
}

func (s *ReactDeclarationContext) SceneName() ISceneNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISceneNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISceneNameContext)
}

func (s *ReactDeclarationContext) AnimateDeclaration() IAnimateDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnimateDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAnimateDeclarationContext)
}

func (s *ReactDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReactDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReactDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.EnterReactDeclaration(s)
	}
}

func (s *ReactDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.ExitReactDeclaration(s)
	}
}

func (s *ReactDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DesignVisitor:
		return t.VisitReactDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DesignParser) ReactDeclaration() (localctx IReactDeclarationContext) {
	localctx = NewReactDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, DesignParserRULE_reactDeclaration)
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
		p.SetState(156)
		p.Match(DesignParserREACT)
	}
	p.SetState(158)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DesignParserIDENTIFIER {
		{
			p.SetState(157)
			p.SceneName()
		}

	}
	{
		p.SetState(160)
		p.Match(DesignParserCOLON)
	}
	{
		p.SetState(161)
		p.ReactAction()
	}
	p.SetState(163)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DesignParserWITHTEXT {
		{
			p.SetState(162)
			p.AnimateDeclaration()
		}

	}

	return localctx
}

// IAnimateDeclarationContext is an interface to support dynamic dispatch.
type IAnimateDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAnimateDeclarationContext differentiates from other interfaces.
	IsAnimateDeclarationContext()
}

type AnimateDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAnimateDeclarationContext() *AnimateDeclarationContext {
	var p = new(AnimateDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DesignParserRULE_animateDeclaration
	return p
}

func (*AnimateDeclarationContext) IsAnimateDeclarationContext() {}

func NewAnimateDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AnimateDeclarationContext {
	var p = new(AnimateDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesignParserRULE_animateDeclaration

	return p
}

func (s *AnimateDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *AnimateDeclarationContext) WITHTEXT() antlr.TerminalNode {
	return s.GetToken(DesignParserWITHTEXT, 0)
}

func (s *AnimateDeclarationContext) ANIMATE() antlr.TerminalNode {
	return s.GetToken(DesignParserANIMATE, 0)
}

func (s *AnimateDeclarationContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(DesignParserLPAREN, 0)
}

func (s *AnimateDeclarationContext) AnimateName() IAnimateNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnimateNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAnimateNameContext)
}

func (s *AnimateDeclarationContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(DesignParserRPAREN, 0)
}

func (s *AnimateDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AnimateDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AnimateDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.EnterAnimateDeclaration(s)
	}
}

func (s *AnimateDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.ExitAnimateDeclaration(s)
	}
}

func (s *AnimateDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DesignVisitor:
		return t.VisitAnimateDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DesignParser) AnimateDeclaration() (localctx IAnimateDeclarationContext) {
	localctx = NewAnimateDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, DesignParserRULE_animateDeclaration)

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
		p.Match(DesignParserWITHTEXT)
	}
	{
		p.SetState(166)
		p.Match(DesignParserANIMATE)
	}
	{
		p.SetState(167)
		p.Match(DesignParserLPAREN)
	}
	{
		p.SetState(168)
		p.AnimateName()
	}
	{
		p.SetState(169)
		p.Match(DesignParserRPAREN)
	}

	return localctx
}

// IReactActionContext is an interface to support dynamic dispatch.
type IReactActionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReactActionContext differentiates from other interfaces.
	IsReactActionContext()
}

type ReactActionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReactActionContext() *ReactActionContext {
	var p = new(ReactActionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DesignParserRULE_reactAction
	return p
}

func (*ReactActionContext) IsReactActionContext() {}

func NewReactActionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReactActionContext {
	var p = new(ReactActionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesignParserRULE_reactAction

	return p
}

func (s *ReactActionContext) GetParser() antlr.Parser { return s.parser }

func (s *ReactActionContext) GotoAction() IGotoActionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGotoActionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGotoActionContext)
}

func (s *ReactActionContext) ShowAction() IShowActionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IShowActionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IShowActionContext)
}

func (s *ReactActionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReactActionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReactActionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.EnterReactAction(s)
	}
}

func (s *ReactActionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.ExitReactAction(s)
	}
}

func (s *ReactActionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DesignVisitor:
		return t.VisitReactAction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DesignParser) ReactAction() (localctx IReactActionContext) {
	localctx = NewReactActionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, DesignParserRULE_reactAction)

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

	p.SetState(173)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DesignParserGOTO_KEY:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(171)
			p.GotoAction()
		}

	case DesignParserSHOW_KEY:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(172)
			p.ShowAction()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IGotoActionContext is an interface to support dynamic dispatch.
type IGotoActionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGotoActionContext differentiates from other interfaces.
	IsGotoActionContext()
}

type GotoActionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGotoActionContext() *GotoActionContext {
	var p = new(GotoActionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DesignParserRULE_gotoAction
	return p
}

func (*GotoActionContext) IsGotoActionContext() {}

func NewGotoActionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GotoActionContext {
	var p = new(GotoActionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesignParserRULE_gotoAction

	return p
}

func (s *GotoActionContext) GetParser() antlr.Parser { return s.parser }

func (s *GotoActionContext) GOTO_KEY() antlr.TerminalNode {
	return s.GetToken(DesignParserGOTO_KEY, 0)
}

func (s *GotoActionContext) ComponentName() IComponentNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComponentNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComponentNameContext)
}

func (s *GotoActionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GotoActionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GotoActionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.EnterGotoAction(s)
	}
}

func (s *GotoActionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.ExitGotoAction(s)
	}
}

func (s *GotoActionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DesignVisitor:
		return t.VisitGotoAction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DesignParser) GotoAction() (localctx IGotoActionContext) {
	localctx = NewGotoActionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, DesignParserRULE_gotoAction)

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
		p.Match(DesignParserGOTO_KEY)
	}
	{
		p.SetState(176)
		p.ComponentName()
	}

	return localctx
}

// IShowActionContext is an interface to support dynamic dispatch.
type IShowActionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsShowActionContext differentiates from other interfaces.
	IsShowActionContext()
}

type ShowActionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyShowActionContext() *ShowActionContext {
	var p = new(ShowActionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DesignParserRULE_showAction
	return p
}

func (*ShowActionContext) IsShowActionContext() {}

func NewShowActionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ShowActionContext {
	var p = new(ShowActionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesignParserRULE_showAction

	return p
}

func (s *ShowActionContext) GetParser() antlr.Parser { return s.parser }

func (s *ShowActionContext) SHOW_KEY() antlr.TerminalNode {
	return s.GetToken(DesignParserSHOW_KEY, 0)
}

func (s *ShowActionContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(DesignParserSTRING_LITERAL, 0)
}

func (s *ShowActionContext) DOT() antlr.TerminalNode {
	return s.GetToken(DesignParserDOT, 0)
}

func (s *ShowActionContext) ComponentName() IComponentNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComponentNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComponentNameContext)
}

func (s *ShowActionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ShowActionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ShowActionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.EnterShowAction(s)
	}
}

func (s *ShowActionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.ExitShowAction(s)
	}
}

func (s *ShowActionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DesignVisitor:
		return t.VisitShowAction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DesignParser) ShowAction() (localctx IShowActionContext) {
	localctx = NewShowActionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, DesignParserRULE_showAction)

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
		p.SetState(178)
		p.Match(DesignParserSHOW_KEY)
	}
	{
		p.SetState(179)
		p.Match(DesignParserSTRING_LITERAL)
	}
	{
		p.SetState(180)
		p.Match(DesignParserDOT)
	}
	{
		p.SetState(181)
		p.ComponentName()
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
	p.EnterRule(localctx, 32, DesignParserRULE_actionName)

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
	p.EnterRule(localctx, 34, DesignParserRULE_componentValue)

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
		p.SetState(185)
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
	p.EnterRule(localctx, 36, DesignParserRULE_componentName)

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
		p.SetState(187)
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
	p.EnterRule(localctx, 38, DesignParserRULE_sceneName)

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
		p.SetState(189)
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
	p.EnterRule(localctx, 40, DesignParserRULE_animateName)

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

	return localctx
}

// IPageDeclarationContext is an interface to support dynamic dispatch.
type IPageDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPageDeclarationContext differentiates from other interfaces.
	IsPageDeclarationContext()
}

type PageDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPageDeclarationContext() *PageDeclarationContext {
	var p = new(PageDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DesignParserRULE_pageDeclaration
	return p
}

func (*PageDeclarationContext) IsPageDeclarationContext() {}

func NewPageDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PageDeclarationContext {
	var p = new(PageDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesignParserRULE_pageDeclaration

	return p
}

func (s *PageDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *PageDeclarationContext) PAGE() antlr.TerminalNode {
	return s.GetToken(DesignParserPAGE, 0)
}

func (s *PageDeclarationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(DesignParserIDENTIFIER, 0)
}

func (s *PageDeclarationContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(DesignParserLBRACE, 0)
}

func (s *PageDeclarationContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(DesignParserRBRACE, 0)
}

func (s *PageDeclarationContext) AllComponentBodyDeclaration() []IComponentBodyDeclarationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IComponentBodyDeclarationContext)(nil)).Elem())
	var tst = make([]IComponentBodyDeclarationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IComponentBodyDeclarationContext)
		}
	}

	return tst
}

func (s *PageDeclarationContext) ComponentBodyDeclaration(i int) IComponentBodyDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComponentBodyDeclarationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IComponentBodyDeclarationContext)
}

func (s *PageDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PageDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PageDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.EnterPageDeclaration(s)
	}
}

func (s *PageDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.ExitPageDeclaration(s)
	}
}

func (s *PageDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DesignVisitor:
		return t.VisitPageDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DesignParser) PageDeclaration() (localctx IPageDeclarationContext) {
	localctx = NewPageDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, DesignParserRULE_pageDeclaration)
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
		p.SetState(193)
		p.Match(DesignParserPAGE)
	}
	{
		p.SetState(194)
		p.Match(DesignParserIDENTIFIER)
	}
	{
		p.SetState(195)
		p.Match(DesignParserLBRACE)
	}
	p.SetState(199)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DesignParserIDENTIFIER {
		{
			p.SetState(196)
			p.ComponentBodyDeclaration()
		}

		p.SetState(201)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(202)
		p.Match(DesignParserRBRACE)
	}

	return localctx
}

// IComponentDeclarationContext is an interface to support dynamic dispatch.
type IComponentDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsComponentDeclarationContext differentiates from other interfaces.
	IsComponentDeclarationContext()
}

type ComponentDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComponentDeclarationContext() *ComponentDeclarationContext {
	var p = new(ComponentDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DesignParserRULE_componentDeclaration
	return p
}

func (*ComponentDeclarationContext) IsComponentDeclarationContext() {}

func NewComponentDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComponentDeclarationContext {
	var p = new(ComponentDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesignParserRULE_componentDeclaration

	return p
}

func (s *ComponentDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *ComponentDeclarationContext) COMPONENT() antlr.TerminalNode {
	return s.GetToken(DesignParserCOMPONENT, 0)
}

func (s *ComponentDeclarationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(DesignParserIDENTIFIER, 0)
}

func (s *ComponentDeclarationContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(DesignParserLBRACE, 0)
}

func (s *ComponentDeclarationContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(DesignParserRBRACE, 0)
}

func (s *ComponentDeclarationContext) AllComponentBodyDeclaration() []IComponentBodyDeclarationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IComponentBodyDeclarationContext)(nil)).Elem())
	var tst = make([]IComponentBodyDeclarationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IComponentBodyDeclarationContext)
		}
	}

	return tst
}

func (s *ComponentDeclarationContext) ComponentBodyDeclaration(i int) IComponentBodyDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComponentBodyDeclarationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IComponentBodyDeclarationContext)
}

func (s *ComponentDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComponentDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComponentDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.EnterComponentDeclaration(s)
	}
}

func (s *ComponentDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.ExitComponentDeclaration(s)
	}
}

func (s *ComponentDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DesignVisitor:
		return t.VisitComponentDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DesignParser) ComponentDeclaration() (localctx IComponentDeclarationContext) {
	localctx = NewComponentDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, DesignParserRULE_componentDeclaration)
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
		p.SetState(204)
		p.Match(DesignParserCOMPONENT)
	}
	{
		p.SetState(205)
		p.Match(DesignParserIDENTIFIER)
	}
	{
		p.SetState(206)
		p.Match(DesignParserLBRACE)
	}
	p.SetState(210)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DesignParserIDENTIFIER {
		{
			p.SetState(207)
			p.ComponentBodyDeclaration()
		}

		p.SetState(212)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(213)
		p.Match(DesignParserRBRACE)
	}

	return localctx
}

// IComponentBodyDeclarationContext is an interface to support dynamic dispatch.
type IComponentBodyDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsComponentBodyDeclarationContext differentiates from other interfaces.
	IsComponentBodyDeclarationContext()
}

type ComponentBodyDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComponentBodyDeclarationContext() *ComponentBodyDeclarationContext {
	var p = new(ComponentBodyDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DesignParserRULE_componentBodyDeclaration
	return p
}

func (*ComponentBodyDeclarationContext) IsComponentBodyDeclarationContext() {}

func NewComponentBodyDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComponentBodyDeclarationContext {
	var p = new(ComponentBodyDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesignParserRULE_componentBodyDeclaration

	return p
}

func (s *ComponentBodyDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *ComponentBodyDeclarationContext) AllComponentName() []IComponentNameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IComponentNameContext)(nil)).Elem())
	var tst = make([]IComponentNameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IComponentNameContext)
		}
	}

	return tst
}

func (s *ComponentBodyDeclarationContext) ComponentName(i int) IComponentNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComponentNameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IComponentNameContext)
}

func (s *ComponentBodyDeclarationContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(DesignParserCOMMA)
}

func (s *ComponentBodyDeclarationContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(DesignParserCOMMA, i)
}

func (s *ComponentBodyDeclarationContext) ConfigKey() IConfigKeyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConfigKeyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConfigKeyContext)
}

func (s *ComponentBodyDeclarationContext) COLON() antlr.TerminalNode {
	return s.GetToken(DesignParserCOLON, 0)
}

func (s *ComponentBodyDeclarationContext) ConfigValue() IConfigValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConfigValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConfigValueContext)
}

func (s *ComponentBodyDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComponentBodyDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComponentBodyDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.EnterComponentBodyDeclaration(s)
	}
}

func (s *ComponentBodyDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.ExitComponentBodyDeclaration(s)
	}
}

func (s *ComponentBodyDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DesignVisitor:
		return t.VisitComponentBodyDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DesignParser) ComponentBodyDeclaration() (localctx IComponentBodyDeclarationContext) {
	localctx = NewComponentBodyDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, DesignParserRULE_componentBodyDeclaration)
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

	p.SetState(227)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(215)
			p.ComponentName()
		}
		p.SetState(220)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == DesignParserCOMMA {
			{
				p.SetState(216)
				p.Match(DesignParserCOMMA)
			}
			{
				p.SetState(217)
				p.ComponentName()
			}

			p.SetState(222)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(223)
			p.ConfigKey()
		}
		{
			p.SetState(224)
			p.Match(DesignParserCOLON)
		}
		{
			p.SetState(225)
			p.ConfigValue()
		}

	}

	return localctx
}

// ILayoutDeclarationContext is an interface to support dynamic dispatch.
type ILayoutDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLayoutDeclarationContext differentiates from other interfaces.
	IsLayoutDeclarationContext()
}

type LayoutDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLayoutDeclarationContext() *LayoutDeclarationContext {
	var p = new(LayoutDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DesignParserRULE_layoutDeclaration
	return p
}

func (*LayoutDeclarationContext) IsLayoutDeclarationContext() {}

func NewLayoutDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LayoutDeclarationContext {
	var p = new(LayoutDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesignParserRULE_layoutDeclaration

	return p
}

func (s *LayoutDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *LayoutDeclarationContext) LAYOUT() antlr.TerminalNode {
	return s.GetToken(DesignParserLAYOUT, 0)
}

func (s *LayoutDeclarationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(DesignParserIDENTIFIER, 0)
}

func (s *LayoutDeclarationContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(DesignParserLBRACE, 0)
}

func (s *LayoutDeclarationContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(DesignParserRBRACE, 0)
}

func (s *LayoutDeclarationContext) AllLayoutRow() []ILayoutRowContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILayoutRowContext)(nil)).Elem())
	var tst = make([]ILayoutRowContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILayoutRowContext)
		}
	}

	return tst
}

func (s *LayoutDeclarationContext) LayoutRow(i int) ILayoutRowContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILayoutRowContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILayoutRowContext)
}

func (s *LayoutDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LayoutDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LayoutDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.EnterLayoutDeclaration(s)
	}
}

func (s *LayoutDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.ExitLayoutDeclaration(s)
	}
}

func (s *LayoutDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DesignVisitor:
		return t.VisitLayoutDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DesignParser) LayoutDeclaration() (localctx ILayoutDeclarationContext) {
	localctx = NewLayoutDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, DesignParserRULE_layoutDeclaration)
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
		p.SetState(229)
		p.Match(DesignParserLAYOUT)
	}
	{
		p.SetState(230)
		p.Match(DesignParserIDENTIFIER)
	}
	{
		p.SetState(231)
		p.Match(DesignParserLBRACE)
	}
	p.SetState(235)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DesignParserT__3 || _la == DesignParserT__4 {
		{
			p.SetState(232)
			p.LayoutRow()
		}

		p.SetState(237)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(238)
		p.Match(DesignParserRBRACE)
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

func (s *LayoutRowContext) LayoutLines() ILayoutLinesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILayoutLinesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILayoutLinesContext)
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
	p.EnterRule(localctx, 50, DesignParserRULE_layoutRow)

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

	p.SetState(250)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DesignParserT__3:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(240)
			p.Match(DesignParserT__3)
		}
		p.SetState(244)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(241)
					p.Match(DesignParserT__3)
				}

			}
			p.SetState(246)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext())
		}

	case DesignParserT__4:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(247)
			p.LayoutLines()
		}
		{
			p.SetState(248)
			p.Match(DesignParserT__4)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ILayoutLinesContext is an interface to support dynamic dispatch.
type ILayoutLinesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLayoutLinesContext differentiates from other interfaces.
	IsLayoutLinesContext()
}

type LayoutLinesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLayoutLinesContext() *LayoutLinesContext {
	var p = new(LayoutLinesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DesignParserRULE_layoutLines
	return p
}

func (*LayoutLinesContext) IsLayoutLinesContext() {}

func NewLayoutLinesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LayoutLinesContext {
	var p = new(LayoutLinesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesignParserRULE_layoutLines

	return p
}

func (s *LayoutLinesContext) GetParser() antlr.Parser { return s.parser }

func (s *LayoutLinesContext) AllLayoutLine() []ILayoutLineContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILayoutLineContext)(nil)).Elem())
	var tst = make([]ILayoutLineContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILayoutLineContext)
		}
	}

	return tst
}

func (s *LayoutLinesContext) LayoutLine(i int) ILayoutLineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILayoutLineContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILayoutLineContext)
}

func (s *LayoutLinesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LayoutLinesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LayoutLinesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.EnterLayoutLines(s)
	}
}

func (s *LayoutLinesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.ExitLayoutLines(s)
	}
}

func (s *LayoutLinesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DesignVisitor:
		return t.VisitLayoutLines(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DesignParser) LayoutLines() (localctx ILayoutLinesContext) {
	localctx = NewLayoutLinesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, DesignParserRULE_layoutLines)

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
		p.SetState(252)
		p.LayoutLine()
	}
	p.SetState(256)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(253)
				p.LayoutLine()
			}

		}
		p.SetState(258)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 54, DesignParserRULE_layoutLine)

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
		p.Match(DesignParserT__4)
	}
	{
		p.SetState(260)
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

func (s *ComponentUseDeclarationContext) ComponentLayoutValue() IComponentLayoutValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComponentLayoutValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComponentLayoutValueContext)
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
	p.EnterRule(localctx, 56, DesignParserRULE_componentUseDeclaration)
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

	p.SetState(272)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DesignParserDECIMAL_LITERAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(262)
			p.Match(DesignParserDECIMAL_LITERAL)
		}

	case DesignParserPOSITION:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(263)
			p.Match(DesignParserPOSITION)
		}

	case DesignParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(264)
			p.ComponentName()
		}
		p.SetState(269)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DesignParserLPAREN {
			{
				p.SetState(265)
				p.Match(DesignParserLPAREN)
			}
			{
				p.SetState(266)
				p.ComponentLayoutValue()
			}
			{
				p.SetState(267)
				p.Match(DesignParserRPAREN)
			}

		}

	case DesignParserSTRING_LITERAL:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(271)
			p.Match(DesignParserSTRING_LITERAL)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IComponentLayoutValueContext is an interface to support dynamic dispatch.
type IComponentLayoutValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsComponentLayoutValueContext differentiates from other interfaces.
	IsComponentLayoutValueContext()
}

type ComponentLayoutValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComponentLayoutValueContext() *ComponentLayoutValueContext {
	var p = new(ComponentLayoutValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DesignParserRULE_componentLayoutValue
	return p
}

func (*ComponentLayoutValueContext) IsComponentLayoutValueContext() {}

func NewComponentLayoutValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComponentLayoutValueContext {
	var p = new(ComponentLayoutValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesignParserRULE_componentLayoutValue

	return p
}

func (s *ComponentLayoutValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ComponentLayoutValueContext) DIGITS_IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(DesignParserDIGITS_IDENTIFIER, 0)
}

func (s *ComponentLayoutValueContext) POSITION() antlr.TerminalNode {
	return s.GetToken(DesignParserPOSITION, 0)
}

func (s *ComponentLayoutValueContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(DesignParserSTRING_LITERAL, 0)
}

func (s *ComponentLayoutValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComponentLayoutValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComponentLayoutValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.EnterComponentLayoutValue(s)
	}
}

func (s *ComponentLayoutValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.ExitComponentLayoutValue(s)
	}
}

func (s *ComponentLayoutValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DesignVisitor:
		return t.VisitComponentLayoutValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DesignParser) ComponentLayoutValue() (localctx IComponentLayoutValueContext) {
	localctx = NewComponentLayoutValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, DesignParserRULE_componentLayoutValue)
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
		p.SetState(274)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-19)&-(0x1f+1)) == 0 && ((1<<uint((_la-19)))&((1<<(DesignParserPOSITION-19))|(1<<(DesignParserSTRING_LITERAL-19))|(1<<(DesignParserDIGITS_IDENTIFIER-19)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IStyleDeclarationContext is an interface to support dynamic dispatch.
type IStyleDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStyleDeclarationContext differentiates from other interfaces.
	IsStyleDeclarationContext()
}

type StyleDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStyleDeclarationContext() *StyleDeclarationContext {
	var p = new(StyleDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DesignParserRULE_styleDeclaration
	return p
}

func (*StyleDeclarationContext) IsStyleDeclarationContext() {}

func NewStyleDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StyleDeclarationContext {
	var p = new(StyleDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesignParserRULE_styleDeclaration

	return p
}

func (s *StyleDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *StyleDeclarationContext) STYLE() antlr.TerminalNode {
	return s.GetToken(DesignParserSTYLE, 0)
}

func (s *StyleDeclarationContext) StyleName() IStyleNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStyleNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStyleNameContext)
}

func (s *StyleDeclarationContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(DesignParserLBRACE, 0)
}

func (s *StyleDeclarationContext) StyleBody() IStyleBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStyleBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStyleBodyContext)
}

func (s *StyleDeclarationContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(DesignParserRBRACE, 0)
}

func (s *StyleDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StyleDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StyleDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.EnterStyleDeclaration(s)
	}
}

func (s *StyleDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.ExitStyleDeclaration(s)
	}
}

func (s *StyleDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DesignVisitor:
		return t.VisitStyleDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DesignParser) StyleDeclaration() (localctx IStyleDeclarationContext) {
	localctx = NewStyleDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, DesignParserRULE_styleDeclaration)

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
		p.SetState(276)
		p.Match(DesignParserSTYLE)
	}
	{
		p.SetState(277)
		p.StyleName()
	}
	{
		p.SetState(278)
		p.Match(DesignParserLBRACE)
	}
	{
		p.SetState(279)
		p.StyleBody()
	}
	{
		p.SetState(280)
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
	p.EnterRule(localctx, 62, DesignParserRULE_styleName)

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
		p.SetState(282)
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

func (s *StyleBodyContext) AllConfigDeclaration() []IConfigDeclarationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IConfigDeclarationContext)(nil)).Elem())
	var tst = make([]IConfigDeclarationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IConfigDeclarationContext)
		}
	}

	return tst
}

func (s *StyleBodyContext) ConfigDeclaration(i int) IConfigDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConfigDeclarationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IConfigDeclarationContext)
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
	p.EnterRule(localctx, 64, DesignParserRULE_styleBody)
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
	p.SetState(289)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DesignParserIDENTIFIER {
		{
			p.SetState(284)
			p.ConfigDeclaration()
		}
		{
			p.SetState(285)
			p.Match(DesignParserT__5)
		}

		p.SetState(291)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ILibraryDeclarationContext is an interface to support dynamic dispatch.
type ILibraryDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLibraryDeclarationContext differentiates from other interfaces.
	IsLibraryDeclarationContext()
}

type LibraryDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLibraryDeclarationContext() *LibraryDeclarationContext {
	var p = new(LibraryDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DesignParserRULE_libraryDeclaration
	return p
}

func (*LibraryDeclarationContext) IsLibraryDeclarationContext() {}

func NewLibraryDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LibraryDeclarationContext {
	var p = new(LibraryDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesignParserRULE_libraryDeclaration

	return p
}

func (s *LibraryDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *LibraryDeclarationContext) LIBRARY() antlr.TerminalNode {
	return s.GetToken(DesignParserLIBRARY, 0)
}

func (s *LibraryDeclarationContext) LibraryName() ILibraryNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILibraryNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILibraryNameContext)
}

func (s *LibraryDeclarationContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(DesignParserLBRACE, 0)
}

func (s *LibraryDeclarationContext) LibraryBody() ILibraryBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILibraryBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILibraryBodyContext)
}

func (s *LibraryDeclarationContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(DesignParserRBRACE, 0)
}

func (s *LibraryDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LibraryDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LibraryDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.EnterLibraryDeclaration(s)
	}
}

func (s *LibraryDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.ExitLibraryDeclaration(s)
	}
}

func (s *LibraryDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DesignVisitor:
		return t.VisitLibraryDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DesignParser) LibraryDeclaration() (localctx ILibraryDeclarationContext) {
	localctx = NewLibraryDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, DesignParserRULE_libraryDeclaration)

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
		p.Match(DesignParserLIBRARY)
	}
	{
		p.SetState(293)
		p.LibraryName()
	}
	{
		p.SetState(294)
		p.Match(DesignParserLBRACE)
	}
	{
		p.SetState(295)
		p.LibraryBody()
	}
	{
		p.SetState(296)
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
	p.EnterRule(localctx, 68, DesignParserRULE_libraryBody)
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
	p.SetState(301)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DesignParserIDENTIFIER {
		{
			p.SetState(298)
			p.Express()
		}

		p.SetState(303)
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
	p.EnterRule(localctx, 70, DesignParserRULE_express)

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

	p.SetState(314)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(304)
			p.ConfigKey()
		}
		{
			p.SetState(305)
			p.Match(DesignParserT__6)
		}
		{
			p.SetState(306)
			p.ConfigValue()
		}
		{
			p.SetState(307)
			p.Match(DesignParserT__5)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(309)
			p.Match(DesignParserIDENTIFIER)
		}
		{
			p.SetState(310)
			p.Match(DesignParserLBRACK)
		}
		{
			p.SetState(311)
			p.LibraryCall()
		}
		{
			p.SetState(312)
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
	p.EnterRule(localctx, 72, DesignParserRULE_libraryCall)
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
		p.SetState(316)
		p.LibraryName()
	}
	{
		p.SetState(317)
		p.Match(DesignParserDOT)
	}
	{
		p.SetState(318)
		p.Match(DesignParserIDENTIFIER)
	}
	p.SetState(326)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DesignParserCOMMA {
		{
			p.SetState(319)
			p.Match(DesignParserCOMMA)
		}
		{
			p.SetState(320)
			p.LibraryName()
		}
		{
			p.SetState(321)
			p.Match(DesignParserDOT)
		}
		{
			p.SetState(322)
			p.Match(DesignParserIDENTIFIER)
		}

		p.SetState(328)
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
	p.EnterRule(localctx, 74, DesignParserRULE_libraryName)

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
		p.SetState(329)
		p.Match(DesignParserIDENTIFIER)
	}

	return localctx
}
