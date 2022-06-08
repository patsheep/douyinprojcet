package main

import (
	"github.com/gin-gonic/gin"
	"github.com/RaymondCode/simple-demo/src/dao"

)

func main() {

	dao.ConnectDb()
	r := gin.Default()

	initRouter(r)

	r.Run(":8089") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
