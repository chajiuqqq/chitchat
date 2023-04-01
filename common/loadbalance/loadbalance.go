package loadbalance

import (
	"fmt"
	"math/rand"

	"github.com/hashicorp/consul/api"
)

type LoadBalance interface {
	// 获取负载均衡的实例
	GetInstance(srvs []*api.CatalogService) (*api.CatalogService, error)
}

type RandomLoadBalance struct {
}

func NewRandomLoadBalance() *RandomLoadBalance {
	return &RandomLoadBalance{}
}

func (r *RandomLoadBalance) GetInstance(srvs []*api.CatalogService) (*api.CatalogService, error) {
	if srvs == nil || len(srvs) == 0 {
		return nil, fmt.Errorf("service not found")
	}
	return srvs[rand.Intn(len(srvs))], nil
}
