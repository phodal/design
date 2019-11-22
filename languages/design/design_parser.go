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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 32, 170,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	3, 2, 7, 2, 48, 10, 2, 12, 2, 14, 2, 51, 11, 2, 3, 2, 3, 2, 7, 2, 55, 10,
	2, 12, 2, 14, 2, 58, 11, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 5, 3, 69, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 5, 5, 76, 10, 5,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7,
	3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 10, 3, 10, 7, 10, 104, 10, 10, 12, 10, 14, 10, 107, 11, 10, 3, 10, 3,
	10, 3, 11, 7, 11, 112, 10, 11, 12, 11, 14, 11, 115, 11, 11, 3, 12, 7, 12,
	118, 10, 12, 12, 12, 14, 12, 121, 11, 12, 3, 13, 7, 13, 124, 10, 13, 12,
	13, 14, 13, 127, 11, 13, 3, 14, 7, 14, 130, 10, 14, 12, 14, 14, 14, 133,
	11, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 5, 16, 142, 10,
	16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19,
	3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 22, 3, 22, 3,
	23, 3, 23, 5, 23, 166, 10, 23, 3, 23, 3, 23, 3, 23, 2, 2, 24, 2, 4, 6,
	8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42,
	44, 2, 2, 2, 166, 2, 49, 3, 2, 2, 2, 4, 68, 3, 2, 2, 2, 6, 70, 3, 2, 2,
	2, 8, 75, 3, 2, 2, 2, 10, 77, 3, 2, 2, 2, 12, 83, 3, 2, 2, 2, 14, 89, 3,
	2, 2, 2, 16, 95, 3, 2, 2, 2, 18, 101, 3, 2, 2, 2, 20, 113, 3, 2, 2, 2,
	22, 119, 3, 2, 2, 2, 24, 125, 3, 2, 2, 2, 26, 131, 3, 2, 2, 2, 28, 134,
	3, 2, 2, 2, 30, 141, 3, 2, 2, 2, 32, 143, 3, 2, 2, 2, 34, 147, 3, 2, 2,
	2, 36, 151, 3, 2, 2, 2, 38, 155, 3, 2, 2, 2, 40, 158, 3, 2, 2, 2, 42, 161,
	3, 2, 2, 2, 44, 163, 3, 2, 2, 2, 46, 48, 5, 44, 23, 2, 47, 46, 3, 2, 2,
	2, 48, 51, 3, 2, 2, 2, 49, 47, 3, 2, 2, 2, 49, 50, 3, 2, 2, 2, 50, 52,
	3, 2, 2, 2, 51, 49, 3, 2, 2, 2, 52, 56, 5, 4, 3, 2, 53, 55, 5, 4, 3, 2,
	54, 53, 3, 2, 2, 2, 55, 58, 3, 2, 2, 2, 56, 54, 3, 2, 2, 2, 56, 57, 3,
	2, 2, 2, 57, 59, 3, 2, 2, 2, 58, 56, 3, 2, 2, 2, 59, 60, 7, 2, 2, 3, 60,
	3, 3, 2, 2, 2, 61, 69, 5, 6, 4, 2, 62, 69, 5, 10, 6, 2, 63, 69, 5, 28,
	15, 2, 64, 69, 5, 12, 7, 2, 65, 69, 5, 14, 8, 2, 66, 69, 5, 16, 9, 2, 67,
	69, 5, 18, 10, 2, 68, 61, 3, 2, 2, 2, 68, 62, 3, 2, 2, 2, 68, 63, 3, 2,
	2, 2, 68, 64, 3, 2, 2, 2, 68, 65, 3, 2, 2, 2, 68, 66, 3, 2, 2, 2, 68, 67,
	3, 2, 2, 2, 69, 5, 3, 2, 2, 2, 70, 71, 7, 8, 2, 2, 71, 72, 7, 3, 2, 2,
	72, 73, 7, 30, 2, 2, 73, 7, 3, 2, 2, 2, 74, 76, 5, 44, 23, 2, 75, 74, 3,
	2, 2, 2, 75, 76, 3, 2, 2, 2, 76, 9, 3, 2, 2, 2, 77, 78, 7, 9, 2, 2, 78,
	79, 7, 30, 2, 2, 79, 80, 7, 27, 2, 2, 80, 81, 5, 20, 11, 2, 81, 82, 7,
	28, 2, 2, 82, 11, 3, 2, 2, 2, 83, 84, 7, 16, 2, 2, 84, 85, 7, 30, 2, 2,
	85, 86, 7, 27, 2, 2, 86, 87, 5, 22, 12, 2, 87, 88, 7, 28, 2, 2, 88, 13,
	3, 2, 2, 2, 89, 90, 7, 17, 2, 2, 90, 91, 7, 30, 2, 2, 91, 92, 7, 27, 2,
	2, 92, 93, 5, 24, 13, 2, 93, 94, 7, 28, 2, 2, 94, 15, 3, 2, 2, 2, 95, 96,
	7, 12, 2, 2, 96, 97, 7, 30, 2, 2, 97, 98, 7, 27, 2, 2, 98, 99, 5, 26, 14,
	2, 99, 100, 7, 28, 2, 2, 100, 17, 3, 2, 2, 2, 101, 105, 7, 4, 2, 2, 102,
	104, 7, 32, 2, 2, 103, 102, 3, 2, 2, 2, 104, 107, 3, 2, 2, 2, 105, 103,
	3, 2, 2, 2, 105, 106, 3, 2, 2, 2, 106, 108, 3, 2, 2, 2, 107, 105, 3, 2,
	2, 2, 108, 109, 7, 4, 2, 2, 109, 19, 3, 2, 2, 2, 110, 112, 5, 30, 16, 2,
	111, 110, 3, 2, 2, 2, 112, 115, 3, 2, 2, 2, 113, 111, 3, 2, 2, 2, 113,
	114, 3, 2, 2, 2, 114, 21, 3, 2, 2, 2, 115, 113, 3, 2, 2, 2, 116, 118, 5,
	30, 16, 2, 117, 116, 3, 2, 2, 2, 118, 121, 3, 2, 2, 2, 119, 117, 3, 2,
	2, 2, 119, 120, 3, 2, 2, 2, 120, 23, 3, 2, 2, 2, 121, 119, 3, 2, 2, 2,
	122, 124, 5, 30, 16, 2, 123, 122, 3, 2, 2, 2, 124, 127, 3, 2, 2, 2, 125,
	123, 3, 2, 2, 2, 125, 126, 3, 2, 2, 2, 126, 25, 3, 2, 2, 2, 127, 125, 3,
	2, 2, 2, 128, 130, 5, 30, 16, 2, 129, 128, 3, 2, 2, 2, 130, 133, 3, 2,
	2, 2, 131, 129, 3, 2, 2, 2, 131, 132, 3, 2, 2, 2, 132, 27, 3, 2, 2, 2,
	133, 131, 3, 2, 2, 2, 134, 135, 5, 30, 16, 2, 135, 29, 3, 2, 2, 2, 136,
	142, 5, 32, 17, 2, 137, 142, 5, 34, 18, 2, 138, 142, 5, 36, 19, 2, 139,
	142, 5, 40, 21, 2, 140, 142, 5, 38, 20, 2, 141, 136, 3, 2, 2, 2, 141, 137,
	3, 2, 2, 2, 141, 138, 3, 2, 2, 2, 141, 139, 3, 2, 2, 2, 141, 140, 3, 2,
	2, 2, 142, 31, 3, 2, 2, 2, 143, 144, 7, 30, 2, 2, 144, 145, 7, 5, 2, 2,
	145, 146, 7, 30, 2, 2, 146, 33, 3, 2, 2, 2, 147, 148, 7, 30, 2, 2, 148,
	149, 7, 6, 2, 2, 149, 150, 7, 30, 2, 2, 150, 35, 3, 2, 2, 2, 151, 152,
	7, 30, 2, 2, 152, 153, 7, 3, 2, 2, 153, 154, 7, 30, 2, 2, 154, 37, 3, 2,
	2, 2, 155, 156, 7, 12, 2, 2, 156, 157, 7, 30, 2, 2, 157, 39, 3, 2, 2, 2,
	158, 159, 7, 16, 2, 2, 159, 160, 7, 30, 2, 2, 160, 41, 3, 2, 2, 2, 161,
	162, 7, 12, 2, 2, 162, 43, 3, 2, 2, 2, 163, 165, 7, 7, 2, 2, 164, 166,
	7, 25, 2, 2, 165, 164, 3, 2, 2, 2, 165, 166, 3, 2, 2, 2, 166, 167, 3, 2,
	2, 2, 167, 168, 7, 30, 2, 2, 168, 45, 3, 2, 2, 2, 13, 49, 56, 68, 75, 105,
	113, 119, 125, 131, 141, 165,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "':'", "'```'", "'='", "'->'", "'#'", "'DesignSystem'", "'Design'",
	"'Project'", "'Page'", "'Layer'", "'Function'", "'Library'", "'Unit'",
	"'Template'", "'Component'", "'position'", "'Flow'", "'behavior'", "",
	"", "", "", "", "", "'{'", "'}'", "'\"'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "DESIGN_SYSTEM", "DESIGN", "PROJECT", "Page", "LAYER",
	"Function", "Library", "Unit", "TEMPLATE", "COMPONENT", "Position", "Flow",
	"Behavior", "WS", "COMMENT", "LINE_COMMENT", "EmptyLine", "Space", "NewLine",
	"LBRACE", "RBRACE", "Quote", "IDENTIFIER", "STRING_LITERAL", "HTML_STRING",
}

var ruleNames = []string{
	"designIt", "declaration", "designSystemDeclaration", "commentBlockDeclaration",
	"designBlockDeclaration", "templateBlockDeclaration", "componentBlockDeclaration",
	"layerBlockDeclaration", "codeBlockDeclaration", "designBodyDeclaration",
	"templateBodyDeclaration", "componentBodyDeclaration", "layerBodyDeclaration",
	"expressDeclaration", "express", "equalExpress", "useExpress", "valueExpress",
	"layerExpress", "templateExpress", "layer", "commentDeclaration",
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
	DesignParserEOF            = antlr.TokenEOF
	DesignParserT__0           = 1
	DesignParserT__1           = 2
	DesignParserT__2           = 3
	DesignParserT__3           = 4
	DesignParserT__4           = 5
	DesignParserDESIGN_SYSTEM  = 6
	DesignParserDESIGN         = 7
	DesignParserPROJECT        = 8
	DesignParserPage           = 9
	DesignParserLAYER          = 10
	DesignParserFunction       = 11
	DesignParserLibrary        = 12
	DesignParserUnit           = 13
	DesignParserTEMPLATE       = 14
	DesignParserCOMPONENT      = 15
	DesignParserPosition       = 16
	DesignParserFlow           = 17
	DesignParserBehavior       = 18
	DesignParserWS             = 19
	DesignParserCOMMENT        = 20
	DesignParserLINE_COMMENT   = 21
	DesignParserEmptyLine      = 22
	DesignParserSpace          = 23
	DesignParserNewLine        = 24
	DesignParserLBRACE         = 25
	DesignParserRBRACE         = 26
	DesignParserQuote          = 27
	DesignParserIDENTIFIER     = 28
	DesignParserSTRING_LITERAL = 29
	DesignParserHTML_STRING    = 30
)

// DesignParser rules.
const (
	DesignParserRULE_designIt                  = 0
	DesignParserRULE_declaration               = 1
	DesignParserRULE_designSystemDeclaration   = 2
	DesignParserRULE_commentBlockDeclaration   = 3
	DesignParserRULE_designBlockDeclaration    = 4
	DesignParserRULE_templateBlockDeclaration  = 5
	DesignParserRULE_componentBlockDeclaration = 6
	DesignParserRULE_layerBlockDeclaration     = 7
	DesignParserRULE_codeBlockDeclaration      = 8
	DesignParserRULE_designBodyDeclaration     = 9
	DesignParserRULE_templateBodyDeclaration   = 10
	DesignParserRULE_componentBodyDeclaration  = 11
	DesignParserRULE_layerBodyDeclaration      = 12
	DesignParserRULE_expressDeclaration        = 13
	DesignParserRULE_express                   = 14
	DesignParserRULE_equalExpress              = 15
	DesignParserRULE_useExpress                = 16
	DesignParserRULE_valueExpress              = 17
	DesignParserRULE_layerExpress              = 18
	DesignParserRULE_templateExpress           = 19
	DesignParserRULE_layer                     = 20
	DesignParserRULE_commentDeclaration        = 21
)

// IDesignItContext is an interface to support dynamic dispatch.
type IDesignItContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDesignItContext differentiates from other interfaces.
	IsDesignItContext()
}

type DesignItContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDesignItContext() *DesignItContext {
	var p = new(DesignItContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DesignParserRULE_designIt
	return p
}

func (*DesignItContext) IsDesignItContext() {}

func NewDesignItContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DesignItContext {
	var p = new(DesignItContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesignParserRULE_designIt

	return p
}

func (s *DesignItContext) GetParser() antlr.Parser { return s.parser }

func (s *DesignItContext) AllDeclaration() []IDeclarationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDeclarationContext)(nil)).Elem())
	var tst = make([]IDeclarationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDeclarationContext)
		}
	}

	return tst
}

func (s *DesignItContext) Declaration(i int) IDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclarationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *DesignItContext) EOF() antlr.TerminalNode {
	return s.GetToken(DesignParserEOF, 0)
}

func (s *DesignItContext) AllCommentDeclaration() []ICommentDeclarationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICommentDeclarationContext)(nil)).Elem())
	var tst = make([]ICommentDeclarationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICommentDeclarationContext)
		}
	}

	return tst
}

func (s *DesignItContext) CommentDeclaration(i int) ICommentDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommentDeclarationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICommentDeclarationContext)
}

func (s *DesignItContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DesignItContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DesignItContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.EnterDesignIt(s)
	}
}

func (s *DesignItContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.ExitDesignIt(s)
	}
}

func (s *DesignItContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DesignVisitor:
		return t.VisitDesignIt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DesignParser) DesignIt() (localctx IDesignItContext) {
	localctx = NewDesignItContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, DesignParserRULE_designIt)
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
	p.SetState(47)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DesignParserT__4 {
		{
			p.SetState(44)
			p.CommentDeclaration()
		}

		p.SetState(49)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(50)
		p.Declaration()
	}
	p.SetState(54)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<DesignParserT__1)|(1<<DesignParserDESIGN_SYSTEM)|(1<<DesignParserDESIGN)|(1<<DesignParserLAYER)|(1<<DesignParserTEMPLATE)|(1<<DesignParserCOMPONENT)|(1<<DesignParserIDENTIFIER))) != 0 {
		{
			p.SetState(51)
			p.Declaration()
		}

		p.SetState(56)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(57)
		p.Match(DesignParserEOF)
	}

	return localctx
}

// IDeclarationContext is an interface to support dynamic dispatch.
type IDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeclarationContext differentiates from other interfaces.
	IsDeclarationContext()
}

type DeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationContext() *DeclarationContext {
	var p = new(DeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DesignParserRULE_declaration
	return p
}

func (*DeclarationContext) IsDeclarationContext() {}

func NewDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationContext {
	var p = new(DeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesignParserRULE_declaration

	return p
}

func (s *DeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationContext) DesignSystemDeclaration() IDesignSystemDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDesignSystemDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDesignSystemDeclarationContext)
}

func (s *DeclarationContext) DesignBlockDeclaration() IDesignBlockDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDesignBlockDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDesignBlockDeclarationContext)
}

func (s *DeclarationContext) ExpressDeclaration() IExpressDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressDeclarationContext)
}

func (s *DeclarationContext) TemplateBlockDeclaration() ITemplateBlockDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITemplateBlockDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITemplateBlockDeclarationContext)
}

func (s *DeclarationContext) ComponentBlockDeclaration() IComponentBlockDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComponentBlockDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComponentBlockDeclarationContext)
}

func (s *DeclarationContext) LayerBlockDeclaration() ILayerBlockDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILayerBlockDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILayerBlockDeclarationContext)
}

func (s *DeclarationContext) CodeBlockDeclaration() ICodeBlockDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICodeBlockDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICodeBlockDeclarationContext)
}

func (s *DeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.EnterDeclaration(s)
	}
}

func (s *DeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.ExitDeclaration(s)
	}
}

func (s *DeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DesignVisitor:
		return t.VisitDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DesignParser) Declaration() (localctx IDeclarationContext) {
	localctx = NewDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, DesignParserRULE_declaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(66)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(59)
			p.DesignSystemDeclaration()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(60)
			p.DesignBlockDeclaration()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(61)
			p.ExpressDeclaration()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(62)
			p.TemplateBlockDeclaration()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(63)
			p.ComponentBlockDeclaration()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(64)
			p.LayerBlockDeclaration()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(65)
			p.CodeBlockDeclaration()
		}

	}

	return localctx
}

// IDesignSystemDeclarationContext is an interface to support dynamic dispatch.
type IDesignSystemDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDesignSystemDeclarationContext differentiates from other interfaces.
	IsDesignSystemDeclarationContext()
}

type DesignSystemDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDesignSystemDeclarationContext() *DesignSystemDeclarationContext {
	var p = new(DesignSystemDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DesignParserRULE_designSystemDeclaration
	return p
}

func (*DesignSystemDeclarationContext) IsDesignSystemDeclarationContext() {}

func NewDesignSystemDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DesignSystemDeclarationContext {
	var p = new(DesignSystemDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesignParserRULE_designSystemDeclaration

	return p
}

func (s *DesignSystemDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *DesignSystemDeclarationContext) DESIGN_SYSTEM() antlr.TerminalNode {
	return s.GetToken(DesignParserDESIGN_SYSTEM, 0)
}

func (s *DesignSystemDeclarationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(DesignParserIDENTIFIER, 0)
}

func (s *DesignSystemDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DesignSystemDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DesignSystemDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.EnterDesignSystemDeclaration(s)
	}
}

func (s *DesignSystemDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.ExitDesignSystemDeclaration(s)
	}
}

func (s *DesignSystemDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DesignVisitor:
		return t.VisitDesignSystemDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DesignParser) DesignSystemDeclaration() (localctx IDesignSystemDeclarationContext) {
	localctx = NewDesignSystemDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, DesignParserRULE_designSystemDeclaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.SetState(68)
		p.Match(DesignParserDESIGN_SYSTEM)
	}
	{
		p.SetState(69)
		p.Match(DesignParserT__0)
	}
	{
		p.SetState(70)
		p.Match(DesignParserIDENTIFIER)
	}

	return localctx
}

// ICommentBlockDeclarationContext is an interface to support dynamic dispatch.
type ICommentBlockDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCommentBlockDeclarationContext differentiates from other interfaces.
	IsCommentBlockDeclarationContext()
}

type CommentBlockDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommentBlockDeclarationContext() *CommentBlockDeclarationContext {
	var p = new(CommentBlockDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DesignParserRULE_commentBlockDeclaration
	return p
}

func (*CommentBlockDeclarationContext) IsCommentBlockDeclarationContext() {}

func NewCommentBlockDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommentBlockDeclarationContext {
	var p = new(CommentBlockDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesignParserRULE_commentBlockDeclaration

	return p
}

func (s *CommentBlockDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *CommentBlockDeclarationContext) CommentDeclaration() ICommentDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommentDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommentDeclarationContext)
}

func (s *CommentBlockDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommentBlockDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommentBlockDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.EnterCommentBlockDeclaration(s)
	}
}

func (s *CommentBlockDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.ExitCommentBlockDeclaration(s)
	}
}

func (s *CommentBlockDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DesignVisitor:
		return t.VisitCommentBlockDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DesignParser) CommentBlockDeclaration() (localctx ICommentBlockDeclarationContext) {
	localctx = NewCommentBlockDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, DesignParserRULE_commentBlockDeclaration)
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

	if _la == DesignParserT__4 {
		{
			p.SetState(72)
			p.CommentDeclaration()
		}

	}

	return localctx
}

// IDesignBlockDeclarationContext is an interface to support dynamic dispatch.
type IDesignBlockDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDesignBlockDeclarationContext differentiates from other interfaces.
	IsDesignBlockDeclarationContext()
}

type DesignBlockDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDesignBlockDeclarationContext() *DesignBlockDeclarationContext {
	var p = new(DesignBlockDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DesignParserRULE_designBlockDeclaration
	return p
}

func (*DesignBlockDeclarationContext) IsDesignBlockDeclarationContext() {}

func NewDesignBlockDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DesignBlockDeclarationContext {
	var p = new(DesignBlockDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesignParserRULE_designBlockDeclaration

	return p
}

func (s *DesignBlockDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *DesignBlockDeclarationContext) DESIGN() antlr.TerminalNode {
	return s.GetToken(DesignParserDESIGN, 0)
}

func (s *DesignBlockDeclarationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(DesignParserIDENTIFIER, 0)
}

func (s *DesignBlockDeclarationContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(DesignParserLBRACE, 0)
}

func (s *DesignBlockDeclarationContext) DesignBodyDeclaration() IDesignBodyDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDesignBodyDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDesignBodyDeclarationContext)
}

func (s *DesignBlockDeclarationContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(DesignParserRBRACE, 0)
}

func (s *DesignBlockDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DesignBlockDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DesignBlockDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.EnterDesignBlockDeclaration(s)
	}
}

func (s *DesignBlockDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.ExitDesignBlockDeclaration(s)
	}
}

func (s *DesignBlockDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DesignVisitor:
		return t.VisitDesignBlockDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DesignParser) DesignBlockDeclaration() (localctx IDesignBlockDeclarationContext) {
	localctx = NewDesignBlockDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, DesignParserRULE_designBlockDeclaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.SetState(75)
		p.Match(DesignParserDESIGN)
	}
	{
		p.SetState(76)
		p.Match(DesignParserIDENTIFIER)
	}
	{
		p.SetState(77)
		p.Match(DesignParserLBRACE)
	}
	{
		p.SetState(78)
		p.DesignBodyDeclaration()
	}
	{
		p.SetState(79)
		p.Match(DesignParserRBRACE)
	}

	return localctx
}

// ITemplateBlockDeclarationContext is an interface to support dynamic dispatch.
type ITemplateBlockDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTemplateBlockDeclarationContext differentiates from other interfaces.
	IsTemplateBlockDeclarationContext()
}

type TemplateBlockDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTemplateBlockDeclarationContext() *TemplateBlockDeclarationContext {
	var p = new(TemplateBlockDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DesignParserRULE_templateBlockDeclaration
	return p
}

func (*TemplateBlockDeclarationContext) IsTemplateBlockDeclarationContext() {}

func NewTemplateBlockDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TemplateBlockDeclarationContext {
	var p = new(TemplateBlockDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesignParserRULE_templateBlockDeclaration

	return p
}

func (s *TemplateBlockDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *TemplateBlockDeclarationContext) TEMPLATE() antlr.TerminalNode {
	return s.GetToken(DesignParserTEMPLATE, 0)
}

func (s *TemplateBlockDeclarationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(DesignParserIDENTIFIER, 0)
}

func (s *TemplateBlockDeclarationContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(DesignParserLBRACE, 0)
}

func (s *TemplateBlockDeclarationContext) TemplateBodyDeclaration() ITemplateBodyDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITemplateBodyDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITemplateBodyDeclarationContext)
}

func (s *TemplateBlockDeclarationContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(DesignParserRBRACE, 0)
}

func (s *TemplateBlockDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TemplateBlockDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TemplateBlockDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.EnterTemplateBlockDeclaration(s)
	}
}

func (s *TemplateBlockDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.ExitTemplateBlockDeclaration(s)
	}
}

func (s *TemplateBlockDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DesignVisitor:
		return t.VisitTemplateBlockDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DesignParser) TemplateBlockDeclaration() (localctx ITemplateBlockDeclarationContext) {
	localctx = NewTemplateBlockDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, DesignParserRULE_templateBlockDeclaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.SetState(81)
		p.Match(DesignParserTEMPLATE)
	}
	{
		p.SetState(82)
		p.Match(DesignParserIDENTIFIER)
	}
	{
		p.SetState(83)
		p.Match(DesignParserLBRACE)
	}
	{
		p.SetState(84)
		p.TemplateBodyDeclaration()
	}
	{
		p.SetState(85)
		p.Match(DesignParserRBRACE)
	}

	return localctx
}

// IComponentBlockDeclarationContext is an interface to support dynamic dispatch.
type IComponentBlockDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsComponentBlockDeclarationContext differentiates from other interfaces.
	IsComponentBlockDeclarationContext()
}

type ComponentBlockDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComponentBlockDeclarationContext() *ComponentBlockDeclarationContext {
	var p = new(ComponentBlockDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DesignParserRULE_componentBlockDeclaration
	return p
}

func (*ComponentBlockDeclarationContext) IsComponentBlockDeclarationContext() {}

func NewComponentBlockDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComponentBlockDeclarationContext {
	var p = new(ComponentBlockDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesignParserRULE_componentBlockDeclaration

	return p
}

func (s *ComponentBlockDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *ComponentBlockDeclarationContext) COMPONENT() antlr.TerminalNode {
	return s.GetToken(DesignParserCOMPONENT, 0)
}

func (s *ComponentBlockDeclarationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(DesignParserIDENTIFIER, 0)
}

func (s *ComponentBlockDeclarationContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(DesignParserLBRACE, 0)
}

func (s *ComponentBlockDeclarationContext) ComponentBodyDeclaration() IComponentBodyDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComponentBodyDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComponentBodyDeclarationContext)
}

func (s *ComponentBlockDeclarationContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(DesignParserRBRACE, 0)
}

func (s *ComponentBlockDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComponentBlockDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComponentBlockDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.EnterComponentBlockDeclaration(s)
	}
}

func (s *ComponentBlockDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.ExitComponentBlockDeclaration(s)
	}
}

func (s *ComponentBlockDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DesignVisitor:
		return t.VisitComponentBlockDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DesignParser) ComponentBlockDeclaration() (localctx IComponentBlockDeclarationContext) {
	localctx = NewComponentBlockDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, DesignParserRULE_componentBlockDeclaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.Match(DesignParserCOMPONENT)
	}
	{
		p.SetState(88)
		p.Match(DesignParserIDENTIFIER)
	}
	{
		p.SetState(89)
		p.Match(DesignParserLBRACE)
	}
	{
		p.SetState(90)
		p.ComponentBodyDeclaration()
	}
	{
		p.SetState(91)
		p.Match(DesignParserRBRACE)
	}

	return localctx
}

// ILayerBlockDeclarationContext is an interface to support dynamic dispatch.
type ILayerBlockDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLayerBlockDeclarationContext differentiates from other interfaces.
	IsLayerBlockDeclarationContext()
}

type LayerBlockDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLayerBlockDeclarationContext() *LayerBlockDeclarationContext {
	var p = new(LayerBlockDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DesignParserRULE_layerBlockDeclaration
	return p
}

func (*LayerBlockDeclarationContext) IsLayerBlockDeclarationContext() {}

func NewLayerBlockDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LayerBlockDeclarationContext {
	var p = new(LayerBlockDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesignParserRULE_layerBlockDeclaration

	return p
}

func (s *LayerBlockDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *LayerBlockDeclarationContext) LAYER() antlr.TerminalNode {
	return s.GetToken(DesignParserLAYER, 0)
}

func (s *LayerBlockDeclarationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(DesignParserIDENTIFIER, 0)
}

func (s *LayerBlockDeclarationContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(DesignParserLBRACE, 0)
}

func (s *LayerBlockDeclarationContext) LayerBodyDeclaration() ILayerBodyDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILayerBodyDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILayerBodyDeclarationContext)
}

func (s *LayerBlockDeclarationContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(DesignParserRBRACE, 0)
}

func (s *LayerBlockDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LayerBlockDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LayerBlockDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.EnterLayerBlockDeclaration(s)
	}
}

func (s *LayerBlockDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.ExitLayerBlockDeclaration(s)
	}
}

func (s *LayerBlockDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DesignVisitor:
		return t.VisitLayerBlockDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DesignParser) LayerBlockDeclaration() (localctx ILayerBlockDeclarationContext) {
	localctx = NewLayerBlockDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, DesignParserRULE_layerBlockDeclaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.Match(DesignParserLAYER)
	}
	{
		p.SetState(94)
		p.Match(DesignParserIDENTIFIER)
	}
	{
		p.SetState(95)
		p.Match(DesignParserLBRACE)
	}
	{
		p.SetState(96)
		p.LayerBodyDeclaration()
	}
	{
		p.SetState(97)
		p.Match(DesignParserRBRACE)
	}

	return localctx
}

// ICodeBlockDeclarationContext is an interface to support dynamic dispatch.
type ICodeBlockDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCodeBlockDeclarationContext differentiates from other interfaces.
	IsCodeBlockDeclarationContext()
}

type CodeBlockDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCodeBlockDeclarationContext() *CodeBlockDeclarationContext {
	var p = new(CodeBlockDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DesignParserRULE_codeBlockDeclaration
	return p
}

func (*CodeBlockDeclarationContext) IsCodeBlockDeclarationContext() {}

func NewCodeBlockDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CodeBlockDeclarationContext {
	var p = new(CodeBlockDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesignParserRULE_codeBlockDeclaration

	return p
}

func (s *CodeBlockDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *CodeBlockDeclarationContext) AllHTML_STRING() []antlr.TerminalNode {
	return s.GetTokens(DesignParserHTML_STRING)
}

func (s *CodeBlockDeclarationContext) HTML_STRING(i int) antlr.TerminalNode {
	return s.GetToken(DesignParserHTML_STRING, i)
}

func (s *CodeBlockDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CodeBlockDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CodeBlockDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.EnterCodeBlockDeclaration(s)
	}
}

func (s *CodeBlockDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.ExitCodeBlockDeclaration(s)
	}
}

func (s *CodeBlockDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DesignVisitor:
		return t.VisitCodeBlockDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DesignParser) CodeBlockDeclaration() (localctx ICodeBlockDeclarationContext) {
	localctx = NewCodeBlockDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, DesignParserRULE_codeBlockDeclaration)
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
		p.Match(DesignParserT__1)
	}
	p.SetState(103)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DesignParserHTML_STRING {
		{
			p.SetState(100)
			p.Match(DesignParserHTML_STRING)
		}

		p.SetState(105)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(106)
		p.Match(DesignParserT__1)
	}

	return localctx
}

// IDesignBodyDeclarationContext is an interface to support dynamic dispatch.
type IDesignBodyDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDesignBodyDeclarationContext differentiates from other interfaces.
	IsDesignBodyDeclarationContext()
}

type DesignBodyDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDesignBodyDeclarationContext() *DesignBodyDeclarationContext {
	var p = new(DesignBodyDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DesignParserRULE_designBodyDeclaration
	return p
}

func (*DesignBodyDeclarationContext) IsDesignBodyDeclarationContext() {}

func NewDesignBodyDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DesignBodyDeclarationContext {
	var p = new(DesignBodyDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesignParserRULE_designBodyDeclaration

	return p
}

func (s *DesignBodyDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *DesignBodyDeclarationContext) AllExpress() []IExpressContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressContext)(nil)).Elem())
	var tst = make([]IExpressContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressContext)
		}
	}

	return tst
}

func (s *DesignBodyDeclarationContext) Express(i int) IExpressContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressContext)
}

func (s *DesignBodyDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DesignBodyDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DesignBodyDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.EnterDesignBodyDeclaration(s)
	}
}

func (s *DesignBodyDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.ExitDesignBodyDeclaration(s)
	}
}

func (s *DesignBodyDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DesignVisitor:
		return t.VisitDesignBodyDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DesignParser) DesignBodyDeclaration() (localctx IDesignBodyDeclarationContext) {
	localctx = NewDesignBodyDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, DesignParserRULE_designBodyDeclaration)
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
	p.SetState(111)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<DesignParserLAYER)|(1<<DesignParserTEMPLATE)|(1<<DesignParserIDENTIFIER))) != 0 {
		{
			p.SetState(108)
			p.Express()
		}

		p.SetState(113)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ITemplateBodyDeclarationContext is an interface to support dynamic dispatch.
type ITemplateBodyDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTemplateBodyDeclarationContext differentiates from other interfaces.
	IsTemplateBodyDeclarationContext()
}

type TemplateBodyDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTemplateBodyDeclarationContext() *TemplateBodyDeclarationContext {
	var p = new(TemplateBodyDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DesignParserRULE_templateBodyDeclaration
	return p
}

func (*TemplateBodyDeclarationContext) IsTemplateBodyDeclarationContext() {}

func NewTemplateBodyDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TemplateBodyDeclarationContext {
	var p = new(TemplateBodyDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesignParserRULE_templateBodyDeclaration

	return p
}

func (s *TemplateBodyDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *TemplateBodyDeclarationContext) AllExpress() []IExpressContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressContext)(nil)).Elem())
	var tst = make([]IExpressContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressContext)
		}
	}

	return tst
}

func (s *TemplateBodyDeclarationContext) Express(i int) IExpressContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressContext)
}

func (s *TemplateBodyDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TemplateBodyDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TemplateBodyDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.EnterTemplateBodyDeclaration(s)
	}
}

func (s *TemplateBodyDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.ExitTemplateBodyDeclaration(s)
	}
}

func (s *TemplateBodyDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DesignVisitor:
		return t.VisitTemplateBodyDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DesignParser) TemplateBodyDeclaration() (localctx ITemplateBodyDeclarationContext) {
	localctx = NewTemplateBodyDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, DesignParserRULE_templateBodyDeclaration)
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
	p.SetState(117)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<DesignParserLAYER)|(1<<DesignParserTEMPLATE)|(1<<DesignParserIDENTIFIER))) != 0 {
		{
			p.SetState(114)
			p.Express()
		}

		p.SetState(119)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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

func (s *ComponentBodyDeclarationContext) AllExpress() []IExpressContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressContext)(nil)).Elem())
	var tst = make([]IExpressContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressContext)
		}
	}

	return tst
}

func (s *ComponentBodyDeclarationContext) Express(i int) IExpressContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressContext)
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
	p.EnterRule(localctx, 22, DesignParserRULE_componentBodyDeclaration)
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
	p.SetState(123)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<DesignParserLAYER)|(1<<DesignParserTEMPLATE)|(1<<DesignParserIDENTIFIER))) != 0 {
		{
			p.SetState(120)
			p.Express()
		}

		p.SetState(125)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ILayerBodyDeclarationContext is an interface to support dynamic dispatch.
type ILayerBodyDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLayerBodyDeclarationContext differentiates from other interfaces.
	IsLayerBodyDeclarationContext()
}

type LayerBodyDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLayerBodyDeclarationContext() *LayerBodyDeclarationContext {
	var p = new(LayerBodyDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DesignParserRULE_layerBodyDeclaration
	return p
}

func (*LayerBodyDeclarationContext) IsLayerBodyDeclarationContext() {}

func NewLayerBodyDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LayerBodyDeclarationContext {
	var p = new(LayerBodyDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesignParserRULE_layerBodyDeclaration

	return p
}

func (s *LayerBodyDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *LayerBodyDeclarationContext) AllExpress() []IExpressContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressContext)(nil)).Elem())
	var tst = make([]IExpressContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressContext)
		}
	}

	return tst
}

func (s *LayerBodyDeclarationContext) Express(i int) IExpressContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressContext)
}

func (s *LayerBodyDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LayerBodyDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LayerBodyDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.EnterLayerBodyDeclaration(s)
	}
}

func (s *LayerBodyDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.ExitLayerBodyDeclaration(s)
	}
}

func (s *LayerBodyDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DesignVisitor:
		return t.VisitLayerBodyDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DesignParser) LayerBodyDeclaration() (localctx ILayerBodyDeclarationContext) {
	localctx = NewLayerBodyDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, DesignParserRULE_layerBodyDeclaration)
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
	p.SetState(129)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<DesignParserLAYER)|(1<<DesignParserTEMPLATE)|(1<<DesignParserIDENTIFIER))) != 0 {
		{
			p.SetState(126)
			p.Express()
		}

		p.SetState(131)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IExpressDeclarationContext is an interface to support dynamic dispatch.
type IExpressDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressDeclarationContext differentiates from other interfaces.
	IsExpressDeclarationContext()
}

type ExpressDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressDeclarationContext() *ExpressDeclarationContext {
	var p = new(ExpressDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DesignParserRULE_expressDeclaration
	return p
}

func (*ExpressDeclarationContext) IsExpressDeclarationContext() {}

func NewExpressDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressDeclarationContext {
	var p = new(ExpressDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesignParserRULE_expressDeclaration

	return p
}

func (s *ExpressDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressDeclarationContext) Express() IExpressContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressContext)
}

func (s *ExpressDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.EnterExpressDeclaration(s)
	}
}

func (s *ExpressDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.ExitExpressDeclaration(s)
	}
}

func (s *ExpressDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DesignVisitor:
		return t.VisitExpressDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DesignParser) ExpressDeclaration() (localctx IExpressDeclarationContext) {
	localctx = NewExpressDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, DesignParserRULE_expressDeclaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.SetState(132)
		p.Express()
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

func (s *ExpressContext) EqualExpress() IEqualExpressContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEqualExpressContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEqualExpressContext)
}

func (s *ExpressContext) UseExpress() IUseExpressContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUseExpressContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUseExpressContext)
}

func (s *ExpressContext) ValueExpress() IValueExpressContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueExpressContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueExpressContext)
}

func (s *ExpressContext) TemplateExpress() ITemplateExpressContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITemplateExpressContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITemplateExpressContext)
}

func (s *ExpressContext) LayerExpress() ILayerExpressContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILayerExpressContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILayerExpressContext)
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
	p.EnterRule(localctx, 28, DesignParserRULE_express)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(134)
			p.EqualExpress()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(135)
			p.UseExpress()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(136)
			p.ValueExpress()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(137)
			p.TemplateExpress()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(138)
			p.LayerExpress()
		}

	}

	return localctx
}

// IEqualExpressContext is an interface to support dynamic dispatch.
type IEqualExpressContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEqualExpressContext differentiates from other interfaces.
	IsEqualExpressContext()
}

type EqualExpressContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEqualExpressContext() *EqualExpressContext {
	var p = new(EqualExpressContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DesignParserRULE_equalExpress
	return p
}

func (*EqualExpressContext) IsEqualExpressContext() {}

func NewEqualExpressContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EqualExpressContext {
	var p = new(EqualExpressContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesignParserRULE_equalExpress

	return p
}

func (s *EqualExpressContext) GetParser() antlr.Parser { return s.parser }

func (s *EqualExpressContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(DesignParserIDENTIFIER)
}

func (s *EqualExpressContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(DesignParserIDENTIFIER, i)
}

func (s *EqualExpressContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualExpressContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EqualExpressContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.EnterEqualExpress(s)
	}
}

func (s *EqualExpressContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.ExitEqualExpress(s)
	}
}

func (s *EqualExpressContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DesignVisitor:
		return t.VisitEqualExpress(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DesignParser) EqualExpress() (localctx IEqualExpressContext) {
	localctx = NewEqualExpressContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, DesignParserRULE_equalExpress)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.Match(DesignParserIDENTIFIER)
	}
	{
		p.SetState(142)
		p.Match(DesignParserT__2)
	}
	{
		p.SetState(143)
		p.Match(DesignParserIDENTIFIER)
	}

	return localctx
}

// IUseExpressContext is an interface to support dynamic dispatch.
type IUseExpressContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUseExpressContext differentiates from other interfaces.
	IsUseExpressContext()
}

type UseExpressContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUseExpressContext() *UseExpressContext {
	var p = new(UseExpressContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DesignParserRULE_useExpress
	return p
}

func (*UseExpressContext) IsUseExpressContext() {}

func NewUseExpressContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UseExpressContext {
	var p = new(UseExpressContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesignParserRULE_useExpress

	return p
}

func (s *UseExpressContext) GetParser() antlr.Parser { return s.parser }

func (s *UseExpressContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(DesignParserIDENTIFIER)
}

func (s *UseExpressContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(DesignParserIDENTIFIER, i)
}

func (s *UseExpressContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UseExpressContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UseExpressContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.EnterUseExpress(s)
	}
}

func (s *UseExpressContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.ExitUseExpress(s)
	}
}

func (s *UseExpressContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DesignVisitor:
		return t.VisitUseExpress(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DesignParser) UseExpress() (localctx IUseExpressContext) {
	localctx = NewUseExpressContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, DesignParserRULE_useExpress)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.Match(DesignParserIDENTIFIER)
	}
	{
		p.SetState(146)
		p.Match(DesignParserT__3)
	}
	{
		p.SetState(147)
		p.Match(DesignParserIDENTIFIER)
	}

	return localctx
}

// IValueExpressContext is an interface to support dynamic dispatch.
type IValueExpressContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueExpressContext differentiates from other interfaces.
	IsValueExpressContext()
}

type ValueExpressContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueExpressContext() *ValueExpressContext {
	var p = new(ValueExpressContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DesignParserRULE_valueExpress
	return p
}

func (*ValueExpressContext) IsValueExpressContext() {}

func NewValueExpressContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueExpressContext {
	var p = new(ValueExpressContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesignParserRULE_valueExpress

	return p
}

func (s *ValueExpressContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueExpressContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(DesignParserIDENTIFIER)
}

func (s *ValueExpressContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(DesignParserIDENTIFIER, i)
}

func (s *ValueExpressContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueExpressContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueExpressContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.EnterValueExpress(s)
	}
}

func (s *ValueExpressContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.ExitValueExpress(s)
	}
}

func (s *ValueExpressContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DesignVisitor:
		return t.VisitValueExpress(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DesignParser) ValueExpress() (localctx IValueExpressContext) {
	localctx = NewValueExpressContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, DesignParserRULE_valueExpress)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.Match(DesignParserIDENTIFIER)
	}
	{
		p.SetState(150)
		p.Match(DesignParserT__0)
	}
	{
		p.SetState(151)
		p.Match(DesignParserIDENTIFIER)
	}

	return localctx
}

// ILayerExpressContext is an interface to support dynamic dispatch.
type ILayerExpressContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLayerExpressContext differentiates from other interfaces.
	IsLayerExpressContext()
}

type LayerExpressContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLayerExpressContext() *LayerExpressContext {
	var p = new(LayerExpressContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DesignParserRULE_layerExpress
	return p
}

func (*LayerExpressContext) IsLayerExpressContext() {}

func NewLayerExpressContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LayerExpressContext {
	var p = new(LayerExpressContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesignParserRULE_layerExpress

	return p
}

func (s *LayerExpressContext) GetParser() antlr.Parser { return s.parser }

func (s *LayerExpressContext) LAYER() antlr.TerminalNode {
	return s.GetToken(DesignParserLAYER, 0)
}

func (s *LayerExpressContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(DesignParserIDENTIFIER, 0)
}

func (s *LayerExpressContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LayerExpressContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LayerExpressContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.EnterLayerExpress(s)
	}
}

func (s *LayerExpressContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.ExitLayerExpress(s)
	}
}

func (s *LayerExpressContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DesignVisitor:
		return t.VisitLayerExpress(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DesignParser) LayerExpress() (localctx ILayerExpressContext) {
	localctx = NewLayerExpressContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, DesignParserRULE_layerExpress)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.Match(DesignParserLAYER)
	}
	{
		p.SetState(154)
		p.Match(DesignParserIDENTIFIER)
	}

	return localctx
}

// ITemplateExpressContext is an interface to support dynamic dispatch.
type ITemplateExpressContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTemplateExpressContext differentiates from other interfaces.
	IsTemplateExpressContext()
}

type TemplateExpressContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTemplateExpressContext() *TemplateExpressContext {
	var p = new(TemplateExpressContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DesignParserRULE_templateExpress
	return p
}

func (*TemplateExpressContext) IsTemplateExpressContext() {}

func NewTemplateExpressContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TemplateExpressContext {
	var p = new(TemplateExpressContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesignParserRULE_templateExpress

	return p
}

func (s *TemplateExpressContext) GetParser() antlr.Parser { return s.parser }

func (s *TemplateExpressContext) TEMPLATE() antlr.TerminalNode {
	return s.GetToken(DesignParserTEMPLATE, 0)
}

func (s *TemplateExpressContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(DesignParserIDENTIFIER, 0)
}

func (s *TemplateExpressContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TemplateExpressContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TemplateExpressContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.EnterTemplateExpress(s)
	}
}

func (s *TemplateExpressContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.ExitTemplateExpress(s)
	}
}

func (s *TemplateExpressContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DesignVisitor:
		return t.VisitTemplateExpress(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DesignParser) TemplateExpress() (localctx ITemplateExpressContext) {
	localctx = NewTemplateExpressContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, DesignParserRULE_templateExpress)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.Match(DesignParserTEMPLATE)
	}
	{
		p.SetState(157)
		p.Match(DesignParserIDENTIFIER)
	}

	return localctx
}

// ILayerContext is an interface to support dynamic dispatch.
type ILayerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLayerContext differentiates from other interfaces.
	IsLayerContext()
}

type LayerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLayerContext() *LayerContext {
	var p = new(LayerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DesignParserRULE_layer
	return p
}

func (*LayerContext) IsLayerContext() {}

func NewLayerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LayerContext {
	var p = new(LayerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesignParserRULE_layer

	return p
}

func (s *LayerContext) GetParser() antlr.Parser { return s.parser }

func (s *LayerContext) LAYER() antlr.TerminalNode {
	return s.GetToken(DesignParserLAYER, 0)
}

func (s *LayerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LayerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LayerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.EnterLayer(s)
	}
}

func (s *LayerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.ExitLayer(s)
	}
}

func (s *LayerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DesignVisitor:
		return t.VisitLayer(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DesignParser) Layer() (localctx ILayerContext) {
	localctx = NewLayerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, DesignParserRULE_layer)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.Match(DesignParserLAYER)
	}

	return localctx
}

// ICommentDeclarationContext is an interface to support dynamic dispatch.
type ICommentDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCommentDeclarationContext differentiates from other interfaces.
	IsCommentDeclarationContext()
}

type CommentDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommentDeclarationContext() *CommentDeclarationContext {
	var p = new(CommentDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DesignParserRULE_commentDeclaration
	return p
}

func (*CommentDeclarationContext) IsCommentDeclarationContext() {}

func NewCommentDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommentDeclarationContext {
	var p = new(CommentDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesignParserRULE_commentDeclaration

	return p
}

func (s *CommentDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *CommentDeclarationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(DesignParserIDENTIFIER, 0)
}

func (s *CommentDeclarationContext) Space() antlr.TerminalNode {
	return s.GetToken(DesignParserSpace, 0)
}

func (s *CommentDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommentDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommentDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.EnterCommentDeclaration(s)
	}
}

func (s *CommentDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesignListener); ok {
		listenerT.ExitCommentDeclaration(s)
	}
}

func (s *CommentDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DesignVisitor:
		return t.VisitCommentDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DesignParser) CommentDeclaration() (localctx ICommentDeclarationContext) {
	localctx = NewCommentDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, DesignParserRULE_commentDeclaration)
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
		p.SetState(161)
		p.Match(DesignParserT__4)
	}
	p.SetState(163)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DesignParserSpace {
		{
			p.SetState(162)
			p.Match(DesignParserSpace)
		}

	}
	{
		p.SetState(165)
		p.Match(DesignParserIDENTIFIER)
	}

	return localctx
}
