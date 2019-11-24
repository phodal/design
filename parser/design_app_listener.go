package parser

import (
	. "../languages/design"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"reflect"
)

var projectConfigs = make(map[string]string)
var components = make(map[string]DComponent)
var flows []DFlow
var layouts []DLayout
var libraries []DLibrary

func NewDesignAppListener() *DesignAppListener {
	return &DesignAppListener{}
}

type DesignAppListener struct {
	BaseDesignListener
}

func (s *DesignAppListener) EnterConfigDeclaration(ctx *ConfigDeclarationContext) {
	projectConfigs[ctx.ConfigKey().GetText()] = ctx.ConfigValue().GetText()
}

func (s *DesignAppListener) EnterFlowDeclaration(ctx *FlowDeclarationContext) {
	flowName := ctx.IDENTIFIER().GetText()
	flow := *&DFlow{
		Interactions: nil,
		FlowName:     flowName,
	}

	declarationContexts := ctx.AllInteractionDeclaration()

	var interactions []DInteraction
	interaction := CreateInteraction()
	interactions = buildInteractions(declarationContexts, interaction, interactions)

	flow.Interactions = interactions
	flows = append(flows, flow)
}

func buildInteractions(declarationContexts []IInteractionDeclarationContext, interaction *DInteraction, interactions []DInteraction) []DInteraction {
	for _, context := range declarationContexts {
		childTypes := reflect.TypeOf(context.GetChild(0)).String()

		switch childTypes {
		case "*parser.SeeDeclarationContext":
			seeCtx := context.GetChild(0).(*SeeDeclarationContext)
			componentName := ""
			componentData := ""
			if seeCtx.IDENTIFIER() != nil {
				componentName = seeCtx.IDENTIFIER().GetText()
			} else {
				componentName = seeCtx.ComponentName().GetText()
				componentData = seeCtx.STRING_LITERAL().GetText()
			}
			seeModel := &DSee{
				ComponentName: componentName,
				Data:          componentData,
			}

			if interaction.See.ComponentName != "" {
				interactions = append(interactions, *interaction)
			}
			interaction = CreateInteraction()
			interaction.See = *seeModel
		case "*parser.DoDeclarationContext":
			doCtx := context.GetChild(0).(*DoDeclarationContext)
			doModel := &DDo{
				ComponentName: doCtx.ComponentName().GetText(),
				Data:          doCtx.STRING_LITERAL().GetText(),
				UIEvent:       doCtx.ActionName().GetText(),
			}
			interaction.Do = *doModel
		case "*parser.ReactDeclarationContext":
			reactCtx := context.GetChild(0).(*ReactDeclarationContext)
			sceneName := ""
			if reactCtx.SceneName() != nil {
				sceneName = reactCtx.SceneName().GetText()
			}
			animateName := ""
			if reactCtx.AnimateDeclaration() != nil {
				animateName = reactCtx.AnimateDeclaration().(*AnimateDeclarationContext).AnimateName().GetText()
			}
			actionName, reactComponentName, reactComponentData := buildAction(reactCtx)

			reactModel := &DReact{
				SceneName:          sceneName,
				ReactAction:        actionName,
				ReactComponentName: reactComponentName,
				ReactComponentData: reactComponentData,
				AnimateName:        animateName,
			}
			interaction.React = append(interaction.React, *reactModel)
		}
	}
	return interactions
}

func buildAction(reactCtx *ReactDeclarationContext) (string, string, string) {
	actionName := ""
	reactComponentName := ""
	reactComponentData := ""
	firstChild := reactCtx.ReactAction().GetChild(0)
	if reflect.TypeOf(firstChild).String() == "*parser.ShowActionContext" {
		showCtx := firstChild.(*ShowActionContext)
		reactComponentData = showCtx.STRING_LITERAL().GetText()
		reactComponentName = showCtx.ComponentName().GetText()
		actionName = showCtx.SHOW_KEY().GetText()
	} else if reflect.TypeOf(firstChild).String() == "*parser.GotoActionContext" {
		goCtx := firstChild.(*GotoActionContext)
		reactComponentName = goCtx.ComponentName().GetText()
		actionName = goCtx.GOTO_KEY().GetText()
	}
	return actionName, reactComponentName, reactComponentData
}

func (s *DesignAppListener) EnterComponentDeclaration(ctx *ComponentDeclarationContext) {
	componentName := ctx.IDENTIFIER().GetText()
	component := components[componentName]
	if component.Name == "" {
		components[componentName] = *CreateDComponent(componentName)
	}
	componentConfigs := make(map[string]string)
	declarations := ctx.AllComponentBodyDeclaration()
	for _, declaration := range declarations {
		childTypes := reflect.TypeOf(declaration.GetChild(0)).String()
		if childTypes == "*parser.ComponentNameContext" {
			childCtx := declaration.GetChild(0).(*ComponentNameContext)
			childComponent := *CreateDComponent(childCtx.GetText())
			component.ChildComponents = append(component.ChildComponents, childComponent)
		} else if childTypes == "*parser.ConfigKeyContext" {
			configKey := declaration.GetChild(0).(*ConfigKeyContext).GetText()
			configValue := declaration.GetChild(2).(*ConfigValueContext).GetText()

			componentConfigs[configKey] = configValue
		}

	}

	component.Configs = componentConfigs
	components[componentName] = component
}

func (s *DesignAppListener) EnterLayoutDeclaration(ctx *LayoutDeclarationContext) {
	layout := DLayout{ctx.IDENTIFIER().GetText(), nil}
	for _, row := range ctx.AllLayoutRow() {
		// TODO: refactor
		if reflect.TypeOf(row.GetChild(0)).String() != "*parser.LayoutLinesContext" {
			continue
		}

		lines := row.GetChild(0).(*LayoutLinesContext).AllLayoutLine()
		row := &DLayoutRow{"", "", nil}

		for _, line := range lines {
			declaration := line.(*LayoutLineContext).ComponentUseDeclaration()
			parseLayoutLine(declaration, row)
		}

		layout.LayoutRows = append(layout.LayoutRows, *row)
	}

	layouts = append(layouts, layout)
}

func parseLayoutLine(declaration IComponentUseDeclarationContext, layout *DLayoutRow) {
	firstChild := declaration.GetChild(0)
	switch reflect.TypeOf(firstChild).String() {
	case "*parser.ComponentNameContext":
		childCtx := firstChild.(*ComponentNameContext)
		componentName := childCtx.IDENTIFIER().GetText()
		layoutValue := ""
		if declaration.GetChildCount() > 2 {
			layoutValue = declaration.GetChild(2).(*ComponentLayoutValueContext).GetText()
		}
		layout.ComponentName = componentName
		layout.LayoutInformation = layoutValue
	default:
		componentValue := firstChild.(*antlr.TerminalNodeImpl).GetText()
		layout.NormalInformation = append(layout.NormalInformation, componentValue)
	}
}

type LibraryPreset struct {
	Key         string
	Value       string
	PresetCalls []PresetCall
}

type PresetCall struct {
	LibraryName   string
	LibraryPreset string
}

type DLibrary struct {
	LibraryName    string
	LibraryPresets []LibraryPreset
}

func (s *DesignAppListener) EnterLibraryDeclaration(ctx *LibraryDeclarationContext) {
	library := &DLibrary{
		LibraryName:    "",
		LibraryPresets: nil,
	}
	library.LibraryName = ctx.LibraryName().GetText()

	for _, express := range ctx.AllLibraryExpress() {
		preset := &LibraryPreset{
			Key:   "",
			Value: "",
			PresetCalls: nil,
		}
		preset.Key = express.(*LibraryExpressContext).PresetKey().GetText()
		pairCtx := express.GetChild(2)
		pairType := reflect.TypeOf(express.GetChild(2))
		switch pairType.String() {
		case "*parser.PresetValueContext":
			preset.Value = pairCtx.(*PresetValueContext).GetText()
		case "*parser.PresetArrayContext":
			for _, call := range pairCtx.(*PresetArrayContext).AllPresetCall() {
				presetCall := &PresetCall{
					LibraryName:   call.(*PresetCallContext).LibraryName().GetText(),
					LibraryPreset: call.(*PresetCallContext).IDENTIFIER().GetText(),
				}

				preset.PresetCalls = append(preset.PresetCalls, *presetCall)
			}
		}

		library.LibraryPresets = append(library.LibraryPresets, *preset)
	}

	libraries = append(libraries, *library)
}

func (s *DesignAppListener) getDesignInformation() {
	fmt.Println(projectConfigs)
	fmt.Println(components)
	fmt.Println(flows)
	fmt.Println(layouts)
	fmt.Println(libraries)
}
