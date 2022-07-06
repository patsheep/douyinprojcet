package dao

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/api"
	"github.com/RaymondCode/simple-demo/config"
	"github.com/RaymondCode/simple-demo/entities"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"strconv"
	"time"
)

var Db *gorm.DB

/*func init() {
	var (
		err error
	)
	Db, err = gorm.Open(mysql.Open("root:97782078@tcp(127.0.0.1:3306)/douyin?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

}*/

func InitDb() {
	username := config.CONFIG.MySQLConfig.Username// 账号
	password := config.CONFIG.MySQLConfig.Password // 密码
	host := config.CONFIG.MySQLConfig.Host         // 数据库地址，可以是Ip或者域名
	port := config.CONFIG.MySQLConfig.Port         // 数据库端口
	dbName := config.CONFIG.MySQLConfig.DBname     // 数据库名
	// dsn := "用户名:密码@tcp(地址:端口)/数据库名"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=True&loc=Local", username, password, host, port, dbName)

	// 配置Gorm连接到MySQL
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN
		DefaultStringSize:         256,   // string 类型字段的默认长度
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Millisecond * 0, // 慢 SQL 阈值
			LogLevel:                  logger.Info,          // 日志级别
			IgnoreRecordNotFoundError: true,                 // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  false,                // 禁用彩色打印
		},
	)
	if db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		Logger: newLogger,
	}); err == nil {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxOpenConns(config.CONFIG.MySQLConfig.MaxOpenConns) // 设置数据库最大连接数
		sqlDB.SetMaxIdleConns(config.CONFIG.MySQLConfig.MaxIdleConns) // 设置上数据库最大闲置连接数
		Db=db
	} else {
		panic("connect server failed")
	}

}


func GetList() []api.Video {
	//node := entities.Tbs{Val: 1}
	var lists []entities.Video

	Db.Table("video").Where("id > ?", 0).Order("id desc").Limit(30).Debug().Find(&lists)
	var resLists []api.Video
	var userIdList []int64

	for i := 0; i < len(lists); i++ {
		userIdList = append(userIdList, lists[i].AuthorId)
	}
	for i := 0; i < len(userIdList); i++ {
		fmt.Println(strconv.Itoa(i) + " " + strconv.FormatInt(userIdList[i], 10))
	}
	var temp []entities.User
	tempmap := map[int64]entities.User{}
	Db.Table("user").Where("id In ?", userIdList).Debug().Find(&temp)
	for i := 0; i < len(temp); i++ {
		fmt.Println("", temp[i])

		tempmap[temp[i].Id] = temp[i]

	}
	for key, val := range tempmap {
		fmt.Println(key, val)
	}
	for i := 0; i < len(lists); i++ {

		var userNode entities.User
		value, ok := tempmap[lists[i].AuthorId]
		if ok {
			userNode = value
		} else {
			fmt.Println("去数据库中查找")
			Db.Table("user").Where("id = ?", lists[i].AuthorId).Debug().Find(&userNode)
		}

		resuserNode := api.User{
			Id:            userNode.Id,
			Name:          userNode.Name,
			FollowCount:   userNode.FollowCount,
			FollowerCount: userNode.FollowerCount,
			IsFollow:      userNode.IsFollow,
		}
		resLists = append(resLists, api.Video{
			Id:            lists[i].Id,
			Author:        resuserNode,
			PlayUrl:       lists[i].PlayUrl,
			CoverUrl:      lists[i].CoverUrl,
			FavoriteCount: lists[i].FavoriteCount,
			CommentCount:  lists[i].CommentCount,
			IsFavorite:    lists[i].IsFavorite,
		})

	}
	for i := 0; i < len(resLists); i++ {
		fmt.Print("%v+\n", resLists[i])
	}
	return resLists

}

func GetListById(id int64) []api.Video {
	fmt.Println("runninggetlistById")
	//node := entities.Tbs{Val: 1}
	var lists []entities.Video

	Db.Table("video").Where("author_id = ?", id).Order("id desc").Limit(30).Debug().Find(&lists)
	var resLists []api.Video
	var userNode entities.User
	Db.Table("user").Where("id = ?", id).Debug().Find(&userNode)
	for i := 0; i < len(lists); i++ {
		resuserNode := api.User{
			Id:            userNode.Id,
			Name:          userNode.Name,
			FollowCount:   userNode.FollowCount,
			FollowerCount: userNode.FollowerCount,
			IsFollow:      userNode.IsFollow,
		}
		resLists = append(resLists, api.Video{
			Id:            lists[i].Id,
			Author:        resuserNode,
			PlayUrl:       lists[i].PlayUrl,
			CoverUrl:      lists[i].CoverUrl,
			FavoriteCount: lists[i].FavoriteCount,
			CommentCount:  lists[i].CommentCount,
			IsFavorite:    lists[i].IsFavorite,
		})

	}
	for i := 0; i < len(resLists); i++ {
		fmt.Print("%v+\n", resLists[i])
	}
	return resLists

}

type OSSKey struct {
	Key    string `form:"key"`
	Secret string `form:"secret"`
}

func GetOSSKEy() []string {
	var res OSSKey
	Db.Table("osskey").Debug().First(&res)
	return []string{res.Key, res.Secret}
}
