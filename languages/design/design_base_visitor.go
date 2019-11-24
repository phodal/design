// Code generated from Design.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Design

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseDesignVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseDesignVisitor) VisitStart(ctx *StartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitComment(ctx *CommentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitConfigDecalartion(ctx *ConfigDecalartionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitConfigKey(ctx *ConfigKeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitConfigValue(ctx *ConfigValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitDecalartions(ctx *DecalartionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitFlowDecalartion(ctx *FlowDecalartionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitFlowBodyDecalartion(ctx *FlowBodyDecalartionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitSeeDecalartion(ctx *SeeDecalartionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitDoDecalartion(ctx *DoDecalartionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitReactDecalartion(ctx *ReactDecalartionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitAnimateDecalartion(ctx *AnimateDecalartionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitActionKey(ctx *ActionKeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitActionName(ctx *ActionNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitComponentValue(ctx *ComponentValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitComponentName(ctx *ComponentNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitSceneName(ctx *SceneNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitAnimateName(ctx *AnimateNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitPageDecalartion(ctx *PageDecalartionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitComponentDecalartion(ctx *ComponentDecalartionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitComponentBodyDecalartion(ctx *ComponentBodyDecalartionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitLayoutDecalaration(ctx *LayoutDecalarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitLayoutBodyDecalartion(ctx *LayoutBodyDecalartionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitLayoutRow(ctx *LayoutRowContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitLayoutLine(ctx *LayoutLineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitComponentUseDeclaration(ctx *ComponentUseDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitStyleDecalartion(ctx *StyleDecalartionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitStyleName(ctx *StyleNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitStyleBody(ctx *StyleBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitLibraryDecalartion(ctx *LibraryDecalartionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitLibraryBody(ctx *LibraryBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitExpress(ctx *ExpressContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitLibraryCall(ctx *LibraryCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitLibraryName(ctx *LibraryNameContext) interface{} {
	return v.VisitChildren(ctx)
}
