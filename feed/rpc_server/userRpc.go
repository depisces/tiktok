package rpc_server

import (
	"context"
	"feed/rpc_server/etcd"
	usersproto "feed/services/to_relation"
	"fmt"
	"github.com/micro/go-micro/v2"
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
