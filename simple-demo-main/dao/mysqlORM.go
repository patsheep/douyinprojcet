package dao

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/entities"

	//"fmt" // 导入 fmt 包，打印字符串是需要用到
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
var Db *gorm.DB


func ConnectDb() {
	var (
		err error
	)
	Db, err = gorm.Open(mysql.Open("root:97782078@tcp(127.0.0.1:3306)/douyin?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

}

func GetList() []entities.Video {
	//node := entities.Tbs{Val: 1}
	var lists []entities.Video2;


	Db.Table("video").Where("id > ?", 0).Debug().Find(&lists)
	var resLists []entities.Video;

	for i:=0;i< len(lists);i++  {

		var userNode entities.User2;

		Db.Table("user").Where("id = ?",lists[i].AuthorId).Debug().Find(&userNode)

		 resuserNode:= entities.User{
			Id:           userNode.Id,
			Name:         userNode.Name,
			FollowCount:  userNode.FollowCount,
			FollowerCount: userNode.FollowerCount,
			IsFollow:      userNode.IsFollow,
		};
		resLists=append(resLists, entities.Video{
			Id:            lists[i].Id,
			Author:        resuserNode,
			PlayUrl:       lists[i].PlayUrl,
			CoverUrl:      lists[i].CoverUrl,
			FavoriteCount: lists[i].FavoriteCount,
			CommentCount:  lists[i].CommentCount,
			IsFavorite:    lists[i].IsFavorite,
		})





	}
	for i:=0;i<len(resLists);i++{
		fmt.Print("%v+\n",resLists[i])
	}
	return resLists

}

type OSSKey struct {
	Key           string  `form:"key"`
	Secret        string `form:"secret"`

}
func GetOSSKEy() []string{
	var res OSSKey;
	Db.Table("osskey").Debug().First(&res)
	return []string{res.Key,res.Secret}
}

