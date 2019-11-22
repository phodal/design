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

// EnterTemplateBlockDeclaration is called when production templateBlockDeclaration is entered.
func (s *BaseDesignListener) EnterTemplateBlockDeclaration(ctx *TemplateBlockDeclarationContext) {}

// ExitTemplateBlockDeclaration is called when production templateBlockDeclaration is exited.
func (s *BaseDesignListener) ExitTemplateBlockDeclaration(ctx *TemplateBlockDeclarationContext) {}

// EnterComponentBlockDeclaration is called when production componentBlockDeclaration is entered.
func (s *BaseDesignListener) EnterComponentBlockDeclaration(ctx *ComponentBlockDeclarationContext) {}

// ExitComponentBlockDeclaration is called when production componentBlockDeclaration is exited.
func (s *BaseDesignListener) ExitComponentBlockDeclaration(ctx *ComponentBlockDeclarationContext) {}

// EnterLayerBlockDeclaration is called when production layerBlockDeclaration is entered.
func (s *BaseDesignListener) EnterLayerBlockDeclaration(ctx *LayerBlockDeclarationContext) {}

// ExitLayerBlockDeclaration is called when production layerBlockDeclaration is exited.
func (s *BaseDesignListener) ExitLayerBlockDeclaration(ctx *LayerBlockDeclarationContext) {}

// EnterCodeBlockDeclaration is called when production codeBlockDeclaration is entered.
func (s *BaseDesignListener) EnterCodeBlockDeclaration(ctx *CodeBlockDeclarationContext) {}

// ExitCodeBlockDeclaration is called when production codeBlockDeclaration is exited.
func (s *BaseDesignListener) ExitCodeBlockDeclaration(ctx *CodeBlockDeclarationContext) {}

// EnterDesignBodyDeclaration is called when production designBodyDeclaration is entered.
func (s *BaseDesignListener) EnterDesignBodyDeclaration(ctx *DesignBodyDeclarationContext) {}

// ExitDesignBodyDeclaration is called when production designBodyDeclaration is exited.
func (s *BaseDesignListener) ExitDesignBodyDeclaration(ctx *DesignBodyDeclarationContext) {}

// EnterTemplateBodyDeclaration is called when production templateBodyDeclaration is entered.
func (s *BaseDesignListener) EnterTemplateBodyDeclaration(ctx *TemplateBodyDeclarationContext) {}

// ExitTemplateBodyDeclaration is called when production templateBodyDeclaration is exited.
func (s *BaseDesignListener) ExitTemplateBodyDeclaration(ctx *TemplateBodyDeclarationContext) {}

// EnterComponentBodyDeclaration is called when production componentBodyDeclaration is entered.
func (s *BaseDesignListener) EnterComponentBodyDeclaration(ctx *ComponentBodyDeclarationContext) {}

// ExitComponentBodyDeclaration is called when production componentBodyDeclaration is exited.
func (s *BaseDesignListener) ExitComponentBodyDeclaration(ctx *ComponentBodyDeclarationContext) {}

// EnterLayerBodyDeclaration is called when production layerBodyDeclaration is entered.
func (s *BaseDesignListener) EnterLayerBodyDeclaration(ctx *LayerBodyDeclarationContext) {}

// ExitLayerBodyDeclaration is called when production layerBodyDeclaration is exited.
func (s *BaseDesignListener) ExitLayerBodyDeclaration(ctx *LayerBodyDeclarationContext) {}

// EnterExpressDeclaration is called when production expressDeclaration is entered.
func (s *BaseDesignListener) EnterExpressDeclaration(ctx *ExpressDeclarationContext) {}

// ExitExpressDeclaration is called when production expressDeclaration is exited.
func (s *BaseDesignListener) ExitExpressDeclaration(ctx *ExpressDeclarationContext) {}

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

// EnterLayerExpress is called when production layerExpress is entered.
func (s *BaseDesignListener) EnterLayerExpress(ctx *LayerExpressContext) {}

// ExitLayerExpress is called when production layerExpress is exited.
func (s *BaseDesignListener) ExitLayerExpress(ctx *LayerExpressContext) {}

// EnterTemplateExpress is called when production templateExpress is entered.
func (s *BaseDesignListener) EnterTemplateExpress(ctx *TemplateExpressContext) {}

// ExitTemplateExpress is called when production templateExpress is exited.
func (s *BaseDesignListener) ExitTemplateExpress(ctx *TemplateExpressContext) {}

// EnterExpressKey is called when production expressKey is entered.
func (s *BaseDesignListener) EnterExpressKey(ctx *ExpressKeyContext) {}

// ExitExpressKey is called when production expressKey is exited.
func (s *BaseDesignListener) ExitExpressKey(ctx *ExpressKeyContext) {}

// EnterExpressValue is called when production expressValue is entered.
func (s *BaseDesignListener) EnterExpressValue(ctx *ExpressValueContext) {}

// ExitExpressValue is called when production expressValue is exited.
func (s *BaseDesignListener) ExitExpressValue(ctx *ExpressValueContext) {}

// EnterLayer is called when production layer is entered.
func (s *BaseDesignListener) EnterLayer(ctx *LayerContext) {}

// ExitLayer is called when production layer is exited.
func (s *BaseDesignListener) ExitLayer(ctx *LayerContext) {}

// EnterCommentDeclaration is called when production commentDeclaration is entered.
func (s *BaseDesignListener) EnterCommentDeclaration(ctx *CommentDeclarationContext) {}

// ExitCommentDeclaration is called when production commentDeclaration is exited.
func (s *BaseDesignListener) ExitCommentDeclaration(ctx *CommentDeclarationContext) {}
