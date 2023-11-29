package models

import (
	"errors"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"time"
)

type Barrage struct {
	Id          int
	Content     string
	AddTime     int64
	UserId      int
	Status      int
	CurrentTime int
	EpisodesId  int
	VideoId     int
}

type BarrageData struct {
	Id          int    `json:"id"`
	Content     string `json:"content"`
	CurrentTime int    `json:"currentTime"`
}

// 弹幕列表
func BarrageList(episodesId int, startTime int, endTime int) (int64, []BarrageData, error) {
	var barrages []BarrageData

	sql := "select `id`,`content`,`current_time` from barrage  where `status` = 1 and `episodes_id`=? and `current_time` >= ? and `current_time` < ? order by `current_time`"
	num, err := orm.NewOrm().Raw(sql, episodesId, startTime, endTime).QueryRows(&barrages)
	if err != nil {
		logs.Error(err)
		return 0, barrages, errors.New("内部异常")
	}
	return num, barrages, err
}

// 保存弹幕
func Savebarrage(episodesId int, videoId int, uid int, content string, currentTime int) error {

	_, err := orm.NewOrm().Raw("insert into barrage (content,`current_time`,user_id,episodes_id,video_id,status,add_time) values (?,?,?,?,?,?,?)", content, currentTime, uid, episodesId, videoId, 1, time.Now().Unix()).Exec()
	if err != nil {
		logs.Error(fmt.Errorf("保存用户信息异常:%w", err))
		return errors.New("内部异常")
	}
	return nil
}
