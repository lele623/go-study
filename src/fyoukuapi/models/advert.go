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

func init() {
	orm.RegisterModel(new(Advert))
}

// 通过频道ID获取频道广告
func GetChannelAdvertById(channelId int) ([]Advert, bool) {
	var (
		advert []Advert
	)

	query := orm.NewOrm().QueryTable("advert")
	query = query.Filter("ChannelId", channelId)
	query = query.Filter("Status", advertStatusOn)
	err := query.One(&advert)
	if err != nil {
		return advert, false
	}
	return advert, true
}
