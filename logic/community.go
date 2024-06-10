package logic

import (
	"bullmoon/dao/mysql"
	"bullmoon/models"
)

func GetCommunityList() (data []*models.Community, err error) {
	data, err = mysql.GetCommunityList()
	// 这里是展示社区列表逻辑，即使数据库为空，返回空列表即可
	//if err != nil {
	//	return nil, err
	//}
	return
}

func GetCommunityDetail(id int64) (data *models.CommunityDetail, err error) {
	data, err = mysql.GetCommunityDetailById(id)
	return
}
