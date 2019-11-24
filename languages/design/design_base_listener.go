// Code generated from Design.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Design

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseDesignListener is a complete listener for a parse tree produced by DesignParser.
type BaseDesignListener struct{}

var _ DesignListener = &BaseDesignListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDesignListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDesignListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDesignListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDesignListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BaseDesignListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseDesignListener) ExitStart(ctx *StartContext) {}

// EnterComment is called when production comment is entered.
func (s *BaseDesignListener) EnterComment(ctx *CommentContext) {}

// ExitComment is called when production comment is exited.
func (s *BaseDesignListener) ExitComment(ctx *CommentContext) {}

// EnterConfigDecalartion is called when production configDecalartion is entered.
func (s *BaseDesignListener) EnterConfigDecalartion(ctx *ConfigDecalartionContext) {}

// ExitConfigDecalartion is called when production configDecalartion is exited.
func (s *BaseDesignListener) ExitConfigDecalartion(ctx *ConfigDecalartionContext) {}

// EnterConfigKey is called when production configKey is entered.
func (s *BaseDesignListener) EnterConfigKey(ctx *ConfigKeyContext) {}

// ExitConfigKey is called when production configKey is exited.
func (s *BaseDesignListener) ExitConfigKey(ctx *ConfigKeyContext) {}

// EnterConfigValue is called when production configValue is entered.
func (s *BaseDesignListener) EnterConfigValue(ctx *ConfigValueContext) {}

// ExitConfigValue is called when production configValue is exited.
func (s *BaseDesignListener) ExitConfigValue(ctx *ConfigValueContext) {}

// EnterDecalartions is called when production decalartions is entered.
func (s *BaseDesignListener) EnterDecalartions(ctx *DecalartionsContext) {}

// ExitDecalartions is called when production decalartions is exited.
func (s *BaseDesignListener) ExitDecalartions(ctx *DecalartionsContext) {}

// EnterFlowDecalartion is called when production flowDecalartion is entered.
func (s *BaseDesignListener) EnterFlowDecalartion(ctx *FlowDecalartionContext) {}

// ExitFlowDecalartion is called when production flowDecalartion is exited.
func (s *BaseDesignListener) ExitFlowDecalartion(ctx *FlowDecalartionContext) {}

// EnterFlowBodyDecalartion is called when production flowBodyDecalartion is entered.
func (s *BaseDesignListener) EnterFlowBodyDecalartion(ctx *FlowBodyDecalartionContext) {}

// ExitFlowBodyDecalartion is called when production flowBodyDecalartion is exited.
func (s *BaseDesignListener) ExitFlowBodyDecalartion(ctx *FlowBodyDecalartionContext) {}

// EnterSeeDecalartion is called when production seeDecalartion is entered.
func (s *BaseDesignListener) EnterSeeDecalartion(ctx *SeeDecalartionContext) {}

// ExitSeeDecalartion is called when production seeDecalartion is exited.
func (s *BaseDesignListener) ExitSeeDecalartion(ctx *SeeDecalartionContext) {}

// EnterDoDecalartion is called when production doDecalartion is entered.
func (s *BaseDesignListener) EnterDoDecalartion(ctx *DoDecalartionContext) {}

// ExitDoDecalartion is called when production doDecalartion is exited.
func (s *BaseDesignListener) ExitDoDecalartion(ctx *DoDecalartionContext) {}

// EnterReactDecalartion is called when production reactDecalartion is entered.
func (s *BaseDesignListener) EnterReactDecalartion(ctx *ReactDecalartionContext) {}

// ExitReactDecalartion is called when production reactDecalartion is exited.
func (s *BaseDesignListener) ExitReactDecalartion(ctx *ReactDecalartionContext) {}

// EnterAnimateDecalartion is called when production animateDecalartion is entered.
func (s *BaseDesignListener) EnterAnimateDecalartion(ctx *AnimateDecalartionContext) {}

// ExitAnimateDecalartion is called when production animateDecalartion is exited.
func (s *BaseDesignListener) ExitAnimateDecalartion(ctx *AnimateDecalartionContext) {}

// EnterActionKey is called when production actionKey is entered.
func (s *BaseDesignListener) EnterActionKey(ctx *ActionKeyContext) {}

// ExitActionKey is called when production actionKey is exited.
func (s *BaseDesignListener) ExitActionKey(ctx *ActionKeyContext) {}

// EnterActionName is called when production actionName is entered.
func (s *BaseDesignListener) EnterActionName(ctx *ActionNameContext) {}

// ExitActionName is called when production actionName is exited.
func (s *BaseDesignListener) ExitActionName(ctx *ActionNameContext) {}

// EnterComponentValue is called when production componentValue is entered.
func (s *BaseDesignListener) EnterComponentValue(ctx *ComponentValueContext) {}

// ExitComponentValue is called when production componentValue is exited.
func (s *BaseDesignListener) ExitComponentValue(ctx *ComponentValueContext) {}

// EnterComponentName is called when production componentName is entered.
func (s *BaseDesignListener) EnterComponentName(ctx *ComponentNameContext) {}

// ExitComponentName is called when production componentName is exited.
func (s *BaseDesignListener) ExitComponentName(ctx *ComponentNameContext) {}

// EnterSceneName is called when production sceneName is entered.
func (s *BaseDesignListener) EnterSceneName(ctx *SceneNameContext) {}

// ExitSceneName is called when production sceneName is exited.
func (s *BaseDesignListener) ExitSceneName(ctx *SceneNameContext) {}

// EnterAnimateName is called when production animateName is entered.
func (s *BaseDesignListener) EnterAnimateName(ctx *AnimateNameContext) {}

// ExitAnimateName is called when production animateName is exited.
func (s *BaseDesignListener) ExitAnimateName(ctx *AnimateNameContext) {}

// EnterPageDecalartion is called when production pageDecalartion is entered.
func (s *BaseDesignListener) EnterPageDecalartion(ctx *PageDecalartionContext) {}

// ExitPageDecalartion is called when production pageDecalartion is exited.
func (s *BaseDesignListener) ExitPageDecalartion(ctx *PageDecalartionContext) {}

// EnterComponentDecalartion is called when production componentDecalartion is entered.
func (s *BaseDesignListener) EnterComponentDecalartion(ctx *ComponentDecalartionContext) {}

// ExitComponentDecalartion is called when production componentDecalartion is exited.
func (s *BaseDesignListener) ExitComponentDecalartion(ctx *ComponentDecalartionContext) {}

// EnterComponentBodyDecalartion is called when production componentBodyDecalartion is entered.
func (s *BaseDesignListener) EnterComponentBodyDecalartion(ctx *ComponentBodyDecalartionContext) {}

// ExitComponentBodyDecalartion is called when production componentBodyDecalartion is exited.
func (s *BaseDesignListener) ExitComponentBodyDecalartion(ctx *ComponentBodyDecalartionContext) {}

// EnterLayoutDecalaration is called when production layoutDecalaration is entered.
func (s *BaseDesignListener) EnterLayoutDecalaration(ctx *LayoutDecalarationContext) {}

// ExitLayoutDecalaration is called when production layoutDecalaration is exited.
func (s *BaseDesignListener) ExitLayoutDecalaration(ctx *LayoutDecalarationContext) {}

// EnterLayoutBodyDecalartion is called when production layoutBodyDecalartion is entered.
func (s *BaseDesignListener) EnterLayoutBodyDecalartion(ctx *LayoutBodyDecalartionContext) {}

// ExitLayoutBodyDecalartion is called when production layoutBodyDecalartion is exited.
func (s *BaseDesignListener) ExitLayoutBodyDecalartion(ctx *LayoutBodyDecalartionContext) {}

// EnterLayoutRow is called when production layoutRow is entered.
func (s *BaseDesignListener) EnterLayoutRow(ctx *LayoutRowContext) {}

// ExitLayoutRow is called when production layoutRow is exited.
func (s *BaseDesignListener) ExitLayoutRow(ctx *LayoutRowContext) {}

// EnterLayoutLine is called when production layoutLine is entered.
func (s *BaseDesignListener) EnterLayoutLine(ctx *LayoutLineContext) {}

// ExitLayoutLine is called when production layoutLine is exited.
func (s *BaseDesignListener) ExitLayoutLine(ctx *LayoutLineContext) {}

// EnterComponentUseDeclaration is called when production componentUseDeclaration is entered.
func (s *BaseDesignListener) EnterComponentUseDeclaration(ctx *ComponentUseDeclarationContext) {}

// ExitComponentUseDeclaration is called when production componentUseDeclaration is exited.
func (s *BaseDesignListener) ExitComponentUseDeclaration(ctx *ComponentUseDeclarationContext) {}

// EnterStyleDecalartion is called when production styleDecalartion is entered.
func (s *BaseDesignListener) EnterStyleDecalartion(ctx *StyleDecalartionContext) {}

// ExitStyleDecalartion is called when production styleDecalartion is exited.
func (s *BaseDesignListener) ExitStyleDecalartion(ctx *StyleDecalartionContext) {}

// EnterStyleName is called when production styleName is entered.
func (s *BaseDesignListener) EnterStyleName(ctx *StyleNameContext) {}

// ExitStyleName is called when production styleName is exited.
func (s *BaseDesignListener) ExitStyleName(ctx *StyleNameContext) {}

// EnterStyleBody is called when production styleBody is entered.
func (s *BaseDesignListener) EnterStyleBody(ctx *StyleBodyContext) {}

// ExitStyleBody is called when production styleBody is exited.
func (s *BaseDesignListener) ExitStyleBody(ctx *StyleBodyContext) {}

// EnterLibraryDecalartion is called when production libraryDecalartion is entered.
func (s *BaseDesignListener) EnterLibraryDecalartion(ctx *LibraryDecalartionContext) {}

// ExitLibraryDecalartion is called when production libraryDecalartion is exited.
func (s *BaseDesignListener) ExitLibraryDecalartion(ctx *LibraryDecalartionContext) {}

// EnterLibraryBody is called when production libraryBody is entered.
func (s *BaseDesignListener) EnterLibraryBody(ctx *LibraryBodyContext) {}

// ExitLibraryBody is called when production libraryBody is exited.
func (s *BaseDesignListener) ExitLibraryBody(ctx *LibraryBodyContext) {}

// EnterExpress is called when production express is entered.
func (s *BaseDesignListener) EnterExpress(ctx *ExpressContext) {}

// ExitExpress is called when production express is exited.
func (s *BaseDesignListener) ExitExpress(ctx *ExpressContext) {}

// EnterLibraryCall is called when production libraryCall is entered.
func (s *BaseDesignListener) EnterLibraryCall(ctx *LibraryCallContext) {}

// ExitLibraryCall is called when production libraryCall is exited.
func (s *BaseDesignListener) ExitLibraryCall(ctx *LibraryCallContext) {}

// EnterLibraryName is called when production libraryName is entered.
func (s *BaseDesignListener) EnterLibraryName(ctx *LibraryNameContext) {}

// ExitLibraryName is called when production libraryName is exited.
func (s *BaseDesignListener) ExitLibraryName(ctx *LibraryNameContext) {}
