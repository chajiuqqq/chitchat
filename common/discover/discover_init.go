package discover

import "github.com/chajiuqqq/chitchat/common/loadbalance"

var(
	ConsulService DiscoveryClient
	LoadBalance loadbalance.LoadBalance
)

func init(){
	LoadBalance = loadbalance.NewRandomLoadBalance()
	ConsulService = NewConsulClient()
}