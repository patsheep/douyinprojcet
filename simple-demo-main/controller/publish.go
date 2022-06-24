package controller

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/dao"
	"github.com/RaymondCode/simple-demo/entities"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/RaymondCode/simple-demo/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"sync"
)

type VideoListResponse struct {
	entities.Response
	VideoList []entities.Video `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {

	token := c.PostForm("token")

	if _, exist := usersLoginInfo[token]; !exist {
		c.JSON(http.StatusOK, entities.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, entities.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	filename := filepath.Base(data.Filename)
	fmt.Println(filename+"!")
	user := usersLoginInfo[token]
	finalName := fmt.Sprintf("%d_%s", user.Id, filename)

	fmt.Println(filename+"!! "+user.Name);
	saveFile := filepath.Join("./video/", finalName)

	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, entities.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	util.GetSnapshot("D:\\douyin\\douyinprojcet\\simple-demo-main\\video\\"+finalName,"D:\\douyin\\douyinprojcet\\simple-demo-main\\cover\\"+finalName[0:len(finalName)-4],10)
	defer publishToDB(finalName,user.Id)




	c.JSON(http.StatusOK, entities.Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	c.JSON(http.StatusOK, VideoListResponse{
		Response: entities.Response{
			StatusCode: 0,
		},
		VideoList: entities.DemoVideos,
	})
}

func publishToDB(filename string,id int64){
		vd :=entities.Video2{
			AuthorId:      id,
			PlayUrl:       "http://pathcystore.oss-cn-shanghai.aliyuncs.com/video/"+filename,
			CoverUrl:      "http://pathcystore.oss-cn-shanghai.aliyuncs.com/cover/"+filename[0:len(filename)-4]+".jpeg",
			FavoriteCount: 0,
			CommentCount:  0,
			IsFavorite:    false,
		}
		dao.Db.Table("video").Create(&vd)
		var wg sync.WaitGroup
		wg.Add(2)
		go service.UploadFile(filename,&wg)//向OSS上传文件
		go service.UploadCover(filename[0:len(filename)-4],&wg)//向OSS上传封面

		wg.Wait()



}
