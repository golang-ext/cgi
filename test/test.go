package main

import (
	"github.com/golang-ext/cgi"
	"log"
)

func main() {
	if err := cgi.Init("/cgi").Run(":8080", cgi.New("D:/workspace/go/src/thelark.cn/cgi")); err != nil {
		log.Println(err)
	}
}
