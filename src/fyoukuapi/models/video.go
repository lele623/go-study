package models

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
)

type Video struct {
	Id                 int
	Title              string
	SubTitle           string
	Status             int
	AddTime            int64
	Img                string
	Img1               string
	ChannelId          int
	TypeId             int
	RegionId           int
	UserId             int
	EpisodesCount      int
	EpisodesUpdateTime int
	IsEnd              int
	IsHot              int
	IsRecommend        int
	Comment            int
}

type VideoEpisodes struct {
	Id      int
	Title   string
	AddTime int64
	Num     int
	VideoId int
	PlayUrl string
	Status  int
	Comment int
}

// 是否热播
const (
	isHotOn  = 1 //是
	isHotOff = 0 //否
)

// 状态
const (
	statusOn  = 1 //是
	statusOff = 0 //否
)

// 通过频道ID获取热播列表
func GetChannelHotListById(channelId int, page int, limit int) ([]Video, bool) {
	var video []Video

	sql := "select * from video where channel_id = ? and is_hot = ? and status = ? order by episodes_update_time desc limit ?,?"
	num, _ := orm.NewOrm().Raw(sql, channelId, isHotOn, statusOn, page, limit).QueryRows(&video)
	if num == 0 {
		return video, false
	}
	return video, true
}

func GetChannelIdRegionRecommendList(channelId int, regionId int, page int, limit int) ([]Video, bool) {
	var video []Video

	sql := "select * from video where channel_id = ? and region_id = ? order by episodes_update_time desc limit ?,?"
	num, _ := orm.NewOrm().Raw(sql, channelId, regionId, page, limit).QueryRows(&video)
	if num == 0 {
		return video, false
	}
	return video, true
}

func GetChannelTypeRecommendList(channelId int, typeId int, page int, limit int) ([]Video, bool) {
	var video []Video

	sql := "select * from video where channel_id = ? and type_id = ? order by episodes_update_time desc limit ?,?"
	num, _ := orm.NewOrm().Raw(sql, channelId, typeId, page, limit).QueryRows(&video)
	if num == 0 {
		return video, false
	}
	return video, true
}

func GetChannelVideo(param map[string]interface{}) ([]Video, bool) {
	var video []Video

	sql := "SELECT * FROM video WHERE channel_id = ? AND status = 1"
	sqlArgs := []interface{}{param["channel_id"]}
	if param["type_id"] != 0 {
		sql += " AND type_id = ?"
		sqlArgs = append(sqlArgs, param["type_id"])
	}
	if param["region_id"] != 0 {
		sql += " AND region_id = ?"
		sqlArgs = append(sqlArgs, param["region_id"])
	}
	if param["end"] != "" {
		if param["end"] == "n" {
			sql += " AND is_end = 0"
		} else if param["end"] == "y" {
			sql += " AND is_end = 1"
		}
	}
	if param["sort"] != "" {
		sql += " ORDER BY " + fmt.Sprint(param["sort"]) + " DESC"
	} else {
		sql += " ORDER BY add_time DESC"
	}

	sql += " LIMIT ? OFFSET ?"
	sqlArgs = append(sqlArgs, param["limit"], param["page"])
	num, _ := orm.NewOrm().Raw(sql, sqlArgs).QueryRows(&video)
	if num == 0 {
		return video, false
	}
	return video, true
}

func GetVideoInfo(videoId int) ([]Video, bool) {
	var video []Video

	sql := "select * from video where id = ? limit 1"
	num, _ := orm.NewOrm().Raw(sql, videoId).QueryRows(&video)
	if num == 0 {
		return video, false
	}
	return video, true
}

func GetVideoEpisodesList(episodesId int) ([]VideoEpisodes, bool) {
	var videoEpisodes []VideoEpisodes

	sql := "select * from video_episodes where video_id = ? order by num"
	num, _ := orm.NewOrm().Raw(sql, episodesId).QueryRows(&videoEpisodes)
	if num == 0 {
		return videoEpisodes, false
	}
	return videoEpisodes, true
}
