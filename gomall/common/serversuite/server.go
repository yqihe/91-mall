package serversuite

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/config-consul/consul"
	consulServer "github.com/kitex-contrib/config-consul/server"
	registryconsul "github.com/kitex-contrib/registry-consul"
	"os"
)

type CommonServerSuite struct {
	CurrentServiceName string
	RegistryAddr       string
}

func (s CommonServerSuite) Options() []server.Option {
	opts := []server.Option{
		server.WithMetaHandler(transmeta.ServerHTTP2Handler),
	}

	r, err := registryconsul.NewConsulRegister(s.RegistryAddr)
	if err != nil {
		klog.Fatal(err)
	}

	opts = append(opts, server.WithRegistry(r))

	if os.Getenv("CONFIG_CENTER_ENABLED") == "true" {
		consulNodes := os.Getenv("CONFIG_CENTER_NODES")
		if consulNodes != "" {
			consulClient, err := consul.NewClient(consul.Options{})
			if err != nil {
				klog.Error(err)
			} else {
				opts = append(opts, server.WithSuite(consulServer.NewSuite(s.CurrentServiceName, consulClient)))
			}
		}
	}

	opts = append(opts,
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: s.CurrentServiceName,
		}),
	)

	return opts
}
