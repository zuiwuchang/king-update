package configure

import (
	"fmt"
)

//Log ...
type Log struct {
	// 需要打印的 日誌等級
	// "TRACE"
	// "DEBUG"
	// "INFO"
	// "WARN"
	// "ERROR"
	// "FAULT"
	Logs []string

	//是否 顯示 短檔案名
	Short bool
	//是否 顯示 長檔案名
	Long bool
}

func (l *Log) show() {
	fmt.Println("\tLog : {")
	fmt.Println("\t\tLogs :", l.Logs, ",")
	fmt.Println("\t\tShort: ", l.Short, ",")
	fmt.Println("\t\tLong :", l.Long, ",")
	fmt.Println("\t},")
}
