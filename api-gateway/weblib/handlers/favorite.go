package handlers

import (
	"api-gateway/services/fav"
	"github.com/gin-gonic/gin"
	"net/http"

	"strconv"
)

func FavoriteAction(ginCtx *gin.Context) {
	actionType, _ := strconv.Atoi(ginCtx.Query("action_type"))
	token := ginCtx.Query("token")
	vid, _ := strconv.ParseInt(ginCtx.Query("video_id"), 10, 64)
	favService := ginCtx.Keys["favoriteService"].(fav.FavoriteService)
	var req fav.DouyinFavoriteActionRequest
	req.ActionType = int32(actionType)
	req.Token = token
	req.VideoId = vid
	action, _ := favService.FavoriteAction(ginCtx, &req)
	ginCtx.JSON(http.StatusOK, fav.DouyinFavoriteActionResponse{
		StatusCode: action.StatusCode,
		StatusMsg:  action.StatusMsg,
	})
}
func FavoriteList(ginCtx *gin.Context) {
	token := ginCtx.Query("token")
	uid, _ := strconv.ParseInt(ginCtx.Query("user_id"), 10, 64)
	favService := ginCtx.Keys["favoriteService"].(fav.FavoriteService)
	var req fav.DouyinFavoriteListRequest
	req.Token = token
	req.UserId = uid
	action, _ := favService.FavoriteList(ginCtx, &req)
	ginCtx.JSON(http.StatusOK, fav.DouyinFavoriteListResponse{
		StatusCode: action.StatusCode,
		StatusMsg:  action.StatusMsg,
		VideoList:  action.VideoList,
	})
}
