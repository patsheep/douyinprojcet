package service

import (
	"github.com/RaymondCode/simple-demo/dao"
	"github.com/RaymondCode/simple-demo/entities"
	"github.com/RaymondCode/simple-demo/util"
)

func AddNewUser(userid, password string) error {
	md5ps := util.GetMd5Val(password)
	user := entities.User2{
		Id:            4,
		UserId:        userid,
		Password:      md5ps,
		Name:          "",
		FollowCount:   0,
		FollowerCount: 0,
		IsFollow:      false,
	}
	err := dao.AddNewUser(user)
	return err
}
