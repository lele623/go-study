package models

import (
	"errors"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

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

// 获取频道下的地区
func GetChannelRegion(channelId int) ([]ChannelRegion, int64, error) {
	var channelType []ChannelRegion

	sql := "select * from channel_region where channel_id = ? order by sort desc"
	count, err := orm.NewOrm().Raw(sql, channelId).QueryRows(&channelType)
	if err != nil {
		logs.Error(err)
		return channelType, 0, errors.New("内部异常")
	}
	return channelType, count, nil
}

// 获取频道下的类型
func GetChannelType(channelId int) ([]ChannelType, int64, error) {
	var channelType []ChannelType

	sql := "select * from channel_type where channel_id = ? order by sort desc"
	count, err := orm.NewOrm().Raw(sql, channelId).QueryRows(&channelType)
	if err != nil {
		logs.Error(err)
		return channelType, 0, errors.New("内部异常")
	}
	return channelType, count, nil
}
