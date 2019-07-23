package logtool

var InsLog *Logger

//初始化Logger，获取log的配置文件相关值
func LogInit(conf Config) {
	InsLog = newLogger(conf)
	InsLog.SetLog()
}
