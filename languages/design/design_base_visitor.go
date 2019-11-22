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

func (v *BaseDesignVisitor) VisitDesignBodyDeclaration(ctx *DesignBodyDeclarationContext) interface{} {
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

func (v *BaseDesignVisitor) VisitTemplateExpress(ctx *TemplateExpressContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitLayer(ctx *LayerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitComment(ctx *CommentContext) interface{} {
	return v.VisitChildren(ctx)
}
