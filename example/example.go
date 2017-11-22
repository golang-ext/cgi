package main

import (
	"github.com/golang-ext/cgi"
	"log"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Server Running :8080")
	server := cgi.Init("/cgi/")
	server.Add("/hello", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("Hello World"))
	})
	if err := server.Run(":8080", cgi.New("D:/work/src/github.com/golang-ext/cgi/example/cgi_dir/")); err != nil {
		log.Println(err)
	}
}
