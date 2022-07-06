package dao

import (
	"database/sql"
	"fmt"
)

func PublishTempToDB(id,snowFlakeId int64) int64 { //上传临时数据，封面和动画地址默认。
	/*	vd :=entities.Video2{
			AuthorId:      id,
			PlayUrl:       "http://pathcystore.oss-cn-shanghai.aliyuncs.com/verifysource/verify.mp4",
			CoverUrl:      "http://pathcystore.oss-cn-shanghai.aliyuncs.com/verifysource/verify.jpg",
			FavoriteCount: 0,
			CommentCount:  0,
			IsFavorite:    false,
		}
		res:=dao.Db.Table("video").Create(&vd)*/
	//ConnectDb()

	db, err := sql.Open("mysql", "root:97782078@tcp(127.0.0.1:3306)/douyin?charset=utf8&parseTime=True&loc=Local")
	if err = db.Ping(); err != nil {
		fmt.Println("没连上")
	}
	res, err := db.Exec("INSERT INTO video(id,author_id,play_url,cover_url,favorite_count,comment_count,is_favorite) values (?,?,?,?,0,0,false)",snowFlakeId, id, "http://pathcystore.oss-cn-shanghai.aliyuncs.com/verifysource/verify.mp4", "http://pathcystore.oss-cn-shanghai.aliyuncs.com/verifysource/verify.jpg")
	if err != nil {
		fmt.Println("bug了")
	}
	lastIns, _ := res.LastInsertId()
	fmt.Println(lastIns)
	//res.Last
	fmt.Print("新插入的ID为: ")
	fmt.Println(lastIns)
	return lastIns

}

