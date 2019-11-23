package parser

import (
	. "../languages/design"
)

func NewDesignAppListener() *DesignAppListener {
	return &DesignAppListener{}
}

type DesignAppListener struct {
	BaseDesignListener
}


func (s *DesignAppListener) EnterConfigDecalartion(ctx *ConfigDecalartionContext) {
	//fmt.Println(ctx.ConfigKey().GetText(), ctx.ConfigValue().GetText())
}

func (s *DesignAppListener) EnterComponentName(ctx *ComponentNameContext) {
	//fmt.Println(ctx)
}

func (s *DesignAppListener) EnterLayoutDecalaration(ctx *LayoutDecalarationContext) {
	//fmt.Println(ctx.GetText())
}

func (s *DesignAppListener) EnterLayoutRow(ctx *LayoutRowContext) {
	//fmt.Println(ctx.GetText())
}

func (s *DesignAppListener) EnterComponentUseDeclaration(ctx *ComponentUseDeclarationContext) {
	//fmt.Println(ctx.GetText())
}