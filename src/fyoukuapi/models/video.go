package models

import (
	"errors"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
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
	count   int
	VideoId int
	PlayUrl string
	Status  int
	Comment int
}

// 是否热播
const (
	videoIsHotOn  = 1 //是
	videoIsHotOff = 1 //是
)

// 状态
const (
	videoStatusOn  = 1 //是
	videoStatusOff = 0 //否
)

// 根据频道ID获取正在热播视频
func GetChannelHotListById(channelId int, page int, limit int) ([]Video, int64, error) {
	var video []Video

	sql := "select * from video where channel_id = ? and is_hot = ? and status = ? order by episodes_update_time desc limit ?,?"
	count, err := orm.NewOrm().Raw(sql, channelId, videoIsHotOn, videoStatusOn, page, limit).QueryRows(&video)
	if err != nil {
		logs.Error(err)
		return video, 0, errors.New("内部异常")
	}
	return video, count, nil
}

// 按照地区获取推荐
func GetChannelIdRegionRecommendList(channelId int, regionId int, page int, limit int) ([]Video, int64, error) {
	var video []Video

	sql := "select * from video where channel_id = ? and region_id = ? order by episodes_update_time desc limit ?,?"
	count, err := orm.NewOrm().Raw(sql, channelId, regionId, page, limit).QueryRows(&video)
	if err != nil {
		logs.Error(err)
		return video, 0, errors.New("内部异常")
	}
	return video, count, nil
}

// 根据频道下类型ID获取推荐视频
func GetChannelTypeRecommendList(channelId int, typeId int, page int, limit int) ([]Video, int64, error) {
	var video []Video

	sql := "select * from video where channel_id = ? and type_id = ? order by episodes_update_time desc limit ?,?"
	count, err := orm.NewOrm().Raw(sql, channelId, typeId, page, limit).QueryRows(&video)
	if err != nil {
		logs.Error(err)
		return video, 0, errors.New("内部异常")
	}
	return video, count, nil
}

// 频道下根据不同条件和排序方式获取视频信息
func GetChannelVideo(param map[string]interface{}) ([]Video, int64, error) {
	var video []Video

	sql := "SELECT * FROM video WHERE channel_id = ? AND status = 1"
	args := []interface{}{param["channel_id"]}
	if param["type_id"] != 0 {
		sql += " AND type_id = ?"
		args = append(args, param["type_id"])
	}
	if param["region_id"] != 0 {
		sql += " AND region_id = ?"
		args = append(args, param["region_id"])
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
	args = append(args, param["limit"], param["page"])
	count, err := orm.NewOrm().Raw(sql, args).QueryRows(&video)
	if err != nil {
		logs.Error(err)
		return video, 0, errors.New("内部异常")
	}
	return video, count, nil
}

// 根据视频ID获取视频信息
func GetVideoInfo(videoId int) ([]Video, int64, error) {
	var video []Video

	sql := "select * from video where id = ? limit 1"
	count, err := orm.NewOrm().Raw(sql, videoId).QueryRows(&video)
	if err != nil {
		logs.Error(err)
		return video, 0, errors.New("内部异常")
	}
	return video, count, nil
}

// 根据视频ID获取剧集列表
func GetVideoEpisodesList(episodesId int) ([]VideoEpisodes, int64, error) {
	var videoEpisodes []VideoEpisodes

	sql := "select * from video_episodes where video_id = ? order by num"
	count, err := orm.NewOrm().Raw(sql, episodesId).QueryRows(&videoEpisodes)
	if err != nil {
		logs.Error(err)
		return videoEpisodes, 0, errors.New("内部异常")
	}
	return videoEpisodes, count, nil
}
