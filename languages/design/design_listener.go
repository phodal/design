// Code generated from Design.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Design

import "github.com/antlr/antlr4/runtime/Go/antlr"

// DesignListener is a complete listener for a parse tree produced by DesignParser.
type DesignListener interface {
	antlr.ParseTreeListener

	// EnterDesignIt is called when entering the designIt production.
	EnterDesignIt(c *DesignItContext)

	// EnterDeclaration is called when entering the declaration production.
	EnterDeclaration(c *DeclarationContext)

	// EnterDesignSystemDeclaration is called when entering the designSystemDeclaration production.
	EnterDesignSystemDeclaration(c *DesignSystemDeclarationContext)

	// EnterCommentBlockDeclaration is called when entering the commentBlockDeclaration production.
	EnterCommentBlockDeclaration(c *CommentBlockDeclarationContext)

	// EnterDesignBlockDeclaration is called when entering the designBlockDeclaration production.
	EnterDesignBlockDeclaration(c *DesignBlockDeclarationContext)

	// EnterDesignBodyDeclaration is called when entering the designBodyDeclaration production.
	EnterDesignBodyDeclaration(c *DesignBodyDeclarationContext)

	// EnterExpress is called when entering the express production.
	EnterExpress(c *ExpressContext)

	// EnterEqualExpress is called when entering the equalExpress production.
	EnterEqualExpress(c *EqualExpressContext)

	// EnterUseExpress is called when entering the useExpress production.
	EnterUseExpress(c *UseExpressContext)

	// EnterValueExpress is called when entering the valueExpress production.
	EnterValueExpress(c *ValueExpressContext)

	// EnterTemplateExpress is called when entering the templateExpress production.
	EnterTemplateExpress(c *TemplateExpressContext)

	// EnterLayer is called when entering the layer production.
	EnterLayer(c *LayerContext)

	// EnterComment is called when entering the comment production.
	EnterComment(c *CommentContext)

	// ExitDesignIt is called when exiting the designIt production.
	ExitDesignIt(c *DesignItContext)

	// ExitDeclaration is called when exiting the declaration production.
	ExitDeclaration(c *DeclarationContext)

	// ExitDesignSystemDeclaration is called when exiting the designSystemDeclaration production.
	ExitDesignSystemDeclaration(c *DesignSystemDeclarationContext)

	// ExitCommentBlockDeclaration is called when exiting the commentBlockDeclaration production.
	ExitCommentBlockDeclaration(c *CommentBlockDeclarationContext)

	// ExitDesignBlockDeclaration is called when exiting the designBlockDeclaration production.
	ExitDesignBlockDeclaration(c *DesignBlockDeclarationContext)

	// ExitDesignBodyDeclaration is called when exiting the designBodyDeclaration production.
	ExitDesignBodyDeclaration(c *DesignBodyDeclarationContext)

	// ExitExpress is called when exiting the express production.
	ExitExpress(c *ExpressContext)

	// ExitEqualExpress is called when exiting the equalExpress production.
	ExitEqualExpress(c *EqualExpressContext)

	// ExitUseExpress is called when exiting the useExpress production.
	ExitUseExpress(c *UseExpressContext)

	// ExitValueExpress is called when exiting the valueExpress production.
	ExitValueExpress(c *ValueExpressContext)

	// ExitTemplateExpress is called when exiting the templateExpress production.
	ExitTemplateExpress(c *TemplateExpressContext)

	// ExitLayer is called when exiting the layer production.
	ExitLayer(c *LayerContext)

	// ExitComment is called when exiting the comment production.
	ExitComment(c *CommentContext)
}
