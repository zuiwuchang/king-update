package tools

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/zuiwuchang/king-update/dev-update"
	"os"
	"path/filepath"
)

const (
	// UpdateFile 檔案升級 配置檔
	UpdateFile = "update.json"
)

type execCreate struct {
}

func (fCtx execCreate) Do() {
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

			fCtx.insertUpdate(uPkg, rules, path)
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
func (fCtx execCreate) insertUpdate(uPkg *update.Update, rules []update.Rule, path string) {
	var e error
	var ok bool
	for _, rule := range rules {
		if ok, e = fCtx.insertUpdateRule(uPkg, rule, path); e != nil {
			fmt.Println(e)
			os.Exit(1)
		} else if ok {
			break
		}
	}
}
func (execCreate) insertUpdateRule(uPkg *update.Update, rule update.Rule, path string) (ok bool, e error) {
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

func initCreate() (cmd *cobra.Command) {
	var exec execCreate
	cmd = &cobra.Command{
		Use:   "create",
		Short: "create package file  --  update.jsonnet",
		Long: `create package file  --  update.jsonnet
	update.jsonnet  --  auto update files and hash
`,
		Run: func(cmd *cobra.Command, args []string) {
			exec.Do()
			fmt.Println(`success : 
	create update.jsonnet`)
		},
	}
	return
}
