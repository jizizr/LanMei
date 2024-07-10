package common

import (
	"context"
	"errors"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/discovery"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/hashicorp/consul/api"
	"github.com/jizizr/LanMei/server/rpc_gen/kitex_gen/bot"
	"github.com/jizizr/LanMei/server/rpc_gen/kitex_gen/rpc"
	"github.com/jizizr/LanMei/server/rpc_gen/kitex_gen/rpc/rpcservice"
	consul "github.com/kitex-contrib/registry-consul"
	"log"
	"sync"
	"time"
)

type ServiceManager struct {
	services map[string]string
	command  *sync.Map
	text     *sync.Map
	c        *api.Client
	r        discovery.Resolver
}

type RpcClientWithDescription struct {
	Client rpcservice.Client
	Desc   string
}

func (cmd *ServiceManager) sync(services map[string][]string) {
	for service := range cmd.services {
		if _, ok := services[service]; !ok {
			cmd.command.Delete(cmd.services[service])
			cmd.text.Delete(service)
			delete(cmd.services, service)
		}
	}
	for service := range services {
		if _, ok := cmd.services[service]; !ok {
			c, err := rpcservice.NewClient(service, client.WithResolver(cmd.r))
			if err != nil {
				klog.Error("Error creating client: %v", err)
			}
			var t rpc.CmdType
			err = errors.New("")
			dur := 1000
			for i := 0; err != nil && i < 3; i++ {
				time.Sleep(time.Duration(dur) * time.Millisecond)
				dur *= 2
				fmt.Println("try to get service type: ", service)
				t, err = c.Type(context.Background(), &rpc.Empty{})
			}
			if err != nil {
				klog.Error("Error getting service type: ", err, ",service: ", service)
				continue
			}
			switch t {
			case rpc.CmdType_COMMAND:
				desc, err := c.Command(context.Background(), &rpc.Empty{})
				if err != nil {
					klog.Error("Error getting command description: ", err, ",service: ", service)
					continue
				}
				fmt.Println(desc.Cmd, desc.Description)
				cmd.services[service] = desc.Cmd
				cmd.command.Store(desc.Cmd, RpcClientWithDescription{
					Client: c,
					Desc:   desc.Description,
				})
			case rpc.CmdType_TEXT:
				cmd.services[service] = ""
				cmd.text.Store(service, c)
			}
		}
	}
}

func (cmd *ServiceManager) CallCommand(command string, message *bot.Message) (bool, error) {
	if v, ok := cmd.command.Load(command); ok {
		c := v.(RpcClientWithDescription).Client
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		f, err := c.Call(ctx, message)
		if f == false && err != nil {
			f, err = c.Call(ctx, message)
		}
		return f, err
	}
	return false, nil
}

func (cmd *ServiceManager) Watch() {
	var lastIndex uint64
	for {
		// Query the catalog for a list of services
		services, meta, err := cmd.c.Catalog().Services(&api.QueryOptions{
			WaitIndex: lastIndex,
			WaitTime:  5 * time.Second,
		})
		if err != nil {
			log.Fatalf("Error fetching services: %v", err)
		}

		// Check if there is any change in the services
		if meta.LastIndex != lastIndex {
			lastIndex = meta.LastIndex
			delete(services, "consul")
			cmd.sync(services)
		}
	}
}

func NewServiceManager(registryAddress string) (*ServiceManager, error) {
	config := api.DefaultConfig()
	config.Address = registryAddress
	client, err := api.NewClient(config)
	if err != nil {
		return nil, err
	}
	r, err := consul.NewConsulResolver(registryAddress)
	if err != nil {
		return nil, err
	}
	return &ServiceManager{
		services: make(map[string]string),
		command:  &sync.Map{},
		text:     &sync.Map{},
		c:        client,
		r:        r,
	}, nil
}
