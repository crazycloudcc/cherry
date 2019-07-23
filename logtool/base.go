package logtool

import (
	"os"
)

/************************************************************************/
// 常量, 变量, 结构体, 接口定义.
/************************************************************************/

const (
	LOG_LEVEL_DEBUG int32 = iota
	LOG_LEVEL_INFO
	LOG_LEVEL_WARN
	LOG_LEVEL_ERROR
	LOG_LEVEL_FATAL
)

//运行时参数
type Logger struct {
	conf      Config
	writeSize int32
	fileName  string
	file      *os.File
	saveLog   bool
	hour      int32
	logLevel  int32
}
