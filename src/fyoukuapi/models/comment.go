package models

import (
	"context"
	"errors"
	"fmt"
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

type CommentList struct {
	Stamp        int    `json:"stamp"`
	Avatar       string `json:"avatar"`
	Name         string `json:"name"`
	AddTime      int64
	Content      string `json:"content"`
	PraiseCount  int    `json:"praiseCount"`
	AddTimeTitle string `json:"addTimeTitle"`
}

// 根据剧集数获取评论列表
func GetCommentList(episodesId int, page int, limit int) ([]CommentList, int64, error) {
	var list []CommentList

	sql := "select c.stamp,c.add_time,c.content,c.praise_count,u.avatar,u.name from comment c left join user u on c.user_id=u.id where c.episodes_id = ? and c.status = ? order by c.add_time desc limit ?,?"
	count, err := orm.NewOrm().Raw(sql, episodesId, 1, page, limit).QueryRows(&list)
	if err != nil {
		logs.Error(err)
		return list, 0, errors.New("内部异常")
	}
	return list, count, nil
}

// 保存评论
func CommentSave(uid int, videoId int, episodesId int, content string) ([]Comment, error) {
	var comment []Comment

	//事务
	err := orm.NewOrm().DoTx(func(ctx context.Context, txOrm orm.TxOrmer) error {
		//保存评论
		saveSql := "insert into comment (content,user_id, episodes_id, video_id, stamp, status, add_time) values (?,?,?,?,?,?,?)"
		result, err := txOrm.Raw(saveSql, content, uid, episodesId, videoId, 0, 1, time.Now().Unix()).Exec()
		if err != nil {
			return fmt.Errorf("保存评论异常:%w", err)
		}
		rowsAffected, _ := result.RowsAffected()
		if rowsAffected == 0 {
			return errors.New("保存评论失败，评论数未发生变化")
		}

		//更新视频总评论数
		updateVideoSql := "update video set comment=comment+1 where id = ?"
		result, err = txOrm.Raw(updateVideoSql, videoId).Exec()
		if err != nil {
			return fmt.Errorf("更新视频总评论数异常:%w", err)
		}
		rowsAffected, _ = result.RowsAffected()
		if rowsAffected == 0 {
			return errors.New("更新视频总评论数失败，评论数未发生变化")
		}

		//更新视频剧集评论数
		updateEpisodesSql := "update video_episodes set comment=comment+1 where id = ?"
		result, err = txOrm.Raw(updateEpisodesSql, episodesId).Exec()
		if err != nil {
			return fmt.Errorf("更新视频剧集评论数异常:%w", err)
		}
		rowsAffected, _ = result.RowsAffected()
		if rowsAffected == 0 {
			return errors.New("更新视频剧集评论数失败，评论数未发生变化")
		}
		return nil
	})
	if err != nil {
		logs.Error(err)
		return comment, errors.New("内部异常")
	}
	return comment, nil
}
