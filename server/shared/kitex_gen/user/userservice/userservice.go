// Code generated by Kitex v0.4.4. DO NOT EDIT.

package userservice

import (
	"context"
	user "github.com/CyanAsterisk/TikGok/server/shared/kitex_gen/user"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return userServiceServiceInfo
}

var userServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "UserService"
	handlerType := (*user.UserService)(nil)
	methods := map[string]kitex.MethodInfo{
		"Register":         kitex.NewMethodInfo(registerHandler, newUserServiceRegisterArgs, newUserServiceRegisterResult, false),
		"Login":            kitex.NewMethodInfo(loginHandler, newUserServiceLoginArgs, newUserServiceLoginResult, false),
		"GetUserInfo":      kitex.NewMethodInfo(getUserInfoHandler, newUserServiceGetUserInfoArgs, newUserServiceGetUserInfoResult, false),
		"BatchGetUserInfo": kitex.NewMethodInfo(batchGetUserInfoHandler, newUserServiceBatchGetUserInfoArgs, newUserServiceBatchGetUserInfoResult, false),
		"GetFollowList":    kitex.NewMethodInfo(getFollowListHandler, newUserServiceGetFollowListArgs, newUserServiceGetFollowListResult, false),
		"GetFollowerList":  kitex.NewMethodInfo(getFollowerListHandler, newUserServiceGetFollowerListArgs, newUserServiceGetFollowerListResult, false),
		"GetFriendList":    kitex.NewMethodInfo(getFriendListHandler, newUserServiceGetFriendListArgs, newUserServiceGetFriendListResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "user",
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

func registerHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceRegisterArgs)
	realResult := result.(*user.UserServiceRegisterResult)
	success, err := handler.(user.UserService).Register(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceRegisterArgs() interface{} {
	return user.NewUserServiceRegisterArgs()
}

func newUserServiceRegisterResult() interface{} {
	return user.NewUserServiceRegisterResult()
}

func loginHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceLoginArgs)
	realResult := result.(*user.UserServiceLoginResult)
	success, err := handler.(user.UserService).Login(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceLoginArgs() interface{} {
	return user.NewUserServiceLoginArgs()
}

func newUserServiceLoginResult() interface{} {
	return user.NewUserServiceLoginResult()
}

func getUserInfoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceGetUserInfoArgs)
	realResult := result.(*user.UserServiceGetUserInfoResult)
	success, err := handler.(user.UserService).GetUserInfo(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceGetUserInfoArgs() interface{} {
	return user.NewUserServiceGetUserInfoArgs()
}

func newUserServiceGetUserInfoResult() interface{} {
	return user.NewUserServiceGetUserInfoResult()
}

func batchGetUserInfoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceBatchGetUserInfoArgs)
	realResult := result.(*user.UserServiceBatchGetUserInfoResult)
	success, err := handler.(user.UserService).BatchGetUserInfo(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceBatchGetUserInfoArgs() interface{} {
	return user.NewUserServiceBatchGetUserInfoArgs()
}

func newUserServiceBatchGetUserInfoResult() interface{} {
	return user.NewUserServiceBatchGetUserInfoResult()
}

func getFollowListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceGetFollowListArgs)
	realResult := result.(*user.UserServiceGetFollowListResult)
	success, err := handler.(user.UserService).GetFollowList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceGetFollowListArgs() interface{} {
	return user.NewUserServiceGetFollowListArgs()
}

func newUserServiceGetFollowListResult() interface{} {
	return user.NewUserServiceGetFollowListResult()
}

func getFollowerListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceGetFollowerListArgs)
	realResult := result.(*user.UserServiceGetFollowerListResult)
	success, err := handler.(user.UserService).GetFollowerList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceGetFollowerListArgs() interface{} {
	return user.NewUserServiceGetFollowerListArgs()
}

func newUserServiceGetFollowerListResult() interface{} {
	return user.NewUserServiceGetFollowerListResult()
}

func getFriendListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceGetFriendListArgs)
	realResult := result.(*user.UserServiceGetFriendListResult)
	success, err := handler.(user.UserService).GetFriendList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceGetFriendListArgs() interface{} {
	return user.NewUserServiceGetFriendListArgs()
}

func newUserServiceGetFriendListResult() interface{} {
	return user.NewUserServiceGetFriendListResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Register(ctx context.Context, req *user.DouyinUserRegisterRequest) (r *user.DouyinUserRegisterResponse, err error) {
	var _args user.UserServiceRegisterArgs
	_args.Req = req
	var _result user.UserServiceRegisterResult
	if err = p.c.Call(ctx, "Register", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Login(ctx context.Context, req *user.DouyinUserLoginRequest) (r *user.DouyinUserLoginResponse, err error) {
	var _args user.UserServiceLoginArgs
	_args.Req = req
	var _result user.UserServiceLoginResult
	if err = p.c.Call(ctx, "Login", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetUserInfo(ctx context.Context, req *user.DouyinGetUserRequest) (r *user.DouyinGetUserResponse, err error) {
	var _args user.UserServiceGetUserInfoArgs
	_args.Req = req
	var _result user.UserServiceGetUserInfoResult
	if err = p.c.Call(ctx, "GetUserInfo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) BatchGetUserInfo(ctx context.Context, req *user.DouyinBatchGetUserRequest) (r *user.DouyinBatchGetUserResonse, err error) {
	var _args user.UserServiceBatchGetUserInfoArgs
	_args.Req = req
	var _result user.UserServiceBatchGetUserInfoResult
	if err = p.c.Call(ctx, "BatchGetUserInfo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetFollowList(ctx context.Context, req *user.DouyinGetRelationFollowListRequest) (r *user.DouyinGetRelationFollowListResponse, err error) {
	var _args user.UserServiceGetFollowListArgs
	_args.Req = req
	var _result user.UserServiceGetFollowListResult
	if err = p.c.Call(ctx, "GetFollowList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetFollowerList(ctx context.Context, req *user.DouyinGetRelationFollowerListRequest) (r *user.DouyinGetRelationFollowerListResponse, err error) {
	var _args user.UserServiceGetFollowerListArgs
	_args.Req = req
	var _result user.UserServiceGetFollowerListResult
	if err = p.c.Call(ctx, "GetFollowerList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetFriendList(ctx context.Context, req *user.DouyinGetRelationFriendListRequest) (r *user.DouyinGetRelationFriendListResponse, err error) {
	var _args user.UserServiceGetFriendListArgs
	_args.Req = req
	var _result user.UserServiceGetFriendListResult
	if err = p.c.Call(ctx, "GetFriendList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
