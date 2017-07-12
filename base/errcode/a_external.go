/*
 * [file desc]
 * author: CC
 * email : crazycloudcc@gmail.com
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

var errfmt *base.Map

/************************************************************************/
// export functions.
/************************************************************************/

func NewError(code int32, args ...interface{}) *ErrCode {
	f := errfmt.Get(code).(string)
	// TODO.
	ec := new(ErrCode)
	ec.Code = code
	ec.Err = errors.New(fmt.Sprintf(f, args...))
	return ec
}

/************************************************************************/
// moudule functions.
/************************************************************************/

func init() {
	errfmt = base.NewMap(1)
	errfmt.Add(int32(1), "unknown error.")
}

/************************************************************************/
// unit tests.
/************************************************************************/
