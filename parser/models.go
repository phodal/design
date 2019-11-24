package parser


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

func CreateDComponent(componentName string) *DComponent {
	return &DComponent{componentName, nil, nil}
}
