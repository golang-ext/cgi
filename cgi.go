package cgi

import (
	"os"
	"net/http/cgi"
)

var (
	GOROOT = os.Getenv("GOROOT")
	GOPATH = os.Getenv("GOPATH")
	GOEXE  = GOROOT + "/bin/go"
)

type CGI struct {
	runPath string // eg."D:/workspace/go/src/thelark.cn/cgi" 运行路径
	goExe   string // eg."E:/Go/bin/go" 脚本路径
	goRoot  string
	goPath  string
	*cgi.Handler
}

func New(runPath string) *CGI {
	return &CGI{runPath: runPath, goExe: GOEXE, goRoot: GOROOT, goPath: GOPATH, Handler: new(cgi.Handler)}
}
func (c *CGI) SetGoRoot(goRoot string) {
	c.goRoot = goRoot
	c.goExe = goRoot + "/bin/go"
}
func (c *CGI) GetGoRoot() string {
	return c.goRoot
}
func (c *CGI) SetGoPath(goPath string) {
	c.goPath = goPath
}
func (c *CGI) GetGoPath() string {
	return c.goPath
}
func (c *CGI) SetGoExe(goExe string) {
	c.goExe = goExe
}
func (c *CGI) GetGoExe() string {
	return c.goExe
}
