package mysql

import (
	"bullmoon/models"
	"strings"

	"github.com/jmoiron/sqlx"
)

func CreatePost(p *models.Post) (err error) {
	sqlStr := "insert into post(post_id, title, content, author_id, community_id) values (?,?,?,?,?)"
	_, err = db.Exec(sqlStr, p.ID, p.Title, p.Content, p.AuthorID, p.CommunityID)
	return
}

func GetPostDetailByID(pid int64) (*models.Post, error) {
	data := new(models.Post)
	sqlStr := "select post_id, title, content, author_id, community_id, create_time from post where post_id=?"
	err := db.Get(data, sqlStr, pid)
	return data, err
}

func GetPostList(page int64, size int64) (posts []*models.Post, err error) {
	posts = make([]*models.Post, 0, size)
	// 分页查询 limit 1,3 会从第二行往后查三条
	sqlStr := "select post_id, title, content, author_id, community_id, create_time from post limit ?,?"
	err = db.Select(&posts, sqlStr, (page-1)*size, size)
	return
}

func GetPostListByID(ids []string) (posts []*models.Post, err error) {
	sqlStr := "select post_id, title, content, author_id, community_id, create_time from post " +
		"where post_id in (?) " +
		"order by FIND_IN_SET(post_id, ?)"
	quire, args, err := sqlx.In(sqlStr, ids, strings.Join(ids, ","))
	if err != nil {
		return nil, err
	}
	quire = db.Rebind(quire)
	err = db.Select(&posts, quire, args...)
	return
}
