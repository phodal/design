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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 29, 107,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 3, 2, 3, 2, 3, 2, 3, 2, 7, 2, 33, 10, 2, 12, 2, 14,
	2, 36, 11, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 44, 10, 3, 3, 4,
	3, 4, 3, 5, 7, 5, 49, 10, 5, 12, 5, 14, 5, 52, 11, 5, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 7, 7, 7, 61, 10, 7, 12, 7, 14, 7, 64, 11, 7, 3, 8,
	3, 8, 3, 8, 3, 8, 5, 8, 70, 10, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 7, 14, 94, 10, 14, 12, 14, 14, 14,
	97, 11, 14, 3, 14, 7, 14, 100, 10, 14, 12, 14, 14, 14, 103, 11, 14, 3,
	14, 3, 14, 3, 14, 3, 101, 2, 15, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22,
	24, 26, 2, 3, 3, 3, 25, 25, 2, 105, 2, 34, 3, 2, 2, 2, 4, 43, 3, 2, 2,
	2, 6, 45, 3, 2, 2, 2, 8, 50, 3, 2, 2, 2, 10, 53, 3, 2, 2, 2, 12, 62, 3,
	2, 2, 2, 14, 69, 3, 2, 2, 2, 16, 71, 3, 2, 2, 2, 18, 75, 3, 2, 2, 2, 20,
	79, 3, 2, 2, 2, 22, 83, 3, 2, 2, 2, 24, 89, 3, 2, 2, 2, 26, 91, 3, 2, 2,
	2, 28, 33, 5, 26, 14, 2, 29, 30, 5, 26, 14, 2, 30, 31, 5, 4, 3, 2, 31,
	33, 3, 2, 2, 2, 32, 28, 3, 2, 2, 2, 32, 29, 3, 2, 2, 2, 33, 36, 3, 2, 2,
	2, 34, 32, 3, 2, 2, 2, 34, 35, 3, 2, 2, 2, 35, 37, 3, 2, 2, 2, 36, 34,
	3, 2, 2, 2, 37, 38, 9, 2, 2, 2, 38, 3, 3, 2, 2, 2, 39, 44, 3, 2, 2, 2,
	40, 44, 5, 6, 4, 2, 41, 44, 5, 8, 5, 2, 42, 44, 5, 10, 6, 2, 43, 39, 3,
	2, 2, 2, 43, 40, 3, 2, 2, 2, 43, 41, 3, 2, 2, 2, 43, 42, 3, 2, 2, 2, 44,
	5, 3, 2, 2, 2, 45, 46, 7, 8, 2, 2, 46, 7, 3, 2, 2, 2, 47, 49, 5, 26, 14,
	2, 48, 47, 3, 2, 2, 2, 49, 52, 3, 2, 2, 2, 50, 48, 3, 2, 2, 2, 50, 51,
	3, 2, 2, 2, 51, 9, 3, 2, 2, 2, 52, 50, 3, 2, 2, 2, 53, 54, 7, 9, 2, 2,
	54, 55, 7, 28, 2, 2, 55, 56, 7, 26, 2, 2, 56, 57, 5, 12, 7, 2, 57, 58,
	7, 27, 2, 2, 58, 11, 3, 2, 2, 2, 59, 61, 5, 14, 8, 2, 60, 59, 3, 2, 2,
	2, 61, 64, 3, 2, 2, 2, 62, 60, 3, 2, 2, 2, 62, 63, 3, 2, 2, 2, 63, 13,
	3, 2, 2, 2, 64, 62, 3, 2, 2, 2, 65, 70, 5, 16, 9, 2, 66, 70, 5, 18, 10,
	2, 67, 70, 5, 20, 11, 2, 68, 70, 5, 22, 12, 2, 69, 65, 3, 2, 2, 2, 69,
	66, 3, 2, 2, 2, 69, 67, 3, 2, 2, 2, 69, 68, 3, 2, 2, 2, 70, 15, 3, 2, 2,
	2, 71, 72, 7, 28, 2, 2, 72, 73, 7, 3, 2, 2, 73, 74, 7, 28, 2, 2, 74, 17,
	3, 2, 2, 2, 75, 76, 7, 28, 2, 2, 76, 77, 7, 4, 2, 2, 77, 78, 7, 28, 2,
	2, 78, 19, 3, 2, 2, 2, 79, 80, 7, 28, 2, 2, 80, 81, 7, 5, 2, 2, 81, 82,
	7, 28, 2, 2, 82, 21, 3, 2, 2, 2, 83, 84, 7, 16, 2, 2, 84, 85, 7, 28, 2,
	2, 85, 86, 7, 6, 2, 2, 86, 87, 7, 29, 2, 2, 87, 88, 7, 6, 2, 2, 88, 23,
	3, 2, 2, 2, 89, 90, 7, 12, 2, 2, 90, 25, 3, 2, 2, 2, 91, 95, 7, 7, 2, 2,
	92, 94, 7, 24, 2, 2, 93, 92, 3, 2, 2, 2, 94, 97, 3, 2, 2, 2, 95, 93, 3,
	2, 2, 2, 95, 96, 3, 2, 2, 2, 96, 101, 3, 2, 2, 2, 97, 95, 3, 2, 2, 2, 98,
	100, 11, 2, 2, 2, 99, 98, 3, 2, 2, 2, 100, 103, 3, 2, 2, 2, 101, 102, 3,
	2, 2, 2, 101, 99, 3, 2, 2, 2, 102, 104, 3, 2, 2, 2, 103, 101, 3, 2, 2,
	2, 104, 105, 7, 25, 2, 2, 105, 27, 3, 2, 2, 2, 10, 32, 34, 43, 50, 62,
	69, 95, 101,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'='", "'->'", "':'", "'```'", "'#'", "'DesignSystem'", "'design'",
	"'project'", "'page'", "'layer'", "'function'", "'library'", "'unit'",
	"'template'", "'position'", "'Flow'", "'behavior'", "", "", "", "", "",
	"", "'{'", "'}'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "DesignSystem", "Design", "Project", "Page", "Layer",
	"Function", "Library", "Unit", "Template", "Position", "Flow", "Behavior",
	"WS", "COMMENT", "LINE_COMMENT", "EmptyLine", "Space", "NewLine", "LBRACE",
	"RBRACE", "IDENTIFIER", "Code",
}

var ruleNames = []string{
	"designIt", "declaration", "designSystemDeclaration", "commentBlockDeclaration",
	"designBlockDeclaration", "designBodyDeclaration", "express", "equalExpress",
	"useExpress", "valueExpress", "templateExpress", "layer", "comment",
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
	DesignParserEOF          = antlr.TokenEOF
	DesignParserT__0         = 1
	DesignParserT__1         = 2
	DesignParserT__2         = 3
	DesignParserT__3         = 4
	DesignParserT__4         = 5
	DesignParserDesignSystem = 6
	DesignParserDesign       = 7
	DesignParserProject      = 8
	DesignParserPage         = 9
	DesignParserLayer        = 10
	DesignParserFunction     = 11
	DesignParserLibrary      = 12
	DesignParserUnit         = 13
	DesignParserTemplate     = 14
	DesignParserPosition     = 15
	DesignParserFlow         = 16
	DesignParserBehavior     = 17
	DesignParserWS           = 18
	DesignParserCOMMENT      = 19
	DesignParserLINE_COMMENT = 20
	DesignParserEmptyLine    = 21
	DesignParserSpace        = 22
	DesignParserNewLine      = 23
	DesignParserLBRACE       = 24
	DesignParserRBRACE       = 25
	DesignParserIDENTIFIER   = 26
	DesignParserCode         = 27
)

// DesignParser rules.
const (
	DesignParserRULE_designIt                = 0
	DesignParserRULE_declaration             = 1
	DesignParserRULE_designSystemDeclaration = 2
	DesignParserRULE_commentBlockDeclaration = 3
	DesignParserRULE_designBlockDeclaration  = 4
	DesignParserRULE_designBodyDeclaration   = 5
	DesignParserRULE_express                 = 6
	DesignParserRULE_equalExpress            = 7
	DesignParserRULE_useExpress              = 8
	DesignParserRULE_valueExpress            = 9
	DesignParserRULE_templateExpress         = 10
	DesignParserRULE_layer                   = 11
	DesignParserRULE_comment                 = 12
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

func (s *DesignItContext) NewLine() antlr.TerminalNode {
	return s.GetToken(DesignParserNewLine, 0)
}

func (s *DesignItContext) EOF() antlr.TerminalNode {
	return s.GetToken(DesignParserEOF, 0)
}

func (s *DesignItContext) AllComment() []ICommentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICommentContext)(nil)).Elem())
	var tst = make([]ICommentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICommentContext)
		}
	}

	return tst
}

func (s *DesignItContext) Comment(i int) ICommentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICommentContext)
}

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
	p.SetState(32)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DesignParserT__4 {
		p.SetState(30)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(26)
				p.Comment()
			}

		case 2:
			{
				p.SetState(27)
				p.Comment()
			}
			{
				p.SetState(28)
				p.Declaration()
			}

		}

		p.SetState(34)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(35)
		_la = p.GetTokenStream().LA(1)

		if !(_la == DesignParserEOF || _la == DesignParserNewLine) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
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

func (s *DeclarationContext) CommentBlockDeclaration() ICommentBlockDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommentBlockDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommentBlockDeclarationContext)
}

func (s *DeclarationContext) DesignBlockDeclaration() IDesignBlockDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDesignBlockDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDesignBlockDeclarationContext)
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

	p.SetState(41)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(38)
			p.DesignSystemDeclaration()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(39)
			p.CommentBlockDeclaration()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(40)
			p.DesignBlockDeclaration()
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

func (s *DesignSystemDeclarationContext) DesignSystem() antlr.TerminalNode {
	return s.GetToken(DesignParserDesignSystem, 0)
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
		p.SetState(43)
		p.Match(DesignParserDesignSystem)
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

func (s *CommentBlockDeclarationContext) AllComment() []ICommentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICommentContext)(nil)).Elem())
	var tst = make([]ICommentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICommentContext)
		}
	}

	return tst
}

func (s *CommentBlockDeclarationContext) Comment(i int) ICommentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICommentContext)
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

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
	p.SetState(48)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(45)
				p.Comment()
			}

		}
		p.SetState(50)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
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

func (s *DesignBlockDeclarationContext) Design() antlr.TerminalNode {
	return s.GetToken(DesignParserDesign, 0)
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
		p.SetState(51)
		p.Match(DesignParserDesign)
	}
	{
		p.SetState(52)
		p.Match(DesignParserIDENTIFIER)
	}
	{
		p.SetState(53)
		p.Match(DesignParserLBRACE)
	}
	{
		p.SetState(54)
		p.DesignBodyDeclaration()
	}
	{
		p.SetState(55)
		p.Match(DesignParserRBRACE)
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
	p.EnterRule(localctx, 10, DesignParserRULE_designBodyDeclaration)
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
	p.SetState(60)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DesignParserTemplate || _la == DesignParserIDENTIFIER {
		{
			p.SetState(57)
			p.Express()
		}

		p.SetState(62)
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
	p.EnterRule(localctx, 12, DesignParserRULE_express)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(67)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(63)
			p.EqualExpress()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(64)
			p.UseExpress()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(65)
			p.ValueExpress()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(66)
			p.TemplateExpress()
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
	p.EnterRule(localctx, 14, DesignParserRULE_equalExpress)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.SetState(69)
		p.Match(DesignParserIDENTIFIER)
	}
	{
		p.SetState(70)
		p.Match(DesignParserT__0)
	}
	{
		p.SetState(71)
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
	p.EnterRule(localctx, 16, DesignParserRULE_useExpress)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.SetState(73)
		p.Match(DesignParserIDENTIFIER)
	}
	{
		p.SetState(74)
		p.Match(DesignParserT__1)
	}
	{
		p.SetState(75)
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
	p.EnterRule(localctx, 18, DesignParserRULE_valueExpress)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.SetState(77)
		p.Match(DesignParserIDENTIFIER)
	}
	{
		p.SetState(78)
		p.Match(DesignParserT__2)
	}
	{
		p.SetState(79)
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

func (s *TemplateExpressContext) Template() antlr.TerminalNode {
	return s.GetToken(DesignParserTemplate, 0)
}

func (s *TemplateExpressContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(DesignParserIDENTIFIER, 0)
}

func (s *TemplateExpressContext) Code() antlr.TerminalNode {
	return s.GetToken(DesignParserCode, 0)
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
	p.EnterRule(localctx, 20, DesignParserRULE_templateExpress)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.Match(DesignParserTemplate)
	}
	{
		p.SetState(82)
		p.Match(DesignParserIDENTIFIER)
	}
	{
		p.SetState(83)
		p.Match(DesignParserT__3)
	}
	{
		p.SetState(84)
		p.Match(DesignParserCode)
	}
	{
		p.SetState(85)
		p.Match(DesignParserT__3)
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

func (s *LayerContext) Layer() antlr.TerminalNode {
	return s.GetToken(DesignParserLayer, 0)
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
	p.EnterRule(localctx, 22, DesignParserRULE_layer)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.Match(DesignParserLayer)
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

func (s *CommentContext) NewLine() antlr.TerminalNode {
	return s.GetToken(DesignParserNewLine, 0)
}

func (s *CommentContext) AllSpace() []antlr.TerminalNode {
	return s.GetTokens(DesignParserSpace)
}

func (s *CommentContext) Space(i int) antlr.TerminalNode {
	return s.GetToken(DesignParserSpace, i)
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
	p.EnterRule(localctx, 24, DesignParserRULE_comment)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.SetState(89)
		p.Match(DesignParserT__4)
	}
	p.SetState(93)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(90)
				p.Match(DesignParserSpace)
			}

		}
		p.SetState(95)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
	}
	p.SetState(99)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())

	for _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1+1 {
			p.SetState(96)
			p.MatchWildcard()

		}
		p.SetState(101)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())
	}
	{
		p.SetState(102)
		p.Match(DesignParserNewLine)
	}

	return localctx
}
