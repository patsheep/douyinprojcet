package dao

import (
	"errors"
	"fmt"
	"github.com/RaymondCode/simple-demo/entities"
	"github.com/RaymondCode/simple-demo/util"
	"gorm.io/gorm"
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
func GetUserByIdAndPassword(userId string, password string) (*entities.User2, error) {
	//检查用户名是否存在
	var user entities.User2
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
	fmt.Println("runhere")
	return &user, nil
}

func AddNewUser(user2 entities.User2) error {
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
