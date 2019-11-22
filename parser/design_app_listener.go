package parser

import (
	. "../languages/design"
	"fmt"
	"reflect"
)

func NewDesignAppListener() *DesignAppListener {
	return &DesignAppListener{}
}

type DesignAppListener struct {
	BaseDesignListener
}

func (s *DesignAppListener) EnterExpress(ctx *ExpressContext) {
	//fmt.Println(ctx.GetText())
}

func (s *DesignAppListener) EnterDesignSystemDeclaration(ctx *DesignSystemDeclarationContext) {
	fmt.Println(ctx.IDENTIFIER())
}

func (s *DesignAppListener) EnterComponentBlockDeclaration(ctx *ComponentBlockDeclarationContext) {
	body := ctx.ComponentBodyDeclaration().GetChild(0)
	if body != nil {
		childComponent := body.GetChild(0)
		if reflect.TypeOf(childComponent).String() == "*parser.UseExpressContext" {
			context := childComponent.(*UseExpressContext)
			fmt.Println(context.ExpressKey().GetText(), context.ExpressValue().GetText())
		}
	}
}


func (s *DesignAppListener) EnterCodeBlockDeclaration(ctx *CodeBlockDeclarationContext) {
	fmt.Println(ctx.AllHTML_STRING())
}