package cgi

import (
	"net/http"
	"fmt"
	"path"
	"encoding/json"
)

type Server struct {
	reqFuncMap map[string]func(w http.ResponseWriter, req *http.Request)
	CGIPath    string
}

func (s *Server) Add(path string, handle func(w http.ResponseWriter, req *http.Request)) {
	s.reqFuncMap[path] = handle
}

func Init(cgiPath string) *Server {
	return &Server{CGIPath: cgiPath, reqFuncMap: make(map[string]func(w http.ResponseWriter, req *http.Request))}
}
func (s *Server) Run(addr string, cgi *CGI) error {
	http.HandleFunc(s.CGIPath, CGIHandler(cgi))
	for path, handle := range s.reqFuncMap {
		http.HandleFunc(path, handle)
	}
	return http.ListenAndServe(addr, nil)
}
func CGIHandler(cgi *CGI) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		dir := cgi.runPath
		cgi.Path = cgi.goExe
		cgi.Dir = dir
		cgi.Env = append(cgi.Env, fmt.Sprintf("GOROOT=%s", cgi.goRoot))
		cgi.Env = append(cgi.Env, fmt.Sprintf("GOPATH=%s", cgi.goPath))
		scriptPath := dir + path.Base(req.URL.Path) // 设置 CGI 可执行文件的工作目录
		context := make(map[string]interface{})
		context["params"] = req.URL.Query()
		jsonBytes, _ := json.Marshal(context)
		args := append([]string{"run", scriptPath}, string(jsonBytes))
		cgi.Args = args
		cgi.ServeHTTP(w, req) // 启用 http server 重新实现 http Handler interfac{}
	}
}
