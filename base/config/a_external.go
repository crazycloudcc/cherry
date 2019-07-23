/*
 * [file desc]
 * author: CC
 * email : 151503324@qq.com
 * date  : 2017.06.17
 */
package config

import "errors"

/************************************************************************/
// constants, variables, structs, interfaces.
/************************************************************************/

/************************************************************************/
// export functions.
/************************************************************************/

/* read json data from file.
 * param fn: file name.
 * param ft: file type. (example: json, xlsx, txt, ...)
 * param reply: read data.
 * return: function error info.
 */
func Read(fn string, ft string, reply interface{}) error {
	switch ft {
	case JSON:
		return jsonRead(fn, reply)
	}
	return errors.New("Read: Unkown Type File!")
}

/* write json data from file.
 * param fn: file name.
 * param ft: file type. (example: json, xlsx, txt, ...)
 * param data: write data.
 * return: function error info.
 */
func Write(fn string, ft string, data interface{}) error {
	return nil
}

/************************************************************************/
// moudule functions.
/************************************************************************/

/************************************************************************/
// unit tests.
/************************************************************************/
