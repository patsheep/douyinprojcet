package controller

import (
	"github.com/RaymondCode/simple-demo/api"
	"github.com/RaymondCode/simple-demo/dao"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type CommentListResponse struct {
	api.Response
	CommentList []api.Comment `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	api.Response
	Comment api.Comment `json:"comment,omitempty"`
}

const DATE_TIME_FORMAT  = "2006-01-02 15:04:05"
// CommentAction no practical effect, just check if token is valid
//三个参数：token,video_id,action_type(1-发布，2-删除),comment_text
func CommentAction(c *gin.Context) {

	token := c.Query("token")
	actionType := c.Query("action_type")

	if res, _ := service.GetToken(token); res==0 {
		if actionType == "1" {
			user_id :=strings.Split(token, ":")[0]

			text := c.Query("comment_text")
			video_id :=c.Query("video_id")
			commentRes:=dao.InsertAndGetComment(user_id,text,video_id)
			user := dao.GetUserById(user_id)
			c.JSON(http.StatusOK, CommentActionResponse{Response: api.Response{StatusCode: 0},
				Comment: api.Comment{
					Id:         commentRes.ID,
					User:       api.User{
						Id:            user.Id,
						Name:          user.Name,
						FollowCount:   user.FollowCount,
						FollowerCount: user.FollowerCount,
						IsFollow:      user.IsFollow,
					},
					Content:    commentRes.Content,
					CreateDate: commentRes.CreateTime.Format(DATE_TIME_FORMAT),
				}})
			return
		}
		c.JSON(http.StatusOK, api.Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, api.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// CommentList all videos have same demo comment list
//token和video_id
func CommentList(c *gin.Context) {

	videoId :=c.Query("video_id")
	tempList:=dao.GetVideoCommentList(videoId)
	var res []api.Comment
	for i:=0;i< len(tempList);i++{
		user := dao.GetUserByUserIdInt64(tempList[i].UserID)
		res=append(res, api.Comment{
			Id:         tempList[i].ID,
			User:       api.User{
				Id:            user.Id,
				Name:          user.Name,
				FollowCount:   user.FollowCount,
				FollowerCount: user.FollowerCount,
				IsFollow:      user.IsFollow,
			},
			Content:    tempList[i].Content,
			CreateDate: tempList[i].CreateTime.Format(DATE_TIME_FORMAT),
		})
	}
	c.JSON(http.StatusOK, CommentListResponse{
		Response:    api.Response{StatusCode: 0},
		CommentList: res,
	})
}
