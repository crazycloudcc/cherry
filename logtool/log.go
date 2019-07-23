package logtool

import (
	"errors"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"time"
)

func newLogger(conf Config) *Logger {
	m := new(Logger)
	m.conf = conf
	return m
}

//log初始化设置，包括logpath和logsize的判断以及新建log日志保存路径
func (this *Logger) SetLog() {
	this.logLevel = this.IntLogLevel()

	if this.conf.Path == "" {
		this.saveLog = false
		return
	} else if this.conf.Size == 0 {
		this.saveLog = false
		err := errors.New("file size is nil")
		fmt.Println(err)
		return
	} else {
		//创建文件夹
		_, err := os.Stat(this.conf.Path)
		if err != nil {
			err = os.Mkdir(this.conf.Path, os.ModePerm)
			if err != nil {
				fmt.Println("create path failed:", err)
				this.saveLog = false
				return
			}
		}
		//创建以时间戳命名的文件
		this.hour = int32(time.Now().Hour())
		creatTime := time.Now().Format("2006-01-02-15")
		this.fileName = creatTime + "-" + strconv.Itoa(i) + ".txt"
		this.file, err = os.Create(this.conf.Path + this.fileName)
		if err != nil {
			fmt.Println("create file faild:", err)
			this.saveLog = false
			return
		}
		this.saveLog = true
	}
}

//DEBUG
func (this *Logger) LogDebug(args ...interface{}) {
	if this.logLevel <= LOG_LEVEL_DEBUG {
		this.doLog("[DEBUG]", args...)
	}
}

//INFO
func (this *Logger) LogInfo(args ...interface{}) {
	if this.logLevel <= LOG_LEVEL_INFO {
		this.doLog("\033[32;1m[INFO]\033[0m", args...)
	}
}

//WARN
func (this *Logger) LogWarn(args ...interface{}) {
	if this.logLevel <= LOG_LEVEL_WARN {
		this.doLog("\033[33;1m[WARN]\033[0m", args...)
	}
}

//ERROR
func (this *Logger) LogError(args ...interface{}) {
	if this.logLevel <= LOG_LEVEL_ERROR {
		this.doLog("\033[31;1m[ERROR]\033[0m", args...)
	}
}

//FATEL
func (this *Logger) LogFatal(args ...interface{}) {
	if this.logLevel <= LOG_LEVEL_FATAL {
		this.doLog("\033[31;1m[FATAL]\033[0m", args...)
	}
}

//输入打印log日志并且判断是否需要保存log日志
func (this *Logger) doLog(prefix string, args ...interface{}) {
	fmt.Println(prefix, args)
	_, file, line, _ := runtime.Caller(2) // depth
	if this.saveLog {
		this.save(fmt.Sprintln(prefix, file, line, args))
	}
}

//将配置文件中的LogLevel字符串转为int类型
func (this *Logger) IntLogLevel() int32 {
	// fmt.Println(this.conf.Level)
	if this.conf.Level == "LOG_LEVEL_DEBUG" {
		return 0
	} else if this.conf.Level == "LOG_LEVEL_INFO" {
		return 1
	} else if this.conf.Level == "LOG_LEVEL_WARN" {
		return 2
	} else if this.conf.Level == "LOG_LEVEL_ERROR" {
		return 3
	} else if this.conf.Level == "LOG_LEVEL_FATAL" {
		return 4
	} else {
		fmt.Println(errors.New("log.go - input log_level error"))
		return 0
	}
}
