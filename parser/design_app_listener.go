package parser

import (
	. "../languages/design"
	"fmt"
	"reflect"
)

var projectConfigs = make(map[string]string)

type DConfig struct {
	Key   string
	Value string
}

type DComponent struct {
	Name            string `json:"name"`
	ChildComponents []DComponent
	Configs         map[string]string
}

type DSee struct {
	ComponentName string
	Data          string
}

type DDo struct {
	ComponentName string
	Data          string
	UIEvent       string // event
}

type DReact struct {
	ContextName        string
	ReactEvent         string
	ReactComponentName string
	AnimateName        string
}

type DInteraction struct {
	See   DSee
	Do    DDo
	React map[string]DReact
}

type DFlow struct {
	Interactions []DInteraction
	FlowName     string
}

func (s *DesignAppListener) createComponent(componentName string) *DComponent {
	return &DComponent{componentName, nil, nil}
}

var components = make(map[string]DComponent)

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

}

func (s *DesignAppListener) EnterComponentDeclaration(ctx *ComponentDeclarationContext) {
	componentName := ctx.IDENTIFIER().GetText()
	component := components[componentName]
	if component.Name == "" {
		components[componentName] = *s.createComponent(componentName)
	}
	componentConfigs := make(map[string]string)
	declarations := ctx.AllComponentBodyDeclaration()
	for _, declaration := range declarations {
		childTypes := reflect.TypeOf(declaration.GetChild(0)).String()
		if childTypes == "*parser.ComponentNameContext" {
			childCtx := declaration.GetChild(0).(*ComponentNameContext)
			childComponent := *s.createComponent(childCtx.GetText())
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
}
