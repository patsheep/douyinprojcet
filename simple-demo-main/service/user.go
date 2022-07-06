package service

import (
	"github.com/RaymondCode/simple-demo/dao"
	"github.com/RaymondCode/simple-demo/entities"
	"github.com/RaymondCode/simple-demo/util"
	"github.com/RaymondCode/simple-demo/util/snowflake"
)

func AddNewUser(userid, password string) error {
	md5ps := util.GetMd5Val(password)
	user := entities.User{
		Id:            snowflake.MakeInt64SnowFlakeId(),
		UserId:        userid,
		Password:      md5ps,
		Name:          userid,
		FollowCount:   0,
		FollowerCount: 0,
		IsFollow:      false,
	}
	err := dao.AddNewUser(user)
	return err
}
