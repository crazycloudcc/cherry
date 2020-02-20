package nettcp

/*
 * encode/decode msg data.
 * author: CC
 * email : 151503324@qq.com
 * date  : 2017.06.17
 */

import "github.com/crazycloudcc/cherry/base"

/************************************************************************/
// constants, variables, structs, interfaces.
/************************************************************************/

// Msg struct.
type Msg struct {
	Len      uint16
	Crypt    byte
	Zip      byte
	ID       uint32
	MetaData []byte
}

/************************************************************************/
// export functions.
/************************************************************************/

/************************************************************************/
// moudule functions.
/************************************************************************/

// encode data.
func msgToBytes(msg *Msg) []byte {
	ret := make([]byte, MsgHeadLen+int(msg.Len))
	len := ret[:MsgLenIndexE]
	base.Uint16ToByte(len, uint16(msg.Len))
	ret[MsgCryptFlagIndex] = msg.Crypt
	ret[MsgZipFlagIndex] = msg.Zip
	msgid := ret[MsgMsgIDIndexS:MsgMsgIDIndexE]
	base.Uint32ToByte(msgid, msg.ID)
	copy(ret[MsgHeadLen:], msg.MetaData)
	return ret
}

// decode data.
func bytesToMsg(buf []byte) *Msg {
	ret := new(Msg)
	ret.Len = base.ByteToUint16(buf[:MsgLenIndexE])
	ret.Crypt = buf[MsgCryptFlagIndex]
	ret.Zip = buf[MsgZipFlagIndex]
	ret.ID = base.ByteToUint32(buf[MsgMsgIDIndexS:MsgMsgIDIndexE])
	if ret.Len > 0 {
		ret.MetaData = make([]byte, ret.Len)
		copy(ret.MetaData, buf[MsgHeadLen:])
	}
	return ret
}

/************************************************************************/
// unit tests.
/************************************************************************/
