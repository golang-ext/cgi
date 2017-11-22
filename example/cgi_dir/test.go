package main

import (
	"fmt"
	"html/template"
	"os"
)

func init() {
	//设置header 头 记得要 加 两个\n
	//否则访问 会有问题  cgi: no headers 错误
	//如果一个 \n 会将结果输出到 终端 并提示cgi: bogus header line: This is  gocgi test
	fmt.Print("Content-Type: text/html;charset=utf-8\n\n")
}

func main() {
	fmt.Println(os.Args)
	fmt.Println("This is test gocgi test")
	okok()
}

func okok() {
	tmpl, _ := template.New("test").Parse("<br/>Hello {{.}}")
	err := tmpl.Execute(os.Stdout, "thelark")
	fmt.Println(err)
}
