/*
 * redis connect proxy. (redigo)
 * author: CC
 * email : 151503324@qq.com
 * date  : 2017.06.17
 */
package dbproxy

import (
	"bytes"
	"cherry/base"
	"encoding/gob"
	"encoding/json"
	"fmt"

	"github.com/garyburd/redigo/redis"
)

/************************************************************************/
// constants, variables, structs, interfaces.
/************************************************************************/

type Redis struct {
	conf  RedisConfig
	rConn redis.Conn
	wConn redis.Conn
}

/************************************************************************/
// export functions.
/************************************************************************/

func (this *Redis) ReadCommand(cmd string, args ...interface{}) (interface{}, error) {
	return this.rConn.Do(cmd, args...)
}

func (this *Redis) ReadCommands(cmd string, args ...interface{}) error {
	return this.rConn.Send(cmd, args...)
}

func (this *Redis) ReadCommandsFlush() error {
	return this.rConn.Flush()
}

func (this *Redis) ReadCommandsReply() (interface{}, error) {
	return this.rConn.Receive()
}

func (this *Redis) WriteCommand(cmd string, args ...interface{}) (interface{}, error) {
	return this.wConn.Do(cmd, args...)
}

func (this *Redis) WriteCommands(cmd string, args ...interface{}) error {
	return this.wConn.Send(cmd, args...)
}

func (this *Redis) WriteCommandsFlush() error {
	return this.wConn.Flush()
}

func (this *Redis) WriteCommandsReply() (interface{}, error) {
	return this.wConn.Receive()
}

// convert interface{} to []interface{}
func RedisValues(reply interface{}, err error) ([]interface{}, error) {
	return redis.Values(reply, err)
}

// convert interface{} to redis data.
func RedisMarshalWithGob(data interface{}) *bytes.Buffer {
	buf := bytes.NewBuffer(nil)
	encoder := gob.NewEncoder(buf)
	err := encoder.Encode(data)
	if err != nil {
		base.LogError("RedisMarshalWithGob:", err)
		return nil
	}
	return buf
}

// convert redis data to interface{}.
func RedisUnmarshalWithGob(data []byte, reply interface{}) {
	buf := bytes.NewBuffer(data)
	decoder := gob.NewDecoder(buf)
	err := decoder.Decode(reply)
	if err != nil {
		base.LogError("RedisUnmarshalWithGob:", err)
	}
}

// convert interface{} to redis data. (json type.)
func RedisMarshal(data interface{}) []byte {
	buf, err := json.Marshal(data)
	if err != nil {
		base.LogError("RedisMarshal:", err)
		return nil
	}
	// base.LogDebug("dbprox/redis.go - buf info: ", string(buf))
	return buf
}

// convert redis data to interface{}. (json type.)
func RedisUnmarshal(data []byte, reply interface{}) {
	err := json.Unmarshal(data, reply)
	if err != nil {
		base.LogError("RedisUnmarshal:", err)
	}
}

/************************************************************************/
// moudule functions.
/************************************************************************/

// create redis connect proxy and connect to redis server.
func newRedisProxy(conf RedisConfig) *Redis {
	server := fmt.Sprintf("%s:%d", conf.Host, conf.Port)

	rconn, err := redis.Dial(conf.Type, server, redis.DialDatabase(int(conf.Database)), redis.DialPassword(conf.Password))
	if err != nil {
		base.LogError("newRedisProxy redis.Dial r:", err)
		return nil
	}

	wconn, err := redis.Dial(conf.Type, server, redis.DialDatabase(int(conf.Database)), redis.DialPassword(conf.Password))
	if err != nil {
		base.LogError("newRedisProxy redis.Dial w:", err)
		return nil
	}

	r := new(Redis)
	r.conf = conf
	r.rConn = rconn
	r.wConn = wconn

	base.LogInfo("newRedisProxy connected to", conf)
	return r
}

/************************************************************************/
// unit tests.
/************************************************************************/
