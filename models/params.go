package models

const (
	OrderTime  = "time"
	OrderScore = "score"
)

type ParamSignUp struct {
	UserName   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

type ParamLogin struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ParamVotedData struct {
	// userid从上下文中获取
	PostID    string `json:"post_id" binding:"required"`              // 帖子id
	Direction int8   `json:"direction,string" binding:"oneof=1 0 -1"` // 赞成票（1）还是反对票（-1）或者是取消投票（0）
}

type ParamPostList struct {
	Page  int64  `json:"page" form:"page"`   // 页码
	Size  int64  `json:"size" form:"size"`   // 每页数量
	Order string `json:"order" form:"order"` // 按照时间还是分数排序
}
