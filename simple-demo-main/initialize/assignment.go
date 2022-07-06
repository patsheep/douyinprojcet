package initialize

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/config"
	"github.com/RaymondCode/simple-demo/dao"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/RaymondCode/simple-demo/util/snowflake"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func AssignMent()  {
	Viper()
	dao.InitDb()
	service.InitRdsPool()
	service.InitOss()
	snowflake.InitSnowFlakeSeed()
	config.PROJECTPATH=GetRunPath2()

}
//获取程序执行目录
func GetRunPath2() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))
	ret := path[:index]
	fmt.Println("视频存储url为"+ret+config.VIDEO_ADDR)
	return ret
}

