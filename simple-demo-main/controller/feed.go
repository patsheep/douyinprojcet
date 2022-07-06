package controller

import (
	"github.com/RaymondCode/simple-demo/api"
	"github.com/RaymondCode/simple-demo/dao"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type FeedResponse struct {
	api.Response
	VideoList []api.Video `json:"video_list,omitempty"`
	NextTime  int64       `json:"next_time,omitempty"`

}

// Feed same demo video list for every request
func Feed(c *gin.Context) {

	c.JSON(http.StatusOK, FeedResponse{
		Response: api.Response{StatusCode: 0},
		//VideoList: DemoVideos,
		VideoList: dao.GetList(),
		NextTime:  time.Now().Unix(),

	})
}
