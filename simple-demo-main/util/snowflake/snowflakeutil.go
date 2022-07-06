package snowflake

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/config"
)

var SnowFlakeSeed *Node
func InitSnowFlakeSeed(){
	id :=config.CONFIG.SnowFlakeConfig.MechineId
	var err error
	SnowFlakeSeed,err=NewNode(id)
	if err!=nil{
		fmt.Println("errorcreatingSnowFlakeSeed")
	}else{
		fmt.Print("已生成雪花ID种子,当前编号为")
		fmt.Println(id)
	}
}

func MakeInt64SnowFlakeId() int64 {
	return SnowFlakeSeed.Generate().Int64()
}
