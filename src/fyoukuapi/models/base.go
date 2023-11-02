package models

import "github.com/beego/beego/v2/client/orm"

type ChannelRegion struct {
	Id        int
	Name      string
	Status    string
	AddTime   int64
	ChannelId int
	Sort      int
}
type ChannelType struct {
	Id        int
	Name      string
	Status    string
	AddTime   int64
	ChannelId int
	Sort      int
}

// 获取地区
func GetChannelRegion(channelId int) ([]ChannelRegion, bool) {
	var channelType []ChannelRegion

	sql := "select * from channel_region where channel_id = ? order by sort desc"
	num, _ := orm.NewOrm().Raw(sql, channelId).QueryRows(&channelType)
	if num == 0 {
		return channelType, false
	}
	return channelType, true
}

// 获取类型
func GetChannelType(channelId int) ([]ChannelType, bool) {
	var channelType []ChannelType

	sql := "select * from channel_type where channel_id = ? order by sort desc"
	num, _ := orm.NewOrm().Raw(sql, channelId).QueryRows(&channelType)
	if num == 0 {
		return channelType, false
	}
	return channelType, true
}
