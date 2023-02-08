package initialize

import (
	"github.com/patsheep/douyinproject/config"
	"github.com/spf13/viper"
	"log"
)

func Viper() {
	//viper.SetConfigType("yaml") // 如果配置文件的名称中没有扩展名，则需要配置此项
	//viper.AddConfigPath("./vipertest/") // 设置读取路径：就是在此路径下搜索配置文件。
	var conf = config.SystemConf{}

	//viper.AddConfigPath("$HOME/.appname")  // 多次调用以添加多个搜索路径
	viper.SetConfigFile("./config/application.yaml") // 设置被读取文件的全名，包括扩展名。
	//viper.SetConfigName("server") // 设置被读取文件的名字： 这个方法 和 SetConfigFile实际上仅使用一个就够了
	viper.ReadInConfig() // 读取配置文件： 这一步将配置文件变成了 Go语言的配置文件对象包含了 map，string 等对象。

	err := viper.Unmarshal(&conf)
	config.CONFIG = conf
	if err != nil {
		log.Panic("viper反序列化错误")
	}

	//	fmt.Printf("%+v",conf)

	// 控制台输出： map[first:panda last:8z] 99 panda [Coding Movie Swimming]
}
