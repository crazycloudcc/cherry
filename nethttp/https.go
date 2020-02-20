package nethttp

/*
 * web https module.
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

// HTTPS TODO
type HTTPS struct {
	conf Config // 配置文件.
}

/************************************************************************/
// export functions.
/************************************************************************/

// NewHTTPS create service proxy.
func NewHTTPS(conf Config) *HTTPS {
	h := new(HTTPS)
	h.conf = conf
	return h
}

// Start service.
func (owner *HTTPS) Start() {
	go owner.doStart()
	base.LogInfo("nethttp.HTTPS Started with:", owner.conf.String())
}

// Stop service.
func (owner *HTTPS) Stop() {
	// TODO.
}

// RegHandler TODO
func (owner *HTTPS) RegHandler(uri string, handleFunc WebHandleFunc) {
	http.HandleFunc(uri, handleFunc)
}

/************************************************************************/
// moudule functions.
/************************************************************************/

func (owner *HTTPS) doStart() {
	url := fmt.Sprintf("%s:%d", owner.conf.Host, owner.conf.Port)
	err := http.ListenAndServeTLS(url, "crt/server.crt", "crt/server.key", nil)
	if err != nil {
		base.LogError("nethttp.HTTPS doStart Error:", err)
	}
}

/************************************************************************/
// unit tests.
/************************************************************************/
