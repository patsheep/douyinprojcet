package controller

import (
	"fmt"
	"github.com/patsheep/douyinproject/api"
	"github.com/patsheep/douyinproject/dao"
	"github.com/patsheep/douyinproject/util/snowflake"

	"github.com/gin-gonic/gin"
	"github.com/patsheep/douyinproject/service"
	"net/http"
	"strings"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin
var usersLoginInfo = map[string]api.User{
	"zhangleidouyin": {
		Id:            1,
		Name:          "zhanglei",
		FollowCount:   10,
		FollowerCount: 5,
		IsFollow:      true,
	},
	"123123456": {
		Id:            1,
		Name:          "123",
		FollowCount:   0,
		FollowerCount: 0,
		IsFollow:      false,
	},
	"12345:1656903457": {
		Id:            1,
		Name:          "qwq",
		FollowCount:   0,
		FollowerCount: 0,
		IsFollow:      false,
	},
}

var userIdSequence = int64(1)

type UserLoginResponse struct {
	api.Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	api.Response
	User api.User `json:"user"`
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token := username + password

	if _, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: api.Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
	} else {
		userIdSequence := snowflake.MakeInt64SnowFlakeId()
		newUser := api.User{
			Id:   userIdSequence,
			Name: username,
		}
		usersLoginInfo[token] = newUser
		err := service.AddNewUser(username, password)
		if err != nil {
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: api.Response{StatusCode: 1, StatusMsg: "error please try again"},
			})
			return
		}
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: api.Response{StatusCode: 0},
			UserId:   userIdSequence,
			Token:    username + password,
		})
	}
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	fmt.Println(username + " " + password)
	if user, err := dao.GetUserByIdAndPassword(username, password); err == nil {

		token, _ := service.GenerateToken(username)

		c.JSON(http.StatusOK, UserLoginResponse{
			Response: api.Response{StatusCode: 0},
			UserId:   user.Id,
			Token:    token,
		})
	} else {

		c.JSON(http.StatusOK, UserLoginResponse{
			Response: api.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}

func UserInfo(c *gin.Context) {
	token := c.Query("token")
	user := dao.GetUserById(strings.Split(string(token), ":")[0])
	if exist, _ := service.GetToken(token); exist == 0 {
		c.JSON(http.StatusOK, UserResponse{
			Response: api.Response{StatusCode: 0},
			User: api.User{
				Id:            user.Id,
				Name:          user.Name,
				FollowCount:   0,
				FollowerCount: 0,
				IsFollow:      false,
			},
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: api.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}
