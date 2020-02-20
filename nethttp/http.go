package nethttp

/*
 * web http module.
 * usage: 1-NewHttp, 2-RegHandler, 3-Start.
 * author: CC
 * email : 151503324@qq.com
 * date  : 2017.06.17
 */

import (
	"fmt"
	"net/http"

	"github.com/crazycloudcc/cherry/base"
)

/************************************************************************/
// constants, variables, structs, interfaces.
/************************************************************************/

// HTTP TODO
type HTTP struct {
	conf Config // 配置文件.
}

/************************************************************************/
// export functions.
/************************************************************************/

// NewHTTP TODO
func NewHTTP(conf Config) *HTTP {
	h := new(HTTP)
	h.conf = conf
	return h
}

// Start service.
func (owner *HTTP) Start() {
	go owner.doStart()
	base.LogInfo("nethttp.Http Started with:", owner.conf.String())
}

// Stop service.
func (owner *HTTP) Stop() {
	// TODO.
}

// RegHandler TODO
func (owner *HTTP) RegHandler(uri string, handleFunc WebHandleFunc) {
	http.HandleFunc(uri, handleFunc)
}

/************************************************************************/
// moudule functions.
/************************************************************************/

func (owner *HTTP) doStart() {
	url := fmt.Sprintf("0.0.0.0:%d", owner.conf.Port)
	err := http.ListenAndServe(url, nil)
	if err != nil {
		base.LogError("nethttp.Http doStart Error:", err)
	}
}

/************************************************************************/
// unit tests.
/************************************************************************/
