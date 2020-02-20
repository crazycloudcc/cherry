package dbproxy

/*
 * [file desc]
 * author: CC
 * email : 151503324@qq.com
 * date  : 2017.06.17
 */

/************************************************************************/
// constants, variables, structs, interfaces.
/************************************************************************/

// InsRedisRemote remote redis instance.
var InsRedisRemote *Redis

// InsRedisLocal local redis instance.
var InsRedisLocal *Redis

/************************************************************************/
// export functions.
/************************************************************************/

// RedisConnectRemote TODO.
func RedisConnectRemote(conf RedisConfig) bool {
	InsRedisRemote = newRedisProxy(conf)
	if InsRedisRemote == nil {
		return false
	}
	return true
}

// RedisConnectLocal TODO.
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
