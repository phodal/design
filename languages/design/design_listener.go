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

	// EnterConfigDecalartion is called when entering the configDecalartion production.
	EnterConfigDecalartion(c *ConfigDecalartionContext)

	// EnterConfigKey is called when entering the configKey production.
	EnterConfigKey(c *ConfigKeyContext)

	// EnterConfigValue is called when entering the configValue production.
	EnterConfigValue(c *ConfigValueContext)

	// EnterDecalartions is called when entering the decalartions production.
	EnterDecalartions(c *DecalartionsContext)

	// EnterFlowDecalartion is called when entering the flowDecalartion production.
	EnterFlowDecalartion(c *FlowDecalartionContext)

	// EnterFlowBodyDecalartion is called when entering the flowBodyDecalartion production.
	EnterFlowBodyDecalartion(c *FlowBodyDecalartionContext)

	// EnterSeeDecalartion is called when entering the seeDecalartion production.
	EnterSeeDecalartion(c *SeeDecalartionContext)

	// EnterDoDecalartion is called when entering the doDecalartion production.
	EnterDoDecalartion(c *DoDecalartionContext)

	// EnterReactDecalartion is called when entering the reactDecalartion production.
	EnterReactDecalartion(c *ReactDecalartionContext)

	// EnterAnimateDecalartion is called when entering the animateDecalartion production.
	EnterAnimateDecalartion(c *AnimateDecalartionContext)

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

	// EnterPageDecalartion is called when entering the pageDecalartion production.
	EnterPageDecalartion(c *PageDecalartionContext)

	// EnterComponentDecalartion is called when entering the componentDecalartion production.
	EnterComponentDecalartion(c *ComponentDecalartionContext)

	// EnterComponentBodyDecalartion is called when entering the componentBodyDecalartion production.
	EnterComponentBodyDecalartion(c *ComponentBodyDecalartionContext)

	// EnterLayoutDecalaration is called when entering the layoutDecalaration production.
	EnterLayoutDecalaration(c *LayoutDecalarationContext)

	// EnterLayoutBodyDecalartion is called when entering the layoutBodyDecalartion production.
	EnterLayoutBodyDecalartion(c *LayoutBodyDecalartionContext)

	// EnterEmptyLine is called when entering the emptyLine production.
	EnterEmptyLine(c *EmptyLineContext)

	// EnterLayoutLine is called when entering the layoutLine production.
	EnterLayoutLine(c *LayoutLineContext)

	// EnterComponentUseDeclaration is called when entering the componentUseDeclaration production.
	EnterComponentUseDeclaration(c *ComponentUseDeclarationContext)

	// EnterStyleDecalartion is called when entering the styleDecalartion production.
	EnterStyleDecalartion(c *StyleDecalartionContext)

	// EnterStyleName is called when entering the styleName production.
	EnterStyleName(c *StyleNameContext)

	// EnterStyleBody is called when entering the styleBody production.
	EnterStyleBody(c *StyleBodyContext)

	// EnterLibraryDecalartion is called when entering the libraryDecalartion production.
	EnterLibraryDecalartion(c *LibraryDecalartionContext)

	// EnterLibraryBody is called when entering the libraryBody production.
	EnterLibraryBody(c *LibraryBodyContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitComment is called when exiting the comment production.
	ExitComment(c *CommentContext)

	// ExitConfigDecalartion is called when exiting the configDecalartion production.
	ExitConfigDecalartion(c *ConfigDecalartionContext)

	// ExitConfigKey is called when exiting the configKey production.
	ExitConfigKey(c *ConfigKeyContext)

	// ExitConfigValue is called when exiting the configValue production.
	ExitConfigValue(c *ConfigValueContext)

	// ExitDecalartions is called when exiting the decalartions production.
	ExitDecalartions(c *DecalartionsContext)

	// ExitFlowDecalartion is called when exiting the flowDecalartion production.
	ExitFlowDecalartion(c *FlowDecalartionContext)

	// ExitFlowBodyDecalartion is called when exiting the flowBodyDecalartion production.
	ExitFlowBodyDecalartion(c *FlowBodyDecalartionContext)

	// ExitSeeDecalartion is called when exiting the seeDecalartion production.
	ExitSeeDecalartion(c *SeeDecalartionContext)

	// ExitDoDecalartion is called when exiting the doDecalartion production.
	ExitDoDecalartion(c *DoDecalartionContext)

	// ExitReactDecalartion is called when exiting the reactDecalartion production.
	ExitReactDecalartion(c *ReactDecalartionContext)

	// ExitAnimateDecalartion is called when exiting the animateDecalartion production.
	ExitAnimateDecalartion(c *AnimateDecalartionContext)

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

	// ExitPageDecalartion is called when exiting the pageDecalartion production.
	ExitPageDecalartion(c *PageDecalartionContext)

	// ExitComponentDecalartion is called when exiting the componentDecalartion production.
	ExitComponentDecalartion(c *ComponentDecalartionContext)

	// ExitComponentBodyDecalartion is called when exiting the componentBodyDecalartion production.
	ExitComponentBodyDecalartion(c *ComponentBodyDecalartionContext)

	// ExitLayoutDecalaration is called when exiting the layoutDecalaration production.
	ExitLayoutDecalaration(c *LayoutDecalarationContext)

	// ExitLayoutBodyDecalartion is called when exiting the layoutBodyDecalartion production.
	ExitLayoutBodyDecalartion(c *LayoutBodyDecalartionContext)

	// ExitEmptyLine is called when exiting the emptyLine production.
	ExitEmptyLine(c *EmptyLineContext)

	// ExitLayoutLine is called when exiting the layoutLine production.
	ExitLayoutLine(c *LayoutLineContext)

	// ExitComponentUseDeclaration is called when exiting the componentUseDeclaration production.
	ExitComponentUseDeclaration(c *ComponentUseDeclarationContext)

	// ExitStyleDecalartion is called when exiting the styleDecalartion production.
	ExitStyleDecalartion(c *StyleDecalartionContext)

	// ExitStyleName is called when exiting the styleName production.
	ExitStyleName(c *StyleNameContext)

	// ExitStyleBody is called when exiting the styleBody production.
	ExitStyleBody(c *StyleBodyContext)

	// ExitLibraryDecalartion is called when exiting the libraryDecalartion production.
	ExitLibraryDecalartion(c *LibraryDecalartionContext)

	// ExitLibraryBody is called when exiting the libraryBody production.
	ExitLibraryBody(c *LibraryBodyContext)
}
