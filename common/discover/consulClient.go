package discover

import (
	"log"

	capi "github.com/hashicorp/consul/api"
)

type consulClient struct {
	Client *capi.Client
}

func NewConsulClient() *consulClient {

	config := capi.DefaultConfig()
	config.Address = "consul:8500"
	c, err := capi.NewClient(config)
	if err != nil {
		log.Panic("fail to create discoveryClient,", err)
	}
	return &consulClient{
		Client: c,
	}
}

/**
 * 服务注册接口
 * @param serviceName 服务名
 * @param instanceId 服务实例Id
 * @param instancePort 服务实例端口
 * @param healthCheckUrl 健康检查地址
 * @param instanceHost 服务实例地址
 * @param meta 服务实例元数据
 */
func (cc consulClient) Register(serviceName, instanceId, healthCheckUrl string, instanceHost string, instancePort int, meta map[string]string, logger *log.Logger) error {

	// 创建服务实例
	service := &capi.AgentServiceRegistration{
		Name:    serviceName,
		Port:    instancePort,
		Address: instanceHost,
		Check: &capi.AgentServiceCheck{
			HTTP:     healthCheckUrl,
			Interval: "10s",
			Timeout:  "2s",
		},
	}

	// 注册服务
	err := cc.Client.Agent().ServiceRegister(service)
	return err
}

/**
 * 服务注销接口
 * @param instanceId 服务实例Id
 */
func (cc consulClient) DeRegister(instanceId string, logger *log.Logger) bool {
	// 删除服务实例
	err := cc.Client.Agent().ServiceDeregister(instanceId)
	return err == nil
}

/**
 * 发现服务实例接口
 * @param serviceName 服务名
 */
func (cc consulClient) DiscoverServices(serviceName string, logger *log.Logger) []*capi.CatalogService {
	srvs, _, err := cc.Client.Catalog().Service(serviceName, "", nil)
	if err != nil {
		return nil
	}
	return srvs
}
