package handlers

import (
	"api-gateway/services/publish"
	"bytes"
	"context"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strconv"
	"time"
)

func Publish(ginCtx *gin.Context) {
	var publishReq publish.DouyinPublishActionRequest

	publishReq.Title = ginCtx.PostForm("title")
	publishReq.Token = ginCtx.PostForm("token")
	fileHeader, _ := ginCtx.FormFile("data")

	file, err := fileHeader.Open()
	if err != nil {
		PanicIfPublishError(err)
		return
	}
	defer file.Close()

	buf := bytes.NewBuffer(nil)
	if _, err = io.Copy(buf, file); err != nil {
		PanicIfPublishError(err)
		return
	}

	publishReq.Data = buf.Bytes()

	ctx, _ := context.WithTimeout(ginCtx, time.Minute*1)
	// 从gin.Key中取出服务实例
	publishService := ginCtx.Keys["publishService"].(publish.PublishService)
	publishResp, err := publishService.Publish(ctx, &publishReq)
	PanicIfPublishError(err)

	ginCtx.JSON(http.StatusOK, publish.DouyinPublishActionResponse{
		StatusCode: publishResp.StatusCode,
		StatusMsg:  publishResp.StatusMsg,
	})
}

func PublishList(ginCtx *gin.Context) {
	var publishReq publish.DouyinPublishListRequest
	publishReq.Token = ginCtx.Query("token")
	ctx, _ := context.WithTimeout(ginCtx, time.Minute*1)

	//user_id绑定req.userId
	userId, err := strconv.ParseInt(ginCtx.Query("user_id"), 10, 64)
	PanicIfPublishError(err)
	publishReq.UserId = userId

	// 从gin.Key中取出服务实例
	publishService := ginCtx.Keys["publishService"].(publish.PublishService)
	publishResp, err := publishService.PublishList(ctx, &publishReq)
	PanicIfPublishError(err)

	ginCtx.JSON(http.StatusOK, publish.DouyinPublishListResponse{
		StatusCode: publishResp.StatusCode,
		StatusMsg:  publishResp.StatusMsg,
		VideoList:  publishResp.VideoList,
	})
}
