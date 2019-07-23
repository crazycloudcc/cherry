/*
 * [file desc]
 * author: CC
 * email : 151503324@qq.com
 * date  : 2017.06.17
 */
package dbproxy

/************************************************************************/
// constants, variables, structs, interfaces.
/************************************************************************/

var InsRedisRemote *Redis // remote redis.
var InsRedisLocal *Redis  // local redis.

/************************************************************************/
// export functions.
/************************************************************************/

func RedisConnectRemote(conf RedisConfig) bool {
	InsRedisRemote = newRedisProxy(conf)
	if InsRedisRemote == nil {
		return false
	}
	return true
}

func RedisConnectLocal(conf RedisConfig) bool {
	InsRedisLocal = newRedisProxy(conf)
	if InsRedisLocal == nil {
		return false
	}
	return true
}

/************************************************************************/
// moudule functions.
/************************************************************************/

/************************************************************************/
// unit tests.
/************************************************************************/
