package main

import (
	"github.com/phodal/design/parser"
)

func main()  {
	//cmd.Execute()
	app := NewDesignApp()
	commentStruct := app.Start("examples/demo.design")
}