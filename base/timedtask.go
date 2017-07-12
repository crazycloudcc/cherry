/*
 * timed task.
 * author: CC
 * email : crazycloudcc@gmail.com
 * date  : 2017.06.17
 */
package base

import (
	"time"
)

/************************************************************************/
// constants, variables, structs, interfaces.
/************************************************************************/

type TimedTaskCallback func()

type TimedTask struct {
	ID       uint64            // 定时任务唯一ID.
	Callback TimedTaskCallback // 定时任务事件.
	Timer    *time.Timer       // 定时任务管理器.
}

var timedTaskSeed uint64

/************************************************************************/
// 模块对外接口.
/************************************************************************/

// 添加定时任务 - 超时调用一次.
func SetTimeout(duration time.Duration, callback TimedTaskCallback) *TimedTask {
	defer seedTick()
	tt := new(TimedTask)
	tt.ID = timedTaskSeed
	tt.Callback = callback
	tt.Timer = time.AfterFunc(duration, tt.Callback)
	return tt
}

// 添加定时任务 - 间隔时间调用一次.
func SetInterval(duration time.Duration, callback TimedTaskCallback) *TimedTask {
	defer seedTick()
	tt := new(TimedTask)
	tt.ID = timedTaskSeed
	tt.Callback = func() {
		callback()
		tt.Timer = time.AfterFunc(duration, tt.Callback)
	}
	tt.Timer = time.AfterFunc(duration, tt.Callback)
	return tt
}

/************************************************************************/
// 模块内功能实现
/************************************************************************/

func init() {
	timedTaskSeed = 1
}

func seedTick() {
	timedTaskSeed++
}

/************************************************************************/
// 模块内功能调试
/************************************************************************/

func UnitTest_TimedTask() {

	// LogDebug("SetTimeout function. ", time.Now())
	// SetTimeout(100*1e9, func() {
	// 	LogDebug("SetTimeout function. ", time.Now())
	// })

	LogDebug("SetInterval function. ", time.Now())
	SetInterval(1*1e9, func() {
		LogDebug("SetInterval function. ", time.Now())
	})
}
