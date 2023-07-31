package main

import (
	"context"
	"net"
	"strconv"

	"github.com/CyanAsterisk/TikGok/server/cmd/user/config"
	"github.com/CyanAsterisk/TikGok/server/cmd/user/dao"
	"github.com/CyanAsterisk/TikGok/server/cmd/user/initialize"
	"github.com/CyanAsterisk/TikGok/server/cmd/user/pkg"
	"github.com/CyanAsterisk/TikGok/server/shared/consts"
	user "github.com/CyanAsterisk/TikGok/server/shared/kitex_gen/user/userservice"
	"github.com/CyanAsterisk/TikGok/server/shared/middleware"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/utils"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
)

func main() {
	// initialization

	// init logger for kitex (klog)
	initialize.InitLogger()
	// cmd line args
	IP, Port := initialize.InitFlag()

	// read local nacos config
	// read remote config from nacos and save to config.GlobalServerConfig
	// TODO: create registry and get registry info
	registry, registryInfo := initialize.InitNacos(Port)

	// create DB client and write log and trace to Otel
	db := initialize.InitDB()

	// TODO: provider from kitex
	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(config.GlobalServerConfig.Name),
		provider.WithExportEndpoint(config.GlobalServerConfig.OtelInfo.EndPoint),
		provider.WithInsecure(),
	)
	defer p.Shutdown(context.Background())

	rc := initialize.InitRedis()

	// rpc client
	socialClient := initialize.InitSocial()
	chatClient := initialize.InitChat()
	interactClient := initialize.InitInteraction()

	{ // TODO: rpc server
		// clients for rpc server
		impl := &UserServiceImpl{
			Jwt:                middleware.NewJWT(config.GlobalServerConfig.JWTInfo.SigningKey),
			SocialManager:      pkg.NewSocialManager(socialClient),
			InteractionManager: pkg.NewInteractionManager(interactClient),
			ChatManager:        pkg.NewChatManager(chatClient),
			RedisManager:       pkg.NewRedisManager(rc),
			Dao:                dao.NewUser(db),
		}

		// Create new server.
		srv := user.NewServer(impl,
			server.WithServiceAddr(utils.NewNetAddr(consts.TCP, net.JoinHostPort(IP, strconv.Itoa(Port)))),
			server.WithRegistry(registry),
			server.WithRegistryInfo(registryInfo),
			server.WithLimit(&limit.Option{MaxConnections: 2000, MaxQPS: 500}),
			server.WithSuite(tracing.NewServerSuite()),
			server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: config.GlobalServerConfig.Name}),
		)

		err := srv.Run()
		if err != nil {
			klog.Fatal(err)
		}
	}
}
