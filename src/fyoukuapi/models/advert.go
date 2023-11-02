package models

import (
	"github.com/beego/beego/v2/client/orm"
)

type Advert struct {
	Id        int
	Title     string
	SubTitle  string
	ChannelId int
	Img       string
	Sort      string
	AddTime   int64
	Url       string
	Status    int
}

// 状态
const (
	advertStatusOn  = 1 //开启
	advertStatusOff = 0 //关闭
)

// 通过频道ID获取频道广告
func GetChannelAdvertById(channelId int) ([]Advert, bool) {
	var advert []Advert

	sql := "select * from advert where channel_id = ? and status = ? limit 1"
	num, _ := orm.NewOrm().Raw(sql, channelId, advertStatusOn).QueryRows(&advert)
	if num == 0 {
		return advert, false
	}
	return advert, true
}
