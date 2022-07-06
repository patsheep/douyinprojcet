package controller

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/api"
	"github.com/RaymondCode/simple-demo/config"
	"github.com/RaymondCode/simple-demo/dao"
	"github.com/RaymondCode/simple-demo/kafka"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/RaymondCode/simple-demo/util/snowflake"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
)

type VideoListResponse struct {
	api.Response
	VideoList []api.Video `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {

	token := c.PostForm("token")
	fmt.Println("tokenis:"+token)
	if exist,_ := service.GetToken(token); exist!=0 {
		c.JSON(http.StatusOK, api.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, api.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	filename := filepath.Base(data.Filename)
	fmt.Printf("文件大小为: %d 类型为 %s\n", data.Size, filename[len(filename)-3:])

	if filename[len(filename)-3:] != "mp4" {
		c.JSON(http.StatusOK, api.Response{
			StatusCode: 1,
			StatusMsg:  "不支持的文件类型",
		})
		return
	}
	user_id :=dao.GetIdByUserId(strings.Split(token, ":")[0])

	user_idInt64, _ := strconv.ParseInt(user_id,10,64)
	var key = snowflake.MakeInt64SnowFlakeId()
	dao.PublishTempToDB(user_idInt64,key)
	finalName := fmt.Sprintf("%s_%d_%s", user_id, key, filename)
	saveFile := config.PROJECTPATH+config.VIDEO_ADDR+finalName

	fmt.Println("saveFileDst"+saveFile)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, api.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	//defer publishToDB(finalName,user.Id)
	defer kafka.ProducerSend(finalName, key)

	c.JSON(http.StatusOK, api.Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {

	token := c.Query("token") //Query是获得get请求的参数
	fmt.Printf("token为%s\n",token)
	if exist,_ := service.GetToken(token); exist!=0 {
		c.JSON(http.StatusOK, api.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	userId:=strings.Split(string(token), ":")[0]
	Id,_ :=strconv.ParseInt(dao.GetIdByUserId(userId),10,64)
	fmt.Println(userId)
	c.JSON(http.StatusOK, VideoListResponse{
		Response: api.Response{
			StatusCode: 0,
		},
		VideoList: dao.GetListById(Id),
	})
}


