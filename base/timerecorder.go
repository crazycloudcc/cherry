/*
 * time recorder.
 * author: CC
 * email : 151503324@qq.com
 * date  : 2017.06.17
 */
package base

import (
	"fmt"
	"sync"
	"time"
)

/************************************************************************/
// constants, variables, structs, interfaces.
/************************************************************************/

type TimeRecorder struct {
	sync.Mutex
	BeginTime    time.Time
	formatPrefix string
}

/************************************************************************/
// export functions.
/************************************************************************/

func NewTimeRecorder() *TimeRecorder {
	tr := new(TimeRecorder)
	tr.BeginTime = time.Now()
	return tr
}

func (this *TimeRecorder) Begin(formatPrefix string) {
	this.Lock()
	defer this.Unlock()
	this.formatPrefix = formatPrefix
	this.BeginTime = time.Now()
	// LogDebug(fmt.Sprintf("%s, begin time:[%v].", this.formatPrefix, this.BeginTime.Format("2006/01/02 15:04:05.000")))
}

func (this *TimeRecorder) End() {
	this.Lock()
	defer this.Unlock()
	// LogDebug(fmt.Sprintf("%s, end time:[%v] cost time:[%v].",
	// 	this.formatPrefix, time.Now().Format("2006/01/02 15:04:05.000"), time.Now().Sub(this.BeginTime)))
	LogInfo(fmt.Sprintf("%s, cost time:[%v].", this.formatPrefix, time.Now().Sub(this.BeginTime)))
}

/************************************************************************/
// moudule functions.
/************************************************************************/

/************************************************************************/
// unit tests.
/************************************************************************/
