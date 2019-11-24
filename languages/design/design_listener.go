// Code generated from Design.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Design

import "github.com/antlr/antlr4/runtime/Go/antlr"

// DesignListener is a complete listener for a parse tree produced by DesignParser.
type DesignListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterComment is called when entering the comment production.
	EnterComment(c *CommentContext)

	// EnterConfigDeclaration is called when entering the configDeclaration production.
	EnterConfigDeclaration(c *ConfigDeclarationContext)

	// EnterConfigKey is called when entering the configKey production.
	EnterConfigKey(c *ConfigKeyContext)

	// EnterConfigValue is called when entering the configValue production.
	EnterConfigValue(c *ConfigValueContext)

	// EnterUnit is called when entering the unit production.
	EnterUnit(c *UnitContext)

	// EnterDecalartions is called when entering the decalartions production.
	EnterDecalartions(c *DecalartionsContext)

	// EnterFlowDeclaration is called when entering the flowDeclaration production.
	EnterFlowDeclaration(c *FlowDeclarationContext)

	// EnterFlowBodyDeclaration is called when entering the flowBodyDeclaration production.
	EnterFlowBodyDeclaration(c *FlowBodyDeclarationContext)

	// EnterSeeDeclaration is called when entering the seeDeclaration production.
	EnterSeeDeclaration(c *SeeDeclarationContext)

	// EnterDoDeclaration is called when entering the doDeclaration production.
	EnterDoDeclaration(c *DoDeclarationContext)

	// EnterReactDeclaration is called when entering the reactDeclaration production.
	EnterReactDeclaration(c *ReactDeclarationContext)

	// EnterAnimateDeclaration is called when entering the animateDeclaration production.
	EnterAnimateDeclaration(c *AnimateDeclarationContext)

	// EnterActionKey is called when entering the actionKey production.
	EnterActionKey(c *ActionKeyContext)

	// EnterActionName is called when entering the actionName production.
	EnterActionName(c *ActionNameContext)

	// EnterComponentValue is called when entering the componentValue production.
	EnterComponentValue(c *ComponentValueContext)

	// EnterComponentName is called when entering the componentName production.
	EnterComponentName(c *ComponentNameContext)

	// EnterSceneName is called when entering the sceneName production.
	EnterSceneName(c *SceneNameContext)

	// EnterAnimateName is called when entering the animateName production.
	EnterAnimateName(c *AnimateNameContext)

	// EnterPageDeclaration is called when entering the pageDeclaration production.
	EnterPageDeclaration(c *PageDeclarationContext)

	// EnterComponentDeclaration is called when entering the componentDeclaration production.
	EnterComponentDeclaration(c *ComponentDeclarationContext)

	// EnterComponentBodyDeclaration is called when entering the componentBodyDeclaration production.
	EnterComponentBodyDeclaration(c *ComponentBodyDeclarationContext)

	// EnterLayoutDeclaration is called when entering the layoutDeclaration production.
	EnterLayoutDeclaration(c *LayoutDeclarationContext)

	// EnterLayoutBodyDeclaration is called when entering the layoutBodyDeclaration production.
	EnterLayoutBodyDeclaration(c *LayoutBodyDeclarationContext)

	// EnterLayoutRow is called when entering the layoutRow production.
	EnterLayoutRow(c *LayoutRowContext)

	// EnterLayoutLine is called when entering the layoutLine production.
	EnterLayoutLine(c *LayoutLineContext)

	// EnterComponentUseDeclaration is called when entering the componentUseDeclaration production.
	EnterComponentUseDeclaration(c *ComponentUseDeclarationContext)

	// EnterStyleDeclaration is called when entering the styleDeclaration production.
	EnterStyleDeclaration(c *StyleDeclarationContext)

	// EnterStyleName is called when entering the styleName production.
	EnterStyleName(c *StyleNameContext)

	// EnterStyleBody is called when entering the styleBody production.
	EnterStyleBody(c *StyleBodyContext)

	// EnterLibraryDeclaration is called when entering the libraryDeclaration production.
	EnterLibraryDeclaration(c *LibraryDeclarationContext)

	// EnterLibraryBody is called when entering the libraryBody production.
	EnterLibraryBody(c *LibraryBodyContext)

	// EnterExpress is called when entering the express production.
	EnterExpress(c *ExpressContext)

	// EnterLibraryCall is called when entering the libraryCall production.
	EnterLibraryCall(c *LibraryCallContext)

	// EnterLibraryName is called when entering the libraryName production.
	EnterLibraryName(c *LibraryNameContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitComment is called when exiting the comment production.
	ExitComment(c *CommentContext)

	// ExitConfigDeclaration is called when exiting the configDeclaration production.
	ExitConfigDeclaration(c *ConfigDeclarationContext)

	// ExitConfigKey is called when exiting the configKey production.
	ExitConfigKey(c *ConfigKeyContext)

	// ExitConfigValue is called when exiting the configValue production.
	ExitConfigValue(c *ConfigValueContext)

	// ExitUnit is called when exiting the unit production.
	ExitUnit(c *UnitContext)

	// ExitDecalartions is called when exiting the decalartions production.
	ExitDecalartions(c *DecalartionsContext)

	// ExitFlowDeclaration is called when exiting the flowDeclaration production.
	ExitFlowDeclaration(c *FlowDeclarationContext)

	// ExitFlowBodyDeclaration is called when exiting the flowBodyDeclaration production.
	ExitFlowBodyDeclaration(c *FlowBodyDeclarationContext)

	// ExitSeeDeclaration is called when exiting the seeDeclaration production.
	ExitSeeDeclaration(c *SeeDeclarationContext)

	// ExitDoDeclaration is called when exiting the doDeclaration production.
	ExitDoDeclaration(c *DoDeclarationContext)

	// ExitReactDeclaration is called when exiting the reactDeclaration production.
	ExitReactDeclaration(c *ReactDeclarationContext)

	// ExitAnimateDeclaration is called when exiting the animateDeclaration production.
	ExitAnimateDeclaration(c *AnimateDeclarationContext)

	// ExitActionKey is called when exiting the actionKey production.
	ExitActionKey(c *ActionKeyContext)

	// ExitActionName is called when exiting the actionName production.
	ExitActionName(c *ActionNameContext)

	// ExitComponentValue is called when exiting the componentValue production.
	ExitComponentValue(c *ComponentValueContext)

	// ExitComponentName is called when exiting the componentName production.
	ExitComponentName(c *ComponentNameContext)

	// ExitSceneName is called when exiting the sceneName production.
	ExitSceneName(c *SceneNameContext)

	// ExitAnimateName is called when exiting the animateName production.
	ExitAnimateName(c *AnimateNameContext)

	// ExitPageDeclaration is called when exiting the pageDeclaration production.
	ExitPageDeclaration(c *PageDeclarationContext)

	// ExitComponentDeclaration is called when exiting the componentDeclaration production.
	ExitComponentDeclaration(c *ComponentDeclarationContext)

	// ExitComponentBodyDeclaration is called when exiting the componentBodyDeclaration production.
	ExitComponentBodyDeclaration(c *ComponentBodyDeclarationContext)

	// ExitLayoutDeclaration is called when exiting the layoutDeclaration production.
	ExitLayoutDeclaration(c *LayoutDeclarationContext)

	// ExitLayoutBodyDeclaration is called when exiting the layoutBodyDeclaration production.
	ExitLayoutBodyDeclaration(c *LayoutBodyDeclarationContext)

	// ExitLayoutRow is called when exiting the layoutRow production.
	ExitLayoutRow(c *LayoutRowContext)

	// ExitLayoutLine is called when exiting the layoutLine production.
	ExitLayoutLine(c *LayoutLineContext)

	// ExitComponentUseDeclaration is called when exiting the componentUseDeclaration production.
	ExitComponentUseDeclaration(c *ComponentUseDeclarationContext)

	// ExitStyleDeclaration is called when exiting the styleDeclaration production.
	ExitStyleDeclaration(c *StyleDeclarationContext)

	// ExitStyleName is called when exiting the styleName production.
	ExitStyleName(c *StyleNameContext)

	// ExitStyleBody is called when exiting the styleBody production.
	ExitStyleBody(c *StyleBodyContext)

	// ExitLibraryDeclaration is called when exiting the libraryDeclaration production.
	ExitLibraryDeclaration(c *LibraryDeclarationContext)

	// ExitLibraryBody is called when exiting the libraryBody production.
	ExitLibraryBody(c *LibraryBodyContext)

	// ExitExpress is called when exiting the express production.
	ExitExpress(c *ExpressContext)

	// ExitLibraryCall is called when exiting the libraryCall production.
	ExitLibraryCall(c *LibraryCallContext)

	// ExitLibraryName is called when exiting the libraryName production.
	ExitLibraryName(c *LibraryNameContext)
}
