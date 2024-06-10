package logic

import (
	"bullmoon/dao/redis"
	"bullmoon/models"
	"strconv"

	"go.uber.org/zap"
)

// 再次使用简易投票算法
// 432分   86400/200 -> 200票可以让时间戳分数增加一天，续一天展示

/*
投票的几种情况：
direction=1时，有两种情况
	1.之前没有投过票，现在投赞成票
	2.之前投过反对票，现在改投赞成票
direction=0时，有两种情况
	1.之前投过赞成票，现在取消投票
	2.之前投过反对票，现在取消投票
direction=-1时，有两种情况
	1.之前没有投过票，现在投反对票
	2.之前投过反对票，现在改投反对票

投票的限制：
每个帖子自发表日起一个星期允许用户投票，超过一个星期禁止投票
	1.到期之后将redis中的保存赞成及反对票数存储到MySQL
	2.到期之后删除  KeyPostVotedZSetPF

*/

func PostVote(userid int64, p *models.ParamVotedData) (err error) {
	zap.L().Debug("logic.PostVote",
		zap.Int64("userID", userid),
		zap.String("postID", p.PostID),
		zap.Int8("direction", p.Direction))
	return redis.VoteForPost(strconv.Itoa(int(userid)), p.PostID, float64(p.Direction))
}
