package controllers

import (
	"fmt"
	"fyoukuapi/models"
)

type VideoController struct {
	BaseController
}

// 获取频道顶部广告
func (this *VideoController) ChannelAdvert() {
	var channelId int

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
	param := make(map[string]interface{})
	param["channel_id"], _ = this.GetInt("channelId", 0)
	if param["channel_id"] == 0 {
		this.JsonResult(1, "必须指定频道")
	}
	param["region_id"], _ = this.GetInt("regionId", 0)
	param["type_id"], _ = this.GetInt("typeId", 0)
	param["end"] = this.GetString("end", "")
	param["sort"] = this.GetString("sort", "")
	param["page"], _ = this.GetInt("page", 0)
	param["limit"], _ = this.GetInt("limit", 2)

	fmt.Println(param)
	video, res := models.GetChannelVideo(param)
	if !res {
		this.JsonResult(1, "没有相关内容")
	}
	this.JsonResult(0, "查询成功", video)
}

// 视频详情
func (this *VideoController) VideoInfo() {
	var videoId int

	videoId, _ = this.GetInt("videoId", 0)
	if videoId == 0 {
		this.JsonResult(1, "必须指定视频ID")
	}
	video, res := models.GetVideoInfo(videoId)
	if !res {
		this.JsonResult(1, "没有相关内容")
	}
	this.JsonResult(0, "查询成功", video)
}

// 视频剧集列表
func (this *VideoController) VideoEpisodesList() {
	var videoId int

	videoId, _ = this.GetInt("videoId", 0)
	if videoId == 0 {
		this.JsonResult(1, "必须指定视频ID")
	}
	videoEpisodes, res := models.GetVideoEpisodesList(videoId)
	if !res {
		this.JsonResult(1, "没有相关内容")
	}
	this.JsonResult(0, "查询成功", videoEpisodes)
}
