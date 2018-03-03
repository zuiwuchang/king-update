package update

import (
	"fmt"
	"os"
)

const (
	// UpdateModeIgnore 忽略 檔案升級
	UpdateModeIgnore = iota
	// UpdateModeTrunc 覆蓋檔案
	UpdateModeTrunc
	// UpdateModeCreate 不存在時 才創建
	UpdateModeCreate
)

// File 需要升級的 檔案
type File struct {
	//檔案 相對 安裝位置的 路徑
	Path string
	//檔案 hash 值
	Hash string
	//檔案 posix 權限
	Mode os.FileMode
}

// NewFile 傳入檔案路徑 創建檔案信息
func NewFile(filepath string) (f File, e error) {
	var file *os.File
	file, e = os.Open(filepath)
	if e != nil {
		return
	}
	defer file.Close()
	var info os.FileInfo
	info, e = file.Stat()
	if e != nil {
		return
	}
	if info.IsDir() {
		e = fmt.Errorf("%s is not a file", filepath)
		return
	}

	var hash string
	if hash, e = HashReader(file); e != nil {
		return
	}

	f = File{
		Path: filepath,
		Mode: info.Mode(),
		Hash: hash,
	}
	return
}
