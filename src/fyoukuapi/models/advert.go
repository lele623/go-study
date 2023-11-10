package models

import (
	"errors"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
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
func GetChannelAdvertById(channelId int) ([]Advert, int64, error) {
	var advert []Advert

	sql := "select * from advert where channel_id = ? and status = ? limit 1"
	count, err := orm.NewOrm().Raw(sql, channelId, advertStatusOn).QueryRows(&advert)
	if err != nil {
		logs.Error(err)
		return advert, 0, errors.New("内部异常")
	}
	return advert, count, nil
}
