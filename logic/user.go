package logic

import (
	"bullmoon/dao/mysql"
	"bullmoon/models"
	"bullmoon/pkg/jwt"
	"bullmoon/pkg/snowflake"
)

// 存放业务逻辑的代码
func SignUp(p *models.ParamSignUp) (err error) {
	// 1.判断用户是否存在
	err = mysql.CheckUserExist(p.UserName)
	if err != nil {
		return err
	}
	// 2.生成雪花id
	id := snowflake.GenID()
	user := &models.User{
		UserId:   id,
		Username: p.UserName,
		Password: p.Password,
	}
	// 3.将用户存放进数据库
	err = mysql.InsertUser(user)
	if err != nil {
		return err
	}
	return
}

func Login(p *models.ParamLogin) (user *models.User, err error) {
	user = new(models.User)
	user.Username = p.UserName
	user.Password = p.Password
	//数据库查询
	if err := mysql.LoginCheck(user); err != nil {
		return nil, err
	}

	token, err := jwt.GenToken(user.UserId, user.Username)
	if err != nil {
		return nil, err
	}
	user.Token = token
	return user, nil
}
