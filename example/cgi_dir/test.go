package main

import (
	"fmt"
	"html/template"
	"os"
	"github.com/golang-ext/cgi/utils/method"
)

func main() {
	fmt.Println(method.QueryParam("a"))
	fmt.Println("This is test gocgi test")
	okok()
}

func okok() {
	tmpl, _ := template.New("test").Parse("<br/>Hello {{.}}")
	err := tmpl.Execute(os.Stdout, "thelark")
	fmt.Println(err)
}
