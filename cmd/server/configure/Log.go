package configure

// Log 運行 日誌
type Log struct {
	// 需要打印的 日誌
	// "TRACE",
	// "DEBUG",
	// "INFO",
	// "WARN",
	// "ERROR",
	// "FAULT",
	Logs []string
	//是否 顯示 短 檔案名
	Short bool
	//是否 顯示 長 檔案名
	Long bool
}
