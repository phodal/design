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

func (s *DesignAppListener) EnterExpress(ctx *ExpressContext) {
	fmt.Println(ctx.GetText())
}