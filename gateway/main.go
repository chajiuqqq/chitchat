package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/chajiuqqq/chitchat/common/discover"
	"github.com/chajiuqqq/chitchat/common/loadbalance"
	"github.com/chajiuqqq/chitchat/gateway/handler"
	"github.com/hashicorp/consul/api"
)

var (
	gatwayPort  = flag.String("port", "8000", "gateway port")
	metricsPort = flag.String("metricsPort", "8500", "metrics port")
)

func main() {
	flag.Parse()
	discoverClient := discover.NewConsulClient()
	myProxy := handler.NewHystrixHandler(discoverClient, loadbalance.NewRandomLoadBalance())
	errChan := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)

	}()
	go func() {
		log.Println("gateway start at port", *gatwayPort)
		http.ListenAndServe(":"+*gatwayPort, myProxy)
	}()
	go func() {
		handler := hystrix.NewStreamHandler()
		handler.Start()
		http.Handle("/hystrix/stream", handler)
		http.ListenAndServe(":"+*metricsPort, nil)
	}()
	log.Println("gateway exit", <-errChan)
}

func NewProxy(client *api.Client) *httputil.ReverseProxy {
	director := func(req *http.Request) {
		paths := strings.Split(req.URL.Path, "/")
		srvName := paths[1]
		srvs, _, err := client.Catalog().Service(srvName, "", nil)
		if err != nil {
			log.Println("get service fail", err)
		}
		if len(srvs) == 0 {
			log.Println("service not found")
		}
		dstPath := strings.Join(paths[2:], "/")
		target := srvs[rand.Int()%len(srvs)]
		log.Println("targetService ID", target.ServiceID)
		req.URL.Scheme = "http"
		req.URL.Host = fmt.Sprintf("%s:%d", target.ServiceAddress, target.ServicePort)
		req.URL.Path = "/" + dstPath

	}
	return &httputil.ReverseProxy{Director: director}

}
