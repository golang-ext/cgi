package main

import (
	"github.com/golang-ext/cgi"
	"log"
	"fmt"
)

func main() {
	fmt.Println("Server Running :8080")
	if err := cgi.Init("/cgi/").Run(":8080", cgi.New("D:/work/src/github.com/golang-ext/cgi/example/cgi_dir/")); err != nil {
		log.Println(err)
	}
}
