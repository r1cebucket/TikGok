package rpc

import (
	"github.com/CyanAsterisk/TikGok/server/cmd/api/config"
	"github.com/CyanAsterisk/TikGok/server/shared/consts"
	"github.com/CyanAsterisk/TikGok/server/shared/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/loadbalance"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	nacos "github.com/kitex-contrib/registry-nacos/resolver"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

// InitUser init user service.
func initUser() {
	// init resolver
	// Read configuration information from nacos
	sc := []constant.ServerConfig{
		{
			IpAddr: config.GlobalNacosConfig.Host,
			Port:   config.GlobalNacosConfig.Port,
		},
	}

	cc := constant.ClientConfig{
		NamespaceId:         config.GlobalNacosConfig.Namespace,
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              consts.NacosLogDir,
		CacheDir:            consts.NacosCacheDir,
		LogLevel:            consts.NacosLogLevel,
	}

	nacosCli, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		})
	// create resolver from nacos
	resolver := nacos.NewNacosResolver(nacosCli, nacos.WithGroup(consts.UserGroup))
	if err != nil {
		klog.Fatalf("new nacos client failed: %s", err.Error())
	}

	// init provider from kitex:
	// 	tarce: init trace provider
	// 	metric: inti metirc pusher
	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(config.GlobalServerConfig.Name),
		provider.WithExportEndpoint(config.GlobalServerConfig.OtelInfo.EndPoint), // collector address
		provider.WithInsecure(),
	)

	// create a user service client
	c, err := userservice.NewClient(
		config.GlobalServerConfig.UserSrvInfo.Name,
		client.WithResolver(resolver),                              // service discovery
		client.WithLoadBalancer(loadbalance.NewWeightedBalancer()), // load balance from kitex
		client.WithMuxConnection(1),                                // multiplexing (#conn)
		client.WithSuite(tracing.NewClientSuite()),                 // trace from kitex otel
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: config.GlobalServerConfig.UserSrvInfo.Name}),
	)
	if err != nil {
		klog.Fatalf("ERROR: cannot init client: %v\n", err)
	}
	config.GlobalUserClient = c
}
