/*
 * [file desc]
 * author: CC
 * email : crazycloudcc@gmail.com
 * date  : 2017.06.17
 */
package logger

import (
	"fmt"
)

/************************************************************************/
// constants, variables, structs, interfaces.
/************************************************************************/

type Config struct {
	LogLevel  int
	LogMode   int
	LogPrefix int
}

/************************************************************************/
// export functions.
/************************************************************************/

func (this *Config) String() string {
	ret := fmt.Sprintf("logger.Config info: level[%d], mode[%d], prefix[%d].",
		this.LogLevel, this.LogMode, this.LogPrefix)
	return ret
}

/************************************************************************/
// moudule functions.
/************************************************************************/

/************************************************************************/
// unit tests.
/************************************************************************/
