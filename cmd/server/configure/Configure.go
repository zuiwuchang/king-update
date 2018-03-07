package configure

import (
	"encoding/json"
	"github.com/google/go-jsonnet"
	kStrings "github.com/zuiwuchang/king-go/strings"
	"io/ioutil"
	"strings"
)

// Configure 配置
type Configure struct {
	// 服務器 工作地址
	Addr string
	//ssl證書 位置
	Crt string
	//ssl證書key 位置
	Key string

	Users []User

	//軟件包 存放 位置
	Folders []string

	//運行 日誌
	Log Log
}

func (c *Configure) format() {
	c.Addr = strings.TrimSpace(c.Addr)
	c.Crt = strings.TrimSpace(c.Crt)
	c.Key = strings.TrimSpace(c.Key)
}
func (c *Configure) String() string {
	b, _ := json.MarshalIndent(c, "", "\t")
	return kStrings.BytesToString(b)
}

var gCnf Configure

// GetConfigure 返回 配置
func GetConfigure() *Configure {
	return &gCnf
}

// Init 初始化 配置
func Init(filename string) (e error) {
	var b []byte
	b, e = ioutil.ReadFile(filename)
	if e != nil {
		return
	}

	var str string
	var vm = jsonnet.MakeVM()
	str, e = vm.EvaluateSnippet("", kStrings.BytesToString(b))
	if e != nil {
		return
	}
	b = kStrings.StringToBytes(str)
	e = json.Unmarshal(b, &gCnf)
	if e != nil {
		return
	}

	gCnf.format()
	return
}
