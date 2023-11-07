package models

import (
	"errors"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"time"
)

type Comment struct {
	Id          int
	Content     string
	AddTime     int64
	UserId      int
	Status      int
	Stamp       int
	PraiseCount int
	EpisodesId  int
	VideoId     int
}

// 根据剧集数获取评论列表
func GetCommentList(episodesId int, page int, limit int) ([]Comment, int64) {
	var comment []Comment

	sql := "select * from comment where episodes_id = ? and status = ? order by add_time desc limit ?,?"
	num, _ := orm.NewOrm().Raw(sql, episodesId, 1, page, limit).QueryRows(&comment)
	return comment, num
}

// 保存评论
func GetCommentSave(uid int, videoId int, episodesId int, content string) ([]Comment, error) {
	var comment []Comment

	sql := "insert into comment (content,user_id, episodes_id, video_id, stamp, status, add_time) values (?,?,?,?,?,?,?)"
	_, err := orm.NewOrm().Raw(sql, content, uid, episodesId, videoId, 0, 1, time.Now()).Exec()
	if err != nil {
		logs.Error(err)
		return comment, errors.New("内部异常")
	}
	return comment, nil
}
