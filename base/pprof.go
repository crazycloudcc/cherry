/*
 * go pprof.
 * author: CC
 * email : crazycloudcc@gmail.com
 * date  : 2017.06.17
 */
package base

import (
	"os"
	"runtime/pprof"
	"sync"
)

/************************************************************************/
// constants, variables, structs, interfaces.
/************************************************************************/

type Pprof struct {
	sync.Mutex
	profile *string
	fp      *os.File
}

/************************************************************************/
// 模块对外接口.
/************************************************************************/

// 开启pprof
func (this *Pprof) Open(fname string) {
	// this.profile = flag.String("cpuprofile", "", "write cpu profile")
	// flag.Parse()
	// f, err := os.Create(*this.profile)
	f, err := os.Create(fname)
	if err != nil {
		LogError("os.Create:", err)
		return
	}
	this.fp = f
	pprof.StartCPUProfile(this.fp)
}

// 关闭pprof
func (this *Pprof) Close() {
	// defer os.Exit(1)
	pprof.StopCPUProfile()
	this.fp.Close()
	LogInfo("pprof.go - Closed")
}

/************************************************************************/
// 模块内功能实现
/************************************************************************/
