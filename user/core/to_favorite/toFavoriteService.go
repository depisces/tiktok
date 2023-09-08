package to_favorite

import (
	"context"
	"errors"
	"user/model"
	proto "user/services/to_favorite"
)

type ToFavoriteService struct {
}

func (ToFavoriteService) UpdateTotalFavorited(ctx context.Context, req *proto.UpdateTotalFavoritedRequest, resp *proto.UpdateTotalFavoritedResponse) error {
	if req.UserId <= 0 || (req.Type != 1 && req.Type != 2) {
		resp.StatusCode = -1
		return errors.New("传入的userId或者type有误")
	}
	//查一下，这个userId能否查到，查不到报错，查到了返回count
	if _, err := model.NewUserDaoInstance().FindUserById(req.UserId); err != nil {
		return errors.New("传入的VideoId查不到")
	}
	//调用数据库的修改功能
	if req.Type == 1 {
		//增加
		model.NewUserDaoInstance().AddTotalFavorited(req.UserId, req.Count)
	} else if req.Type == 2 {
		//减少
		model.NewUserDaoInstance().ReduceTotalFavorited(req.UserId, req.Count)
	}

	resp.StatusCode = 0
	return nil
}

func (ToFavoriteService) UpdateFavoriteCount(ctx context.Context, req *proto.UpdateFavoriteCountRequest, resp *proto.UpdateFavoriteCountResponse) error {
	if req.UserId <= 0 || (req.Type != 1 && req.Type != 2) {
		resp.StatusCode = -1
		return errors.New("传入的userId或者type有误")
	}
	//查一下，这个userId能否查到，查不到报错，查到了返回count
	if _, err := model.NewUserDaoInstance().FindUserById(req.UserId); err != nil {
		return errors.New("传入的VideoId查不到")
	}
	//调用数据库的修改功能
	if req.Type == 1 {
		//增加
		model.NewUserDaoInstance().AddFavoriteCount(req.UserId, req.Count)
	} else if req.Type == 2 {
		//减少
		model.NewUserDaoInstance().ReduceFavoriteCount(req.UserId, req.Count)
	}

	resp.StatusCode = 0
	return nil
}
