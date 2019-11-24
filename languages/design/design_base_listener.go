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

// EnterConfigDeclaration is called when production configDeclaration is entered.
func (s *BaseDesignListener) EnterConfigDeclaration(ctx *ConfigDeclarationContext) {}

// ExitConfigDeclaration is called when production configDeclaration is exited.
func (s *BaseDesignListener) ExitConfigDeclaration(ctx *ConfigDeclarationContext) {}

// EnterConfigKey is called when production configKey is entered.
func (s *BaseDesignListener) EnterConfigKey(ctx *ConfigKeyContext) {}

// ExitConfigKey is called when production configKey is exited.
func (s *BaseDesignListener) ExitConfigKey(ctx *ConfigKeyContext) {}

// EnterConfigValue is called when production configValue is entered.
func (s *BaseDesignListener) EnterConfigValue(ctx *ConfigValueContext) {}

// ExitConfigValue is called when production configValue is exited.
func (s *BaseDesignListener) ExitConfigValue(ctx *ConfigValueContext) {}

// EnterUnit is called when production unit is entered.
func (s *BaseDesignListener) EnterUnit(ctx *UnitContext) {}

// ExitUnit is called when production unit is exited.
func (s *BaseDesignListener) ExitUnit(ctx *UnitContext) {}

// EnterDecalartions is called when production decalartions is entered.
func (s *BaseDesignListener) EnterDecalartions(ctx *DecalartionsContext) {}

// ExitDecalartions is called when production decalartions is exited.
func (s *BaseDesignListener) ExitDecalartions(ctx *DecalartionsContext) {}

// EnterFlowDeclaration is called when production flowDeclaration is entered.
func (s *BaseDesignListener) EnterFlowDeclaration(ctx *FlowDeclarationContext) {}

// ExitFlowDeclaration is called when production flowDeclaration is exited.
func (s *BaseDesignListener) ExitFlowDeclaration(ctx *FlowDeclarationContext) {}

// EnterFlowBodyDeclaration is called when production flowBodyDeclaration is entered.
func (s *BaseDesignListener) EnterFlowBodyDeclaration(ctx *FlowBodyDeclarationContext) {}

// ExitFlowBodyDeclaration is called when production flowBodyDeclaration is exited.
func (s *BaseDesignListener) ExitFlowBodyDeclaration(ctx *FlowBodyDeclarationContext) {}

// EnterSeeDeclaration is called when production seeDeclaration is entered.
func (s *BaseDesignListener) EnterSeeDeclaration(ctx *SeeDeclarationContext) {}

// ExitSeeDeclaration is called when production seeDeclaration is exited.
func (s *BaseDesignListener) ExitSeeDeclaration(ctx *SeeDeclarationContext) {}

// EnterDoDeclaration is called when production doDeclaration is entered.
func (s *BaseDesignListener) EnterDoDeclaration(ctx *DoDeclarationContext) {}

// ExitDoDeclaration is called when production doDeclaration is exited.
func (s *BaseDesignListener) ExitDoDeclaration(ctx *DoDeclarationContext) {}

// EnterReactDeclaration is called when production reactDeclaration is entered.
func (s *BaseDesignListener) EnterReactDeclaration(ctx *ReactDeclarationContext) {}

// ExitReactDeclaration is called when production reactDeclaration is exited.
func (s *BaseDesignListener) ExitReactDeclaration(ctx *ReactDeclarationContext) {}

// EnterAnimateDeclaration is called when production animateDeclaration is entered.
func (s *BaseDesignListener) EnterAnimateDeclaration(ctx *AnimateDeclarationContext) {}

// ExitAnimateDeclaration is called when production animateDeclaration is exited.
func (s *BaseDesignListener) ExitAnimateDeclaration(ctx *AnimateDeclarationContext) {}

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

// EnterPageDeclaration is called when production pageDeclaration is entered.
func (s *BaseDesignListener) EnterPageDeclaration(ctx *PageDeclarationContext) {}

// ExitPageDeclaration is called when production pageDeclaration is exited.
func (s *BaseDesignListener) ExitPageDeclaration(ctx *PageDeclarationContext) {}

// EnterComponentDeclaration is called when production componentDeclaration is entered.
func (s *BaseDesignListener) EnterComponentDeclaration(ctx *ComponentDeclarationContext) {}

// ExitComponentDeclaration is called when production componentDeclaration is exited.
func (s *BaseDesignListener) ExitComponentDeclaration(ctx *ComponentDeclarationContext) {}

// EnterComponentBodyDeclaration is called when production componentBodyDeclaration is entered.
func (s *BaseDesignListener) EnterComponentBodyDeclaration(ctx *ComponentBodyDeclarationContext) {}

// ExitComponentBodyDeclaration is called when production componentBodyDeclaration is exited.
func (s *BaseDesignListener) ExitComponentBodyDeclaration(ctx *ComponentBodyDeclarationContext) {}

// EnterLayoutDeclaration is called when production layoutDeclaration is entered.
func (s *BaseDesignListener) EnterLayoutDeclaration(ctx *LayoutDeclarationContext) {}

// ExitLayoutDeclaration is called when production layoutDeclaration is exited.
func (s *BaseDesignListener) ExitLayoutDeclaration(ctx *LayoutDeclarationContext) {}

// EnterLayoutBodyDeclaration is called when production layoutBodyDeclaration is entered.
func (s *BaseDesignListener) EnterLayoutBodyDeclaration(ctx *LayoutBodyDeclarationContext) {}

// ExitLayoutBodyDeclaration is called when production layoutBodyDeclaration is exited.
func (s *BaseDesignListener) ExitLayoutBodyDeclaration(ctx *LayoutBodyDeclarationContext) {}

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

// EnterStyleDeclaration is called when production styleDeclaration is entered.
func (s *BaseDesignListener) EnterStyleDeclaration(ctx *StyleDeclarationContext) {}

// ExitStyleDeclaration is called when production styleDeclaration is exited.
func (s *BaseDesignListener) ExitStyleDeclaration(ctx *StyleDeclarationContext) {}

// EnterStyleName is called when production styleName is entered.
func (s *BaseDesignListener) EnterStyleName(ctx *StyleNameContext) {}

// ExitStyleName is called when production styleName is exited.
func (s *BaseDesignListener) ExitStyleName(ctx *StyleNameContext) {}

// EnterStyleBody is called when production styleBody is entered.
func (s *BaseDesignListener) EnterStyleBody(ctx *StyleBodyContext) {}

// ExitStyleBody is called when production styleBody is exited.
func (s *BaseDesignListener) ExitStyleBody(ctx *StyleBodyContext) {}

// EnterLibraryDeclaration is called when production libraryDeclaration is entered.
func (s *BaseDesignListener) EnterLibraryDeclaration(ctx *LibraryDeclarationContext) {}

// ExitLibraryDeclaration is called when production libraryDeclaration is exited.
func (s *BaseDesignListener) ExitLibraryDeclaration(ctx *LibraryDeclarationContext) {}

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
