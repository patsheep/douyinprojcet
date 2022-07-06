package controller

import (
	"github.com/RaymondCode/simple-demo/api"
	"github.com/RaymondCode/simple-demo/dao"
	"github.com/RaymondCode/simple-demo/entities"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type UserListResponse struct {
	api.Response
	UserList []api.User `json:"user_list"`
}

// RelationAction no practical effect, just check if token is valid
//三个参数，token,to_user_id,action_type(1-关注，2-取消关注）
func RelationAction(c *gin.Context) {
	token := c.Query("token")
	action_type :=c.Query("action_type")
	to_user_id := c.Query("to_user_id")
	if res, _ := service.GetToken(token); res==0 {
		from_user_id :=dao.GetIdByUserId(strings.Split(token, ":")[0])
		if from_user_id==to_user_id{
			c.JSON(http.StatusOK, api.Response{StatusCode: 1, StatusMsg: "Can not follow self"})
			return
		}
		service.AddNewFollowRelation(from_user_id,to_user_id,action_type)
		c.JSON(http.StatusOK, api.Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, api.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// FollowList all users have same follow list
func FollowList(c *gin.Context) {
	id :=c.Query("user_id")

	userIdList:=service.GetFollowList(id)
	userList:=dao.GetUserListByIdArray(userIdList)
	res :=transUserList(userList)
	c.JSON(http.StatusOK, UserListResponse{
		Response: api.Response{
			StatusCode: 0,
		},
		UserList: res,
	})
}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {
	id :=c.Query("user_id")

	userIdList:=service.GetFollowerList(id)
	userList:=dao.GetUserListByIdArray(userIdList)
	res :=transUserList(userList)
	c.JSON(http.StatusOK, UserListResponse{
		Response: api.Response{
			StatusCode: 0,
		},
		UserList: res,
	})
}

func transUserList(userList []entities.User)[]api.User  {
	var res []api.User
	for i:=0;i<len(userList);i++{
		res= append(res, api.User{
			Id:            userList[i].Id,
			Name:          userList[i].Name,
			FollowCount:   userList[i].FollowCount,
			FollowerCount: userList[i].FollowerCount,
			IsFollow:      userList[i].IsFollow,
		})
	}
	return res
}
