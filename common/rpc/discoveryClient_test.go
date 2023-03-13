package rpc

import (
	"fmt"
	"log"
	"testing"
)

func TestDiscoverService(t *testing.T) {
	srvName := "threadService"
	srv, err := NewDiscoveryClient().DiscoverService(srvName)
	if err != nil {
		log.Panic(err)
	}
	address := fmt.Sprintf("%s:%d", srv.Address, srv.Port)
	if address != "127.0.0.1:8000" {
		t.Fatal("error address:", address)
	}
	t.Log(address)
}

func TestHealthCheck(t *testing.T) {
	rpc := NewRpcClient()
	NewDiscoveryClient().HealthCheck(rpc)
}
