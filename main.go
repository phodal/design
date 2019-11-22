package main

import (
	. "./parser"
)

func main()  {
	//cmd.Execute()
	app := NewDesignApp()
	app.Start("examples/demo.design")
}
