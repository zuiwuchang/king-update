package update

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/go-jsonnet"
	"github.com/zuiwuchang/king-go/strings"
	"io/ioutil"
	"regexp"
	gStrings "strings"
)

// ErrorNoRule 沒有定義 任何 更新規則
var ErrorNoRule = errors.New("not found any rules")

// Rule 一個 規則
type Rule struct {
	Regexp *regexp.Regexp
	Path   string
	Mode   int
}

// Match 返回 傳入字符串 是否和 當前規則 匹配
func (r *Rule) Match(str string) (yes bool) {
	if r.Regexp == nil {
		if r.Path == str {
			yes = true
			return
		}
	} else {
		yes = r.Regexp.MatchString(str)
	}
	return
}

type ruleJSON struct {
	Regexp string
	Path   string
	Mode   int
}

// UnmarshalRules 解析包 升級配置到 go 結構
func UnmarshalRules(b []byte) ([]Rule, error) {
	return UnmarshalRulesString(strings.BytesToString(b))
}

// UnmarshalRulesFile 解析包 升級配置到 go 結構
func UnmarshalRulesFile(filename string) ([]Rule, error) {
	b, e := ioutil.ReadFile(filename)
	if e != nil {
		return nil, e
	}
	return UnmarshalRulesString(strings.BytesToString(b))
}

// UnmarshalRulesString 解析包 升級配置到 go 結構
func UnmarshalRulesString(str string) (rules []Rule, e error) {
	vm := jsonnet.MakeVM()
	str, e = vm.EvaluateSnippet("", str)
	if e != nil {
		return
	}
	b := strings.StringToBytes(str)
	var rulesJSON []ruleJSON
	e = json.Unmarshal(b, &rulesJSON)
	if e != nil {
		return
	}

	n := len(rulesJSON)
	if n == 0 {
		e = ErrorNoRule
		return
	}
	rules = make([]Rule, n)
	for i, node := range rulesJSON {
		if node.Mode < UpdateModeIgnore || node.Mode > UpdateModeCreate {
			e = fmt.Errorf("unknow rule mode : %v", node)
			rules = nil
			return
		}
		rules[i].Mode = node.Mode

		path := gStrings.TrimSpace(node.Path)
		if path == "" {
			rules[i].Regexp, e = regexp.Compile(node.Regexp)
			if e != nil {
				rules = nil
				return
			}
		} else {
			rules[i].Path = path
		}
	}
	return
}
