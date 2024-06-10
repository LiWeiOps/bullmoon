package redis

import (
	"context"
	"errors"
	"fmt"
	"math"
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	oneWeekInSeconds         = 7 * 24 * 3600
	scorePreVote     float64 = 432
)

var (
	ErrVoteTimeExpire = errors.New("投票时间已过")
	ErrVoteRepested   = errors.New("不允许重复投票")
)

func CreatePost(postID int64) (err error) {
	ctx := context.Background()
	pipeline := rdb.TxPipeline()
	pipeline.ZAdd(ctx, getRedisKey(KeyPostTimeZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID},
	)

	rdb.ZAdd(ctx, getRedisKey(KeyPostScoreZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID},
	)
	_, err = pipeline.Exec(ctx)
	return
}

func VoteForPost(userID, postID string, value float64) (err error) {
	// 1.判断帖子是否在投票时间限制内
	// 去redis取帖子发布时间
	ctx := context.Background()
	postTime := rdb.ZScore(ctx, getRedisKey(KeyPostTimeZSet), postID).Val()
	fmt.Println("postTime:", postTime)
	if float64(time.Now().Unix())-postTime > oneWeekInSeconds {
		return ErrVoteTimeExpire
	}
	// 2和3放到一个pipeline中执行
	// 2.更新帖子分数
	// 先查询当前用户给当前帖子的投票记录
	ov := rdb.ZScore(ctx, getRedisKey(KeyPostVotedZSetPF)+postID, userID).Val()
	// 如果这一次投票与前一次一致，不允许重复投票
	if value == ov {
		return ErrVoteRepested
	}
	var op float64
	if value > ov {
		op = 1
	} else {
		op = -1
	}
	diff := math.Abs(ov - value) // 计算两次投票差值
	pipeline := rdb.Pipeline()
	pipeline.ZIncrBy(ctx, getRedisKey(KeyPostScoreZSet), op*diff*scorePreVote, postID)
	// 3.更新用户为该帖子的投票记录
	if value == 0 {
		// 如果取消投票就直接删除该用户投票记录
		pipeline.ZRem(ctx, getRedisKey(KeyPostVotedZSetPF)+postID, userID)
	} else {
		pipeline.ZAdd(ctx, getRedisKey(KeyPostVotedZSetPF)+postID, redis.Z{
			Score:  value,
			Member: userID,
		})
	}
	_, err = pipeline.Exec(ctx)
	return
}
