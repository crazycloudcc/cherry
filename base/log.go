/*
 * logger.
 * author: CC
 * email : crazycloudcc@gmail.com
 * date  : 2017.06.17
 */
package base

import (
	"fmt"
	"time"
	// "runtime"
)

/************************************************************************/
// constants, variables, structs, interfaces.
/************************************************************************/

const (
	LOG_LEVEL_DEBUG = iota
	LOG_LEVEL_INFO
	LOG_LEVEL_WARN
	LOG_LEVEL_ERROR
	LOG_LEVEL_FATAL
)

const (
	FLAG_NONE = 1 << iota
	FLAG_FILE
	FILE_LINE
	FILE_TIME
)

var logLevel int = LOG_LEVEL_DEBUG
var logFlag int = FLAG_NONE

// var doLog func(prefix string, args ...interface{})

/************************************************************************/
// 模块对外接口.
/************************************************************************/

// 设置Log输出等级.
func SetLogLevel(level int) {
	// log.SetFlags(log.Lshortfile | log.LstdFlags)
	logLevel = level
}

// 设置Log输出属性.
func SetLogFlag(flag int) {
	logFlag = flag
}

func LogDebug(args ...interface{}) {
	if logLevel <= LOG_LEVEL_DEBUG {
		doLog("[DEBUG]", args...)
	}
}

func LogInfo(args ...interface{}) {
	if logLevel <= LOG_LEVEL_INFO {
		doLog("\033[32;1m[INFO]\033[0m", args...)
	}
}

func LogWarn(args ...interface{}) {
	if logLevel <= LOG_LEVEL_WARN {
		doLog("\033[33;1m[WARN]\033[0m", args...)
	}
}

func LogError(args ...interface{}) {
	if logLevel <= LOG_LEVEL_ERROR {
		doLog("\033[31;1m[ERROR]\033[0m", args...)
	}
}

func LogFatal(args ...interface{}) {
	if logLevel <= LOG_LEVEL_FATAL {
		doLog("\033[31;1m[FATAL]\033[0m", args...)
		panic(nil)
	}
}

/************************************************************************/
// 模块内功能实现
/************************************************************************/

// // 显示日期, 文件名, 行数.
// func doLog(prefix string, args ...interface{}) {
// 	_, file, line, _ := runtime.Caller(2) // depth
// 	// log.Println(file, line, args)
// 	fmt.Println(prefix, time.Now().Format("2006/01/02 15:04:05.000"), file, line, args)
// }

func doLog(prefix string, args ...interface{}) {
	fmt.Println(time.Now().Format("2006/01/02 15:04:05.000"), prefix, args)
}
