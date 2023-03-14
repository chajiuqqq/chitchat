package rpc

import (
	"errors"
	"log"
	"time"

	capi "github.com/hashicorp/consul/api"
)

type discoveryClient struct {
	*capi.Client
	services map[string]*capi.AgentService
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
	myDiscoveryClient.services = make(map[string]*capi.AgentService)
	return myDiscoveryClient
}

func (d *discoveryClient) DiscoverService(service string) (*capi.AgentService, error) {
	srvs, err := d.Client.Agent().Services()
	if err != nil {
		return nil, err
	}
	if srv, ok := srvs[service]; ok {
		d.services[service] = srv
		return srv, nil
	}
	return nil, errors.New("No such service.")

}

func (d *discoveryClient) HealthCheck(c Consumer) {
	go func() {
		for {
			select {
			case <-time.Tick(time.Second):
				// log.Println("check...")
				for srvName, srv := range d.services {
					status, checkInfo, err := d.Client.Agent().AgentHealthServiceByID(srvName)
					if err != nil {
						log.Println("fail to get service health,", err)
					}
					if status == "critical" {
						log.Println("service critical:", srvName)
					}
					if statusChanged(srv, checkInfo) {
						log.Println(srvName,"status changed")
						done, err := c.Reload(checkInfo.Service)
						<-done
						log.Println("reload finished")
						d.services[srvName] = checkInfo.Service
						if err != nil {
							log.Panic(srvName, "service reload error:", err)
						}
					}
				}
			}
		}
	}()
}

func statusChanged(srv *capi.AgentService, ck *capi.AgentServiceChecksInfo) (changed bool) {
	if srv.Address != ck.Service.Address || srv.Port != ck.Service.Port {
		return true
	}
	return false
}
