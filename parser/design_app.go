package parser

import (
	. "../languages/design"
	"encoding/json"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func NewDesignApp() *DesignApp {
	return &DesignApp{}
}

type DesignApp struct {
}

func (j *DesignApp) Start(path string) []byte {
	context := (*DesignApp)(nil).ProcessFile(path).Start()
	listener := NewDesignAppListener()

	antlr.NewParseTreeWalker().Walk(listener, context)

	information := listener.getDesignInformation()

	content, err := json.Marshal(information)
	if err != nil {
		fmt.Println("error:", err)
	}

	return content
}

func (j *DesignApp) ProcessFile(path string) *DesignParser {
	is, _ := antlr.NewFileStream(path)
	lexer := NewDesignLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, 0);
	parser := NewDesignParser(stream)
	return parser
}
