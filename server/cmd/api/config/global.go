package config

import (
	"github.com/CyanAsterisk/TikGok/server/cmd/api/pkg/uploadService"
	"github.com/CyanAsterisk/TikGok/server/shared/kitex_gen/chat/chatservice"
	"github.com/CyanAsterisk/TikGok/server/shared/kitex_gen/interaction/interactionserver"
	"github.com/CyanAsterisk/TikGok/server/shared/kitex_gen/sociality/socialityservice"
	"github.com/CyanAsterisk/TikGok/server/shared/kitex_gen/user/userservice"
	"github.com/CyanAsterisk/TikGok/server/shared/kitex_gen/video/videoservice"
)

var (
	GlobalServerConfig  = &ServerConfig{} // from nacos (remote)
	GlobalNacosConfig   = &NacosConfig{}  // local nacos config
	GlobalUploadService *uploadService.Service

	// init rpc client
	GlobalChatClient        chatservice.Client
	GlobalUserClient        userservice.Client
	GlobalVideoClient       videoservice.Client
	GlobalSocialClient      socialityservice.Client
	GlobalInteractionClient interactionserver.Client
)
