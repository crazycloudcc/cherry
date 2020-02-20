package dbproxy

/*
 * redis connect proxy. (redigo)
 * author: CC
 * email : 151503324@qq.com
 * date  : 2017.06.17
 */

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"

	"github.com/crazycloudcc/cherry/base"

	"github.com/garyburd/redigo/redis"
)

/************************************************************************/
// constants, variables, structs, interfaces.
/************************************************************************/

// Redis TODO
type Redis struct {
	conf  RedisConfig
	rConn redis.Conn
	wConn redis.Conn
}

/************************************************************************/
// export functions.
/************************************************************************/

// ReadCommand TODO
func (owner *Redis) ReadCommand(cmd string, args ...interface{}) (interface{}, error) {
	return owner.rConn.Do(cmd, args...)
}

// ReadCommands TODO
func (owner *Redis) ReadCommands(cmd string, args ...interface{}) error {
	return owner.rConn.Send(cmd, args...)
}

// ReadCommandsFlush TODO
func (owner *Redis) ReadCommandsFlush() error {
	return owner.rConn.Flush()
}

// ReadCommandsReply TODO
func (owner *Redis) ReadCommandsReply() (interface{}, error) {
	return owner.rConn.Receive()
}

// WriteCommand TODO
func (owner *Redis) WriteCommand(cmd string, args ...interface{}) (interface{}, error) {
	return owner.wConn.Do(cmd, args...)
}

// WriteCommands TODO
func (owner *Redis) WriteCommands(cmd string, args ...interface{}) error {
	return owner.wConn.Send(cmd, args...)
}

// WriteCommandsFlush TODO
func (owner *Redis) WriteCommandsFlush() error {
	return owner.wConn.Flush()
}

// WriteCommandsReply TODO
func (owner *Redis) WriteCommandsReply() (interface{}, error) {
	return owner.wConn.Receive()
}

// RedisValues convert interface{} to []interface{}
func RedisValues(reply interface{}, err error) ([]interface{}, error) {
	return redis.Values(reply, err)
}

// RedisMarshalWithGob convert interface{} to redis data.
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

// RedisUnmarshalWithGob convert redis data to interface{}.
func RedisUnmarshalWithGob(data []byte, reply interface{}) {
	buf := bytes.NewBuffer(data)
	decoder := gob.NewDecoder(buf)
	err := decoder.Decode(reply)
	if err != nil {
		base.LogError("RedisUnmarshalWithGob:", err)
	}
}

// RedisMarshal convert interface{} to redis data. (json type.)
func RedisMarshal(data interface{}) []byte {
	buf, err := json.Marshal(data)
	if err != nil {
		base.LogError("RedisMarshal:", err)
		return nil
	}
	// base.LogDebug("dbprox/redis.go - buf info: ", string(buf))
	return buf
}

// RedisUnmarshal convert redis data to interface{}. (json type.)
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
