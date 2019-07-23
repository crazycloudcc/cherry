/*
 * [file desc]
 * author: CC
 * email : 151503324@qq.com
 * date  : 2017.06.17
 */
package random

import "math/rand"

/************************************************************************/
// constants, variables, structs, interfaces.
/************************************************************************/

type randInt32 struct {
	tick int32
}

/************************************************************************/
// export functions.
/************************************************************************/

func (this *randInt32) Value(min, max int32) int32 {
	this.tick++
	return min + rand.Int31n(max-min)
}

/************************************************************************/
// moudule functions.
/************************************************************************/

/************************************************************************/
// unit tests.
/************************************************************************/
