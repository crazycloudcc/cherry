package nettcp

/*
 * [file desc]
 * author: CC
 * email : 151503324@qq.com
 * date  : 2017.06.17
 */

import (
	"errors"

	"github.com/crazycloudcc/cherry/base"
)

/************************************************************************/
// constants, variables, structs, interfaces.
/************************************************************************/

var (
	errFooMsgHead   = errors.New("[nettcp.parser] msg head error")
	errFooMsgDecode = errors.New("[nettcp.parser] msg decode error")
	errFooMsgEncode = errors.New("[nettcp.parser] msg encode error")
)

/************************************************************************/
// export functions.
/************************************************************************/

// ReadMsgLen 获得消息长度.
// 参数: 1-长度为2字节的消息头信息.
// 返回: 1-消息长度, 2-错误信息.
func ReadMsgLen(buf []byte) (uint16, error) {
	if len(buf) != MsgLenIndexE {
		return 0, errFooMsgHead
	}
	msgLen := base.ByteToUint16(buf)
	return msgLen, nil
}

// Unmarshal 消息解码.
// 参数: 1-消息体二进制数据.(包含8字节的消息头)
// 返回: 1-格式化后的消息结构体, 2-错误信息.
func Unmarshal(buf []byte) (*Msg, error) {
	ret := bytesToMsg(buf)
	if ret.Zip == MsgZipFlag {
		// TODO.
	}
	if ret.Crypt == MsgCryptFlag {
		// TODO.
	}
	return ret, nil
}

// Marshal 消息编码.
// 参数: 1-消息结构体.
// 返回: 1-消息体二进制数据, 2-错误信息.
func Marshal(msg *Msg) ([]byte, error) {
	if msg.Crypt == MsgCryptFlag {
		// TODO.
	}
	if msg.Zip == MsgZipFlag {
		// TODO.
	}
	ret := msgToBytes(msg)
	return ret, nil
}

/************************************************************************/
// moudule functions.
/************************************************************************/

/************************************************************************/
// unit tests.
/************************************************************************/
