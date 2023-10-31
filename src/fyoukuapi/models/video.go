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
	img                string
	img1               string
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
func GetChannelHotListById(channelId int, limit int) ([]Video, bool) {
	var (
		video []Video
	)

	query := orm.NewOrm().QueryTable("video")
	query = query.Filter("ChannelId", channelId)
	query = query.Filter("IsHot", isHotOn)
	query = query.Filter("Status", statusOn)
	query = query.Limit(limit)
	_, err := query.All(&video)
	if err != nil {
		return video, false
	}
	return video, true
}
