package cgi

import (
	"net/http"
	"fmt"
)

type Server struct {
	CGIPath string
}

func Init(cgiPath string) *Server {
	return &Server{CGIPath: cgiPath}
}
func (s *Server) Run(addr string, cgi *CGI) error {
	http.HandleFunc(s.CGIPath, s.CGIHandler(cgi))
	return http.ListenAndServe(addr, nil)
}
func (s *Server) CGIHandler(cgi *CGI) func(w http.ResponseWriter, req *http.Request) {
	dir := cgi.runPath
	cgi.Path = cgi.goExe
	cgi.Dir = dir
	cgi.Env = append(cgi.Env, fmt.Sprintf("GOROOT=%s", cgi.goRoot))
	return func(w http.ResponseWriter, req *http.Request) {
		scriptPath := dir + req.URL.Path // 设置 CGI 可执行文件的工作目录
		args := []string{"run", scriptPath}
		cgi.Args = append(cgi.Args, args...)
		cgi.ServeHTTP(w, req) // 启用 http server 重新实现 http Handler interfac{}
	}
}

