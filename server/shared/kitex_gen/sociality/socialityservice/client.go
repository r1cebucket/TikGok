// Code generated by Kitex v0.4.4. DO NOT EDIT.

package socialityservice

import (
	"context"
	sociality "github.com/CyanAsterisk/TikGok/server/shared/kitex_gen/sociality"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Action(ctx context.Context, req *sociality.DouyinRelationActionRequest, callOptions ...callopt.Option) (r *sociality.DouyinRelationActionResponse, err error)
	GetFollowingList(ctx context.Context, req *sociality.DouyinGetRelationFollowListRequest, callOptions ...callopt.Option) (r *sociality.DouyinGetRelationFollowListResponse, err error)
	GetFollowerList(ctx context.Context, req *sociality.DouyinGetRelationFollowerListRequest, callOptions ...callopt.Option) (r *sociality.DouyinGetRelationFollowerListResponse, err error)
	GetFriendList(ctx context.Context, req *sociality.DouyinGetRelationFriendListRequest, callOptions ...callopt.Option) (r *sociality.DouyinGetRelationFriendListResponse, err error)
	CheckFollow(ctx context.Context, req *sociality.DouyinCheckFollowRequest, callOptions ...callopt.Option) (r *sociality.DouyinCheckFollowResponse, err error)
	GetFollowerCount(ctx context.Context, req *sociality.DouyinGetFollowerCountRequest, callOptions ...callopt.Option) (r *sociality.DouyinGetFollowerCountResponse, err error)
	GetFollowingCount(ctx context.Context, req *sociality.DouyinGetFollowingCountRequest, callOptions ...callopt.Option) (r *sociality.DouyinGetFollowingCountResponse, err error)
	BatchCheckFollow(ctx context.Context, req *sociality.DouyinBatchCheckFollowRequest, callOptions ...callopt.Option) (r *sociality.DouyinBatchCheckFollowResponse, err error)
	BatchGetFollowerCount(ctx context.Context, req *sociality.DouyinBatchGetFollowerCountRequest, callOptions ...callopt.Option) (r *sociality.DouyinBatchGetFollowerCountResponse, err error)
	BatchGetFollowingCount(ctx context.Context, req *sociality.DouyinBatchGetFollowingCountRequest, callOptions ...callopt.Option) (r *sociality.DouyinBatchGetFollowingCountResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kSocialityServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kSocialityServiceClient struct {
	*kClient
}

func (p *kSocialityServiceClient) Action(ctx context.Context, req *sociality.DouyinRelationActionRequest, callOptions ...callopt.Option) (r *sociality.DouyinRelationActionResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Action(ctx, req)
}

func (p *kSocialityServiceClient) GetFollowingList(ctx context.Context, req *sociality.DouyinGetRelationFollowListRequest, callOptions ...callopt.Option) (r *sociality.DouyinGetRelationFollowListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetFollowingList(ctx, req)
}

func (p *kSocialityServiceClient) GetFollowerList(ctx context.Context, req *sociality.DouyinGetRelationFollowerListRequest, callOptions ...callopt.Option) (r *sociality.DouyinGetRelationFollowerListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetFollowerList(ctx, req)
}

func (p *kSocialityServiceClient) GetFriendList(ctx context.Context, req *sociality.DouyinGetRelationFriendListRequest, callOptions ...callopt.Option) (r *sociality.DouyinGetRelationFriendListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetFriendList(ctx, req)
}

func (p *kSocialityServiceClient) CheckFollow(ctx context.Context, req *sociality.DouyinCheckFollowRequest, callOptions ...callopt.Option) (r *sociality.DouyinCheckFollowResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CheckFollow(ctx, req)
}

func (p *kSocialityServiceClient) GetFollowerCount(ctx context.Context, req *sociality.DouyinGetFollowerCountRequest, callOptions ...callopt.Option) (r *sociality.DouyinGetFollowerCountResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetFollowerCount(ctx, req)
}

func (p *kSocialityServiceClient) GetFollowingCount(ctx context.Context, req *sociality.DouyinGetFollowingCountRequest, callOptions ...callopt.Option) (r *sociality.DouyinGetFollowingCountResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetFollowingCount(ctx, req)
}

func (p *kSocialityServiceClient) BatchCheckFollow(ctx context.Context, req *sociality.DouyinBatchCheckFollowRequest, callOptions ...callopt.Option) (r *sociality.DouyinBatchCheckFollowResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.BatchCheckFollow(ctx, req)
}

func (p *kSocialityServiceClient) BatchGetFollowerCount(ctx context.Context, req *sociality.DouyinBatchGetFollowerCountRequest, callOptions ...callopt.Option) (r *sociality.DouyinBatchGetFollowerCountResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.BatchGetFollowerCount(ctx, req)
}

func (p *kSocialityServiceClient) BatchGetFollowingCount(ctx context.Context, req *sociality.DouyinBatchGetFollowingCountRequest, callOptions ...callopt.Option) (r *sociality.DouyinBatchGetFollowingCountResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.BatchGetFollowingCount(ctx, req)
}
