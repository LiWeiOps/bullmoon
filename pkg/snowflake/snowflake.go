package snowflake

import (
	"time"

	sf "github.com/bwmarrin/snowflake"
)

var node *sf.Node

func Init(startTime string, machineID int64) (err error) {
	// 获取指定时间的时间类型
	var st time.Time
	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		return
	}
	// 指定雪花算法的起始时间
	sf.Epoch = st.UnixNano() / 1000000
	// 根据machineID创造雪花id生成器
	node, err = sf.NewNode(machineID)
	return
}
func GenID() int64 {
	// 可以将生成的id转化为指定类型
	return node.Generate().Int64()
}
