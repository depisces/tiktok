package rpc_server

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	"publish/rpc_server/etcd"
	from_user_proto "publish/services/from_user"
	usersproto "publish/services/to_relation"
)

func GetUsersInfo(userId []int64, token string) ([]*usersproto.User, error) {
	userMicroService := micro.NewService(micro.Registry(etcdInit.EtcdReg))
	usersService := usersproto.NewToRelationService("rpcUserService", userMicroService.Client())

	var req usersproto.GetUsersByIdsRequest

	req.UserId = userId
	req.Token = token

	resp, err := usersService.GetUsersByIds(context.TODO(), &req)
	if err != nil {
		fmt.Println("调用远程UserInfo服务失败，具体错误如下")
		fmt.Println(err)
	}

	return resp.UserList, err
}

func UpdateWorkCount(uid int64, count int32, actionType int32) bool {
	toPublishMicroService := micro.NewService(micro.Registry(etcdInit.EtcdReg))
	toPublishService := from_user_proto.NewToPublishService("rpcUserService", toPublishMicroService.Client())
	var req from_user_proto.UpdateWorkCountRequest
	req.UserId = uid
	req.Count = count
	req.Type = actionType
	resp, err := toPublishService.UpdateWorkCount(context.TODO(), &req)
	if err != nil || resp.StatusCode != 0 {
		fmt.Println("work_count维护失败:", err)
		return false
	}
	return true

}
