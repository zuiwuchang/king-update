package main

import (
	"encoding/json"
	"fmt"
	"github.com/zuiwuchang/king-update/king-update"
	"os"
	"time"
)

const (
	// TemplateFile 包模板 檔案
	TemplateFile = "../.default.jsonnet"
	// PackageFile 包配置 檔案
	PackageFile = "package.jsonnet"
	// RuleFile 升級規則 定義
	RuleFile = "rules.jsonnet"
)

// Init 依據 package.json 配置檔案 初始化 軟件包
func Init() {
	createPackageFile()
	createfiles()
}
func createfiles() {
	f, e := os.Create(RuleFile)
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
	defer f.Close()
	_, e = f.WriteString(`// 在此檔案中 定義 升級規則
// 每條 規則 包含一個 Regexp(正則定義的 匹配規則)/Path(檔案相對路徑) 和一個 Mode(升級模式)
// 規則 優先級 安出現的 先後順序

`)
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
	_, e = f.WriteString(
		fmt.Sprintf("// 檔案升級\nlocal modeIgnore = %v;\n// 覆蓋檔案\nlocal modeTrunc = %v;\n//不存在時 才創建\nlocal modeCreate = %v;",
			update.UpdateModeIgnore,
			update.UpdateModeTrunc,
			update.UpdateModeCreate,
		),
	)
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
	_, e = f.WriteString(`

//定義規則
[
	//忽略 所有測試 檔案 *.test.*
	{
		Regexp:"\\.test\\.",
		Mode:modeIgnore,
	},
	{
		//如果已經存在 配置檔案 則 不創建
		Path:"my.jsonnet",
		Mode:modeCreate,
	},
	{
		//將剩下的所有 檔案都 覆蓋舊版本
		Mode:modeTrunc,
	},
]`)
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
}
func createPackageFile() {
	pkg, e := update.UnmarshalPackageFile(TemplateFile)
	if e != nil {
		pkg = &update.Package{}
	}

	f, e := os.Create(PackageFile)
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
	defer f.Close()
	pkg.Version = time.Now().Format("20060102150405")
	b, e := json.MarshalIndent(pkg, "", "\t")
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
	_, e = f.Write(b)
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
}
