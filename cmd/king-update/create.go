package main

import (
	"fmt"
	"github.com/zuiwuchang/king-update/king-update"
	"os"
	"path/filepath"
)

const (
	// UpdateFile 檔案升級 配置檔
	UpdateFile = "update.json"
)

// Create 創建 檔案升級 配置檔
func Create() {
	rules, e := update.UnmarshalRulesFile(RuleFile)
	if e != nil {
		fmt.Println(e)
		return
	}
	uPkg := update.NewUpdate()
	filepath.Walk(
		".",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				fmt.Println(err)
				return err
			}
			if info.IsDir() {
				return nil
			}

			insertUpdate(uPkg, rules, path)
			return nil
		},
	)
	if uPkg.Empty() {
		fmt.Println("no found any update files")
		os.Exit(1)
		return
	}

	e = uPkg.MarshalFile(UpdateFile)
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}

}
func insertUpdate(uPkg *update.Update, rules []update.Rule, path string) {
	var e error
	var ok bool
	for _, rule := range rules {
		if ok, e = insertUpdateRule(uPkg, rule, path); e != nil {
			fmt.Println(e)
			os.Exit(1)
		} else if ok {
			break
		}
	}
}
func insertUpdateRule(uPkg *update.Update, rule update.Rule, path string) (ok bool, e error) {
	//驗證 規則 是否 匹配
	if !rule.Match(path) {
		return
	}
	ok = true

	//要忽略的 檔案
	if rule.Mode == update.UpdateModeIgnore {
		return
	}

	var keys map[string]update.File
	if rule.Mode == update.UpdateModeTrunc {
		keys = uPkg.Truch
	} else {
		keys = uPkg.Create
	}

	var file update.File
	file, e = update.NewFile(path)
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}

	keys[path] = file
	return
}
