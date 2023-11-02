package models

import (
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
} // 是否热播
// 状态
const (
	isHotOn  = 1 //是
	isHotOff = 0 //否
)
const (
	statusOn  = 1 //是
	statusOff = 0 //否
)

func init() {
	orm.RegisterModel(new(Video))
}

// 通过频道ID获取热播列表
func GetChannelHotListById(channelId int, page int, limit int) ([]Video, bool) {
	var (
		video []Video
	)

	sql := "select * from video where channel_id = ? and is_hot = ? and status = ? order by episodes_update_time desc limit ?,?"
	num, _ := orm.NewOrm().Raw(sql, channelId, isHotOn, statusOn, page, limit).QueryRows(&video)
	if num == 0 {
		return video, false
	}
	return video, true
}

func GetChannelIdRegionRecommendList(channelId int, regionId int, page int, limit int) ([]Video, bool) {
	var (
		video []Video
	)

	sql := "select * from video where channel_id = ? and region_id = ? order by episodes_update_time desc limit ?,?"
	num, _ := orm.NewOrm().Raw(sql, channelId, regionId, page, limit).QueryRows(&video)
	if num == 0 {
		return video, false
	}
	return video, true
}

func GetChannelTypeRecommendList(channelId int, typeId int, page int, limit int) ([]Video, bool) {
	var (
		video []Video
	)

	sql := "select * from video where channel_id = ? and type_id = ? order by episodes_update_time desc limit ?,?"
	num, _ := orm.NewOrm().Raw(sql, channelId, typeId, page, limit).QueryRows(&video)
	if num == 0 {
		return video, false
	}
	return video, true
}

func GetChannelVideo(condition map[string]interface{}) ([]Video, bool) {
	var (
		video []Video
		args  []interface{}
	)

	sql := "select * from video where channel_id = ? and status = 1"
	args = append(args, condition["channelId"])
	for i, v := range condition {
		if v == "" || v == 0 {
			continue
		}
		switch i {
		case "region_id":
			sql += " and region_id = ?"
			args = append(args, v)
		case "typeId":
			sql += " and type_id = ?"
			args = append(args, v)
		case "end":
			if condition["end"] == "n" {
				v = 0
			}
			if condition["end"] == "y" {
				v = 1
			}
			sql += " and is_end = ?"
			args = append(args, v)
		}
	}

	if condition["sort"] != "" {
		sql += " order by ? desc"
		args = append(args, condition["sort"])
	} else {
		sql += " order by add_time desc"
	}

	sql += " limit ?,?"
	args = append(args, condition["page"])
	args = append(args, condition["limit"])
	num, _ := orm.NewOrm().Raw(sql, args...).QueryRows(&video)
	if num == 0 {
		return video, false
	}
	return video, true
}
