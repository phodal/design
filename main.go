package main

import (
	. "./parser"
	"io/ioutil"
	"os"
)

func main()  {
	//cmd.Execute()
	app := NewDesignApp()
	bytes := app.Start("examples/login.design")

	_ = ioutil.WriteFile("output.json", bytes, 0644)

	_, _ = os.Stdout.Write(bytes)
}
