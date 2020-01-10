package netwebsocket

/*
 * network config.
 * author: CC
 * email : 151503324@qq.com
 * date  : 2017.06.17
 */

import (
	"fmt"
)

/************************************************************************/
// constants, variables, structs, interfaces.
/************************************************************************/

// Config struct.
type Config struct {
	Flag int32
	Type string
	Host string
	Port int32
	Max  int32
}

/************************************************************************/
// export functions.
/************************************************************************/

func (owner *Config) String() string {
	ret := fmt.Sprintf("Type[%s], Host[%s], Port[%d], Max[%d]", owner.Type, owner.Host, owner.Port, owner.Max)
	return ret
}

/************************************************************************/
// moudule functions.
/************************************************************************/

/************************************************************************/
// unit tests.
/************************************************************************/
