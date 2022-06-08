package entities

import "gorm.io/gorm"


type Tbs struct {
	gorm.Model
	Val int `form:"val"`

}

func (Tbs) TableName() string {
	return "User"
}

