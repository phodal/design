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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 49, 358,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 3, 2, 3, 2, 3, 2, 7, 2, 88,
	10, 2, 12, 2, 14, 2, 91, 11, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 5, 6, 105, 10, 6, 3, 6, 3, 6, 5, 6, 109,
	10, 6, 3, 6, 3, 6, 5, 6, 113, 10, 6, 3, 6, 3, 6, 3, 6, 5, 6, 118, 10, 6,
	3, 6, 5, 6, 121, 10, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8,
	3, 8, 5, 8, 132, 10, 8, 3, 9, 3, 9, 3, 9, 3, 9, 7, 9, 138, 10, 9, 12, 9,
	14, 9, 141, 11, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 5, 10, 148, 10, 10,
	3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 5, 11, 155, 10, 11, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 5, 13, 167, 10, 13,
	3, 13, 3, 13, 3, 13, 5, 13, 172, 10, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 15, 3, 15, 5, 15, 182, 10, 15, 3, 16, 3, 16, 3, 16, 3, 17,
	3, 17, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 19, 3, 19, 3, 20, 3, 20, 3,
	21, 3, 21, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3, 23, 7, 23, 206, 10, 23,
	12, 23, 14, 23, 209, 11, 23, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 3, 24,
	7, 24, 217, 10, 24, 12, 24, 14, 24, 220, 11, 24, 3, 24, 3, 24, 3, 25, 3,
	25, 3, 25, 7, 25, 227, 10, 25, 12, 25, 14, 25, 230, 11, 25, 3, 25, 3, 25,
	3, 25, 3, 25, 5, 25, 236, 10, 25, 3, 26, 3, 26, 3, 26, 3, 26, 7, 26, 242,
	10, 26, 12, 26, 14, 26, 245, 11, 26, 3, 26, 3, 26, 3, 27, 3, 27, 7, 27,
	251, 10, 27, 12, 27, 14, 27, 254, 11, 27, 3, 27, 3, 27, 3, 27, 5, 27, 259,
	10, 27, 3, 28, 3, 28, 7, 28, 263, 10, 28, 12, 28, 14, 28, 266, 11, 28,
	3, 29, 3, 29, 3, 29, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 5,
	30, 278, 10, 30, 3, 30, 5, 30, 281, 10, 30, 3, 31, 3, 31, 3, 32, 3, 32,
	3, 32, 3, 32, 3, 32, 3, 32, 3, 33, 3, 33, 3, 34, 3, 34, 3, 34, 7, 34, 296,
	10, 34, 12, 34, 14, 34, 299, 11, 34, 3, 35, 3, 35, 3, 35, 3, 35, 7, 35,
	305, 10, 35, 12, 35, 14, 35, 308, 11, 35, 3, 35, 3, 35, 3, 36, 3, 36, 3,
	36, 3, 36, 5, 36, 316, 10, 36, 3, 36, 5, 36, 319, 10, 36, 3, 36, 3, 36,
	3, 36, 7, 36, 324, 10, 36, 12, 36, 14, 36, 327, 11, 36, 3, 36, 3, 36, 5,
	36, 331, 10, 36, 3, 37, 3, 37, 3, 37, 3, 37, 3, 38, 3, 38, 3, 39, 3, 39,
	3, 40, 3, 40, 3, 40, 3, 40, 7, 40, 345, 10, 40, 12, 40, 14, 40, 348, 11,
	40, 3, 40, 3, 40, 3, 41, 3, 41, 3, 41, 3, 41, 3, 42, 3, 42, 3, 42, 2, 2,
	43, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36,
	38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72,
	74, 76, 78, 80, 82, 2, 4, 3, 2, 3, 5, 5, 2, 21, 21, 26, 26, 47, 47, 2,
	359, 2, 89, 3, 2, 2, 2, 4, 94, 3, 2, 2, 2, 6, 96, 3, 2, 2, 2, 8, 100, 3,
	2, 2, 2, 10, 120, 3, 2, 2, 2, 12, 122, 3, 2, 2, 2, 14, 131, 3, 2, 2, 2,
	16, 133, 3, 2, 2, 2, 18, 147, 3, 2, 2, 2, 20, 149, 3, 2, 2, 2, 22, 156,
	3, 2, 2, 2, 24, 164, 3, 2, 2, 2, 26, 173, 3, 2, 2, 2, 28, 181, 3, 2, 2,
	2, 30, 183, 3, 2, 2, 2, 32, 186, 3, 2, 2, 2, 34, 191, 3, 2, 2, 2, 36, 193,
	3, 2, 2, 2, 38, 195, 3, 2, 2, 2, 40, 197, 3, 2, 2, 2, 42, 199, 3, 2, 2,
	2, 44, 201, 3, 2, 2, 2, 46, 212, 3, 2, 2, 2, 48, 235, 3, 2, 2, 2, 50, 237,
	3, 2, 2, 2, 52, 258, 3, 2, 2, 2, 54, 260, 3, 2, 2, 2, 56, 267, 3, 2, 2,
	2, 58, 280, 3, 2, 2, 2, 60, 282, 3, 2, 2, 2, 62, 284, 3, 2, 2, 2, 64, 290,
	3, 2, 2, 2, 66, 297, 3, 2, 2, 2, 68, 300, 3, 2, 2, 2, 70, 330, 3, 2, 2,
	2, 72, 332, 3, 2, 2, 2, 74, 336, 3, 2, 2, 2, 76, 338, 3, 2, 2, 2, 78, 340,
	3, 2, 2, 2, 80, 351, 3, 2, 2, 2, 82, 355, 3, 2, 2, 2, 84, 88, 5, 4, 3,
	2, 85, 88, 5, 6, 4, 2, 86, 88, 5, 14, 8, 2, 87, 84, 3, 2, 2, 2, 87, 85,
	3, 2, 2, 2, 87, 86, 3, 2, 2, 2, 88, 91, 3, 2, 2, 2, 89, 87, 3, 2, 2, 2,
	89, 90, 3, 2, 2, 2, 90, 92, 3, 2, 2, 2, 91, 89, 3, 2, 2, 2, 92, 93, 7,
	2, 2, 3, 93, 3, 3, 2, 2, 2, 94, 95, 7, 45, 2, 2, 95, 5, 3, 2, 2, 2, 96,
	97, 5, 8, 5, 2, 97, 98, 7, 41, 2, 2, 98, 99, 5, 10, 6, 2, 99, 7, 3, 2,
	2, 2, 100, 101, 7, 45, 2, 2, 101, 9, 3, 2, 2, 2, 102, 104, 7, 47, 2, 2,
	103, 105, 5, 12, 7, 2, 104, 103, 3, 2, 2, 2, 104, 105, 3, 2, 2, 2, 105,
	121, 3, 2, 2, 2, 106, 108, 7, 48, 2, 2, 107, 109, 5, 12, 7, 2, 108, 107,
	3, 2, 2, 2, 108, 109, 3, 2, 2, 2, 109, 121, 3, 2, 2, 2, 110, 112, 7, 49,
	2, 2, 111, 113, 5, 12, 7, 2, 112, 111, 3, 2, 2, 2, 112, 113, 3, 2, 2, 2,
	113, 121, 3, 2, 2, 2, 114, 117, 7, 45, 2, 2, 115, 116, 7, 43, 2, 2, 116,
	118, 7, 45, 2, 2, 117, 115, 3, 2, 2, 2, 117, 118, 3, 2, 2, 2, 118, 121,
	3, 2, 2, 2, 119, 121, 7, 26, 2, 2, 120, 102, 3, 2, 2, 2, 120, 106, 3, 2,
	2, 2, 120, 110, 3, 2, 2, 2, 120, 114, 3, 2, 2, 2, 120, 119, 3, 2, 2, 2,
	121, 11, 3, 2, 2, 2, 122, 123, 9, 2, 2, 2, 123, 13, 3, 2, 2, 2, 124, 132,
	5, 6, 4, 2, 125, 132, 5, 16, 9, 2, 126, 132, 5, 44, 23, 2, 127, 132, 5,
	62, 32, 2, 128, 132, 5, 46, 24, 2, 129, 132, 5, 68, 35, 2, 130, 132, 5,
	50, 26, 2, 131, 124, 3, 2, 2, 2, 131, 125, 3, 2, 2, 2, 131, 126, 3, 2,
	2, 2, 131, 127, 3, 2, 2, 2, 131, 128, 3, 2, 2, 2, 131, 129, 3, 2, 2, 2,
	131, 130, 3, 2, 2, 2, 132, 15, 3, 2, 2, 2, 133, 134, 7, 12, 2, 2, 134,
	135, 7, 45, 2, 2, 135, 139, 7, 35, 2, 2, 136, 138, 5, 18, 10, 2, 137, 136,
	3, 2, 2, 2, 138, 141, 3, 2, 2, 2, 139, 137, 3, 2, 2, 2, 139, 140, 3, 2,
	2, 2, 140, 142, 3, 2, 2, 2, 141, 139, 3, 2, 2, 2, 142, 143, 7, 36, 2, 2,
	143, 17, 3, 2, 2, 2, 144, 148, 5, 20, 11, 2, 145, 148, 5, 22, 12, 2, 146,
	148, 5, 24, 13, 2, 147, 144, 3, 2, 2, 2, 147, 145, 3, 2, 2, 2, 147, 146,
	3, 2, 2, 2, 148, 19, 3, 2, 2, 2, 149, 154, 7, 13, 2, 2, 150, 155, 7, 45,
	2, 2, 151, 152, 7, 26, 2, 2, 152, 153, 7, 42, 2, 2, 153, 155, 5, 38, 20,
	2, 154, 150, 3, 2, 2, 2, 154, 151, 3, 2, 2, 2, 155, 21, 3, 2, 2, 2, 156,
	157, 7, 14, 2, 2, 157, 158, 7, 37, 2, 2, 158, 159, 5, 34, 18, 2, 159, 160,
	7, 38, 2, 2, 160, 161, 7, 26, 2, 2, 161, 162, 7, 42, 2, 2, 162, 163, 5,
	38, 20, 2, 163, 23, 3, 2, 2, 2, 164, 166, 7, 15, 2, 2, 165, 167, 5, 40,
	21, 2, 166, 165, 3, 2, 2, 2, 166, 167, 3, 2, 2, 2, 167, 168, 3, 2, 2, 2,
	168, 169, 7, 41, 2, 2, 169, 171, 5, 28, 15, 2, 170, 172, 5, 26, 14, 2,
	171, 170, 3, 2, 2, 2, 171, 172, 3, 2, 2, 2, 172, 25, 3, 2, 2, 2, 173, 174,
	7, 16, 2, 2, 174, 175, 7, 17, 2, 2, 175, 176, 7, 33, 2, 2, 176, 177, 5,
	42, 22, 2, 177, 178, 7, 34, 2, 2, 178, 27, 3, 2, 2, 2, 179, 182, 5, 30,
	16, 2, 180, 182, 5, 32, 17, 2, 181, 179, 3, 2, 2, 2, 181, 180, 3, 2, 2,
	2, 182, 29, 3, 2, 2, 2, 183, 184, 7, 10, 2, 2, 184, 185, 5, 38, 20, 2,
	185, 31, 3, 2, 2, 2, 186, 187, 7, 11, 2, 2, 187, 188, 7, 26, 2, 2, 188,
	189, 7, 42, 2, 2, 189, 190, 5, 38, 20, 2, 190, 33, 3, 2, 2, 2, 191, 192,
	7, 45, 2, 2, 192, 35, 3, 2, 2, 2, 193, 194, 7, 45, 2, 2, 194, 37, 3, 2,
	2, 2, 195, 196, 7, 45, 2, 2, 196, 39, 3, 2, 2, 2, 197, 198, 7, 45, 2, 2,
	198, 41, 3, 2, 2, 2, 199, 200, 7, 45, 2, 2, 200, 43, 3, 2, 2, 2, 201, 202,
	7, 22, 2, 2, 202, 203, 7, 45, 2, 2, 203, 207, 7, 35, 2, 2, 204, 206, 5,
	48, 25, 2, 205, 204, 3, 2, 2, 2, 206, 209, 3, 2, 2, 2, 207, 205, 3, 2,
	2, 2, 207, 208, 3, 2, 2, 2, 208, 210, 3, 2, 2, 2, 209, 207, 3, 2, 2, 2,
	210, 211, 7, 36, 2, 2, 211, 45, 3, 2, 2, 2, 212, 213, 7, 23, 2, 2, 213,
	214, 7, 45, 2, 2, 214, 218, 7, 35, 2, 2, 215, 217, 5, 48, 25, 2, 216, 215,
	3, 2, 2, 2, 217, 220, 3, 2, 2, 2, 218, 216, 3, 2, 2, 2, 218, 219, 3, 2,
	2, 2, 219, 221, 3, 2, 2, 2, 220, 218, 3, 2, 2, 2, 221, 222, 7, 36, 2, 2,
	222, 47, 3, 2, 2, 2, 223, 228, 5, 38, 20, 2, 224, 225, 7, 43, 2, 2, 225,
	227, 5, 38, 20, 2, 226, 224, 3, 2, 2, 2, 227, 230, 3, 2, 2, 2, 228, 226,
	3, 2, 2, 2, 228, 229, 3, 2, 2, 2, 229, 236, 3, 2, 2, 2, 230, 228, 3, 2,
	2, 2, 231, 232, 5, 8, 5, 2, 232, 233, 7, 41, 2, 2, 233, 234, 5, 10, 6,
	2, 234, 236, 3, 2, 2, 2, 235, 223, 3, 2, 2, 2, 235, 231, 3, 2, 2, 2, 236,
	49, 3, 2, 2, 2, 237, 238, 7, 20, 2, 2, 238, 239, 7, 45, 2, 2, 239, 243,
	7, 35, 2, 2, 240, 242, 5, 52, 27, 2, 241, 240, 3, 2, 2, 2, 242, 245, 3,
	2, 2, 2, 243, 241, 3, 2, 2, 2, 243, 244, 3, 2, 2, 2, 244, 246, 3, 2, 2,
	2, 245, 243, 3, 2, 2, 2, 246, 247, 7, 36, 2, 2, 247, 51, 3, 2, 2, 2, 248,
	252, 7, 6, 2, 2, 249, 251, 7, 6, 2, 2, 250, 249, 3, 2, 2, 2, 251, 254,
	3, 2, 2, 2, 252, 250, 3, 2, 2, 2, 252, 253, 3, 2, 2, 2, 253, 259, 3, 2,
	2, 2, 254, 252, 3, 2, 2, 2, 255, 256, 5, 54, 28, 2, 256, 257, 7, 7, 2,
	2, 257, 259, 3, 2, 2, 2, 258, 248, 3, 2, 2, 2, 258, 255, 3, 2, 2, 2, 259,
	53, 3, 2, 2, 2, 260, 264, 5, 56, 29, 2, 261, 263, 5, 56, 29, 2, 262, 261,
	3, 2, 2, 2, 263, 266, 3, 2, 2, 2, 264, 262, 3, 2, 2, 2, 264, 265, 3, 2,
	2, 2, 265, 55, 3, 2, 2, 2, 266, 264, 3, 2, 2, 2, 267, 268, 7, 7, 2, 2,
	268, 269, 5, 58, 30, 2, 269, 57, 3, 2, 2, 2, 270, 281, 7, 48, 2, 2, 271,
	281, 7, 21, 2, 2, 272, 277, 5, 38, 20, 2, 273, 274, 7, 33, 2, 2, 274, 275,
	5, 60, 31, 2, 275, 276, 7, 34, 2, 2, 276, 278, 3, 2, 2, 2, 277, 273, 3,
	2, 2, 2, 277, 278, 3, 2, 2, 2, 278, 281, 3, 2, 2, 2, 279, 281, 7, 26, 2,
	2, 280, 270, 3, 2, 2, 2, 280, 271, 3, 2, 2, 2, 280, 272, 3, 2, 2, 2, 280,
	279, 3, 2, 2, 2, 281, 59, 3, 2, 2, 2, 282, 283, 9, 3, 2, 2, 283, 61, 3,
	2, 2, 2, 284, 285, 7, 24, 2, 2, 285, 286, 5, 64, 33, 2, 286, 287, 7, 35,
	2, 2, 287, 288, 5, 66, 34, 2, 288, 289, 7, 36, 2, 2, 289, 63, 3, 2, 2,
	2, 290, 291, 7, 45, 2, 2, 291, 65, 3, 2, 2, 2, 292, 293, 5, 6, 4, 2, 293,
	294, 7, 8, 2, 2, 294, 296, 3, 2, 2, 2, 295, 292, 3, 2, 2, 2, 296, 299,
	3, 2, 2, 2, 297, 295, 3, 2, 2, 2, 297, 298, 3, 2, 2, 2, 298, 67, 3, 2,
	2, 2, 299, 297, 3, 2, 2, 2, 300, 301, 7, 25, 2, 2, 301, 302, 5, 82, 42,
	2, 302, 306, 7, 35, 2, 2, 303, 305, 5, 70, 36, 2, 304, 303, 3, 2, 2, 2,
	305, 308, 3, 2, 2, 2, 306, 304, 3, 2, 2, 2, 306, 307, 3, 2, 2, 2, 307,
	309, 3, 2, 2, 2, 308, 306, 3, 2, 2, 2, 309, 310, 7, 36, 2, 2, 310, 69,
	3, 2, 2, 2, 311, 312, 5, 74, 38, 2, 312, 315, 7, 9, 2, 2, 313, 316, 5,
	76, 39, 2, 314, 316, 5, 78, 40, 2, 315, 313, 3, 2, 2, 2, 315, 314, 3, 2,
	2, 2, 316, 318, 3, 2, 2, 2, 317, 319, 7, 8, 2, 2, 318, 317, 3, 2, 2, 2,
	318, 319, 3, 2, 2, 2, 319, 331, 3, 2, 2, 2, 320, 321, 5, 74, 38, 2, 321,
	325, 7, 35, 2, 2, 322, 324, 5, 72, 37, 2, 323, 322, 3, 2, 2, 2, 324, 327,
	3, 2, 2, 2, 325, 323, 3, 2, 2, 2, 325, 326, 3, 2, 2, 2, 326, 328, 3, 2,
	2, 2, 327, 325, 3, 2, 2, 2, 328, 329, 7, 36, 2, 2, 329, 331, 3, 2, 2, 2,
	330, 311, 3, 2, 2, 2, 330, 320, 3, 2, 2, 2, 331, 71, 3, 2, 2, 2, 332, 333,
	5, 74, 38, 2, 333, 334, 7, 9, 2, 2, 334, 335, 5, 76, 39, 2, 335, 73, 3,
	2, 2, 2, 336, 337, 7, 45, 2, 2, 337, 75, 3, 2, 2, 2, 338, 339, 5, 10, 6,
	2, 339, 77, 3, 2, 2, 2, 340, 341, 7, 37, 2, 2, 341, 346, 5, 80, 41, 2,
	342, 343, 7, 43, 2, 2, 343, 345, 5, 80, 41, 2, 344, 342, 3, 2, 2, 2, 345,
	348, 3, 2, 2, 2, 346, 344, 3, 2, 2, 2, 346, 347, 3, 2, 2, 2, 347, 349,
	3, 2, 2, 2, 348, 346, 3, 2, 2, 2, 349, 350, 7, 38, 2, 2, 350, 79, 3, 2,
	2, 2, 351, 352, 5, 82, 42, 2, 352, 353, 7, 42, 2, 2, 353, 354, 7, 45, 2,
	2, 354, 81, 3, 2, 2, 2, 355, 356, 7, 45, 2, 2, 356, 83, 3, 2, 2, 2, 33,
	87, 89, 104, 108, 112, 117, 120, 131, 139, 147, 154, 166, 171, 181, 207,
	218, 228, 235, 243, 252, 258, 264, 277, 280, 297, 306, 315, 318, 325, 330,
	346,
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
	"styleName", "styleBody", "libraryDeclaration", "libraryExpress", "keyValue",
	"presetKey", "presetValue", "presetArray", "presetCall", "libraryName",
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
	DesignParserRULE_libraryExpress           = 34
	DesignParserRULE_keyValue                 = 35
	DesignParserRULE_presetKey                = 36
	DesignParserRULE_presetValue              = 37
	DesignParserRULE_presetArray              = 38
	DesignParserRULE_presetCall               = 39
	DesignParserRULE_libraryName              = 40
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
	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<DesignParserFLOW)|(1<<DesignParserLAYOUT)|(1<<DesignParserPAGE)|(1<<DesignParserCOMPONENT)|(1<<DesignParserSTYLE)|(1<<DesignParserLIBRARY))) != 0) || _la == DesignParserIDENTIFIER {
		p.SetState(85)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(82)
				p.Comment()
			}

		case 2:
			{
				p.SetState(83)
				p.ConfigDeclaration()
			}

		case 3:
			{
				p.SetState(84)
				p.Decalartions()
			}

		}

		p.SetState(89)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(90)
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
		p.SetState(92)
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
		p.SetState(94)
		p.ConfigKey()
	}
	{
		p.SetState(95)
		p.Match(DesignParserCOLON)
	}
	{
		p.SetState(96)
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
		p.SetState(98)
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

	p.SetState(118)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DesignParserDIGITS_IDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(100)
			p.Match(DesignParserDIGITS_IDENTIFIER)
		}
		p.SetState(102)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<DesignParserT__0)|(1<<DesignParserT__1)|(1<<DesignParserT__2))) != 0 {
			{
				p.SetState(101)
				p.Unit()
			}

		}

	case DesignParserDECIMAL_LITERAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(104)
			p.Match(DesignParserDECIMAL_LITERAL)
		}
		p.SetState(106)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<DesignParserT__0)|(1<<DesignParserT__1)|(1<<DesignParserT__2))) != 0 {
			{
				p.SetState(105)
				p.Unit()
			}

		}

	case DesignParserFLOAT_LITERAL:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(108)
			p.Match(DesignParserFLOAT_LITERAL)
		}
		p.SetState(110)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<DesignParserT__0)|(1<<DesignParserT__1)|(1<<DesignParserT__2))) != 0 {
			{
				p.SetState(109)
				p.Unit()
			}

		}

	case DesignParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(112)
			p.Match(DesignParserIDENTIFIER)
		}
		p.SetState(115)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DesignParserCOMMA {
			{
				p.SetState(113)
				p.Match(DesignParserCOMMA)
			}
			{
				p.SetState(114)
				p.Match(DesignParserIDENTIFIER)
			}

		}

	case DesignParserSTRING_LITERAL:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(117)
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
		p.SetState(120)
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

	p.SetState(129)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DesignParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(122)
			p.ConfigDeclaration()
		}

	case DesignParserFLOW:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(123)
			p.FlowDeclaration()
		}

	case DesignParserPAGE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(124)
			p.PageDeclaration()
		}

	case DesignParserSTYLE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(125)
			p.StyleDeclaration()
		}

	case DesignParserCOMPONENT:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(126)
			p.ComponentDeclaration()
		}

	case DesignParserLIBRARY:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(127)
			p.LibraryDeclaration()
		}

	case DesignParserLAYOUT:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(128)
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
		p.SetState(131)
		p.Match(DesignParserFLOW)
	}
	{
		p.SetState(132)
		p.Match(DesignParserIDENTIFIER)
	}
	{
		p.SetState(133)
		p.Match(DesignParserLBRACE)
	}
	p.SetState(137)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<DesignParserSEE)|(1<<DesignParserDO)|(1<<DesignParserREACT))) != 0 {
		{
			p.SetState(134)
			p.InteractionDeclaration()
		}

		p.SetState(139)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(140)
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

	p.SetState(145)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DesignParserSEE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(142)
			p.SeeDeclaration()
		}

	case DesignParserDO:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(143)
			p.DoDeclaration()
		}

	case DesignParserREACT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(144)
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
		p.SetState(147)
		p.Match(DesignParserSEE)
	}
	p.SetState(152)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DesignParserIDENTIFIER:
		{
			p.SetState(148)
			p.Match(DesignParserIDENTIFIER)
		}

	case DesignParserSTRING_LITERAL:
		{
			p.SetState(149)
			p.Match(DesignParserSTRING_LITERAL)
		}
		{
			p.SetState(150)
			p.Match(DesignParserDOT)
		}
		{
			p.SetState(151)
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
		p.SetState(154)
		p.Match(DesignParserDO)
	}
	{
		p.SetState(155)
		p.Match(DesignParserLBRACK)
	}
	{
		p.SetState(156)
		p.ActionName()
	}
	{
		p.SetState(157)
		p.Match(DesignParserRBRACK)
	}
	{
		p.SetState(158)
		p.Match(DesignParserSTRING_LITERAL)
	}
	{
		p.SetState(159)
		p.Match(DesignParserDOT)
	}
	{
		p.SetState(160)
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
		p.SetState(162)
		p.Match(DesignParserREACT)
	}
	p.SetState(164)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DesignParserIDENTIFIER {
		{
			p.SetState(163)
			p.SceneName()
		}

	}
	{
		p.SetState(166)
		p.Match(DesignParserCOLON)
	}
	{
		p.SetState(167)
		p.ReactAction()
	}
	p.SetState(169)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DesignParserWITHTEXT {
		{
			p.SetState(168)
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
		p.SetState(171)
		p.Match(DesignParserWITHTEXT)
	}
	{
		p.SetState(172)
		p.Match(DesignParserANIMATE)
	}
	{
		p.SetState(173)
		p.Match(DesignParserLPAREN)
	}
	{
		p.SetState(174)
		p.AnimateName()
	}
	{
		p.SetState(175)
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

	p.SetState(179)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DesignParserGOTO_KEY:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(177)
			p.GotoAction()
		}

	case DesignParserSHOW_KEY:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(178)
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
		p.SetState(181)
		p.Match(DesignParserGOTO_KEY)
	}
	{
		p.SetState(182)
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
		p.SetState(184)
		p.Match(DesignParserSHOW_KEY)
	}
	{
		p.SetState(185)
		p.Match(DesignParserSTRING_LITERAL)
	}
	{
		p.SetState(186)
		p.Match(DesignParserDOT)
	}
	{
		p.SetState(187)
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
		p.SetState(189)
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
		p.SetState(191)
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
		p.SetState(193)
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
		p.SetState(195)
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
		p.SetState(197)
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
		p.SetState(199)
		p.Match(DesignParserPAGE)
	}
	{
		p.SetState(200)
		p.Match(DesignParserIDENTIFIER)
	}
	{
		p.SetState(201)
		p.Match(DesignParserLBRACE)
	}
	p.SetState(205)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DesignParserIDENTIFIER {
		{
			p.SetState(202)
			p.ComponentBodyDeclaration()
		}

		p.SetState(207)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(208)
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
		p.SetState(210)
		p.Match(DesignParserCOMPONENT)
	}
	{
		p.SetState(211)
		p.Match(DesignParserIDENTIFIER)
	}
	{
		p.SetState(212)
		p.Match(DesignParserLBRACE)
	}
	p.SetState(216)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DesignParserIDENTIFIER {
		{
			p.SetState(213)
			p.ComponentBodyDeclaration()
		}

		p.SetState(218)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(219)
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

	p.SetState(233)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(221)
			p.ComponentName()
		}
		p.SetState(226)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == DesignParserCOMMA {
			{
				p.SetState(222)
				p.Match(DesignParserCOMMA)
			}
			{
				p.SetState(223)
				p.ComponentName()
			}

			p.SetState(228)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(229)
			p.ConfigKey()
		}
		{
			p.SetState(230)
			p.Match(DesignParserCOLON)
		}
		{
			p.SetState(231)
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
		p.SetState(235)
		p.Match(DesignParserLAYOUT)
	}
	{
		p.SetState(236)
		p.Match(DesignParserIDENTIFIER)
	}
	{
		p.SetState(237)
		p.Match(DesignParserLBRACE)
	}
	p.SetState(241)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DesignParserT__3 || _la == DesignParserT__4 {
		{
			p.SetState(238)
			p.LayoutRow()
		}

		p.SetState(243)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(244)
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

	p.SetState(256)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DesignParserT__3:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(246)
			p.Match(DesignParserT__3)
		}
		p.SetState(250)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(247)
					p.Match(DesignParserT__3)
				}

			}
			p.SetState(252)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext())
		}

	case DesignParserT__4:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(253)
			p.LayoutLines()
		}
		{
			p.SetState(254)
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
		p.SetState(258)
		p.LayoutLine()
	}
	p.SetState(262)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(259)
				p.LayoutLine()
			}

		}
		p.SetState(264)
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
		p.SetState(265)
		p.Match(DesignParserT__4)
	}
	{
		p.SetState(266)
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

	p.SetState(278)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DesignParserDECIMAL_LITERAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(268)
			p.Match(DesignParserDECIMAL_LITERAL)
		}

	case DesignParserPOSITION:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(269)
			p.Match(DesignParserPOSITION)
		}

	case DesignParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(270)
			p.ComponentName()
		}
		p.SetState(275)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DesignParserLPAREN {
			{
				p.SetState(271)
				p.Match(DesignParserLPAREN)
			}
			{
				p.SetState(272)
				p.ComponentLayoutValue()
			}
			{
				p.SetState(273)
				p.Match(DesignParserRPAREN)
			}

		}

	case DesignParserSTRING_LITERAL:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(277)
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
		p.SetState(280)
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
		p.SetState(282)
		p.Match(DesignParserSTYLE)
	}
	{
		p.SetState(283)
		p.StyleName()
	}
	{
		p.SetState(284)
		p.Match(DesignParserLBRACE)
	}
	{
		p.SetState(285)
		p.StyleBody()
	}
	{
		p.SetState(286)
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
		p.SetState(288)
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
	p.SetState(295)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DesignParserIDENTIFIER {
		{
			p.SetState(290)
			p.ConfigDeclaration()
		}
		{
			p.SetState(291)
			p.Match(DesignParserT__5)
		}

		p.SetState(297)
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

func (s *LibraryDeclarationContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(DesignParserRBRACE, 0)
}

func (s *LibraryDeclarationContext) AllLibraryExpress() []ILibraryExpressContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILibraryExpressContext)(nil)).Elem())
	var tst = make([]ILibraryExpressContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILibraryExpressContext)
		}
	}

	return tst
}

func (s *LibraryDeclarationContext) LibraryExpress(i int) ILibraryExpressContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILibraryExpressContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILibraryExpressContext)
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
		p.SetState(298)
		p.Match(DesignParserLIBRARY)
	}
	{
		p.SetState(299)
		p.LibraryName()
	}
	{
		p.SetState(300)
		p.Match(DesignParserLBRACE)
	}
	p.SetState(304)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DesignParserIDENTIFIER {
		{
			p.SetState(301)
			p.LibraryExpress()
		}

		p.SetState(306)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(307)
		p.Match(DesignParserRBRACE)
	}

	return localctx
}

// ILibraryExpressContext is an interface to support dynamic dispatch.
type ILibraryExpressContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLibraryExpressContext differentiates from other interfaces.
	IsLibraryExpressContext()
}

type LibraryExpressContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLibraryExpressContext() *LibraryExpressContext {
	var p = new(LibraryExpressContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DesignParserRULE_libraryExpress
	return p
}

func (*LibraryExpressContext) IsLibraryExpressContext() {}

func NewLibraryExpressContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LibraryExpressContext {
	var p = new(LibraryExpressContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesignParserRULE_libraryExpress

	return p
}

func (s *LibraryExpressContext) GetParser() antlr.Parser { return s.parser }

func (s *LibraryExpressContext) PresetKey() IPresetKeyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPresetKeyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPresetKeyContext)
}

func (s *LibraryExpressContext) PresetValue() IPresetValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPresetValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPresetValueContext)
}

func (s *LibraryExpressContext) PresetArray() IPresetArrayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPresetArrayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPresetArrayContext)
}

func (s *LibraryExpressContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(DesignParserLBRACE, 0)
}

func (s *LibraryExpressContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(DesignParserRBRACE, 0)
}

func (s *LibraryExpressContext) AllKeyValue() []IKeyValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IKeyValueContext)(nil)).Elem())
	var tst = make([]IKeyValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IKeyValueContext)
		}
	}

	return tst
}

func (s *LibraryExpressContext) KeyValue(i int) IKeyValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKeyValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IKeyValueContext)
}

func (s *LibraryExpressContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LibraryExpressContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LibraryExpressContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.EnterLibraryExpress(s)
	}
}

func (s *LibraryExpressContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.ExitLibraryExpress(s)
	}
}

func (s *LibraryExpressContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DesignVisitor:
		return t.VisitLibraryExpress(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DesignParser) LibraryExpress() (localctx ILibraryExpressContext) {
	localctx = NewLibraryExpressContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, DesignParserRULE_libraryExpress)
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

	p.SetState(328)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(309)
			p.PresetKey()
		}
		{
			p.SetState(310)
			p.Match(DesignParserT__6)
		}
		p.SetState(313)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case DesignParserSTRING_LITERAL, DesignParserIDENTIFIER, DesignParserDIGITS_IDENTIFIER, DesignParserDECIMAL_LITERAL, DesignParserFLOAT_LITERAL:
			{
				p.SetState(311)
				p.PresetValue()
			}

		case DesignParserLBRACK:
			{
				p.SetState(312)
				p.PresetArray()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}
		p.SetState(316)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DesignParserT__5 {
			{
				p.SetState(315)
				p.Match(DesignParserT__5)
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(318)
			p.PresetKey()
		}
		{
			p.SetState(319)
			p.Match(DesignParserLBRACE)
		}
		p.SetState(323)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == DesignParserIDENTIFIER {
			{
				p.SetState(320)
				p.KeyValue()
			}

			p.SetState(325)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(326)
			p.Match(DesignParserRBRACE)
		}

	}

	return localctx
}

// IKeyValueContext is an interface to support dynamic dispatch.
type IKeyValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKeyValueContext differentiates from other interfaces.
	IsKeyValueContext()
}

type KeyValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeyValueContext() *KeyValueContext {
	var p = new(KeyValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DesignParserRULE_keyValue
	return p
}

func (*KeyValueContext) IsKeyValueContext() {}

func NewKeyValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeyValueContext {
	var p = new(KeyValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesignParserRULE_keyValue

	return p
}

func (s *KeyValueContext) GetParser() antlr.Parser { return s.parser }

func (s *KeyValueContext) PresetKey() IPresetKeyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPresetKeyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPresetKeyContext)
}

func (s *KeyValueContext) PresetValue() IPresetValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPresetValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPresetValueContext)
}

func (s *KeyValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeyValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeyValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.EnterKeyValue(s)
	}
}

func (s *KeyValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.ExitKeyValue(s)
	}
}

func (s *KeyValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DesignVisitor:
		return t.VisitKeyValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DesignParser) KeyValue() (localctx IKeyValueContext) {
	localctx = NewKeyValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, DesignParserRULE_keyValue)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.SetState(330)
		p.PresetKey()
	}
	{
		p.SetState(331)
		p.Match(DesignParserT__6)
	}
	{
		p.SetState(332)
		p.PresetValue()
	}

	return localctx
}

// IPresetKeyContext is an interface to support dynamic dispatch.
type IPresetKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPresetKeyContext differentiates from other interfaces.
	IsPresetKeyContext()
}

type PresetKeyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPresetKeyContext() *PresetKeyContext {
	var p = new(PresetKeyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DesignParserRULE_presetKey
	return p
}

func (*PresetKeyContext) IsPresetKeyContext() {}

func NewPresetKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PresetKeyContext {
	var p = new(PresetKeyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesignParserRULE_presetKey

	return p
}

func (s *PresetKeyContext) GetParser() antlr.Parser { return s.parser }

func (s *PresetKeyContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(DesignParserIDENTIFIER, 0)
}

func (s *PresetKeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PresetKeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PresetKeyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.EnterPresetKey(s)
	}
}

func (s *PresetKeyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.ExitPresetKey(s)
	}
}

func (s *PresetKeyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DesignVisitor:
		return t.VisitPresetKey(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DesignParser) PresetKey() (localctx IPresetKeyContext) {
	localctx = NewPresetKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, DesignParserRULE_presetKey)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.SetState(334)
		p.Match(DesignParserIDENTIFIER)
	}

	return localctx
}

// IPresetValueContext is an interface to support dynamic dispatch.
type IPresetValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPresetValueContext differentiates from other interfaces.
	IsPresetValueContext()
}

type PresetValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPresetValueContext() *PresetValueContext {
	var p = new(PresetValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DesignParserRULE_presetValue
	return p
}

func (*PresetValueContext) IsPresetValueContext() {}

func NewPresetValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PresetValueContext {
	var p = new(PresetValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesignParserRULE_presetValue

	return p
}

func (s *PresetValueContext) GetParser() antlr.Parser { return s.parser }

func (s *PresetValueContext) ConfigValue() IConfigValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConfigValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConfigValueContext)
}

func (s *PresetValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PresetValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PresetValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.EnterPresetValue(s)
	}
}

func (s *PresetValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.ExitPresetValue(s)
	}
}

func (s *PresetValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DesignVisitor:
		return t.VisitPresetValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DesignParser) PresetValue() (localctx IPresetValueContext) {
	localctx = NewPresetValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, DesignParserRULE_presetValue)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.SetState(336)
		p.ConfigValue()
	}

	return localctx
}

// IPresetArrayContext is an interface to support dynamic dispatch.
type IPresetArrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPresetArrayContext differentiates from other interfaces.
	IsPresetArrayContext()
}

type PresetArrayContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPresetArrayContext() *PresetArrayContext {
	var p = new(PresetArrayContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DesignParserRULE_presetArray
	return p
}

func (*PresetArrayContext) IsPresetArrayContext() {}

func NewPresetArrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PresetArrayContext {
	var p = new(PresetArrayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesignParserRULE_presetArray

	return p
}

func (s *PresetArrayContext) GetParser() antlr.Parser { return s.parser }

func (s *PresetArrayContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(DesignParserLBRACK, 0)
}

func (s *PresetArrayContext) AllPresetCall() []IPresetCallContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPresetCallContext)(nil)).Elem())
	var tst = make([]IPresetCallContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPresetCallContext)
		}
	}

	return tst
}

func (s *PresetArrayContext) PresetCall(i int) IPresetCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPresetCallContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPresetCallContext)
}

func (s *PresetArrayContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(DesignParserRBRACK, 0)
}

func (s *PresetArrayContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(DesignParserCOMMA)
}

func (s *PresetArrayContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(DesignParserCOMMA, i)
}

func (s *PresetArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PresetArrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PresetArrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.EnterPresetArray(s)
	}
}

func (s *PresetArrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.ExitPresetArray(s)
	}
}

func (s *PresetArrayContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DesignVisitor:
		return t.VisitPresetArray(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DesignParser) PresetArray() (localctx IPresetArrayContext) {
	localctx = NewPresetArrayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, DesignParserRULE_presetArray)
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
		p.SetState(338)
		p.Match(DesignParserLBRACK)
	}
	{
		p.SetState(339)
		p.PresetCall()
	}
	p.SetState(344)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DesignParserCOMMA {
		{
			p.SetState(340)
			p.Match(DesignParserCOMMA)
		}
		{
			p.SetState(341)
			p.PresetCall()
		}

		p.SetState(346)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(347)
		p.Match(DesignParserRBRACK)
	}

	return localctx
}

// IPresetCallContext is an interface to support dynamic dispatch.
type IPresetCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPresetCallContext differentiates from other interfaces.
	IsPresetCallContext()
}

type PresetCallContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPresetCallContext() *PresetCallContext {
	var p = new(PresetCallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DesignParserRULE_presetCall
	return p
}

func (*PresetCallContext) IsPresetCallContext() {}

func NewPresetCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PresetCallContext {
	var p = new(PresetCallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesignParserRULE_presetCall

	return p
}

func (s *PresetCallContext) GetParser() antlr.Parser { return s.parser }

func (s *PresetCallContext) LibraryName() ILibraryNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILibraryNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILibraryNameContext)
}

func (s *PresetCallContext) DOT() antlr.TerminalNode {
	return s.GetToken(DesignParserDOT, 0)
}

func (s *PresetCallContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(DesignParserIDENTIFIER, 0)
}

func (s *PresetCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PresetCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PresetCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.EnterPresetCall(s)
	}
}

func (s *PresetCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.ExitPresetCall(s)
	}
}

func (s *PresetCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DesignVisitor:
		return t.VisitPresetCall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DesignParser) PresetCall() (localctx IPresetCallContext) {
	localctx = NewPresetCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, DesignParserRULE_presetCall)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.SetState(349)
		p.LibraryName()
	}
	{
		p.SetState(350)
		p.Match(DesignParserDOT)
	}
	{
		p.SetState(351)
		p.Match(DesignParserIDENTIFIER)
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
	p.EnterRule(localctx, 80, DesignParserRULE_libraryName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.SetState(353)
		p.Match(DesignParserIDENTIFIER)
	}

	return localctx
}
