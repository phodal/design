package parser

type DesignInformation struct {
	ProjectConfigs map[string]string     `json:"projectConfigs"`
	Flows          []DFlow               `json:"flows"`
	Components     map[string]DComponent `json:"components"`
	Layouts        []DLayout             `json:"layouts"`
	Libraries      []DLibrary            `json:"libraries"`
}

type DConfig struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type DComponent struct {
	Name            string            `json:"name"`
	ChildComponents []DComponent      `json:"childComponents"`
	Configs         map[string]string `json:"configs"`
}

type DSee struct {
	ComponentName string `json:"componentName"`
	Data          string `json:"data"`
}

type DDo struct {
	UIEvent       string `json:"uiEvent"`
	ComponentName string `json:"componentName"`
	Data          string `json:"data"`
}

type DReact struct {
	SceneName          string `json:"sceneName"`
	ReactAction        string `json:"reactAction"`
	ReactComponentName string `json:"reactComponentName"`
	AnimateName        string `json:"animateName"`
	ReactComponentData string `json:"reactComponentData"`
}

type DInteraction struct {
	See   DSee     `json:"see"`
	Do    DDo      `json:"do"`
	React []DReact `json:"react"`
}

type DFlow struct {
	Interactions []DInteraction `json:"interactions"`
	FlowName     string         `json:"flowName"`
}

func CreateDComponent(componentName string) *DComponent {
	return &DComponent{componentName, nil, nil}
}

type DLayout struct {
	LayoutName string       `json:"layoutName"`
	LayoutRows []DLayoutRow `json:"layoutRows"`
}

type DLayoutRow struct {
	ComponentName     string `json:"componentName"`
	LayoutInformation string `json:"layoutInformation"`
	NormalInformation string `json:"normalInformation"`
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

type DProperty struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type LibraryPreset struct {
	Key           string       `json:"key"`
	Value         string       `json:"value"`
	PresetCalls   []PresetCall `json:"presetCalls"`
	SubProperties []DProperty  `json:"subProperties"`
}

type PresetCall struct {
	LibraryName   string `json:"libraryName"`
	LibraryPreset string `json:"libraryPreset"`
}

type DLibrary struct {
	LibraryName    string          `json:"libraryName"`
	LibraryPresets []LibraryPreset `json:"libraryPresets"`
}
