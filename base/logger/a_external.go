/*
 * [file desc]
 * author: CC
 * email : crazycloudcc@gmail.com
 * date  : 2017.06.17
 */
package logger

/************************************************************************/
// constants, variables, structs, interfaces.
/************************************************************************/

/************************************************************************/
// export functions.
/************************************************************************/

func LoadConfig(conf *Config) {
	insConfig.LogLevel = conf.LogLevel
	insConfig.LogMode = conf.LogMode
	insConfig.LogPrefix = conf.LogPrefix
}

func Debug(args ...interface{}) {
	insConsole.debug(args...)
	insLogger.debug(args...)
}

func Warn(args ...interface{}) {
	insConsole.warn(args...)
	insLogger.warn(args...)
}

func Error(args ...interface{}) {
	insConsole.error(args...)
	insLogger.error(args...)
}

func Fatal(args ...interface{}) {
	insConsole.fatal(args...)
	insLogger.fatal(args...)
}

func Info(args ...interface{}) {
	insConsole.info(args...)
	insLogger.info(args...)
}

/************************************************************************/
// moudule functions.
/************************************************************************/

/************************************************************************/
// unit tests.
/************************************************************************/
