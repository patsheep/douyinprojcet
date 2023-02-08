package dao

import (
	"fmt"
	"github.com/patsheep/douyinproject/entities"
	"strconv"
)

func GetFavoriteVideoList(idList []string) []entities.Video {
	var videoIdList []int64
	for i := 0; i < len(idList); i++ {
		val, _ := strconv.ParseInt(idList[i], 10, 64)
		videoIdList = append(videoIdList, val)
	}
	var temp []entities.Video

	Db.Table("video").Where("id In ?", videoIdList).Debug().Find(&temp)
	for i := 0; i < len(temp); i++ {
		fmt.Printf("%+d\n", temp[i])

	}
	return temp

}
