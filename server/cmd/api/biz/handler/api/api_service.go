// Code generated by hertz generator.

package api

import (
	"context"

	"github.com/CyanAsterisk/TikGok/server/cmd/api/biz/model/api"
	"github.com/CyanAsterisk/TikGok/server/cmd/api/global"
	"github.com/CyanAsterisk/TikGok/server/cmd/api/tools"
	"github.com/CyanAsterisk/TikGok/server/shared/consts"
	"github.com/CyanAsterisk/TikGok/server/shared/errno"
	"github.com/CyanAsterisk/TikGok/server/shared/kitex_gen/chat"
	"github.com/CyanAsterisk/TikGok/server/shared/kitex_gen/interaction"
	"github.com/CyanAsterisk/TikGok/server/shared/kitex_gen/sociality"
	"github.com/CyanAsterisk/TikGok/server/shared/kitex_gen/user"
	"github.com/CyanAsterisk/TikGok/server/shared/kitex_gen/video"
	"github.com/CyanAsterisk/TikGok/server/shared/middleware"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

// Register .
// @router /douyin/user/register [POST]
func Register(ctx context.Context, c *app.RequestContext) {
	resp := new(api.DouyinUserRegisterResponse)
	var req api.DouyinUserRegisterRequest
	err := c.BindAndValidate(&req)
	if err != nil {
		resp.StatusCode = int32(errno.ParamsEr.ErrCode)
		resp.StatusMsg = errno.ParamsEr.ErrMsg
		errno.SendResponse(c, resp)
		return
	}
	res, err := global.UserClient.Register(ctx, &user.DouyinUserRegisterRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		hlog.Errorf("rpc call user server err:%s", err.Error())
		resp.StatusCode = int32(errno.RPCUserErr.ErrCode)
		resp.StatusMsg = errno.RPCUserErr.ErrMsg
		errno.SendResponse(c, resp)
		return
	}
	resp.StatusCode = res.BaseResp.StatusCode
	resp.StatusMsg = res.BaseResp.StatusMsg
	resp.Token = res.Token
	resp.UserID = res.UserId
	errno.SendResponse(c, resp)
}

// Login .
// @router /douyin/user/login [POST]
func Login(ctx context.Context, c *app.RequestContext) {
	resp := new(api.DouyinUserLoginResponse)
	var req api.DouyinUserLoginRequest
	err := c.BindAndValidate(&req)
	if err != nil {
		resp.StatusCode = int32(errno.ParamsEr.ErrCode)
		resp.StatusMsg = errno.ParamsEr.ErrMsg
		errno.SendResponse(c, resp)
		return
	}

	res, err := global.UserClient.Login(ctx, &user.DouyinUserLoginRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		resp.StatusCode = int32(errno.RPCUserErr.ErrCode)
		resp.StatusMsg = errno.RPCUserErr.ErrMsg
		errno.SendResponse(c, resp)
		return
	}
	resp.StatusCode = res.BaseResp.StatusCode
	resp.StatusMsg = res.BaseResp.StatusMsg
	resp.Token = res.Token
	resp.UserID = res.UserId
	errno.SendResponse(c, resp)
}

// GetUserInfo .
// @router /douyin/user [GET]
func GetUserInfo(ctx context.Context, c *app.RequestContext) {
	resp := new(api.DouyinUserResponse)
	var req api.DouyinUserRequest
	err := c.BindAndValidate(&req)
	if err != nil {
		resp.StatusCode = int32(errno.ParamsEr.ErrCode)
		resp.StatusMsg = errno.ParamsEr.Error()
		errno.SendResponse(c, resp)
		return
	}
	aid, flag := c.Get(consts.AccountID)
	if !flag {
		resp.StatusCode = int32(errno.ServiceErr.ErrCode)
		resp.StatusMsg = errno.ServiceErr.ErrMsg
		errno.SendResponse(c, resp)
		return
	}
	res, err := global.UserClient.GetUserInfo(ctx, &user.DouyinUserRequest{
		ViewerId: aid.(int64),
		OwnerId:  req.UserID,
	})
	if err != nil {
		resp.StatusCode = int32(errno.RPCUserErr.ErrCode)
		resp.StatusMsg = errno.RPCUserErr.ErrMsg
		errno.SendResponse(c, resp)
		return
	}
	resp.StatusCode = res.BaseResp.StatusCode
	resp.StatusMsg = res.BaseResp.StatusMsg
	resp.User = tools.User(res.User)
	errno.SendResponse(c, resp)
}

// Feed .
// @router /douyin/feed [GET]
func Feed(ctx context.Context, c *app.RequestContext) {
	resp := new(api.DouyinFeedResponse)
	var err error
	var req api.DouyinFeedRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		resp.StatusCode = int32(errno.ParamsEr.ErrCode)
		resp.StatusMsg = errno.ParamsEr.ErrMsg
		errno.SendResponse(c, resp)
		return
	}
	aid := int64(0)
	if req.Token != "" {
		j := middleware.NewJWT(global.ServerConfig.JWTInfo.SigningKey)
		claims, err := j.ParseToken(req.Token)
		if err != nil {
			resp.StatusCode = int32(errno.ParamsEr.ErrCode)
			resp.StatusMsg = "bad token"
			errno.SendResponse(c, resp)
			return
		}
		aid = claims.ID
	}
	res, err := global.VideoClient.Feed(ctx, &video.DouyinFeedRequest{
		LatestTime: req.LatestTime,
		ViewerId:   aid,
	})
	if err != nil {
		hlog.Errorf("feed err:%s", err.Error())
		resp.StatusCode = int32(errno.RPCVideoErr.ErrCode)
		resp.StatusMsg = errno.RPCVideoErr.ErrMsg
		errno.SendResponse(c, resp)
		return
	}
	resp.StatusCode = res.BaseResp.StatusCode
	resp.StatusMsg = res.BaseResp.StatusMsg
	resp.NextTime = res.NextTime
	resp.VideoList = tools.Videos(res.VideoList)
	errno.SendResponse(c, resp)
}

// PublishVideo .
// @router /douyin/publish/action [POST]
func PublishVideo(ctx context.Context, c *app.RequestContext) {
	resp := new(api.DouyinPublishActionResponse)
	var err error
	var req api.DouyinPublishActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		resp.StatusCode = int32(errno.ParamsEr.ErrCode)
		resp.StatusMsg = errno.ParamsEr.Error()
		errno.SendResponse(c, resp)
		return
	}
	aid, flag := c.Get(consts.AccountID)
	if !flag {
		resp.StatusCode = int32(errno.RPCVideoErr.ErrCode)
		resp.StatusMsg = errno.RPCVideoErr.ErrMsg
		errno.SendResponse(c, resp)
		return
	}

	fileHeader, err := c.Request.FormFile("data")
	if err != nil {
		resp.StatusCode = int32(errno.ParamsEr.ErrCode)
		resp.StatusMsg = errno.ParamsEr.ErrMsg
		errno.SendResponse(c, resp)
		return
	}

	playUrl, coverUrl, err := tools.UpLoadFile(fileHeader)

	res, err := global.VideoClient.PublishVideo(ctx, &video.DouyinPublishActionRequest{
		UserId:   aid.(int64),
		PlayUrl:  playUrl,
		CoverUrl: coverUrl,
		Title:    req.Title,
	})
	if err != nil {
		resp.StatusCode = int32(errno.RPCVideoErr.ErrCode)
		resp.StatusMsg = errno.RPCVideoErr.ErrMsg
		errno.SendResponse(c, resp)
		return
	}
	resp.StatusCode = res.BaseResp.StatusCode
	resp.StatusMsg = res.BaseResp.StatusMsg
	errno.SendResponse(c, resp)
}

// VideoList .
// @router /douyin/publish/list [GET]
func VideoList(ctx context.Context, c *app.RequestContext) {
	resp := new(api.DouyinPublishListResponse)
	var err error
	var req api.DouyinPublishListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		resp.StatusCode = int32(errno.ParamsEr.ErrCode)
		resp.StatusMsg = errno.ParamsEr.ErrMsg
		errno.SendResponse(c, resp)
		return
	}

	aid, flag := c.Get(consts.AccountID)
	if !flag {
		resp.StatusCode = int32(errno.RPCVideoErr.ErrCode)
		resp.StatusMsg = errno.RPCVideoErr.ErrMsg
		errno.SendResponse(c, resp)
		return
	}
	res, err := global.VideoClient.VideoList(ctx, &video.DouyinPublishListRequest{
		ViewerId: req.UserID,
		OwnerId:  aid.(int64),
	})
	if err != nil {
		resp.StatusCode = int32(errno.RPCVideoErr.ErrCode)
		resp.StatusMsg = errno.RPCVideoErr.ErrMsg
		errno.SendResponse(c, resp)
		return
	}
	resp.StatusCode = res.BaseResp.StatusCode
	resp.StatusMsg = res.BaseResp.StatusMsg
	resp.VideoList = tools.Videos(res.VideoList)
	errno.SendResponse(c, resp)
}

// Favorite .
// @router /douyin/favorite/action [POST]
func Favorite(ctx context.Context, c *app.RequestContext) {
	resp := new(api.DouyinFavoriteActionResponse)
	var err error
	var req api.DouyinFavoriteActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		resp.StatusCode = int32(errno.ParamsEr.ErrCode)
		resp.StatusMsg = errno.ParamsEr.ErrMsg
		errno.SendResponse(c, resp)
		return
	}

	aid, flag := c.Get(consts.AccountID)
	if !flag {
		resp.StatusCode = int32(errno.ServiceErr.ErrCode)
		resp.StatusMsg = errno.ServiceErr.ErrMsg
		errno.SendResponse(c, resp)
		return
	}
	res, err := global.InteractionClient.Favorite(ctx, &interaction.DouyinFavoriteActionRequest{
		UserId:     aid.(int64),
		VideoId:    req.VideoID,
		ActionType: req.ActionType,
	})
	if err != nil {
		resp.StatusCode = int32(errno.VideoServerErr.ErrCode)
		resp.StatusMsg = errno.VideoServerErr.ErrMsg
		errno.SendResponse(c, resp)
		return
	}
	resp.StatusCode = res.BaseResp.StatusCode
	resp.StatusMsg = res.BaseResp.StatusMsg
	errno.SendResponse(c, resp)
}

// FavoriteList .
// @router /douyin/favorite/list [GET]
func FavoriteList(ctx context.Context, c *app.RequestContext) {
	resp := new(api.DouyinFavoriteListResponse)
	var err error
	var req api.DouyinFavoriteListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		resp.StatusCode = int32(errno.ParamsEr.ErrCode)
		resp.StatusMsg = errno.ParamsEr.ErrMsg
		errno.SendResponse(c, resp)
		return
	}

	aid, flag := c.Get(consts.AccountID)
	if !flag {
		resp.StatusCode = int32(errno.ServiceErr.ErrCode)
		resp.StatusMsg = errno.ServiceErr.ErrMsg
		errno.SendResponse(c, resp)
		return
	}
	res, err := global.InteractionClient.FavoriteList(ctx, &interaction.DouyinFavoriteListRequest{
		OwnerId:  req.UserID,
		ViewerId: aid.(int64),
	})
	if err != nil {
		resp.StatusCode = int32(errno.RPCInteractionErr.ErrCode)
		resp.StatusMsg = errno.RPCInteractionErr.ErrMsg
		errno.SendResponse(c, resp)
		return
	}
	resp.StatusCode = res.BaseResp.StatusCode
	resp.StatusMsg = res.BaseResp.StatusMsg
	resp.VideoList = tools.Videos(res.VideoList)
	errno.SendResponse(c, resp)
}

// Comment .
// @router /douyin/comment/action [POST]
func Comment(ctx context.Context, c *app.RequestContext) {
	resp := new(api.DouyinCommentActionResponse)
	var err error
	var req api.DouyinCommentActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		resp.StatusCode = int32(errno.ParamsEr.ErrCode)
		resp.StatusMsg = errno.ParamsEr.ErrMsg
		errno.SendResponse(c, resp)
		return
	}

	aid, flag := c.Get(consts.AccountID)
	if !flag {
		resp.StatusCode = int32(errno.ServiceErr.ErrCode)
		resp.StatusMsg = errno.ServiceErr.ErrMsg
		errno.SendResponse(c, resp)
		return
	}
	res, err := global.InteractionClient.Comment(ctx, &interaction.DouyinCommentActionRequest{
		UserId:      aid.(int64),
		VideoId:     req.VideoID,
		ActionType:  req.ActionType,
		CommentText: req.CommentText,
		CommentId:   req.CommentID,
	})
	if err != nil {
		resp.StatusCode = int32(errno.RPCInteractionErr.ErrCode)
		resp.StatusMsg = errno.RPCInteractionErr.ErrMsg
		errno.SendResponse(c, resp)
		return
	}
	resp.StatusCode = res.BaseResp.StatusCode
	resp.StatusMsg = res.BaseResp.StatusMsg
	if resp.StatusCode == int32(errno.Success.ErrCode) {
		resp.Comment = tools.Comment(res.Comment)
	}
	errno.SendResponse(c, resp)
}

// CommentList .
// @router /douyin/comment/list [GET]
func CommentList(ctx context.Context, c *app.RequestContext) {
	resp := new(api.DouyinCommentListResponse)
	var err error
	var req api.DouyinCommentListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		resp.StatusCode = int32(errno.ParamsEr.ErrCode)
		resp.StatusMsg = errno.ParamsEr.ErrMsg
		errno.SendResponse(c, resp)
		return
	}
	res, err := global.InteractionClient.CommentList(ctx, &interaction.DouyinCommentListRequest{
		VideoId: req.VideoID,
	})
	if err != nil {
		resp.StatusCode = int32(errno.RPCInteractionErr.ErrCode)
		resp.StatusMsg = errno.RPCInteractionErr.ErrMsg
		errno.SendResponse(c, resp)
		return
	}
	resp.StatusCode = res.BaseResp.StatusCode
	resp.StatusMsg = res.BaseResp.StatusMsg
	resp.CommentList = tools.Comments(res.CommentList)
	errno.SendResponse(c, resp)
}

// Action .
// @router /douyin/relation/action [POST]
func Action(ctx context.Context, c *app.RequestContext) {
	resp := new(api.DouyinRelationActionResponse)
	var err error
	var req api.DouyinRelationActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		resp.StatusCode = int32(errno.ParamsEr.ErrCode)
		resp.StatusMsg = errno.ParamsEr.ErrMsg
		errno.SendResponse(c, resp)
		return
	}

	aid, flag := c.Get(consts.AccountID)
	if !flag {
		resp.StatusCode = int32(errno.ServiceErr.ErrCode)
		resp.StatusMsg = errno.ServiceErr.ErrMsg
		errno.SendResponse(c, resp)
		return
	}
	res, err := global.SocialClient.Action(ctx, &sociality.DouyinRelationActionRequest{
		UserId:     aid.(int64),
		ToUserId:   req.ToUserID,
		ActionType: req.ActionType,
	})
	if err != nil {
		resp.StatusCode = int32(errno.RPCSocialityErr.ErrCode)
		resp.StatusMsg = errno.RPCSocialityErr.ErrMsg
		errno.SendResponse(c, resp)
		return
	}
	resp.StatusCode = res.BaseResp.StatusCode
	resp.StatusMsg = res.BaseResp.StatusMsg
	errno.SendResponse(c, resp)
}

// FollowingList .
// @router /douyin/relation/follow/list [GET]
func FollowingList(ctx context.Context, c *app.RequestContext) {
	resp := new(api.DouyinRelationFollowListResponse)
	var err error
	var req api.DouyinRelationFollowListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		resp.StatusCode = int32(errno.ParamsEr.ErrCode)
		resp.StatusMsg = errno.ParamsEr.ErrMsg
		errno.SendResponse(c, resp)
		return
	}
	aid, flag := c.Get(consts.AccountID)
	if !flag {
		resp.StatusCode = int32(errno.ServiceErr.ErrCode)
		resp.StatusMsg = errno.ServiceErr.ErrMsg
		errno.SendResponse(c, resp)
		return
	}
	res, err := global.SocialClient.FollowingList(ctx, &sociality.DouyinRelationFollowListRequest{
		OwnerId:  req.UserID,
		ViewerId: aid.(int64),
	})
	if err != nil {
		resp.StatusCode = int32(errno.RPCSocialityErr.ErrCode)
		resp.StatusMsg = errno.RPCSocialityErr.ErrMsg
		errno.SendResponse(c, resp)
		return
	}
	resp.StatusCode = res.BaseResp.StatusCode
	resp.StatusMsg = res.BaseResp.StatusMsg
	resp.UserList = tools.Users(res.UserList)
	errno.SendResponse(c, resp)
}

// FollowerList .
// @router /douyin/relation/follower/list [GET]
func FollowerList(ctx context.Context, c *app.RequestContext) {
	resp := new(api.DouyinRelationFollowerListResponse)
	var err error
	var req api.DouyinRelationFollowerListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		resp.StatusCode = int32(errno.ParamsEr.ErrCode)
		resp.StatusMsg = errno.ParamsEr.ErrMsg
		errno.SendResponse(c, resp)
		return
	}
	aid, flag := c.Get(consts.AccountID)
	if !flag {
		resp.StatusCode = int32(errno.ServiceErr.ErrCode)
		resp.StatusMsg = errno.ServiceErr.ErrMsg
		errno.SendResponse(c, resp)
		return
	}
	res, err := global.SocialClient.FollowerList(ctx, &sociality.DouyinRelationFollowerListRequest{
		OwnerId:  req.UserID,
		ViewerId: aid.(int64),
	})
	if err != nil {
		resp.StatusCode = int32(errno.RPCSocialityErr.ErrCode)
		resp.StatusMsg = errno.RPCSocialityErr.ErrMsg
		errno.SendResponse(c, resp)
		return
	}
	resp.StatusCode = res.BaseResp.StatusCode
	resp.StatusMsg = res.BaseResp.StatusMsg
	resp.UserList = tools.Users(res.UserList)
	errno.SendResponse(c, resp)
}

// FriendList .
// @router /douyin/relation/friend/list [GET]
func FriendList(ctx context.Context, c *app.RequestContext) {
	resp := new(api.DouyinRelationFriendListResponse)
	var err error
	var req api.DouyinRelationFriendListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		resp.StatusCode = int32(errno.ParamsEr.ErrCode)
		resp.StatusMsg = errno.ParamsEr.ErrMsg
		errno.SendResponse(c, resp)
		return
	}
	aid, flag := c.Get(consts.AccountID)
	if !flag {
		resp.StatusCode = int32(errno.ServiceErr.ErrCode)
		resp.StatusMsg = errno.ServiceErr.ErrMsg
		errno.SendResponse(c, resp)
		return
	}
	res, err := global.SocialClient.FriendList(ctx, &sociality.DouyinRelationFriendListRequest{
		OwnerId:  req.UserID,
		ViewerId: aid.(int64),
	})
	if err != nil {
		resp.StatusCode = int32(errno.RPCSocialityErr.ErrCode)
		resp.StatusMsg = errno.RPCSocialityErr.ErrMsg
		errno.SendResponse(c, resp)
		return
	}
	resp.StatusCode = res.BaseResp.StatusCode
	resp.StatusMsg = res.BaseResp.StatusMsg
	resp.UserList = tools.FUsers(res.UserList)
	errno.SendResponse(c, resp)
}

// ChatHistory .
// @router /douyin/message/chat [GET]
func ChatHistory(ctx context.Context, c *app.RequestContext) {
	resp := new(api.DouyinMessageChatResponse)
	var err error
	var req api.DouyinMessageChatRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		resp.StatusCode = int32(errno.ParamsEr.ErrCode)
		resp.StatusMsg = errno.ParamsEr.ErrMsg
		errno.SendResponse(c, resp)
		return
	}
	aid, flag := c.Get(consts.AccountID)
	if !flag {
		resp.StatusCode = int32(errno.ServiceErr.ErrCode)
		resp.StatusMsg = errno.ServiceErr.ErrMsg
		errno.SendResponse(c, resp)
		return
	}

	res, err := global.ChatClient.ChatHistory(ctx, &chat.DouyinMessageChatRequest{
		UserId:   aid.(int64),
		ToUserId: req.ToUserID,
	})
	if err != nil {
		resp.StatusCode = int32(errno.RPCChatErr.ErrCode)
		resp.StatusMsg = errno.RPCChatErr.ErrMsg
		errno.SendResponse(c, resp)
		return
	}
	resp.StatusCode = res.BaseResp.StatusCode
	resp.StatusMsg = res.BaseResp.StatusMsg
	resp.MessageList = tools.Messages(res.MessageList)
	errno.SendResponse(c, resp)
}

// SentMessage .
// @router /douyin/message/chat [POST]
func SentMessage(ctx context.Context, c *app.RequestContext) {
	resp := new(api.DouyinMessageActionResponse)
	var err error
	var req api.DouyinMessageActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		resp.StatusCode = int32(errno.ParamsEr.ErrCode)
		resp.StatusMsg = errno.ParamsEr.ErrMsg
		errno.SendResponse(c, resp)
		return
	}
	aid, flag := c.Get(consts.AccountID)
	if !flag {
		resp.StatusCode = int32(errno.ServiceErr.ErrCode)
		resp.StatusMsg = errno.ServiceErr.ErrMsg
		errno.SendResponse(c, resp)
		return
	}
	res, err := global.ChatClient.SentMessage(ctx, &chat.DouyinMessageActionRequest{
		UserId:     aid.(int64),
		ToUserId:   req.ToUserID,
		ActionType: req.ActionType,
		Content:    req.Content,
	})
	if err != nil {
		resp.StatusCode = int32(errno.RPCChatErr.ErrCode)
		resp.StatusMsg = errno.RPCChatErr.ErrMsg
		errno.SendResponse(c, resp)
		return
	}
	resp.StatusCode = res.BaseResp.StatusCode
	resp.StatusMsg = res.BaseResp.StatusMsg
	errno.SendResponse(c, resp)
}
