/*
 * [file desc]
 * author: CC
 * email : 151503324@qq.com
 * date  : 2017.06.17
 */
package timedtask

import "time"

/************************************************************************/
// constants, variables, structs, interfaces.
/************************************************************************/

type TimedTask struct {
	id       uint64
	taskType int32
	duration time.Duration
	callback TimedTaskCallback
	timer    *time.Timer
}

/************************************************************************/
// export functions.
/************************************************************************/

func (this *TimedTask) Reset() {
	this.timer.Reset(this.duration)
}

func (this *TimedTask) Stop() {
	this.timer.Stop()
}

/************************************************************************/
// moudule functions.
/************************************************************************/

func newTimedTask(t int32, duration time.Duration, callback TimedTaskCallback) *TimedTask {
	defer seedTick()
	tt := new(TimedTask)
	tt.id = seed
	tt.taskType = t
	tt.duration = duration
	switch tt.taskType {
	case 0: // timeout type.
		tt.callback = callback
	case 1: // interval type.
		tt.callback = func() {
			callback()
			tt.timer = time.AfterFunc(duration, tt.callback)
		}
	}
	tt.timer = time.AfterFunc(duration, tt.callback)
	return tt
}

/************************************************************************/
// unit tests.
/************************************************************************/
