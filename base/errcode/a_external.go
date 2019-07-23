/*
 * [file desc]
 * author: CC
 * email : 151503324@qq.com
 * date  : 2017.06.17
 */
package errcode

import (
	"cherry/base"
	"errors"
	"fmt"
)

/************************************************************************/
// constants, variables, structs, interfaces.
/************************************************************************/

/************************************************************************/
// export functions.
/************************************************************************/

// init module from json file.
func Init(fn string) {
}

// create error code.
func NewErrorCode(code int32, args ...interface{}) *ErrCode {
	v := errfmt.Get(code)
	ec := new(ErrCode)
	ec.Code = code
	if v == nil {
		v = errfmt.Get(Unknown)
		ec.Err = errors.New(v.(string))
	} else {
		f := v.(string)
		if args != nil {
			ec.Err = errors.New(fmt.Sprintf(f, args...))
		}
	}
	return ec
}

/************************************************************************/
// moudule functions.
/************************************************************************/

// register error code.
func regErrorCode(code int32, fmtString string) {
	if errfmt.Get(code) != nil {
		base.LogError("error code: ", code, fmtString, "is already not exists")
		return
	}
	errfmt.Add(code, fmtString)
}

/************************************************************************/
// unit tests.
/************************************************************************/
