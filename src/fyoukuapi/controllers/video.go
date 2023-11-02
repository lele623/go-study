package controllers

import (
	"fyoukuapi/models"
)

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
		page      int
		limit     int
	)

	channelId, _ = this.GetInt("channelId", 0)
	page, _ = this.GetInt("page", 0)
	limit, _ = this.GetInt("limit", 9)
	if channelId == 0 {
		this.JsonResult(1, "必须指定频道")
	}
	advert, res := models.GetChannelHotListById(channelId, page, limit)
	if !res {
		this.JsonResult(1, "没有相关内容")
	}
	this.JsonResult(0, "查询成功", advert)
}

// 按照地区获取推荐
func (this *VideoController) ChannelRegionRecommendList() {
	var (
		channelId int
		regionId  int
		page      int
		limit     int
	)

	channelId, _ = this.GetInt("channelId", 0)
	if channelId == 0 {
		this.JsonResult(1, "必须指定频道")
	}
	regionId, _ = this.GetInt("regionId", 0)
	if regionId == 0 {
		this.JsonResult(1, "必须指定频道地区")
	}
	page, _ = this.GetInt("page", 0)
	limit, _ = this.GetInt("limit", 9)

	video, res := models.GetChannelIdRegionRecommendList(channelId, regionId, page, limit)
	if !res {
		this.JsonResult(1, "没有相关内容")
	}
	this.JsonResult(0, "查询成功", video)
}

// 按照类型获取推荐
func (this *VideoController) ChannelTypeRecommendList() {
	var (
		channelId int
		typeId    int
		page      int
		limit     int
	)

	channelId, _ = this.GetInt("channelId", 0)
	if channelId == 0 {
		this.JsonResult(1, "必须指定频道")
	}
	typeId, _ = this.GetInt("typeId", 0)
	if typeId == 0 {
		this.JsonResult(1, "必须指定频道类型")
	}
	page, _ = this.GetInt("page", 0)
	limit, _ = this.GetInt("limit", 9)

	video, res := models.GetChannelTypeRecommendList(channelId, typeId, page, limit)
	if !res {
		this.JsonResult(1, "没有相关内容")
	}
	this.JsonResult(0, "查询成功", video)

}

// 视频列表
func (this *VideoController) ChannelVideo() {
	condition := make(map[string]interface{})
	condition["channelId"], _ = this.GetInt("channelId", 0)
	if condition["channelId"] == 0 {
		this.JsonResult(1, "必须指定频道")
	}
	condition["regionId"], _ = this.GetInt("regionId", 0)
	condition["typeId"], _ = this.GetInt("typeId", 0)
	condition["end"] = this.GetString("end", "")
	condition["sort"] = this.GetString("sort", "")
	condition["page"], _ = this.GetInt("page", 0)
	condition["limit"], _ = this.GetInt("limit", 12)

	video, res := models.GetChannelVideo(condition)
	if !res {
		this.JsonResult(1, "没有相关内容")
	}
	this.JsonResult(0, "查询成功", video)

}
