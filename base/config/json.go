/*
 * handle json file.
 * author: CC
 * email : 151503324@qq.com
 * date  : 2017.06.17
 */
package config

import (
	"encoding/json"
	"io/ioutil"
)

/************************************************************************/
// constants, variables, structs, interfaces.
/************************************************************************/

const JSON string = "json"

/************************************************************************/
// export functions.
/************************************************************************/

/************************************************************************/
// moudule functions.
/************************************************************************/

func jsonRead(fn string, reply interface{}) error {
	ioutil.ReadFile(fn)
	buf, err := ioutil.ReadFile(fn)
	if err != nil {
		return err
	}

	err = json.Unmarshal(buf, reply)
	if err != nil {
		return err
	}
	return nil
}

/************************************************************************/
// unit tests.
/************************************************************************/
