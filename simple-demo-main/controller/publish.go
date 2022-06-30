package controller

import (
	"database/sql"
	"fmt"
	"github.com/RaymondCode/simple-demo/dao"
	"github.com/RaymondCode/simple-demo/entities"
	"github.com/RaymondCode/simple-demo/kafkatest"
	"github.com/RaymondCode/simple-demo/service"
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
	fmt.Printf("文件大小为: %d 类型为 %s\n", data.Size, filename[len(filename)-3:])

	if filename[len(filename)-3:] != "mp4" {
		c.JSON(http.StatusOK, entities.Response{
			StatusCode: 1,
			StatusMsg:  "不支持的文件类型",
		})
		return
	}
	fmt.Println(filename + "!")
	user := usersLoginInfo[token]
	var key = publishTempToDB(user.Id)
	finalName := fmt.Sprintf("%d_%d_%s", user.Id, key, filename)
	fmt.Println(filename + "!! " + user.Name)
	saveFile := filepath.Join("./video/", finalName)

	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, entities.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	//defer publishToDB(finalName,user.Id)
	defer kafkatest.ProducerSend(finalName, key)

	c.JSON(http.StatusOK, entities.Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {

	token := c.Query("token") //Query是获得get请求的参数

	if _, exist := usersLoginInfo[token]; !exist {
		c.JSON(http.StatusOK, entities.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}
	user := usersLoginInfo[token]
	fmt.Println(user.Id)
	c.JSON(http.StatusOK, VideoListResponse{
		Response: entities.Response{
			StatusCode: 0,
		},
		VideoList: dao.GetListById(user.Id),
	})
}

func publishToDB(filename string, id int64) {
	vd := entities.Video2{
		AuthorId:      id,
		PlayUrl:       "http://pathcystore.oss-cn-shanghai.aliyuncs.com/video/" + filename,
		CoverUrl:      "http://pathcystore.oss-cn-shanghai.aliyuncs.com/cover/" + filename[0:len(filename)-4] + ".jpeg",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
	}
	dao.Db.Table("video").Create(&vd)
	var wg sync.WaitGroup
	wg.Add(2)
	go service.UploadFile(filename, &wg)                     //向OSS上传文件
	go service.UploadCover(filename[0:len(filename)-4], &wg) //向OSS上传封面
	wg.Wait()

}

func publishTempToDB(id int64) int64 { //上传临时数据，封面和动画地址默认。
	/*	vd :=entities.Video2{
			AuthorId:      id,
			PlayUrl:       "http://pathcystore.oss-cn-shanghai.aliyuncs.com/verifysource/verify.mp4",
			CoverUrl:      "http://pathcystore.oss-cn-shanghai.aliyuncs.com/verifysource/verify.jpg",
			FavoriteCount: 0,
			CommentCount:  0,
			IsFavorite:    false,
		}
		res:=dao.Db.Table("video").Create(&vd)*/
	//ConnectDb()
	db, err := sql.Open("mysql", "root:97782078@tcp(127.0.0.1:3306)/douyin?charset=utf8&parseTime=True&loc=Local")
	if err = db.Ping(); err != nil {
		fmt.Println("没连上")
	}
	res, err := db.Exec("INSERT INTO video(author_id,play_url,cover_url,favorite_count,comment_count,is_favorite) values (?,?,?,0,0,false)", id, "http://pathcystore.oss-cn-shanghai.aliyuncs.com/verifysource/verify.mp4", "http://pathcystore.oss-cn-shanghai.aliyuncs.com/verifysource/verify.jpg")
	if err != nil {
		fmt.Println("bug了")
	}
	lastIns, _ := res.LastInsertId()
	fmt.Println(lastIns)
	//res.Last
	fmt.Print("新插入的ID为: ")
	fmt.Println(lastIns)
	return lastIns

}
