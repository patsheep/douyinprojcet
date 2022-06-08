package dao

import (
	"github.com/RaymondCode/simple-demo/src/entities"

	//"fmt" // 导入 fmt 包，打印字符串是需要用到
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
var Db *gorm.DB

func ConnectDb() {
	var (
		err error
	)
	Db, err = gorm.Open(mysql.Open("root:97782078@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

}

func getList(){
	node := entities.Tbs{Val: 1}
	id:=Db.Table("tb").First(node);
	print(id)
}
