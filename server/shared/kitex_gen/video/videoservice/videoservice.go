// Code generated by Kitex v0.4.4. DO NOT EDIT.

package videoservice

import (
	"context"
	video "github.com/CyanAsterisk/TikGok/server/shared/kitex_gen/video"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return videoServiceServiceInfo
}

var videoServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "VideoService"
	handlerType := (*video.VideoService)(nil)
	methods := map[string]kitex.MethodInfo{
		"Feed":         kitex.NewMethodInfo(feedHandler, newVideoServiceFeedArgs, newVideoServiceFeedResult, false),
		"PublishVideo": kitex.NewMethodInfo(publishVideoHandler, newVideoServicePublishVideoArgs, newVideoServicePublishVideoResult, false),
		"VideoList":    kitex.NewMethodInfo(videoListHandler, newVideoServiceVideoListArgs, newVideoServiceVideoListResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "video",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.4.4",
		Extra:           extra,
	}
	return svcInfo
}

func feedHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*video.VideoServiceFeedArgs)
	realResult := result.(*video.VideoServiceFeedResult)
	success, err := handler.(video.VideoService).Feed(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVideoServiceFeedArgs() interface{} {
	return video.NewVideoServiceFeedArgs()
}

func newVideoServiceFeedResult() interface{} {
	return video.NewVideoServiceFeedResult()
}

func publishVideoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*video.VideoServicePublishVideoArgs)
	realResult := result.(*video.VideoServicePublishVideoResult)
	success, err := handler.(video.VideoService).PublishVideo(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVideoServicePublishVideoArgs() interface{} {
	return video.NewVideoServicePublishVideoArgs()
}

func newVideoServicePublishVideoResult() interface{} {
	return video.NewVideoServicePublishVideoResult()
}

func videoListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*video.VideoServiceVideoListArgs)
	realResult := result.(*video.VideoServiceVideoListResult)
	success, err := handler.(video.VideoService).VideoList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVideoServiceVideoListArgs() interface{} {
	return video.NewVideoServiceVideoListArgs()
}

func newVideoServiceVideoListResult() interface{} {
	return video.NewVideoServiceVideoListResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Feed(ctx context.Context, req *video.DouyinFeedRequest) (r *video.DouyinFeedResponse, err error) {
	var _args video.VideoServiceFeedArgs
	_args.Req = req
	var _result video.VideoServiceFeedResult
	if err = p.c.Call(ctx, "Feed", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) PublishVideo(ctx context.Context, req *video.DouyinPublishActionRequest) (r *video.DouyinPublishActionResponse, err error) {
	var _args video.VideoServicePublishVideoArgs
	_args.Req = req
	var _result video.VideoServicePublishVideoResult
	if err = p.c.Call(ctx, "PublishVideo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) VideoList(ctx context.Context, req *video.DouyinPublishListRequest) (r *video.DouyinPublishListResponse, err error) {
	var _args video.VideoServiceVideoListArgs
	_args.Req = req
	var _result video.VideoServiceVideoListResult
	if err = p.c.Call(ctx, "VideoList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
