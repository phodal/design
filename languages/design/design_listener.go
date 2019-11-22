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

	// EnterTemplateBlockDeclaration is called when entering the templateBlockDeclaration production.
	EnterTemplateBlockDeclaration(c *TemplateBlockDeclarationContext)

	// EnterComponentBlockDeclaration is called when entering the componentBlockDeclaration production.
	EnterComponentBlockDeclaration(c *ComponentBlockDeclarationContext)

	// EnterLayerBlockDeclaration is called when entering the layerBlockDeclaration production.
	EnterLayerBlockDeclaration(c *LayerBlockDeclarationContext)

	// EnterCodeBlockDeclaration is called when entering the codeBlockDeclaration production.
	EnterCodeBlockDeclaration(c *CodeBlockDeclarationContext)

	// EnterDesignBodyDeclaration is called when entering the designBodyDeclaration production.
	EnterDesignBodyDeclaration(c *DesignBodyDeclarationContext)

	// EnterTemplateBodyDeclaration is called when entering the templateBodyDeclaration production.
	EnterTemplateBodyDeclaration(c *TemplateBodyDeclarationContext)

	// EnterComponentBodyDeclaration is called when entering the componentBodyDeclaration production.
	EnterComponentBodyDeclaration(c *ComponentBodyDeclarationContext)

	// EnterLayerBodyDeclaration is called when entering the layerBodyDeclaration production.
	EnterLayerBodyDeclaration(c *LayerBodyDeclarationContext)

	// EnterExpressDeclaration is called when entering the expressDeclaration production.
	EnterExpressDeclaration(c *ExpressDeclarationContext)

	// EnterExpress is called when entering the express production.
	EnterExpress(c *ExpressContext)

	// EnterEqualExpress is called when entering the equalExpress production.
	EnterEqualExpress(c *EqualExpressContext)

	// EnterUseExpress is called when entering the useExpress production.
	EnterUseExpress(c *UseExpressContext)

	// EnterValueExpress is called when entering the valueExpress production.
	EnterValueExpress(c *ValueExpressContext)

	// EnterLayerExpress is called when entering the layerExpress production.
	EnterLayerExpress(c *LayerExpressContext)

	// EnterTemplateExpress is called when entering the templateExpress production.
	EnterTemplateExpress(c *TemplateExpressContext)

	// EnterLayer is called when entering the layer production.
	EnterLayer(c *LayerContext)

	// EnterCommentDeclaration is called when entering the commentDeclaration production.
	EnterCommentDeclaration(c *CommentDeclarationContext)

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

	// ExitTemplateBlockDeclaration is called when exiting the templateBlockDeclaration production.
	ExitTemplateBlockDeclaration(c *TemplateBlockDeclarationContext)

	// ExitComponentBlockDeclaration is called when exiting the componentBlockDeclaration production.
	ExitComponentBlockDeclaration(c *ComponentBlockDeclarationContext)

	// ExitLayerBlockDeclaration is called when exiting the layerBlockDeclaration production.
	ExitLayerBlockDeclaration(c *LayerBlockDeclarationContext)

	// ExitCodeBlockDeclaration is called when exiting the codeBlockDeclaration production.
	ExitCodeBlockDeclaration(c *CodeBlockDeclarationContext)

	// ExitDesignBodyDeclaration is called when exiting the designBodyDeclaration production.
	ExitDesignBodyDeclaration(c *DesignBodyDeclarationContext)

	// ExitTemplateBodyDeclaration is called when exiting the templateBodyDeclaration production.
	ExitTemplateBodyDeclaration(c *TemplateBodyDeclarationContext)

	// ExitComponentBodyDeclaration is called when exiting the componentBodyDeclaration production.
	ExitComponentBodyDeclaration(c *ComponentBodyDeclarationContext)

	// ExitLayerBodyDeclaration is called when exiting the layerBodyDeclaration production.
	ExitLayerBodyDeclaration(c *LayerBodyDeclarationContext)

	// ExitExpressDeclaration is called when exiting the expressDeclaration production.
	ExitExpressDeclaration(c *ExpressDeclarationContext)

	// ExitExpress is called when exiting the express production.
	ExitExpress(c *ExpressContext)

	// ExitEqualExpress is called when exiting the equalExpress production.
	ExitEqualExpress(c *EqualExpressContext)

	// ExitUseExpress is called when exiting the useExpress production.
	ExitUseExpress(c *UseExpressContext)

	// ExitValueExpress is called when exiting the valueExpress production.
	ExitValueExpress(c *ValueExpressContext)

	// ExitLayerExpress is called when exiting the layerExpress production.
	ExitLayerExpress(c *LayerExpressContext)

	// ExitTemplateExpress is called when exiting the templateExpress production.
	ExitTemplateExpress(c *TemplateExpressContext)

	// ExitLayer is called when exiting the layer production.
	ExitLayer(c *LayerContext)

	// ExitCommentDeclaration is called when exiting the commentDeclaration production.
	ExitCommentDeclaration(c *CommentDeclarationContext)
}
