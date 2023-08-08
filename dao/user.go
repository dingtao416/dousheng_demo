package dao

import (
	"errors"
)

func IsUserExist(username string) error {
	var userlogin UserLogin
	// 能查询到数据，说明当前用户名存在，err 也不等于 nil
	if err := Db.Where("username=?", username).First(&userlogin).Error; err == nil {
		return errors.New("用户名已存在")
	}
	return nil
}

func AddUser(user *User) error {
	return Db.Create(user).Error
}

// 此处 login 应包含且只包含有效的 username 与 password 字段
func QueryUserLogin(login *UserLogin) error {
	/*
		当使用结构作为条件查询时，GORM 只会查询非零值字段。
		这意味着如果您的字段值为 0、''、false 或其他 零值，
		该字段不会被用于构建查询条件
	*/
	if err := Db.Where(login).First(login).Error; err != nil {
		return errors.New("用户不存在，账号或密码出错")
	}
	return nil
}

func QueryUserInfo(user *User) error {
	if err := Db.Where(user).First(user).Error; err != nil {
		return errors.New("用户不存在")
	}
	return nil
}
