package update

import (
	"encoding/json"
	"github.com/zuiwuchang/king-go/os/fileperm"
	"io/ioutil"
)

// Update 檔案 更新 規則
type Update struct {
	// Truch 更新
	Truch map[string]File
	// Create 更新
	Create map[string]File
}
type updateJSON struct {
	Truch  []File
	Create []File
}

// NewUpdate ...
func NewUpdate() *Update {
	return &Update{
		Truch:  make(map[string]File),
		Create: make(map[string]File),
	}
}

// Empty 返回 更新規則是否為空
func (u *Update) Empty() bool {
	return len(u.Truch) == 0 && len(u.Create) == 0
}

// MarshalFile 序列化到 json 檔案
func (u *Update) MarshalFile(filename string) (e error) {
	var uj updateJSON
	if len(u.Truch) > 0 {
		i := 0
		uj.Truch = make([]File, len(u.Truch))
		for _, node := range u.Truch {
			uj.Truch[i] = node
			i++
		}
	}
	if len(u.Create) > 0 {
		i := 0
		uj.Create = make([]File, len(u.Create))
		for _, node := range u.Create {
			uj.Create[i] = node
			i++
		}
	}

	var b []byte
	b, e = json.MarshalIndent(&uj, "", "\t")
	if e != nil {
		return
	}
	e = ioutil.WriteFile(filename, b, fileperm.File)
	return
}
