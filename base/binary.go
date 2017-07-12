/*
 * binary convert functions.
 * author: CC
 * email : crazycloudcc@gmail.com
 * date  : 2017.06.17
 */
package base

import (
	"encoding/binary"
)

/************************************************************************/
// constants, variables, structs, interfaces.
/************************************************************************/

var littleEndian bool

/************************************************************************/
// 模块对外接口.
/************************************************************************/

// 设置小字端.
func SetLittleEndian() {
	littleEndian = true
}

// 设置大字端.
func SetBigEndian() {
	littleEndian = false
}

// 二进制写入uint16到[]byte.
func Uint16ToByte(dst []byte, v uint16) {
	if littleEndian {
		binary.LittleEndian.PutUint16(dst, v)
	} else {
		binary.BigEndian.PutUint16(dst, v)
	}
}

// 二进制写入uint32到[]byte.
func Uint32ToByte(dst []byte, v uint32) {
	if littleEndian {
		binary.LittleEndian.PutUint32(dst, v)
	} else {
		binary.BigEndian.PutUint32(dst, v)
	}
}

// 二进制转uint16.
func ByteToUint16(data []byte) uint16 {
	if littleEndian {
		return binary.LittleEndian.Uint16(data)
	} else {
		return binary.BigEndian.Uint16(data)
	}
}

// 二进制转uint32
func ByteToUint32(data []byte) uint32 {
	if littleEndian {
		return binary.LittleEndian.Uint32(data)
	} else {
		return binary.BigEndian.Uint32(data)
	}
}

// // Uint16转二进制.
// func Uint16ToByte(value uint16) []byte {
// 	ret := make([]byte, 2)
// 	if littleEndian {
// 		binary.LittleEndian.PutUint16(ret, value)
// 	} else {
// 		binary.BigEndian.PutUint16(ret, value)
// 	}
// 	return ret
// }

/************************************************************************/
// 模块内功能实现
/************************************************************************/

func init() {
	littleEndian = true
}
