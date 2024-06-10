package logic

import (
	"bullmoon/dao/mysql"
	"bullmoon/dao/redis"
	"bullmoon/models"
	"bullmoon/pkg/snowflake"

	"go.uber.org/zap"
)

func CreatePost(p *models.Post) (err error) {
	// 获取post_id
	p.ID = snowflake.GenID()
	// 将post存入数据库
	err = mysql.CreatePost(p)
	if err != nil {
		return err
	}
	err = redis.CreatePost(p.ID)
	return
}

// GetPostDetail 返回帖子详情
func GetPostDetail(pid int64) (data *models.ApiPostDetail, err error) {
	// 根据数据拼接并返回想要的数据
	post, err := mysql.GetPostDetailByID(pid)
	if err != nil {
		zap.L().Error("mysql.GetPostDetailByID(pid) failed", zap.Int64("pid", pid), zap.Error(err))
		return
	}
	// 根据userid查询用户
	user, err := mysql.GetUserByID(post.AuthorID)
	if err != nil {
		zap.L().Error("mysql.GetUserByID(post.AuthorID)", zap.Int64("uid", post.AuthorID), zap.Error(err))
		return
	}
	// 根据community_id查询信息
	community, err := mysql.GetCommunityDetailById(post.CommunityID)
	if err != nil {
		zap.L().Error("mysql.GetCommunityDetailById(post.CommunityID)", zap.Int64("cid", post.CommunityID), zap.Error(err))
		return
	}
	data = &models.ApiPostDetail{
		AuthorName:      user.Username,
		Post:            post,
		CommunityDetail: community,
	}
	return
}

// GetPostList 查询帖子列表
func GetPostList(page int64, size int64) (posts []*models.ApiPostDetail, err error) {
	data, err := mysql.GetPostList(page, size)
	if err != nil {
		zap.L().Error("mysql.GetPostList() failed", zap.Error(err))
		return nil, err
	}
	posts = make([]*models.ApiPostDetail, 0, len(data))
	for _, post := range data {
		user, err := mysql.GetUserByID(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserByID(post.AuthorID)", zap.Int64("uid", post.AuthorID), zap.Error(err))
			continue
		}
		// 根据community_id查询信息
		community, err := mysql.GetCommunityDetailById(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetCommunityDetailById(post.CommunityID)", zap.Int64("cid", post.CommunityID), zap.Error(err))
			continue
		}
		apiPost := &models.ApiPostDetail{
			AuthorName:      user.Username,
			Post:            post,
			CommunityDetail: community,
		}
		posts = append(posts, apiPost)
	}
	return
}

func GetPostList2(p *models.ParamPostList) (posts []*models.ApiPostDetail, err error) {
	// 去redis查询id列表
	ids, err := redis.GetPostIDsInOrder(p)
	if err != nil {
		return
	}
	data, err := mysql.GetPostListByID(ids)
	posts = make([]*models.ApiPostDetail, 0, len(data))

	voteData, err := redis.GetPostVoteData(ids)
	if err != nil {
		return nil, err
	}

	for idx, post := range data {
		user, err := mysql.GetUserByID(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserByID(post.AuthorID)", zap.Int64("uid", post.AuthorID), zap.Error(err))
			continue
		}
		// 根据community_id查询信息
		community, err := mysql.GetCommunityDetailById(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetCommunityDetailById(post.CommunityID)", zap.Int64("cid", post.CommunityID), zap.Error(err))
			continue
		}
		apiPost := &models.ApiPostDetail{
			AuthorName:      user.Username,
			VoteNum:         voteData[idx],
			Post:            post,
			CommunityDetail: community,
		}
		posts = append(posts, apiPost)
	}
	return
}
