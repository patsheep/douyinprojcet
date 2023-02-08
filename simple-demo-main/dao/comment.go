package dao

import (
	"fmt"
	"github.com/patsheep/douyinproject/entities"
	"github.com/patsheep/douyinproject/util/snowflake"
	"gorm.io/gorm"
	"strconv"
)

func InsertAndGetComment(userId, content, videoId string) entities.Comment {
	fmt.Println(userId + " " + content + " " + videoId)
	uid, _ := strconv.ParseInt(userId, 10, 64)
	vid, _ := strconv.ParseInt(videoId, 10, 64)
	idval := snowflake.MakeInt64SnowFlakeId()
	node := entities.Comment{
		ID:      idval,
		UserID:  uid,
		VideoID: vid,
		Content: content,
	}
	Db.Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		if err := tx.Table("comment").Create(&node).Error; err != nil {
			// 返回任何错误都会回滚事务
			fmt.Println("create new comment error")
			return nil
		}
		if err := tx.Table("comment").Where("id = ?", idval).Limit(1).Find(&node).Error; err != nil {
			// 返回任何错误都会回滚事务
			fmt.Println("create new comment error")
			return nil
		}

		// 返回 nil 提交事务
		return nil
	})
	fmt.Printf("%+v\n", node)
	return node

}

func GetVideoCommentList(videoId string) []entities.Comment {
	var lists []entities.Comment
	Db.Table("comment").Where("video_id = ?", videoId).Debug().Find(&lists)
	return lists
}
