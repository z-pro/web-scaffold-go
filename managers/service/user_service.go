package service

import (
	"gim-master/logic/dao"
	"gim-master/logic/model"
)

type userService struct{}

var UserService = new(userService)

// Add 添加用户
func (*userService) Add(user model.User) (int64, error) {
	return dao.UserDao.Add(user)
}
