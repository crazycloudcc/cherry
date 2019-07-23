/*
 * binary convert functions.
 * author: CC
 * email : 151503324@qq.com
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
// export functions.
/************************************************************************/

// set little endian.
func SetLittleEndian() {
	littleEndian = true
}

// set big endian.
func SetBigEndian() {
	littleEndian = false
}

// covert uint16 to []byte.
func Uint16ToByte(dst []byte, v uint16) {
	if littleEndian {
		binary.LittleEndian.PutUint16(dst, v)
	} else {
		binary.BigEndian.PutUint16(dst, v)
	}
}

// covert uint32 to []byte.
func Uint32ToByte(dst []byte, v uint32) {
	if littleEndian {
		binary.LittleEndian.PutUint32(dst, v)
	} else {
		binary.BigEndian.PutUint32(dst, v)
	}
}

// covert []byte to uint16.
func ByteToUint16(data []byte) uint16 {
	if littleEndian {
		return binary.LittleEndian.Uint16(data)
	} else {
		return binary.BigEndian.Uint16(data)
	}
}

// covert []byte to uint32.
func ByteToUint32(data []byte) uint32 {
	if littleEndian {
		return binary.LittleEndian.Uint32(data)
	} else {
		return binary.BigEndian.Uint32(data)
	}
}

/************************************************************************/
// moudule functions.
/************************************************************************/

func init() {
	littleEndian = true
}

/************************************************************************/
// unit tests.
/************************************************************************/
