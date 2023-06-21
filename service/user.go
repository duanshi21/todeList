package service

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"sync"
	"todoList/pkg/ctl"
	"todoList/pkg/utils"
	"todoList/repository/db/dao"
	"todoList/repository/db/model"
	"todoList/types"
)

var UserSrvIns *UserSrv
var UserSrvOnce sync.Once

type UserSrv struct {
}

func GetUserSrv() *UserSrv {
	UserSrvOnce.Do(func() {
		UserSrvIns = &UserSrv{}
	})
	return UserSrvIns
}

// 注册
func (s *UserSrv) UserRegister(ctx context.Context, req *types.UserRegisterReq) (resp interface{}, err error) {
	userDao := dao.NewUserDao(ctx)
	u, err := userDao.FindUserByUserName(req.UserName)
	switch err {
	case gorm.ErrRecordNotFound:
		u = &model.User{
			Username: req.UserName,
		}
		// 密码进行加密存储
		if err = u.SetPassword(req.Password); err != nil {
			utils.LoggerObj.Info(err)
			return
		}
		if err = userDao.CreateUser(u); err != nil {
			utils.LoggerObj.Info(err)
			return
		}
		return ctl.RespSuccess(nil), nil
	case nil:
		err = errors.New("用户已存在")
		return
	default:
		return
	}
}

// 登录
func (s *UserSrv) UserLogin(ctx context.Context, req *types.UserLoginReq) (resp interface{}, err error) {
	// 1、找到该用户
	userDao := dao.NewUserDao(ctx)
	u, err := userDao.FindUserByUserName(req.UserName)
	if err != nil {
		// 2、说明没有该用户，直接return
		utils.LoggerObj.Error(err)
		return
	}
	// 3、校验密码
	if !u.CheckPassword(req.Password) {
		err = errors.New("账号/密码错误")
		utils.LoggerObj.Error(err)
		return
	}
	// 分发token

	return
}
