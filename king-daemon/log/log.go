package log

import (
	klog "github.com/zuiwuchang/king-go/log"
	"github.com/zuiwuchang/king-update/king-daemon/configure"
	"log"
)

// Loggers ...
var Loggers = klog.NewDebugLoggers2("[ku-daemon] ")

// Trace ...
var Trace *log.Logger

// Debug ...
var Debug *log.Logger

// Info ...
var Info *log.Logger

// Warn ...
var Warn *log.Logger

// Error ...
var Error *log.Logger

// Fault ...
var Fault *log.Logger

// Init 初始化 日誌
func Init() {
	logers := Loggers

	cnf := configure.GetConfigure()
	flag := log.LstdFlags
	if cnf.Log.Short {
		flag |= log.Lshortfile
	} else if cnf.Log.Long {
		flag |= log.Llongfile
	}

	logs := make(map[string]bool)
	for _, lv := range cnf.Log.Logs {
		logs[lv] = true
	}
	if _, ok := logs["TRACE"]; ok {
		Trace = logers.Trace
		Trace.SetFlags(flag)

	}

	if _, ok := logs["DEBUG"]; ok {
		Debug = logers.Debug
		Debug.SetFlags(flag)
	}

	if _, ok := logs["INFO"]; ok {
		Info = logers.Info
		Info.SetFlags(flag)
	}

	if _, ok := logs["WARN"]; ok {
		Warn = logers.Warn
		Warn.SetFlags(flag)
	}

	if _, ok := logs["ERROR"]; ok {
		Error = logers.Error
		Error.SetFlags(flag)
	}

	if _, ok := logs["FAULT"]; ok {
		Fault = logers.Fault
		Fault.SetFlags(flag)

	}
}
