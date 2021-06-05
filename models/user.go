package models

import (
	"fmt"
	"go-demo/dao"
)

// 用户模型
type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func GetUserList() (userList []*User, err error) {
	if err = dao.DB.Find(&userList).Error; err != nil {
		return nil, err
	}
	return
}

func GetUser(id string) (user *User, err error) {
	user = new(User)
	if err = dao.DB.Where("id=?", id).First(user).Error; err != nil {
		return nil, fmt.Errorf("no record of user id : %v, error is : %w", id, err)
	}
	return
}
