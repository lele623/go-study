package controllers

import "fyoukuapi/models"

type VideoController struct {
	BaseController
}

// 获取频道顶部广告
func (this *VideoController) ChannelAdvert() {
	var (
		channelId int
	)
	channelId, _ = this.GetInt("channelId", 0)
	if channelId == 0 {
		this.JsonResult(1, "必须指定频道")
	}
	advert, res := models.GetChannelAdvertById(channelId)
	if !res {
		this.JsonResult(1, "没有相关内容")
	}
	this.JsonResult(0, "查询成功", advert)
}

// 获取热播列表
func (this *VideoController) ChannelHotList() {
	var (
		channelId int
		limit     int
	)
	channelId, _ = this.GetInt("channelId", 0)
	channelId, _ = this.GetInt("limit", 9)
	if channelId == 0 {
		this.JsonResult(1, "必须指定频道")
	}
	advert, res := models.GetChannelHotListById(channelId, limit)
	if !res {
		this.JsonResult(1, "没有相关内容")
	}
	this.JsonResult(0, "查询成功", advert)

}
