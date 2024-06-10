package mysql

import (
	"bullmoon/models"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
)

var secret string = "wangliwei.com"

// 每一步数据库操作封装成函数
// 等待logic层调用
func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

func InsertUser(user *models.User) (err error) {
	// 对密码进行加密存储
	user.Password = encryptPassword(user.Password)
	// sql入库操作
	sqlStr := "insert into user(user_id, username, password) values(?,?,?)"
	_, err = db.Exec(sqlStr, user.UserId, user.Username, user.Password)
	if err != nil {
		return err
	}
	return
}

func CheckUserExist(username string) (err error) {
	sqlStr := "select count(user_id) from user where username=?"
	var count int
	err = db.Get(&count, sqlStr, username)
	if err != nil {
		return err
	}
	if count > 0 {
		return ErrorUserExist
	}
	return
}

func LoginCheck(user *models.User) (err error) {
	oPassword := user.Password
	sqlStr := "select user_id, username, password from user where username=?"
	err = db.Get(user, sqlStr, user.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			return ErrorUserNotExist
		}
		return fmt.Errorf("用户登录查询出错，err: %v", err)
	}

	password := encryptPassword(oPassword)
	if password != user.Password {
		return ErrorInvalidPassword
	}

	return
}

func GetUserByID(id int64) (data *models.User, err error) {
	data = new(models.User)
	sqlStr := "select user_id, username from user where user_id=?"
	err = db.Get(data, sqlStr, id)
	return
}
