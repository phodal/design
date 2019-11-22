// Code generated from Design.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Design

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseDesignListener is a complete listener for a parse tree produced by DesignParser.
type BaseDesignListener struct{}

var _ DesignListener = &BaseDesignListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDesignListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDesignListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDesignListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDesignListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterDesignIt is called when production designIt is entered.
func (s *BaseDesignListener) EnterDesignIt(ctx *DesignItContext) {}

// ExitDesignIt is called when production designIt is exited.
func (s *BaseDesignListener) ExitDesignIt(ctx *DesignItContext) {}

// EnterDeclaration is called when production declaration is entered.
func (s *BaseDesignListener) EnterDeclaration(ctx *DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *BaseDesignListener) ExitDeclaration(ctx *DeclarationContext) {}

// EnterDesignSystemDeclaration is called when production designSystemDeclaration is entered.
func (s *BaseDesignListener) EnterDesignSystemDeclaration(ctx *DesignSystemDeclarationContext) {}

// ExitDesignSystemDeclaration is called when production designSystemDeclaration is exited.
func (s *BaseDesignListener) ExitDesignSystemDeclaration(ctx *DesignSystemDeclarationContext) {}

// EnterCommentBlockDeclaration is called when production commentBlockDeclaration is entered.
func (s *BaseDesignListener) EnterCommentBlockDeclaration(ctx *CommentBlockDeclarationContext) {}

// ExitCommentBlockDeclaration is called when production commentBlockDeclaration is exited.
func (s *BaseDesignListener) ExitCommentBlockDeclaration(ctx *CommentBlockDeclarationContext) {}

// EnterDesignBlockDeclaration is called when production designBlockDeclaration is entered.
func (s *BaseDesignListener) EnterDesignBlockDeclaration(ctx *DesignBlockDeclarationContext) {}

// ExitDesignBlockDeclaration is called when production designBlockDeclaration is exited.
func (s *BaseDesignListener) ExitDesignBlockDeclaration(ctx *DesignBlockDeclarationContext) {}

// EnterDesignBodyDeclaration is called when production designBodyDeclaration is entered.
func (s *BaseDesignListener) EnterDesignBodyDeclaration(ctx *DesignBodyDeclarationContext) {}

// ExitDesignBodyDeclaration is called when production designBodyDeclaration is exited.
func (s *BaseDesignListener) ExitDesignBodyDeclaration(ctx *DesignBodyDeclarationContext) {}

// EnterExpress is called when production express is entered.
func (s *BaseDesignListener) EnterExpress(ctx *ExpressContext) {}

// ExitExpress is called when production express is exited.
func (s *BaseDesignListener) ExitExpress(ctx *ExpressContext) {}

// EnterEqualExpress is called when production equalExpress is entered.
func (s *BaseDesignListener) EnterEqualExpress(ctx *EqualExpressContext) {}

// ExitEqualExpress is called when production equalExpress is exited.
func (s *BaseDesignListener) ExitEqualExpress(ctx *EqualExpressContext) {}

// EnterUseExpress is called when production useExpress is entered.
func (s *BaseDesignListener) EnterUseExpress(ctx *UseExpressContext) {}

// ExitUseExpress is called when production useExpress is exited.
func (s *BaseDesignListener) ExitUseExpress(ctx *UseExpressContext) {}

// EnterValueExpress is called when production valueExpress is entered.
func (s *BaseDesignListener) EnterValueExpress(ctx *ValueExpressContext) {}

// ExitValueExpress is called when production valueExpress is exited.
func (s *BaseDesignListener) ExitValueExpress(ctx *ValueExpressContext) {}

// EnterTemplateExpress is called when production templateExpress is entered.
func (s *BaseDesignListener) EnterTemplateExpress(ctx *TemplateExpressContext) {}

// ExitTemplateExpress is called when production templateExpress is exited.
func (s *BaseDesignListener) ExitTemplateExpress(ctx *TemplateExpressContext) {}

// EnterLayer is called when production layer is entered.
func (s *BaseDesignListener) EnterLayer(ctx *LayerContext) {}

// ExitLayer is called when production layer is exited.
func (s *BaseDesignListener) ExitLayer(ctx *LayerContext) {}

// EnterComment is called when production comment is entered.
func (s *BaseDesignListener) EnterComment(ctx *CommentContext) {}

// ExitComment is called when production comment is exited.
func (s *BaseDesignListener) ExitComment(ctx *CommentContext) {}
