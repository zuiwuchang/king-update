package configure

import (
	"encoding/json"
	//"fmt"
	"github.com/google/go-jsonnet"
	kStrings "github.com/zuiwuchang/king-go/strings"
	"io/ioutil"
)

// Configure ...
type Configure struct {
	//工作地址
	LAddr string

	//日誌
	Log Log
}

// String ...
func (c *Configure) String() string {
	b, _ := json.MarshalIndent(c, "", "\t")
	return string(b)
}

var gCnf Configure

// GetConfigure 返回 全局的 唯一 配置 檔
func GetConfigure() *Configure {
	return &gCnf
}

// Init 初始化 配置 到 go 結構
func Init(filename string) (e error) {
	//read file
	var b []byte
	b, e = ioutil.ReadFile(filename)
	if e != nil {
		return
	}

	//jsonnet
	var strJSON string
	vm := jsonnet.MakeVM()
	strJSON, e = vm.EvaluateSnippet("", kStrings.BytesToString(b))
	if e != nil {
		return
	}

	//json
	e = json.Unmarshal(kStrings.StringToBytes(strJSON), &gCnf)
	if e != nil {
		return
	}

	/*g_cnf.Timeout *= time.Second

	//格式化 參數
	g_cnf.Http2.format()
	g_cnf.DB.format()
	g_cnf.Path.format()
	g_cnf.Request.format()
	*/
	return
}
