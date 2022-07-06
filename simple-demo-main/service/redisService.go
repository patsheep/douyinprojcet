package service

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/config"
	"github.com/garyburd/redigo/redis"
	"strconv"
	"strings"
	"time"
)
var pool *redis.Pool
/*func init(){

	pool=&redis.Pool{
		MaxIdle: 8,
		MaxActive: 0,
		IdleTimeout: 100,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp",
				"localhost:6379",

				redis.DialDatabase(int(0)))
		},
	}
	fmt.Println("redis连接池建立完成")
}*/

func InitRdsPool(){
	rdsPool:=&redis.Pool{
		MaxIdle: 8,
		MaxActive: 0,
		IdleTimeout: 100,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp",
				fmt.Sprintf("%s:%d", config.CONFIG.RedisConfig.Host, config.CONFIG.RedisConfig.Port),

				redis.DialDatabase(int(0)))
		},
	}
	pool=rdsPool
}
func GenerateToken(user_id string)(string,error){
	conn:=pool.Get()
	defer conn.Close()
	nowtime :=time.Now().Unix()
	TokenVal := user_id+":"+strconv.FormatInt(nowtime, 10)
	_, err := conn.Do("Set", "Token:"+user_id, TokenVal)
	if err!=nil{
		return "",err
	}
	return TokenVal,nil

}

func SetToken(user_id string){
	conn:=pool.Get()
	defer conn.Close()
	nowtime :=time.Now().Unix()
	_, err := conn.Do("Set", "Token:"+user_id, user_id+"_"+strconv.FormatInt(nowtime, 10))
	if err!=nil{
		return
	}
/*	_, err = conn.Do("Expire",  "Token:"+user_id,100)
	if err!=nil{
		return
	}*/
}
//从redis获取token
//返回值含义：
//0.成功
//1.账号在其他地方登录（找到token但时间不对）
//2.请重新登录（token已经过期）
func GetToken(token string)(int,error)  {
	userid:=strings.Split(string(token), ":")[0]
	conn:=pool.Get()
	defer conn.Close()
	res,err :=redis.String(conn.Do("Get","Token:"+userid))
	if err!=nil{

		fmt.Println("登录已经过期")
		fmt.Println(err)
		return 2,err

	}
	if res==""{
		fmt.Println("登录已经过期")
		return 2,err
	}
	if res!=token{
		return 1,err
	}
	return  0,err
}

func AddFavorite(videoId,userId,favoriteType string){
	coon:=pool.Get()
	defer coon.Close()
	if favoriteType=="1"{
		redis.String(coon.Do("HINCRBY","userlikevideo",videoId+":"+userId,1))
	   redis.String(coon.Do("HINCRBY","videofavorite",videoId,1))
	   redis.String(coon.Do("LPUSH","FavoriteList:"+userId,videoId))
	} else{
		redis.String(coon.Do("HINCRY","Favorite:"+videoId,-1))
		redis.String(coon.Do("HINCRBY","userlikevideo",videoId+":"+userId,-1))
	}

}

func GetFavoriteVideoIdList(userid string)(*[]string) {
	coon:=pool.Get()
	defer coon.Close()
	var res []string

	len,err:=redis.Int(coon.Do("LLEN","FavoriteList:"+userid))
	if err!=nil{
		fmt.Println("bug")
	}
	for i:=0;i<len ;i++  {
		vdId,_:=redis.String(coon.Do("LINDEX","FavoriteList:"+userid,i))
		isActive,_:=redis.String(coon.Do("HGET","userlikevideo",vdId+":"+userid))
		fmt.Println(vdId+" !"+isActive)
		if isActive=="1"{
			res=append(res, vdId)
		}
	}

	return &res
}

func AddNewFollowRelation(from_id,to_id,act_type string){
	coon:=pool.Get()
	defer coon.Close()
	if act_type=="1"{
		nowtime:=time.Now().Unix()
		//用redis.bool会报错
		_,err:=redis.Bool(coon.Do("ZADD","FollowerList:"+from_id,nowtime,to_id))
		if err!=nil{
			fmt.Println(err)
		}
		_,err=redis.Bool(coon.Do("ZADD","FollowList:"+to_id,nowtime,from_id))
		if err!=nil{
			fmt.Println(err)
		}
	}else {
		_,err:=redis.Bool(coon.Do("ZREM","FollowerList:"+from_id,to_id))
		if err!=nil{
			fmt.Println(err)
		}
		_,err=redis.Bool(coon.Do("ZREM","FollowList:"+to_id,from_id))
		if err!=nil{
			fmt.Println(err)
		}
	}


}

func GetFollowList(id string)[]int64{
	coon:=pool.Get()
	defer coon.Close()
	list,err:=redis.Int64s(coon.Do("ZRANGE","FollowList:"+id,0,-1))
	if err!=nil{
		fmt.Println(err)
	}
	return list
}

func GetFollowerList(id string)[]int64{
	coon:=pool.Get()
	defer coon.Close()
	list,err:=redis.Int64s(coon.Do("ZRANGE","FollowerList:"+id,0,-1))
	if err!=nil{
		fmt.Println(err)
	}
	return list
}

