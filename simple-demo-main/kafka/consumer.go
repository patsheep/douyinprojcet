package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/patsheep/douyinproject/config"
	"github.com/patsheep/douyinproject/dao"
	"github.com/patsheep/douyinproject/service"
	"github.com/patsheep/douyinproject/util"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	wg sync.WaitGroup //同步等待组
	//在类型上，它是一个结构体。一个WaitGroup的用途是等待一个goroutine的集合执行完成。
	//主goroutine调用了Add()方法来设置要等待的goroutine的数量。
	//然后，每个goroutine都会执行并且执行完成后调用Done()这个方法。
	//与此同时，可以使用Wait()方法来阻塞，直到所有的goroutine都执行完成。
)

func RunConsumer() {
	//获取消费者对象 可以设置多个IP地址和端口号，使用逗号进行分割
	consumer, err := sarama.NewConsumer(strings.Split("localhost:9092", ","), nil)
	//获取失败
	if err != nil {
		fmt.Println("Failed to start consumer: %s", err)
		return
	}
	//对该topic进行监听
	partitionList, err := consumer.Partitions("go_kafka")
	if err != nil {
		fmt.Println("Failed to get the list of partitions: ", err)
		return
	}
	//打印分区
	fmt.Println(partitionList)
	//获取分区和片偏移
	for partition := range partitionList {
		pc, err := consumer.ConsumePartition("go_kafka", int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("Failed to start consumer for partition %d: %s\n", partition, err)
			return
		}
		//延迟执行
		defer pc.AsyncClose()
		//启动多线程
		go func(pc sarama.PartitionConsumer) {
			wg.Add(1)
			//获得message的信息
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d, Offset:%d, Key:%s, Value:%s", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
				fmt.Println()
				list := strings.Split(string(msg.Value), " ")
				finalName := list[0]

				key, _ := strconv.ParseInt(list[1], 10, 64)
				if err != nil {
					fmt.Println("bug")
				}
				util.GetSnapshot(config.PROJECTPATH+config.VIDEO_ADDR+finalName, config.PROJECTPATH+config.COVER_ADDR+finalName[0:len(finalName)-4], 10)
				fmt.Println("封面图生成完成")
				publishToDB(finalName, key)
				fmt.Println("审核完成#")
			}
			wg.Done()
		}(pc)
	}
	//线程休眠
	time.Sleep(10 * time.Second)
	wg.Wait()
	consumer.Close()
}
func Consumerget() {
	//获取消费者对象 可以设置多个IP地址和端口号，使用逗号进行分割
	consumer, err := sarama.NewConsumer(strings.Split("localhost:9092", ","), nil)
	//获取失败
	if err != nil {
		fmt.Println("Failed to start consumer: %s", err)
		return
	}
	//对该topic进行监听
	partitionList, err := consumer.Partitions("go_kafka")
	if err != nil {
		fmt.Println("Failed to get the list of partitions: ", err)
		return
	}
	//打印分区
	fmt.Println(partitionList)
	//获取分区和片偏移
	for partition := range partitionList {
		pc, err := consumer.ConsumePartition("go_kafka", int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("Failed to start consumer for partition %d: %s\n", partition, err)
			return
		}
		//延迟执行
		defer pc.AsyncClose()
		//启动多线程
		go func(pc sarama.PartitionConsumer) {
			wg.Add(1)
			//获得message的信息
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d, Offset:%d, Key:%s, Value:%s", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
				fmt.Println()

			}

			wg.Done()
		}(pc)
	}
	//线程休眠
	time.Sleep(10 * time.Second)
	wg.Wait()
	consumer.Close()

}

func publishToDB(filename string, id int64) {
	/*	vd :=entities.Video2{
		AuthorId:      id,
		PlayUrl:       "http://pathcystore.oss-cn-shanghai.aliyuncs.com/video/"+filename,
		CoverUrl:      "http://pathcystore.oss-cn-shanghai.aliyuncs.com/cover/"+filename[0:len(filename)-4]+".jpeg",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
	}*/
	dao.Db.Table("video").Where("id", id).Updates(map[string]interface{}{"play_url": config.CONFIG.OssConfig.Endpoint + "/video/" + filename, "cover_url": config.CONFIG.OssConfig.Endpoint + "/cover/" + filename[0:len(filename)-4] + ".jpeg"})
	var wg sync.WaitGroup
	wg.Add(2)
	go service.UploadFile(filename, id, &wg)                     //向OSS上传文件
	go service.UploadCover(filename[0:len(filename)-4], id, &wg) //向OSS上传封面
	wg.Wait()

}
