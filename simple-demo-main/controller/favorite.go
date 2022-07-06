package controller

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/api"
	"github.com/RaymondCode/simple-demo/dao"
	"github.com/RaymondCode/simple-demo/service"
	"strings"

	"github.com/gin-gonic/gin"
	"net/http"
)

// FavoriteAction no practical effect, just check if token is valid
//三个参数，token,video_id,action_type
func FavoriteAction(c *gin.Context) {
	token := c.Query("token")
	if exist, _ := service.GetToken(token); exist!=0 {
		c.JSON(http.StatusOK, api.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	fmt.Println(token)
	user_id :=strings.Split(token, ":")[0]
	action_type :=c.Query("action_type")
	video_id :=c.Query("video_id")
	service.AddFavorite(video_id,user_id,action_type)
	c.JSON(http.StatusOK, api.Response{StatusCode: 0, StatusMsg: "点赞成功"})

}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	token := c.Query("token")
	if exist, _ := service.GetToken(token); exist!=0 {
		c.JSON(http.StatusOK, api.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}
	user_id :=strings.Split(token, ":")[0]
	videoIdList :=service.GetFavoriteVideoIdList(user_id)
	videoList:=dao.GetFavoriteVideoList(*videoIdList)
	var resList []api.Video
	for _,node := range videoList  {
		user :=dao.GetUserByIdInt64(node.AuthorId)

		resList=append(resList, api.Video{
			Id:            node.Id,
			Author:        api.User{
				Id:            user.Id,
				Name:          user.Name,
				FollowCount:   user.FollowCount,
				FollowerCount: user.FollowerCount,
				IsFollow:      user.IsFollow,
			},
			PlayUrl:       node.PlayUrl,
			CoverUrl:      node.CoverUrl,
			FavoriteCount: node.FavoriteCount,
			CommentCount:  node.CommentCount,
			IsFavorite:    true,
		})

		
	}
	fmt.Printf("结果长度为：%d",len(resList))
	for _,node :=range resList{
		fmt.Println("%+v\n",node)
	}
	c.JSON(http.StatusOK, VideoListResponse{
		Response: api.Response{
			StatusCode: 0,
		},
		VideoList: resList,
	})
}
