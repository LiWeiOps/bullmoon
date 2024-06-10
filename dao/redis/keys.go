package redis

// redis key
// redis key注意使用命名空间的方式，方便查询和拆分

const (
	KeyPrefix          = "bullmoon:"
	KeyPostTimeZSet    = "post:time"   // zset;帖子及发帖时间
	KeyPostScoreZSet   = "post:score"  // zset;帖子及投票分数
	KeyPostVotedZSetPF = "post:voted:" // zset;记录用户及其投票类型 参数是post_id
)

func getRedisKey(key string) string {
	return KeyPrefix + key
}
