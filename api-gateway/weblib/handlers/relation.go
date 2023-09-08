package handlers

import (
	"api-gateway/services/relation"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func RelationAction(ginCtx *gin.Context) {
	var relationReq relation.DouyinRelationActionRequest
	//获取request的信息
	relationReq.ToUserId, _ = strconv.ParseInt(ginCtx.Query("to_user_id"), 10, 64)
	actionType, _ := strconv.Atoi(ginCtx.Query("action_type"))
	relationReq.ActionType = int32(actionType)
	relationReq.Token = ginCtx.Query("token")

	// 从gin.Key中取出服务实例
	relationService := ginCtx.Keys["relationService"].(relation.RelationService)
	//调用comment微服务，将context的上下文传入
	relationResp, _ := relationService.RelationAction(context.Background(), &relationReq)

	//返回
	ginCtx.JSON(http.StatusOK, relation.DouyinRelationActionResponse{
		StatusCode: relationResp.StatusCode,
		StatusMsg:  relationResp.StatusMsg,
	})
}

func FollowerList(ginCtx *gin.Context) {
	var relationReq relation.DouyinRelationFollowerListRequest

	relationReq.Token = ginCtx.Query("token")
	relationReq.UserId, _ = strconv.ParseInt(ginCtx.Query("user_id"), 10, 64)

	relationService := ginCtx.Keys["relationService"].(relation.RelationService)
	resp, _ := relationService.FollowerList(context.Background(), &relationReq)

	ginCtx.JSON(http.StatusOK, relation.DouyinRelationFollowerListResponse{
		StatusCode: resp.StatusCode,
		StatusMsg:  resp.StatusMsg,
		UserList:   resp.UserList,
	})
}

func FollowList(ginCtx *gin.Context) {
	var relationReq relation.DouyinRelationFollowListRequest

	relationReq.Token = ginCtx.Query("token")
	relationReq.UserId, _ = strconv.ParseInt(ginCtx.Query("user_id"), 10, 64)

	relationService := ginCtx.Keys["relationService"].(relation.RelationService)
	resp, _ := relationService.FollowList(context.Background(), &relationReq)

	ginCtx.JSON(http.StatusOK, relation.DouyinRelationFollowListResponse{
		StatusCode: resp.StatusCode,
		StatusMsg:  resp.StatusMsg,
		UserList:   resp.UserList,
	})
}

func FriendList(ginCtx *gin.Context) {
	var relationReq relation.DouyinRelationFriendListRequest

	relationReq.Token = ginCtx.Query("token")
	relationReq.UserId, _ = strconv.ParseInt(ginCtx.Query("user_id"), 10, 64)

	relationService := ginCtx.Keys["relationService"].(relation.RelationService)
	resp, _ := relationService.FriendList(context.Background(), &relationReq)

	ginCtx.JSON(http.StatusOK, relation.DouyinRelationFriendListResponse{
		StatusCode: resp.StatusCode,
		StatusMsg:  resp.StatusMsg,
		UserList:   resp.UserList,
	})
}
