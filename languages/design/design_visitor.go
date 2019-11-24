// Code generated from Design.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Design

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by DesignParser.
type DesignVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by DesignParser#start.
	VisitStart(ctx *StartContext) interface{}

	// Visit a parse tree produced by DesignParser#comment.
	VisitComment(ctx *CommentContext) interface{}

	// Visit a parse tree produced by DesignParser#configDecalartion.
	VisitConfigDecalartion(ctx *ConfigDecalartionContext) interface{}

	// Visit a parse tree produced by DesignParser#configKey.
	VisitConfigKey(ctx *ConfigKeyContext) interface{}

	// Visit a parse tree produced by DesignParser#configValue.
	VisitConfigValue(ctx *ConfigValueContext) interface{}

	// Visit a parse tree produced by DesignParser#unit.
	VisitUnit(ctx *UnitContext) interface{}

	// Visit a parse tree produced by DesignParser#decalartions.
	VisitDecalartions(ctx *DecalartionsContext) interface{}

	// Visit a parse tree produced by DesignParser#flowDecalartion.
	VisitFlowDecalartion(ctx *FlowDecalartionContext) interface{}

	// Visit a parse tree produced by DesignParser#flowBodyDecalartion.
	VisitFlowBodyDecalartion(ctx *FlowBodyDecalartionContext) interface{}

	// Visit a parse tree produced by DesignParser#seeDecalartion.
	VisitSeeDecalartion(ctx *SeeDecalartionContext) interface{}

	// Visit a parse tree produced by DesignParser#doDecalartion.
	VisitDoDecalartion(ctx *DoDecalartionContext) interface{}

	// Visit a parse tree produced by DesignParser#reactDecalartion.
	VisitReactDecalartion(ctx *ReactDecalartionContext) interface{}

	// Visit a parse tree produced by DesignParser#animateDecalartion.
	VisitAnimateDecalartion(ctx *AnimateDecalartionContext) interface{}

	// Visit a parse tree produced by DesignParser#actionKey.
	VisitActionKey(ctx *ActionKeyContext) interface{}

	// Visit a parse tree produced by DesignParser#actionName.
	VisitActionName(ctx *ActionNameContext) interface{}

	// Visit a parse tree produced by DesignParser#componentValue.
	VisitComponentValue(ctx *ComponentValueContext) interface{}

	// Visit a parse tree produced by DesignParser#componentName.
	VisitComponentName(ctx *ComponentNameContext) interface{}

	// Visit a parse tree produced by DesignParser#sceneName.
	VisitSceneName(ctx *SceneNameContext) interface{}

	// Visit a parse tree produced by DesignParser#animateName.
	VisitAnimateName(ctx *AnimateNameContext) interface{}

	// Visit a parse tree produced by DesignParser#pageDecalartion.
	VisitPageDecalartion(ctx *PageDecalartionContext) interface{}

	// Visit a parse tree produced by DesignParser#componentDecalartion.
	VisitComponentDecalartion(ctx *ComponentDecalartionContext) interface{}

	// Visit a parse tree produced by DesignParser#componentBodyDecalartion.
	VisitComponentBodyDecalartion(ctx *ComponentBodyDecalartionContext) interface{}

	// Visit a parse tree produced by DesignParser#layoutDecalaration.
	VisitLayoutDecalaration(ctx *LayoutDecalarationContext) interface{}

	// Visit a parse tree produced by DesignParser#layoutBodyDecalartion.
	VisitLayoutBodyDecalartion(ctx *LayoutBodyDecalartionContext) interface{}

	// Visit a parse tree produced by DesignParser#layoutRow.
	VisitLayoutRow(ctx *LayoutRowContext) interface{}

	// Visit a parse tree produced by DesignParser#layoutLine.
	VisitLayoutLine(ctx *LayoutLineContext) interface{}

	// Visit a parse tree produced by DesignParser#componentUseDeclaration.
	VisitComponentUseDeclaration(ctx *ComponentUseDeclarationContext) interface{}

	// Visit a parse tree produced by DesignParser#styleDecalartion.
	VisitStyleDecalartion(ctx *StyleDecalartionContext) interface{}

	// Visit a parse tree produced by DesignParser#styleName.
	VisitStyleName(ctx *StyleNameContext) interface{}

	// Visit a parse tree produced by DesignParser#styleBody.
	VisitStyleBody(ctx *StyleBodyContext) interface{}

	// Visit a parse tree produced by DesignParser#libraryDecalartion.
	VisitLibraryDecalartion(ctx *LibraryDecalartionContext) interface{}

	// Visit a parse tree produced by DesignParser#libraryBody.
	VisitLibraryBody(ctx *LibraryBodyContext) interface{}

	// Visit a parse tree produced by DesignParser#express.
	VisitExpress(ctx *ExpressContext) interface{}

	// Visit a parse tree produced by DesignParser#libraryCall.
	VisitLibraryCall(ctx *LibraryCallContext) interface{}

	// Visit a parse tree produced by DesignParser#libraryName.
	VisitLibraryName(ctx *LibraryNameContext) interface{}
}
