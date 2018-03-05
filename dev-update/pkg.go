package update

import (
	"encoding/json"
	"github.com/google/go-jsonnet"
	"github.com/zuiwuchang/king-go/strings"
	"io/ioutil"
)

// Package 一個 軟件包 包含了 軟件包各種信息 更新規則
type Package struct {
	//軟件包 名稱 作為區分軟件包的標識
	Name string
	//軟件包 版本號 使用 utf8 字符串 比較 較小的作為上個版本
	Version string

	//軟件 官網
	Web string

	//功能描述
	Description string
	//相對上個版本的 的修改描述
	Modify string
	//如果為 true 則 不會從上個 版本 自動更新
	DisableAuto bool
	//程式作者
	Author string
	//聯繫作者
	AuthorEmail string
	//當前版本 維護人員
	Maintainer string
	//聯繫 維護人員
	MaintainerEmail string
}

// UnmarshalPackage 解析包 配置到 go 結構
func UnmarshalPackage(b []byte) (*Package, error) {
	return UnmarshalPackageString(strings.BytesToString(b))
}

// UnmarshalPackageFile 解析包 配置到 go 結構
func UnmarshalPackageFile(filename string) (*Package, error) {
	b, e := ioutil.ReadFile(filename)
	if e != nil {
		return nil, e
	}
	return UnmarshalPackageString(strings.BytesToString(b))
}

// UnmarshalPackageString 解析包 配置到 go 結構
func UnmarshalPackageString(str string) (pkg *Package, e error) {
	vm := jsonnet.MakeVM()
	str, e = vm.EvaluateSnippet("", str)
	if e != nil {
		return
	}
	b := strings.StringToBytes(str)
	var p Package
	e = json.Unmarshal(b, &p)
	if e != nil {
		return
	}
	pkg = &p
	return
}
