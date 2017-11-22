package method

import (
	"fmt"
	"os"
	"encoding/json"
)

var params = make(map[string][]string)

func init() {
	//设置header 头 记得要 加 两个\n
	//否则访问 会有问题  cgi: no headers 错误
	//如果一个 \n 会将结果输出到 终端 并提示cgi: bogus header line: This is  gocgi test
	fmt.Print("Content-Type: text/html;charset=utf-8\n\n")
	var context map[string]map[string][]string
	json.Unmarshal([]byte(os.Args[1]), &context)
	params = context["params"]

}
func QueryParam(name string) string {
	return params[name][0]
}
