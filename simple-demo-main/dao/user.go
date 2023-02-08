package dao

import (
	"errors"
	"fmt"
	"github.com/patsheep/douyinproject/entities"
	"github.com/patsheep/douyinproject/util"
	"gorm.io/gorm"
	"strconv"
	"sync"
)

type UserDao struct{}

var (
	userDao  *UserDao
	userOnce sync.Once
)

func NewUserDaoInstance() *UserDao {
	userOnce.Do(func() {
		userDao = &UserDao{}
	})
	return userDao
}
func GetUserByIdAndPassword(userId string, password string) (*entities.User, error) {
	//检查用户名是否存在
	var user entities.User
	result := Db.Table("user").Where("user_id = ?", userId).Limit(1).Find(&user)
	if result.RowsAffected == 0 {
		err := errors.New("username does not exist")
		return nil, err
	}
	//检查密码是否正确

	if util.GetMd5Val(password) != user.Password {
		err := errors.New("wrong password")
		return nil, err
	}
	return &user, nil
}

func GetUserById(tokenString string) entities.User {
	var res entities.User
	Db.Table("user").Where("user_id = ?", tokenString).Limit(1).Find(&res)
	fmt.Printf("%+v\n", res)
	return res
}

func GetUserByIdInt64(userId int64) entities.User {
	var res entities.User
	id := strconv.FormatInt(userId, 10)
	Db.Table("user").Where("id = ?", id).Limit(1).Find(&res)

	return res
}

func GetIdByUserId(userid string) string {
	var res entities.User
	Db.Table("user").Where("user_id = ?", userid).Limit(1).Find(&res)

	return strconv.FormatInt(res.Id, 10)

}
func GetUserByUserIdInt64(userId int64) entities.User {
	var res entities.User
	id := strconv.FormatInt(userId, 10)
	Db.Table("user").Where("user_id = ?", id).Limit(1).Find(&res)
	return res

}
func AddNewUser(user2 entities.User) error {
	err := Db.Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		if err := tx.Table("user").Create(user2).Error; err != nil {
			// 返回任何错误都会回滚事务
			fmt.Println("create new user error")
			return err
		}

		// 返回 nil 提交事务
		return nil
	})
	return err
}

func GetUserListByIdArray(list []int64) []entities.User {
	var temp []entities.User

	Db.Table("user").Where("id In ?", list).Debug().Find(&temp)
	return temp
}

type node struct {
	Id       int
	Userid   int
	Username string
	Password string
	Salt     int
}

func AddUserAccount(val int, wg *sync.WaitGroup) {

	Db.Table("user_account").Create(&node{
		Userid:   val,
		Username: "",
		Password: "",
		Salt:     0,
	})
	defer wg.Done()
}
