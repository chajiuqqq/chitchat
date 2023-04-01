package handler

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httputil"
	"strings"
	"sync"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/chajiuqqq/chitchat/common/discover"
	"github.com/chajiuqqq/chitchat/common/loadbalance"
)

type HystrixHandler struct {
	hystrixs   map[string]bool
	hystrixMux *sync.Mutex

	discoverClient discover.DiscoveryClient
	loadBalance    loadbalance.LoadBalance
}

func NewHystrixHandler(discoverClient discover.DiscoveryClient, loadBalance loadbalance.LoadBalance) *HystrixHandler {
	return &HystrixHandler{
		hystrixs:       make(map[string]bool),
		hystrixMux:     &sync.Mutex{},
		discoverClient: discoverClient,
		loadBalance:    loadBalance,
	}
}

func (h *HystrixHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	reqPath := req.URL.Path
	if reqPath == "" {
		return
	}
	pathArray := strings.Split(reqPath, "/")
	srvName := pathArray[1]
	if srvName == "" {
		rw.WriteHeader(404)
		return
	}
	if _, ok := h.hystrixs[srvName]; !ok {
		h.hystrixMux.Lock()
		if _, ok := h.hystrixs[srvName]; !ok {
			hystrix.ConfigureCommand(srvName, hystrix.CommandConfig{})
			h.hystrixs[srvName] = true
		}
		h.hystrixMux.Unlock()
	}
	err:=hystrix.Do(srvName, func() error {
		srvs := h.discoverClient.DiscoverServices(srvName, nil)
		srv, err := h.loadBalance.GetInstance(srvs)
		if err != nil {
			return err
		}

		director := func(req *http.Request) {

			dstPath := strings.Join(pathArray[2:], "/")
			req.URL.Scheme = "http"
			req.URL.Host = fmt.Sprintf("%s:%d", srv.ServiceAddress, srv.ServicePort)
			req.URL.Path = "/" + dstPath

		}
		var proxyErr error
		errHandler := func(rw http.ResponseWriter ,req *http.Request, err error) {
				proxyErr = err
		}

		proxy := &httputil.ReverseProxy{
			Director: director,
			ErrorHandler: errHandler,
		}
		proxy.ServeHTTP(rw,req)
		return proxyErr

	}, func(e error) error {
		return errors.New("fallback error:" + e.Error())
	})

	if err!=nil{
		rw.WriteHeader(500)
		rw.Write([]byte(err.Error()))
	}


}
