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

	// Visit a parse tree produced by DesignParser#configDeclaration.
	VisitConfigDeclaration(ctx *ConfigDeclarationContext) interface{}

	// Visit a parse tree produced by DesignParser#configKey.
	VisitConfigKey(ctx *ConfigKeyContext) interface{}

	// Visit a parse tree produced by DesignParser#configValue.
	VisitConfigValue(ctx *ConfigValueContext) interface{}

	// Visit a parse tree produced by DesignParser#unit.
	VisitUnit(ctx *UnitContext) interface{}

	// Visit a parse tree produced by DesignParser#decalartions.
	VisitDecalartions(ctx *DecalartionsContext) interface{}

	// Visit a parse tree produced by DesignParser#flowDeclaration.
	VisitFlowDeclaration(ctx *FlowDeclarationContext) interface{}

	// Visit a parse tree produced by DesignParser#interactionDeclaration.
	VisitInteractionDeclaration(ctx *InteractionDeclarationContext) interface{}

	// Visit a parse tree produced by DesignParser#seeDeclaration.
	VisitSeeDeclaration(ctx *SeeDeclarationContext) interface{}

	// Visit a parse tree produced by DesignParser#doDeclaration.
	VisitDoDeclaration(ctx *DoDeclarationContext) interface{}

	// Visit a parse tree produced by DesignParser#reactDeclaration.
	VisitReactDeclaration(ctx *ReactDeclarationContext) interface{}

	// Visit a parse tree produced by DesignParser#animateDeclaration.
	VisitAnimateDeclaration(ctx *AnimateDeclarationContext) interface{}

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

	// Visit a parse tree produced by DesignParser#pageDeclaration.
	VisitPageDeclaration(ctx *PageDeclarationContext) interface{}

	// Visit a parse tree produced by DesignParser#componentDeclaration.
	VisitComponentDeclaration(ctx *ComponentDeclarationContext) interface{}

	// Visit a parse tree produced by DesignParser#componentBodyDeclaration.
	VisitComponentBodyDeclaration(ctx *ComponentBodyDeclarationContext) interface{}

	// Visit a parse tree produced by DesignParser#layoutDeclaration.
	VisitLayoutDeclaration(ctx *LayoutDeclarationContext) interface{}

	// Visit a parse tree produced by DesignParser#layoutBodyDeclaration.
	VisitLayoutBodyDeclaration(ctx *LayoutBodyDeclarationContext) interface{}

	// Visit a parse tree produced by DesignParser#layoutRow.
	VisitLayoutRow(ctx *LayoutRowContext) interface{}

	// Visit a parse tree produced by DesignParser#layoutLine.
	VisitLayoutLine(ctx *LayoutLineContext) interface{}

	// Visit a parse tree produced by DesignParser#componentUseDeclaration.
	VisitComponentUseDeclaration(ctx *ComponentUseDeclarationContext) interface{}

	// Visit a parse tree produced by DesignParser#styleDeclaration.
	VisitStyleDeclaration(ctx *StyleDeclarationContext) interface{}

	// Visit a parse tree produced by DesignParser#styleName.
	VisitStyleName(ctx *StyleNameContext) interface{}

	// Visit a parse tree produced by DesignParser#styleBody.
	VisitStyleBody(ctx *StyleBodyContext) interface{}

	// Visit a parse tree produced by DesignParser#libraryDeclaration.
	VisitLibraryDeclaration(ctx *LibraryDeclarationContext) interface{}

	// Visit a parse tree produced by DesignParser#libraryBody.
	VisitLibraryBody(ctx *LibraryBodyContext) interface{}

	// Visit a parse tree produced by DesignParser#express.
	VisitExpress(ctx *ExpressContext) interface{}

	// Visit a parse tree produced by DesignParser#libraryCall.
	VisitLibraryCall(ctx *LibraryCallContext) interface{}

	// Visit a parse tree produced by DesignParser#libraryName.
	VisitLibraryName(ctx *LibraryNameContext) interface{}
}
