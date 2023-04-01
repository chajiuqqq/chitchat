package client

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/chajiuqqq/chitchat/common/discover"
	"github.com/chajiuqqq/chitchat/common/loadbalance"
	"google.golang.org/grpc"
)

var (
	ErrServiceNotFound = errors.New("No such RPC service")
	DefaultLoadBalance = loadbalance.NewRandomLoadBalance()
)

// interface
type ClientManager interface {
	Invoke(path string, hystrixName string, ctx context.Context, inVal interface{}, outVal interface{}) error
}

// struct
type DefaultManager struct {
	serviceName    string
	discoverClient discover.DiscoveryClient
	loadBalance    loadbalance.LoadBalance
}

func (d *DefaultManager) Invoke(path string, hystrixName string, ctx context.Context, inVal interface{}, outVal interface{}) error {
	err := hystrix.Do(hystrixName, func() error {
		srvs := d.discoverClient.DiscoverServices(d.serviceName, nil)
		targetSrv, err := d.loadBalance.GetInstance(srvs)
		if err != nil {
			return ErrServiceNotFound
		}
		if conn, err := grpc.Dial(targetSrv.ServiceAddress+":"+strconv.Itoa(targetSrv.ServicePort), grpc.WithInsecure(), grpc.WithTimeout(time.Second)); err == nil {
			if err = conn.Invoke(ctx, path, inVal, outVal); err != nil {
				return err
			}
		}
		return err
	}, func(err error) error {
		return err
	})

	return err
}
