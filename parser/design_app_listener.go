package parser

import (
	. "../languages/design"
	"fmt"
	"reflect"
)

var projectConfigs = make(map[string]string)
var components = make(map[string]DComponent)
var flows []DFlow

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

			interactions = append(interactions, *interaction)
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

		}
	}

	flow.Interactions = interactions
	flows = append(flows, flow)
}

func CreateInteraction() *DInteraction {
	seeModel := &DSee{"", ""}
	doModel := &DDo{"", "", ""}
	return &DInteraction{
		See:   *seeModel,
		Do:    *doModel,
		React: nil,
	}
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
	//fmt.Println(ctx.GetText())
}

func (s *DesignAppListener) EnterLayoutRow(ctx *LayoutRowContext) {
	//fmt.Println(ctx.GetText())
}

func (s *DesignAppListener) EnterComponentUseDeclaration(ctx *ComponentUseDeclarationContext) {
	//fmt.Println(ctx.GetText())
}

func (s *DesignAppListener) getDesignInformation() {
	fmt.Println(projectConfigs)
	fmt.Println(components)
	fmt.Println(flows)
}
