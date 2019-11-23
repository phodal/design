package parser

import (
	. "../languages/design"
	"fmt"
)

func NewDesignAppListener() *DesignAppListener {
	return &DesignAppListener{}
}

type DesignAppListener struct {
	BaseDesignListener
}


func (s *DesignAppListener) EnterConfigDecalartion(ctx *ConfigDecalartionContext) {
	fmt.Println(ctx.ConfigKey().GetText(), ctx.ConfigValue().GetText())
}