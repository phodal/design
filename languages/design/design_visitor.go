// Code generated from Design.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Design

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by DesignParser.
type DesignVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by DesignParser#designIt.
	VisitDesignIt(ctx *DesignItContext) interface{}

	// Visit a parse tree produced by DesignParser#declaration.
	VisitDeclaration(ctx *DeclarationContext) interface{}

	// Visit a parse tree produced by DesignParser#designSystemDeclaration.
	VisitDesignSystemDeclaration(ctx *DesignSystemDeclarationContext) interface{}

	// Visit a parse tree produced by DesignParser#commentBlockDeclaration.
	VisitCommentBlockDeclaration(ctx *CommentBlockDeclarationContext) interface{}

	// Visit a parse tree produced by DesignParser#designBlockDeclaration.
	VisitDesignBlockDeclaration(ctx *DesignBlockDeclarationContext) interface{}

	// Visit a parse tree produced by DesignParser#designBodyDeclaration.
	VisitDesignBodyDeclaration(ctx *DesignBodyDeclarationContext) interface{}

	// Visit a parse tree produced by DesignParser#express.
	VisitExpress(ctx *ExpressContext) interface{}

	// Visit a parse tree produced by DesignParser#equalExpress.
	VisitEqualExpress(ctx *EqualExpressContext) interface{}

	// Visit a parse tree produced by DesignParser#useExpress.
	VisitUseExpress(ctx *UseExpressContext) interface{}

	// Visit a parse tree produced by DesignParser#valueExpress.
	VisitValueExpress(ctx *ValueExpressContext) interface{}

	// Visit a parse tree produced by DesignParser#templateExpress.
	VisitTemplateExpress(ctx *TemplateExpressContext) interface{}

	// Visit a parse tree produced by DesignParser#layer.
	VisitLayer(ctx *LayerContext) interface{}

	// Visit a parse tree produced by DesignParser#comment.
	VisitComment(ctx *CommentContext) interface{}
}
