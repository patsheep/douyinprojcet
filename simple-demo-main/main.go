package main

import (
	"github.com/gin-gonic/gin"
	"github.com/patsheep/douyinproject/initialize"
	"github.com/patsheep/douyinproject/kafka"
)

func main() {
	/*	fmt.Println(service.GetToken("weqwe"))
		time.Sleep(time.Second)
		fmt.Println(service.GetToken("weqwe"))*/

	initialize.AssignMent()

	// go kafkatest.RunProducer();
	go kafka.RunConsumer()

	//	dao.GetList()
	//util.GetSnapshot("D:\\CF\\stg.mp4","D:\\CF\\BEAR",20)

	r := gin.Default()

	initRouter(r)

	r.Run(":8089") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
