package main

import (
	"github.com/RaymondCode/simple-demo/dao"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
)

func main() {

	dao.ConnectDb()
	service.OSSkeyinit()



//	dao.GetList()
	//util.GetSnapshot("D:\\CF\\stg.mp4","D:\\CF\\BEAR",20)

	r := gin.Default()

	initRouter(r)

	r.Run(":8089") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
