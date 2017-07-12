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

const (
	LOG_LEVEL_DEBUG = 1 << iota
	LOG_LEVEL_WARN
	LOG_LEVEL_ERROR
	LOG_LEVEL_FATAL
	LOG_LEVEL_INFO

	LOG_MODE_CONSOLE = 1 << iota
	LOG_MODE_LOGGER

	LOG_PREFIX_FILE = 1 << iota
	LOG_PREFIX_LINE
	LOG_PREFIX_TIME
)

var insConfig *Config
var insConsole *Console
var insLogger *Logger

/************************************************************************/
// export functions.
/************************************************************************/

/************************************************************************/
// moudule functions.
/************************************************************************/

func init() {
	insConfig = new(Config)
	insConfig.LogLevel = LOG_LEVEL_DEBUG |
		LOG_LEVEL_WARN |
		LOG_LEVEL_ERROR |
		LOG_LEVEL_FATAL |
		LOG_LEVEL_INFO

	insConfig.LogMode = LOG_MODE_CONSOLE | LOG_MODE_LOGGER
	insConfig.LogPrefix = LOG_PREFIX_FILE |
		LOG_PREFIX_LINE |
		LOG_PREFIX_TIME

	insConsole = new(Console)
	insLogger = new(Logger)
}

/************************************************************************/
// unit tests.
/************************************************************************/
