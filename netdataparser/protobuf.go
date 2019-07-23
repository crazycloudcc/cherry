package netdataparser

/*
 * netdataparser protobuf data.
 * author: CC
 * email : 151503324@qq.com
 * date  : 2017.06.17
 */

import (
	"github.com/golang/protobuf/proto"
)

/************************************************************************/
// constants, variables, structs, interfaces.
/************************************************************************/

/************************************************************************/
// export functions.
/************************************************************************/

// Unmarshal TODO
func Unmarshal(buf []byte, pb proto.Message) error {
	return proto.Unmarshal(buf, pb)
}

// Marshal TODO
func Marshal(pb proto.Message) ([]byte, error) {
	return proto.Marshal(pb)
}

/************************************************************************/
// moudule functions.
/************************************************************************/

/************************************************************************/
// unit tests.
/************************************************************************/
