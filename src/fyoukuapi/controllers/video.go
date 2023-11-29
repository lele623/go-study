package controllers

import (
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
	advert, count, err := models.GetChannelAdvertById(channelId)
	if err != nil {
		this.JsonResult(1, err.Error())
	}
	if count == 0 {
		this.JsonResult(1, "没有相关内容")
	}
	this.JsonResult(0, "操作成功", advert)
}

// 根据频道ID获取正在热播视频
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
	advert, count, err := models.GetChannelHotListById(channelId, page, limit)
	if err != nil {
		this.JsonResult(1, err.Error())
	}
	if count == 0 {
		this.JsonResult(1, "没有相关内容")
	}
	this.JsonResult(0, "操作成功", advert)
}

// 根据频道下的地区ID获取推荐视频
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

	video, count, err := models.GetChannelIdRegionRecommendList(channelId, regionId, page, limit)
	if err != nil {
		this.JsonResult(1, err.Error())
	}
	if count == 0 {
		this.JsonResult(1, "没有相关内容")
	}
	this.JsonResult(0, "操作成功", video)
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

	video, count, err := models.GetChannelTypeRecommendList(channelId, typeId, page, limit)
	if err != nil {
		this.JsonResult(1, err.Error())
	}
	if count == 0 {
		this.JsonResult(1, "没有相关内容")
	}
	this.JsonResult(0, "操作成功", video)
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

	video, count, err := models.GetChannelVideo(param)
	if err != nil {
		this.JsonResult(1, err.Error())
	}
	if count == 0 {
		this.JsonResult(1, "没有相关内容")
	}
	this.JsonResult(0, "操作成功", video)
}

// 视频详情
func (this *VideoController) VideoInfo() {
	var videoId int

	videoId, _ = this.GetInt("videoId", 0)
	if videoId == 0 {
		this.JsonResult(1, "必须指定视频ID")
	}
	video, count, err := models.GetVideoInfo(videoId)
	if err != nil {
		this.JsonResult(1, err.Error())
	}
	if count == 0 {
		this.JsonResult(1, "没有相关内容")
	}
	this.JsonResult(0, "操作成功", video)
}

// 视频剧集列表
func (this *VideoController) VideoEpisodesList() {
	var videoId int

	videoId, _ = this.GetInt("videoId", 0)
	if videoId == 0 {
		this.JsonResult(1, "必须指定视频ID")
	}
	videoEpisodes, count, err := models.GetVideoEpisodesList(videoId)
	if err != nil {
		this.JsonResult(1, err.Error())
	}
	if count == 0 {
		this.JsonResult(1, "没有相关内容")
	}
	this.JsonResult(0, "操作成功", videoEpisodes)
}

// 根据频道ID获取排行榜
func (this *VideoController) ChannelRanking() {
	var channelId int

	channelId, _ = this.GetInt("channelId", 0)
	if channelId == 0 {
		this.JsonResult(1, "必须指定频道")
	}
	video, _, err := models.GetChannelRanking(channelId)
	if err != nil {
		this.JsonResult(1, err.Error())
	}
	this.JsonResult(0, "操作成功", video)
}

// 根据类型ID获取排行榜
func (this *VideoController) TypeRanking() {
	var typeId int

	typeId, _ = this.GetInt("typeId", 0)
	if typeId == 0 {
		this.JsonResult(1, "必须指定类型")
	}
	video, _, err := models.GetTypeRanking(typeId)
	if err != nil {
		this.JsonResult(1, err.Error())
	}
	this.JsonResult(0, "操作成功", video)
}

// 保存视频
func (this *VideoController) VideoSave() {
	var (
		playUrl   string
		title     string
		subTitle  string
		uid       int
		channelId int
		typeId    int
		regionId  int
	)

	playUrl = this.GetString("playUrl", "")
	if playUrl == "" {
		this.JsonResult(1, "视频地址不能为空")
	}
	uid, _ = this.GetInt("uid", 0)
	if uid == 0 {
		this.JsonResult(1, "请先登录")
	}
	title = this.GetString("title", "")
	subTitle = this.GetString("subTitle", "")
	channelId, _ = this.GetInt("channelId", 0)
	typeId, _ = this.GetInt("typeId", 0)
	regionId, _ = this.GetInt("regionId", 0)

	err := models.VideoSave(playUrl, title, subTitle, uid, channelId, typeId, regionId)
	if err != nil {
		this.JsonResult(1, err.Error())
	}
	this.JsonResult(0, "操作成功")
}
