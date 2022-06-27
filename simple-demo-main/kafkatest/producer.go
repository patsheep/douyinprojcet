package kafkatest

import (
	"fmt"
	"github.com/Shopify/sarama"
	"strconv"
	"time"
)

//消息生产者
func RunProducer() {
	//获取配置类
	config := sarama.NewConfig()                              //配置类实例（指针类型）
	config.Producer.RequiredAcks = sarama.WaitForAll          //代理需要的确认可靠性级别(默认为WaitForLocal)
	config.Producer.Partitioner = sarama.NewRandomPartitioner //生成用于选择要发送消息的分区的分区(默认为散列消息键)。
	config.Producer.Return.Successes = true                   //如果启用，成功传递的消息将在成功通道(默认禁用)。
	//获取客户端对象
	client, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		//获取客户端失败
		fmt.Println("producer close, err:", err)
		return
	}
	//延迟执行，类似于栈，等到其他代码都执行完毕后再执行
	defer client.Close()
	//一直循环
	for {
		//获取Message对象
		msg := &sarama.ProducerMessage{}
		//设置topic
		msg.Topic = "go_kafka"
		//设置Message值
		msg.Value = sarama.StringEncoder("this is a good test, my message is good")
		//发送消息，返回pid、片偏移
		pid, offset, err := client.SendMessage(msg)
		//发送失败
		if err != nil {
			fmt.Println("send message failed,", err)
			return
		}
		//打印返回结果
		fmt.Printf("pid:%v offset:%v\n", pid, offset)
		//线程休眠下
		time.Sleep(10 * time.Second)
	}
}
func ProducerSend(str string, key int64) {
	//获取配置类
	config := sarama.NewConfig()                              //配置类实例（指针类型）
	config.Producer.RequiredAcks = sarama.WaitForAll          //代理需要的确认可靠性级别(默认为WaitForLocal)
	config.Producer.Partitioner = sarama.NewRandomPartitioner //生成用于选择要发送消息的分区的分区(默认为散列消息键)。
	config.Producer.Return.Successes = true                   //如果启用，成功传递的消息将在成功通道(默认禁用)。
	//获取客户端对象
	client, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		//获取客户端失败
		fmt.Println("producer close, err:", err)
		return
	}
	//延迟执行，类似于栈，等到其他代码都执行完毕后再执行
	defer client.Close()
	//一直循环

	//获取Message对象
	msg := &sarama.ProducerMessage{}
	//设置topic
	msg.Topic = "go_kafka"
	//设置Message值

	msg.Value = sarama.StringEncoder(str + " " + strconv.FormatInt(key, 10))
	//发送消息，返回pid、片偏移
	pid, offset, err := client.SendMessage(msg)
	//发送失败
	if err != nil {
		fmt.Println("send message failed,", err)
		return
	}
	//打印返回结果
	fmt.Printf("pid:%v offset:%v\n", pid, offset)
	//线程休眠下

}
