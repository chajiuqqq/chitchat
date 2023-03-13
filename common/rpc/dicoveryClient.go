package rpc

import (
	"errors"
	"log"

	capi "github.com/hashicorp/consul/api"
)

type discoveryClient struct {
	*capi.Client
}

var myDiscoveryClient *discoveryClient

func NewDiscoveryClient() *discoveryClient {
	if myDiscoveryClient != nil {
		return myDiscoveryClient
	}
	c, err := capi.NewClient(capi.DefaultConfig())
	if err != nil {
		log.Panic("fail to create discoveryClient,", err)
	}
	myDiscoveryClient = new(discoveryClient)
	myDiscoveryClient.Client = c
	return myDiscoveryClient
}

func (d *discoveryClient) DiscoverService(service string) (*capi.AgentService, error) {
	srvs, err := d.Client.Agent().Services()
	if err != nil {
		return nil, err
	}
	if srv, ok := srvs[service]; ok {
		return srv, nil
	}
	return nil, errors.New("No such service.")

}
