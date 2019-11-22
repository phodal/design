// Code generated from Design.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Design

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseDesignVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseDesignVisitor) VisitDesignIt(ctx *DesignItContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitDeclaration(ctx *DeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitDesignSystemDeclaration(ctx *DesignSystemDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitCommentBlockDeclaration(ctx *CommentBlockDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitDesignBlockDeclaration(ctx *DesignBlockDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitTemplateBlockDeclaration(ctx *TemplateBlockDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitComponentBlockDeclaration(ctx *ComponentBlockDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitLayerBlockDeclaration(ctx *LayerBlockDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitCodeBlockDeclaration(ctx *CodeBlockDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitDesignBodyDeclaration(ctx *DesignBodyDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitTemplateBodyDeclaration(ctx *TemplateBodyDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitComponentBodyDeclaration(ctx *ComponentBodyDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitLayerBodyDeclaration(ctx *LayerBodyDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitExpressDeclaration(ctx *ExpressDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitExpress(ctx *ExpressContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitEqualExpress(ctx *EqualExpressContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitUseExpress(ctx *UseExpressContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitValueExpress(ctx *ValueExpressContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitLayerExpress(ctx *LayerExpressContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitTemplateExpress(ctx *TemplateExpressContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitExpressKey(ctx *ExpressKeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitExpressValue(ctx *ExpressValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitLayer(ctx *LayerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitCommentDeclaration(ctx *CommentDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}
